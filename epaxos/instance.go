package epaxos

import "github.com/dragon-distributed/dkv/epaxos/epaxospb"

type Instance struct {
	*epaxospb.Instance

	State epaxospb.InstanceState

	PreAcceptCount int

	IsPreAcceptReject bool

	AcceptCount int

	isExecuted bool
}

type InstanceSlice []*Instance

func (p InstanceSlice) Len() int { return len(p) }
func (p InstanceSlice) Less(i, j int) bool {
	if p[i].Seq < p[j].Seq {
		return true
	}
	if p[i].Seq == p[j].Seq {
		if p[i].NodeID < p[j].NodeID {
			return true
		}
	}
	return false
}
func (p InstanceSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
