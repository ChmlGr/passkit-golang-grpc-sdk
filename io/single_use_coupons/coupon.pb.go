// Code generated by protoc-gen-go. DO NOT EDIT.
// source: io/single_use_coupons/coupon.proto

package single_use_coupons

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	math "math"
	io "github.com/PassKit/passkit-golang-grpc-sdk/io"
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

type CouponStatus int32

const (
	// Default status when a coupon is issued.
	CouponStatus_UNREDEEMED CouponStatus = 0
	// Set after a coupon is redeemed (the redeem endpoint is called).
	CouponStatus_REDEEMED CouponStatus = 1
)

var CouponStatus_name = map[int32]string{
	0: "UNREDEEMED",
	1: "REDEEMED",
}

var CouponStatus_value = map[string]int32{
	"UNREDEEMED": 0,
	"REDEEMED":   1,
}

func (x CouponStatus) String() string {
	return proto.EnumName(CouponStatus_name, int32(x))
}

func (CouponStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cd71d037819dd437, []int{0}
}

// The Coupon Details
type Coupon struct {
	// PassKit generated coupon id (22 characters).
	// @tag: validateGeneric:"required" validateCreate:"-" validateUpdate:"required_without=ExternalId"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validateGeneric:"required" validateCreate:"-" validateUpdate:"required_without=ExternalId"`
	// This can be used to set the 'external' ID of the coupon (i.e. the unique coupon code as it's being used in your system). If provided then this can be used to query & update coupon. This field will be treated as unique within the campaign, and cannot be updated at a later stage.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"required_without=Id"
	ExternalId string `protobuf:"bytes,2,opt,name=externalId,proto3" json:"externalId,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"required_without=Id"`
	// The coupon offer id that the coupon belongs to.
	// @tag: validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"omitempty"
	OfferId string `protobuf:"bytes,3,opt,name=offerId,proto3" json:"offerId,omitempty" validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"omitempty"`
	// The coupon campaign id the the coupon belongs to.
	// @tag: validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required_without=Id"
	CampaignId string `protobuf:"bytes,4,opt,name=campaignId,proto3" json:"campaignId,omitempty" validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required_without=Id"`
	// Optional field if you would like to add PII to the individual coupon. Person is stored encrypted.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	Person *io.Person `protobuf:"bytes,5,opt,name=person,proto3" json:"person,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Any other meta fields for the coupon that are not covered in the protocol (i.e. you might want to captures sign-up location, or favourite perfume, for segmenting later on).
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	MetaData map[string]string `protobuf:"bytes,6,rep,name=metaData,proto3" json:"metaData,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Optional SKU. Can be used in the barcode by setting ${coupon.sku} in the template barcode settings.
	Sku string `protobuf:"bytes,7,opt,name=sku,proto3" json:"sku,omitempty"`
	// Indicates if the user of the coupon opted out to receive marketing promotions (defaults to false).
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	OptOut bool `protobuf:"varint,8,opt,name=optOut,proto3" json:"optOut,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Holds the coupon status.
	Status CouponStatus `protobuf:"varint,10,opt,name=status,proto3,enum=single_use_coupons.CouponStatus" json:"status,omitempty"`
	// Optional field to set when the coupon expires. When the expiry date hits, it will automatically gray out the barcode. Only applicable if the template for the offer is set to EXPIRE_ON_VARIABLE_DATE_TIME.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	ExpiryDate *timestamp.Timestamp `protobuf:"bytes,11,opt,name=expiryDate,proto3" json:"expiryDate,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Contains details about the redemption (if the coupon is in a redeemed state).
	RedemptionDetails *RedemptionDetails `protobuf:"bytes,12,opt,name=redemptionDetails,proto3" json:"redemptionDetails,omitempty"`
	// Optional Pass level overrides for proximity & images.
	PassOverrides *io.PassOverrides `protobuf:"bytes,13,opt,name=passOverrides,proto3" json:"passOverrides,omitempty"`
	// Pass Metadata hold meta data about the pass like status, lifecycle, utm details and installation data.
	PassMetaData *io.Metadata `protobuf:"bytes,14,opt,name=passMetaData,proto3" json:"passMetaData,omitempty"`
	// The date the coupon was created.
	// @tag: validateGeneric:"-" validateCreate:"-" validateUpdate:"-"
	Created *timestamp.Timestamp `protobuf:"bytes,15,opt,name=created,proto3" json:"created,omitempty" validateGeneric:"-" validateCreate:"-" validateUpdate:"-"`
	// The date the coupon was last updated.
	// @tag: validateGeneric:"-" validateCreate:"-" validateUpdate:"-"
	Updated              *timestamp.Timestamp `protobuf:"bytes,16,opt,name=updated,proto3" json:"updated,omitempty" validateGeneric:"-" validateCreate:"-" validateUpdate:"-"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Coupon) Reset()         { *m = Coupon{} }
func (m *Coupon) String() string { return proto.CompactTextString(m) }
func (*Coupon) ProtoMessage()    {}
func (*Coupon) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd71d037819dd437, []int{0}
}

func (m *Coupon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coupon.Unmarshal(m, b)
}
func (m *Coupon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coupon.Marshal(b, m, deterministic)
}
func (m *Coupon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coupon.Merge(m, src)
}
func (m *Coupon) XXX_Size() int {
	return xxx_messageInfo_Coupon.Size(m)
}
func (m *Coupon) XXX_DiscardUnknown() {
	xxx_messageInfo_Coupon.DiscardUnknown(m)
}

var xxx_messageInfo_Coupon proto.InternalMessageInfo

func (m *Coupon) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Coupon) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

func (m *Coupon) GetOfferId() string {
	if m != nil {
		return m.OfferId
	}
	return ""
}

func (m *Coupon) GetCampaignId() string {
	if m != nil {
		return m.CampaignId
	}
	return ""
}

func (m *Coupon) GetPerson() *io.Person {
	if m != nil {
		return m.Person
	}
	return nil
}

func (m *Coupon) GetMetaData() map[string]string {
	if m != nil {
		return m.MetaData
	}
	return nil
}

func (m *Coupon) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *Coupon) GetOptOut() bool {
	if m != nil {
		return m.OptOut
	}
	return false
}

func (m *Coupon) GetStatus() CouponStatus {
	if m != nil {
		return m.Status
	}
	return CouponStatus_UNREDEEMED
}

func (m *Coupon) GetExpiryDate() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiryDate
	}
	return nil
}

func (m *Coupon) GetRedemptionDetails() *RedemptionDetails {
	if m != nil {
		return m.RedemptionDetails
	}
	return nil
}

func (m *Coupon) GetPassOverrides() *io.PassOverrides {
	if m != nil {
		return m.PassOverrides
	}
	return nil
}

func (m *Coupon) GetPassMetaData() *io.Metadata {
	if m != nil {
		return m.PassMetaData
	}
	return nil
}

func (m *Coupon) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Coupon) GetUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

type RedemptionDetails struct {
	// The date the coupon was redeemed (set by the system when the redeem endpoint is called).
	RedemptionDate *timestamp.Timestamp `protobuf:"bytes,1,opt,name=redemptionDate,proto3" json:"redemptionDate,omitempty"`
	// Optional field that contains the redemption code provided on redeem (either the fixed one if set, or user provided).
	RedemptionCode string `protobuf:"bytes,2,opt,name=redemptionCode,proto3" json:"redemptionCode,omitempty"`
	// Option redemption GPS lat, lon &  alt.
	Lat float64 `protobuf:"fixed64,3,opt,name=lat,proto3" json:"lat,omitempty"`
	// Longitude.
	Lon float64 `protobuf:"fixed64,4,opt,name=lon,proto3" json:"lon,omitempty"`
	// Altitude in metres.
	Alt int32 `protobuf:"varint,5,opt,name=alt,proto3" json:"alt,omitempty"`
	// Optional redemption user agent, can be used in case where a Merchant uses the PassKit scanning solution combined with their own POS redemption and wants to differentiate the source of redemption.
	RedemptionSource string `protobuf:"bytes,6,opt,name=redemptionSource,proto3" json:"redemptionSource,omitempty"`
	// Optional redemption reference, can be used for staff name, username, POS ID, location ID, etc.
	RedemptionReference string `protobuf:"bytes,7,opt,name=redemptionReference,proto3" json:"redemptionReference,omitempty"`
	// Optional transaction reference, can be used to consolidate baskets at a later stage via a manual process.
	TransactionReference string `protobuf:"bytes,8,opt,name=transactionReference,proto3" json:"transactionReference,omitempty"`
	// Optional total transaction amount.
	TransactionAmount    float64  `protobuf:"fixed64,9,opt,name=transactionAmount,proto3" json:"transactionAmount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedemptionDetails) Reset()         { *m = RedemptionDetails{} }
func (m *RedemptionDetails) String() string { return proto.CompactTextString(m) }
func (*RedemptionDetails) ProtoMessage()    {}
func (*RedemptionDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd71d037819dd437, []int{1}
}

func (m *RedemptionDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedemptionDetails.Unmarshal(m, b)
}
func (m *RedemptionDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedemptionDetails.Marshal(b, m, deterministic)
}
func (m *RedemptionDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedemptionDetails.Merge(m, src)
}
func (m *RedemptionDetails) XXX_Size() int {
	return xxx_messageInfo_RedemptionDetails.Size(m)
}
func (m *RedemptionDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_RedemptionDetails.DiscardUnknown(m)
}

var xxx_messageInfo_RedemptionDetails proto.InternalMessageInfo

func (m *RedemptionDetails) GetRedemptionDate() *timestamp.Timestamp {
	if m != nil {
		return m.RedemptionDate
	}
	return nil
}

func (m *RedemptionDetails) GetRedemptionCode() string {
	if m != nil {
		return m.RedemptionCode
	}
	return ""
}

func (m *RedemptionDetails) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *RedemptionDetails) GetLon() float64 {
	if m != nil {
		return m.Lon
	}
	return 0
}

func (m *RedemptionDetails) GetAlt() int32 {
	if m != nil {
		return m.Alt
	}
	return 0
}

func (m *RedemptionDetails) GetRedemptionSource() string {
	if m != nil {
		return m.RedemptionSource
	}
	return ""
}

func (m *RedemptionDetails) GetRedemptionReference() string {
	if m != nil {
		return m.RedemptionReference
	}
	return ""
}

func (m *RedemptionDetails) GetTransactionReference() string {
	if m != nil {
		return m.TransactionReference
	}
	return ""
}

func (m *RedemptionDetails) GetTransactionAmount() float64 {
	if m != nil {
		return m.TransactionAmount
	}
	return 0
}

type ExternalIdRequest struct {
	// Coupon Campaign ID
	// @tag: validateGeneric:"required"
	CouponCampaignId string `protobuf:"bytes,1,opt,name=couponCampaignId,proto3" json:"couponCampaignId,omitempty" validateGeneric:"required"`
	// External ID
	// @tag: validateGeneric:"required"
	ExternalId           string   `protobuf:"bytes,2,opt,name=externalId,proto3" json:"externalId,omitempty" validateGeneric:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExternalIdRequest) Reset()         { *m = ExternalIdRequest{} }
func (m *ExternalIdRequest) String() string { return proto.CompactTextString(m) }
func (*ExternalIdRequest) ProtoMessage()    {}
func (*ExternalIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd71d037819dd437, []int{2}
}

func (m *ExternalIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExternalIdRequest.Unmarshal(m, b)
}
func (m *ExternalIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExternalIdRequest.Marshal(b, m, deterministic)
}
func (m *ExternalIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalIdRequest.Merge(m, src)
}
func (m *ExternalIdRequest) XXX_Size() int {
	return xxx_messageInfo_ExternalIdRequest.Size(m)
}
func (m *ExternalIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalIdRequest proto.InternalMessageInfo

func (m *ExternalIdRequest) GetCouponCampaignId() string {
	if m != nil {
		return m.CouponCampaignId
	}
	return ""
}

func (m *ExternalIdRequest) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

type ListRequestDeprecated struct {
	// Coupon Campaign ID
	// @tag: validateGeneric:"required" validateCreate:"required" validateUpdate:"required"
	CouponCampaignId string `protobuf:"bytes,1,opt,name=couponCampaignId,proto3" json:"couponCampaignId,omitempty" validateGeneric:"required" validateCreate:"required" validateUpdate:"required"`
	// External ID
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	Pagination           *io.Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListRequestDeprecated) Reset()         { *m = ListRequestDeprecated{} }
func (m *ListRequestDeprecated) String() string { return proto.CompactTextString(m) }
func (*ListRequestDeprecated) ProtoMessage()    {}
func (*ListRequestDeprecated) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd71d037819dd437, []int{3}
}

func (m *ListRequestDeprecated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequestDeprecated.Unmarshal(m, b)
}
func (m *ListRequestDeprecated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequestDeprecated.Marshal(b, m, deterministic)
}
func (m *ListRequestDeprecated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequestDeprecated.Merge(m, src)
}
func (m *ListRequestDeprecated) XXX_Size() int {
	return xxx_messageInfo_ListRequestDeprecated.Size(m)
}
func (m *ListRequestDeprecated) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequestDeprecated.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequestDeprecated proto.InternalMessageInfo

func (m *ListRequestDeprecated) GetCouponCampaignId() string {
	if m != nil {
		return m.CouponCampaignId
	}
	return ""
}

func (m *ListRequestDeprecated) GetPagination() *io.Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type ListRequest struct {
	// Coupon Campaign ID
	// @tag: validateGeneric:"required" validateCreate:"required" validateUpdate:"required"
	CouponCampaignId string `protobuf:"bytes,1,opt,name=couponCampaignId,proto3" json:"couponCampaignId,omitempty" validateGeneric:"required" validateCreate:"required" validateUpdate:"required"`
	// External ID
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	Filters              *io.Filters `protobuf:"bytes,2,opt,name=filters,proto3" json:"filters,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd71d037819dd437, []int{4}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetCouponCampaignId() string {
	if m != nil {
		return m.CouponCampaignId
	}
	return ""
}

func (m *ListRequest) GetFilters() *io.Filters {
	if m != nil {
		return m.Filters
	}
	return nil
}

func init() {
	proto.RegisterEnum("single_use_coupons.CouponStatus", CouponStatus_name, CouponStatus_value)
	proto.RegisterType((*Coupon)(nil), "single_use_coupons.Coupon")
	proto.RegisterMapType((map[string]string)(nil), "single_use_coupons.Coupon.MetaDataEntry")
	proto.RegisterType((*RedemptionDetails)(nil), "single_use_coupons.RedemptionDetails")
	proto.RegisterType((*ExternalIdRequest)(nil), "single_use_coupons.ExternalIdRequest")
	proto.RegisterType((*ListRequestDeprecated)(nil), "single_use_coupons.ListRequestDeprecated")
	proto.RegisterType((*ListRequest)(nil), "single_use_coupons.ListRequest")
}

func init() {
	proto.RegisterFile("io/single_use_coupons/coupon.proto", fileDescriptor_cd71d037819dd437)
}

var fileDescriptor_cd71d037819dd437 = []byte{
	// 877 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x9e, 0x9c, 0xc6, 0x71, 0x8e, 0x13, 0xcf, 0xe6, 0xb2, 0x4e, 0xf0, 0xc5, 0x46, 0x18, 0x68,
	0xe1, 0x05, 0xa9, 0x54, 0x78, 0x03, 0x5a, 0x64, 0x57, 0x6d, 0xec, 0x61, 0x59, 0xd7, 0x25, 0x50,
	0xd6, 0x9b, 0xdd, 0x64, 0xac, 0x74, 0xac, 0x12, 0x91, 0x48, 0x95, 0xa4, 0xb2, 0xe5, 0x35, 0xf6,
	0x18, 0x7b, 0x9c, 0x3e, 0xc2, 0xde, 0x61, 0xf7, 0x03, 0x29, 0x29, 0x56, 0xe2, 0x14, 0x41, 0xaf,
	0x44, 0x7e, 0xe7, 0xfb, 0x0e, 0x79, 0xfe, 0x28, 0x98, 0x70, 0x19, 0x6a, 0x2e, 0xd2, 0x0c, 0xcf,
	0x4b, 0x8d, 0xe7, 0xb1, 0x2c, 0x0b, 0x29, 0x74, 0x58, 0x7d, 0x83, 0x42, 0x49, 0x23, 0x09, 0x59,
	0x27, 0x8c, 0xbf, 0x49, 0xa5, 0x4c, 0x33, 0x0c, 0x1d, 0xe3, 0x6d, 0xb9, 0x0c, 0x0d, 0xcf, 0x51,
	0x1b, 0x96, 0x17, 0x95, 0x68, 0x3c, 0xe6, 0x32, 0x8c, 0x65, 0x9e, 0x4b, 0x11, 0x16, 0x2c, 0xe5,
	0x82, 0x19, 0xde, 0x38, 0x1c, 0x3f, 0x5c, 0xd9, 0x96, 0x3c, 0x33, 0xa8, 0x6a, 0x7c, 0xaf, 0xad,
	0xd1, 0xba, 0x46, 0xfd, 0x16, 0x8a, 0x4a, 0x4b, 0xc1, 0xb2, 0xda, 0xf2, 0xd5, 0xca, 0x92, 0xa3,
	0x51, 0x3c, 0x6e, 0x24, 0x07, 0xee, 0x13, 0x3f, 0x49, 0x51, 0x3c, 0xd1, 0x7f, 0xb2, 0x34, 0x45,
	0x15, 0xca, 0xc2, 0xde, 0x40, 0x87, 0x4c, 0x08, 0x69, 0xdc, 0x6d, 0x6a, 0xf6, 0xe4, 0xdf, 0x2e,
	0x74, 0x8f, 0x5c, 0x5c, 0x64, 0x00, 0x1d, 0x9e, 0xf8, 0x1e, 0xf5, 0xa6, 0xdb, 0x51, 0x87, 0x27,
	0xe4, 0x6b, 0x00, 0xfc, 0xcb, 0xa0, 0x12, 0x2c, 0x3b, 0x4e, 0xfc, 0x8e, 0xc3, 0x5b, 0x08, 0xf1,
	0x61, 0x4b, 0x2e, 0x97, 0xa8, 0x8e, 0x13, 0x7f, 0xc3, 0x19, 0x9b, 0xad, 0x55, 0xc6, 0x2c, 0x2f,
	0x18, 0x4f, 0xc5, 0x71, 0xe2, 0x3f, 0xa8, 0x94, 0x2b, 0x84, 0x4c, 0xa0, 0x5b, 0x45, 0xe3, 0x6f,
	0x52, 0x6f, 0xda, 0x9f, 0x41, 0xc0, 0x65, 0x70, 0xea, 0x90, 0xa8, 0xb6, 0x90, 0x39, 0xf4, 0x72,
	0x34, 0x6c, 0xce, 0x0c, 0xf3, 0xbb, 0x74, 0x63, 0xda, 0x9f, 0x4d, 0x83, 0xf5, 0x5a, 0x04, 0xd5,
	0xdd, 0x83, 0xd7, 0x35, 0x75, 0x21, 0x8c, 0xba, 0x8a, 0xae, 0x95, 0x64, 0x08, 0x1b, 0xfa, 0xa2,
	0xf4, 0xb7, 0xdc, 0x15, 0xec, 0x92, 0x3c, 0x84, 0xae, 0x2c, 0xcc, 0x49, 0x69, 0xfc, 0x1e, 0xf5,
	0xa6, 0xbd, 0xa8, 0xde, 0x91, 0xe7, 0xd0, 0xd5, 0x86, 0x99, 0x52, 0xfb, 0x40, 0xbd, 0xe9, 0x60,
	0x46, 0x3f, 0x7e, 0xda, 0x99, 0xe3, 0x45, 0x35, 0x9f, 0x1c, 0xda, 0x3c, 0x15, 0x5c, 0x5d, 0xcd,
	0x99, 0x41, 0xbf, 0xef, 0x22, 0x1a, 0x07, 0x55, 0x8f, 0x04, 0x4d, 0x8f, 0x04, 0xbf, 0x35, 0x3d,
	0x12, 0xb5, 0xd8, 0xe4, 0x0c, 0x46, 0x0a, 0x13, 0xcc, 0x5d, 0x7d, 0xe6, 0x68, 0x18, 0xcf, 0xb4,
	0xbf, 0xe3, 0x5c, 0x3c, 0xba, 0xeb, 0x02, 0xd1, 0x6d, 0x72, 0xb4, 0xae, 0x27, 0xcf, 0x60, 0xd7,
	0xb6, 0xd0, 0xc9, 0x25, 0x2a, 0xc5, 0x13, 0xd4, 0xfe, 0xae, 0x73, 0x38, 0x72, 0x59, 0x6e, 0x1b,
	0xa2, 0x9b, 0x3c, 0xf2, 0x14, 0x76, 0x2c, 0xd0, 0x24, 0xd3, 0x1f, 0x38, 0xdd, 0x8e, 0xd5, 0x59,
	0x2c, 0x61, 0x86, 0x45, 0x37, 0x18, 0xe4, 0x7b, 0xd8, 0x8a, 0x15, 0x32, 0x83, 0x89, 0xff, 0xf9,
	0xbd, 0x81, 0x37, 0x54, 0xab, 0x2a, 0x8b, 0xc4, 0xa9, 0x86, 0xf7, 0xab, 0x6a, 0xea, 0xf8, 0x07,
	0xd8, 0xbd, 0x51, 0x66, 0x5b, 0xdc, 0x0b, 0xbc, 0xaa, 0x3b, 0xd6, 0x2e, 0xc9, 0x1e, 0x6c, 0x5e,
	0xb2, 0xac, 0xc4, 0xba, 0x5b, 0xab, 0xcd, 0x61, 0xe7, 0xb9, 0x77, 0x58, 0xfe, 0xfd, 0x42, 0x41,
	0xb1, 0x5f, 0xf7, 0xfa, 0xec, 0xe4, 0x27, 0x99, 0x25, 0x9a, 0x9a, 0x77, 0x48, 0x93, 0x2a, 0x6b,
	0xf4, 0x31, 0xb5, 0xf5, 0x44, 0x2a, 0x97, 0x94, 0xd1, 0x52, 0xf0, 0xf7, 0x25, 0x52, 0x2e, 0x12,
	0x7e, 0xc9, 0x93, 0x92, 0x65, 0xb4, 0xaa, 0x00, 0xe5, 0x42, 0x1b, 0x26, 0x62, 0xa4, 0x53, 0x1e,
	0x60, 0x40, 0x19, 0xb5, 0xe9, 0xf8, 0x36, 0xf8, 0xe0, 0x35, 0xcd, 0xff, 0xc1, 0x6b, 0x75, 0xfa,
	0xcf, 0x0f, 0x7a, 0xdb, 0x43, 0x98, 0xfc, 0xd7, 0x81, 0xd1, 0x5a, 0xe5, 0xc8, 0x4b, 0x18, 0xb4,
	0x6a, 0x67, 0x7b, 0xc7, 0xbb, 0x37, 0x19, 0xb7, 0x14, 0xe4, 0x71, 0xdb, 0xc7, 0x91, 0x4c, 0x9a,
	0xc8, 0x6f, 0xa1, 0x36, 0x55, 0x19, 0x33, 0x6e, 0x4e, 0xbd, 0xc8, 0x2e, 0x1d, 0x22, 0x85, 0x1b,
	0x4e, 0x8b, 0x48, 0x61, 0x11, 0x96, 0x19, 0x37, 0x92, 0x9b, 0x91, 0x5d, 0x92, 0x7d, 0x18, 0xae,
	0xfc, 0x9c, 0xc9, 0x52, 0xc5, 0xe8, 0x77, 0x9d, 0xff, 0x35, 0x9c, 0x3c, 0x85, 0x2f, 0x56, 0x58,
	0x84, 0x4b, 0x54, 0x28, 0x62, 0xac, 0x27, 0xef, 0x2e, 0x13, 0x99, 0xc1, 0x9e, 0x51, 0x4c, 0x68,
	0x16, 0xdf, 0x94, 0xf4, 0x9c, 0xe4, 0x4e, 0x1b, 0x39, 0x80, 0x51, 0x0b, 0x7f, 0x91, 0xcb, 0x52,
	0x18, 0x7f, 0xdb, 0xc5, 0xb0, 0x6e, 0x98, 0x9c, 0xc3, 0x68, 0x71, 0xfd, 0x5e, 0x45, 0xf8, 0xbe,
	0x44, 0xed, 0x82, 0xaa, 0x6a, 0x79, 0xb4, 0x7a, 0xa2, 0xaa, 0x16, 0x5a, 0xc3, 0xef, 0x7b, 0x02,
	0x27, 0x1a, 0xbe, 0xfc, 0x85, 0x6b, 0x53, 0xbb, 0x9e, 0x63, 0xa1, 0x30, 0x76, 0x1d, 0xfe, 0x29,
	0x87, 0x04, 0x00, 0xab, 0xbf, 0x84, 0x3b, 0xa4, 0x3f, 0x1b, 0x54, 0xb3, 0xda, 0xa0, 0x51, 0x8b,
	0x31, 0xf9, 0x03, 0xfa, 0xad, 0x43, 0x3f, 0xe9, 0xa8, 0x47, 0xb0, 0x55, 0xfd, 0x74, 0x74, 0x7d,
	0x4e, 0xdf, 0x9e, 0xf3, 0x63, 0x05, 0x45, 0x8d, 0x6d, 0xff, 0x00, 0x76, 0xda, 0x2f, 0x1d, 0x19,
	0x00, 0xbc, 0xf9, 0x35, 0x5a, 0xcc, 0x17, 0x8b, 0xd7, 0x8b, 0xf9, 0xf0, 0x33, 0xb2, 0x03, 0xbd,
	0xeb, 0x9d, 0xf7, 0xb2, 0x00, 0xca, 0x65, 0x60, 0xe7, 0xe0, 0x82, 0x9b, 0xe0, 0xf4, 0x15, 0x97,
	0xc1, 0x99, 0x7b, 0xb9, 0xde, 0x68, 0xac, 0xdc, 0xe8, 0xdf, 0x9f, 0x69, 0xc3, 0xf4, 0xbb, 0x6b,
	0x52, 0x2c, 0xf3, 0x90, 0xcb, 0x30, 0x97, 0x09, 0x66, 0xa1, 0x4e, 0x2e, 0xc2, 0x54, 0x86, 0x77,
	0xfe, 0x8c, 0xff, 0xe9, 0x90, 0xd3, 0x57, 0xb7, 0xbd, 0xbd, 0xed, 0xba, 0xc9, 0xf8, 0xee, 0xff,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0x25, 0x3d, 0x89, 0xc1, 0x07, 0x00, 0x00,
}
