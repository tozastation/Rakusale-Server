// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shop.proto

package shop

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type RequestShop struct {
	ImagePath            string   `protobuf:"bytes,1,opt,name=image_path,json=imagePath,proto3" json:"image_path,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Latitude             float32  `protobuf:"fixed32,3,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float32  `protobuf:"fixed32,4,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Introduction         string   `protobuf:"bytes,5,opt,name=introduction,proto3" json:"introduction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestShop) Reset()         { *m = RequestShop{} }
func (m *RequestShop) String() string { return proto.CompactTextString(m) }
func (*RequestShop) ProtoMessage()    {}
func (*RequestShop) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{0}
}

func (m *RequestShop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestShop.Unmarshal(m, b)
}
func (m *RequestShop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestShop.Marshal(b, m, deterministic)
}
func (m *RequestShop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestShop.Merge(m, src)
}
func (m *RequestShop) XXX_Size() int {
	return xxx_messageInfo_RequestShop.Size(m)
}
func (m *RequestShop) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestShop.DiscardUnknown(m)
}

var xxx_messageInfo_RequestShop proto.InternalMessageInfo

func (m *RequestShop) GetImagePath() string {
	if m != nil {
		return m.ImagePath
	}
	return ""
}

func (m *RequestShop) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RequestShop) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *RequestShop) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *RequestShop) GetIntroduction() string {
	if m != nil {
		return m.Introduction
	}
	return ""
}

type ResponseShop struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ImagePath            string   `protobuf:"bytes,2,opt,name=image_path,json=imagePath,proto3" json:"image_path,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Latitude             float32  `protobuf:"fixed32,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float32  `protobuf:"fixed32,5,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Introduction         string   `protobuf:"bytes,6,opt,name=introduction,proto3" json:"introduction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseShop) Reset()         { *m = ResponseShop{} }
func (m *ResponseShop) String() string { return proto.CompactTextString(m) }
func (*ResponseShop) ProtoMessage()    {}
func (*ResponseShop) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{1}
}

func (m *ResponseShop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseShop.Unmarshal(m, b)
}
func (m *ResponseShop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseShop.Marshal(b, m, deterministic)
}
func (m *ResponseShop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseShop.Merge(m, src)
}
func (m *ResponseShop) XXX_Size() int {
	return xxx_messageInfo_ResponseShop.Size(m)
}
func (m *ResponseShop) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseShop.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseShop proto.InternalMessageInfo

func (m *ResponseShop) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ResponseShop) GetImagePath() string {
	if m != nil {
		return m.ImagePath
	}
	return ""
}

func (m *ResponseShop) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResponseShop) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *ResponseShop) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *ResponseShop) GetIntroduction() string {
	if m != nil {
		return m.Introduction
	}
	return ""
}

// Read
type GetMyShopRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMyShopRequest) Reset()         { *m = GetMyShopRequest{} }
func (m *GetMyShopRequest) String() string { return proto.CompactTextString(m) }
func (*GetMyShopRequest) ProtoMessage()    {}
func (*GetMyShopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{2}
}

func (m *GetMyShopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMyShopRequest.Unmarshal(m, b)
}
func (m *GetMyShopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMyShopRequest.Marshal(b, m, deterministic)
}
func (m *GetMyShopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMyShopRequest.Merge(m, src)
}
func (m *GetMyShopRequest) XXX_Size() int {
	return xxx_messageInfo_GetMyShopRequest.Size(m)
}
func (m *GetMyShopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMyShopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMyShopRequest proto.InternalMessageInfo

func (m *GetMyShopRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetMyShopResponse struct {
	Status               int64         `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Shop                 *ResponseShop `protobuf:"bytes,2,opt,name=shop,proto3" json:"shop,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetMyShopResponse) Reset()         { *m = GetMyShopResponse{} }
func (m *GetMyShopResponse) String() string { return proto.CompactTextString(m) }
func (*GetMyShopResponse) ProtoMessage()    {}
func (*GetMyShopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{3}
}

func (m *GetMyShopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMyShopResponse.Unmarshal(m, b)
}
func (m *GetMyShopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMyShopResponse.Marshal(b, m, deterministic)
}
func (m *GetMyShopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMyShopResponse.Merge(m, src)
}
func (m *GetMyShopResponse) XXX_Size() int {
	return xxx_messageInfo_GetMyShopResponse.Size(m)
}
func (m *GetMyShopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMyShopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMyShopResponse proto.InternalMessageInfo

func (m *GetMyShopResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *GetMyShopResponse) GetShop() *ResponseShop {
	if m != nil {
		return m.Shop
	}
	return nil
}

type ShopsEmpty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShopsEmpty) Reset()         { *m = ShopsEmpty{} }
func (m *ShopsEmpty) String() string { return proto.CompactTextString(m) }
func (*ShopsEmpty) ProtoMessage()    {}
func (*ShopsEmpty) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{4}
}

func (m *ShopsEmpty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShopsEmpty.Unmarshal(m, b)
}
func (m *ShopsEmpty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShopsEmpty.Marshal(b, m, deterministic)
}
func (m *ShopsEmpty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShopsEmpty.Merge(m, src)
}
func (m *ShopsEmpty) XXX_Size() int {
	return xxx_messageInfo_ShopsEmpty.Size(m)
}
func (m *ShopsEmpty) XXX_DiscardUnknown() {
	xxx_messageInfo_ShopsEmpty.DiscardUnknown(m)
}

var xxx_messageInfo_ShopsEmpty proto.InternalMessageInfo

type GetAllShopsResponse struct {
	Status               int64           `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Shops                []*ResponseShop `protobuf:"bytes,2,rep,name=shops,proto3" json:"shops,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetAllShopsResponse) Reset()         { *m = GetAllShopsResponse{} }
func (m *GetAllShopsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllShopsResponse) ProtoMessage()    {}
func (*GetAllShopsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{5}
}

func (m *GetAllShopsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllShopsResponse.Unmarshal(m, b)
}
func (m *GetAllShopsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllShopsResponse.Marshal(b, m, deterministic)
}
func (m *GetAllShopsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllShopsResponse.Merge(m, src)
}
func (m *GetAllShopsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllShopsResponse.Size(m)
}
func (m *GetAllShopsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllShopsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllShopsResponse proto.InternalMessageInfo

func (m *GetAllShopsResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *GetAllShopsResponse) GetShops() []*ResponseShop {
	if m != nil {
		return m.Shops
	}
	return nil
}

// Create
type ShopImage struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShopImage) Reset()         { *m = ShopImage{} }
func (m *ShopImage) String() string { return proto.CompactTextString(m) }
func (*ShopImage) ProtoMessage()    {}
func (*ShopImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{6}
}

func (m *ShopImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShopImage.Unmarshal(m, b)
}
func (m *ShopImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShopImage.Marshal(b, m, deterministic)
}
func (m *ShopImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShopImage.Merge(m, src)
}
func (m *ShopImage) XXX_Size() int {
	return xxx_messageInfo_ShopImage.Size(m)
}
func (m *ShopImage) XXX_DiscardUnknown() {
	xxx_messageInfo_ShopImage.DiscardUnknown(m)
}

var xxx_messageInfo_ShopImage proto.InternalMessageInfo

func (m *ShopImage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PostMyShopRequest struct {
	Token                string       `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Image                *ShopImage   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Shop                 *RequestShop `protobuf:"bytes,3,opt,name=shop,proto3" json:"shop,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PostMyShopRequest) Reset()         { *m = PostMyShopRequest{} }
func (m *PostMyShopRequest) String() string { return proto.CompactTextString(m) }
func (*PostMyShopRequest) ProtoMessage()    {}
func (*PostMyShopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{7}
}

func (m *PostMyShopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostMyShopRequest.Unmarshal(m, b)
}
func (m *PostMyShopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostMyShopRequest.Marshal(b, m, deterministic)
}
func (m *PostMyShopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostMyShopRequest.Merge(m, src)
}
func (m *PostMyShopRequest) XXX_Size() int {
	return xxx_messageInfo_PostMyShopRequest.Size(m)
}
func (m *PostMyShopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostMyShopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostMyShopRequest proto.InternalMessageInfo

func (m *PostMyShopRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PostMyShopRequest) GetImage() *ShopImage {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *PostMyShopRequest) GetShop() *RequestShop {
	if m != nil {
		return m.Shop
	}
	return nil
}

type PostMyShopResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostMyShopResponse) Reset()         { *m = PostMyShopResponse{} }
func (m *PostMyShopResponse) String() string { return proto.CompactTextString(m) }
func (*PostMyShopResponse) ProtoMessage()    {}
func (*PostMyShopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{8}
}

func (m *PostMyShopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostMyShopResponse.Unmarshal(m, b)
}
func (m *PostMyShopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostMyShopResponse.Marshal(b, m, deterministic)
}
func (m *PostMyShopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostMyShopResponse.Merge(m, src)
}
func (m *PostMyShopResponse) XXX_Size() int {
	return xxx_messageInfo_PostMyShopResponse.Size(m)
}
func (m *PostMyShopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostMyShopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostMyShopResponse proto.InternalMessageInfo

func (m *PostMyShopResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

// Update
type PutMyShopRequest struct {
	Token                string       `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Shop                 *RequestShop `protobuf:"bytes,2,opt,name=shop,proto3" json:"shop,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PutMyShopRequest) Reset()         { *m = PutMyShopRequest{} }
func (m *PutMyShopRequest) String() string { return proto.CompactTextString(m) }
func (*PutMyShopRequest) ProtoMessage()    {}
func (*PutMyShopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{9}
}

func (m *PutMyShopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutMyShopRequest.Unmarshal(m, b)
}
func (m *PutMyShopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutMyShopRequest.Marshal(b, m, deterministic)
}
func (m *PutMyShopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutMyShopRequest.Merge(m, src)
}
func (m *PutMyShopRequest) XXX_Size() int {
	return xxx_messageInfo_PutMyShopRequest.Size(m)
}
func (m *PutMyShopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutMyShopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutMyShopRequest proto.InternalMessageInfo

func (m *PutMyShopRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PutMyShopRequest) GetShop() *RequestShop {
	if m != nil {
		return m.Shop
	}
	return nil
}

type PutMyShopResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutMyShopResponse) Reset()         { *m = PutMyShopResponse{} }
func (m *PutMyShopResponse) String() string { return proto.CompactTextString(m) }
func (*PutMyShopResponse) ProtoMessage()    {}
func (*PutMyShopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{10}
}

func (m *PutMyShopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutMyShopResponse.Unmarshal(m, b)
}
func (m *PutMyShopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutMyShopResponse.Marshal(b, m, deterministic)
}
func (m *PutMyShopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutMyShopResponse.Merge(m, src)
}
func (m *PutMyShopResponse) XXX_Size() int {
	return xxx_messageInfo_PutMyShopResponse.Size(m)
}
func (m *PutMyShopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutMyShopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutMyShopResponse proto.InternalMessageInfo

func (m *PutMyShopResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

// Delete
type DeleteMyShopRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ShopId               int64    `protobuf:"varint,2,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMyShopRequest) Reset()         { *m = DeleteMyShopRequest{} }
func (m *DeleteMyShopRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteMyShopRequest) ProtoMessage()    {}
func (*DeleteMyShopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{11}
}

func (m *DeleteMyShopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMyShopRequest.Unmarshal(m, b)
}
func (m *DeleteMyShopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMyShopRequest.Marshal(b, m, deterministic)
}
func (m *DeleteMyShopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMyShopRequest.Merge(m, src)
}
func (m *DeleteMyShopRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteMyShopRequest.Size(m)
}
func (m *DeleteMyShopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMyShopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMyShopRequest proto.InternalMessageInfo

func (m *DeleteMyShopRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *DeleteMyShopRequest) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

type DeleteMyShopResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMyShopResponse) Reset()         { *m = DeleteMyShopResponse{} }
func (m *DeleteMyShopResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteMyShopResponse) ProtoMessage()    {}
func (*DeleteMyShopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{12}
}

func (m *DeleteMyShopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMyShopResponse.Unmarshal(m, b)
}
func (m *DeleteMyShopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMyShopResponse.Marshal(b, m, deterministic)
}
func (m *DeleteMyShopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMyShopResponse.Merge(m, src)
}
func (m *DeleteMyShopResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteMyShopResponse.Size(m)
}
func (m *DeleteMyShopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMyShopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMyShopResponse proto.InternalMessageInfo

func (m *DeleteMyShopResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*RequestShop)(nil), "shop.RequestShop")
	proto.RegisterType((*ResponseShop)(nil), "shop.ResponseShop")
	proto.RegisterType((*GetMyShopRequest)(nil), "shop.GetMyShopRequest")
	proto.RegisterType((*GetMyShopResponse)(nil), "shop.GetMyShopResponse")
	proto.RegisterType((*ShopsEmpty)(nil), "shop.ShopsEmpty")
	proto.RegisterType((*GetAllShopsResponse)(nil), "shop.GetAllShopsResponse")
	proto.RegisterType((*ShopImage)(nil), "shop.ShopImage")
	proto.RegisterType((*PostMyShopRequest)(nil), "shop.PostMyShopRequest")
	proto.RegisterType((*PostMyShopResponse)(nil), "shop.PostMyShopResponse")
	proto.RegisterType((*PutMyShopRequest)(nil), "shop.PutMyShopRequest")
	proto.RegisterType((*PutMyShopResponse)(nil), "shop.PutMyShopResponse")
	proto.RegisterType((*DeleteMyShopRequest)(nil), "shop.DeleteMyShopRequest")
	proto.RegisterType((*DeleteMyShopResponse)(nil), "shop.DeleteMyShopResponse")
}

func init() { proto.RegisterFile("shop.proto", fileDescriptor_0f3030369b20fd61) }

var fileDescriptor_0f3030369b20fd61 = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0xdd, 0xa6, 0xcd, 0x6a, 0xee, 0x16, 0x6d, 0xef, 0x2e, 0xbb, 0x31, 0x28, 0x96, 0x81, 0x95,
	0x82, 0xb6, 0xd5, 0xfa, 0xe2, 0x83, 0x2c, 0x54, 0x56, 0x8a, 0x0f, 0x62, 0x49, 0x05, 0xc1, 0x97,
	0x65, 0xb6, 0x19, 0x9b, 0xa1, 0x69, 0x26, 0x66, 0x26, 0xca, 0xfe, 0x19, 0xff, 0x82, 0xbf, 0xc5,
	0x7f, 0x24, 0x99, 0xa4, 0x69, 0xb6, 0xdf, 0x6f, 0xb9, 0xf7, 0x4e, 0xcf, 0x39, 0xf7, 0x9c, 0x4b,
	0x01, 0xa4, 0x2f, 0xa2, 0x6e, 0x14, 0x0b, 0x25, 0xb0, 0x96, 0x7e, 0x93, 0x3f, 0x15, 0x38, 0x71,
	0xd9, 0xcf, 0x84, 0x49, 0x35, 0xf6, 0x45, 0x84, 0xcf, 0x00, 0xf8, 0x9c, 0x4e, 0xd9, 0x4d, 0x44,
	0x95, 0x6f, 0x57, 0x5a, 0x95, 0xb6, 0xe5, 0x5a, 0xba, 0x33, 0xa2, 0xca, 0x47, 0x84, 0x5a, 0x48,
	0xe7, 0xcc, 0x36, 0xf4, 0x40, 0x7f, 0xa3, 0x03, 0x0f, 0x03, 0xaa, 0xb8, 0x4a, 0x3c, 0x66, 0x57,
	0x5b, 0x95, 0xb6, 0xe1, 0x16, 0x35, 0x3e, 0x05, 0x2b, 0x10, 0xe1, 0x34, 0x1b, 0xd6, 0xf4, 0x70,
	0xd9, 0x40, 0x02, 0x75, 0x1e, 0xaa, 0x58, 0x78, 0xc9, 0x44, 0x71, 0x11, 0xda, 0xa6, 0x46, 0xbd,
	0xd7, 0x23, 0x7f, 0x2b, 0x50, 0x77, 0x99, 0x8c, 0x44, 0x28, 0x99, 0x56, 0xf8, 0x08, 0x0c, 0xee,
	0x69, 0x65, 0x55, 0xd7, 0xe0, 0xde, 0x8a, 0x62, 0x63, 0x9b, 0xe2, 0xea, 0x16, 0xc5, 0xb5, 0x5d,
	0x8a, 0xcd, 0x7d, 0x8a, 0x8f, 0x37, 0x28, 0x6e, 0x43, 0x63, 0xc8, 0xd4, 0xe7, 0xbb, 0x54, 0x6d,
	0x6e, 0x2d, 0x9e, 0x81, 0xa9, 0xc4, 0x8c, 0x85, 0xb9, 0xa3, 0x59, 0x41, 0xc6, 0xd0, 0x2c, 0xbd,
	0xcc, 0x76, 0xc4, 0x73, 0x38, 0x96, 0x8a, 0xaa, 0x44, 0xe6, 0x3b, 0xe6, 0x15, 0xbe, 0x00, 0x9d,
	0x98, 0xde, 0xf0, 0xa4, 0x8f, 0x5d, 0x1d, 0x65, 0xd9, 0x19, 0x37, 0x4b, 0xb4, 0x0e, 0x90, 0x56,
	0xf2, 0xe3, 0x3c, 0x52, 0x77, 0xe4, 0x1b, 0x9c, 0x0e, 0x99, 0x1a, 0x04, 0x81, 0xee, 0xed, 0x25,
	0x69, 0x83, 0x99, 0x82, 0x48, 0xdb, 0x68, 0x55, 0xb7, 0xb0, 0x64, 0x0f, 0xc8, 0x73, 0xb0, 0xd2,
	0xf2, 0x53, 0x6a, 0x74, 0x6a, 0xb2, 0x47, 0x15, 0xd5, 0x60, 0x75, 0x57, 0x7f, 0x93, 0xdf, 0xd0,
	0x1c, 0x09, 0x79, 0x88, 0x0f, 0x78, 0x09, 0xa6, 0x0e, 0x2c, 0xdf, 0xed, 0x71, 0xc6, 0x5a, 0xc0,
	0xbb, 0xd9, 0x14, 0x2f, 0x73, 0x07, 0xaa, 0xfa, 0x55, 0x73, 0xa1, 0xad, 0x38, 0xde, 0xdc, 0x80,
	0x57, 0x80, 0x65, 0xe2, 0xdd, 0x1b, 0x93, 0x2f, 0xd0, 0x18, 0x25, 0x07, 0xaa, 0x2c, 0x07, 0xb0,
	0x95, 0xfe, 0x25, 0x34, 0x4b, 0x80, 0x7b, 0xd8, 0xaf, 0xe1, 0xf4, 0x9a, 0x05, 0x4c, 0xb1, 0x43,
	0x04, 0x5c, 0xc0, 0x83, 0x94, 0xe1, 0x86, 0x7b, 0x5a, 0x43, 0x8a, 0x92, 0x5a, 0xe4, 0x91, 0x2e,
	0x9c, 0xdd, 0x47, 0xd9, 0xcd, 0xda, 0xff, 0x67, 0x80, 0xa9, 0xef, 0x01, 0xaf, 0xc0, 0x2a, 0x2e,
	0x10, 0xcf, 0xb3, 0x95, 0x56, 0x8f, 0xd7, 0xb9, 0x58, 0xeb, 0x67, 0xf8, 0xe4, 0x08, 0xdf, 0xc3,
	0x49, 0xe9, 0xbc, 0xb0, 0xb1, 0x4c, 0x2e, 0xbb, 0x3f, 0xe7, 0x49, 0xf1, 0xdb, 0xd5, 0x1b, 0x24,
	0x47, 0x38, 0x00, 0x58, 0x26, 0x85, 0x39, 0xcd, 0xda, 0xd1, 0x38, 0xf6, 0xfa, 0xa0, 0x80, 0xb8,
	0x02, 0xab, 0x70, 0x7b, 0xb1, 0xc0, 0x6a, 0x9e, 0x8b, 0x05, 0xd6, 0x62, 0x21, 0x47, 0x38, 0x84,
	0x7a, 0xd9, 0x3a, 0xcc, 0xf5, 0x6e, 0x08, 0xc5, 0x71, 0x36, 0x8d, 0x16, 0x40, 0x1f, 0xbe, 0x7e,
	0x77, 0xa7, 0x5c, 0xf9, 0xc9, 0x6d, 0x77, 0x22, 0xe6, 0xbd, 0xfe, 0xeb, 0x37, 0xef, 0x3a, 0x73,
	0x1e, 0x53, 0x3e, 0x63, 0x5c, 0x51, 0xde, 0x11, 0xf1, 0xb4, 0xe7, 0xd2, 0x59, 0x22, 0x69, 0xc0,
	0x3a, 0x83, 0x50, 0x28, 0x9f, 0xc5, 0x9d, 0x31, 0x8b, 0x7f, 0xb1, 0xb8, 0xc7, 0x43, 0xc5, 0xe2,
	0x1f, 0x74, 0xc2, 0x64, 0x4f, 0x66, 0x9d, 0x38, 0x9a, 0xf4, 0x52, 0xb2, 0xdb, 0x63, 0xfd, 0x5f,
	0xfd, 0xf6, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x35, 0xdf, 0x21, 0xc4, 0xb9, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShopsClient is the client API for Shops service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShopsClient interface {
	GetMyShop(ctx context.Context, in *GetMyShopRequest, opts ...grpc.CallOption) (*GetMyShopResponse, error)
	GetAllShops(ctx context.Context, in *ShopsEmpty, opts ...grpc.CallOption) (*GetAllShopsResponse, error)
	PostMyShop(ctx context.Context, in *PostMyShopRequest, opts ...grpc.CallOption) (*PostMyShopResponse, error)
	PutMyShop(ctx context.Context, in *PutMyShopRequest, opts ...grpc.CallOption) (*PutMyShopResponse, error)
	DeleteMyShop(ctx context.Context, in *DeleteMyShopRequest, opts ...grpc.CallOption) (*DeleteMyShopResponse, error)
}

type shopsClient struct {
	cc *grpc.ClientConn
}

func NewShopsClient(cc *grpc.ClientConn) ShopsClient {
	return &shopsClient{cc}
}

func (c *shopsClient) GetMyShop(ctx context.Context, in *GetMyShopRequest, opts ...grpc.CallOption) (*GetMyShopResponse, error) {
	out := new(GetMyShopResponse)
	err := c.cc.Invoke(ctx, "/shop.Shops/GetMyShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopsClient) GetAllShops(ctx context.Context, in *ShopsEmpty, opts ...grpc.CallOption) (*GetAllShopsResponse, error) {
	out := new(GetAllShopsResponse)
	err := c.cc.Invoke(ctx, "/shop.Shops/GetAllShops", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopsClient) PostMyShop(ctx context.Context, in *PostMyShopRequest, opts ...grpc.CallOption) (*PostMyShopResponse, error) {
	out := new(PostMyShopResponse)
	err := c.cc.Invoke(ctx, "/shop.Shops/PostMyShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopsClient) PutMyShop(ctx context.Context, in *PutMyShopRequest, opts ...grpc.CallOption) (*PutMyShopResponse, error) {
	out := new(PutMyShopResponse)
	err := c.cc.Invoke(ctx, "/shop.Shops/PutMyShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopsClient) DeleteMyShop(ctx context.Context, in *DeleteMyShopRequest, opts ...grpc.CallOption) (*DeleteMyShopResponse, error) {
	out := new(DeleteMyShopResponse)
	err := c.cc.Invoke(ctx, "/shop.Shops/DeleteMyShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopsServer is the server API for Shops service.
type ShopsServer interface {
	GetMyShop(context.Context, *GetMyShopRequest) (*GetMyShopResponse, error)
	GetAllShops(context.Context, *ShopsEmpty) (*GetAllShopsResponse, error)
	PostMyShop(context.Context, *PostMyShopRequest) (*PostMyShopResponse, error)
	PutMyShop(context.Context, *PutMyShopRequest) (*PutMyShopResponse, error)
	DeleteMyShop(context.Context, *DeleteMyShopRequest) (*DeleteMyShopResponse, error)
}

func RegisterShopsServer(s *grpc.Server, srv ShopsServer) {
	s.RegisterService(&_Shops_serviceDesc, srv)
}

func _Shops_GetMyShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopsServer).GetMyShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.Shops/GetMyShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopsServer).GetMyShop(ctx, req.(*GetMyShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shops_GetAllShops_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopsEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopsServer).GetAllShops(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.Shops/GetAllShops",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopsServer).GetAllShops(ctx, req.(*ShopsEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shops_PostMyShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostMyShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopsServer).PostMyShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.Shops/PostMyShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopsServer).PostMyShop(ctx, req.(*PostMyShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shops_PutMyShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutMyShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopsServer).PutMyShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.Shops/PutMyShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopsServer).PutMyShop(ctx, req.(*PutMyShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shops_DeleteMyShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMyShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopsServer).DeleteMyShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.Shops/DeleteMyShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopsServer).DeleteMyShop(ctx, req.(*DeleteMyShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Shops_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shop.Shops",
	HandlerType: (*ShopsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMyShop",
			Handler:    _Shops_GetMyShop_Handler,
		},
		{
			MethodName: "GetAllShops",
			Handler:    _Shops_GetAllShops_Handler,
		},
		{
			MethodName: "PostMyShop",
			Handler:    _Shops_PostMyShop_Handler,
		},
		{
			MethodName: "PutMyShop",
			Handler:    _Shops_PutMyShop_Handler,
		},
		{
			MethodName: "DeleteMyShop",
			Handler:    _Shops_DeleteMyShop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop.proto",
}
