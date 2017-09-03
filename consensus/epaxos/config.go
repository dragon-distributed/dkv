package consensus

type (
	TypeConsensusId int32

)


type ConsensusConfig struct {

	consensusId TypeConsensusId
	instance  Instance

	peers     []string
	bindAddr  string

	peerStore PeerStore
	logStore  LogStore

}