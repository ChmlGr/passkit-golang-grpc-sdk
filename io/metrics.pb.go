// Code generated by protoc-gen-go. DO NOT EDIT.
// source: io/common/metrics.proto

package io

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Statuses of the pass. A pass will collect multiple statuses as it passes through its lifecycle.
type LifecycleEvents int32

const (
	LifecycleEvents_NO_STATUS LifecycleEvents = 0
	// Pass has been created and presented to the user to install.
	LifecycleEvents_RENDERED LifecycleEvents = 1
	// Pass has been installed in at least one Apple device.
	LifecycleEvents_INSTALLED_APPLE LifecycleEvents = 2
	// Pass has been installed in at least one Google Pay account.
	LifecycleEvents_INSTALLED_GOOGLE LifecycleEvents = 4
	// Pass has been installed in at least one Android wallet app.
	LifecycleEvents_INSTALLED_ANDROID_OTHER LifecycleEvents = 8
	// Pass has been uninstalled from at least one Apple device.
	LifecycleEvents_UNINSTALLED_APPLE LifecycleEvents = 1024
	// Pass has been uninstalled from at least one Google Pay account.
	LifecycleEvents_UNINSTALLED_GOOGLE LifecycleEvents = 2048
	// Pass has been uninstalled from at least one Android wallet app.
	LifecycleEvents_UNINSTALLED_ANDROID_OTHER LifecycleEvents = 4096
	// An Apple pass bundle has been created.
	LifecycleEvents_APPLE_PASS_CREATED LifecycleEvents = 8192
	// A Google Pay record has been created.
	LifecycleEvents_GOOGLE_PAY_RECORD_CREATED LifecycleEvents = 16384
	// Invalidation request sent to pass.
	LifecycleEvents_INVALIDATE_REQUESTED LifecycleEvents = 262144
	// Pass has been invalidated.
	LifecycleEvents_INVALIDATE_CONFIRMED LifecycleEvents = 524288
)

var LifecycleEvents_name = map[int32]string{
	0:      "NO_STATUS",
	1:      "RENDERED",
	2:      "INSTALLED_APPLE",
	4:      "INSTALLED_GOOGLE",
	8:      "INSTALLED_ANDROID_OTHER",
	1024:   "UNINSTALLED_APPLE",
	2048:   "UNINSTALLED_GOOGLE",
	4096:   "UNINSTALLED_ANDROID_OTHER",
	8192:   "APPLE_PASS_CREATED",
	16384:  "GOOGLE_PAY_RECORD_CREATED",
	262144: "INVALIDATE_REQUESTED",
	524288: "INVALIDATE_CONFIRMED",
}

var LifecycleEvents_value = map[string]int32{
	"NO_STATUS":                 0,
	"RENDERED":                  1,
	"INSTALLED_APPLE":           2,
	"INSTALLED_GOOGLE":          4,
	"INSTALLED_ANDROID_OTHER":   8,
	"UNINSTALLED_APPLE":         1024,
	"UNINSTALLED_GOOGLE":        2048,
	"UNINSTALLED_ANDROID_OTHER": 4096,
	"APPLE_PASS_CREATED":        8192,
	"GOOGLE_PAY_RECORD_CREATED": 16384,
	"INVALIDATE_REQUESTED":      262144,
	"INVALIDATE_CONFIRMED":      524288,
}

func (x LifecycleEvents) String() string {
	return proto.EnumName(LifecycleEvents_name, int32(x))
}

func (LifecycleEvents) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e3ee370c4b50b41a, []int{0}
}

// Status is the best determined status of the pass.
type Status int32

const (
	// Pass has been issued.
	Status_PASS_ISSUED Status = 0
	// Pass has been installed in a wallet.
	Status_PASS_INSTALLED Status = 1
	// Pass has been uninstalled (possibly deleted) from a wallet.
	Status_PASS_UNINSTALLED Status = 2
	// Pass has been invalidated.
	Status_PASS_INVALIDATED Status = 3
)

var Status_name = map[int32]string{
	0: "PASS_ISSUED",
	1: "PASS_INSTALLED",
	2: "PASS_UNINSTALLED",
	3: "PASS_INVALIDATED",
}

var Status_value = map[string]int32{
	"PASS_ISSUED":      0,
	"PASS_INSTALLED":   1,
	"PASS_UNINSTALLED": 2,
	"PASS_INVALIDATED": 3,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e3ee370c4b50b41a, []int{1}
}

// Metrics list issued, installed, uninstalled and invalidated counts.
type Metrics struct {
	// Pass records created.
	Created uint64 `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	// Number of passes installed in Apple Wallet.
	AppleWalletInstalls uint64 `protobuf:"varint,2,opt,name=appleWalletInstalls,proto3" json:"appleWalletInstalls,omitempty"`
	// Number of passes installed in Google Pay.
	GooglePayInstalls uint64 `protobuf:"varint,3,opt,name=googlePayInstalls,proto3" json:"googlePayInstalls,omitempty"`
	// Number of passes installed in Android wallet apps e.g. WalletPasses.
	OtherAndroidInstalls uint64 `protobuf:"varint,4,opt,name=otherAndroidInstalls,proto3" json:"otherAndroidInstalls,omitempty"`
	// Number of passes uninstalled from Apple Wallet.
	AppleWalletUninstalls uint64 `protobuf:"varint,11,opt,name=appleWalletUninstalls,proto3" json:"appleWalletUninstalls,omitempty"`
	// Number of passes uninstalled from Google Pay.
	GooglePayUninstalls uint64 `protobuf:"varint,12,opt,name=googlePayUninstalls,proto3" json:"googlePayUninstalls,omitempty"`
	// Number of passes uninstalled from Android wallet apps e.g. WalletPasses.
	OtherAndroidUninstalls uint64 `protobuf:"varint,13,opt,name=otherAndroidUninstalls,proto3" json:"otherAndroidUninstalls,omitempty"`
	// Number of passes invalidated.
	Invalidations        uint64   `protobuf:"varint,20,opt,name=invalidations,proto3" json:"invalidations,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metrics) Reset()         { *m = Metrics{} }
func (m *Metrics) String() string { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()    {}
func (*Metrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3ee370c4b50b41a, []int{0}
}

func (m *Metrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metrics.Unmarshal(m, b)
}
func (m *Metrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metrics.Marshal(b, m, deterministic)
}
func (m *Metrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metrics.Merge(m, src)
}
func (m *Metrics) XXX_Size() int {
	return xxx_messageInfo_Metrics.Size(m)
}
func (m *Metrics) XXX_DiscardUnknown() {
	xxx_messageInfo_Metrics.DiscardUnknown(m)
}

var xxx_messageInfo_Metrics proto.InternalMessageInfo

func (m *Metrics) GetCreated() uint64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Metrics) GetAppleWalletInstalls() uint64 {
	if m != nil {
		return m.AppleWalletInstalls
	}
	return 0
}

func (m *Metrics) GetGooglePayInstalls() uint64 {
	if m != nil {
		return m.GooglePayInstalls
	}
	return 0
}

func (m *Metrics) GetOtherAndroidInstalls() uint64 {
	if m != nil {
		return m.OtherAndroidInstalls
	}
	return 0
}

func (m *Metrics) GetAppleWalletUninstalls() uint64 {
	if m != nil {
		return m.AppleWalletUninstalls
	}
	return 0
}

func (m *Metrics) GetGooglePayUninstalls() uint64 {
	if m != nil {
		return m.GooglePayUninstalls
	}
	return 0
}

func (m *Metrics) GetOtherAndroidUninstalls() uint64 {
	if m != nil {
		return m.OtherAndroidUninstalls
	}
	return 0
}

func (m *Metrics) GetInvalidations() uint64 {
	if m != nil {
		return m.Invalidations
	}
	return 0
}

type UTM struct {
	// Use source to identify where the request is coming from. Defaults to Unknown.
	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// Use medium to identify a medium such as email, app, or cost-per-click advertising.
	Medium string `protobuf:"bytes,2,opt,name=medium,proto3" json:"medium,omitempty"`
	// Used for keyword analysis. Use campaign to identify a specific product promotion or strategic campaign.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Used for paid search. Use term to note the keywords for the ad that led to the pass.
	Term string `protobuf:"bytes,4,opt,name=term,proto3" json:"term,omitempty"`
	// Used for A/B testing and content-targeted ads. Use content to differentiate ads or links that point to the same URL.
	Content              string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UTM) Reset()         { *m = UTM{} }
func (m *UTM) String() string { return proto.CompactTextString(m) }
func (*UTM) ProtoMessage()    {}
func (*UTM) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3ee370c4b50b41a, []int{1}
}

func (m *UTM) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UTM.Unmarshal(m, b)
}
func (m *UTM) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UTM.Marshal(b, m, deterministic)
}
func (m *UTM) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UTM.Merge(m, src)
}
func (m *UTM) XXX_Size() int {
	return xxx_messageInfo_UTM.Size(m)
}
func (m *UTM) XXX_DiscardUnknown() {
	xxx_messageInfo_UTM.DiscardUnknown(m)
}

var xxx_messageInfo_UTM proto.InternalMessageInfo

func (m *UTM) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *UTM) GetMedium() string {
	if m != nil {
		return m.Medium
	}
	return ""
}

func (m *UTM) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UTM) GetTerm() string {
	if m != nil {
		return m.Term
	}
	return ""
}

func (m *UTM) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type MetaLocation struct {
	// Country of issue.
	Country string `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	// City of issue.
	City string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	// State, region, province or county of issue.
	State string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	// Latitude of issue (provided in the issue call).
	Latitude float64 `protobuf:"fixed64,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// Longitude of issue (provided in the issue call).
	Longitude            float64  `protobuf:"fixed64,5,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetaLocation) Reset()         { *m = MetaLocation{} }
func (m *MetaLocation) String() string { return proto.CompactTextString(m) }
func (*MetaLocation) ProtoMessage()    {}
func (*MetaLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3ee370c4b50b41a, []int{2}
}

func (m *MetaLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaLocation.Unmarshal(m, b)
}
func (m *MetaLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaLocation.Marshal(b, m, deterministic)
}
func (m *MetaLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaLocation.Merge(m, src)
}
func (m *MetaLocation) XXX_Size() int {
	return xxx_messageInfo_MetaLocation.Size(m)
}
func (m *MetaLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaLocation.DiscardUnknown(m)
}

var xxx_messageInfo_MetaLocation proto.InternalMessageInfo

func (m *MetaLocation) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *MetaLocation) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *MetaLocation) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *MetaLocation) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *MetaLocation) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

// Additional data relating to the pass. Metadata can be provided with a create pass request. The only writeable field is UTM. All other fields are system generated.
type Metadata struct {
	// Current status of the pass. Not Writable.
	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=io.Status" json:"status,omitempty"`
	// Events that the pass has attained. Not Writable.
	LifecycleEvents []LifecycleEvents `protobuf:"varint,2,rep,packed,name=lifecycleEvents,proto3,enum=io.LifecycleEvents" json:"lifecycleEvents,omitempty"`
	// Contains UTM data if provided in the URL parameters (utm_source, utm_medium, utm_name, utm_term, utm_content). Can also be used to track channel (web / app, etc).
	Utm *UTM `protobuf:"bytes,3,opt,name=utm,proto3" json:"utm,omitempty"`
	// Id of the physical pass. Not Writeable
	AltId string `protobuf:"bytes,4,opt,name=altId,proto3" json:"altId,omitempty"`
	// IP address of the pass issue. Not Writeable
	IssueIpAddress string `protobuf:"bytes,5,opt,name=issueIpAddress,proto3" json:"issueIpAddress,omitempty"`
	// IP address of the device where the pass was installed on. Not Writeable.
	InstallIpAddress string `protobuf:"bytes,6,opt,name=installIpAddress,proto3" json:"installIpAddress,omitempty"`
	// Location information based of the where the pass landing page was rendered. Not Writeable
	RenderLocation *MetaLocation `protobuf:"bytes,7,opt,name=renderLocation,proto3" json:"renderLocation,omitempty"`
	// IP address of the device where the pass landing page was rendered on. Not Writeable.
	RenderIpAddress string `protobuf:"bytes,8,opt,name=renderIpAddress,proto3" json:"renderIpAddress,omitempty"`
	// User agent provided by the Wallet Application. Not Writable.
	InstallUserAgent string `protobuf:"bytes,9,opt,name=installUserAgent,proto3" json:"installUserAgent,omitempty"`
	// User agent where the landing page was rendered on. Not Writable.
	RenderUserAgent string `protobuf:"bytes,10,opt,name=renderUserAgent,proto3" json:"renderUserAgent,omitempty"`
	// User agent provided by the Wallet Application. Not Writable.
	InstallDeviceAttributes []DeviceAttributes `protobuf:"varint,11,rep,packed,name=installDeviceAttributes,proto3,enum=io.DeviceAttributes" json:"installDeviceAttributes,omitempty"`
	// User agent where the landing page was rendered on. Not Writable.
	RenderDeviceAttributes []DeviceAttributes `protobuf:"varint,12,rep,packed,name=renderDeviceAttributes,proto3,enum=io.DeviceAttributes" json:"renderDeviceAttributes,omitempty"`
	// Date the pass record was created. Not Writable.
	IssueAt *timestamp.Timestamp `protobuf:"bytes,13,opt,name=issueAt,proto3" json:"issueAt,omitempty"`
	// Date the pass was first displayed to the user to install. Not Writable.
	RenderedAt *timestamp.Timestamp `protobuf:"bytes,14,opt,name=renderedAt,proto3" json:"renderedAt,omitempty"`
	// Date the pass was first installed in a mobile wallet. Not Writable.
	FirstInstalledAt *timestamp.Timestamp `protobuf:"bytes,15,opt,name=firstInstalledAt,proto3" json:"firstInstalledAt,omitempty"`
	// Date the pass was last installed in a mobile wallet. For Apple devices, this will update each time the device registers a new token. Not Writable.
	LastInstalledAt *timestamp.Timestamp `protobuf:"bytes,16,opt,name=lastInstalledAt,proto3" json:"lastInstalledAt,omitempty"`
	// Date the pass was first uninstalled from a mobile wallet.  For Apple devices, this represents either the time notifications were disabled, or when the pass was removed from the wallet as it is not possible to distinguish the two events. Not Writable.
	FirstUninstalledAt *timestamp.Timestamp `protobuf:"bytes,17,opt,name=firstUninstalledAt,proto3" json:"firstUninstalledAt,omitempty"`
	// Date the pass was last uninstalled from a mobile wallet. Not Writable.
	LastUninstalledAt *timestamp.Timestamp `protobuf:"bytes,18,opt,name=lastUninstalledAt,proto3" json:"lastUninstalledAt,omitempty"`
	// Date the pass was invalidated. Not Writable.
	InvalidatedAt *timestamp.Timestamp `protobuf:"bytes,19,opt,name=invalidatedAt,proto3" json:"invalidatedAt,omitempty"`
	// Date the pass record was last updated. Can be used for segmenting passes based 'last visit', or 'last-purchase'. Not Writable.
	LastUpdatedAt        *timestamp.Timestamp `protobuf:"bytes,20,opt,name=lastUpdatedAt,proto3" json:"lastUpdatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3ee370c4b50b41a, []int{3}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_PASS_ISSUED
}

func (m *Metadata) GetLifecycleEvents() []LifecycleEvents {
	if m != nil {
		return m.LifecycleEvents
	}
	return nil
}

func (m *Metadata) GetUtm() *UTM {
	if m != nil {
		return m.Utm
	}
	return nil
}

func (m *Metadata) GetAltId() string {
	if m != nil {
		return m.AltId
	}
	return ""
}

func (m *Metadata) GetIssueIpAddress() string {
	if m != nil {
		return m.IssueIpAddress
	}
	return ""
}

func (m *Metadata) GetInstallIpAddress() string {
	if m != nil {
		return m.InstallIpAddress
	}
	return ""
}

func (m *Metadata) GetRenderLocation() *MetaLocation {
	if m != nil {
		return m.RenderLocation
	}
	return nil
}

func (m *Metadata) GetRenderIpAddress() string {
	if m != nil {
		return m.RenderIpAddress
	}
	return ""
}

func (m *Metadata) GetInstallUserAgent() string {
	if m != nil {
		return m.InstallUserAgent
	}
	return ""
}

func (m *Metadata) GetRenderUserAgent() string {
	if m != nil {
		return m.RenderUserAgent
	}
	return ""
}

func (m *Metadata) GetInstallDeviceAttributes() []DeviceAttributes {
	if m != nil {
		return m.InstallDeviceAttributes
	}
	return nil
}

func (m *Metadata) GetRenderDeviceAttributes() []DeviceAttributes {
	if m != nil {
		return m.RenderDeviceAttributes
	}
	return nil
}

func (m *Metadata) GetIssueAt() *timestamp.Timestamp {
	if m != nil {
		return m.IssueAt
	}
	return nil
}

func (m *Metadata) GetRenderedAt() *timestamp.Timestamp {
	if m != nil {
		return m.RenderedAt
	}
	return nil
}

func (m *Metadata) GetFirstInstalledAt() *timestamp.Timestamp {
	if m != nil {
		return m.FirstInstalledAt
	}
	return nil
}

func (m *Metadata) GetLastInstalledAt() *timestamp.Timestamp {
	if m != nil {
		return m.LastInstalledAt
	}
	return nil
}

func (m *Metadata) GetFirstUninstalledAt() *timestamp.Timestamp {
	if m != nil {
		return m.FirstUninstalledAt
	}
	return nil
}

func (m *Metadata) GetLastUninstalledAt() *timestamp.Timestamp {
	if m != nil {
		return m.LastUninstalledAt
	}
	return nil
}

func (m *Metadata) GetInvalidatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.InvalidatedAt
	}
	return nil
}

func (m *Metadata) GetLastUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("io.LifecycleEvents", LifecycleEvents_name, LifecycleEvents_value)
	proto.RegisterEnum("io.Status", Status_name, Status_value)
	proto.RegisterType((*Metrics)(nil), "io.Metrics")
	proto.RegisterType((*UTM)(nil), "io.UTM")
	proto.RegisterType((*MetaLocation)(nil), "io.MetaLocation")
	proto.RegisterType((*Metadata)(nil), "io.Metadata")
}

func init() {
	proto.RegisterFile("io/common/metrics.proto", fileDescriptor_e3ee370c4b50b41a)
}

var fileDescriptor_e3ee370c4b50b41a = []byte{
	// 1029 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xe1, 0x6e, 0xe2, 0x46,
	0x10, 0x3e, 0x82, 0x03, 0x66, 0xc8, 0xc1, 0x32, 0xd0, 0x84, 0xd0, 0xaa, 0x17, 0x45, 0xa7, 0x2a,
	0x8a, 0x2a, 0xa8, 0xd2, 0x53, 0x55, 0x55, 0x6a, 0x75, 0x6e, 0xec, 0xbb, 0xc3, 0x47, 0x80, 0x1a,
	0x68, 0x75, 0xfd, 0x83, 0x1c, 0xbc, 0xc9, 0xad, 0xce, 0xf6, 0x22, 0x7b, 0x49, 0x95, 0x7f, 0x79,
	0x80, 0x3e, 0x49, 0xdf, 0xa0, 0x0f, 0xd0, 0x37, 0xe9, 0x83, 0x54, 0x5e, 0xdb, 0x60, 0x08, 0x29,
	0xff, 0x3c, 0xdf, 0x7c, 0xdf, 0x37, 0xb3, 0x3b, 0xeb, 0x81, 0x23, 0xc6, 0x3b, 0x33, 0xee, 0x79,
	0xdc, 0xef, 0x78, 0x54, 0x04, 0x6c, 0x16, 0xb6, 0xe7, 0x01, 0x17, 0x1c, 0xf7, 0x18, 0x6f, 0xbd,
	0xb8, 0xe5, 0xfc, 0xd6, 0xa5, 0x1d, 0x89, 0x5c, 0x2f, 0x6e, 0x3a, 0x82, 0x79, 0x34, 0x14, 0xb6,
	0x37, 0x8f, 0x49, 0xad, 0xd6, 0x4a, 0x6d, 0x0b, 0x11, 0xb0, 0xeb, 0x85, 0xa0, 0x89, 0xc1, 0xe9,
	0x3f, 0x79, 0x28, 0x5e, 0xc5, 0x96, 0xd8, 0x84, 0xe2, 0x2c, 0xa0, 0xb6, 0xa0, 0x4e, 0x33, 0x77,
	0x92, 0x3b, 0x53, 0xac, 0x34, 0xc4, 0x6f, 0xa0, 0x6e, 0xcf, 0xe7, 0x2e, 0xfd, 0xcd, 0x76, 0x5d,
	0x2a, 0xba, 0x7e, 0x28, 0x6c, 0xd7, 0x0d, 0x9b, 0x7b, 0x92, 0xb5, 0x2d, 0x85, 0x5f, 0x43, 0x2d,
	0x6e, 0x6b, 0x68, 0xdf, 0x2f, 0xf9, 0x79, 0xc9, 0x7f, 0x9c, 0xc0, 0x0b, 0x68, 0x70, 0xf1, 0x91,
	0x06, 0x9a, 0xef, 0x04, 0x9c, 0x39, 0x4b, 0x81, 0x22, 0x05, 0x5b, 0x73, 0xf8, 0x0a, 0x3e, 0xcb,
	0x14, 0x9e, 0xf8, 0x2c, 0x15, 0x95, 0xa5, 0x68, 0x7b, 0x32, 0x3a, 0xc9, 0xb2, 0x7c, 0x46, 0x73,
	0x10, 0x9f, 0x64, 0x4b, 0x0a, 0xbf, 0x83, 0xc3, 0x6c, 0xfd, 0x8c, 0xe8, 0xb9, 0x14, 0x3d, 0x91,
	0xc5, 0x97, 0xf0, 0x9c, 0xf9, 0x77, 0xb6, 0xcb, 0x1c, 0x5b, 0x30, 0xee, 0x87, 0xcd, 0x86, 0xa4,
	0xaf, 0x83, 0xa6, 0xa2, 0xee, 0x93, 0x82, 0xa9, 0xa8, 0x05, 0x52, 0x34, 0x15, 0xb5, 0x48, 0x54,
	0x53, 0x51, 0x55, 0x52, 0x32, 0x15, 0xb5, 0x44, 0xc0, 0x54, 0x54, 0x20, 0x65, 0x53, 0x51, 0x2b,
	0xa4, 0x6a, 0x2a, 0x6a, 0x95, 0x10, 0x53, 0x51, 0x09, 0xa9, 0x99, 0x8a, 0x5a, 0x23, 0x68, 0x2a,
	0x2a, 0x92, 0xba, 0xa9, 0xa8, 0x75, 0xd2, 0x38, 0xfd, 0x03, 0xf2, 0x93, 0xf1, 0x15, 0x1e, 0x42,
	0x21, 0xe4, 0x8b, 0x60, 0x46, 0xe5, 0x04, 0x4b, 0x56, 0x12, 0x45, 0xb8, 0x47, 0x1d, 0xb6, 0xf0,
	0xe4, 0xcc, 0x4a, 0x56, 0x12, 0x21, 0x82, 0xe2, 0xdb, 0x1e, 0x95, 0x93, 0x29, 0x59, 0xf2, 0x3b,
	0xc2, 0x04, 0x0d, 0x3c, 0x79, 0xf9, 0x25, 0x4b, 0x7e, 0xcb, 0xa7, 0xc1, 0x7d, 0x41, 0x7d, 0xd1,
	0xdc, 0x97, 0x70, 0x1a, 0x9e, 0xfe, 0x99, 0x83, 0x83, 0x2b, 0x2a, 0xec, 0x1e, 0x9f, 0xc9, 0x23,
	0xc5, 0xd4, 0x85, 0x2f, 0x82, 0xfb, 0xa4, 0x87, 0x34, 0x8c, 0x8c, 0x67, 0x4c, 0xdc, 0x27, 0x2d,
	0xc8, 0x6f, 0x6c, 0xc0, 0x7e, 0x28, 0x6c, 0x91, 0x76, 0x10, 0x07, 0xd8, 0x02, 0xd5, 0xb5, 0x05,
	0x13, 0x0b, 0x87, 0xca, 0x36, 0x72, 0xd6, 0x32, 0xc6, 0x2f, 0xa0, 0xe4, 0x72, 0xff, 0x36, 0x4e,
	0xee, 0xcb, 0xe4, 0x0a, 0x38, 0xfd, 0x57, 0x05, 0x35, 0x6a, 0xc7, 0xb1, 0x85, 0x8d, 0xa7, 0x50,
	0x88, 0xfc, 0x16, 0xa1, 0xec, 0xa4, 0x72, 0x01, 0x6d, 0xc6, 0xdb, 0x23, 0x89, 0x58, 0x49, 0x06,
	0x7f, 0x84, 0xaa, 0xcb, 0x6e, 0xe8, 0xec, 0x7e, 0xe6, 0x52, 0xe3, 0x8e, 0xfa, 0x22, 0x7a, 0xd6,
	0xf9, 0xb3, 0xca, 0x45, 0x3d, 0x22, 0xf7, 0xd6, 0x53, 0xd6, 0x26, 0x17, 0x8f, 0x21, 0xbf, 0x10,
	0x9e, 0xec, 0xbe, 0x7c, 0x51, 0x8c, 0x24, 0x93, 0xf1, 0x95, 0x15, 0x61, 0xd1, 0xd1, 0x6c, 0x57,
	0x74, 0x9d, 0xe4, 0x22, 0xe3, 0x00, 0xbf, 0x82, 0x0a, 0x0b, 0xc3, 0x05, 0xed, 0xce, 0x35, 0xc7,
	0x09, 0x68, 0x18, 0x26, 0x17, 0xba, 0x81, 0xe2, 0x39, 0x90, 0xe4, 0x29, 0xad, 0x98, 0x05, 0xc9,
	0x7c, 0x84, 0xe3, 0xf7, 0x50, 0x09, 0xa8, 0xef, 0xd0, 0x20, 0x1d, 0x42, 0xb3, 0x28, 0xfb, 0x21,
	0x51, 0x3f, 0xd9, 0xe1, 0x58, 0x1b, 0x3c, 0x3c, 0x83, 0x6a, 0x8c, 0xac, 0x8a, 0xa8, 0xb2, 0xc8,
	0x26, 0x9c, 0xe9, 0x67, 0x12, 0xd2, 0x40, 0xbb, 0x8d, 0x9e, 0x42, 0x69, 0xad, 0x9f, 0x25, 0xbe,
	0x72, 0x5d, 0x51, 0x21, 0xeb, 0xba, 0x62, 0xf6, 0xe1, 0x28, 0x51, 0xeb, 0xf4, 0x8e, 0xcd, 0xa8,
	0xb6, 0xdc, 0x4f, 0xcd, 0xb2, 0x9c, 0x42, 0x23, 0x3a, 0xc2, 0x66, 0xce, 0x7a, 0x4a, 0x84, 0x3d,
	0x38, 0x8c, 0x4b, 0x3c, 0xb2, 0x3b, 0xf8, 0x1f, 0xbb, 0x27, 0x34, 0xf8, 0x0a, 0x8a, 0x72, 0x2a,
	0x9a, 0x90, 0xff, 0x7a, 0xf9, 0xa2, 0xd5, 0x8e, 0x37, 0x44, 0x3b, 0xdd, 0xb5, 0xed, 0x71, 0xba,
	0x6b, 0xad, 0x94, 0x8a, 0x3f, 0x00, 0xc4, 0x7e, 0xd4, 0xd1, 0x44, 0xb3, 0xb2, 0x53, 0x98, 0x61,
	0xe3, 0x1b, 0x20, 0x37, 0x2c, 0x08, 0xd3, 0x3d, 0x2a, 0x1d, 0xaa, 0x3b, 0x1d, 0x1e, 0x69, 0x50,
	0x87, 0xaa, 0x6b, 0xaf, 0xdb, 0x90, 0x9d, 0x36, 0x9b, 0x12, 0x34, 0x01, 0xa5, 0xf3, 0x72, 0xab,
	0x49, 0xa3, 0xda, 0x4e, 0xa3, 0x2d, 0x2a, 0x7c, 0x07, 0xb5, 0xc8, 0x7e, 0xdd, 0x0a, 0x77, 0x5a,
	0x3d, 0x16, 0xe1, 0xeb, 0xcc, 0x62, 0x95, 0x2e, 0xf5, 0x9d, 0x2e, 0xeb, 0x82, 0xc8, 0x41, 0xda,
	0xce, 0x53, 0x87, 0xc6, 0x6e, 0x87, 0x35, 0xc1, 0xf9, 0xdf, 0x7b, 0x50, 0xdd, 0xd8, 0x0d, 0xf8,
	0x1c, 0x4a, 0xfd, 0xc1, 0x74, 0x34, 0xd6, 0xc6, 0x93, 0x11, 0x79, 0x86, 0x07, 0xa0, 0x5a, 0x46,
	0x5f, 0x37, 0x2c, 0x43, 0x27, 0x39, 0xac, 0x43, 0xb5, 0xdb, 0x1f, 0x8d, 0xb5, 0x5e, 0xcf, 0xd0,
	0xa7, 0xda, 0x70, 0xd8, 0x33, 0xc8, 0x1e, 0x36, 0x80, 0xac, 0xc0, 0xb7, 0x83, 0xc1, 0xdb, 0x9e,
	0x41, 0x14, 0xfc, 0x1c, 0x8e, 0x32, 0xd4, 0xbe, 0x6e, 0x0d, 0xba, 0xfa, 0x74, 0x30, 0x7e, 0x67,
	0x58, 0x44, 0xc5, 0x43, 0xa8, 0x4d, 0xfa, 0x9b, 0x4e, 0x0f, 0x2a, 0x1e, 0x01, 0x66, 0xf1, 0xc4,
	0xec, 0x81, 0xe0, 0x97, 0x70, 0xbc, 0x26, 0x58, 0xf3, 0x7b, 0x38, 0x89, 0x84, 0xd2, 0x64, 0x3a,
	0xd4, 0x46, 0xa3, 0xe9, 0xa5, 0x65, 0x68, 0x63, 0x43, 0x27, 0x0f, 0xaf, 0xf1, 0x05, 0x1c, 0xc7,
	0x2e, 0xd3, 0xa1, 0xf6, 0x61, 0x6a, 0x19, 0x97, 0x03, 0x4b, 0x5f, 0xe5, 0x1f, 0x72, 0xd8, 0x82,
	0x46, 0xb7, 0xff, 0xab, 0xd6, 0xeb, 0xea, 0xda, 0xd8, 0x98, 0x5a, 0xc6, 0x2f, 0x13, 0x63, 0x14,
	0xe7, 0xc8, 0x46, 0xee, 0x72, 0xd0, 0x7f, 0xd3, 0xb5, 0xae, 0x64, 0xee, 0xe4, 0xfc, 0x03, 0x14,
	0xe2, 0x1d, 0x8c, 0x55, 0x28, 0xcb, 0xaa, 0xdd, 0xd1, 0x68, 0x62, 0xe8, 0xe4, 0x19, 0x22, 0x54,
	0x62, 0x20, 0xed, 0x97, 0xe4, 0xa2, 0x4b, 0x92, 0x58, 0xe6, 0x14, 0xf1, 0xd5, 0x25, 0xcc, 0xb4,
	0x8a, 0x4e, 0xf2, 0x3f, 0xff, 0x04, 0x55, 0xc6, 0xdb, 0x73, 0x3b, 0x0c, 0x3f, 0x31, 0xd1, 0x1e,
	0xbe, 0x67, 0xfc, 0xf7, 0x97, 0xa1, 0xb0, 0xc3, 0x8f, 0x4b, 0x6c, 0xc6, 0xbd, 0x0e, 0xe3, 0x1d,
	0x8f, 0x3b, 0xd4, 0xed, 0x84, 0xce, 0xa7, 0xce, 0x2d, 0xef, 0x30, 0xfe, 0xd7, 0x9e, 0x32, 0x7c,
	0xdf, 0xe5, 0xd7, 0x05, 0x39, 0xf9, 0x6f, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x09, 0xbf, 0x26,
	0x0e, 0x70, 0x09, 0x00, 0x00,
}
