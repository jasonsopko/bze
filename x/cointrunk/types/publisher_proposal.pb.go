// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cointrunk/publisher_proposal.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type PublisherProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Address     string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Active      bool   `protobuf:"varint,5,opt,name=active,proto3" json:"active,omitempty"`
}

func (m *PublisherProposal) Reset()         { *m = PublisherProposal{} }
func (m *PublisherProposal) String() string { return proto.CompactTextString(m) }
func (*PublisherProposal) ProtoMessage()    {}
func (*PublisherProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_3acb71719072ee70, []int{0}
}
func (m *PublisherProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PublisherProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PublisherProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PublisherProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublisherProposal.Merge(m, src)
}
func (m *PublisherProposal) XXX_Size() int {
	return m.Size()
}
func (m *PublisherProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_PublisherProposal.DiscardUnknown(m)
}

var xxx_messageInfo_PublisherProposal proto.InternalMessageInfo

func (m *PublisherProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PublisherProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PublisherProposal) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PublisherProposal) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *PublisherProposal) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func init() {
	proto.RegisterType((*PublisherProposal)(nil), "bze.cointrunk.PublisherProposal")
}

func init() {
	proto.RegisterFile("cointrunk/publisher_proposal.proto", fileDescriptor_3acb71719072ee70)
}

var fileDescriptor_3acb71719072ee70 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x31, 0x4e, 0xc3, 0x30,
	0x18, 0x85, 0x63, 0x68, 0x0b, 0x18, 0x31, 0x60, 0x21, 0xe4, 0xc9, 0x8a, 0x3a, 0x75, 0x21, 0x1e,
	0xb8, 0x01, 0x13, 0x63, 0xd5, 0x91, 0x05, 0xd9, 0xce, 0x2f, 0x62, 0xe1, 0xda, 0x96, 0xfd, 0x07,
	0x41, 0x4f, 0x01, 0xb7, 0x62, 0xec, 0xc8, 0x88, 0x92, 0x8b, 0x20, 0x99, 0x16, 0xba, 0xbd, 0xef,
	0xbd, 0x6f, 0x79, 0x74, 0x6e, 0x82, 0xf5, 0x98, 0x7a, 0xff, 0x2c, 0x63, 0xaf, 0x9d, 0xcd, 0x1d,
	0xa4, 0xc7, 0x98, 0x42, 0x0c, 0x59, 0xb9, 0x26, 0xa6, 0x80, 0x81, 0x5d, 0xe8, 0x0d, 0x34, 0x7f,
	0xde, 0xfc, 0x83, 0xd0, 0xcb, 0xe5, 0xde, 0x5d, 0xee, 0x54, 0x76, 0x45, 0xa7, 0x68, 0xd1, 0x01,
	0x27, 0x35, 0x59, 0x9c, 0xad, 0x7e, 0x81, 0xd5, 0xf4, 0xbc, 0x85, 0x6c, 0x92, 0x8d, 0x68, 0x83,
	0xe7, 0x47, 0x65, 0x3b, 0xac, 0x18, 0xa3, 0x13, 0xaf, 0xd6, 0xc0, 0x8f, 0xcb, 0x54, 0x32, 0xe3,
	0xf4, 0x44, 0xb5, 0x6d, 0x82, 0x9c, 0xf9, 0xa4, 0xd4, 0x7b, 0x64, 0xd7, 0x74, 0xa6, 0x0c, 0xda,
	0x17, 0xe0, 0xd3, 0x9a, 0x2c, 0x4e, 0x57, 0x3b, 0xba, 0xbb, 0xff, 0x1c, 0x04, 0xd9, 0x0e, 0x82,
	0x7c, 0x0f, 0x82, 0xbc, 0x8f, 0xa2, 0xda, 0x8e, 0xa2, 0xfa, 0x1a, 0x45, 0xf5, 0xd0, 0x3c, 0x59,
	0xec, 0x7a, 0xdd, 0x98, 0xb0, 0x96, 0x7a, 0x03, 0x37, 0xca, 0xc5, 0x4e, 0x21, 0xa8, 0x42, 0xf2,
	0x55, 0xfe, 0xff, 0xc7, 0xb7, 0x08, 0x59, 0xcf, 0xca, 0xe7, 0xdb, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x77, 0x79, 0x0e, 0xad, 0x19, 0x01, 0x00, 0x00,
}

func (m *PublisherProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublisherProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublisherProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPublisherProposal(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPublisherProposal(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintPublisherProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintPublisherProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPublisherProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovPublisherProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PublisherProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovPublisherProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovPublisherProposal(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPublisherProposal(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPublisherProposal(uint64(l))
	}
	if m.Active {
		n += 2
	}
	return n
}

func sovPublisherProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPublisherProposal(x uint64) (n int) {
	return sovPublisherProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PublisherProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublisherProposal
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
			return fmt.Errorf("proto: PublisherProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PublisherProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublisherProposal
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
				return ErrInvalidLengthPublisherProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublisherProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublisherProposal
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
				return ErrInvalidLengthPublisherProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublisherProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublisherProposal
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
				return ErrInvalidLengthPublisherProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublisherProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublisherProposal
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
				return ErrInvalidLengthPublisherProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublisherProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublisherProposal
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
			m.Active = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPublisherProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPublisherProposal
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
func skipPublisherProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPublisherProposal
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
					return 0, ErrIntOverflowPublisherProposal
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
					return 0, ErrIntOverflowPublisherProposal
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
				return 0, ErrInvalidLengthPublisherProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPublisherProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPublisherProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPublisherProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPublisherProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPublisherProposal = fmt.Errorf("proto: unexpected end of group")
)
