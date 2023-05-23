// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/containerd/containerd/runtime/v2/runc/options/oci.proto

package options

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
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

type Options struct {
	// disable pivot root when creating a container
	NoPivotRoot bool `protobuf:"varint,1,opt,name=no_pivot_root,json=noPivotRoot,proto3" json:"no_pivot_root,omitempty"`
	// create a new keyring for the container
	NoNewKeyring bool `protobuf:"varint,2,opt,name=no_new_keyring,json=noNewKeyring,proto3" json:"no_new_keyring,omitempty"`
	// place the shim in a cgroup
	ShimCgroup string `protobuf:"bytes,3,opt,name=shim_cgroup,json=shimCgroup,proto3" json:"shim_cgroup,omitempty"`
	// set the I/O's pipes uid
	IoUid uint32 `protobuf:"varint,4,opt,name=io_uid,json=ioUid,proto3" json:"io_uid,omitempty"`
	// set the I/O's pipes gid
	IoGid uint32 `protobuf:"varint,5,opt,name=io_gid,json=ioGid,proto3" json:"io_gid,omitempty"`
	// binary name of the runc binary
	BinaryName string `protobuf:"bytes,6,opt,name=binary_name,json=binaryName,proto3" json:"binary_name,omitempty"`
	// runc root directory
	Root string `protobuf:"bytes,7,opt,name=root,proto3" json:"root,omitempty"`
	// criu binary path
	CriuPath string `protobuf:"bytes,8,opt,name=criu_path,json=criuPath,proto3" json:"criu_path,omitempty"`
	// enable systemd cgroups
	SystemdCgroup bool `protobuf:"varint,9,opt,name=systemd_cgroup,json=systemdCgroup,proto3" json:"systemd_cgroup,omitempty"`
	// criu image path
	CriuImagePath string `protobuf:"bytes,10,opt,name=criu_image_path,json=criuImagePath,proto3" json:"criu_image_path,omitempty"`
	// criu work path
	CriuWorkPath         string   `protobuf:"bytes,11,opt,name=criu_work_path,json=criuWorkPath,proto3" json:"criu_work_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Options) Reset()      { *m = Options{} }
func (*Options) ProtoMessage() {}
func (*Options) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e5440d739e9a863, []int{0}
}
func (m *Options) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Options) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Options.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Options) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Options.Merge(m, src)
}
func (m *Options) XXX_Size() int {
	return m.Size()
}
func (m *Options) XXX_DiscardUnknown() {
	xxx_messageInfo_Options.DiscardUnknown(m)
}

var xxx_messageInfo_Options proto.InternalMessageInfo

type CheckpointOptions struct {
	// exit the container after a checkpoint
	Exit bool `protobuf:"varint,1,opt,name=exit,proto3" json:"exit,omitempty"`
	// checkpoint open tcp connections
	OpenTcp bool `protobuf:"varint,2,opt,name=open_tcp,json=openTcp,proto3" json:"open_tcp,omitempty"`
	// checkpoint external unix sockets
	ExternalUnixSockets bool `protobuf:"varint,3,opt,name=external_unix_sockets,json=externalUnixSockets,proto3" json:"external_unix_sockets,omitempty"`
	// checkpoint terminals (ptys)
	Terminal bool `protobuf:"varint,4,opt,name=terminal,proto3" json:"terminal,omitempty"`
	// allow checkpointing of file locks
	FileLocks bool `protobuf:"varint,5,opt,name=file_locks,json=fileLocks,proto3" json:"file_locks,omitempty"`
	// restore provided namespaces as empty namespaces
	EmptyNamespaces []string `protobuf:"bytes,6,rep,name=empty_namespaces,json=emptyNamespaces,proto3" json:"empty_namespaces,omitempty"`
	// set the cgroups mode, soft, full, strict
	CgroupsMode string `protobuf:"bytes,7,opt,name=cgroups_mode,json=cgroupsMode,proto3" json:"cgroups_mode,omitempty"`
	// checkpoint image path
	ImagePath string `protobuf:"bytes,8,opt,name=image_path,json=imagePath,proto3" json:"image_path,omitempty"`
	// checkpoint work path
	WorkPath             string   `protobuf:"bytes,9,opt,name=work_path,json=workPath,proto3" json:"work_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckpointOptions) Reset()      { *m = CheckpointOptions{} }
func (*CheckpointOptions) ProtoMessage() {}
func (*CheckpointOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e5440d739e9a863, []int{1}
}
func (m *CheckpointOptions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CheckpointOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CheckpointOptions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CheckpointOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckpointOptions.Merge(m, src)
}
func (m *CheckpointOptions) XXX_Size() int {
	return m.Size()
}
func (m *CheckpointOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckpointOptions.DiscardUnknown(m)
}

var xxx_messageInfo_CheckpointOptions proto.InternalMessageInfo

type ProcessDetails struct {
	// exec process id if the process is managed by a shim
	ExecID               string   `protobuf:"bytes,1,opt,name=exec_id,json=execId,proto3" json:"exec_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessDetails) Reset()      { *m = ProcessDetails{} }
func (*ProcessDetails) ProtoMessage() {}
func (*ProcessDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e5440d739e9a863, []int{2}
}
func (m *ProcessDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProcessDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProcessDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProcessDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessDetails.Merge(m, src)
}
func (m *ProcessDetails) XXX_Size() int {
	return m.Size()
}
func (m *ProcessDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessDetails proto.InternalMessageInfo

func init() {
	defer func(start time.Time) { bug.PrintCustomStack(start) }(time.Now());
	proto.RegisterType((*Options)(nil), "containerd.runc.v1.Options")
	proto.RegisterType((*CheckpointOptions)(nil), "containerd.runc.v1.CheckpointOptions")
	proto.RegisterType((*ProcessDetails)(nil), "containerd.runc.v1.ProcessDetails")
}

func init() {
	defer func(start time.Time) { bug.PrintCustomStack(start) }(time.Now());
	proto.RegisterFile("github.com/containerd/containerd/runtime/v2/runc/options/oci.proto", fileDescriptor_4e5440d739e9a863)
}

var fileDescriptor_4e5440d739e9a863 = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0x87, 0xeb, 0xfe, 0x49, 0xec, 0x4d, 0x93, 0xc2, 0x42, 0x25, 0xd3, 0x8a, 0x34, 0x94, 0x82,
	0xc2, 0x25, 0x11, 0x45, 0x9c, 0xb8, 0xa0, 0xb6, 0x08, 0x55, 0x40, 0xa9, 0x0c, 0x15, 0xa8, 0x97,
	0x95, 0xbb, 0x1e, 0x9c, 0x51, 0xe2, 0x1d, 0xcb, 0xbb, 0x69, 0xd2, 0x1b, 0xef, 0xc5, 0x0b, 0xf4,
	0xc8, 0x91, 0x13, 0xa2, 0xb9, 0xf1, 0x16, 0x68, 0xd7, 0x4e, 0xdb, 0x33, 0x27, 0xcf, 0x7e, 0xf3,
	0xf3, 0x78, 0xfd, 0xad, 0x96, 0xed, 0xa5, 0x68, 0x06, 0xe3, 0xb3, 0x9e, 0xa4, 0xac, 0x2f, 0x49,
	0x99, 0x18, 0x15, 0x14, 0xc9, 0xed, 0xb2, 0x18, 0x2b, 0x83, 0x19, 0xf4, 0xcf, 0x77, 0x6d, 0x29,
	0xfb, 0x94, 0x1b, 0x24, 0xa5, 0xfb, 0x24, 0xb1, 0x97, 0x17, 0x64, 0x88, 0xf3, 0x9b, 0x74, 0xcf,
	0x46, 0x7a, 0xe7, 0xcf, 0x37, 0xee, 0xa7, 0x94, 0x92, 0x6b, 0xf7, 0x6d, 0x55, 0x26, 0xb7, 0xff,
	0x2e, 0xb2, 0xfa, 0xc7, 0xf2, 0x7d, 0xbe, 0xcd, 0x9a, 0x8a, 0x44, 0x8e, 0xe7, 0x64, 0x44, 0x41,
	0x64, 0x42, 0xaf, 0xe3, 0x75, 0xfd, 0xa8, 0xa1, 0xe8, 0xd8, 0xb2, 0x88, 0xc8, 0xf0, 0x1d, 0xd6,
	0x52, 0x24, 0x14, 0x4c, 0xc4, 0x10, 0x2e, 0x0a, 0x54, 0x69, 0xb8, 0xe8, 0x42, 0xab, 0x8a, 0x8e,
	0x60, 0xf2, 0xae, 0x64, 0x7c, 0x8b, 0x35, 0xf4, 0x00, 0x33, 0x21, 0xd3, 0x82, 0xc6, 0x79, 0xb8,
	0xd4, 0xf1, 0xba, 0x41, 0xc4, 0x2c, 0xda, 0x77, 0x84, 0xaf, 0xb3, 0x1a, 0x92, 0x18, 0x63, 0x12,
	0x2e, 0x77, 0xbc, 0x6e, 0x33, 0x5a, 0x41, 0x3a, 0xc1, 0xa4, 0xc2, 0x29, 0x26, 0xe1, 0xca, 0x1c,
	0xbf, 0xc5, 0xc4, 0x8e, 0x3b, 0x43, 0x15, 0x17, 0x17, 0x42, 0xc5, 0x19, 0x84, 0xb5, 0x72, 0x5c,
	0x89, 0x8e, 0xe2, 0x0c, 0x38, 0x67, 0xcb, 0x6e, 0xc3, 0x75, 0xd7, 0x71, 0x35, 0xdf, 0x64, 0x81,
	0x2c, 0x70, 0x2c, 0xf2, 0xd8, 0x0c, 0x42, 0xdf, 0x35, 0x7c, 0x0b, 0x8e, 0x63, 0x33, 0xe0, 0x4f,
	0x58, 0x4b, 0x5f, 0x68, 0x03, 0x59, 0x32, 0xdf, 0x63, 0xe0, 0x7e, 0xa3, 0x59, 0xd1, 0x6a, 0x9b,
	0x4f, 0xd9, 0x9a, 0x9b, 0x81, 0x59, 0x9c, 0x42, 0x39, 0x89, 0xb9, 0x49, 0x4d, 0x8b, 0x0f, 0x2d,
	0x75, 0xe3, 0x76, 0x58, 0xcb, 0xe5, 0x26, 0x54, 0x0c, 0xcb, 0x58, 0xc3, 0xc5, 0x56, 0x2d, 0xfd,
	0x42, 0xc5, 0xd0, 0xa6, 0xb6, 0x7f, 0x2c, 0xb2, 0xbb, 0xfb, 0x03, 0x90, 0xc3, 0x9c, 0x50, 0x99,
	0xb9, 0x75, 0xce, 0x96, 0x61, 0x8a, 0x73, 0xd9, 0xae, 0xe6, 0x0f, 0x98, 0x4f, 0x39, 0x28, 0x61,
	0x64, 0x5e, 0xf9, 0xad, 0xdb, 0xf5, 0x67, 0x99, 0xf3, 0x5d, 0xb6, 0x0e, 0x53, 0x03, 0x85, 0x8a,
	0x47, 0x62, 0xac, 0x70, 0x2a, 0x34, 0xc9, 0x21, 0x18, 0xed, 0x24, 0xfb, 0xd1, 0xbd, 0x79, 0xf3,
	0x44, 0xe1, 0xf4, 0x53, 0xd9, 0xe2, 0x1b, 0xcc, 0x37, 0x50, 0x64, 0xa8, 0xe2, 0x91, 0xf3, 0xed,
	0x47, 0xd7, 0x6b, 0xfe, 0x90, 0xb1, 0x6f, 0x38, 0x02, 0x31, 0x22, 0x39, 0xd4, 0x4e, 0xbb, 0x1f,
	0x05, 0x96, 0xbc, 0xb7, 0x80, 0x3f, 0x63, 0x77, 0x20, 0xcb, 0x4d, 0x69, 0x5e, 0xe7, 0xb1, 0x04,
	0x1d, 0xd6, 0x3a, 0x4b, 0xdd, 0x20, 0x5a, 0x73, 0xfc, 0xe8, 0x1a, 0xf3, 0x47, 0x6c, 0xb5, 0x74,
	0xa9, 0x45, 0x46, 0x09, 0x54, 0x87, 0xd1, 0xa8, 0xd8, 0x07, 0x4a, 0xc0, 0x7e, 0xec, 0x96, 0xca,
	0xf2, 0x50, 0x02, 0xbc, 0xd6, 0xb8, 0xc9, 0x82, 0x1b, 0x83, 0x41, 0x79, 0x64, 0x93, 0xb9, 0xbd,
	0x97, 0xac, 0x75, 0x5c, 0x90, 0x04, 0xad, 0x0f, 0xc0, 0xc4, 0x38, 0xd2, 0xfc, 0x31, 0xab, 0xc3,
	0x14, 0xa4, 0xc0, 0xc4, 0xc9, 0x0b, 0xf6, 0xd8, 0xec, 0xf7, 0x56, 0xed, 0xcd, 0x14, 0xe4, 0xe1,
	0x41, 0x54, 0xb3, 0xad, 0xc3, 0x64, 0xef, 0xf4, 0xf2, 0xaa, 0xbd, 0xf0, 0xeb, 0xaa, 0xbd, 0xf0,
	0x7d, 0xd6, 0xf6, 0x2e, 0x67, 0x6d, 0xef, 0xe7, 0xac, 0xed, 0xfd, 0x99, 0xb5, 0xbd, 0xd3, 0xd7,
	0xff, 0x7b, 0xd1, 0x5e, 0x55, 0xcf, 0xaf, 0x0b, 0x67, 0x35, 0x77, 0x8b, 0x5e, 0xfc, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0x90, 0x50, 0x79, 0xf2, 0xb5, 0x03, 0x00, 0x00,
}

func (m *Options) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Options) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Options) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.CriuWorkPath) > 0 {
		i -= len(m.CriuWorkPath)
		copy(dAtA[i:], m.CriuWorkPath)
		i = encodeVarintOci(dAtA, i, uint64(len(m.CriuWorkPath)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.CriuImagePath) > 0 {
		i -= len(m.CriuImagePath)
		copy(dAtA[i:], m.CriuImagePath)
		i = encodeVarintOci(dAtA, i, uint64(len(m.CriuImagePath)))
		i--
		dAtA[i] = 0x52
	}
	if m.SystemdCgroup {
		i--
		if m.SystemdCgroup {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if len(m.CriuPath) > 0 {
		i -= len(m.CriuPath)
		copy(dAtA[i:], m.CriuPath)
		i = encodeVarintOci(dAtA, i, uint64(len(m.CriuPath)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Root) > 0 {
		i -= len(m.Root)
		copy(dAtA[i:], m.Root)
		i = encodeVarintOci(dAtA, i, uint64(len(m.Root)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.BinaryName) > 0 {
		i -= len(m.BinaryName)
		copy(dAtA[i:], m.BinaryName)
		i = encodeVarintOci(dAtA, i, uint64(len(m.BinaryName)))
		i--
		dAtA[i] = 0x32
	}
	if m.IoGid != 0 {
		i = encodeVarintOci(dAtA, i, uint64(m.IoGid))
		i--
		dAtA[i] = 0x28
	}
	if m.IoUid != 0 {
		i = encodeVarintOci(dAtA, i, uint64(m.IoUid))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ShimCgroup) > 0 {
		i -= len(m.ShimCgroup)
		copy(dAtA[i:], m.ShimCgroup)
		i = encodeVarintOci(dAtA, i, uint64(len(m.ShimCgroup)))
		i--
		dAtA[i] = 0x1a
	}
	if m.NoNewKeyring {
		i--
		if m.NoNewKeyring {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.NoPivotRoot {
		i--
		if m.NoPivotRoot {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CheckpointOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckpointOptions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CheckpointOptions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.WorkPath) > 0 {
		i -= len(m.WorkPath)
		copy(dAtA[i:], m.WorkPath)
		i = encodeVarintOci(dAtA, i, uint64(len(m.WorkPath)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.ImagePath) > 0 {
		i -= len(m.ImagePath)
		copy(dAtA[i:], m.ImagePath)
		i = encodeVarintOci(dAtA, i, uint64(len(m.ImagePath)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.CgroupsMode) > 0 {
		i -= len(m.CgroupsMode)
		copy(dAtA[i:], m.CgroupsMode)
		i = encodeVarintOci(dAtA, i, uint64(len(m.CgroupsMode)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.EmptyNamespaces) > 0 {
		for iNdEx := len(m.EmptyNamespaces) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.EmptyNamespaces[iNdEx])
			copy(dAtA[i:], m.EmptyNamespaces[iNdEx])
			i = encodeVarintOci(dAtA, i, uint64(len(m.EmptyNamespaces[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if m.FileLocks {
		i--
		if m.FileLocks {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.Terminal {
		i--
		if m.Terminal {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.ExternalUnixSockets {
		i--
		if m.ExternalUnixSockets {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.OpenTcp {
		i--
		if m.OpenTcp {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Exit {
		i--
		if m.Exit {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ProcessDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProcessDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProcessDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ExecID) > 0 {
		i -= len(m.ExecID)
		copy(dAtA[i:], m.ExecID)
		i = encodeVarintOci(dAtA, i, uint64(len(m.ExecID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOci(dAtA []byte, offset int, v uint64) int {
	offset -= sovOci(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Options) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NoPivotRoot {
		n += 2
	}
	if m.NoNewKeyring {
		n += 2
	}
	l = len(m.ShimCgroup)
	if l > 0 {
		n += 1 + l + sovOci(uint64(l))
	}
	if m.IoUid != 0 {
		n += 1 + sovOci(uint64(m.IoUid))
	}
	if m.IoGid != 0 {
		n += 1 + sovOci(uint64(m.IoGid))
	}
	l = len(m.BinaryName)
	if l > 0 {
		n += 1 + l + sovOci(uint64(l))
	}
	l = len(m.Root)
	if l > 0 {
		n += 1 + l + sovOci(uint64(l))
	}
	l = len(m.CriuPath)
	if l > 0 {
		n += 1 + l + sovOci(uint64(l))
	}
	if m.SystemdCgroup {
		n += 2
	}
	l = len(m.CriuImagePath)
	if l > 0 {
		n += 1 + l + sovOci(uint64(l))
	}
	l = len(m.CriuWorkPath)
	if l > 0 {
		n += 1 + l + sovOci(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CheckpointOptions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Exit {
		n += 2
	}
	if m.OpenTcp {
		n += 2
	}
	if m.ExternalUnixSockets {
		n += 2
	}
	if m.Terminal {
		n += 2
	}
	if m.FileLocks {
		n += 2
	}
	if len(m.EmptyNamespaces) > 0 {
		for _, s := range m.EmptyNamespaces {
			l = len(s)
			n += 1 + l + sovOci(uint64(l))
		}
	}
	l = len(m.CgroupsMode)
	if l > 0 {
		n += 1 + l + sovOci(uint64(l))
	}
	l = len(m.ImagePath)
	if l > 0 {
		n += 1 + l + sovOci(uint64(l))
	}
	l = len(m.WorkPath)
	if l > 0 {
		n += 1 + l + sovOci(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ProcessDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ExecID)
	if l > 0 {
		n += 1 + l + sovOci(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovOci(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOci(x uint64) (n int) {
	return sovOci(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Options) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Options{`,
		`NoPivotRoot:` + fmt.Sprintf("%v", this.NoPivotRoot) + `,`,
		`NoNewKeyring:` + fmt.Sprintf("%v", this.NoNewKeyring) + `,`,
		`ShimCgroup:` + fmt.Sprintf("%v", this.ShimCgroup) + `,`,
		`IoUid:` + fmt.Sprintf("%v", this.IoUid) + `,`,
		`IoGid:` + fmt.Sprintf("%v", this.IoGid) + `,`,
		`BinaryName:` + fmt.Sprintf("%v", this.BinaryName) + `,`,
		`Root:` + fmt.Sprintf("%v", this.Root) + `,`,
		`CriuPath:` + fmt.Sprintf("%v", this.CriuPath) + `,`,
		`SystemdCgroup:` + fmt.Sprintf("%v", this.SystemdCgroup) + `,`,
		`CriuImagePath:` + fmt.Sprintf("%v", this.CriuImagePath) + `,`,
		`CriuWorkPath:` + fmt.Sprintf("%v", this.CriuWorkPath) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CheckpointOptions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CheckpointOptions{`,
		`Exit:` + fmt.Sprintf("%v", this.Exit) + `,`,
		`OpenTcp:` + fmt.Sprintf("%v", this.OpenTcp) + `,`,
		`ExternalUnixSockets:` + fmt.Sprintf("%v", this.ExternalUnixSockets) + `,`,
		`Terminal:` + fmt.Sprintf("%v", this.Terminal) + `,`,
		`FileLocks:` + fmt.Sprintf("%v", this.FileLocks) + `,`,
		`EmptyNamespaces:` + fmt.Sprintf("%v", this.EmptyNamespaces) + `,`,
		`CgroupsMode:` + fmt.Sprintf("%v", this.CgroupsMode) + `,`,
		`ImagePath:` + fmt.Sprintf("%v", this.ImagePath) + `,`,
		`WorkPath:` + fmt.Sprintf("%v", this.WorkPath) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProcessDetails) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProcessDetails{`,
		`ExecID:` + fmt.Sprintf("%v", this.ExecID) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringOci(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Options) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOci
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
			return fmt.Errorf("proto: Options: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Options: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoPivotRoot", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOci
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
			m.NoPivotRoot = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoNewKeyring", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOci
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
			m.NoNewKeyring = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShimCgroup", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOci
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
				return ErrInvalidLengthOci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShimCgroup = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IoUid", wireType)
			}
			m.IoUid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOci
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IoUid |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IoGid", wireType)
			}
			m.IoGid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOci
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IoGid |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BinaryName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOci
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
				return ErrInvalidLengthOci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BinaryName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOci
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
				return ErrInvalidLengthOci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Root = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CriuPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOci
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
				return ErrInvalidLengthOci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CriuPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SystemdCgroup", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOci
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
			m.SystemdCgroup = bool(v != 0)
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CriuImagePath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOci
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
				return ErrInvalidLengthOci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CriuImagePath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CriuWorkPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOci
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
				return ErrInvalidLengthOci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CriuWorkPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOci(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOci
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CheckpointOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOci
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
			return fmt.Errorf("proto: CheckpointOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckpointOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exit", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOci
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
			m.Exit = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpenTcp", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOci
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
			m.OpenTcp = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalUnixSockets", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOci
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
			m.ExternalUnixSockets = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Terminal", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOci
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
			m.Terminal = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileLocks", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOci
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
			m.FileLocks = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmptyNamespaces", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOci
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
				return ErrInvalidLengthOci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EmptyNamespaces = append(m.EmptyNamespaces, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CgroupsMode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOci
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
				return ErrInvalidLengthOci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CgroupsMode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImagePath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOci
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
				return ErrInvalidLengthOci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImagePath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorkPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOci
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
				return ErrInvalidLengthOci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WorkPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOci(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOci
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProcessDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOci
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
			return fmt.Errorf("proto: ProcessDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOci
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
				return ErrInvalidLengthOci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOci(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOci
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipOci(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOci
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
					return 0, ErrIntOverflowOci
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
					return 0, ErrIntOverflowOci
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
				return 0, ErrInvalidLengthOci
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOci
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOci
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOci        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOci          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOci = fmt.Errorf("proto: unexpected end of group")
)
