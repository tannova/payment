// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: stark/api/stark.proto

package stark

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type PaymentType int32

const (
	PaymentType_PAYMENT_UNSPECIFIED PaymentType = 0
	PaymentType_TOPUP               PaymentType = 1
	PaymentType_WITHDRAW            PaymentType = 2
)

// Enum value maps for PaymentType.
var (
	PaymentType_name = map[int32]string{
		0: "PAYMENT_UNSPECIFIED",
		1: "TOPUP",
		2: "WITHDRAW",
	}
	PaymentType_value = map[string]int32{
		"PAYMENT_UNSPECIFIED": 0,
		"TOPUP":               1,
		"WITHDRAW":            2,
	}
)

func (x PaymentType) Enum() *PaymentType {
	p := new(PaymentType)
	*p = x
	return p
}

func (x PaymentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentType) Descriptor() protoreflect.EnumDescriptor {
	return file_stark_api_stark_proto_enumTypes[0].Descriptor()
}

func (PaymentType) Type() protoreflect.EnumType {
	return &file_stark_api_stark_proto_enumTypes[0]
}

func (x PaymentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentType.Descriptor instead.
func (PaymentType) EnumDescriptor() ([]byte, []int) {
	return file_stark_api_stark_proto_rawDescGZIP(), []int{0}
}

type MethodType int32

const (
	MethodType_METHOD_UNSPECIFIED MethodType = 0
	MethodType_T                  MethodType = 1 // Bank Transfer
	MethodType_I                  MethodType = 2 // Internet Banking
	MethodType_P                  MethodType = 3 // Phone card/Telco
	MethodType_E                  MethodType = 4 // E Wallet
	MethodType_C                  MethodType = 5 // Cryptocurrency
)

// Enum value maps for MethodType.
var (
	MethodType_name = map[int32]string{
		0: "METHOD_UNSPECIFIED",
		1: "T",
		2: "I",
		3: "P",
		4: "E",
		5: "C",
	}
	MethodType_value = map[string]int32{
		"METHOD_UNSPECIFIED": 0,
		"T":                  1,
		"I":                  2,
		"P":                  3,
		"E":                  4,
		"C":                  5,
	}
)

func (x MethodType) Enum() *MethodType {
	p := new(MethodType)
	*p = x
	return p
}

func (x MethodType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MethodType) Descriptor() protoreflect.EnumDescriptor {
	return file_stark_api_stark_proto_enumTypes[1].Descriptor()
}

func (MethodType) Type() protoreflect.EnumType {
	return &file_stark_api_stark_proto_enumTypes[1]
}

func (x MethodType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MethodType.Descriptor instead.
func (MethodType) EnumDescriptor() ([]byte, []int) {
	return file_stark_api_stark_proto_rawDescGZIP(), []int{1}
}

type Status int32

const (
	Status_STATUS_UNSPECIFIED Status = 0
	Status_CREATED            Status = 1
	Status_CANCELED           Status = 2
	Status_REJECTING          Status = 3
	Status_REJECTED           Status = 4
	Status_REJECT_FAILED      Status = 5
	Status_APPROVED           Status = 6
	Status_APPROVE_FAILED     Status = 7
	Status_SUBMITTED          Status = 8
	Status_SUBMIT_FAILED      Status = 9
	Status_COMPLETED          Status = 10
	// for calling to UMO and prevent teller submit manual
	Status_SUBMITTING Status = 11
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0:  "STATUS_UNSPECIFIED",
		1:  "CREATED",
		2:  "CANCELED",
		3:  "REJECTING",
		4:  "REJECTED",
		5:  "REJECT_FAILED",
		6:  "APPROVED",
		7:  "APPROVE_FAILED",
		8:  "SUBMITTED",
		9:  "SUBMIT_FAILED",
		10: "COMPLETED",
		11: "SUBMITTING",
	}
	Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"CREATED":            1,
		"CANCELED":           2,
		"REJECTING":          3,
		"REJECTED":           4,
		"REJECT_FAILED":      5,
		"APPROVED":           6,
		"APPROVE_FAILED":     7,
		"SUBMITTED":          8,
		"SUBMIT_FAILED":      9,
		"COMPLETED":          10,
		"SUBMITTING":         11,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_stark_api_stark_proto_enumTypes[2].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_stark_api_stark_proto_enumTypes[2]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_stark_api_stark_proto_rawDescGZIP(), []int{2}
}

type Currency int32

const (
	Currency_CURRENCY_UNSPECIFIED Currency = 0
	Currency_VND                  Currency = 1
	Currency_USDT                 Currency = 2
)

// Enum value maps for Currency.
var (
	Currency_name = map[int32]string{
		0: "CURRENCY_UNSPECIFIED",
		1: "VND",
		2: "USDT",
	}
	Currency_value = map[string]int32{
		"CURRENCY_UNSPECIFIED": 0,
		"VND":                  1,
		"USDT":                 2,
	}
)

func (x Currency) Enum() *Currency {
	p := new(Currency)
	*p = x
	return p
}

func (x Currency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Currency) Descriptor() protoreflect.EnumDescriptor {
	return file_stark_api_stark_proto_enumTypes[3].Descriptor()
}

func (Currency) Type() protoreflect.EnumType {
	return &file_stark_api_stark_proto_enumTypes[3]
}

func (x Currency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Currency.Descriptor instead.
func (Currency) EnumDescriptor() ([]byte, []int) {
	return file_stark_api_stark_proto_rawDescGZIP(), []int{3}
}

type Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt  *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedBy  string               `protobuf:"bytes,5,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedBy  string               `protobuf:"bytes,6,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	ApprovedBy string               `protobuf:"bytes,7,opt,name=approved_by,json=approvedBy,proto3" json:"approved_by,omitempty"`
	MerchantId int64                `protobuf:"varint,8,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	Method     MethodType           `protobuf:"varint,9,opt,name=method,proto3,enum=mcuc.stark.MethodType" json:"method,omitempty"`
	Type       PaymentType          `protobuf:"varint,10,opt,name=type,proto3,enum=mcuc.stark.PaymentType" json:"type,omitempty"`
	Status     Status               `protobuf:"varint,11,opt,name=status,proto3,enum=mcuc.stark.Status" json:"status,omitempty"`
}

func (x *Payment) Reset() {
	*x = Payment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stark_api_stark_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_stark_api_stark_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_stark_api_stark_proto_rawDescGZIP(), []int{0}
}

func (x *Payment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Payment) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Payment) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Payment) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Payment) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *Payment) GetApprovedBy() string {
	if x != nil {
		return x.ApprovedBy
	}
	return ""
}

func (x *Payment) GetMerchantId() int64 {
	if x != nil {
		return x.MerchantId
	}
	return 0
}

func (x *Payment) GetMethod() MethodType {
	if x != nil {
		return x.Method
	}
	return MethodType_METHOD_UNSPECIFIED
}

func (x *Payment) GetType() PaymentType {
	if x != nil {
		return x.Type
	}
	return PaymentType_PAYMENT_UNSPECIFIED
}

func (x *Payment) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_UNSPECIFIED
}

type PaymentDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//	*PaymentDetail_Banking
	//	*PaymentDetail_EWallet
	//	*PaymentDetail_Telco
	//	*PaymentDetail_Crypto
	Payload isPaymentDetail_Payload `protobuf_oneof:"payload"`
}

func (x *PaymentDetail) Reset() {
	*x = PaymentDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stark_api_stark_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentDetail) ProtoMessage() {}

func (x *PaymentDetail) ProtoReflect() protoreflect.Message {
	mi := &file_stark_api_stark_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentDetail.ProtoReflect.Descriptor instead.
func (*PaymentDetail) Descriptor() ([]byte, []int) {
	return file_stark_api_stark_proto_rawDescGZIP(), []int{1}
}

func (m *PaymentDetail) GetPayload() isPaymentDetail_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *PaymentDetail) GetBanking() *BankingPaymentDetail {
	if x, ok := x.GetPayload().(*PaymentDetail_Banking); ok {
		return x.Banking
	}
	return nil
}

func (x *PaymentDetail) GetEWallet() *EWalletPaymentDetail {
	if x, ok := x.GetPayload().(*PaymentDetail_EWallet); ok {
		return x.EWallet
	}
	return nil
}

func (x *PaymentDetail) GetTelco() *TelcoPaymentDetail {
	if x, ok := x.GetPayload().(*PaymentDetail_Telco); ok {
		return x.Telco
	}
	return nil
}

func (x *PaymentDetail) GetCrypto() *CryptoPaymentDetail {
	if x, ok := x.GetPayload().(*PaymentDetail_Crypto); ok {
		return x.Crypto
	}
	return nil
}

type isPaymentDetail_Payload interface {
	isPaymentDetail_Payload()
}

type PaymentDetail_Banking struct {
	Banking *BankingPaymentDetail `protobuf:"bytes,1,opt,name=banking,proto3,oneof"`
}

type PaymentDetail_EWallet struct {
	EWallet *EWalletPaymentDetail `protobuf:"bytes,2,opt,name=e_wallet,json=eWallet,proto3,oneof"`
}

type PaymentDetail_Telco struct {
	Telco *TelcoPaymentDetail `protobuf:"bytes,3,opt,name=telco,proto3,oneof"`
}

type PaymentDetail_Crypto struct {
	Crypto *CryptoPaymentDetail `protobuf:"bytes,4,opt,name=crypto,proto3,oneof"`
}

func (*PaymentDetail_Banking) isPaymentDetail_Payload() {}

func (*PaymentDetail_EWallet) isPaymentDetail_Payload() {}

func (*PaymentDetail_Telco) isPaymentDetail_Payload() {}

func (*PaymentDetail_Crypto) isPaymentDetail_Payload() {}

type PaymentWithDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payment       *Payment       `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
	PaymentDetail *PaymentDetail `protobuf:"bytes,2,opt,name=payment_detail,json=paymentDetail,proto3" json:"payment_detail,omitempty"`
}

func (x *PaymentWithDetail) Reset() {
	*x = PaymentWithDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stark_api_stark_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentWithDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentWithDetail) ProtoMessage() {}

func (x *PaymentWithDetail) ProtoReflect() protoreflect.Message {
	mi := &file_stark_api_stark_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentWithDetail.ProtoReflect.Descriptor instead.
func (*PaymentWithDetail) Descriptor() ([]byte, []int) {
	return file_stark_api_stark_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentWithDetail) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

func (x *PaymentWithDetail) GetPaymentDetail() *PaymentDetail {
	if x != nil {
		return x.PaymentDetail
	}
	return nil
}

type Revision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy string               `protobuf:"bytes,3,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	PaymentId int64                `protobuf:"varint,4,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	Status    Status               `protobuf:"varint,5,opt,name=status,proto3,enum=mcuc.stark.Status" json:"status,omitempty"`
	Note      string               `protobuf:"bytes,6,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *Revision) Reset() {
	*x = Revision{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stark_api_stark_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Revision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Revision) ProtoMessage() {}

func (x *Revision) ProtoReflect() protoreflect.Message {
	mi := &file_stark_api_stark_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Revision.ProtoReflect.Descriptor instead.
func (*Revision) Descriptor() ([]byte, []int) {
	return file_stark_api_stark_proto_rawDescGZIP(), []int{3}
}

func (x *Revision) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Revision) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Revision) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Revision) GetPaymentId() int64 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

func (x *Revision) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_UNSPECIFIED
}

func (x *Revision) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

// Callback to Jarvis
type CompletePaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentId         int64          `protobuf:"varint,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	PaymentMethod     MethodType     `protobuf:"varint,2,opt,name=payment_method,json=paymentMethod,proto3,enum=mcuc.stark.MethodType" json:"payment_method,omitempty"`
	PaymentType       PaymentType    `protobuf:"varint,3,opt,name=payment_type,json=paymentType,proto3,enum=mcuc.stark.PaymentType" json:"payment_type,omitempty"`
	PaymentStatus     Status         `protobuf:"varint,4,opt,name=payment_status,json=paymentStatus,proto3,enum=mcuc.stark.Status" json:"payment_status,omitempty"`
	PaymentDetail     *PaymentDetail `protobuf:"bytes,5,opt,name=payment_detail,json=paymentDetail,proto3" json:"payment_detail,omitempty"`
	MexCurrentBalance uint64         `protobuf:"varint,6,opt,name=mex_current_balance,json=mexCurrentBalance,proto3" json:"mex_current_balance,omitempty"`
}

func (x *CompletePaymentRequest) Reset() {
	*x = CompletePaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stark_api_stark_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletePaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletePaymentRequest) ProtoMessage() {}

func (x *CompletePaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stark_api_stark_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletePaymentRequest.ProtoReflect.Descriptor instead.
func (*CompletePaymentRequest) Descriptor() ([]byte, []int) {
	return file_stark_api_stark_proto_rawDescGZIP(), []int{4}
}

func (x *CompletePaymentRequest) GetPaymentId() int64 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

func (x *CompletePaymentRequest) GetPaymentMethod() MethodType {
	if x != nil {
		return x.PaymentMethod
	}
	return MethodType_METHOD_UNSPECIFIED
}

func (x *CompletePaymentRequest) GetPaymentType() PaymentType {
	if x != nil {
		return x.PaymentType
	}
	return PaymentType_PAYMENT_UNSPECIFIED
}

func (x *CompletePaymentRequest) GetPaymentStatus() Status {
	if x != nil {
		return x.PaymentStatus
	}
	return Status_STATUS_UNSPECIFIED
}

func (x *CompletePaymentRequest) GetPaymentDetail() *PaymentDetail {
	if x != nil {
		return x.PaymentDetail
	}
	return nil
}

func (x *CompletePaymentRequest) GetMexCurrentBalance() uint64 {
	if x != nil {
		return x.MexCurrentBalance
	}
	return 0
}

// CompletePaymentReply
type CompletePaymentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CompletePaymentReply) Reset() {
	*x = CompletePaymentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stark_api_stark_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletePaymentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletePaymentReply) ProtoMessage() {}

func (x *CompletePaymentReply) ProtoReflect() protoreflect.Message {
	mi := &file_stark_api_stark_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletePaymentReply.ProtoReflect.Descriptor instead.
func (*CompletePaymentReply) Descriptor() ([]byte, []int) {
	return file_stark_api_stark_proto_rawDescGZIP(), []int{5}
}

var File_stark_api_stark_proto protoreflect.FileDescriptor

var file_stark_api_stark_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x61, 0x72,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74,
	0x61, 0x72, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x73, 0x74,
	0x61, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x65, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x74, 0x6f, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x73, 0x74, 0x61, 0x72,
	0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x6c, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x98, 0x03, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72,
	0x6b, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x12, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa4, 0x02,
	0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x43, 0x0a, 0x07, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x70, 0x65,
	0x70, 0x70, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x48, 0x00, 0x52, 0x07, 0x62, 0x61, 0x6e,
	0x6b, 0x69, 0x6e, 0x67, 0x12, 0x42, 0x0a, 0x08, 0x65, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74,
	0x61, 0x72, 0x6b, 0x2e, 0x74, 0x6f, 0x6e, 0x79, 0x2e, 0x45, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x48, 0x00, 0x52,
	0x07, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x3d, 0x0a, 0x05, 0x74, 0x65, 0x6c, 0x63,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73,
	0x74, 0x61, 0x72, 0x6b, 0x2e, 0x6d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x2e, 0x54, 0x65, 0x6c, 0x63,
	0x6f, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x48, 0x00,
	0x52, 0x05, 0x74, 0x65, 0x6c, 0x63, 0x6f, 0x12, 0x40, 0x0a, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73,
	0x74, 0x61, 0x72, 0x6b, 0x2e, 0x75, 0x6c, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x48,
	0x00, 0x52, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x63,
	0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x0e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0d, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0xd3, 0x01, 0x0a, 0x08,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x12, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x22, 0xdf, 0x02, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b,
	0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x3a, 0x0a, 0x0c, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12,
	0x2e, 0x6d, 0x63, 0x75, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x40, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x63, 0x75, 0x63,
	0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x65, 0x78, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x11, 0x6d, 0x65, 0x78, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2a, 0x3f, 0x0a, 0x0b, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x41,
	0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x4f, 0x50, 0x55, 0x50, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52, 0x41, 0x57, 0x10, 0x02, 0x2a, 0x47, 0x0a, 0x0a,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x45,
	0x54, 0x48, 0x4f, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x54, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01, 0x49, 0x10, 0x02,
	0x12, 0x05, 0x0a, 0x01, 0x50, 0x10, 0x03, 0x12, 0x05, 0x0a, 0x01, 0x45, 0x10, 0x04, 0x12, 0x05,
	0x0a, 0x01, 0x43, 0x10, 0x05, 0x2a, 0xce, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x49, 0x4e, 0x47,
	0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x04,
	0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x10,
	0x06, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x42, 0x4d, 0x49, 0x54, 0x54,
	0x45, 0x44, 0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x55, 0x42, 0x4d, 0x49, 0x54, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c,
	0x45, 0x54, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x55, 0x42, 0x4d, 0x49, 0x54,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x0b, 0x2a, 0x37, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x56, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x44, 0x54, 0x10, 0x02, 0x42,
	0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x63,
	0x75, 0x63, 0x2f, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x3b, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stark_api_stark_proto_rawDescOnce sync.Once
	file_stark_api_stark_proto_rawDescData = file_stark_api_stark_proto_rawDesc
)

func file_stark_api_stark_proto_rawDescGZIP() []byte {
	file_stark_api_stark_proto_rawDescOnce.Do(func() {
		file_stark_api_stark_proto_rawDescData = protoimpl.X.CompressGZIP(file_stark_api_stark_proto_rawDescData)
	})
	return file_stark_api_stark_proto_rawDescData
}

var file_stark_api_stark_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_stark_api_stark_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_stark_api_stark_proto_goTypes = []interface{}{
	(PaymentType)(0),               // 0: mcuc.stark.PaymentType
	(MethodType)(0),                // 1: mcuc.stark.MethodType
	(Status)(0),                    // 2: mcuc.stark.Status
	(Currency)(0),                  // 3: mcuc.stark.Currency
	(*Payment)(nil),                // 4: mcuc.stark.Payment
	(*PaymentDetail)(nil),          // 5: mcuc.stark.PaymentDetail
	(*PaymentWithDetail)(nil),      // 6: mcuc.stark.PaymentWithDetail
	(*Revision)(nil),               // 7: mcuc.stark.Revision
	(*CompletePaymentRequest)(nil), // 8: mcuc.stark.CompletePaymentRequest
	(*CompletePaymentReply)(nil),   // 9: mcuc.stark.CompletePaymentReply
	(*timestamp.Timestamp)(nil),    // 10: google.protobuf.Timestamp
	(*BankingPaymentDetail)(nil),   // 11: mcuc.stark.pepper.BankingPaymentDetail
	(*EWalletPaymentDetail)(nil),   // 12: mcuc.stark.tony.EWalletPaymentDetail
	(*TelcoPaymentDetail)(nil),     // 13: mcuc.stark.morgan.TelcoPaymentDetail
	(*CryptoPaymentDetail)(nil),    // 14: mcuc.stark.ultron.CryptoPaymentDetail
}
var file_stark_api_stark_proto_depIdxs = []int32{
	10, // 0: mcuc.stark.Payment.created_at:type_name -> google.protobuf.Timestamp
	10, // 1: mcuc.stark.Payment.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 2: mcuc.stark.Payment.method:type_name -> mcuc.stark.MethodType
	0,  // 3: mcuc.stark.Payment.type:type_name -> mcuc.stark.PaymentType
	2,  // 4: mcuc.stark.Payment.status:type_name -> mcuc.stark.Status
	11, // 5: mcuc.stark.PaymentDetail.banking:type_name -> mcuc.stark.pepper.BankingPaymentDetail
	12, // 6: mcuc.stark.PaymentDetail.e_wallet:type_name -> mcuc.stark.tony.EWalletPaymentDetail
	13, // 7: mcuc.stark.PaymentDetail.telco:type_name -> mcuc.stark.morgan.TelcoPaymentDetail
	14, // 8: mcuc.stark.PaymentDetail.crypto:type_name -> mcuc.stark.ultron.CryptoPaymentDetail
	4,  // 9: mcuc.stark.PaymentWithDetail.payment:type_name -> mcuc.stark.Payment
	5,  // 10: mcuc.stark.PaymentWithDetail.payment_detail:type_name -> mcuc.stark.PaymentDetail
	10, // 11: mcuc.stark.Revision.created_at:type_name -> google.protobuf.Timestamp
	2,  // 12: mcuc.stark.Revision.status:type_name -> mcuc.stark.Status
	1,  // 13: mcuc.stark.CompletePaymentRequest.payment_method:type_name -> mcuc.stark.MethodType
	0,  // 14: mcuc.stark.CompletePaymentRequest.payment_type:type_name -> mcuc.stark.PaymentType
	2,  // 15: mcuc.stark.CompletePaymentRequest.payment_status:type_name -> mcuc.stark.Status
	5,  // 16: mcuc.stark.CompletePaymentRequest.payment_detail:type_name -> mcuc.stark.PaymentDetail
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_stark_api_stark_proto_init() }
func file_stark_api_stark_proto_init() {
	if File_stark_api_stark_proto != nil {
		return
	}
	file_stark_api_morgan_proto_init()
	file_stark_api_pepper_proto_init()
	file_stark_api_tony_proto_init()
	file_stark_api_ultron_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_stark_api_stark_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payment); i {
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
		file_stark_api_stark_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentDetail); i {
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
		file_stark_api_stark_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentWithDetail); i {
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
		file_stark_api_stark_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Revision); i {
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
		file_stark_api_stark_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompletePaymentRequest); i {
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
		file_stark_api_stark_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompletePaymentReply); i {
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
	file_stark_api_stark_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*PaymentDetail_Banking)(nil),
		(*PaymentDetail_EWallet)(nil),
		(*PaymentDetail_Telco)(nil),
		(*PaymentDetail_Crypto)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stark_api_stark_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stark_api_stark_proto_goTypes,
		DependencyIndexes: file_stark_api_stark_proto_depIdxs,
		EnumInfos:         file_stark_api_stark_proto_enumTypes,
		MessageInfos:      file_stark_api_stark_proto_msgTypes,
	}.Build()
	File_stark_api_stark_proto = out.File
	file_stark_api_stark_proto_rawDesc = nil
	file_stark_api_stark_proto_goTypes = nil
	file_stark_api_stark_proto_depIdxs = nil
}