package epaxos

import "github.com/dragon-distributed/dkv/epaxos/epaxospb"

const (
	ProgressStateActive   = 0
	ProgressStateUnActive = 1
)

type ProgressStateType uint64

var prstmap = [...]string{
	"ProgressStateActive",
	"ProgressStateUnActive",
}

func (st ProgressStateType) String() string { return prstmap[uint64(st)] }

type Progress struct {
	ID epaxospb.TypeNodeID

	State ProgressStateType

	ProcessingMap map[epaxospb.TypeInstanceID]*Instance
}

func (p *Progress) IsPaused() bool {
	if p.State == ProgressStateActive {
		return true
	} else {
		return false
	}
}
