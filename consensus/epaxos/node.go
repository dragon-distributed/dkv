package consensus

import "golang.org/x/net/context"

type ConsensusNode struct {

	conf *ConsensusConfig
	transport Transport


}

func NewConsensusNode(conf *ConsensusConfig) (*ConsensusNode,error){

	consensusNode := &ConsensusNode{
		conf : conf,
	}

	return consensusNode,nil

}


func (c *ConsensusNode) Start() (err error){


	//start listener


	//start loop/handler


}

func (c *ConsensusNode) Propose(data []byte, ctx context.Context){

}


