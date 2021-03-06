// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: listOpts.proto

/*
	Package types is a generated protocol buffer package.

	It is generated from these files:
		listOpts.proto
		uuid.proto
		nil.proto
		atTimestamps.proto

	It has these top-level messages:
		ListOpts
		UUID
		Nil
		CreatedAt
		UpdatedAt
		ArchivedAt
		DeletedAt
		Timestamps
*/
package types

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ListOpts struct {
	// limit specifies the maximum number of results to be returned by the
	// server. The server may further constrain the maximum number of results
	// returned in a single page. If the limit is 0, the server will decide the
	// number of results to be returned.
	Limit int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// offset is the number of results that will be offset
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// The string value should follow SQL syntax: comma separated list of fields.
	// For example: "foo,bar". The default sorting order is ascending. To specify
	// descending order for a field, a suffix " desc" should be appended to the
	// field name. For example: "foo desc,bar".
	// Redundant space characters in the syntax are insignificant. "foo,bar desc" and "  foo ,  bar  desc  " are equivalent.
	OrderBy string `protobuf:"bytes,3,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
}

func (m *ListOpts) Reset()                    { *m = ListOpts{} }
func (m *ListOpts) String() string            { return proto.CompactTextString(m) }
func (*ListOpts) ProtoMessage()               {}
func (*ListOpts) Descriptor() ([]byte, []int) { return fileDescriptorListOpts, []int{0} }

func (m *ListOpts) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListOpts) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListOpts) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

func init() {
	proto.RegisterType((*ListOpts)(nil), "canaryhealth.protobuf.ListOpts")
}
func (m *ListOpts) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListOpts) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Limit != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintListOpts(dAtA, i, uint64(m.Limit))
	}
	if m.Offset != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintListOpts(dAtA, i, uint64(m.Offset))
	}
	if len(m.OrderBy) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintListOpts(dAtA, i, uint64(len(m.OrderBy)))
		i += copy(dAtA[i:], m.OrderBy)
	}
	return i, nil
}

func encodeVarintListOpts(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ListOpts) Size() (n int) {
	var l int
	_ = l
	if m.Limit != 0 {
		n += 1 + sovListOpts(uint64(m.Limit))
	}
	if m.Offset != 0 {
		n += 1 + sovListOpts(uint64(m.Offset))
	}
	l = len(m.OrderBy)
	if l > 0 {
		n += 1 + l + sovListOpts(uint64(l))
	}
	return n
}

func sovListOpts(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozListOpts(x uint64) (n int) {
	return sovListOpts(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ListOpts) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListOpts
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
			return fmt.Errorf("proto: ListOpts: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListOpts: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListOpts
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListOpts
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListOpts
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
				return ErrInvalidLengthListOpts
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListOpts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthListOpts
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
func skipListOpts(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowListOpts
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
					return 0, ErrIntOverflowListOpts
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
					return 0, ErrIntOverflowListOpts
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
				return 0, ErrInvalidLengthListOpts
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowListOpts
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
				next, err := skipListOpts(dAtA[start:])
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
	ErrInvalidLengthListOpts = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowListOpts   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("listOpts.proto", fileDescriptorListOpts) }

var fileDescriptorListOpts = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xc9, 0x2c, 0x2e,
	0xf1, 0x2f, 0x28, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4d, 0x4e, 0xcc, 0x4b,
	0x2c, 0xaa, 0xcc, 0x48, 0x4d, 0xcc, 0x29, 0xc9, 0x80, 0x88, 0x25, 0x95, 0xa6, 0x29, 0x05, 0x73,
	0x71, 0xf8, 0x40, 0x15, 0x0a, 0x89, 0x70, 0xb1, 0xe6, 0x64, 0xe6, 0x66, 0x96, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0xb0, 0x06, 0x41, 0x38, 0x42, 0x62, 0x5c, 0x6c, 0xf9, 0x69, 0x69, 0xc5, 0xa9, 0x25,
	0x12, 0x4c, 0x60, 0x61, 0x28, 0x4f, 0x48, 0x92, 0x8b, 0x23, 0xbf, 0x28, 0x25, 0xb5, 0x28, 0x3e,
	0xa9, 0x52, 0x82, 0x59, 0x81, 0x51, 0x83, 0x33, 0x88, 0x1d, 0xcc, 0x77, 0xaa, 0x74, 0x12, 0x3f,
	0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63,
	0x88, 0x62, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xdb, 0x6b, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xfa, 0x13, 0x6e, 0x83, 0x9d, 0x00, 0x00, 0x00,
}
