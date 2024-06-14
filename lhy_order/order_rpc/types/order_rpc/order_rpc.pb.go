// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: order_rpc.proto

package order_rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    []byte  `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Price   float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	OrderId int32   `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CreateOrderResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateOrderResponse) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type UpdateOrderStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	State  int64 `protobuf:"varint,2,opt,name=state,proto3" json:"state,omitempty"` //-1: 已取消 0:待支付 1:未使用
}

func (x *UpdateOrderStateRequest) Reset() {
	*x = UpdateOrderStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrderStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderStateRequest) ProtoMessage() {}

func (x *UpdateOrderStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderStateRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrderStateRequest) Descriptor() ([]byte, []int) {
	return file_order_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateOrderStateRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateOrderStateRequest) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

type UpdateOrderStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    int32   `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Price      float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	UserId     int32   `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	State      int64   `protobuf:"varint,4,opt,name=state,proto3" json:"state,omitempty"`
	CreateTime int64   `protobuf:"varint,5,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
}

func (x *UpdateOrderStateResponse) Reset() {
	*x = UpdateOrderStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrderStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderStateResponse) ProtoMessage() {}

func (x *UpdateOrderStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderStateResponse.ProtoReflect.Descriptor instead.
func (*UpdateOrderStateResponse) Descriptor() ([]byte, []int) {
	return file_order_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateOrderStateResponse) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *UpdateOrderStateResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateOrderStateResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateOrderStateResponse) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *UpdateOrderStateResponse) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

type GetOrderDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetOrderDetailRequest) Reset() {
	*x = GetOrderDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderDetailRequest) ProtoMessage() {}

func (x *GetOrderDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderDetailRequest.ProtoReflect.Descriptor instead.
func (*GetOrderDetailRequest) Descriptor() ([]byte, []int) {
	return file_order_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrderDetailRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetOrderDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    int32   `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Price      float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	UserId     int32   `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	State      int64   `protobuf:"varint,4,opt,name=state,proto3" json:"state,omitempty"`
	CreateTime int64   `protobuf:"varint,5,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
}

func (x *GetOrderDetailResponse) Reset() {
	*x = GetOrderDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderDetailResponse) ProtoMessage() {}

func (x *GetOrderDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderDetailResponse.ProtoReflect.Descriptor instead.
func (*GetOrderDetailResponse) Descriptor() ([]byte, []int) {
	return file_order_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *GetOrderDetailResponse) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *GetOrderDetailResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GetOrderDetailResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetOrderDetailResponse) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *GetOrderDetailResponse) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

type CreateSeckillOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ItemId int32 `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
}

func (x *CreateSeckillOrderRequest) Reset() {
	*x = CreateSeckillOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSeckillOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSeckillOrderRequest) ProtoMessage() {}

func (x *CreateSeckillOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSeckillOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateSeckillOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *CreateSeckillOrderRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateSeckillOrderRequest) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

type CreateSeckillOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    int64   `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Price      float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	UserId     int32   `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	State      int64   `protobuf:"varint,4,opt,name=state,proto3" json:"state,omitempty"`
	CreateTime int64   `protobuf:"varint,5,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
}

func (x *CreateSeckillOrderResponse) Reset() {
	*x = CreateSeckillOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_rpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSeckillOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSeckillOrderResponse) ProtoMessage() {}

func (x *CreateSeckillOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_rpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSeckillOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateSeckillOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_rpc_proto_rawDescGZIP(), []int{7}
}

func (x *CreateSeckillOrderResponse) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *CreateSeckillOrderResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateSeckillOrderResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateSeckillOrderResponse) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *CreateSeckillOrderResponse) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

var File_order_rpc_proto protoreflect.FileDescriptor

var file_order_rpc_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x22, 0x2d, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x9a, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x30,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x98, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x19, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x9c, 0x01, 0x0a, 0x1a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xcb, 0x03, 0x0a, 0x05, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x1d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5b, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x70,
	0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x20, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x24, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63,
	0x6b, 0x69, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_rpc_proto_rawDescOnce sync.Once
	file_order_rpc_proto_rawDescData = file_order_rpc_proto_rawDesc
)

func file_order_rpc_proto_rawDescGZIP() []byte {
	file_order_rpc_proto_rawDescOnce.Do(func() {
		file_order_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_rpc_proto_rawDescData)
	})
	return file_order_rpc_proto_rawDescData
}

var file_order_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_order_rpc_proto_goTypes = []interface{}{
	(*CreateOrderRequest)(nil),         // 0: order_rpc.CreateOrderRequest
	(*CreateOrderResponse)(nil),        // 1: order_rpc.CreateOrderResponse
	(*UpdateOrderStateRequest)(nil),    // 2: order_rpc.UpdateOrderStateRequest
	(*UpdateOrderStateResponse)(nil),   // 3: order_rpc.UpdateOrderStateResponse
	(*GetOrderDetailRequest)(nil),      // 4: order_rpc.GetOrderDetailRequest
	(*GetOrderDetailResponse)(nil),     // 5: order_rpc.GetOrderDetailResponse
	(*CreateSeckillOrderRequest)(nil),  // 6: order_rpc.CreateSeckillOrderRequest
	(*CreateSeckillOrderResponse)(nil), // 7: order_rpc.CreateSeckillOrderResponse
}
var file_order_rpc_proto_depIdxs = []int32{
	0, // 0: order_rpc.Order.CreateOrder:input_type -> order_rpc.CreateOrderRequest
	2, // 1: order_rpc.Order.UpdateOrderState:input_type -> order_rpc.UpdateOrderStateRequest
	4, // 2: order_rpc.Order.GetOrderDetail:input_type -> order_rpc.GetOrderDetailRequest
	6, // 3: order_rpc.Order.CreateSeckillOrder:input_type -> order_rpc.CreateSeckillOrderRequest
	6, // 4: order_rpc.Order.CreateRollback:input_type -> order_rpc.CreateSeckillOrderRequest
	1, // 5: order_rpc.Order.CreateOrder:output_type -> order_rpc.CreateOrderResponse
	3, // 6: order_rpc.Order.UpdateOrderState:output_type -> order_rpc.UpdateOrderStateResponse
	5, // 7: order_rpc.Order.GetOrderDetail:output_type -> order_rpc.GetOrderDetailResponse
	7, // 8: order_rpc.Order.CreateSeckillOrder:output_type -> order_rpc.CreateSeckillOrderResponse
	7, // 9: order_rpc.Order.CreateRollback:output_type -> order_rpc.CreateSeckillOrderResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_order_rpc_proto_init() }
func file_order_rpc_proto_init() {
	if File_order_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrderStateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrderStateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderDetailRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderDetailResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSeckillOrderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_rpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSeckillOrderResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_order_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_rpc_proto_goTypes,
		DependencyIndexes: file_order_rpc_proto_depIdxs,
		MessageInfos:      file_order_rpc_proto_msgTypes,
	}.Build()
	File_order_rpc_proto = out.File
	file_order_rpc_proto_rawDesc = nil
	file_order_rpc_proto_goTypes = nil
	file_order_rpc_proto_depIdxs = nil
}
