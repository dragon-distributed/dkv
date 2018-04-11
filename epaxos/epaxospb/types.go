package epaxospb

type TypeNodeID uint64
type TypeKey string
type TypeInstanceID uint64

type TypeNodeIDSlice []TypeNodeID

func (p TypeNodeIDSlice) Len() int           { return len(p) }
func (p TypeNodeIDSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p TypeNodeIDSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type TypeInstanceIDSlice []TypeInstanceID

func (p TypeInstanceIDSlice) Len() int           { return len(p) }
func (p TypeInstanceIDSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p TypeInstanceIDSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
