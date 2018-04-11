package epaxos

import (
	"fmt"
	"github.com/dragon-distributed/dkv/epaxos/epaxospb"
	"sort"
	"testing"
)

// [nodeID,instanceID]
// A[100,1] -> B[101,1]
// B[101,1] -> A[100,1]
func TestSCCSimpleLoop(t *testing.T) {

	instanceA := createInstance(100, 1, nil, 1)
	instanceB := createInstance(101, 1, nil, 1)

	instanceA.Instance.Deps[101] = 1
	instanceB.Instance.Deps[100] = 1

	executor := &Executor{
		ToExecuteInstances:  make(map[epaxospb.TypeNodeID]map[epaxospb.TypeInstanceID]*Instance),
		MinExecutedInstance: make(map[epaxospb.TypeNodeID]epaxospb.TypeInstanceID),
		executeNotifyChan:   make(chan struct{}),
	}

	executor.PutInstanceToExecute(instanceA.NodeID, instanceA)
	executor.PutInstanceToExecute(instanceB.NodeID, instanceB)

	err, scc := executor.getSCC(instanceA)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	for _, v := range scc {
		sort.Sort(InstanceSlice(v))
	}

	fmt.Printf("================= finish find scc\n")
	for _, v := range scc {
		fmt.Printf("++++++++++++++++\n")
		for _, instance := range v {
			fmt.Printf("execute instance, nodeID %v, ID %v, full information %v \n", instance.NodeID, instance.ID, instance)
			instance.isExecuted = true
		}
	}
}

func TestSCC(t *testing.T) {

	instance6 := createInstance(6, 1, nil, 1)
	instance5 := createInstance(5, 1, []*Instance{instance6}, 1)
	instance4 := createInstance(4, 1, []*Instance{instance6}, 2)

	instance3 := createInstance(3, 1, []*Instance{instance4, instance5}, 2)
	instance2 := createInstance(2, 1, []*Instance{instance4}, 2)

	instance1 := createInstance(1, 1, []*Instance{instance2, instance3}, 4)
	instance4.Deps[1] = 1

	executor := &Executor{
		ToExecuteInstances:  make(map[epaxospb.TypeNodeID]map[epaxospb.TypeInstanceID]*Instance),
		MinExecutedInstance: make(map[epaxospb.TypeNodeID]epaxospb.TypeInstanceID),
		executeNotifyChan:   make(chan struct{}),
	}

	executor.PutInstanceToExecute(instance1.NodeID, instance1)
	executor.PutInstanceToExecute(instance2.NodeID, instance2)
	executor.PutInstanceToExecute(instance3.NodeID, instance3)
	executor.PutInstanceToExecute(instance4.NodeID, instance4)
	executor.PutInstanceToExecute(instance5.NodeID, instance5)
	executor.PutInstanceToExecute(instance6.NodeID, instance6)

	err, scc := executor.getSCC(instance1)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	for _, v := range scc {
		sort.Sort(InstanceSlice(v))
	}

	fmt.Printf("================= finish find scc\n")
	for _, v := range scc {
		fmt.Printf("++++++++++++++++\n")
		for _, instance := range v {
			fmt.Printf("execute instance, nodeID %v, ID %v, full information %v \n", instance.NodeID, instance.ID, instance)
			instance.isExecuted = true
		}
	}
}

func createInstance(nodeID epaxospb.TypeNodeID, instanceID epaxospb.TypeInstanceID,
	depsInstance []*Instance, seq uint64) *Instance {

	pbInstance := &epaxospb.Instance{
		NodeID: nodeID,
		ID:     instanceID,
		Deps:   make(map[epaxospb.TypeNodeID]epaxospb.TypeInstanceID),
		Seq:    seq,
	}

	for _, deps := range depsInstance {
		pbInstance.Deps[deps.NodeID] = deps.ID
	}

	instance := &Instance{
		Instance:   pbInstance,
		isExecuted: false,
	}

	return instance

}
