// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: coreum/asset/ft/v1/authz.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// MintAuthorization allows the grantee to mint up to mint_limit coin from
// the granter's account.
type MintAuthorization struct {
	MintLimit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=mint_limit,json=mintLimit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"mint_limit"`
}

func (m *MintAuthorization) Reset()         { *m = MintAuthorization{} }
func (m *MintAuthorization) String() string { return proto.CompactTextString(m) }
func (*MintAuthorization) ProtoMessage()    {}
func (*MintAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6a458149a08610, []int{0}
}
func (m *MintAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MintAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MintAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MintAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintAuthorization.Merge(m, src)
}
func (m *MintAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *MintAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_MintAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_MintAuthorization proto.InternalMessageInfo

func (m *MintAuthorization) GetMintLimit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.MintLimit
	}
	return nil
}

// BurnAuthorization allows the grantee to burn up to burn_limit coin from
// the granter's account.
type BurnAuthorization struct {
	BurnLimit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=burn_limit,json=burnLimit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"burn_limit"`
}

func (m *BurnAuthorization) Reset()         { *m = BurnAuthorization{} }
func (m *BurnAuthorization) String() string { return proto.CompactTextString(m) }
func (*BurnAuthorization) ProtoMessage()    {}
func (*BurnAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6a458149a08610, []int{1}
}
func (m *BurnAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BurnAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BurnAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BurnAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BurnAuthorization.Merge(m, src)
}
func (m *BurnAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *BurnAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_BurnAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_BurnAuthorization proto.InternalMessageInfo

func (m *BurnAuthorization) GetBurnLimit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.BurnLimit
	}
	return nil
}

func init() {
	proto.RegisterType((*MintAuthorization)(nil), "coreum.asset.ft.v1.MintAuthorization")
	proto.RegisterType((*BurnAuthorization)(nil), "coreum.asset.ft.v1.BurnAuthorization")
}

func init() { proto.RegisterFile("coreum/asset/ft/v1/authz.proto", fileDescriptor_8e6a458149a08610) }

var fileDescriptor_8e6a458149a08610 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xce, 0x2f, 0x4a,
	0x2d, 0xcd, 0xd5, 0x4f, 0x2c, 0x2e, 0x4e, 0x2d, 0xd1, 0x4f, 0x2b, 0xd1, 0x2f, 0x33, 0xd4, 0x4f,
	0x2c, 0x2d, 0xc9, 0xa8, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x82, 0xc8, 0xeb, 0x81,
	0xe5, 0xf5, 0xd2, 0x4a, 0xf4, 0xca, 0x0c, 0xa5, 0x04, 0x13, 0x73, 0x33, 0xf3, 0xf2, 0xf5, 0xc1,
	0x24, 0x44, 0x99, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x98, 0xa9, 0x0f, 0x62, 0x41, 0x45, 0x25,
	0x93, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xe3, 0x21, 0x12, 0x10, 0x0e, 0x54, 0x4a, 0x0e, 0xc2, 0xd3,
	0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x4f, 0xce, 0xcf,
	0xcc, 0x83, 0xc8, 0x2b, 0x9d, 0x65, 0xe4, 0x12, 0xf4, 0xcd, 0xcc, 0x2b, 0x71, 0x2c, 0x2d, 0xc9,
	0xc8, 0x2f, 0xca, 0xac, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0x13, 0xca, 0xe7, 0xe2, 0xca, 0xcd, 0xcc,
	0x2b, 0x89, 0xcf, 0xc9, 0xcc, 0xcd, 0x2c, 0x91, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x92, 0xd4,
	0x83, 0x1a, 0x0c, 0x32, 0x4a, 0x0f, 0x6a, 0x94, 0x9e, 0x73, 0x7e, 0x66, 0x9e, 0x93, 0xe9, 0x89,
	0x7b, 0xf2, 0x0c, 0xab, 0xee, 0xcb, 0x6b, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7,
	0xe7, 0x42, 0x5d, 0x01, 0xa5, 0x74, 0x8b, 0x53, 0xb2, 0xf5, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0xc1,
	0x1a, 0x8a, 0x57, 0x3c, 0xdf, 0xa0, 0xc5, 0x10, 0xc4, 0x09, 0xb2, 0xc3, 0x07, 0x64, 0x85, 0x95,
	0xfb, 0xa9, 0x2d, 0xba, 0x4a, 0x50, 0xf3, 0x21, 0xc1, 0x02, 0xb3, 0x00, 0xc5, 0x61, 0x5d, 0xcf,
	0x37, 0x68, 0xc9, 0x20, 0x19, 0x89, 0xe1, 0x72, 0xb0, 0x7f, 0x9c, 0x4a, 0x8b, 0xf2, 0x30, 0xfc,
	0x93, 0x54, 0x5a, 0x94, 0x47, 0x6b, 0xff, 0x80, 0xec, 0xa0, 0xc8, 0x3f, 0x18, 0x2e, 0x77, 0x0a,
	0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96,
	0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x33, 0x24, 0xc7, 0x39, 0x83,
	0x13, 0x8f, 0x5b, 0x7e, 0x69, 0x5e, 0x0a, 0x58, 0x9b, 0x3e, 0x34, 0xb5, 0x95, 0x19, 0xeb, 0x57,
	0x20, 0x92, 0x1c, 0xd8, 0xc1, 0x49, 0x6c, 0xe0, 0x88, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x21, 0xfb, 0x13, 0xa8, 0x92, 0x02, 0x00, 0x00,
}

func (m *MintAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MintAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MintAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MintLimit) > 0 {
		for iNdEx := len(m.MintLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MintLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAuthz(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *BurnAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BurnAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BurnAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BurnLimit) > 0 {
		for iNdEx := len(m.BurnLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BurnLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAuthz(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuthz(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuthz(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MintAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MintLimit) > 0 {
		for _, e := range m.MintLimit {
			l = e.Size()
			n += 1 + l + sovAuthz(uint64(l))
		}
	}
	return n
}

func (m *BurnAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.BurnLimit) > 0 {
		for _, e := range m.BurnLimit {
			l = e.Size()
			n += 1 + l + sovAuthz(uint64(l))
		}
	}
	return n
}

func sovAuthz(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuthz(x uint64) (n int) {
	return sovAuthz(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MintAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: MintAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MintAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintLimit = append(m.MintLimit, types.Coin{})
			if err := m.MintLimit[len(m.MintLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func (m *BurnAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: BurnAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BurnAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BurnLimit = append(m.BurnLimit, types.Coin{})
			if err := m.BurnLimit[len(m.BurnLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func skipAuthz(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
				return 0, ErrInvalidLengthAuthz
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuthz
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuthz
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuthz        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthz          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuthz = fmt.Errorf("proto: unexpected end of group")
)
