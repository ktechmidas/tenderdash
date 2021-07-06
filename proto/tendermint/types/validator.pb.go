// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/types/validator.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	crypto "github.com/tendermint/tendermint/proto/tendermint/crypto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ValidatorSet struct {
	Validators         []*Validator     `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators,omitempty"`
	Proposer           *Validator       `protobuf:"bytes,2,opt,name=proposer,proto3" json:"proposer,omitempty"`
	TotalVotingPower   int64            `protobuf:"varint,3,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"total_voting_power,omitempty"`
	ThresholdPublicKey crypto.PublicKey `protobuf:"bytes,4,opt,name=threshold_public_key,json=thresholdPublicKey,proto3" json:"threshold_public_key"`
	QuorumType         int32            `protobuf:"varint,5,opt,name=quorum_type,json=quorumType,proto3" json:"quorum_type,omitempty"`
	QuorumHash         []byte           `protobuf:"bytes,6,opt,name=quorum_hash,json=quorumHash,proto3" json:"quorum_hash,omitempty"`
	HasPublicKeys      bool             `protobuf:"varint,7,opt,name=has_public_keys,json=hasPublicKeys,proto3" json:"has_public_keys,omitempty"`
}

func (m *ValidatorSet) Reset()         { *m = ValidatorSet{} }
func (m *ValidatorSet) String() string { return proto.CompactTextString(m) }
func (*ValidatorSet) ProtoMessage()    {}
func (*ValidatorSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e92274df03d3088, []int{0}
}
func (m *ValidatorSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSet.Merge(m, src)
}
func (m *ValidatorSet) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSet.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSet proto.InternalMessageInfo

func (m *ValidatorSet) GetValidators() []*Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *ValidatorSet) GetProposer() *Validator {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func (m *ValidatorSet) GetTotalVotingPower() int64 {
	if m != nil {
		return m.TotalVotingPower
	}
	return 0
}

func (m *ValidatorSet) GetThresholdPublicKey() crypto.PublicKey {
	if m != nil {
		return m.ThresholdPublicKey
	}
	return crypto.PublicKey{}
}

func (m *ValidatorSet) GetQuorumType() int32 {
	if m != nil {
		return m.QuorumType
	}
	return 0
}

func (m *ValidatorSet) GetQuorumHash() []byte {
	if m != nil {
		return m.QuorumHash
	}
	return nil
}

func (m *ValidatorSet) GetHasPublicKeys() bool {
	if m != nil {
		return m.HasPublicKeys
	}
	return false
}

type Validator struct {
	PubKey           *crypto.PublicKey `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	VotingPower      int64             `protobuf:"varint,3,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
	ProposerPriority int64             `protobuf:"varint,4,opt,name=proposer_priority,json=proposerPriority,proto3" json:"proposer_priority,omitempty"`
	ProTxHash        []byte            `protobuf:"bytes,5,opt,name=pro_tx_hash,json=proTxHash,proto3" json:"pro_tx_hash,omitempty"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e92274df03d3088, []int{1}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetPubKey() *crypto.PublicKey {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *Validator) GetVotingPower() int64 {
	if m != nil {
		return m.VotingPower
	}
	return 0
}

func (m *Validator) GetProposerPriority() int64 {
	if m != nil {
		return m.ProposerPriority
	}
	return 0
}

func (m *Validator) GetProTxHash() []byte {
	if m != nil {
		return m.ProTxHash
	}
	return nil
}

type SimpleValidator struct {
	PubKey      *crypto.PublicKey `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	VotingPower int64             `protobuf:"varint,2,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
}

func (m *SimpleValidator) Reset()         { *m = SimpleValidator{} }
func (m *SimpleValidator) String() string { return proto.CompactTextString(m) }
func (*SimpleValidator) ProtoMessage()    {}
func (*SimpleValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e92274df03d3088, []int{2}
}
func (m *SimpleValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SimpleValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SimpleValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SimpleValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleValidator.Merge(m, src)
}
func (m *SimpleValidator) XXX_Size() int {
	return m.Size()
}
func (m *SimpleValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleValidator.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleValidator proto.InternalMessageInfo

func (m *SimpleValidator) GetPubKey() *crypto.PublicKey {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *SimpleValidator) GetVotingPower() int64 {
	if m != nil {
		return m.VotingPower
	}
	return 0
}

func init() {
	proto.RegisterType((*ValidatorSet)(nil), "tendermint.types.ValidatorSet")
	proto.RegisterType((*Validator)(nil), "tendermint.types.Validator")
	proto.RegisterType((*SimpleValidator)(nil), "tendermint.types.SimpleValidator")
}

func init() { proto.RegisterFile("tendermint/types/validator.proto", fileDescriptor_4e92274df03d3088) }

var fileDescriptor_4e92274df03d3088 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6e, 0xd3, 0x40,
	0x14, 0xce, 0x24, 0x4d, 0xda, 0xbe, 0x04, 0x35, 0x8c, 0xba, 0xb0, 0x4a, 0xe5, 0x9a, 0x2c, 0x90,
	0x25, 0x90, 0x2d, 0x81, 0x50, 0x17, 0xdd, 0x75, 0x85, 0xc4, 0x26, 0xb8, 0x51, 0x17, 0x6c, 0x2c,
	0x3b, 0x19, 0xc5, 0xa3, 0x38, 0x79, 0xc3, 0xcc, 0x38, 0xd4, 0xb7, 0xe0, 0x38, 0x1c, 0xa1, 0xcb,
	0x2e, 0x59, 0x21, 0x94, 0xdc, 0x80, 0x13, 0x20, 0x8f, 0x89, 0x6d, 0x15, 0x50, 0xd8, 0x8d, 0xde,
	0xf7, 0xbd, 0xf9, 0x7e, 0xa4, 0x07, 0x8e, 0x66, 0xab, 0x19, 0x93, 0x4b, 0xbe, 0xd2, 0xbe, 0xce,
	0x05, 0x53, 0xfe, 0x3a, 0x4a, 0xf9, 0x2c, 0xd2, 0x28, 0x3d, 0x21, 0x51, 0x23, 0x1d, 0xd6, 0x0c,
	0xcf, 0x30, 0xce, 0x4e, 0xe7, 0x38, 0x47, 0x03, 0xfa, 0xc5, 0xab, 0xe4, 0x9d, 0x9d, 0x37, 0x7e,
	0x9a, 0xca, 0x5c, 0x68, 0xf4, 0x17, 0x2c, 0x57, 0x25, 0x3a, 0xfa, 0xd9, 0x86, 0xc1, 0xed, 0xee,
	0xe7, 0x1b, 0xa6, 0xe9, 0x15, 0x40, 0xa5, 0xa4, 0x2c, 0xe2, 0x74, 0xdc, 0xfe, 0xeb, 0x67, 0xde,
	0x63, 0x2d, 0xaf, 0xda, 0x09, 0x1a, 0x74, 0x7a, 0x09, 0x47, 0x42, 0xa2, 0x40, 0xc5, 0xa4, 0xd5,
	0x76, 0xc8, 0xbe, 0xd5, 0x8a, 0x4c, 0x5f, 0x01, 0xd5, 0xa8, 0xa3, 0x34, 0x5c, 0xa3, 0xe6, 0xab,
	0x79, 0x28, 0xf0, 0x33, 0x93, 0x56, 0xc7, 0x21, 0x6e, 0x27, 0x18, 0x1a, 0xe4, 0xd6, 0x00, 0xe3,
	0x62, 0x4e, 0x27, 0x70, 0xaa, 0x13, 0xc9, 0x54, 0x82, 0xe9, 0x2c, 0x14, 0x59, 0x9c, 0xf2, 0x69,
	0xb8, 0x60, 0xb9, 0x75, 0x60, 0x24, 0xcf, 0x9b, 0x92, 0x65, 0x62, 0x6f, 0x6c, 0x48, 0xef, 0x59,
	0x7e, 0x7d, 0x70, 0xff, 0xfd, 0xa2, 0x15, 0xd0, 0x6a, 0xbf, 0x42, 0xe8, 0x05, 0xf4, 0x3f, 0x65,
	0x28, 0xb3, 0x65, 0x58, 0xf8, 0xb4, 0xba, 0x0e, 0x71, 0xbb, 0x01, 0x94, 0xa3, 0x49, 0x2e, 0x58,
	0x83, 0x90, 0x44, 0x2a, 0xb1, 0x7a, 0x0e, 0x71, 0x07, 0x3b, 0xc2, 0xbb, 0x48, 0x25, 0xf4, 0x05,
	0x9c, 0x24, 0x91, 0x6a, 0x38, 0x52, 0xd6, 0xa1, 0x43, 0xdc, 0xa3, 0xe0, 0x49, 0x12, 0xa9, 0x4a,
	0x48, 0x8d, 0xbe, 0x12, 0x38, 0xae, 0x5a, 0xa0, 0x57, 0x70, 0x28, 0xb2, 0xd8, 0x04, 0x68, 0xff,
	0x67, 0x00, 0x12, 0xf4, 0x44, 0x16, 0x17, 0xa6, 0x9f, 0xc3, 0xe0, 0x2f, 0x95, 0xf5, 0xd7, 0x8d,
	0xb6, 0x5e, 0xc2, 0xd3, 0x5d, 0xcf, 0xa1, 0x90, 0x1c, 0x25, 0xd7, 0x65, 0x55, 0x9d, 0x60, 0xb8,
	0x03, 0xc6, 0xbf, 0xe7, 0xd4, 0x86, 0xbe, 0x90, 0x18, 0xea, 0xbb, 0x32, 0x63, 0xd7, 0x64, 0x3c,
	0x16, 0x12, 0x27, 0x77, 0x45, 0xc4, 0xd1, 0x02, 0x4e, 0x6e, 0xf8, 0x52, 0xa4, 0xac, 0xf6, 0xff,
	0xb6, 0xf6, 0x4f, 0xf6, 0xfb, 0xff, 0xa7, 0xf3, 0xf6, 0x1f, 0xce, 0xaf, 0x3f, 0xdc, 0x6f, 0x6c,
	0xf2, 0xb0, 0xb1, 0xc9, 0x8f, 0x8d, 0x4d, 0xbe, 0x6c, 0xed, 0xd6, 0xc3, 0xd6, 0x6e, 0x7d, 0xdb,
	0xda, 0xad, 0x8f, 0x97, 0x73, 0xae, 0x93, 0x2c, 0xf6, 0xa6, 0xb8, 0xf4, 0x9b, 0x97, 0x52, 0x3f,
	0xcb, 0x3b, 0x78, 0x7c, 0x45, 0x71, 0xcf, 0xcc, 0xdf, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x71,
	0x41, 0x97, 0x2e, 0x60, 0x03, 0x00, 0x00,
}

func (m *ValidatorSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HasPublicKeys {
		i--
		if m.HasPublicKeys {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.QuorumHash) > 0 {
		i -= len(m.QuorumHash)
		copy(dAtA[i:], m.QuorumHash)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.QuorumHash)))
		i--
		dAtA[i] = 0x32
	}
	if m.QuorumType != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.QuorumType))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.ThresholdPublicKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.TotalVotingPower != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.TotalVotingPower))
		i--
		dAtA[i] = 0x18
	}
	if m.Proposer != nil {
		{
			size, err := m.Proposer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintValidator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintValidator(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ProTxHash) > 0 {
		i -= len(m.ProTxHash)
		copy(dAtA[i:], m.ProTxHash)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.ProTxHash)))
		i--
		dAtA[i] = 0x2a
	}
	if m.ProposerPriority != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.ProposerPriority))
		i--
		dAtA[i] = 0x20
	}
	if m.VotingPower != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.VotingPower))
		i--
		dAtA[i] = 0x18
	}
	if m.PubKey != nil {
		{
			size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintValidator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *SimpleValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SimpleValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SimpleValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VotingPower != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.VotingPower))
		i--
		dAtA[i] = 0x10
	}
	if m.PubKey != nil {
		{
			size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintValidator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintValidator(dAtA []byte, offset int, v uint64) int {
	offset -= sovValidator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovValidator(uint64(l))
		}
	}
	if m.Proposer != nil {
		l = m.Proposer.Size()
		n += 1 + l + sovValidator(uint64(l))
	}
	if m.TotalVotingPower != 0 {
		n += 1 + sovValidator(uint64(m.TotalVotingPower))
	}
	l = m.ThresholdPublicKey.Size()
	n += 1 + l + sovValidator(uint64(l))
	if m.QuorumType != 0 {
		n += 1 + sovValidator(uint64(m.QuorumType))
	}
	l = len(m.QuorumHash)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	if m.HasPublicKeys {
		n += 2
	}
	return n
}

func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PubKey != nil {
		l = m.PubKey.Size()
		n += 1 + l + sovValidator(uint64(l))
	}
	if m.VotingPower != 0 {
		n += 1 + sovValidator(uint64(m.VotingPower))
	}
	if m.ProposerPriority != 0 {
		n += 1 + sovValidator(uint64(m.ProposerPriority))
	}
	l = len(m.ProTxHash)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	return n
}

func (m *SimpleValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PubKey != nil {
		l = m.PubKey.Size()
		n += 1 + l + sovValidator(uint64(l))
	}
	if m.VotingPower != 0 {
		n += 1 + sovValidator(uint64(m.VotingPower))
	}
	return n
}

func sovValidator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozValidator(x uint64) (n int) {
	return sovValidator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ValidatorSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, &Validator{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proposer == nil {
				m.Proposer = &Validator{}
			}
			if err := m.Proposer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalVotingPower", wireType)
			}
			m.TotalVotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalVotingPower |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThresholdPublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ThresholdPublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuorumType", wireType)
			}
			m.QuorumType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QuorumType |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuorumHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QuorumHash = append(m.QuorumHash[:0], dAtA[iNdEx:postIndex]...)
			if m.QuorumHash == nil {
				m.QuorumHash = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasPublicKeys", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.HasPublicKeys = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PubKey == nil {
				m.PubKey = &crypto.PublicKey{}
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingPower", wireType)
			}
			m.VotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingPower |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposerPriority", wireType)
			}
			m.ProposerPriority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposerPriority |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProTxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProTxHash = append(m.ProTxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ProTxHash == nil {
				m.ProTxHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func (m *SimpleValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SimpleValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SimpleValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PubKey == nil {
				m.PubKey = &crypto.PublicKey{}
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingPower", wireType)
			}
			m.VotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingPower |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func skipValidator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowValidator
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
			if length < 0 {
				return 0, ErrInvalidLengthValidator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupValidator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthValidator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthValidator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValidator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupValidator = fmt.Errorf("proto: unexpected end of group")
)
