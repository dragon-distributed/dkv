package epaxos

import (
	"fmt"
	"github.com/dragon-distributed/dkv/epaxos/epaxospb"
	"sync"
)

var (
	FastPathCount = 0
	SlowPathCount = 0
)

type EPaxos struct {
	NodeID epaxospb.TypeNodeID

	// TODO. cmds snapshot clean
	Cmds map[epaxospb.TypeNodeID]map[string]*epaxospb.Instance

	// the max instance id for each node
	MaxInstanceID epaxospb.TypeInstanceID

	// waiting send messages
	//sendMsgs []*epaxospb.Message

	transport Transport

	progress       map[epaxospb.TypeNodeID]*Progress
	progressRWLock sync.RWMutex

	// store for instance
	store Store

	ballot map[epaxospb.TypeNodeID]epaxospb.Ballot

	// store processing instance
	//processingInstances map[epaxospb.TypeNodeID]map[epaxospb.TypeInstanceID]*Instance

	Executor *Executor
}

func NewEPaxos(store Store, nodeID epaxospb.TypeNodeID, transport Transport) *EPaxos {

	ep := &EPaxos{
		store:     store,
		NodeID:    nodeID,
		Cmds:      make(map[epaxospb.TypeNodeID]map[string]*epaxospb.Instance),
		progress:  make(map[epaxospb.TypeNodeID]*Progress),
		ballot:    make(map[epaxospb.TypeNodeID]epaxospb.Ballot),
		transport: transport,
	}

	ep.Cmds[nodeID] = make(map[string]*epaxospb.Instance)
	ep.ballot[nodeID] = epaxospb.Ballot{Epoch: 1, NodeId: uint64(nodeID)}
	ep.Executor = NewExecutor()

	return ep
}

func (ep *EPaxos) addEPaxosPeer(ID epaxospb.TypeNodeID) {

	ep.progress[ID] = &Progress{
		ID:            ID,
		State:         ProgressStateActive,
		ProcessingMap: make(map[epaxospb.TypeInstanceID]*Instance),
	}
	ep.Cmds[ID] = make(map[string]*epaxospb.Instance)
}

func (ep *EPaxos) Propose(cmds []*epaxospb.Cmd) {

	instances := make([]*epaxospb.Instance, 0, len(cmds))
	for _, cmd := range cmds {

		ep.MaxInstanceID = ep.MaxInstanceID + 1

		localDeps := ep.scanConflict(cmd)
		deps := make(map[epaxospb.TypeNodeID]epaxospb.TypeInstanceID)
		var maxSeq uint64 = 0
		for nodeID, instance := range localDeps {
			deps[nodeID] = instance.ID
			if instance.Seq > maxSeq {
				maxSeq = instance.Seq
			}
		}
		maxSeq = maxSeq + 1

		fmt.Printf("Propose epaxos maxInstanceID :%v,localDeps :%v, deps:%v,maxSeq:%v \n",
			ep.MaxInstanceID, localDeps, deps, maxSeq)

		instance := &epaxospb.Instance{
			ID:     ep.MaxInstanceID,
			Cmd:    cmd,
			Deps:   deps,
			Seq:    maxSeq,
			NodeID: ep.NodeID,
		}
		instances = append(instances, instance)

		ep.updateAllProcessingInstance(instances)

		// update local cmds
		cloneInstance := ep.cloneInstanceWithoutCmd(instance)
		cloneInstance.State = epaxospb.InstanceState_PreAccepted
		ep.Cmds[ep.NodeID][cmd.Key] = cloneInstance
	}

	message := &epaxospb.Message{
		Type:     epaxospb.MessageType_MsgPreAccept,
		Instance: instances,
		From:     ep.NodeID,
		Ballot:   ep.ballot[ep.NodeID],
	}

	// like raft , write to its disk in parallel with replicating to other nodes
	ep.broadcast(message)

	ep.store.Save(instances)

}

func (ep *EPaxos) OnPreAccept(msg *epaxospb.Message) {

	fmt.Printf("epaxos OnPreAccept, msg: %v\n", msg)

	if !ep.isBallotMatch(msg) {
		ep.respWhenBallotDismatch(msg)
		return
	}

	// reply without cmd to reduce network cost
	acceptInstancesWithoutCmd := make([]*epaxospb.Instance, 0)
	rejectInstancesWithoutCmd := make([]*epaxospb.Instance, 0)

	for _, requestInstance := range msg.Instance {

		// save to mem first
		ep.processingInstances[requestInstance.NodeID][requestInstance.ID] = &Instance{
			Instance: requestInstance,
			State:    epaxospb.InstanceState_PreAccept,
		}

		isAccept := true

		// scan local conflict before update
		localDeps := ep.scanConflict(requestInstance.Cmd)
		var localMaxSeq uint64 = 0
		for _, instance := range localDeps {
			if instance.Seq > localMaxSeq {
				localMaxSeq = instance.Seq
			}
		}

		//update local cmds, dependencies first
		instanceWithoutCmd := ep.cloneInstanceWithoutCmd(requestInstance)
		ep.Cmds[requestInstance.NodeID][requestInstance.Cmd.Key] = instanceWithoutCmd

		// update resp when deps or seq not equal
		if localMaxSeq+1 > requestInstance.Seq {
			isAccept = false
			requestInstance.Seq = localMaxSeq + 1
		}

		if !ep.isDepsEqual(localDeps, requestInstance.Deps) {
			isAccept = false
			if requestInstance.Deps == nil {
				requestInstance.Deps = make(map[epaxospb.TypeNodeID]epaxospb.TypeInstanceID)
			}
			ep.mergeDeps(localDeps, requestInstance.Deps)
		}

		if isAccept {
			acceptInstancesWithoutCmd = append(acceptInstancesWithoutCmd, instanceWithoutCmd)
		} else {
			rejectInstancesWithoutCmd = append(rejectInstancesWithoutCmd, instanceWithoutCmd)
		}
	}

	ep.store.Save(msg.Instance)

	// reply accept instance
	if len(acceptInstancesWithoutCmd) > 0 {
		acceptMessage := &epaxospb.Message{
			Type:     epaxospb.MessageType_MsgPreAccpetResp,
			Instance: acceptInstancesWithoutCmd,
			From:     ep.NodeID,
			Ballot:   ep.ballot[ep.NodeID],
			Reject:   false,
			To:       msg.From,
		}
		ep.send(acceptMessage)
	}

	// reply reject instances
	if len(rejectInstancesWithoutCmd) > 0 {
		rejectMessage := &epaxospb.Message{
			Type:       epaxospb.MessageType_MsgPreAccpetResp,
			Instance:   rejectInstancesWithoutCmd,
			From:       ep.NodeID,
			Ballot:     ep.ballot[ep.NodeID],
			Reject:     true,
			RejectType: epaxospb.RespRejectType_PreAcceptReject,
			To:         msg.From,
		}
		ep.send(rejectMessage)
	}
}

func (ep *EPaxos) OnPreAcceptResp(msg *epaxospb.Message) {

	fmt.Printf("epaxos OnPreAcceptResp, msg: %v\n", msg)

	if !ep.isBallotMatch(msg) {
		ep.respWhenBallotDismatch(msg)
		return
	}

	// update processing instances and dispatch instance which reach quorum
	commitPhaseInstancesWithoutCmd := make([]*epaxospb.Instance, 0)
	commitPhaseInstances := make([]*Instance, 0)
	acceptPhaseInstancesWithoutCmd := make([]*epaxospb.Instance, 0)

	for _, respInstance := range msg.Instance {

		localProcessingInstance := ep.processingInstances[respInstance.NodeID][respInstance.ID]

		if localProcessingInstance == nil {
			fmt.Println("cannot found local instance when OnPreAcceptResp, ID:", respInstance.ID)
			continue
		}

		if localProcessingInstance.State != epaxospb.InstanceState_PreAccept {
			fmt.Println(respInstance.ID, " not in PreAccept state, it in ",
				localProcessingInstance.State.String())
			continue
		}

		localProcessingInstance.PreAcceptCount = localProcessingInstance.PreAcceptCount + 1

		if msg.Reject && msg.RejectType == epaxospb.RespRejectType_PreAcceptReject {
			localProcessingInstance.IsPreAcceptReject = true
			if localProcessingInstance.Seq < respInstance.Seq {
				localProcessingInstance.Seq = respInstance.Seq
			}
			ep.mergeDepsWithID(respInstance.Deps, localProcessingInstance.Deps)
		}

		if ep.isQuorum(localProcessingInstance.PreAcceptCount + 1) {

			if localProcessingInstance.IsPreAcceptReject {
				localProcessingInstance.State = epaxospb.InstanceState_Accept
				acceptPhaseInstancesWithoutCmd = append(acceptPhaseInstancesWithoutCmd,
					ep.cloneInstanceWithoutCmd(localProcessingInstance.Instance))
				SlowPathCount++
			} else {
				localProcessingInstance.State = epaxospb.InstanceState_Commit
				commitPhaseInstancesWithoutCmd = append(commitPhaseInstancesWithoutCmd,
					ep.cloneInstanceWithoutCmd(localProcessingInstance.Instance))
				commitPhaseInstances = append(commitPhaseInstances, localProcessingInstance)
				delete(ep.processingInstances[ep.NodeID], respInstance.ID)
				FastPathCount++
			}

		}
	}

	//commit phase handle
	if len(commitPhaseInstancesWithoutCmd) > 0 {
		message := &epaxospb.Message{
			Type:     epaxospb.MessageType_MsgCommit,
			Instance: commitPhaseInstancesWithoutCmd,
			From:     ep.NodeID,
			Ballot:   ep.ballot[ep.NodeID],
		}
		ep.broadcast(message)

		// save after broadcast
		ep.store.Update(commitPhaseInstancesWithoutCmd)

		// put to execute
		for _, toExecute := range commitPhaseInstances {
			ep.Executor.PutInstanceToExecute(toExecute.NodeID, toExecute)
		}

	}

	//accept phase handle
	if len(acceptPhaseInstancesWithoutCmd) > 0 {
		message := &epaxospb.Message{
			Type:     epaxospb.MessageType_MsgAccept,
			Instance: acceptPhaseInstancesWithoutCmd,
			From:     ep.NodeID,
			Ballot:   ep.ballot[ep.NodeID],
		}

		ep.broadcast(message)

		// save after broadcast
		ep.store.Update(acceptPhaseInstancesWithoutCmd)
	}
}

func (ep *EPaxos) OnAccept(msg *epaxospb.Message) {

	fmt.Printf("epaxos OnAccept, msg: %v\n", msg)

	if !ep.isBallotMatch(msg) {
		ep.respWhenBallotDismatch(msg)
		return
	}

	// on accept msg without cmds
	for _, requestInstance := range msg.Instance {

		localProcessingInstance := ep.processingInstances[requestInstance.NodeID][requestInstance.ID]
		if localProcessingInstance == nil {
			fmt.Println("cannot found local instance when OnPreAcceptResp, ID:", requestInstance.ID)
			continue
		}

		//accept phase update local cmds, dependencies
		ep.Cmds[localProcessingInstance.NodeID][localProcessingInstance.Cmd.Key] = requestInstance
		localProcessingInstance.Deps = requestInstance.Deps
		localProcessingInstance.Seq = requestInstance.Seq

	}

	ep.store.Update(msg.Instance)

	message := &epaxospb.Message{
		Type:     epaxospb.MessageType_MsgAcceptResp,
		Instance: msg.Instance,
		From:     ep.NodeID,
		Ballot:   ep.ballot[ep.NodeID],
		To:       msg.From,
	}

	ep.send(message)
}

func (ep *EPaxos) OnAcceptResp(msg *epaxospb.Message) {

	fmt.Printf("epaxos OnAcceptResp, msg: %v\n", msg)

	if !ep.isBallotMatch(msg) {
		ep.respWhenBallotDismatch(msg)
		return
	}

	commitPhaseInstancesWithoutCmd := make([]*epaxospb.Instance, 0)
	commitPhaseInstances := make([]*Instance, 0)

	for _, respInstance := range msg.Instance {

		localProcessingInstance := ep.processingInstances[ep.NodeID][respInstance.ID]

		if localProcessingInstance == nil {
			fmt.Println("cannot found local instance when OnAcceptResp, ID:", respInstance.ID)
			continue
		}

		if localProcessingInstance.State != epaxospb.InstanceState_Accept {
			fmt.Println(respInstance.ID, " not in Accept state, it in ",
				localProcessingInstance.State.String())
			continue
		}

		localProcessingInstance.AcceptCount = localProcessingInstance.AcceptCount + 1

		if ep.isQuorum(localProcessingInstance.AcceptCount + 1) {
			localProcessingInstance.State = epaxospb.InstanceState_Commit
			commitPhaseInstancesWithoutCmd = append(commitPhaseInstancesWithoutCmd, respInstance)
			commitPhaseInstances = append(commitPhaseInstances, localProcessingInstance)
			delete(ep.processingInstances[respInstance.NodeID], respInstance.ID)
		}
	}

	//commit phase handle
	if len(commitPhaseInstancesWithoutCmd) > 0 {
		message := &epaxospb.Message{
			Type:     epaxospb.MessageType_MsgCommit,
			Instance: commitPhaseInstancesWithoutCmd,
			From:     ep.NodeID,
			Ballot:   ep.ballot[ep.NodeID],
		}
		ep.broadcast(message)

		// save after broadcast
		ep.store.Update(commitPhaseInstancesWithoutCmd)

		// put to execute
		for _, toExecute := range commitPhaseInstances {
			ep.Executor.PutInstanceToExecute(toExecute.NodeID, toExecute)
		}
	}

}

func (ep *EPaxos) OnCommit(msg *epaxospb.Message) {

	fmt.Printf("epaxos OnCommit, msg: %v\n", msg)

	if !ep.isBallotMatch(msg) {
		ep.respWhenBallotDismatch(msg)
		return
	}

	// on accept msg without cmds
	for _, requestInstance := range msg.Instance {

		localProcessingInstance := ep.processingInstances[requestInstance.NodeID][requestInstance.ID]
		if localProcessingInstance == nil {
			fmt.Println("cannot found local instance when OnCommit, ID:", requestInstance.ID)
			panic("cannot found local instance when OnCommit")
			continue
		}

		//accept phase update local cmds, dependencies
		ep.Cmds[localProcessingInstance.NodeID][localProcessingInstance.Cmd.Key] = requestInstance
		localProcessingInstance.Deps = requestInstance.Deps
		localProcessingInstance.Seq = requestInstance.Seq
	}

	ep.store.Update(msg.Instance)

	// put to execute
	for _, toExecute := range msg.Instance {
		ep.Executor.PutInstanceToExecute(toExecute.NodeID,
			ep.processingInstances[toExecute.NodeID][toExecute.ID])
		delete(ep.processingInstances[toExecute.NodeID], toExecute.ID)
	}

}

func (ep *EPaxos) OnBallotDismatch(msg *epaxospb.Message) {
	// TODO.
	panic("ballot dismatch")
}

func (ep *EPaxos) scanConflict(cmd *epaxospb.Cmd) map[epaxospb.TypeNodeID]*epaxospb.Instance {

	conflictMap := make(map[epaxospb.TypeNodeID]*epaxospb.Instance)
	for k, v := range ep.Cmds {
		if v[cmd.Key] != nil {
			conflictMap[k] = v[cmd.Key]
		}
	}
	return conflictMap
}

func (ep *EPaxos) broadcast(msg *epaxospb.Message) {

	if msg.To > 0 {
		panic("broadcast msg should not include msg.To")
	}

	ep.progressRWLock.RLock()
	defer ep.progressRWLock.RUnlock()

	//fmt.Printf("broadcast %v , len(ep.progress):%v\n", msg, len(ep.progress))
	for _, p := range ep.progress {
		msg.To = p.ID
		ep.send(msg)
	}
}

func (ep *EPaxos) send(msg *epaxospb.Message) {
	ep.transport.Send([]*epaxospb.Message{msg})
}

func (ep *EPaxos) cloneInstanceWithoutCmd(instance *epaxospb.Instance) *epaxospb.Instance {

	cloneMap := make(map[epaxospb.TypeNodeID]epaxospb.TypeInstanceID)

	for key, value := range instance.Deps {
		cloneMap[key] = value
	}

	return &epaxospb.Instance{
		ID:     instance.ID,
		Deps:   cloneMap,
		Seq:    instance.Seq,
		NodeID: instance.NodeID,
	}
}

func (ep *EPaxos) isBallotMatch(requestMsg *epaxospb.Message) bool {

	//TODO.
	return true
}

func (ep *EPaxos) respWhenBallotDismatch(requestMsg *epaxospb.Message) {

	rejectMessage := &epaxospb.Message{
		Type:       epaxospb.MessageType_MsgBallotDismatch,
		From:       ep.NodeID,
		Ballot:     ep.ballot[ep.NodeID],
		Reject:     true,
		RejectType: epaxospb.RespRejectType_BallotNotMatch,
		To:         requestMsg.From,
	}
	ep.send(rejectMessage)
}

func (ep *EPaxos) isDepsEqual(first map[epaxospb.TypeNodeID]*epaxospb.Instance,
	second map[epaxospb.TypeNodeID]epaxospb.TypeInstanceID) bool {

	if len(first) != len(second) {
		return false
	}

	for k, v := range first {
		secondInstanceID := second[k]
		if secondInstanceID == 0 {
			return false
		}
		if secondInstanceID != v.ID {
			return false
		}
	}
	return true
}

func (ep *EPaxos) mergeDeps(from map[epaxospb.TypeNodeID]*epaxospb.Instance,
	to map[epaxospb.TypeNodeID]epaxospb.TypeInstanceID) {

	for k, v := range from {
		if to[k] == 0 || to[k] < v.ID {
			to[k] = v.ID
		}
	}
}

func (ep *EPaxos) mergeDepsWithID(from map[epaxospb.TypeNodeID]epaxospb.TypeInstanceID,
	to map[epaxospb.TypeNodeID]epaxospb.TypeInstanceID) {

	for k, v := range from {
		if to[k] == 0 || to[k] < v {
			to[k] = v
		}
	}
}

func (ep *EPaxos) isQuorum(count int) bool {
	if count >= len(ep.progress)/2+1 {
		return true
	} else {
		return false
	}
}

func (ep *EPaxos) updateAllProcessingInstance(instance *epaxospb.Instance) {
	for _, p := range ep.progress {
		p.ProcessingMap[instance.ID] = instance
	}
}
