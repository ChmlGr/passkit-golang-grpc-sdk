// Code generated by protoc-gen-go. DO NOT EDIT.
// source: io/common/pass.proto

package io

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// PassObject contains pass-object level details that are common across protocols.
type PassOverrides struct {
	// Image assets hold image data for this pass.
	//
	// Types that are valid to be assigned to ImageAssets:
	//	*PassOverrides_ImageIds
	//	*PassOverrides_Images
	ImageAssets isPassOverrides_ImageAssets `protobuf_oneof:"ImageAssets"`
	// A list of up to 10 GPS locations where this pass should be presented on the lock-screen. Overrides any template level values.
	Locations []*GPSLocation `protobuf:"bytes,3,rep,name=locations,proto3" json:"locations,omitempty"`
	// A list of up to 10 Beacon UUIDs that should trigger the pass to be presented on the lock-screen. Overrides any template level values.
	Beacons []*Beacon `protobuf:"bytes,4,rep,name=beacons,proto3" json:"beacons,omitempty"`
	// A list of (personalized) links to be shown on the pass. Overrides any template level values.
	Links                []*Link  `protobuf:"bytes,5,rep,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PassOverrides) Reset()         { *m = PassOverrides{} }
func (m *PassOverrides) String() string { return proto.CompactTextString(m) }
func (*PassOverrides) ProtoMessage()    {}
func (*PassOverrides) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7e1b4d2f709df4, []int{0}
}

func (m *PassOverrides) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PassOverrides.Unmarshal(m, b)
}
func (m *PassOverrides) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PassOverrides.Marshal(b, m, deterministic)
}
func (m *PassOverrides) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PassOverrides.Merge(m, src)
}
func (m *PassOverrides) XXX_Size() int {
	return xxx_messageInfo_PassOverrides.Size(m)
}
func (m *PassOverrides) XXX_DiscardUnknown() {
	xxx_messageInfo_PassOverrides.DiscardUnknown(m)
}

var xxx_messageInfo_PassOverrides proto.InternalMessageInfo

type isPassOverrides_ImageAssets interface {
	isPassOverrides_ImageAssets()
}

type PassOverrides_ImageIds struct {
	ImageIds *ImageIds `protobuf:"bytes,1,opt,name=imageIds,proto3,oneof"`
}

type PassOverrides_Images struct {
	Images *ImageData `protobuf:"bytes,2,opt,name=images,proto3,oneof"`
}

func (*PassOverrides_ImageIds) isPassOverrides_ImageAssets() {}

func (*PassOverrides_Images) isPassOverrides_ImageAssets() {}

func (m *PassOverrides) GetImageAssets() isPassOverrides_ImageAssets {
	if m != nil {
		return m.ImageAssets
	}
	return nil
}

func (m *PassOverrides) GetImageIds() *ImageIds {
	if x, ok := m.GetImageAssets().(*PassOverrides_ImageIds); ok {
		return x.ImageIds
	}
	return nil
}

func (m *PassOverrides) GetImages() *ImageData {
	if x, ok := m.GetImageAssets().(*PassOverrides_Images); ok {
		return x.Images
	}
	return nil
}

func (m *PassOverrides) GetLocations() []*GPSLocation {
	if m != nil {
		return m.Locations
	}
	return nil
}

func (m *PassOverrides) GetBeacons() []*Beacon {
	if m != nil {
		return m.Beacons
	}
	return nil
}

func (m *PassOverrides) GetLinks() []*Link {
	if m != nil {
		return m.Links
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PassOverrides) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PassOverrides_ImageIds)(nil),
		(*PassOverrides_Images)(nil),
	}
}

func init() {
	proto.RegisterType((*PassOverrides)(nil), "io.PassOverrides")
}

func init() { proto.RegisterFile("io/common/pass.proto", fileDescriptor_6c7e1b4d2f709df4) }

var fileDescriptor_6c7e1b4d2f709df4 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x4d, 0x4b, 0xfb, 0x40,
	0x10, 0xc6, 0xdb, 0xf4, 0xe5, 0xdf, 0xff, 0xd6, 0x52, 0x08, 0x0a, 0xb1, 0xa7, 0x22, 0xa2, 0x45,
	0xda, 0x2c, 0x54, 0xbc, 0x0a, 0x0d, 0x82, 0x2d, 0x2d, 0x18, 0xea, 0xcd, 0xdb, 0x36, 0x59, 0xe2,
	0x90, 0x66, 0x27, 0x64, 0x16, 0x5f, 0x6e, 0x7e, 0x06, 0x3f, 0x81, 0x67, 0x3f, 0xa5, 0xec, 0x26,
	0xa6, 0x5e, 0x76, 0x87, 0xe7, 0xf9, 0x3d, 0xbb, 0xc3, 0x0c, 0x3b, 0x06, 0xe4, 0x11, 0x66, 0x19,
	0x2a, 0x9e, 0x0b, 0x22, 0x3f, 0x2f, 0x50, 0xa3, 0xeb, 0x00, 0x8e, 0x8c, 0x03, 0x99, 0x48, 0x64,
	0x79, 0x96, 0xce, 0xe8, 0xe4, 0xc0, 0xef, 0x41, 0xa5, 0x55, 0x60, 0x74, 0xfa, 0xe7, 0x99, 0x02,
	0xdf, 0x20, 0x03, 0xfd, 0x5e, 0x59, 0x53, 0x7b, 0x45, 0xb3, 0x44, 0xaa, 0x19, 0xbd, 0x8a, 0x24,
	0x91, 0x05, 0xc7, 0x5c, 0x03, 0x2a, 0xe2, 0x42, 0x29, 0xd4, 0xc2, 0xd6, 0x25, 0x7d, 0xf6, 0xe1,
	0xb0, 0x41, 0x28, 0x88, 0x1e, 0x5e, 0x64, 0x51, 0x40, 0x2c, 0xc9, 0xbd, 0x62, 0x3d, 0xdb, 0xc0,
	0x2a, 0x26, 0xaf, 0x39, 0x6e, 0x4e, 0xfa, 0xf3, 0x23, 0x1f, 0xd0, 0x5f, 0x55, 0xda, 0xb2, 0xb1,
	0xad, 0x7d, 0xf7, 0x92, 0x75, 0x6d, 0x4d, 0x9e, 0x63, 0xc9, 0x41, 0x4d, 0xde, 0x09, 0x2d, 0x96,
	0x8d, 0x6d, 0x65, 0xbb, 0x37, 0xec, 0xff, 0x1e, 0xa3, 0xf2, 0x67, 0xaf, 0x35, 0x6e, 0x4d, 0xfa,
	0xf3, 0xa1, 0x61, 0xef, 0xc3, 0xc7, 0x4d, 0xa5, 0x07, 0xdd, 0xcf, 0x45, 0xeb, 0xab, 0xc9, 0xb6,
	0x07, 0xd2, 0x9d, 0xb2, 0x7f, 0x3b, 0x29, 0x22, 0x13, 0x6a, 0xdb, 0x10, 0x33, 0xa1, 0xc0, 0x4a,
	0x35, 0xff, 0x8b, 0xb8, 0x17, 0xac, 0x63, 0x67, 0xe4, 0x75, 0x2c, 0xdb, 0x33, 0xec, 0x06, 0x54,
	0x5a, 0x93, 0xa5, 0x1d, 0x0c, 0x58, 0xdf, 0xf6, 0xb8, 0x20, 0x92, 0x9a, 0x82, 0x5b, 0x36, 0x04,
	0xf4, 0xcd, 0x36, 0x52, 0xd0, 0x7e, 0xb8, 0x06, 0x7c, 0x3a, 0x27, 0x2d, 0xe8, 0xb9, 0xd6, 0x22,
	0xcc, 0x38, 0x20, 0xcf, 0x30, 0x96, 0x7b, 0x4e, 0x71, 0xca, 0x13, 0xe4, 0x80, 0xdf, 0x4e, 0x3b,
	0x5c, 0xaf, 0x70, 0xd7, 0xb5, 0x93, 0xbc, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x8b, 0x64,
	0xfa, 0xdb, 0x01, 0x00, 0x00,
}