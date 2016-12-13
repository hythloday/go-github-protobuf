// Code generated by protoc-gen-gogo.
// source: public_event.proto
// DO NOT EDIT!

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

type PublicEvent struct {
	Repository   *Repository   `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	Sender       *User         `protobuf:"bytes,2,opt,name=sender" json:"sender,omitempty"`
	Installation *Installation `protobuf:"bytes,3,opt,name=installation" json:"installation,omitempty"`
	Organization *User         `protobuf:"bytes,4,opt,name=organization" json:"organization,omitempty"`
}

func (m *PublicEvent) Reset()                    { *m = PublicEvent{} }
func (m *PublicEvent) String() string            { return proto.CompactTextString(m) }
func (*PublicEvent) ProtoMessage()               {}
func (*PublicEvent) Descriptor() ([]byte, []int) { return fileDescriptorPublicEvent, []int{0} }

func init() {
	proto.RegisterType((*PublicEvent)(nil), "github.PublicEvent")
}
func (m *PublicEvent) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PublicEvent) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Repository != nil {
		data[i] = 0xa
		i++
		i = encodeVarintPublicEvent(data, i, uint64(m.Repository.Size()))
		n1, err := m.Repository.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Sender != nil {
		data[i] = 0x12
		i++
		i = encodeVarintPublicEvent(data, i, uint64(m.Sender.Size()))
		n2, err := m.Sender.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Installation != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintPublicEvent(data, i, uint64(m.Installation.Size()))
		n3, err := m.Installation.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Organization != nil {
		data[i] = 0x22
		i++
		i = encodeVarintPublicEvent(data, i, uint64(m.Organization.Size()))
		n4, err := m.Organization.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func encodeFixed64PublicEvent(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32PublicEvent(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintPublicEvent(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *PublicEvent) Size() (n int) {
	var l int
	_ = l
	if m.Repository != nil {
		l = m.Repository.Size()
		n += 1 + l + sovPublicEvent(uint64(l))
	}
	if m.Sender != nil {
		l = m.Sender.Size()
		n += 1 + l + sovPublicEvent(uint64(l))
	}
	if m.Installation != nil {
		l = m.Installation.Size()
		n += 1 + l + sovPublicEvent(uint64(l))
	}
	if m.Organization != nil {
		l = m.Organization.Size()
		n += 1 + l + sovPublicEvent(uint64(l))
	}
	return n
}

func sovPublicEvent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPublicEvent(x uint64) (n int) {
	return sovPublicEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PublicEvent) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PublicEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PublicEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repository", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPublicEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Repository == nil {
				m.Repository = &Repository{}
			}
			if err := m.Repository.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPublicEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Sender == nil {
				m.Sender = &User{}
			}
			if err := m.Sender.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Installation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPublicEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Installation == nil {
				m.Installation = &Installation{}
			}
			if err := m.Installation.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Organization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPublicEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Organization == nil {
				m.Organization = &User{}
			}
			if err := m.Organization.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicEvent(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicEvent
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
func skipPublicEvent(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPublicEvent
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowPublicEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowPublicEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPublicEvent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPublicEvent
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipPublicEvent(data[start:])
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
	ErrInvalidLengthPublicEvent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPublicEvent   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("public_event.proto", fileDescriptorPublicEvent) }

var fileDescriptorPublicEvent = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0x4d, 0xca,
	0xc9, 0x4c, 0x8e, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0x92, 0xd2, 0x85, 0xd0, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9,
	0xf9, 0xe9, 0xf9, 0xfa, 0x60, 0xe9, 0xa4, 0xd2, 0x34, 0x30, 0x0f, 0xcc, 0x01, 0xb3, 0x20, 0xda,
	0xa4, 0xb8, 0x4a, 0x8b, 0x53, 0x8b, 0xa0, 0x6c, 0x81, 0xa2, 0xd4, 0x82, 0xfc, 0xe2, 0xcc, 0x92,
	0xfc, 0xa2, 0x4a, 0xa8, 0x88, 0x50, 0x66, 0x5e, 0x71, 0x49, 0x62, 0x4e, 0x4e, 0x62, 0x49, 0x66,
	0x7e, 0x1e, 0x44, 0x4c, 0xe9, 0x32, 0x23, 0x17, 0x77, 0x00, 0xd8, 0x7e, 0x57, 0x90, 0xf5, 0x42,
	0x46, 0x5c, 0x5c, 0x08, 0x7d, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x42, 0x7a, 0x50, 0x57,
	0x04, 0xc1, 0x65, 0x82, 0x90, 0x54, 0x09, 0xa9, 0x70, 0xb1, 0x15, 0xa7, 0xe6, 0xa5, 0xa4, 0x16,
	0x49, 0x30, 0x81, 0xd5, 0xf3, 0xc0, 0xd4, 0x87, 0x02, 0x5d, 0x13, 0x04, 0x95, 0x13, 0xb2, 0xe0,
	0xe2, 0x41, 0xb6, 0x5f, 0x82, 0x19, 0xac, 0x56, 0x04, 0xa6, 0xd6, 0x13, 0x49, 0x2e, 0x08, 0x45,
	0xa5, 0x90, 0x01, 0x17, 0x4f, 0x7e, 0x51, 0x7a, 0x62, 0x5e, 0x66, 0x15, 0x44, 0x27, 0x0b, 0x16,
	0x5b, 0x50, 0x54, 0x38, 0x89, 0x9c, 0x78, 0x28, 0xc7, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x05, 0x20,
	0x7e, 0x00, 0xc4, 0x33, 0x1e, 0xcb, 0x31, 0x24, 0xb1, 0x81, 0xbd, 0x6c, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xe4, 0x2e, 0xf1, 0x3c, 0x71, 0x01, 0x00, 0x00,
}
