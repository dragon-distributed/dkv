package epaxosnew

import (
	"github.com/dragon-distributed/dkv/epaxosnew/pbnew"
)

type TransportCallback interface {
	Receive(message *pbnew.Message) error
}

type Transport interface {
	Send(messages []*pbnew.Message)

	Start() error

	AddPeer(addr string, nodeID uint64) error
}

//
//type TcpTransport struct {
//	callback    TransportCallback
//	Addr        string
//	Ln          *net.TCPListener
//	peers       map[epaxospb.TypeNodeID]*peer
//	peersRWLock sync.RWMutex
//	ShutdownCh  chan struct{}
//}
//
//type peer struct {
//	addr string
//	conn net.Conn
//}
//
//func NewTcpTransport(callback TransportCallback, listenAddr string) *TcpTransport {
//	t := &TcpTransport{
//		callback: callback,
//		Addr:     listenAddr,
//		peers:    make(map[epaxospb.TypeNodeID]*peer),
//	}
//	return t
//}
//
//func (t *TcpTransport) AddPeer(addr string, nodeID epaxospb.TypeNodeID) error {
//
//	t.peersRWLock.Lock()
//	t.peers[nodeID] = &peer{addr: addr}
//	t.peersRWLock.Unlock()
//	//go t.handleRead(conn)
//
//	return nil
//}
//
//func (t *TcpTransport) Send(messages []*epaxospb.Message) {
//
//	t.peersRWLock.RLock()
//	defer t.peersRWLock.RUnlock()
//	headerBytes := make([]byte, 4)
//	for _, msg := range messages {
//		if t.peers[msg.To] == nil {
//			fmt.Errorf("message need send to %v cannot find in transport peers \n ", msg.To)
//			continue
//		}
//		writedBody := 0
//		writedHeader := 0
//		size := msg.Size()
//		bodyBytes, err := msg.Marshal()
//		if err != nil {
//			panic(err)
//		}
//
//		binary.BigEndian.PutUint32(headerBytes, uint32(size))
//
//		for writedHeader != 4 {
//			n, err := t.peers[msg.To].conn.Write(headerBytes[writedHeader:])
//			if err != nil {
//				panic(err)
//			}
//			writedHeader = writedHeader + n
//		}
//		for writedBody != msg.Size() {
//			n, err := t.peers[msg.To].conn.Write(bodyBytes[writedBody:])
//			if err != nil {
//				panic(err)
//			}
//			writedBody = writedBody + n
//		}
//	}
//}
//
//func (t *TcpTransport) Start() error {
//
//	addr, err := net.ResolveTCPAddr("tcp", t.Addr)
//	if err != nil {
//		fmt.Printf("[kafka_server.go-Start]:ResolveTCPAddr error,s.Addr=%s,error=%s \n ", t.Addr, err.Error())
//		return err
//	}
//
//	ln, err := net.ListenTCP("tcp", addr)
//	if err != nil {
//		fmt.Printf("[kafka_server.go-Start]:ListenTCP error,error=%s \n", err.Error())
//		return err
//	}
//	t.Ln = ln
//	fmt.Printf("transport start \n")
//	go t.handleConnectAction()
//	go t.connectPeerLoop()
//
//	return nil
//}
//
//func (t *TcpTransport) connectPeerLoop() {
//	for {
//		timeout := time.After(time.Second * 3)
//		<-timeout
//		t.peersRWLock.RLock()
//		for _, v := range t.peers {
//			if v.conn == nil {
//				conn, err := net.Dial("tcp", v.addr)
//				if err != nil {
//					fmt.Printf("cannot connect to %v \n", v.addr)
//				} else {
//					v.conn = conn
//					go t.handleRead(conn)
//					fmt.Printf("connect to %v success \n", v.addr)
//				}
//			}
//		}
//
//		t.peersRWLock.RUnlock()
//	}
//}
//
//func (t *TcpTransport) handleConnectAction() {
//	for {
//		select {
//		case <-t.ShutdownCh:
//			fmt.Printf("[kafka_server.go-Start]: server closed \n")
//			return
//		default:
//			conn, err := t.Ln.Accept()
//			if err != nil {
//				fmt.Printf("[kafka_server.go-Start]:listener accept failed: %v \n", err)
//				continue
//			}
//			go t.handleRead(conn)
//		}
//	}
//}
//
//func (t *TcpTransport) handleRead(conn net.Conn) {
//	defer func() {
//		conn.Close()
//	}()
//
//	p := make([]byte, 4)
//	for {
//		_, err := io.ReadFull(conn, p[:])
//		if err != nil {
//			panic(err)
//		}
//		size := binary.BigEndian.Uint32(p)
//
//		b := make([]byte, size)
//		_, err = io.ReadFull(conn, b)
//		if err != nil {
//			panic(err)
//			break
//		}
//
//		msg := &epaxospb.Message{}
//		err = msg.Unmarshal(b)
//		if err != nil {
//			panic(err)
//			break
//		}
//		t.callback.Receive(msg)
//	}
//}
