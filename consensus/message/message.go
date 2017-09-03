package transport


type MessageType uint32;

const (

	MsgPreAccept MessageType = 0
	MsgPreAcceptOK MessageType  = 1

)

type Message struct {

	msg PbMessage
	msgType MessageType

}

type PbMessage interface {
	Size() (n int)
	Marshal() (data []byte, err error)
	Unmarshal(data []byte) error
}