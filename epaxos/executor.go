package epaxos

import (
	"errors"
	"fmt"
	"github.com/dragon-distributed/dkv/epaxos/epaxospb"
	"sort"
	"sync"
)

var (
	depNodeNotFound = errors.New("depNodeNotFound")
)

type Executor struct {
	ToExecuteInstances map[epaxospb.TypeNodeID]map[epaxospb.TypeInstanceID]*Instance

	// the min executed instance id
	MinExecutedInstance map[epaxospb.TypeNodeID]epaxospb.TypeInstanceID

	mapLock sync.RWMutex

	executeNotifyChan chan struct{}

	addExecuteChan chan *Instance
}

func NewExecutor() *Executor {

	e := &Executor{
		executeNotifyChan: make(chan struct{}, 1),
	}

	e.ToExecuteInstances = make(map[epaxospb.TypeNodeID]map[epaxospb.TypeInstanceID]*Instance)
	e.MinExecutedInstance = make(map[epaxospb.TypeNodeID]epaxospb.TypeInstanceID)
	e.addExecuteChan = make(chan *Instance, 50000000)

	return e
}

func (e *Executor) Start() {

	for {
		select {
		case addExec := <-e.addExecuteChan:

			if e.ToExecuteInstances[addExec.NodeID] == nil {
				e.ToExecuteInstances[addExec.NodeID] = make(map[epaxospb.TypeInstanceID]*Instance)
			}
			e.ToExecuteInstances[addExec.NodeID][addExec.ID] = addExec
			select {
			case e.executeNotifyChan <- struct{}{}:
			default:
			}
		case <-e.executeNotifyChan:

			for {
				canContinueExecute := make([]bool, len(e.ToExecuteInstances))
				for i := 0; i < len(e.ToExecuteInstances); i++ {
					canContinueExecute[i] = true
				}

				var scc [][]*Instance = nil

				index := -1
				for nodeID, instanceMap := range e.ToExecuteInstances {

					index++

					executedInstanceID := e.MinExecutedInstance[nodeID]

					// need execute orderly
					if instanceMap[executedInstanceID+1] == nil {
						canContinueExecute[index] = false
						continue // for next nodeID
					}

					// scan next instance to execute
					if instanceMap[executedInstanceID+1].isExecuted {
						for {
							delete(instanceMap, executedInstanceID+1)
							// consider next
							executedInstanceID = executedInstanceID + 1
							e.MinExecutedInstance[nodeID] = executedInstanceID
							if instanceMap[executedInstanceID+1] == nil {
								canContinueExecute[index] = false
								break
							}
							if !instanceMap[executedInstanceID+1].isExecuted {
								break
							}
						}
						if canContinueExecute[index] == false {
							continue // for next nodeID
						}
					}

					fmt.Printf("minExecutedInstance map, map:%v, \n", e.MinExecutedInstance)
					//fmt.Printf("toExecuteInstances map, map:%v, \n", e.toExecuteInstances)

					var err error
					err, scc = e.getSCC(instanceMap[executedInstanceID+1])
					if err != nil {
						fmt.Printf("============ err: %v \n", err)
						canContinueExecute[index] = false
						continue // for next nodeID
					}

					// sort scc
					for _, v := range scc {
						sort.Sort(InstanceSlice(v))
					}

					// execute scc
					e.executeSCC(scc)
				}

				canConinue := false
				for i := 0; i < len(e.ToExecuteInstances); i++ {
					if canContinueExecute[i] == true {
						canConinue = true
						break
					}
				}
				if !canConinue {
					break
				}

			}
		}
	}
}

type tNode struct {
	dfn       int
	low       int
	instance  *Instance
	isInStack bool
}

func (e *Executor) executeSCC(scc [][]*Instance) {
	for _, v := range scc {
		for _, instance := range v {
			fmt.Printf("execute instance %v \n", instance)
			instance.isExecuted = true
		}
	}
}

func (e *Executor) getSCC(instance *Instance) (error, [][]*Instance) {

	//fmt.Println("===getSCC :%v", instance)

	startIndex := 1
	scc := make([][]*Instance, 0)
	var err error
	visitMap := make(map[epaxospb.TypeNodeID]map[epaxospb.TypeInstanceID]*tNode)

	scc, err = e.visit(instance, &startIndex, &Stack{}, scc, visitMap)
	if err != nil {
		return err, nil
	}
	return nil, scc
}

func (e *Executor) visit(instance *Instance, index *int, stack *Stack,
	scc [][]*Instance, visitMap map[epaxospb.TypeNodeID]map[epaxospb.TypeInstanceID]*tNode) ([][]*Instance, error) {

	fmt.Printf("visit instance %v\n", instance)
	node := &tNode{
		dfn:       *index,
		low:       *index,
		instance:  instance,
		isInStack: true,
	}
	stack.Push(node)
	if visitMap[instance.NodeID] == nil {
		visitMap[instance.NodeID] = make(map[epaxospb.TypeInstanceID]*tNode)
	}
	visitMap[instance.NodeID][instance.ID] = node
	*index++

	// order iterator map
	var depNodeIDKeys []epaxospb.TypeNodeID
	for k := range node.instance.Deps {
		depNodeIDKeys = append(depNodeIDKeys, k)
	}
	sort.Sort(epaxospb.TypeNodeIDSlice(depNodeIDKeys))

	for _, depNodeID := range depNodeIDKeys {

		fmt.Printf("for visit sub node, use dep Node ID %v\n", depNodeID)
		needDepsInstanceID := node.instance.Deps[depNodeID]
		if needDepsInstanceID <= e.MinExecutedInstance[depNodeID] {
			continue
		}

		// find deps instance in execute map
		var depInstanceInExecute *Instance = e.getToExecuteInstance(depNodeID, needDepsInstanceID)
		if depInstanceInExecute == nil {
			//return nil, depNodeNotFound
			errStr := fmt.Sprintf("depNodeNotFound, instanceID:%d from node:%d , needDepsInstanceID is %d in nodeID %d",
				instance.ID, instance.NodeID, needDepsInstanceID, depNodeID)
			return nil, errors.New(errStr)
		}

		// is visit before
		if !e.isVisitBefore(depInstanceInExecute, visitMap) {
			var err error
			scc, err = e.visit(depInstanceInExecute, index, stack, scc, visitMap)
			if err != nil {
				return nil, err
			}
			depTNode := getTNodeFromVisitMap(depInstanceInExecute.NodeID, depInstanceInExecute.ID, visitMap)
			if depTNode.low < node.low {
				node.low = depTNode.low
			}
		} else {
			depTNode := getTNodeFromVisitMap(depInstanceInExecute.NodeID, depInstanceInExecute.ID, visitMap)

			if depTNode.isInStack {
				if depTNode.dfn < node.low {
					node.low = depTNode.dfn
				}
			}
		}
	}

	if node.dfn == node.low {
		sc := make([]*Instance, 0)
		for {
			stackTNode := stack.Pop()
			stackTNode.isInStack = false
			sc = append(sc, stackTNode.instance)

			if stackTNode.instance.NodeID == node.instance.NodeID &&
				stackTNode.instance.ID == node.instance.ID {
				break
			}
		}
		scc = append(scc, sc)
	}
	return scc, nil

}

func (e *Executor) getToExecuteInstance(nodeID epaxospb.TypeNodeID, instanceID epaxospb.TypeInstanceID) *Instance {

	var executeInstance *Instance = nil
	executeNodeMap := e.ToExecuteInstances[nodeID]
	if executeNodeMap != nil {
		executeInstance = executeNodeMap[instanceID]
	}
	return executeInstance
}

func (e *Executor) isVisitBefore(instance *Instance, visitMap map[epaxospb.TypeNodeID]map[epaxospb.TypeInstanceID]*tNode) bool {

	if getTNodeFromVisitMap(instance.NodeID, instance.ID, visitMap) != nil {
		return true
	}
	return false
}

func getTNodeFromVisitMap(nodeID epaxospb.TypeNodeID, instanceID epaxospb.TypeInstanceID,
	visitMap map[epaxospb.TypeNodeID]map[epaxospb.TypeInstanceID]*tNode) *tNode {

	if visitMap[nodeID] != nil {
		if visitMap[nodeID][instanceID] != nil {
			return visitMap[nodeID][instanceID]
		}
	}

	return nil
}

func (e *Executor) PutInstanceToExecute(nodeID epaxospb.TypeNodeID, instance *Instance) {

	//e.mapLock.Lock()
	//defer e.mapLock.Unlock()
	//
	//fmt.Printf("PutInstanceToExecute, %v, nodeID:%v \n", instance, nodeID)
	//if e.toExecuteInstances[nodeID] == nil {
	//	e.toExecuteInstances[nodeID] = make(map[epaxospb.TypeInstanceID]*Instance)
	//}
	//e.toExecuteInstances[nodeID][instance.ID] = instance
	//select {
	//case e.executeNotifyChan <- struct{}{}:
	//default:
	//}
	e.addExecuteChan <- instance
}

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value *tNode
	next  *Element
}

// Return the stack's length
func (s *Stack) Len() int {
	return s.size
}

// Push a new element onto the stack
func (s *Stack) Push(value *tNode) {
	s.top = &Element{value, s.top}
	s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *Stack) Pop() (value *tNode) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}
