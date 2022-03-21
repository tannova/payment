// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: api/postman.proto

package postman

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type SendMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipients []string       `protobuf:"bytes,1,rep,name=recipients,proto3" json:"recipients,omitempty"`
	Subject    string         `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Content    string         `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Params     []*EmailParams `protobuf:"bytes,4,rep,name=params,proto3" json:"params,omitempty"`
}

func (x *SendMailRequest) Reset() {
	*x = SendMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_postman_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailRequest) ProtoMessage() {}

func (x *SendMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_postman_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailRequest.ProtoReflect.Descriptor instead.
func (*SendMailRequest) Descriptor() ([]byte, []int) {
	return file_api_postman_proto_rawDescGZIP(), []int{0}
}

func (x *SendMailRequest) GetRecipients() []string {
	if x != nil {
		return x.Recipients
	}
	return nil
}

func (x *SendMailRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SendMailRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SendMailRequest) GetParams() []*EmailParams {
	if x != nil {
		return x.Params
	}
	return nil
}

type EmailParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EmailParams) Reset() {
	*x = EmailParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_postman_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailParams) ProtoMessage() {}

func (x *EmailParams) ProtoReflect() protoreflect.Message {
	mi := &file_api_postman_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailParams.ProtoReflect.Descriptor instead.
func (*EmailParams) Descriptor() ([]byte, []int) {
	return file_api_postman_proto_rawDescGZIP(), []int{1}
}

func (x *EmailParams) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *EmailParams) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SendSmsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipient string `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Content   string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SendSmsRequest) Reset() {
	*x = SendSmsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_postman_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsRequest) ProtoMessage() {}

func (x *SendSmsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_postman_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsRequest.ProtoReflect.Descriptor instead.
func (*SendSmsRequest) Descriptor() ([]byte, []int) {
	return file_api_postman_proto_rawDescGZIP(), []int{2}
}

func (x *SendSmsRequest) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *SendSmsRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_api_postman_proto protoreflect.FileDescriptor

var file_api_postman_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x35, 0x0a, 0x0b, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x48, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x83, 0x01, 0x0a, 0x07,
	0x50, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x12, 0x3c, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x61, 0x69, 0x6c, 0x12, 0x18, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3a, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73,
	0x12, 0x17, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53,
	0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x72, 0x65, 0x79, 0x68, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_postman_proto_rawDescOnce sync.Once
	file_api_postman_proto_rawDescData = file_api_postman_proto_rawDesc
)

func file_api_postman_proto_rawDescGZIP() []byte {
	file_api_postman_proto_rawDescOnce.Do(func() {
		file_api_postman_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_postman_proto_rawDescData)
	})
	return file_api_postman_proto_rawDescData
}

var file_api_postman_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_postman_proto_goTypes = []interface{}{
	(*SendMailRequest)(nil), // 0: postman.SendMailRequest
	(*EmailParams)(nil),     // 1: postman.EmailParams
	(*SendSmsRequest)(nil),  // 2: postman.SendSmsRequest
	(*empty.Empty)(nil),     // 3: google.protobuf.Empty
}
var file_api_postman_proto_depIdxs = []int32{
	1, // 0: postman.SendMailRequest.params:type_name -> postman.EmailParams
	0, // 1: postman.Postman.SendMail:input_type -> postman.SendMailRequest
	2, // 2: postman.Postman.SendSms:input_type -> postman.SendSmsRequest
	3, // 3: postman.Postman.SendMail:output_type -> google.protobuf.Empty
	3, // 4: postman.Postman.SendSms:output_type -> google.protobuf.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_postman_proto_init() }
func file_api_postman_proto_init() {
	if File_api_postman_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_postman_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMailRequest); i {
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
		file_api_postman_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailParams); i {
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
		file_api_postman_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsRequest); i {
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
			RawDescriptor: file_api_postman_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_postman_proto_goTypes,
		DependencyIndexes: file_api_postman_proto_depIdxs,
		MessageInfos:      file_api_postman_proto_msgTypes,
	}.Build()
	File_api_postman_proto = out.File
	file_api_postman_proto_rawDesc = nil
	file_api_postman_proto_goTypes = nil
	file_api_postman_proto_depIdxs = nil
}
