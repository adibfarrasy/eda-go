// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: storespb/messages.proto

package storespb

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

type Store struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location      string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Participating bool   `protobuf:"varint,4,opt,name=participating,proto3" json:"participating,omitempty"`
}

func (x *Store) Reset() {
	*x = Store{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storespb_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Store) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Store) ProtoMessage() {}

func (x *Store) ProtoReflect() protoreflect.Message {
	mi := &file_storespb_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Store.ProtoReflect.Descriptor instead.
func (*Store) Descriptor() ([]byte, []int) {
	return file_storespb_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Store) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Store) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Store) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Store) GetParticipating() bool {
	if x != nil {
		return x.Participating
	}
	return false
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StoreId     string  `protobuf:"bytes,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	Name        string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Sku         string  `protobuf:"bytes,5,opt,name=sku,proto3" json:"sku,omitempty"`
	Price       float64 `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storespb_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_storespb_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_storespb_messages_proto_rawDescGZIP(), []int{1}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *Product) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_storespb_messages_proto protoreflect.FileDescriptor

var file_storespb_messages_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x73, 0x70, 0x62, 0x22, 0x6d, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x22, 0x92, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6b,
	0x75, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x89, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x70, 0x62, 0x42, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x65, 0x64, 0x61, 0x2d, 0x69,
	0x6e, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x68, 0x36, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x73, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x73, 0x70, 0x62, 0xca, 0x02, 0x08, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x70,
	0x62, 0xe2, 0x02, 0x14, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x70, 0x62, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storespb_messages_proto_rawDescOnce sync.Once
	file_storespb_messages_proto_rawDescData = file_storespb_messages_proto_rawDesc
)

func file_storespb_messages_proto_rawDescGZIP() []byte {
	file_storespb_messages_proto_rawDescOnce.Do(func() {
		file_storespb_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_storespb_messages_proto_rawDescData)
	})
	return file_storespb_messages_proto_rawDescData
}

var file_storespb_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_storespb_messages_proto_goTypes = []interface{}{
	(*Store)(nil),   // 0: storespb.Store
	(*Product)(nil), // 1: storespb.Product
}
var file_storespb_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_storespb_messages_proto_init() }
func file_storespb_messages_proto_init() {
	if File_storespb_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storespb_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Store); i {
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
		file_storespb_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
			RawDescriptor: file_storespb_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storespb_messages_proto_goTypes,
		DependencyIndexes: file_storespb_messages_proto_depIdxs,
		MessageInfos:      file_storespb_messages_proto_msgTypes,
	}.Build()
	File_storespb_messages_proto = out.File
	file_storespb_messages_proto_rawDesc = nil
	file_storespb_messages_proto_goTypes = nil
	file_storespb_messages_proto_depIdxs = nil
}
