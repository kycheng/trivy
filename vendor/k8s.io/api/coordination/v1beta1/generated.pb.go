/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/kubernetes/vendor/k8s.io/api/coordination/v1beta1/generated.proto

package v1beta1

import (
	fmt "fmt"
	"time"

	io "io"

	"github.com/aquasecurity/trivy/pkg/bug"
	proto "github.com/gogo/protobuf/proto"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

func (m *Lease) Reset()      { *m = Lease{} }
func (*Lease) ProtoMessage() {}
func (*Lease) Descriptor() ([]byte, []int) {
	return fileDescriptor_daca6bcd2ff63a80, []int{0}
}
func (m *Lease) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Lease) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Lease) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lease.Merge(m, src)
}
func (m *Lease) XXX_Size() int {
	return m.Size()
}
func (m *Lease) XXX_DiscardUnknown() {
	xxx_messageInfo_Lease.DiscardUnknown(m)
}

var xxx_messageInfo_Lease proto.InternalMessageInfo

func (m *LeaseList) Reset()      { *m = LeaseList{} }
func (*LeaseList) ProtoMessage() {}
func (*LeaseList) Descriptor() ([]byte, []int) {
	return fileDescriptor_daca6bcd2ff63a80, []int{1}
}
func (m *LeaseList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LeaseList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *LeaseList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseList.Merge(m, src)
}
func (m *LeaseList) XXX_Size() int {
	return m.Size()
}
func (m *LeaseList) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseList.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseList proto.InternalMessageInfo

func (m *LeaseSpec) Reset()      { *m = LeaseSpec{} }
func (*LeaseSpec) ProtoMessage() {}
func (*LeaseSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_daca6bcd2ff63a80, []int{2}
}
func (m *LeaseSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LeaseSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *LeaseSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseSpec.Merge(m, src)
}
func (m *LeaseSpec) XXX_Size() int {
	return m.Size()
}
func (m *LeaseSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseSpec.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseSpec proto.InternalMessageInfo

func init() {
	defer func(start time.Time) { bug.PrintCustomStack(start) }(time.Now())
	proto.RegisterType((*Lease)(nil), "k8s.io.api.coordination.v1beta1.Lease")
	proto.RegisterType((*LeaseList)(nil), "k8s.io.api.coordination.v1beta1.LeaseList")
	proto.RegisterType((*LeaseSpec)(nil), "k8s.io.api.coordination.v1beta1.LeaseSpec")
}

func init() {
	defer func(start time.Time) { bug.PrintCustomStack(start) }(time.Now())
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/api/coordination/v1beta1/generated.proto", fileDescriptor_daca6bcd2ff63a80)
}

var fileDescriptor_daca6bcd2ff63a80 = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xc1, 0x6e, 0xd3, 0x4e,
	0x10, 0xc6, 0xe3, 0xb6, 0x91, 0x9a, 0xcd, 0xbf, 0xfd, 0x47, 0x56, 0x0e, 0x56, 0x0e, 0x76, 0x95,
	0x03, 0xaa, 0x90, 0xd8, 0x25, 0x15, 0x42, 0x88, 0x13, 0x58, 0x20, 0xb5, 0xc2, 0x15, 0x92, 0xdb,
	0x13, 0xea, 0x81, 0xb5, 0x3d, 0x38, 0x4b, 0x6a, 0xaf, 0xd9, 0x5d, 0x07, 0xf5, 0xc6, 0x23, 0x70,
	0xe5, 0x45, 0xe0, 0x15, 0x72, 0xec, 0xb1, 0x27, 0x8b, 0x98, 0x17, 0x41, 0xde, 0xb8, 0x4d, 0x48,
	0x8a, 0x12, 0x71, 0xf3, 0xce, 0xcc, 0xf7, 0x9b, 0x6f, 0xbe, 0x83, 0xd1, 0xf1, 0xe8, 0x99, 0xc4,
	0x8c, 0x93, 0x51, 0x1e, 0x80, 0x48, 0x41, 0x81, 0x24, 0x63, 0x48, 0x23, 0x2e, 0x48, 0xdd, 0xa0,
	0x19, 0x23, 0x21, 0xe7, 0x22, 0x62, 0x29, 0x55, 0x8c, 0xa7, 0x64, 0x3c, 0x08, 0x40, 0xd1, 0x01,
	0x89, 0x21, 0x05, 0x41, 0x15, 0x44, 0x38, 0x13, 0x5c, 0x71, 0xd3, 0x99, 0x09, 0x30, 0xcd, 0x18,
	0x5e, 0x14, 0xe0, 0x5a, 0xd0, 0x7b, 0x14, 0x33, 0x35, 0xcc, 0x03, 0x1c, 0xf2, 0x84, 0xc4, 0x3c,
	0xe6, 0x44, 0xeb, 0x82, 0xfc, 0x83, 0x7e, 0xe9, 0x87, 0xfe, 0x9a, 0xf1, 0x7a, 0x4f, 0xe6, 0x06,
	0x12, 0x1a, 0x0e, 0x59, 0x0a, 0xe2, 0x8a, 0x64, 0xa3, 0xb8, 0x2a, 0x48, 0x92, 0x80, 0xa2, 0x64,
	0xbc, 0xe2, 0xa2, 0x47, 0xfe, 0xa6, 0x12, 0x79, 0xaa, 0x58, 0x02, 0x2b, 0x82, 0xa7, 0xeb, 0x04,
	0x32, 0x1c, 0x42, 0x42, 0x97, 0x75, 0xfd, 0x1f, 0x06, 0x6a, 0x7a, 0x40, 0x25, 0x98, 0xef, 0xd1,
	0x6e, 0xe5, 0x26, 0xa2, 0x8a, 0x5a, 0xc6, 0x81, 0x71, 0xd8, 0x3e, 0x7a, 0x8c, 0xe7, 0x59, 0xdc,
	0x41, 0x71, 0x36, 0x8a, 0xab, 0x82, 0xc4, 0xd5, 0x34, 0x1e, 0x0f, 0xf0, 0xdb, 0xe0, 0x23, 0x84,
	0xea, 0x14, 0x14, 0x75, 0xcd, 0x49, 0xe1, 0x34, 0xca, 0xc2, 0x41, 0xf3, 0x9a, 0x7f, 0x47, 0x35,
	0x3d, 0xb4, 0x23, 0x33, 0x08, 0xad, 0x2d, 0x4d, 0x7f, 0x88, 0xd7, 0x24, 0x8d, 0xb5, 0xaf, 0xb3,
	0x0c, 0x42, 0xf7, 0xbf, 0x9a, 0xbb, 0x53, 0xbd, 0x7c, 0x4d, 0xe9, 0x7f, 0x37, 0x50, 0x4b, 0x4f,
	0x78, 0x4c, 0x2a, 0xf3, 0x62, 0xc5, 0x3d, 0xde, 0xcc, 0x7d, 0xa5, 0xd6, 0xde, 0x3b, 0xf5, 0x8e,
	0xdd, 0xdb, 0xca, 0x82, 0xf3, 0x37, 0xa8, 0xc9, 0x14, 0x24, 0xd2, 0xda, 0x3a, 0xd8, 0x3e, 0x6c,
	0x1f, 0x3d, 0xd8, 0xcc, 0xba, 0xbb, 0x57, 0x23, 0x9b, 0x27, 0x95, 0xd8, 0x9f, 0x31, 0xfa, 0xdf,
	0xb6, 0x6b, 0xe3, 0xd5, 0x31, 0xe6, 0x73, 0xb4, 0x3f, 0xe4, 0x97, 0x11, 0x88, 0x93, 0x08, 0x52,
	0xc5, 0xd4, 0x95, 0xb6, 0xdf, 0x72, 0xcd, 0xb2, 0x70, 0xf6, 0x8f, 0xff, 0xe8, 0xf8, 0x4b, 0x93,
	0xa6, 0x87, 0xba, 0x97, 0x15, 0xe8, 0x55, 0x2e, 0xf4, 0xfa, 0x33, 0x08, 0x79, 0x1a, 0x49, 0x1d,
	0x70, 0xd3, 0xb5, 0xca, 0xc2, 0xe9, 0x7a, 0xf7, 0xf4, 0xfd, 0x7b, 0x55, 0x66, 0x80, 0xda, 0x34,
	0xfc, 0x94, 0x33, 0x01, 0xe7, 0x2c, 0x01, 0x6b, 0x5b, 0xa7, 0x48, 0x36, 0x4b, 0xf1, 0x94, 0x85,
	0x82, 0x57, 0x32, 0xf7, 0xff, 0xb2, 0x70, 0xda, 0x2f, 0xe7, 0x1c, 0x7f, 0x11, 0x6a, 0x5e, 0xa0,
	0x96, 0x80, 0x14, 0x3e, 0xeb, 0x0d, 0x3b, 0xff, 0xb6, 0x61, 0xaf, 0x2c, 0x9c, 0x96, 0x7f, 0x4b,
	0xf1, 0xe7, 0x40, 0xf3, 0x05, 0xea, 0xe8, 0xcb, 0xce, 0x05, 0x4d, 0x25, 0xab, 0x6e, 0x93, 0x56,
	0x53, 0x67, 0xd1, 0x2d, 0x0b, 0xa7, 0xe3, 0x2d, 0xf5, 0xfc, 0x95, 0x69, 0xf7, 0xf5, 0x64, 0x6a,
	0x37, 0xae, 0xa7, 0x76, 0xe3, 0x66, 0x6a, 0x37, 0xbe, 0x94, 0xb6, 0x31, 0x29, 0x6d, 0xe3, 0xba,
	0xb4, 0x8d, 0x9b, 0xd2, 0x36, 0x7e, 0x96, 0xb6, 0xf1, 0xf5, 0x97, 0xdd, 0x78, 0xe7, 0xac, 0xf9,
	0xa9, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x19, 0x0e, 0xd7, 0x8f, 0x04, 0x00, 0x00,
}

func (m *Lease) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Lease) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Lease) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *LeaseList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ListMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *LeaseSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LeaseTransitions != nil {
		i = encodeVarintGenerated(dAtA, i, uint64(*m.LeaseTransitions))
		i--
		dAtA[i] = 0x28
	}
	if m.RenewTime != nil {
		{
			size, err := m.RenewTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.AcquireTime != nil {
		{
			size, err := m.AcquireTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.LeaseDurationSeconds != nil {
		i = encodeVarintGenerated(dAtA, i, uint64(*m.LeaseDurationSeconds))
		i--
		dAtA[i] = 0x10
	}
	if m.HolderIdentity != nil {
		i -= len(*m.HolderIdentity)
		copy(dAtA[i:], *m.HolderIdentity)
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.HolderIdentity)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Lease) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *LeaseList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *LeaseSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HolderIdentity != nil {
		l = len(*m.HolderIdentity)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.LeaseDurationSeconds != nil {
		n += 1 + sovGenerated(uint64(*m.LeaseDurationSeconds))
	}
	if m.AcquireTime != nil {
		l = m.AcquireTime.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.RenewTime != nil {
		l = m.RenewTime.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.LeaseTransitions != nil {
		n += 1 + sovGenerated(uint64(*m.LeaseTransitions))
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Lease) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Lease{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "LeaseSpec", "LeaseSpec", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *LeaseList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]Lease{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(strings.Replace(f.String(), "Lease", "Lease", 1), `&`, ``, 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&LeaseList{`,
		`ListMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ListMeta), "ListMeta", "v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + repeatedStringForItems + `,`,
		`}`,
	}, "")
	return s
}
func (this *LeaseSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LeaseSpec{`,
		`HolderIdentity:` + valueToStringGenerated(this.HolderIdentity) + `,`,
		`LeaseDurationSeconds:` + valueToStringGenerated(this.LeaseDurationSeconds) + `,`,
		`AcquireTime:` + strings.Replace(fmt.Sprintf("%v", this.AcquireTime), "MicroTime", "v1.MicroTime", 1) + `,`,
		`RenewTime:` + strings.Replace(fmt.Sprintf("%v", this.RenewTime), "MicroTime", "v1.MicroTime", 1) + `,`,
		`LeaseTransitions:` + valueToStringGenerated(this.LeaseTransitions) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Lease) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: Lease: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lease: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *LeaseList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: LeaseList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, Lease{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *LeaseSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: LeaseSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HolderIdentity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.HolderIdentity = &s
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseDurationSeconds", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.LeaseDurationSeconds = &v
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AcquireTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AcquireTime == nil {
				m.AcquireTime = &v1.MicroTime{}
			}
			if err := m.AcquireTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RenewTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RenewTime == nil {
				m.RenewTime = &v1.MicroTime{}
			}
			if err := m.RenewTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseTransitions", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.LeaseTransitions = &v
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)
