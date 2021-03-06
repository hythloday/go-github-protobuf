// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: installation.proto

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Installation struct {
	Id              int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Account         *User  `protobuf:"bytes,2,opt,name=account" json:"account,omitempty"`
	AccessTokensUrl string `protobuf:"bytes,3,opt,name=access_tokens_url,json=accessTokensUrl,proto3" json:"access_tokens_url,omitempty"`
	RepositoriesUrl string `protobuf:"bytes,4,opt,name=repositories_url,json=repositoriesUrl,proto3" json:"repositories_url,omitempty"`
	HtmlUrl         string `protobuf:"bytes,5,opt,name=html_url,json=htmlUrl,proto3" json:"html_url,omitempty"`
}

func (m *Installation) Reset()                    { *m = Installation{} }
func (m *Installation) String() string            { return proto.CompactTextString(m) }
func (*Installation) ProtoMessage()               {}
func (*Installation) Descriptor() ([]byte, []int) { return fileDescriptorInstallation, []int{0} }

func init() {
	proto.RegisterType((*Installation)(nil), "github.Installation")
}
func (m *Installation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Installation) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintInstallation(dAtA, i, uint64(m.Id))
	}
	if m.Account != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintInstallation(dAtA, i, uint64(m.Account.Size()))
		n1, err := m.Account.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.AccessTokensUrl) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintInstallation(dAtA, i, uint64(len(m.AccessTokensUrl)))
		i += copy(dAtA[i:], m.AccessTokensUrl)
	}
	if len(m.RepositoriesUrl) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintInstallation(dAtA, i, uint64(len(m.RepositoriesUrl)))
		i += copy(dAtA[i:], m.RepositoriesUrl)
	}
	if len(m.HtmlUrl) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintInstallation(dAtA, i, uint64(len(m.HtmlUrl)))
		i += copy(dAtA[i:], m.HtmlUrl)
	}
	return i, nil
}

func encodeFixed64Installation(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Installation(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintInstallation(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Installation) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovInstallation(uint64(m.Id))
	}
	if m.Account != nil {
		l = m.Account.Size()
		n += 1 + l + sovInstallation(uint64(l))
	}
	l = len(m.AccessTokensUrl)
	if l > 0 {
		n += 1 + l + sovInstallation(uint64(l))
	}
	l = len(m.RepositoriesUrl)
	if l > 0 {
		n += 1 + l + sovInstallation(uint64(l))
	}
	l = len(m.HtmlUrl)
	if l > 0 {
		n += 1 + l + sovInstallation(uint64(l))
	}
	return n
}

func sovInstallation(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozInstallation(x uint64) (n int) {
	return sovInstallation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Installation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInstallation
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
			return fmt.Errorf("proto: Installation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Installation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInstallation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInstallation
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
				return ErrInvalidLengthInstallation
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Account == nil {
				m.Account = &User{}
			}
			if err := m.Account.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessTokensUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInstallation
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
				return ErrInvalidLengthInstallation
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessTokensUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RepositoriesUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInstallation
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
				return ErrInvalidLengthInstallation
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RepositoriesUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HtmlUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInstallation
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
				return ErrInvalidLengthInstallation
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HtmlUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInstallation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInstallation
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
func skipInstallation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInstallation
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
					return 0, ErrIntOverflowInstallation
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
					return 0, ErrIntOverflowInstallation
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
				return 0, ErrInvalidLengthInstallation
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowInstallation
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
				next, err := skipInstallation(dAtA[start:])
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
	ErrInvalidLengthInstallation = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInstallation   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("installation.proto", fileDescriptorInstallation) }

var fileDescriptorInstallation = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x4b, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xeb, 0x40, 0x5b, 0x30, 0x15, 0x0f, 0x8b, 0x45, 0xe8, 0xc2, 0x8a, 0x58, 0xa0, 0x80,
	0x44, 0x2a, 0xc1, 0x0d, 0xd8, 0xb1, 0x8d, 0xc8, 0xba, 0x4a, 0x5c, 0x93, 0x5a, 0xb8, 0x99, 0xca,
	0x1e, 0xdf, 0x85, 0xa3, 0x70, 0x84, 0x2e, 0x39, 0x02, 0x84, 0x8b, 0xa0, 0x8e, 0x41, 0xca, 0xca,
	0xff, 0xe3, 0xf3, 0x8c, 0x86, 0x0b, 0xd3, 0x79, 0xac, 0xad, 0xad, 0xd1, 0x40, 0x57, 0x6c, 0x1d,
	0x20, 0x88, 0x49, 0x6b, 0x70, 0x1d, 0x9a, 0xf9, 0x7d, 0x7c, 0x0b, 0x05, 0x9b, 0x45, 0x0b, 0x2d,
	0x2c, 0xa8, 0x6e, 0xc2, 0x2b, 0x39, 0x32, 0xa4, 0xe2, 0xb7, 0x39, 0x0f, 0x5e, 0xbb, 0xa8, 0xaf,
	0x3f, 0x18, 0x9f, 0x3d, 0x0f, 0x26, 0x8b, 0x53, 0x9e, 0x98, 0x55, 0xca, 0x32, 0x96, 0x8f, 0xcb,
	0xc4, 0xac, 0xc4, 0x0d, 0x9f, 0xd6, 0x4a, 0x41, 0xe8, 0x30, 0x4d, 0x32, 0x96, 0x9f, 0x3c, 0xcc,
	0x8a, 0xbf, 0x6d, 0x95, 0xd7, 0xae, 0xfc, 0x2f, 0xc5, 0x1d, 0xbf, 0xa8, 0x95, 0xd2, 0xde, 0x2f,
	0x11, 0xde, 0x74, 0xe7, 0x97, 0xc1, 0xd9, 0xf4, 0x20, 0x63, 0xf9, 0x71, 0x79, 0x16, 0x8b, 0x17,
	0xca, 0x2b, 0x67, 0xc5, 0x2d, 0x3f, 0x77, 0x7a, 0x0b, 0xde, 0x20, 0x38, 0xa3, 0x23, 0x7a, 0x18,
	0xd1, 0x61, 0xbe, 0x47, 0xaf, 0xf8, 0xd1, 0x1a, 0x37, 0x96, 0x90, 0x31, 0x21, 0xd3, 0xbd, 0xaf,
	0x9c, 0x7d, 0xba, 0xdc, 0x7d, 0xcb, 0xd1, 0xae, 0x97, 0xec, 0xb3, 0x97, 0xec, 0xab, 0x97, 0xec,
	0xfd, 0x47, 0x8e, 0x9a, 0x09, 0xdd, 0xf5, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x08, 0x29, 0x0d,
	0xaa, 0x30, 0x01, 0x00, 0x00,
}
