// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: plugin.proto

/*
Package runtime is a generated protocol buffer package.

It is generated from these files:

	plugin.proto

It has these top-level messages:

	PluginSpec
	PluginPrivilege
*/
package runtime

import (
	fmt "fmt"
	"time"

	"github.com/aquasecurity/trivy/pkg/bug"
	proto "github.com/gogo/protobuf/proto"

	math "math"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// PluginSpec defines the base payload which clients can specify for creating
// a service with the plugin runtime.
type PluginSpec struct {
	Name       string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Remote     string             `protobuf:"bytes,2,opt,name=remote,proto3" json:"remote,omitempty"`
	Privileges []*PluginPrivilege `protobuf:"bytes,3,rep,name=privileges" json:"privileges,omitempty"`
	Disabled   bool               `protobuf:"varint,4,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Env        []string           `protobuf:"bytes,5,rep,name=env" json:"env,omitempty"`
}

func (m *PluginSpec) Reset()                    { *m = PluginSpec{} }
func (m *PluginSpec) String() string            { return proto.CompactTextString(m) }
func (*PluginSpec) ProtoMessage()               {}
func (*PluginSpec) Descriptor() ([]byte, []int) { return fileDescriptorPlugin, []int{0} }

func (m *PluginSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PluginSpec) GetRemote() string {
	if m != nil {
		return m.Remote
	}
	return ""
}

func (m *PluginSpec) GetPrivileges() []*PluginPrivilege {
	if m != nil {
		return m.Privileges
	}
	return nil
}

func (m *PluginSpec) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *PluginSpec) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

// PluginPrivilege describes a permission the user has to accept
// upon installing a plugin.
type PluginPrivilege struct {
	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Value       []string `protobuf:"bytes,3,rep,name=value" json:"value,omitempty"`
}

func (m *PluginPrivilege) Reset()                    { *m = PluginPrivilege{} }
func (m *PluginPrivilege) String() string            { return proto.CompactTextString(m) }
func (*PluginPrivilege) ProtoMessage()               {}
func (*PluginPrivilege) Descriptor() ([]byte, []int) { return fileDescriptorPlugin, []int{1} }

func (m *PluginPrivilege) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PluginPrivilege) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PluginPrivilege) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	defer func(start time.Time) { bug.PrintCustomStack(start) }(time.Now())
	proto.RegisterType((*PluginSpec)(nil), "PluginSpec")
	proto.RegisterType((*PluginPrivilege)(nil), "PluginPrivilege")
}
func (m *PluginSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PluginSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPlugin(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Remote) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPlugin(dAtA, i, uint64(len(m.Remote)))
		i += copy(dAtA[i:], m.Remote)
	}
	if len(m.Privileges) > 0 {
		for _, msg := range m.Privileges {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintPlugin(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Disabled {
		dAtA[i] = 0x20
		i++
		if m.Disabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Env) > 0 {
		for _, s := range m.Env {
			dAtA[i] = 0x2a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *PluginPrivilege) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PluginPrivilege) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPlugin(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPlugin(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if len(m.Value) > 0 {
		for _, s := range m.Value {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintPlugin(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PluginSpec) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPlugin(uint64(l))
	}
	l = len(m.Remote)
	if l > 0 {
		n += 1 + l + sovPlugin(uint64(l))
	}
	if len(m.Privileges) > 0 {
		for _, e := range m.Privileges {
			l = e.Size()
			n += 1 + l + sovPlugin(uint64(l))
		}
	}
	if m.Disabled {
		n += 2
	}
	if len(m.Env) > 0 {
		for _, s := range m.Env {
			l = len(s)
			n += 1 + l + sovPlugin(uint64(l))
		}
	}
	return n
}

func (m *PluginPrivilege) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPlugin(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovPlugin(uint64(l))
	}
	if len(m.Value) > 0 {
		for _, s := range m.Value {
			l = len(s)
			n += 1 + l + sovPlugin(uint64(l))
		}
	}
	return n
}

func sovPlugin(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPlugin(x uint64) (n int) {
	return sovPlugin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PluginSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlugin
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
			return fmt.Errorf("proto: PluginSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PluginSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Remote", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Remote = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Privileges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Privileges = append(m.Privileges, &PluginPrivilege{})
			if err := m.Privileges[len(m.Privileges)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Disabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
			m.Disabled = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Env", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Env = append(m.Env, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlugin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlugin
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
func (m *PluginPrivilege) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlugin
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
			return fmt.Errorf("proto: PluginPrivilege: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PluginPrivilege: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlugin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlugin
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
func skipPlugin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlugin
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
					return 0, ErrIntOverflowPlugin
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
					return 0, ErrIntOverflowPlugin
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
				return 0, ErrInvalidLengthPlugin
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPlugin
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
				next, err := skipPlugin(dAtA[start:])
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
	ErrInvalidLengthPlugin = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlugin   = fmt.Errorf("proto: integer overflow")
)

func init() {
	defer func(start time.Time) { bug.PrintCustomStack(start) }(time.Now())
	proto.RegisterFile("plugin.proto", fileDescriptorPlugin)
}

var fileDescriptorPlugin = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4d, 0x4b, 0xc3, 0x30,
	0x18, 0xc7, 0x89, 0xdd, 0xc6, 0xfa, 0x4c, 0x70, 0x04, 0x91, 0xe2, 0xa1, 0x94, 0x9d, 0x7a, 0x6a,
	0x45, 0x2f, 0x82, 0x37, 0x0f, 0x9e, 0x47, 0xbc, 0x09, 0x1e, 0xd2, 0xf6, 0xa1, 0x06, 0x9b, 0x17,
	0x92, 0xb4, 0xe2, 0x37, 0xf1, 0x23, 0x79, 0xf4, 0x23, 0x48, 0x3f, 0x89, 0x98, 0x75, 0x32, 0x64,
	0xa7, 0xff, 0x4b, 0xc2, 0x9f, 0x1f, 0x0f, 0x9c, 0x9a, 0xae, 0x6f, 0x85, 0x2a, 0x8c, 0xd5, 0x5e,
	0x6f, 0x3e, 0x08, 0xc0, 0x36, 0x14, 0x8f, 0x06, 0x6b, 0x4a, 0x61, 0xa6, 0xb8, 0xc4, 0x84, 0x64,
	0x24, 0x8f, 0x59, 0xf0, 0xf4, 0x02, 0x16, 0x16, 0xa5, 0xf6, 0x98, 0x9c, 0x84, 0x76, 0x4a, 0xf4,
	0x0a, 0xc0, 0x58, 0x31, 0x88, 0x0e, 0x5b, 0x74, 0x49, 0x94, 0x45, 0xf9, 0xea, 0x7a, 0x5d, 0xec,
	0xc6, 0xb6, 0xfb, 0x07, 0x76, 0xf0, 0x87, 0x5e, 0xc2, 0xb2, 0x11, 0x8e, 0x57, 0x1d, 0x36, 0xc9,
	0x2c, 0x23, 0xf9, 0x92, 0xfd, 0x65, 0xba, 0x86, 0x08, 0xd5, 0x90, 0xcc, 0xb3, 0x28, 0x8f, 0xd9,
	0xaf, 0xdd, 0x3c, 0xc3, 0xd9, 0xbf, 0xb1, 0xa3, 0x78, 0x19, 0xac, 0x1a, 0x74, 0xb5, 0x15, 0xc6,
	0x0b, 0xad, 0x26, 0xc6, 0xc3, 0x8a, 0x9e, 0xc3, 0x7c, 0xe0, 0x5d, 0x8f, 0x81, 0x31, 0x66, 0xbb,
	0x70, 0xff, 0xf0, 0x39, 0xa6, 0xe4, 0x6b, 0x4c, 0xc9, 0xf7, 0x98, 0x92, 0xa7, 0xdb, 0x56, 0xf8,
	0x97, 0xbe, 0x2a, 0x6a, 0x2d, 0xcb, 0x46, 0xd7, 0xaf, 0x68, 0xf7, 0xc2, 0x8d, 0x28, 0xfd, 0xbb,
	0x41, 0x57, 0xba, 0x37, 0x6e, 0x65, 0x69, 0x7b, 0xe5, 0x85, 0xc4, 0xbb, 0x49, 0xab, 0x45, 0x38,
	0xe4, 0xcd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0xa8, 0xd9, 0x9b, 0x58, 0x01, 0x00, 0x00,
}
