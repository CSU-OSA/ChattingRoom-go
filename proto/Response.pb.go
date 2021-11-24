// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: proto/Response.proto

package proto

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

type Response_Type int32

const (
	Response_RESULT  Response_Type = 0
	Response_MESSAGE Response_Type = 1
)

// Enum value maps for Response_Type.
var (
	Response_Type_name = map[int32]string{
		0: "RESULT",
		1: "MESSAGE",
	}
	Response_Type_value = map[string]int32{
		"RESULT":  0,
		"MESSAGE": 1,
	}
)

func (x Response_Type) Enum() *Response_Type {
	p := new(Response_Type)
	*p = x
	return p
}

func (x Response_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Response_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_Response_proto_enumTypes[0].Descriptor()
}

func (Response_Type) Type() protoreflect.EnumType {
	return &file_proto_Response_proto_enumTypes[0]
}

func (x Response_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Response_Type.Descriptor instead.
func (Response_Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_Response_proto_rawDescGZIP(), []int{0, 0}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    Response_Type      `protobuf:"varint,1,opt,name=type,proto3,enum=Response_Type" json:"type,omitempty"`
	Message []*ResponseMessage `protobuf:"bytes,2,rep,name=message,proto3" json:"message,omitempty"`
	Result  *Result            `protobuf:"bytes,3,opt,name=result,proto3,oneof" json:"result,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_Response_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetType() Response_Type {
	if x != nil {
		return x.Type
	}
	return Response_RESULT
}

func (x *Response) GetMessage() []*ResponseMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Response) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_proto_Response_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *Result) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type ResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel  string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	FromNick string `protobuf:"bytes,2,opt,name=fromNick,proto3" json:"fromNick,omitempty"`
	Content  string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ResponseMessage) Reset() {
	*x = ResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMessage) ProtoMessage() {}

func (x *ResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMessage.ProtoReflect.Descriptor instead.
func (*ResponseMessage) Descriptor() ([]byte, []int) {
	return file_proto_Response_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseMessage) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *ResponseMessage) GetFromNick() string {
	if x != nil {
		return x.FromNick
	}
	return ""
}

func (x *ResponseMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_proto_Response_proto protoreflect.FileDescriptor

var file_proto_Response_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x88, 0x01, 0x01, 0x22, 0x1f, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x38, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22,
	0x61, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x72, 0x6f, 0x6d, 0x4e, 0x69, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x72, 0x6f, 0x6d, 0x4e, 0x69, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x42, 0x0a, 0x48, 0x01, 0x5a, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_Response_proto_rawDescOnce sync.Once
	file_proto_Response_proto_rawDescData = file_proto_Response_proto_rawDesc
)

func file_proto_Response_proto_rawDescGZIP() []byte {
	file_proto_Response_proto_rawDescOnce.Do(func() {
		file_proto_Response_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_Response_proto_rawDescData)
	})
	return file_proto_Response_proto_rawDescData
}

var file_proto_Response_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_Response_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_Response_proto_goTypes = []interface{}{
	(Response_Type)(0),      // 0: Response.Type
	(*Response)(nil),        // 1: Response
	(*Result)(nil),          // 2: Result
	(*ResponseMessage)(nil), // 3: ResponseMessage
}
var file_proto_Response_proto_depIdxs = []int32{
	0, // 0: Response.type:type_name -> Response.Type
	3, // 1: Response.message:type_name -> ResponseMessage
	2, // 2: Response.result:type_name -> Result
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_Response_proto_init() }
func file_proto_Response_proto_init() {
	if File_proto_Response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_Response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_Response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_proto_Response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMessage); i {
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
	file_proto_Response_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_Response_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_Response_proto_goTypes,
		DependencyIndexes: file_proto_Response_proto_depIdxs,
		EnumInfos:         file_proto_Response_proto_enumTypes,
		MessageInfos:      file_proto_Response_proto_msgTypes,
	}.Build()
	File_proto_Response_proto = out.File
	file_proto_Response_proto_rawDesc = nil
	file_proto_Response_proto_goTypes = nil
	file_proto_Response_proto_depIdxs = nil
}
