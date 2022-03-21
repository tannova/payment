// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: stark/api/vision.proto

package stark

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	api "gitlab.com/mcuc/monorepo/backend/groot/pkg/api"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetPaymentDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPaymentDetailRequest) Reset() {
	*x = GetPaymentDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stark_api_vision_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentDetailRequest) ProtoMessage() {}

func (x *GetPaymentDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stark_api_vision_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentDetailRequest.ProtoReflect.Descriptor instead.
func (*GetPaymentDetailRequest) Descriptor() ([]byte, []int) {
	return file_stark_api_vision_proto_rawDescGZIP(), []int{0}
}

func (x *GetPaymentDetailRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetPaymentDetailReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payment       *Payment       `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
	PaymentDetail *PaymentDetail `protobuf:"bytes,2,opt,name=payment_detail,json=paymentDetail,proto3" json:"payment_detail,omitempty"`
	Revisions     []*Revision    `protobuf:"bytes,3,rep,name=revisions,proto3" json:"revisions,omitempty"`
}

func (x *GetPaymentDetailReply) Reset() {
	*x = GetPaymentDetailReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stark_api_vision_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentDetailReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentDetailReply) ProtoMessage() {}

func (x *GetPaymentDetailReply) ProtoReflect() protoreflect.Message {
	mi := &file_stark_api_vision_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentDetailReply.ProtoReflect.Descriptor instead.
func (*GetPaymentDetailReply) Descriptor() ([]byte, []int) {
	return file_stark_api_vision_proto_rawDescGZIP(), []int{1}
}

func (x *GetPaymentDetailReply) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

func (x *GetPaymentDetailReply) GetPaymentDetail() *PaymentDetail {
	if x != nil {
		return x.PaymentDetail
	}
	return nil
}

func (x *GetPaymentDetailReply) GetRevisions() []*Revision {
	if x != nil {
		return x.Revisions
	}
	return nil
}

type ListPaymentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page              uint32               `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size              uint32               `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	From              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To                *timestamp.Timestamp `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	PaymentIds        []int64              `protobuf:"varint,5,rep,packed,name=payment_ids,json=paymentIds,proto3" json:"payment_ids,omitempty"`
	PaymentTypes      []PaymentType        `protobuf:"varint,6,rep,packed,name=payment_types,json=paymentTypes,proto3,enum=mcuc.stark.PaymentType" json:"payment_types,omitempty"`
	Methods           []MethodType         `protobuf:"varint,7,rep,packed,name=methods,proto3,enum=mcuc.stark.MethodType" json:"methods,omitempty"`
	Statuses          []Status             `protobuf:"varint,8,rep,packed,name=statuses,proto3,enum=mcuc.stark.Status" json:"statuses,omitempty"`
	MerchantIds       []int64              `protobuf:"varint,9,rep,packed,name=merchant_ids,json=merchantIds,proto3" json:"merchant_ids,omitempty"`
	BankNames         []BankName           `protobuf:"varint,10,rep,packed,name=bank_names,json=bankNames,proto3,enum=mcuc.stark.pepper.BankName" json:"bank_names,omitempty"`
	EWalletNames      []EWalletName        `protobuf:"varint,11,rep,packed,name=e_wallet_names,json=eWalletNames,proto3,enum=mcuc.stark.tony.EWalletName" json:"e_wallet_names,omitempty"`
	TelcoNames        []api.TelcoName      `protobuf:"varint,12,rep,packed,name=telco_names,json=telcoNames,proto3,enum=mcuc.groot.TelcoName" json:"telco_names,omitempty"`
	CryptoWalletName  []CryptoWalletName   `protobuf:"varint,13,rep,packed,name=crypto_wallet_name,json=cryptoWalletName,proto3,enum=mcuc.stark.ultron.CryptoWalletName" json:"crypto_wallet_name,omitempty"`
	OrderAscCreatedAt bool                 `protobuf:"varint,14,opt,name=order_asc_created_at,json=orderAscCreatedAt,proto3" json:"order_asc_created_at,omitempty"`
	OrderAscUpdatedAt bool                 `protobuf:"varint,15,opt,name=order_asc_updated_at,json=orderAscUpdatedAt,proto3" json:"order_asc_updated_at,omitempty"`
}

func (x *ListPaymentsRequest) Reset() {
	*x = ListPaymentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stark_api_vision_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPaymentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPaymentsRequest) ProtoMessage() {}

func (x *ListPaymentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stark_api_vision_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPaymentsRequest.ProtoReflect.Descriptor instead.
func (*ListPaymentsRequest) Descriptor() ([]byte, []int) {
	return file_stark_api_vision_proto_rawDescGZIP(), []int{2}
}

func (x *ListPaymentsRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListPaymentsRequest) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListPaymentsRequest) GetFrom() *timestamp.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *ListPaymentsRequest) GetTo() *timestamp.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *ListPaymentsRequest) GetPaymentIds() []int64 {
	if x != nil {
		return x.PaymentIds
	}
	return nil
}

func (x *ListPaymentsRequest) GetPaymentTypes() []PaymentType {
	if x != nil {
		return x.PaymentTypes
	}
	return nil
}

func (x *ListPaymentsRequest) GetMethods() []MethodType {
	if x != nil {
		return x.Methods
	}
	return nil
}

func (x *ListPaymentsRequest) GetStatuses() []Status {
	if x != nil {
		return x.Statuses
	}
	return nil
}

func (x *ListPaymentsRequest) GetMerchantIds() []int64 {
	if x != nil {
		return x.MerchantIds
	}
	return nil
}

func (x *ListPaymentsRequest) GetBankNames() []BankName {
	if x != nil {
		return x.BankNames
	}
	return nil
}

func (x *ListPaymentsRequest) GetEWalletNames() []EWalletName {
	if x != nil {
		return x.EWalletNames
	}
	return nil
}

func (x *ListPaymentsRequest) GetTelcoNames() []api.TelcoName {
	if x != nil {
		return x.TelcoNames
	}
	return nil
}

func (x *ListPaymentsRequest) GetCryptoWalletName() []CryptoWalletName {
	if x != nil {
		return x.CryptoWalletName
	}
	return nil
}

func (x *ListPaymentsRequest) GetOrderAscCreatedAt() bool {
	if x != nil {
		return x.OrderAscCreatedAt
	}
	return false
}

func (x *ListPaymentsRequest) GetOrderAscUpdatedAt() bool {
	if x != nil {
		return x.OrderAscUpdatedAt
	}
	return false
}

type ListPaymentsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records     []*PaymentWithDetail `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	Total       uint64               `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	CurrentPage uint32               `protobuf:"varint,3,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
}

func (x *ListPaymentsReply) Reset() {
	*x = ListPaymentsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stark_api_vision_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPaymentsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPaymentsReply) ProtoMessage() {}

func (x *ListPaymentsReply) ProtoReflect() protoreflect.Message {
	mi := &file_stark_api_vision_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPaymentsReply.ProtoReflect.Descriptor instead.
func (*ListPaymentsReply) Descriptor() ([]byte, []int) {
	return file_stark_api_vision_proto_rawDescGZIP(), []int{3}
}

func (x *ListPaymentsReply) GetRecords() []*PaymentWithDetail {
	if x != nil {
		return x.Records
	}
	return nil
}

func (x *ListPaymentsReply) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListPaymentsReply) GetCurrentPage() uint32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

type GetPaymentInfoByPaymentCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetPaymentInfoByPaymentCodeRequest) Reset() {
	*x = GetPaymentInfoByPaymentCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stark_api_vision_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentInfoByPaymentCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentInfoByPaymentCodeRequest) ProtoMessage() {}

func (x *GetPaymentInfoByPaymentCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stark_api_vision_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentInfoByPaymentCodeRequest.ProtoReflect.Descriptor instead.
func (*GetPaymentInfoByPaymentCodeRequest) Descriptor() ([]byte, []int) {
	return file_stark_api_vision_proto_rawDescGZIP(), []int{4}
}

func (x *GetPaymentInfoByPaymentCodeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GetPaymentInfoByPaymentCodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MerchantId     int64      `protobuf:"varint,1,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	MerchantUserId int64      `protobuf:"varint,2,opt,name=merchant_user_id,json=merchantUserId,proto3" json:"merchant_user_id,omitempty"`
	PaymentMethod  MethodType `protobuf:"varint,3,opt,name=payment_method,json=paymentMethod,proto3,enum=mcuc.stark.MethodType" json:"payment_method,omitempty"`
	// Types that are assignable to Payload:
	//	*GetPaymentInfoByPaymentCodeReply_BankName
	//	*GetPaymentInfoByPaymentCodeReply_EWalletName
	Payload isGetPaymentInfoByPaymentCodeReply_Payload `protobuf_oneof:"payload"`
}

func (x *GetPaymentInfoByPaymentCodeReply) Reset() {
	*x = GetPaymentInfoByPaymentCodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stark_api_vision_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentInfoByPaymentCodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentInfoByPaymentCodeReply) ProtoMessage() {}

func (x *GetPaymentInfoByPaymentCodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_stark_api_vision_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentInfoByPaymentCodeReply.ProtoReflect.Descriptor instead.
func (*GetPaymentInfoByPaymentCodeReply) Descriptor() ([]byte, []int) {
	return file_stark_api_vision_proto_rawDescGZIP(), []int{5}
}

func (x *GetPaymentInfoByPaymentCodeReply) GetMerchantId() int64 {
	if x != nil {
		return x.MerchantId
	}
	return 0
}

func (x *GetPaymentInfoByPaymentCodeReply) GetMerchantUserId() int64 {
	if x != nil {
		return x.MerchantUserId
	}
	return 0
}

func (x *GetPaymentInfoByPaymentCodeReply) GetPaymentMethod() MethodType {
	if x != nil {
		return x.PaymentMethod
	}
	return MethodType_METHOD_UNSPECIFIED
}

func (m *GetPaymentInfoByPaymentCodeReply) GetPayload() isGetPaymentInfoByPaymentCodeReply_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *GetPaymentInfoByPaymentCodeReply) GetBankName() BankName {
	if x, ok := x.GetPayload().(*GetPaymentInfoByPaymentCodeReply_BankName); ok {
		return x.BankName
	}
	return BankName_BANK_UNSPECIFIED
}

func (x *GetPaymentInfoByPaymentCodeReply) GetEWalletName() EWalletName {
	if x, ok := x.GetPayload().(*GetPaymentInfoByPaymentCodeReply_EWalletName); ok {
		return x.EWalletName
	}
	return EWalletName_EWALLET_NAME_UNSPECIFIED
}

type isGetPaymentInfoByPaymentCodeReply_Payload interface {
	isGetPaymentInfoByPaymentCodeReply_Payload()
}

type GetPaymentInfoByPaymentCodeReply_BankName struct {
	BankName BankName `protobuf:"varint,4,opt,name=bank_name,json=bankName,proto3,enum=mcuc.stark.pepper.BankName,oneof"`
}

type GetPaymentInfoByPaymentCodeReply_EWalletName struct {
	EWalletName EWalletName `protobuf:"varint,5,opt,name=e_wallet_name,json=eWalletName,proto3,enum=mcuc.stark.tony.EWalletName,oneof"`
}

func (*GetPaymentInfoByPaymentCodeReply_BankName) isGetPaymentInfoByPaymentCodeReply_Payload() {}

func (*GetPaymentInfoByPaymentCodeReply_EWalletName) isGetPaymentInfoByPaymentCodeReply_Payload() {}

var File_stark_api_vision_proto protoreflect.FileDescriptor

var file_stark_api_vision_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73,
	0x74, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x72, 0x6f, 0x6f, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x72, 0x6f, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x73, 0x74,
	0x61, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x65, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x74, 0x61,
	0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x6c, 0x74,
	0x72, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0xbc, 0x01,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e,
	0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x32, 0x0a, 0x09, 0x72, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x63,
	0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xea, 0x05, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x2e, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x2a, 0x0a, 0x02,
	0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x12, 0x3c, 0x0a, 0x0d, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e,
	0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6d, 0x63,
	0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x12, 0x3a, 0x0a, 0x0a,
	0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x1b, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x70, 0x65,
	0x70, 0x70, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x09, 0x62,
	0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x0e, 0x65, 0x5f, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x1c, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x74, 0x6f,
	0x6e, 0x79, 0x2e, 0x45, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x0c,
	0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x0b,
	0x74, 0x65, 0x6c, 0x63, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x67, 0x72, 0x6f, 0x6f, 0x74, 0x2e, 0x54,
	0x65, 0x6c, 0x63, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x0a, 0x74, 0x65, 0x6c, 0x63, 0x6f, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x12, 0x51, 0x0a, 0x12, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x23, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x75, 0x6c,
	0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x10, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x14, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x61, 0x73, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x73, 0x63, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2f, 0x0a, 0x14, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x61, 0x73, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x73, 0x63,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x37, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x22, 0x38, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x79, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xb7, 0x02, 0x0a, 0x20,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b,
	0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x3a, 0x0a, 0x09, 0x62, 0x61,
	0x6e, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e,
	0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x70, 0x65, 0x70, 0x70, 0x65,
	0x72, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x08, 0x62, 0x61,
	0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x65, 0x5f, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e,
	0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x74, 0x6f, 0x6e, 0x79, 0x2e,
	0x45, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x65,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0xdc, 0x02, 0x0a, 0x06, 0x56, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x68, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x2a, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72,
	0x6b, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5c, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x2e, 0x6d, 0x63, 0x75,
	0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x89, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x35, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e,
	0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x33, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x42, 0x79, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x63, 0x75, 0x63, 0x2f, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stark_api_vision_proto_rawDescOnce sync.Once
	file_stark_api_vision_proto_rawDescData = file_stark_api_vision_proto_rawDesc
)

func file_stark_api_vision_proto_rawDescGZIP() []byte {
	file_stark_api_vision_proto_rawDescOnce.Do(func() {
		file_stark_api_vision_proto_rawDescData = protoimpl.X.CompressGZIP(file_stark_api_vision_proto_rawDescData)
	})
	return file_stark_api_vision_proto_rawDescData
}

var file_stark_api_vision_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_stark_api_vision_proto_goTypes = []interface{}{
	(*GetPaymentDetailRequest)(nil),            // 0: mcuc.stark.vision.GetPaymentDetailRequest
	(*GetPaymentDetailReply)(nil),              // 1: mcuc.stark.vision.GetPaymentDetailReply
	(*ListPaymentsRequest)(nil),                // 2: mcuc.stark.vision.ListPaymentsRequest
	(*ListPaymentsReply)(nil),                  // 3: mcuc.stark.vision.ListPaymentsReply
	(*GetPaymentInfoByPaymentCodeRequest)(nil), // 4: mcuc.stark.vision.GetPaymentInfoByPaymentCodeRequest
	(*GetPaymentInfoByPaymentCodeReply)(nil),   // 5: mcuc.stark.vision.GetPaymentInfoByPaymentCodeReply
	(*Payment)(nil),                            // 6: mcuc.stark.Payment
	(*PaymentDetail)(nil),                      // 7: mcuc.stark.PaymentDetail
	(*Revision)(nil),                           // 8: mcuc.stark.Revision
	(*timestamp.Timestamp)(nil),                // 9: google.protobuf.Timestamp
	(PaymentType)(0),                           // 10: mcuc.stark.PaymentType
	(MethodType)(0),                            // 11: mcuc.stark.MethodType
	(Status)(0),                                // 12: mcuc.stark.Status
	(BankName)(0),                              // 13: mcuc.stark.pepper.BankName
	(EWalletName)(0),                           // 14: mcuc.stark.tony.EWalletName
	(api.TelcoName)(0),                         // 15: mcuc.groot.TelcoName
	(CryptoWalletName)(0),                      // 16: mcuc.stark.ultron.CryptoWalletName
	(*PaymentWithDetail)(nil),                  // 17: mcuc.stark.PaymentWithDetail
}
var file_stark_api_vision_proto_depIdxs = []int32{
	6,  // 0: mcuc.stark.vision.GetPaymentDetailReply.payment:type_name -> mcuc.stark.Payment
	7,  // 1: mcuc.stark.vision.GetPaymentDetailReply.payment_detail:type_name -> mcuc.stark.PaymentDetail
	8,  // 2: mcuc.stark.vision.GetPaymentDetailReply.revisions:type_name -> mcuc.stark.Revision
	9,  // 3: mcuc.stark.vision.ListPaymentsRequest.from:type_name -> google.protobuf.Timestamp
	9,  // 4: mcuc.stark.vision.ListPaymentsRequest.to:type_name -> google.protobuf.Timestamp
	10, // 5: mcuc.stark.vision.ListPaymentsRequest.payment_types:type_name -> mcuc.stark.PaymentType
	11, // 6: mcuc.stark.vision.ListPaymentsRequest.methods:type_name -> mcuc.stark.MethodType
	12, // 7: mcuc.stark.vision.ListPaymentsRequest.statuses:type_name -> mcuc.stark.Status
	13, // 8: mcuc.stark.vision.ListPaymentsRequest.bank_names:type_name -> mcuc.stark.pepper.BankName
	14, // 9: mcuc.stark.vision.ListPaymentsRequest.e_wallet_names:type_name -> mcuc.stark.tony.EWalletName
	15, // 10: mcuc.stark.vision.ListPaymentsRequest.telco_names:type_name -> mcuc.groot.TelcoName
	16, // 11: mcuc.stark.vision.ListPaymentsRequest.crypto_wallet_name:type_name -> mcuc.stark.ultron.CryptoWalletName
	17, // 12: mcuc.stark.vision.ListPaymentsReply.records:type_name -> mcuc.stark.PaymentWithDetail
	11, // 13: mcuc.stark.vision.GetPaymentInfoByPaymentCodeReply.payment_method:type_name -> mcuc.stark.MethodType
	13, // 14: mcuc.stark.vision.GetPaymentInfoByPaymentCodeReply.bank_name:type_name -> mcuc.stark.pepper.BankName
	14, // 15: mcuc.stark.vision.GetPaymentInfoByPaymentCodeReply.e_wallet_name:type_name -> mcuc.stark.tony.EWalletName
	0,  // 16: mcuc.stark.vision.Vision.GetPaymentDetail:input_type -> mcuc.stark.vision.GetPaymentDetailRequest
	2,  // 17: mcuc.stark.vision.Vision.ListPayments:input_type -> mcuc.stark.vision.ListPaymentsRequest
	4,  // 18: mcuc.stark.vision.Vision.GetPaymentInfoByPaymentCode:input_type -> mcuc.stark.vision.GetPaymentInfoByPaymentCodeRequest
	1,  // 19: mcuc.stark.vision.Vision.GetPaymentDetail:output_type -> mcuc.stark.vision.GetPaymentDetailReply
	3,  // 20: mcuc.stark.vision.Vision.ListPayments:output_type -> mcuc.stark.vision.ListPaymentsReply
	5,  // 21: mcuc.stark.vision.Vision.GetPaymentInfoByPaymentCode:output_type -> mcuc.stark.vision.GetPaymentInfoByPaymentCodeReply
	19, // [19:22] is the sub-list for method output_type
	16, // [16:19] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_stark_api_vision_proto_init() }
func file_stark_api_vision_proto_init() {
	if File_stark_api_vision_proto != nil {
		return
	}
	file_stark_api_stark_proto_init()
	file_stark_api_pepper_proto_init()
	file_stark_api_tony_proto_init()
	file_stark_api_ultron_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_stark_api_vision_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaymentDetailRequest); i {
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
		file_stark_api_vision_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaymentDetailReply); i {
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
		file_stark_api_vision_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPaymentsRequest); i {
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
		file_stark_api_vision_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPaymentsReply); i {
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
		file_stark_api_vision_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaymentInfoByPaymentCodeRequest); i {
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
		file_stark_api_vision_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaymentInfoByPaymentCodeReply); i {
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
	file_stark_api_vision_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*GetPaymentInfoByPaymentCodeReply_BankName)(nil),
		(*GetPaymentInfoByPaymentCodeReply_EWalletName)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stark_api_vision_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stark_api_vision_proto_goTypes,
		DependencyIndexes: file_stark_api_vision_proto_depIdxs,
		MessageInfos:      file_stark_api_vision_proto_msgTypes,
	}.Build()
	File_stark_api_vision_proto = out.File
	file_stark_api_vision_proto_rawDesc = nil
	file_stark_api_vision_proto_goTypes = nil
	file_stark_api_vision_proto_depIdxs = nil
}
