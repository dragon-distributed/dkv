package main

import (
	"fmt"
	"github.com/dragon-distributed/dkv/epaxos"
	"github.com/dragon-distributed/dkv/epaxos/epaxospb"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	args := os.Args
	if len(args) != 2 {
		panic("need 2 args")
	}
	var n *epaxos.Node

	if args[1] == "1" {
		conf := &epaxos.NodeConfig{
			PeerListenAddress: "127.0.0.1:3001",
			InitialPeers:      "1:127.0.0.1:3001,2:127.0.0.1:3002,3:127.0.0.1:3003",
			NodeID:            epaxospb.TypeNodeID(1),
		}
		n = epaxos.NewNode(conf, &epaxos.DummyStore{})
		go n.Start()
	} else if args[1] == "2" {
		conf := &epaxos.NodeConfig{
			PeerListenAddress: "127.0.0.1:3002",
			InitialPeers:      "1:127.0.0.1:3001,2:127.0.0.1:3002,3:127.0.0.1:3003",
			NodeID:            epaxospb.TypeNodeID(2),
		}
		n = epaxos.NewNode(conf, &epaxos.DummyStore{})
		go n.Start()
	} else if args[1] == "3" {
		conf := &epaxos.NodeConfig{
			PeerListenAddress: "127.0.0.1:3003",
			InitialPeers:      "1:127.0.0.1:3001,2:127.0.0.1:3002,3:127.0.0.1:3003",
			NodeID:            epaxospb.TypeNodeID(3),
		}
		n = epaxos.NewNode(conf, &epaxos.DummyStore{})
		go n.Start()

	} else {
		panic("need 1,2,3")
	}

	timeout := time.After(time.Second * 10)
	<-timeout

	go func() {
		for {
			timeout := time.After(time.Second * 1)
			<-timeout
			fmt.Printf("============slowpath:%v, fastpath:%v\n", epaxos.SlowPathCount, epaxos.FastPathCount)
		}

	}()

	count := 1
	for {
		rand.Seed(time.Now().UnixNano())
		//timeout := time.After(time.Millisecond * time.Duration(rand.Int31n(10)+10))
		//<-timeout
		if rand.Int31n(100) <= 10 {
			cmd := &epaxospb.Cmd{
				Key:   "key1",
				Value: []byte("value1"),
				Type:  epaxospb.CmdType_CmdPut,
			}
			n.Propose(cmd)
		} else {
			cmd := &epaxospb.Cmd{
				Key:   "key" + strconv.Itoa(int(rand.Int31n(100000))),
				Value: []byte("value1"),
				Type:  epaxospb.CmdType_CmdPut,
			}
			n.Propose(cmd)
		}

		//fmt.Println("example propose a cmd")

		count++
		if count >= 50000 {
			break
		}
	}

	serve := make(chan struct{})
	<-serve

}
