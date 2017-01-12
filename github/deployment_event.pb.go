// Code generated by protoc-gen-gogo.
// source: deployment_event.proto
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

type DeploymentEvent struct {
	Deployment   *Deployment   `protobuf:"bytes,1,opt,name=deployment" json:"deployment,omitempty"`
	Repository   *Repository   `protobuf:"bytes,2,opt,name=repository" json:"repository,omitempty"`
	Sender       *User         `protobuf:"bytes,3,opt,name=sender" json:"sender,omitempty"`
	Installation *Installation `protobuf:"bytes,4,opt,name=installation" json:"installation,omitempty"`
	Organization *User         `protobuf:"bytes,5,opt,name=organization" json:"organization,omitempty"`
}

func (m *DeploymentEvent) Reset()                    { *m = DeploymentEvent{} }
func (m *DeploymentEvent) String() string            { return proto.CompactTextString(m) }
func (*DeploymentEvent) ProtoMessage()               {}
func (*DeploymentEvent) Descriptor() ([]byte, []int) { return fileDescriptorDeploymentEvent, []int{0} }

func init() {
	proto.RegisterType((*DeploymentEvent)(nil), "github.DeploymentEvent")
}
func (m *DeploymentEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeploymentEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Deployment != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDeploymentEvent(dAtA, i, uint64(m.Deployment.Size()))
		n1, err := m.Deployment.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Repository != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDeploymentEvent(dAtA, i, uint64(m.Repository.Size()))
		n2, err := m.Repository.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Sender != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintDeploymentEvent(dAtA, i, uint64(m.Sender.Size()))
		n3, err := m.Sender.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Installation != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintDeploymentEvent(dAtA, i, uint64(m.Installation.Size()))
		n4, err := m.Installation.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Organization != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintDeploymentEvent(dAtA, i, uint64(m.Organization.Size()))
		n5, err := m.Organization.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}

func encodeFixed64DeploymentEvent(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32DeploymentEvent(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintDeploymentEvent(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DeploymentEvent) Size() (n int) {
	var l int
	_ = l
	if m.Deployment != nil {
		l = m.Deployment.Size()
		n += 1 + l + sovDeploymentEvent(uint64(l))
	}
	if m.Repository != nil {
		l = m.Repository.Size()
		n += 1 + l + sovDeploymentEvent(uint64(l))
	}
	if m.Sender != nil {
		l = m.Sender.Size()
		n += 1 + l + sovDeploymentEvent(uint64(l))
	}
	if m.Installation != nil {
		l = m.Installation.Size()
		n += 1 + l + sovDeploymentEvent(uint64(l))
	}
	if m.Organization != nil {
		l = m.Organization.Size()
		n += 1 + l + sovDeploymentEvent(uint64(l))
	}
	return n
}

func sovDeploymentEvent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDeploymentEvent(x uint64) (n int) {
	return sovDeploymentEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeploymentEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeploymentEvent
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
			return fmt.Errorf("proto: DeploymentEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeploymentEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deployment", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploymentEvent
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
				return ErrInvalidLengthDeploymentEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Deployment == nil {
				m.Deployment = &Deployment{}
			}
			if err := m.Deployment.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repository", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploymentEvent
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
				return ErrInvalidLengthDeploymentEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Repository == nil {
				m.Repository = &Repository{}
			}
			if err := m.Repository.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploymentEvent
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
				return ErrInvalidLengthDeploymentEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Sender == nil {
				m.Sender = &User{}
			}
			if err := m.Sender.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Installation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploymentEvent
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
				return ErrInvalidLengthDeploymentEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Installation == nil {
				m.Installation = &Installation{}
			}
			if err := m.Installation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Organization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploymentEvent
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
				return ErrInvalidLengthDeploymentEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Organization == nil {
				m.Organization = &User{}
			}
			if err := m.Organization.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeploymentEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeploymentEvent
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
func skipDeploymentEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeploymentEvent
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
					return 0, ErrIntOverflowDeploymentEvent
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
					return 0, ErrIntOverflowDeploymentEvent
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
				return 0, ErrInvalidLengthDeploymentEvent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDeploymentEvent
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
				next, err := skipDeploymentEvent(dAtA[start:])
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
	ErrInvalidLengthDeploymentEvent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeploymentEvent   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("deployment_event.proto", fileDescriptorDeploymentEvent) }

var fileDescriptorDeploymentEvent = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x49, 0x2d, 0xc8,
	0xc9, 0xaf, 0xcc, 0x4d, 0xcd, 0x2b, 0x89, 0x4f, 0x2d, 0x03, 0x92, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0x6c, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x52, 0xba, 0x10, 0x5a, 0x2f, 0x39, 0x3f,
	0x57, 0x3f, 0x3d, 0x3f, 0x3d, 0x5f, 0x1f, 0x2c, 0x9d, 0x54, 0x9a, 0x06, 0xe6, 0x81, 0x39, 0x60,
	0x16, 0x44, 0x9b, 0x14, 0x57, 0x69, 0x71, 0x6a, 0x11, 0x94, 0x2d, 0x50, 0x94, 0x5a, 0x90, 0x5f,
	0x9c, 0x59, 0x92, 0x5f, 0x54, 0x09, 0x13, 0x41, 0x58, 0x06, 0x15, 0x11, 0xca, 0xcc, 0x2b, 0x2e,
	0x49, 0xcc, 0xc9, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x83, 0x88, 0x29, 0x75, 0x33, 0x71, 0xf1, 0xbb,
	0xc0, 0x15, 0xba, 0x82, 0x1c, 0x25, 0x64, 0xc4, 0xc5, 0x85, 0xd0, 0x2b, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x6d, 0x24, 0xa4, 0x07, 0x75, 0x1b, 0x42, 0x71, 0x10, 0x92, 0x2a, 0x90, 0x1e, 0x84, 0x0b,
	0x24, 0x98, 0x50, 0xf5, 0x04, 0xc1, 0x65, 0x82, 0x90, 0x54, 0x09, 0xa9, 0x70, 0xb1, 0x15, 0xa7,
	0xe6, 0xa5, 0xa4, 0x16, 0x49, 0x30, 0x83, 0xd5, 0xf3, 0xc0, 0xd4, 0x87, 0x02, 0xfd, 0x15, 0x04,
	0x95, 0x13, 0xb2, 0xe0, 0xe2, 0x41, 0x76, 0xb7, 0x04, 0x0b, 0x58, 0xad, 0x08, 0x4c, 0xad, 0x27,
	0x92, 0x5c, 0x10, 0x8a, 0x4a, 0x21, 0x03, 0x2e, 0x9e, 0xfc, 0xa2, 0xf4, 0xc4, 0xbc, 0xcc, 0x2a,
	0x88, 0x4e, 0x56, 0x2c, 0xb6, 0xa0, 0xa8, 0x70, 0x12, 0x39, 0xf1, 0x50, 0x8e, 0xe1, 0xc4, 0x23,
	0x39, 0xc6, 0x0b, 0x40, 0xfc, 0x00, 0x88, 0x67, 0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0x07, 0x95,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xde, 0xe2, 0xd2, 0x94, 0xbf, 0x01, 0x00, 0x00,
}
