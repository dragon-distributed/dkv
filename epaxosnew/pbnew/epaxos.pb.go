// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: epaxos.proto

/*
	Package pbnew is a generated protocol buffer package.

	It is generated from these files:
		epaxos.proto

	It has these top-level messages:
		Cmd
		Ballot
		Instance
		Message
*/
package pbnew

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type InstanceType int32

const (
	InstanceType_InstanceTypeCmd      InstanceType = 0
	InstanceType_InstanceTypeElection InstanceType = 1
)

var InstanceType_name = map[int32]string{
	0: "InstanceTypeCmd",
	1: "InstanceTypeElection",
}
var InstanceType_value = map[string]int32{
	"InstanceTypeCmd":      0,
	"InstanceTypeElection": 1,
}

func (x InstanceType) String() string {
	return proto.EnumName(InstanceType_name, int32(x))
}
func (InstanceType) EnumDescriptor() ([]byte, []int) { return fileDescriptorEpaxos, []int{0} }

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
func (CmdType) EnumDescriptor() ([]byte, []int) { return fileDescriptorEpaxos, []int{1} }

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
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptorEpaxos, []int{2} }

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
func (InstanceState) EnumDescriptor() ([]byte, []int) { return fileDescriptorEpaxos, []int{3} }

type Cmd struct {
	Key       string  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value     []byte  `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Type      CmdType `protobuf:"varint,3,opt,name=type,proto3,enum=pbnew.CmdType" json:"type,omitempty"`
	RequestID uint64  `protobuf:"varint,4,opt,name=requestID,proto3" json:"requestID,omitempty"`
}

func (m *Cmd) Reset()                    { *m = Cmd{} }
func (m *Cmd) String() string            { return proto.CompactTextString(m) }
func (*Cmd) ProtoMessage()               {}
func (*Cmd) Descriptor() ([]byte, []int) { return fileDescriptorEpaxos, []int{0} }

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

type Ballot struct {
	Epoch  uint64 `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	NodeId uint64 `protobuf:"varint,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (m *Ballot) Reset()                    { *m = Ballot{} }
func (m *Ballot) String() string            { return proto.CompactTextString(m) }
func (*Ballot) ProtoMessage()               {}
func (*Ballot) Descriptor() ([]byte, []int) { return fileDescriptorEpaxos, []int{1} }

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

type Instance struct {
	ID           uint64            `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	NodeID       uint64            `protobuf:"varint,2,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	InstanceType InstanceType      `protobuf:"varint,3,opt,name=instanceType,proto3,enum=pbnew.InstanceType" json:"instanceType,omitempty"`
	State        InstanceState     `protobuf:"varint,4,opt,name=state,proto3,enum=pbnew.InstanceState" json:"state,omitempty"`
	Ballot       *Ballot           `protobuf:"bytes,5,opt,name=ballot" json:"ballot,omitempty"`
	Deps         map[uint64]uint64 `protobuf:"bytes,6,rep,name=deps" json:"deps,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Seq          uint64            `protobuf:"varint,7,opt,name=seq,proto3" json:"seq,omitempty"`
	Cmd          *Cmd              `protobuf:"bytes,8,opt,name=cmd" json:"cmd,omitempty"`
}

func (m *Instance) Reset()                    { *m = Instance{} }
func (m *Instance) String() string            { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()               {}
func (*Instance) Descriptor() ([]byte, []int) { return fileDescriptorEpaxos, []int{2} }

func (m *Instance) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Instance) GetNodeID() uint64 {
	if m != nil {
		return m.NodeID
	}
	return 0
}

func (m *Instance) GetInstanceType() InstanceType {
	if m != nil {
		return m.InstanceType
	}
	return InstanceType_InstanceTypeCmd
}

func (m *Instance) GetState() InstanceState {
	if m != nil {
		return m.State
	}
	return InstanceState_PreAccepted
}

func (m *Instance) GetBallot() *Ballot {
	if m != nil {
		return m.Ballot
	}
	return nil
}

func (m *Instance) GetDeps() map[uint64]uint64 {
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

func (m *Instance) GetCmd() *Cmd {
	if m != nil {
		return m.Cmd
	}
	return nil
}

type Message struct {
	Type      MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=pbnew.MessageType" json:"type,omitempty"`
	Instances []*Instance `protobuf:"bytes,2,rep,name=instances" json:"instances,omitempty"`
	To        uint64      `protobuf:"varint,3,opt,name=to,proto3" json:"to,omitempty"`
	From      uint64      `protobuf:"varint,4,opt,name=from,proto3" json:"from,omitempty"`
	Ballot    *Ballot     `protobuf:"bytes,5,opt,name=ballot" json:"ballot,omitempty"`
	Reject    bool        `protobuf:"varint,6,opt,name=reject,proto3" json:"reject,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptorEpaxos, []int{3} }

func (m *Message) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_MsgPreAccept
}

func (m *Message) GetInstances() []*Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

func (m *Message) GetTo() uint64 {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *Message) GetFrom() uint64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *Message) GetBallot() *Ballot {
	if m != nil {
		return m.Ballot
	}
	return nil
}

func (m *Message) GetReject() bool {
	if m != nil {
		return m.Reject
	}
	return false
}

func init() {
	proto.RegisterType((*Cmd)(nil), "pbnew.Cmd")
	proto.RegisterType((*Ballot)(nil), "pbnew.Ballot")
	proto.RegisterType((*Instance)(nil), "pbnew.Instance")
	proto.RegisterType((*Message)(nil), "pbnew.Message")
	proto.RegisterEnum("pbnew.InstanceType", InstanceType_name, InstanceType_value)
	proto.RegisterEnum("pbnew.CmdType", CmdType_name, CmdType_value)
	proto.RegisterEnum("pbnew.MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("pbnew.InstanceState", InstanceState_name, InstanceState_value)
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
	return i, nil
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
	if m.NodeID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(m.NodeID))
	}
	if m.InstanceType != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(m.InstanceType))
	}
	if m.State != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(m.State))
	}
	if m.Ballot != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(m.Ballot.Size()))
		n1, err := m.Ballot.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Deps) > 0 {
		for k, _ := range m.Deps {
			dAtA[i] = 0x32
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
		dAtA[i] = 0x38
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(m.Seq))
	}
	if m.Cmd != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(m.Cmd.Size()))
		n2, err := m.Cmd.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
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
	if len(m.Instances) > 0 {
		for _, msg := range m.Instances {
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
	if m.Ballot != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintEpaxos(dAtA, i, uint64(m.Ballot.Size()))
		n3, err := m.Ballot.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
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
	return n
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

func (m *Instance) Size() (n int) {
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovEpaxos(uint64(m.ID))
	}
	if m.NodeID != 0 {
		n += 1 + sovEpaxos(uint64(m.NodeID))
	}
	if m.InstanceType != 0 {
		n += 1 + sovEpaxos(uint64(m.InstanceType))
	}
	if m.State != 0 {
		n += 1 + sovEpaxos(uint64(m.State))
	}
	if m.Ballot != nil {
		l = m.Ballot.Size()
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
	if m.Cmd != nil {
		l = m.Cmd.Size()
		n += 1 + l + sovEpaxos(uint64(l))
	}
	return n
}

func (m *Message) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovEpaxos(uint64(m.Type))
	}
	if len(m.Instances) > 0 {
		for _, e := range m.Instances {
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
	if m.Ballot != nil {
		l = m.Ballot.Size()
		n += 1 + l + sovEpaxos(uint64(l))
	}
	if m.Reject {
		n += 2
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
				m.ID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
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
				m.NodeID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InstanceType", wireType)
			}
			m.InstanceType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpaxos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InstanceType |= (InstanceType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
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
			if m.Ballot == nil {
				m.Ballot = &Ballot{}
			}
			if err := m.Ballot.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
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
				m.Deps = make(map[uint64]uint64)
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
				m.Deps[mapkey] = mapvalue
			} else {
				var mapvalue uint64
				m.Deps[mapkey] = mapvalue
			}
			iNdEx = postIndex
		case 7:
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
		case 8:
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
				return fmt.Errorf("proto: wrong wireType = %d for field Instances", wireType)
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
			m.Instances = append(m.Instances, &Instance{})
			if err := m.Instances[len(m.Instances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
				m.To |= (uint64(b) & 0x7F) << shift
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
				m.From |= (uint64(b) & 0x7F) << shift
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
			if m.Ballot == nil {
				m.Ballot = &Ballot{}
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
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xed, 0xda, 0x8e, 0x93, 0x4c, 0x3e, 0xea, 0x4e, 0x43, 0x31, 0xa8, 0x8a, 0xac, 0x48, 0x20,
	0x2b, 0x52, 0x73, 0x08, 0x87, 0x22, 0xa4, 0x1e, 0x68, 0x5c, 0xa1, 0x1c, 0x22, 0x95, 0xa5, 0x77,
	0xe4, 0xda, 0x43, 0x1a, 0x88, 0x63, 0xd7, 0xbb, 0x05, 0x72, 0xe0, 0x7f, 0xf0, 0x7f, 0xb8, 0x70,
	0xe4, 0xc4, 0x19, 0x95, 0x3f, 0x82, 0xbc, 0x5e, 0x12, 0xb7, 0xe2, 0xc0, 0x6d, 0xdf, 0xcc, 0x9b,
	0x7d, 0x7e, 0xf3, 0x56, 0x86, 0x36, 0x65, 0xe1, 0xe7, 0x54, 0x8c, 0xb2, 0x3c, 0x95, 0x29, 0xd6,
	0xb2, 0xcb, 0x15, 0x7d, 0x1a, 0xa4, 0x60, 0x4e, 0x92, 0x18, 0x1d, 0x30, 0x3f, 0xd0, 0xda, 0x65,
	0x1e, 0xf3, 0x9b, 0xbc, 0x38, 0x62, 0x0f, 0x6a, 0x1f, 0xc3, 0xe5, 0x0d, 0xb9, 0x86, 0xc7, 0xfc,
	0x36, 0x2f, 0x01, 0x0e, 0xc0, 0x92, 0xeb, 0x8c, 0x5c, 0xd3, 0x63, 0x7e, 0x77, 0xdc, 0x1d, 0xa9,
	0x4b, 0x46, 0x93, 0x24, 0xbe, 0x58, 0x67, 0xc4, 0x55, 0x0f, 0x0f, 0xa1, 0x99, 0xd3, 0xf5, 0x0d,
	0x09, 0x39, 0x0d, 0x5c, 0xcb, 0x63, 0xbe, 0xc5, 0xb7, 0x85, 0xc1, 0x31, 0xd8, 0xa7, 0xe1, 0x72,
	0x99, 0xca, 0x42, 0x81, 0xb2, 0x34, 0xba, 0x52, 0xaa, 0x16, 0x2f, 0x01, 0x3e, 0x84, 0xfa, 0x2a,
	0x8d, 0xe9, 0xed, 0x22, 0x56, 0xca, 0x16, 0xb7, 0x0b, 0x38, 0x8d, 0x07, 0x3f, 0x0d, 0x68, 0x4c,
	0x57, 0x42, 0x86, 0xab, 0x88, 0xb0, 0x0b, 0xc6, 0x34, 0xd0, 0x83, 0xc6, 0x34, 0xc0, 0x03, 0x28,
	0x69, 0xc1, 0x9d, 0xa1, 0x00, 0x8f, 0xa1, 0xbd, 0xd0, 0x33, 0x17, 0xdb, 0xef, 0xde, 0xd7, 0xdf,
	0x3d, 0xad, 0xb4, 0xf8, 0x1d, 0x22, 0x0e, 0xa1, 0x26, 0x64, 0x28, 0x49, 0x19, 0xe8, 0x8e, 0x7b,
	0xf7, 0x26, 0xde, 0x14, 0x3d, 0x5e, 0x52, 0xf0, 0x09, 0xd8, 0x97, 0xca, 0x92, 0x5b, 0xf3, 0x98,
	0xdf, 0x1a, 0x77, 0x34, 0xb9, 0xf4, 0xc9, 0x75, 0x13, 0x8f, 0xc0, 0x8a, 0x29, 0x13, 0xae, 0xed,
	0x99, 0x7e, 0x6b, 0xfc, 0xe8, 0xde, 0x8d, 0xa3, 0x80, 0x32, 0x71, 0xb6, 0x92, 0xf9, 0x9a, 0x2b,
	0x5a, 0x11, 0x89, 0xa0, 0x6b, 0xb7, 0xae, 0xfc, 0x14, 0x47, 0x3c, 0x04, 0x33, 0x4a, 0x62, 0xb7,
	0xa1, 0x44, 0x60, 0xbb, 0x7b, 0x5e, 0x94, 0x1f, 0x1f, 0x43, 0x73, 0x73, 0x45, 0x35, 0x4f, 0xeb,
	0x1f, 0x79, 0x5a, 0x3a, 0xcf, 0x17, 0xc6, 0x73, 0x36, 0xf8, 0xc6, 0xa0, 0x3e, 0x23, 0x21, 0xc2,
	0x39, 0xe1, 0x53, 0x9d, 0x2f, 0x53, 0xae, 0x51, 0x6b, 0xe8, 0x6e, 0x25, 0xe3, 0x23, 0x68, 0xfe,
	0x5d, 0x97, 0x70, 0x0d, 0x65, 0x68, 0xf7, 0x9e, 0x21, 0xbe, 0x65, 0x14, 0x71, 0xc9, 0x54, 0x2d,
	0xdf, 0xe2, 0x86, 0x4c, 0x11, 0xc1, 0x7a, 0x97, 0xa7, 0x89, 0x7e, 0x1d, 0xea, 0xfc, 0xbf, 0x5b,
	0x3c, 0x00, 0x3b, 0xa7, 0xf7, 0x14, 0x49, 0xd7, 0xf6, 0x98, 0xdf, 0xe0, 0x1a, 0x0d, 0x4f, 0xa0,
	0x5d, 0x8d, 0x13, 0xf7, 0x61, 0xb7, 0x8a, 0x27, 0x49, 0xec, 0xec, 0xa0, 0x0b, 0xbd, 0x6a, 0xf1,
	0x6c, 0x49, 0x91, 0x5c, 0xa4, 0x2b, 0x87, 0x0d, 0x47, 0x50, 0xd7, 0xaf, 0x18, 0x01, 0xec, 0x49,
	0x12, 0xbf, 0x22, 0xe9, 0xec, 0xe8, 0xf3, 0xf9, 0x8d, 0x74, 0x18, 0xb6, 0x14, 0xe5, 0x75, 0xd1,
	0x30, 0x87, 0x5f, 0xa0, 0x55, 0xd9, 0x0a, 0x3a, 0xd0, 0x9e, 0x89, 0xf9, 0x79, 0x4e, 0x2f, 0xa3,
	0x88, 0xb2, 0x62, 0xb2, 0x07, 0xce, 0xa6, 0x92, 0x91, 0xe4, 0x24, 0x32, 0x87, 0x61, 0x07, 0x9a,
	0x33, 0x31, 0xd7, 0x24, 0x03, 0xf7, 0xa0, 0xb3, 0x81, 0x8a, 0x61, 0xe2, 0x03, 0xd8, 0x9b, 0x89,
	0x79, 0x69, 0x3a, 0x58, 0x88, 0x24, 0x94, 0xd1, 0x95, 0x63, 0xe9, 0xc1, 0x49, 0x9a, 0x24, 0x0b,
	0xe9, 0xd4, 0x86, 0x27, 0xd0, 0xb9, 0xf3, 0x14, 0x71, 0x17, 0x5a, 0x1b, 0x75, 0x2a, 0xac, 0xb6,
	0xa1, 0xb1, 0x41, 0x4a, 0xb7, 0x9c, 0x2d, 0xa0, 0x71, 0xea, 0x7c, 0xbf, 0xed, 0xb3, 0x1f, 0xb7,
	0x7d, 0xf6, 0xeb, 0xb6, 0xcf, 0xbe, 0xfe, 0xee, 0xef, 0x5c, 0xda, 0xea, 0xaf, 0xf0, 0xec, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xcd, 0x19, 0xc1, 0x25, 0x04, 0x00, 0x00,
}
