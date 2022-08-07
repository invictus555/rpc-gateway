// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: src/proto/prod.proto

package service

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GrpcGateWayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsId int32 `protobuf:"varint,1,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"` // 传入的商品id
}

func (x *GrpcGateWayRequest) Reset() {
	*x = GrpcGateWayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_prod_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcGateWayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcGateWayRequest) ProtoMessage() {}

func (x *GrpcGateWayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_prod_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcGateWayRequest.ProtoReflect.Descriptor instead.
func (*GrpcGateWayRequest) Descriptor() ([]byte, []int) {
	return file_src_proto_prod_proto_rawDescGZIP(), []int{0}
}

func (x *GrpcGateWayRequest) GetGoodsId() int32 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

type GrpcGateWayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsStock int32 `protobuf:"varint,1,opt,name=goods_stock,json=goodsStock,proto3" json:"goods_stock,omitempty"` // 商品库存
}

func (x *GrpcGateWayResponse) Reset() {
	*x = GrpcGateWayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_prod_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcGateWayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcGateWayResponse) ProtoMessage() {}

func (x *GrpcGateWayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_prod_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcGateWayResponse.ProtoReflect.Descriptor instead.
func (*GrpcGateWayResponse) Descriptor() ([]byte, []int) {
	return file_src_proto_prod_proto_rawDescGZIP(), []int{1}
}

func (x *GrpcGateWayResponse) GetGoodsStock() int32 {
	if x != nil {
		return x.GoodsStock
	}
	return 0
}

var File_src_proto_prod_proto protoreflect.FileDescriptor

var file_src_proto_prod_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a,
	0x12, 0x47, 0x72, 0x70, 0x63, 0x47, 0x61, 0x74, 0x65, 0x57, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x22, 0x36,
	0x0a, 0x13, 0x47, 0x72, 0x70, 0x63, 0x47, 0x61, 0x74, 0x65, 0x57, 0x61, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x32, 0x83, 0x01, 0x0a, 0x12, 0x47, 0x72, 0x70, 0x63, 0x47,
	0x61, 0x74, 0x65, 0x57, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x47, 0x72, 0x70, 0x63, 0x47, 0x61, 0x74, 0x65, 0x57, 0x61, 0x79, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x72, 0x70, 0x63, 0x47, 0x61, 0x74, 0x65, 0x57, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x72, 0x70, 0x63,
	0x47, 0x61, 0x74, 0x65, 0x57, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x2f, 0x7b, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x0d, 0x5a, 0x0b,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_src_proto_prod_proto_rawDescOnce sync.Once
	file_src_proto_prod_proto_rawDescData = file_src_proto_prod_proto_rawDesc
)

func file_src_proto_prod_proto_rawDescGZIP() []byte {
	file_src_proto_prod_proto_rawDescOnce.Do(func() {
		file_src_proto_prod_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_proto_prod_proto_rawDescData)
	})
	return file_src_proto_prod_proto_rawDescData
}

var file_src_proto_prod_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_proto_prod_proto_goTypes = []interface{}{
	(*GrpcGateWayRequest)(nil),  // 0: service.GrpcGateWayRequest
	(*GrpcGateWayResponse)(nil), // 1: service.GrpcGateWayResponse
}
var file_src_proto_prod_proto_depIdxs = []int32{
	0, // 0: service.GrpcGateWayService.GetGrpcGateWayStock:input_type -> service.GrpcGateWayRequest
	1, // 1: service.GrpcGateWayService.GetGrpcGateWayStock:output_type -> service.GrpcGateWayResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_proto_prod_proto_init() }
func file_src_proto_prod_proto_init() {
	if File_src_proto_prod_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_proto_prod_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcGateWayRequest); i {
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
		file_src_proto_prod_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcGateWayResponse); i {
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
			RawDescriptor: file_src_proto_prod_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_proto_prod_proto_goTypes,
		DependencyIndexes: file_src_proto_prod_proto_depIdxs,
		MessageInfos:      file_src_proto_prod_proto_msgTypes,
	}.Build()
	File_src_proto_prod_proto = out.File
	file_src_proto_prod_proto_rawDesc = nil
	file_src_proto_prod_proto_goTypes = nil
	file_src_proto_prod_proto_depIdxs = nil
}