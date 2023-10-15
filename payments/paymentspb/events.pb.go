// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: paymentspb/events.proto

package paymentspb

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

type InvoicePaid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId string `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *InvoicePaid) Reset() {
	*x = InvoicePaid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentspb_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoicePaid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoicePaid) ProtoMessage() {}

func (x *InvoicePaid) ProtoReflect() protoreflect.Message {
	mi := &file_paymentspb_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvoicePaid.ProtoReflect.Descriptor instead.
func (*InvoicePaid) Descriptor() ([]byte, []int) {
	return file_paymentspb_events_proto_rawDescGZIP(), []int{0}
}

func (x *InvoicePaid) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InvoicePaid) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

var File_paymentspb_events_proto protoreflect.FileDescriptor

var file_paymentspb_events_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x70, 0x62, 0x22, 0x38, 0x0a, 0x0b, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x50, 0x61, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x42,
	0x97, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x70, 0x62, 0x42, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x30, 0x65, 0x64, 0x61, 0x2d, 0x69, 0x6e, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2f, 0x63, 0x68, 0x38, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0xca, 0x02, 0x0a, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x70, 0x62, 0xe2, 0x02, 0x16, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_paymentspb_events_proto_rawDescOnce sync.Once
	file_paymentspb_events_proto_rawDescData = file_paymentspb_events_proto_rawDesc
)

func file_paymentspb_events_proto_rawDescGZIP() []byte {
	file_paymentspb_events_proto_rawDescOnce.Do(func() {
		file_paymentspb_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_paymentspb_events_proto_rawDescData)
	})
	return file_paymentspb_events_proto_rawDescData
}

var file_paymentspb_events_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_paymentspb_events_proto_goTypes = []interface{}{
	(*InvoicePaid)(nil), // 0: paymentspb.InvoicePaid
}
var file_paymentspb_events_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_paymentspb_events_proto_init() }
func file_paymentspb_events_proto_init() {
	if File_paymentspb_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_paymentspb_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvoicePaid); i {
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
			RawDescriptor: file_paymentspb_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_paymentspb_events_proto_goTypes,
		DependencyIndexes: file_paymentspb_events_proto_depIdxs,
		MessageInfos:      file_paymentspb_events_proto_msgTypes,
	}.Build()
	File_paymentspb_events_proto = out.File
	file_paymentspb_events_proto_rawDesc = nil
	file_paymentspb_events_proto_goTypes = nil
	file_paymentspb_events_proto_depIdxs = nil
}
