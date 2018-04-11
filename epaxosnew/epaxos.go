package epaxosnew

import "github.com/dragon-distributed/dkv/epaxosnew/pbnew"

type EPaxos struct {
	transport Transport

	myLeader *MyProcess

	cmds *Cmds

	nodeID uint64

	progress map[uint64]*Progress
}

type MyProcess struct {

	// the max instance id for each node
	maxInstanceID uint64
}

type Progress struct {
	nodeID uint64
}

func (ep *EPaxos) Propose(cmd *pbnew.Cmd) {

	if cmd.Type != pbnew.CmdType_CmdPut {
		panic("cmd type should put")
	}

	ep.myLeader.maxInstanceID = ep.myLeader.maxInstanceID + 1

	deps := ep.cmds.containKey(cmd.Key)
	var maxSeq uint64 = 0
	for _, instance := range deps {
		if instance.Seq > maxSeq {
			maxSeq = instance.Seq
		}
	}
	maxSeq = maxSeq + 1

	instance := &pbnew.Instance{
		ID:     ep.myLeader.maxInstanceID,
		Cmd:    cmd,
		Deps:   deps,
		Seq:    maxSeq,
		NodeID: ep.nodeID,
	}

	cloneInstance := ep.cloneInstanceWithoutCmdValue(instance)
	cloneInstance.State = pbnew.InstanceState_PreAccepted
	ep.cmds.updateLocalCmd(cloneInstance)


	message := &pbnew.Message{
		Type:     pbnew.MessageType_MsgPreAccept,
		Instance: instances,
		From:     ep.NodeID,
		Ballot:   ep.ballot[ep.NodeID],
	}

}

func (ep *EPaxos) cloneInstanceWithoutCmdValue(instance *pbnew.Instance) *pbnew.Instance {

	cloneMap := make(map[uint64]uint64)

	for key, value := range instance.Deps {
		cloneMap[key] = value
	}

	cmd := &pbnew.Cmd{
		Key:       instance.Cmd.Key,
		Type:      instance.Cmd.Type,
		RequestID: instance.Cmd.RequestID,
	}

	return &pbnew.Instance{
		ID:     instance.ID,
		Deps:   cloneMap,
		Seq:    instance.Seq,
		NodeID: instance.NodeID,
		Cmd:    cmd,
	}
}
