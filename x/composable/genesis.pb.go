// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lbm/composable/v1alpha1/genesis.proto

package composable

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_line_lbm_sdk_types "github.com/line/lbm-sdk/types"
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

// GenesisState defines the nft module's genesis state.
type GenesisState struct {
	// all the paramaters of the module
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// nfts grouped by class
	Nfts []ClassNFTs `protobuf:"bytes,2,rep,name=nfts,proto3" json:"nfts"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_459f8203f46123fd, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetNfts() []ClassNFTs {
	if m != nil {
		return m.Nfts
	}
	return nil
}

// ClassNFTs defines all nft of a class.
type ClassNFTs struct {
	// class of the following nfts
	Class Class `protobuf:"bytes,1,opt,name=class,proto3" json:"class"`
	// previous nft id of the class
	PreviousId github_com_line_lbm_sdk_types.Uint `protobuf:"bytes,2,opt,name=previous_id,json=previousId,proto3,customtype=github.com/line/lbm-sdk/types.Uint" json:"previous_id"`
	// groups of nft states of the same class
	NftStates []NFTState `protobuf:"bytes,3,rep,name=nft_states,json=nftStates,proto3" json:"nft_states"`
}

func (m *ClassNFTs) Reset()         { *m = ClassNFTs{} }
func (m *ClassNFTs) String() string { return proto.CompactTextString(m) }
func (*ClassNFTs) ProtoMessage()    {}
func (*ClassNFTs) Descriptor() ([]byte, []int) {
	return fileDescriptor_459f8203f46123fd, []int{1}
}
func (m *ClassNFTs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClassNFTs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClassNFTs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClassNFTs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassNFTs.Merge(m, src)
}
func (m *ClassNFTs) XXX_Size() int {
	return m.Size()
}
func (m *ClassNFTs) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassNFTs.DiscardUnknown(m)
}

var xxx_messageInfo_ClassNFTs proto.InternalMessageInfo

func (m *ClassNFTs) GetClass() Class {
	if m != nil {
		return m.Class
	}
	return Class{}
}

func (m *ClassNFTs) GetNftStates() []NFTState {
	if m != nil {
		return m.NftStates
	}
	return nil
}

// NFTState defines state of an nft.
type NFTState struct {
	// metadata of the nft
	Nft NFT `protobuf:"bytes,1,opt,name=nft,proto3" json:"nft"`
	// owner of the nft
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// parent of the nft
	Parent *FullID `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (m *NFTState) Reset()         { *m = NFTState{} }
func (m *NFTState) String() string { return proto.CompactTextString(m) }
func (*NFTState) ProtoMessage()    {}
func (*NFTState) Descriptor() ([]byte, []int) {
	return fileDescriptor_459f8203f46123fd, []int{2}
}
func (m *NFTState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NFTState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NFTState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NFTState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NFTState.Merge(m, src)
}
func (m *NFTState) XXX_Size() int {
	return m.Size()
}
func (m *NFTState) XXX_DiscardUnknown() {
	xxx_messageInfo_NFTState.DiscardUnknown(m)
}

var xxx_messageInfo_NFTState proto.InternalMessageInfo

func (m *NFTState) GetNft() NFT {
	if m != nil {
		return m.Nft
	}
	return NFT{}
}

func (m *NFTState) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *NFTState) GetParent() *FullID {
	if m != nil {
		return m.Parent
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "lbm.composable.v1alpha1.GenesisState")
	proto.RegisterType((*ClassNFTs)(nil), "lbm.composable.v1alpha1.ClassNFTs")
	proto.RegisterType((*NFTState)(nil), "lbm.composable.v1alpha1.NFTState")
}

func init() {
	proto.RegisterFile("lbm/composable/v1alpha1/genesis.proto", fileDescriptor_459f8203f46123fd)
}

var fileDescriptor_459f8203f46123fd = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x33, 0x4d, 0x5b, 0xec, 0xd4, 0xd3, 0x20, 0x18, 0x8a, 0xa4, 0x35, 0x2a, 0x14, 0xc1,
	0x84, 0x56, 0x41, 0x10, 0xf5, 0x50, 0x25, 0x52, 0x84, 0x22, 0xb1, 0x5e, 0xbc, 0x94, 0xa4, 0x9d,
	0xa4, 0xc1, 0xc9, 0x4c, 0xc8, 0x4c, 0xeb, 0xee, 0x67, 0xd8, 0xcb, 0x5e, 0xf6, 0x3b, 0xf5, 0xd8,
	0xe3, 0xb2, 0x2c, 0x65, 0x69, 0xbf, 0xc8, 0x92, 0xc9, 0x64, 0x77, 0x2f, 0xe9, 0xde, 0x32, 0x79,
	0xbf, 0xff, 0x7f, 0xde, 0xfb, 0xcf, 0x83, 0x6f, 0x48, 0x90, 0x38, 0x73, 0x96, 0xa4, 0x8c, 0xfb,
	0x01, 0xc1, 0xce, 0x7a, 0xe0, 0x93, 0x74, 0xe9, 0x0f, 0x9c, 0x08, 0x53, 0xcc, 0x63, 0x6e, 0xa7,
	0x19, 0x13, 0x0c, 0x3d, 0x27, 0x41, 0x62, 0xdf, 0x63, 0x76, 0x89, 0x75, 0x9e, 0x45, 0x2c, 0x62,
	0x92, 0x71, 0xf2, 0xaf, 0x02, 0xef, 0xbc, 0xaa, 0x72, 0x15, 0xa7, 0x29, 0x56, 0x9e, 0xd6, 0x19,
	0x80, 0x4f, 0x7f, 0x14, 0xb7, 0xfc, 0x16, 0xbe, 0xc0, 0xe8, 0x0b, 0x6c, 0xa6, 0x7e, 0xe6, 0x27,
	0xdc, 0x00, 0x3d, 0xd0, 0x6f, 0x0f, 0xbb, 0x76, 0xc5, 0xad, 0xf6, 0x2f, 0x89, 0x8d, 0xea, 0x9b,
	0x5d, 0x57, 0xf3, 0x94, 0x08, 0x7d, 0x86, 0x75, 0x1a, 0x0a, 0x6e, 0xd4, 0x7a, 0x7a, 0xbf, 0x3d,
	0xb4, 0x2a, 0xc5, 0xdf, 0x88, 0xcf, 0xf9, 0xc4, 0x9d, 0x96, 0x7a, 0xa9, 0xb2, 0xae, 0x01, 0x6c,
	0xdd, 0x55, 0xd0, 0x27, 0xd8, 0x98, 0xe7, 0x07, 0xd5, 0x89, 0x79, 0xdc, 0x4c, 0x19, 0x15, 0x12,
	0xf4, 0x13, 0xb6, 0xd3, 0x0c, 0xaf, 0x63, 0xb6, 0xe2, 0xb3, 0x78, 0x61, 0xd4, 0x7a, 0xa0, 0xdf,
	0x1a, 0xbd, 0xcd, 0x89, 0xab, 0x5d, 0xd7, 0x8a, 0x62, 0xb1, 0x5c, 0x05, 0xb9, 0x97, 0x43, 0x62,
	0x8a, 0x1d, 0x12, 0x24, 0xef, 0xf8, 0xe2, 0x9f, 0x4a, 0xe6, 0x4f, 0x4c, 0x85, 0x07, 0x4b, 0xf9,
	0x78, 0x81, 0x5c, 0x08, 0x69, 0x28, 0x66, 0x3c, 0x0f, 0x88, 0x1b, 0xba, 0x1c, 0xed, 0x65, 0x65,
	0x37, 0x13, 0x77, 0x2a, 0xa3, 0x54, 0x0d, 0xb5, 0x68, 0x28, 0xe4, 0x99, 0x5b, 0x17, 0x00, 0x3e,
	0x29, 0xab, 0xe8, 0x03, 0xd4, 0x69, 0x28, 0xd4, 0x6c, 0x2f, 0x8e, 0xb9, 0x29, 0xa3, 0x1c, 0x47,
	0x1d, 0xd8, 0x60, 0xff, 0x29, 0xce, 0xd4, 0x44, 0x79, 0x05, 0x78, 0xc5, 0x2f, 0xf4, 0x51, 0x3e,
	0x1d, 0xa6, 0xc2, 0xd0, 0x1f, 0x79, 0x3a, 0x77, 0x45, 0xc8, 0xf8, 0xbb, 0xa7, 0xf0, 0xd1, 0xd7,
	0xcd, 0xde, 0x04, 0xdb, 0xbd, 0x09, 0x6e, 0xf6, 0x26, 0x38, 0x3f, 0x98, 0xda, 0xf6, 0x60, 0x6a,
	0x97, 0x07, 0x53, 0xfb, 0xfb, 0xba, 0x2a, 0xa9, 0x93, 0x07, 0x9b, 0x15, 0x34, 0xe5, 0x2e, 0xbd,
	0xbf, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x9a, 0x78, 0x2e, 0xc8, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Nfts) > 0 {
		for iNdEx := len(m.Nfts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Nfts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ClassNFTs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassNFTs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClassNFTs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NftStates) > 0 {
		for iNdEx := len(m.NftStates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NftStates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size := m.PreviousId.Size()
		i -= size
		if _, err := m.PreviousId.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Class.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *NFTState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NFTState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NFTState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Parent != nil {
		{
			size, err := m.Parent.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Nft.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Nfts) > 0 {
		for _, e := range m.Nfts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *ClassNFTs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Class.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.PreviousId.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.NftStates) > 0 {
		for _, e := range m.NftStates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *NFTState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Nft.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Parent != nil {
		l = m.Parent.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nfts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nfts = append(m.Nfts, ClassNFTs{})
			if err := m.Nfts[len(m.Nfts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *ClassNFTs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ClassNFTs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassNFTs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Class", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Class.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PreviousId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftStates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftStates = append(m.NftStates, NFTState{})
			if err := m.NftStates[len(m.NftStates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *NFTState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: NFTState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NFTState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nft", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Nft.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Parent == nil {
				m.Parent = &FullID{}
			}
			if err := m.Parent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
