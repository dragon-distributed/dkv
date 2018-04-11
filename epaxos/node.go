package epaxos

import (
	"github.com/dragon-distributed/dkv/epaxos/epaxospb"
	"strconv"
	"strings"
)

var (
	batchProposeSize = 64
)

// epaxos node
type Node struct {
	receiveMsg chan *epaxospb.Message

	proposeMsg chan *epaxospb.Cmd

	transport Transport

	EP *EPaxos

	config *NodeConfig
}

type NodeConfig struct {

	// listen address for peer
	PeerListenAddress string

	// initial peers split by comma
	// [nodeID]:[listenAddress]
	// eg: 0:127.0.0.1:3000,1:127.0.0.1:3001
	InitialPeers string

	// this node id
	NodeID epaxospb.TypeNodeID
}

func NewNode(config *NodeConfig, store Store) *Node {

	node := &Node{
		config:     config,
		receiveMsg: make(chan *epaxospb.Message),
		proposeMsg: make(chan *epaxospb.Cmd, batchProposeSize),
	}

	node.transport = NewTcpTransport(&TransportReceivedCallback{node: node}, config.PeerListenAddress)
	node.EP = NewEPaxos(store, config.NodeID, node.transport)

	return node
}

func (n *Node) Propose(cmd *epaxospb.Cmd) {
	n.proposeMsg <- cmd
}

func (n *Node) Start() {

	n.transport.Start()
	go n.EP.Executor.Start()

	for _, peerStr := range strings.Split(n.config.InitialPeers, ",") {
		temp := strings.Split(peerStr, ":")
		nodeID, err := strconv.Atoi(temp[0])
		if err != nil {
			panic(err)
		}
		peerAddress := temp[1] + ":" + temp[2]

		if epaxospb.TypeNodeID(nodeID) != n.EP.NodeID {
			n.transport.AddPeer(peerAddress, epaxospb.TypeNodeID(nodeID))
		}
		n.EP.addEPaxosPeer(epaxospb.TypeNodeID(nodeID))
	}

	n.run()
}

func (n *Node) run() {

	var proposeChan chan *epaxospb.Cmd
	for {
		if len(n.EP.Executor.addExecuteChan) > 400 {
			proposeChan = nil
		} else {
			proposeChan = n.proposeMsg
		}

		select {
		case msg := <-n.receiveMsg:
			if msg.Type == epaxospb.MessageType_MsgPreAccept {
				n.EP.OnPreAccept(msg)
			} else if msg.Type == epaxospb.MessageType_MsgPreAccpetResp {
				n.EP.OnPreAcceptResp(msg)
			} else if msg.Type == epaxospb.MessageType_MsgAccept {
				n.EP.OnAccept(msg)
			} else if msg.Type == epaxospb.MessageType_MsgAcceptResp {
				n.EP.OnAcceptResp(msg)
			} else if msg.Type == epaxospb.MessageType_MsgCommit {
				n.EP.OnCommit(msg)
			} else {
				panic("unknow msg type: " + epaxospb.MessageType_name[int32(msg.Type)])
			}
		case cmd := <-proposeChan:
			cmds := make([]*epaxospb.Cmd, 0, batchProposeSize)
			cmds = append(cmds, cmd)
			for i := 1; i < batchProposeSize; i++ {
				select {
				case batchCmd := <-n.proposeMsg:
					cmds = append(cmds, batchCmd)
				default:
					i = batchProposeSize
				}
			}
			n.EP.Propose(cmds)
		}
	}
}

type TransportReceivedCallback struct {
	node *Node
}

func (t *TransportReceivedCallback) Receive(message *epaxospb.Message) error {
	t.node.receiveMsg <- message
	return nil
}
