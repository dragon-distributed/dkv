package epaxos

import "github.com/dragon-distributed/dkv/epaxos/epaxospb"

type Store interface {
	Save(instances []*epaxospb.Instance)

	Update(instancesWithoutCmds []*epaxospb.Instance)
}

type DummyStore struct {

}

func (d *DummyStore) Save(instances []*epaxospb.Instance) {

}

func (d *DummyStore) Update(instancesWithoutCmds []*epaxospb.Instance) {

}
