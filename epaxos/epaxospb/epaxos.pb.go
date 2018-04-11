// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: epaxos.proto

/*
	Package epaxospb is a generated protocol buffer package.

	It is generated from these files:
		epaxos.proto

	It has these top-level messages:
		Ballot
		Message
		Instance
		Cmd
*/
package epaxospb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MessageType int32

const (
	MessageType_MsgPreAccept      MessageType = 0
	MessageType_MsgPreAccpetResp  MessageType = 1
	MessageType_MsgAccept         MessageType = 2
	MessageType_MsgAcceptResp     MessageType = 3
	MessageType_MsgBallotDismatch MessageType = 4
	MessageType_MsgCommit         MessageType = 5
)

var MessageType_name = map[int32]string{
	0: "MsgPreAccept",
	1: "MsgPreAccpetResp",
	2: "MsgAccept",
	3: "MsgAcceptResp",
	4: "MsgBallotDismatch",
	5: "MsgCommit",
}
var MessageType_value = map[string]int32{
	"MsgPreAccept":      0,
	"MsgPreAccpetResp":  1,
	"MsgAccept":         2,
	"MsgAcceptResp":     3,
	"MsgBallotDismatch": 4,
	"MsgCommit":         5,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptorEpaxos, []int{0} }

type InstanceState int32

const (
	InstanceState_PreAccepted InstanceState = 0
	InstanceState_Accepted    InstanceState = 1
	InstanceState_Committed   InstanceState = 2
)

var InstanceState_name = map[int32]string{
	0: "PreAccepted",
	1: "Accepted",
	2: "Committed",
}
var InstanceState_value = map[string]int32{
	"PreAccepted": 0,
	"Accepted":    1,
	"Committed":   2,
}

func (x InstanceState) String() string {
	return proto.EnumName(InstanceState_name, int32(x))
}
func (InstanceState) EnumDescriptor() ([]byte, []int) { return fileDescriptorEpaxos, []int{1} }

type RespRejectType int32

const (
	RespRejectType_BallotNotMatch  RespRejectType = 0
	RespRejectType_PreAcceptReject RespRejectType = 1
)

var RespRejectType_name = map[int32]string{
	0: "BallotNotMatch",
	1: "PreAcceptReject",
}
var RespRejectType_value = map[string]int32{
	"BallotNotMatch":  0,
	"PreAcceptReject": 1,
}

func (x RespRejectType) String() string {
	return proto.EnumName(RespRejectType_name, int32(x))
}
func (RespRejectType) EnumDescriptor() ([]byte, []int) { return fileDescriptorEpaxos, []int{2} }

type CmdType int32

const (
	CmdType_CmdGet  CmdType = 0
	CmdType_CmdPut  CmdType = 1
	CmdType_CmdQGet CmdType = 3
)

var CmdType_name = map[int32]string{
	0: "CmdGet",
	1: "CmdPut",
	3: "CmdQGet",
}
var CmdType_value = map[string]int32{
	"CmdGet":  0,
	"CmdPut":  1,
	"CmdQGet": 3,
}

func (x CmdType) String() string {
	return proto.EnumName(CmdType_name, int32(x))
}
func (CmdType) EnumDescriptor() ([]byte, []int) { return fileDescriptorEpaxos, []int{3} }

type Ballot struct {
	Epoch  uint64 `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	NodeId uint64 `protobuf:"varint,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (m *Ballot) Reset()                    { *m = Ballot{} }
func (m *Ballot) String() string            { return proto.CompactTextString(m) }
func (*Ballot) ProtoMessage()               {}
func (*Ballot) Descriptor() ([]byte, []int) { return fileDescriptorEpaxos, []int{0} }

func (m *Ballot) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *Ballot) GetNodeId() uint64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

type Message struct {
	Type       MessageType    `protobuf:"varint,1,opt,name=type,proto3,enum=epaxospb.MessageType" json:"type,omitempty"`
	Instance   []*Instance    `protobuf:"bytes,2,rep,name=instance" json:"instance,omitempty"`
	To         TypeNodeID     `protobuf:"varint,3,opt,name=to,proto3,casttype=TypeNodeID" json:"to,omitempty"`
	From       TypeNodeID     `protobuf:"varint,4,opt,name=from,proto3,casttype=TypeNodeID" json:"from,omitempty"`
	Ballot     Ballot         `protobuf:"bytes,5,opt,name=ballot" json:"ballot"`
	Reject     bool           `protobuf:"varint,6,opt,name=reject,proto3" json:"reject,omitempty"`
	RejectType RespRejectType `protobuf:"varint,7,opt,name=rejectType,proto3,enum=epaxospb.RespRejectType" json:"rejectType,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptorEpaxos, []int{1} }

func (m *Message) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_MsgPreAccept
}

func (m *Message) GetInstance() []*Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *Message) GetTo() TypeNodeID {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *Message) GetFrom() TypeNodeID {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *Message) GetBallot() Ballot {
	if m != nil {
		return m.Ballot
	}
	return Ballot{}
}

func (m *Message) GetReject() bool {
	if m != nil {
		return m.Reject
	}
	return false
}

func (m *Message) GetRejectType() RespRejectType {
	if m != nil {
		return m.RejectType
	}
	return RespRejectType_BallotNotMatch
}

type Instance struct {
	ID     TypeInstanceID                `protobuf:"varint,1,opt,name=ID,proto3,casttype=TypeInstanceID" json:"ID,omitempty"`
	Cmd    *Cmd                          `protobuf:"bytes,2,opt,name=cmd" json:"cmd,omitempty"`
	Deps   map[TypeNodeID]TypeInstanceID `protobuf:"bytes,3,rep,name=deps,castkey=TypeNodeID,castvalue=TypeInstanceID" json:"deps,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Seq    uint64                        `protobuf:"varint,4,opt,name=seq,proto3" json:"seq,omitempty"`
	NodeID TypeNodeID                    `protobuf:"varint,5,opt,name=nodeID,proto3,casttype=TypeNodeID" json:"nodeID,omitempty"`
	State  InstanceState                 `protobuf:"varint,6,opt,name=state,proto3,enum=epaxospb.InstanceState" json:"state,omitempty"`
}

func (m *Instance) Reset()                    { *m = Instance{} }
func (m *Instance) String() string            { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()               {}
func (*Instance) Descriptor() ([]byte, []int) { return fileDescriptorEpaxos, []int{2} }

func (m *Instance) GetID() TypeInstanceID {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Instance) GetCmd() *Cmd {
	if m != nil {
		return m.Cmd
	}
	return nil
}

func (m *Instance) GetDeps() map[TypeNodeID]TypeInstanceID {
	if m != nil {
		return m.Deps
	}
	return nil
}

func (m *Instance) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Instance) GetNodeID() TypeNodeID {
	if m != nil {
		return m.NodeID
	}
	return 0
}

func (m *Instance) GetState() InstanceState {
	if m != nil {
		return m.State
	}
	return InstanceState_PreAccepted
}

type Cmd struct {
	Key          string  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value        []byte  `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Type         CmdType `protobuf:"varint,3,opt,name=type,proto3,enum=epaxospb.CmdType" json:"type,omitempty"`
	RequestID    uint64  `protobuf:"varint,4,opt,name=requestID,proto3" json:"requestID,omitempty"`
	ConnectionID uint64  `protobuf:"varint,5,opt,name=connectionID,proto3" json:"connectionID,omitempty"`
}

func (m *Cmd) Reset()                    { *m = Cmd{} }
func (m *Cmd) String() string            { return proto.CompactTextString(m) }
func (*Cmd) ProtoMessage()               {}
func (*Cmd) Descriptor() ([]byte, []int) { return fileDescriptorEpaxos, []int{3} }

func (m *Cmd) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Cmd) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Cmd) GetType() CmdType {
	if m != nil {
		return m.Type
	}
	return CmdType_CmdGet
}

func (m *Cmd) GetRequestID() uint64 {
	if m != nil {
		return m.RequestID
	}
	return 0
}

func (m *Cmd) GetConnectionID() uint64 {
	if m != nil {
		return m.ConnectionID
	}
	return 0
}

func init() {
	proto.RegisterType((*Ballot)(nil), "epaxospb.Ballot")
	proto.RegisterType((*Message)(nil), "epaxospb.Message")
	proto.RegisterType((*Instance)(nil), "epaxospb.Instance")
	proto.RegisterType((*Cmd)(nil), "epaxospb.Cmd")
	proto.RegisterEnum("epaxospb.MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("epaxospb.InstanceState", InstanceState_name, InstanceState_value)
	proto.RegisterEnum("epaxospb.RespRejectType", RespRejectType_name, RespRejectType_value)
	proto.RegisterEnum("epaxospb.CmdType", CmdType_name, CmdType_value)
}
func (m *Ballot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ballot) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Epoch != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(m.Epoch))
	}
	if m.NodeId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(m.NodeId))
	}
	return i, nil
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(m.Type))
	}
	if len(m.Instance) > 0 {
		for _, msg := range m.Instance {
			dAtA[i] = 0x12
			i++
			i = encodeVarintEpaxos(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.To != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(m.To))
	}
	if m.From != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(m.From))
	}
	dAtA[i] = 0x2a
	i++
	i = encodeVarintEpaxos(dAtA, i, uint64(m.Ballot.Size()))
	n1, err := m.Ballot.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.Reject {
		dAtA[i] = 0x30
		i++
		if m.Reject {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.RejectType != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(m.RejectType))
	}
	return i, nil
}

func (m *Instance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Instance) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(m.ID))
	}
	if m.Cmd != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(m.Cmd.Size()))
		n2, err := m.Cmd.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Deps) > 0 {
		for k, _ := range m.Deps {
			dAtA[i] = 0x1a
			i++
			v := m.Deps[k]
			mapSize := 1 + sovEpaxos(uint64(k)) + 1 + sovEpaxos(uint64(v))
			i = encodeVarintEpaxos(dAtA, i, uint64(mapSize))
			dAtA[i] = 0x8
			i++
			i = encodeVarintEpaxos(dAtA, i, uint64(k))
			dAtA[i] = 0x10
			i++
			i = encodeVarintEpaxos(dAtA, i, uint64(v))
		}
	}
	if m.Seq != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(m.Seq))
	}
	if m.NodeID != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(m.NodeID))
	}
	if m.State != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(m.State))
	}
	return i, nil
}

func (m *Cmd) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Cmd) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	if m.Type != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(m.Type))
	}
	if m.RequestID != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(m.RequestID))
	}
	if m.ConnectionID != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(m.ConnectionID))
	}
	return i, nil
}

func encodeFixed64Epaxos(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Epaxos(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintEpaxos(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Ballot) Size() (n int) {
	var l int
	_ = l
	if m.Epoch != 0 {
		n += 1 + sovEpaxos(uint64(m.Epoch))
	}
	if m.NodeId != 0 {
		n += 1 + sovEpaxos(uint64(m.NodeId))
	}
	return n
}

func (m *Message) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovEpaxos(uint64(m.Type))
	}
	if len(m.Instance) > 0 {
		for _, e := range m.Instance {
			l = e.Size()
			n += 1 + l + sovEpaxos(uint64(l))
		}
	}
	if m.To != 0 {
		n += 1 + sovEpaxos(uint64(m.To))
	}
	if m.From != 0 {
		n += 1 + sovEpaxos(uint64(m.From))
	}
	l = m.Ballot.Size()
	n += 1 + l + sovEpaxos(uint64(l))
	if m.Reject {
		n += 2
	}
	if m.RejectType != 0 {
		n += 1 + sovEpaxos(uint64(m.RejectType))
	}
	return n
}

func (m *Instance) Size() (n int) {
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovEpaxos(uint64(m.ID))
	}
	if m.Cmd != nil {
		l = m.Cmd.Size()
		n += 1 + l + sovEpaxos(uint64(l))
	}
	if len(m.Deps) > 0 {
		for k, v := range m.Deps {
			_ = k
			_ = v
			mapEntrySize := 1 + sovEpaxos(uint64(k)) + 1 + sovEpaxos(uint64(v))
			n += mapEntrySize + 1 + sovEpaxos(uint64(mapEntrySize))
		}
	}
	if m.Seq != 0 {
		n += 1 + sovEpaxos(uint64(m.Seq))
	}
	if m.NodeID != 0 {
		n += 1 + sovEpaxos(uint64(m.NodeID))
	}
	if m.State != 0 {
		n += 1 + sovEpaxos(uint64(m.State))
	}
	return n
}

func (m *Cmd) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovEpaxos(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovEpaxos(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovEpaxos(uint64(m.Type))
	}
	if m.RequestID != 0 {
		n += 1 + sovEpaxos(uint64(m.RequestID))
	}
	if m.ConnectionID != 0 {
		n += 1 + sovEpaxos(uint64(m.ConnectionID))
	}
	return n
}

func sovEpaxos(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEpaxos(x uint64) (n int) {
	return sovEpaxos(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Ballot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEpaxos
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Ballot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ballot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeId", wireType)
			}
			m.NodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeId |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEpaxos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEpaxos
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEpaxos
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (MessageType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Instance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEpaxos
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Instance = append(m.Instance, &Instance{})
			if err := m.Instance[len(m.Instance)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			m.To = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.To |= (TypeNodeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			m.From = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.From |= (TypeNodeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ballot", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEpaxos
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Ballot.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reject", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Reject = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RejectType", wireType)
			}
			m.RejectType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RejectType |= (RespRejectType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEpaxos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEpaxos
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Instance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEpaxos
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Instance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Instance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (TypeInstanceID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cmd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEpaxos
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Cmd == nil {
				m.Cmd = &Cmd{}
			}
			if err := m.Cmd.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEpaxos
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var mapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				mapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if m.Deps == nil {
				m.Deps = make(map[TypeNodeID]TypeInstanceID)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEpaxos
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var mapvalue uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEpaxos
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					mapvalue |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Deps[TypeNodeID(mapkey)] = ((TypeInstanceID)(mapvalue))
			} else {
				var mapvalue TypeInstanceID
				m.Deps[TypeNodeID(mapkey)] = mapvalue
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seq", wireType)
			}
			m.Seq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seq |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			m.NodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeID |= (TypeNodeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= (InstanceState(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEpaxos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEpaxos
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Cmd) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEpaxos
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Cmd: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Cmd: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEpaxos
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEpaxos
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (CmdType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestID", wireType)
			}
			m.RequestID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionID", wireType)
			}
			m.ConnectionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConnectionID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEpaxos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEpaxos
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEpaxos(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEpaxos
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthEpaxos
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEpaxos
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipEpaxos(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthEpaxos = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEpaxos   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("epaxos.proto", fileDescriptorEpaxos) }

var fileDescriptorEpaxos = []byte{
	// 657 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x8d, 0xed, 0xfc, 0xde, 0xfc, 0xd4, 0xbd, 0x5f, 0xfb, 0xd5, 0xaa, 0xaa, 0x24, 0x8a, 0xf4,
	0x7d, 0x0a, 0x91, 0x30, 0x52, 0x58, 0xb4, 0x20, 0xb1, 0x20, 0x31, 0x42, 0x59, 0xa4, 0x2a, 0x03,
	0x7b, 0xe4, 0xda, 0x43, 0x1a, 0xa8, 0x3d, 0x6e, 0x66, 0x8a, 0xc8, 0x82, 0x0d, 0x4b, 0x1e, 0x00,
	0xf1, 0x0c, 0xbc, 0x07, 0x52, 0x97, 0x3c, 0x41, 0x8b, 0xca, 0x5b, 0x74, 0x85, 0x66, 0xc6, 0x71,
	0x12, 0xa5, 0xbb, 0xb9, 0xf7, 0x9c, 0x3b, 0x67, 0xee, 0xf1, 0x91, 0xa1, 0x46, 0x13, 0xff, 0x13,
	0xe3, 0x6e, 0x32, 0x63, 0x82, 0x61, 0x59, 0x57, 0xc9, 0xe9, 0xfe, 0xce, 0x84, 0x4d, 0x98, 0x6a,
	0x3e, 0x92, 0x27, 0x8d, 0x77, 0x0e, 0xa1, 0x38, 0xf0, 0xcf, 0xcf, 0x99, 0xc0, 0x1d, 0x28, 0xd0,
	0x84, 0x05, 0x67, 0x8e, 0xd1, 0x36, 0xba, 0x79, 0xa2, 0x0b, 0xdc, 0x83, 0x52, 0xcc, 0x42, 0xfa,
	0x76, 0x1a, 0x3a, 0xa6, 0xea, 0x17, 0x65, 0x39, 0x0a, 0x3b, 0x3f, 0x4c, 0x28, 0x8d, 0x29, 0xe7,
	0xfe, 0x84, 0xe2, 0x03, 0xc8, 0x8b, 0x79, 0x42, 0xd5, 0x64, 0xa3, 0xbf, 0xeb, 0x2e, 0x34, 0xdd,
	0x94, 0xf0, 0x66, 0x9e, 0x50, 0xa2, 0x28, 0xe8, 0x42, 0x79, 0x1a, 0x73, 0xe1, 0xc7, 0x01, 0x75,
	0xcc, 0xb6, 0xd5, 0xad, 0xf6, 0x71, 0x49, 0x1f, 0xa5, 0x08, 0xc9, 0x38, 0xd8, 0x04, 0x53, 0x30,
	0xc7, 0x92, 0xd2, 0x83, 0xc6, 0xdd, 0x75, 0x0b, 0xe4, 0x5d, 0xc7, 0xf2, 0x09, 0x1e, 0x31, 0x05,
	0xc3, 0x0e, 0xe4, 0xdf, 0xcd, 0x58, 0xe4, 0xe4, 0xef, 0x65, 0x28, 0x0c, 0x5d, 0x28, 0x9e, 0xaa,
	0x1d, 0x9d, 0x42, 0xdb, 0xe8, 0x56, 0xfb, 0xf6, 0x52, 0x51, 0xef, 0x3e, 0xc8, 0x5f, 0x5d, 0xb7,
	0x72, 0x24, 0x65, 0xe1, 0xbf, 0x50, 0x9c, 0xd1, 0xf7, 0x34, 0x10, 0x4e, 0xb1, 0x6d, 0x74, 0xcb,
	0x24, 0xad, 0xf0, 0x08, 0x40, 0x9f, 0xa4, 0x82, 0x53, 0x52, 0xcb, 0x3a, 0xcb, 0xbb, 0x08, 0xe5,
	0x09, 0xc9, 0x70, 0xb2, 0xc2, 0xed, 0xfc, 0x34, 0xa1, 0xbc, 0x58, 0x0e, 0x3b, 0x60, 0x8e, 0x3c,
	0xed, 0xf2, 0x00, 0xef, 0xae, 0x5b, 0x0d, 0x49, 0x59, 0xa0, 0x72, 0xad, 0x91, 0x87, 0x2d, 0xb0,
	0x82, 0x48, 0x5b, 0x5e, 0xed, 0xd7, 0x97, 0x1a, 0xc3, 0x28, 0x24, 0x12, 0x41, 0x02, 0xf9, 0x90,
	0x26, 0xdc, 0xb1, 0x94, 0x87, 0x07, 0x9b, 0x1e, 0xba, 0x1e, 0x4d, 0xf8, 0x8b, 0x58, 0xcc, 0xe6,
	0x83, 0xf6, 0x97, 0x9b, 0x55, 0x57, 0xbe, 0xde, 0x6c, 0x48, 0xaa, 0xbb, 0xd0, 0x06, 0x8b, 0xd3,
	0x0b, 0x6d, 0x25, 0x91, 0x47, 0xfc, 0x1f, 0xf4, 0xe7, 0xf6, 0x94, 0x73, 0x9b, 0xfe, 0xa6, 0x28,
	0x3e, 0x84, 0x02, 0x17, 0xbe, 0xa0, 0xca, 0xb0, 0x46, 0x7f, 0x6f, 0xf3, 0x39, 0xaf, 0x25, 0x4c,
	0x34, 0x6b, 0xff, 0x10, 0x2a, 0xd9, 0xeb, 0xa4, 0xea, 0x07, 0x3a, 0x4f, 0x53, 0x27, 0x8f, 0x32,
	0x89, 0x1f, 0xfd, 0xf3, 0x4b, 0x9a, 0x26, 0x4e, 0x17, 0x4f, 0xcd, 0x23, 0xa3, 0xf3, 0xcd, 0x00,
	0x6b, 0x18, 0x85, 0xab, 0x33, 0x95, 0x7b, 0x66, 0x6a, 0xe9, 0x0c, 0xfe, 0x97, 0x06, 0xd3, 0x52,
	0xcf, 0xda, 0x5e, 0xf3, 0x71, 0x25, 0x94, 0x07, 0x50, 0x99, 0xd1, 0x8b, 0x4b, 0xca, 0xc5, 0xc8,
	0x4b, 0xd7, 0x5f, 0x36, 0xb0, 0x03, 0xb5, 0x80, 0xc5, 0x31, 0x0d, 0xc4, 0x94, 0xc5, 0x0b, 0x2b,
	0xc8, 0x5a, 0xaf, 0xf7, 0x19, 0xaa, 0x2b, 0x59, 0x47, 0x1b, 0x6a, 0x63, 0x3e, 0x39, 0x99, 0xd1,
	0xe7, 0x41, 0x40, 0x13, 0x61, 0xe7, 0x70, 0x07, 0xec, 0xac, 0x93, 0x50, 0x21, 0xb3, 0x62, 0x1b,
	0x58, 0x87, 0xca, 0x98, 0x4f, 0x52, 0x92, 0x89, 0xdb, 0x50, 0xcf, 0x4a, 0xc5, 0xb0, 0x70, 0x17,
	0xb6, 0xc7, 0x7c, 0xa2, 0x63, 0xea, 0x4d, 0x79, 0xe4, 0x8b, 0xe0, 0xcc, 0xce, 0xa7, 0x83, 0x43,
	0x16, 0x45, 0x53, 0x61, 0x17, 0x7a, 0xcf, 0xa0, 0xbe, 0x66, 0x34, 0x6e, 0x41, 0x35, 0x53, 0xa7,
	0xa1, 0x9d, 0xc3, 0x1a, 0x94, 0xb3, 0x4a, 0xe9, 0xea, 0x59, 0x59, 0x9a, 0xbd, 0x27, 0xd0, 0x58,
	0x0f, 0x2f, 0x22, 0x34, 0xb4, 0xe6, 0x31, 0x13, 0x63, 0xa5, 0x99, 0xc3, 0x7f, 0x60, 0x2b, 0xbb,
	0x53, 0x53, 0x6d, 0xa3, 0xe7, 0x42, 0x29, 0xf5, 0x12, 0x01, 0x8a, 0xc3, 0x28, 0x7c, 0x49, 0xe5,
	0xba, 0xfa, 0x7c, 0x72, 0x29, 0x6c, 0x03, 0xab, 0x8a, 0xf2, 0x4a, 0x02, 0xd6, 0xc0, 0xbe, 0xba,
	0x6d, 0x1a, 0xbf, 0x6e, 0x9b, 0xc6, 0xef, 0xdb, 0xa6, 0xf1, 0xfd, 0x4f, 0x33, 0x77, 0x5a, 0x54,
	0x3f, 0xa2, 0xc7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xab, 0x39, 0x0c, 0xb8, 0x04, 0x00,
	0x00,
}