// protoc --proto_path=internal --go_out=generated --go-grpc_out=generated internal/orders/EquityOrderDtls.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: orders/EquityOrderDtls.proto

package __

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

type OrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrdClmMtchAccnt string `protobuf:"bytes,1,opt,name=ord_clm_mtch_accnt,json=ordClmMtchAccnt,proto3" json:"ord_clm_mtch_accnt,omitempty"`
}

func (x *OrderRequest) Reset() {
	*x = OrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_EquityOrderDtls_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRequest) ProtoMessage() {}

func (x *OrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_EquityOrderDtls_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRequest.ProtoReflect.Descriptor instead.
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return file_orders_EquityOrderDtls_proto_rawDescGZIP(), []int{0}
}

func (x *OrderRequest) GetOrdClmMtchAccnt() string {
	if x != nil {
		return x.OrdClmMtchAccnt
	}
	return ""
}

type OrdOrdrDtls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrdClmMtchAccnt string  `protobuf:"bytes,1,opt,name=ord_clm_mtch_accnt,json=ordClmMtchAccnt,proto3" json:"ord_clm_mtch_accnt,omitempty"`
	OrdStckCd       string  `protobuf:"bytes,2,opt,name=ord_stck_cd,json=ordStckCd,proto3" json:"ord_stck_cd,omitempty"`
	OrdOrdrDt       string  `protobuf:"bytes,3,opt,name=ord_ordr_dt,json=ordOrdrDt,proto3" json:"ord_ordr_dt,omitempty"`
	OrdOrdrFlw      string  `protobuf:"bytes,4,opt,name=ord_ordr_flw,json=ordOrdrFlw,proto3" json:"ord_ordr_flw,omitempty"`
	OrdOrdrQty      int32   `protobuf:"varint,5,opt,name=ord_ordr_qty,json=ordOrdrQty,proto3" json:"ord_ordr_qty,omitempty"`
	OrdLmtRt        float64 `protobuf:"fixed64,6,opt,name=ord_lmt_rt,json=ordLmtRt,proto3" json:"ord_lmt_rt,omitempty"`
	OrdOrdrStts     string  `protobuf:"bytes,7,opt,name=ord_ordr_stts,json=ordOrdrStts,proto3" json:"ord_ordr_stts,omitempty"`
}

func (x *OrdOrdrDtls) Reset() {
	*x = OrdOrdrDtls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_EquityOrderDtls_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrdOrdrDtls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdOrdrDtls) ProtoMessage() {}

func (x *OrdOrdrDtls) ProtoReflect() protoreflect.Message {
	mi := &file_orders_EquityOrderDtls_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdOrdrDtls.ProtoReflect.Descriptor instead.
func (*OrdOrdrDtls) Descriptor() ([]byte, []int) {
	return file_orders_EquityOrderDtls_proto_rawDescGZIP(), []int{1}
}

func (x *OrdOrdrDtls) GetOrdClmMtchAccnt() string {
	if x != nil {
		return x.OrdClmMtchAccnt
	}
	return ""
}

func (x *OrdOrdrDtls) GetOrdStckCd() string {
	if x != nil {
		return x.OrdStckCd
	}
	return ""
}

func (x *OrdOrdrDtls) GetOrdOrdrDt() string {
	if x != nil {
		return x.OrdOrdrDt
	}
	return ""
}

func (x *OrdOrdrDtls) GetOrdOrdrFlw() string {
	if x != nil {
		return x.OrdOrdrFlw
	}
	return ""
}

func (x *OrdOrdrDtls) GetOrdOrdrQty() int32 {
	if x != nil {
		return x.OrdOrdrQty
	}
	return 0
}

func (x *OrdOrdrDtls) GetOrdLmtRt() float64 {
	if x != nil {
		return x.OrdLmtRt
	}
	return 0
}

func (x *OrdOrdrDtls) GetOrdOrdrStts() string {
	if x != nil {
		return x.OrdOrdrStts
	}
	return ""
}

type OrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrdDtls []*OrdOrdrDtls `protobuf:"bytes,1,rep,name=ord_dtls,json=ordDtls,proto3" json:"ord_dtls,omitempty"`
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_EquityOrderDtls_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orders_EquityOrderDtls_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_orders_EquityOrderDtls_proto_rawDescGZIP(), []int{2}
}

func (x *OrderResponse) GetOrdDtls() []*OrdOrdrDtls {
	if x != nil {
		return x.OrdDtls
	}
	return nil
}

var File_orders_EquityOrderDtls_proto protoreflect.FileDescriptor

var file_orders_EquityOrderDtls_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x45, 0x71, 0x75, 0x69, 0x74, 0x79, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x44, 0x74, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x3b, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x12, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x6c,
	0x6d, 0x5f, 0x6d, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x63, 0x63, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x6f, 0x72, 0x64, 0x43, 0x6c, 0x6d, 0x4d, 0x74, 0x63, 0x68, 0x41, 0x63,
	0x63, 0x6e, 0x74, 0x22, 0x80, 0x02, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x4f, 0x72, 0x64, 0x72, 0x44,
	0x74, 0x6c, 0x73, 0x12, 0x2b, 0x0a, 0x12, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x6c, 0x6d, 0x5f, 0x6d,
	0x74, 0x63, 0x68, 0x5f, 0x61, 0x63, 0x63, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x6f, 0x72, 0x64, 0x43, 0x6c, 0x6d, 0x4d, 0x74, 0x63, 0x68, 0x41, 0x63, 0x63, 0x6e, 0x74,
	0x12, 0x1e, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x5f, 0x73, 0x74, 0x63, 0x6b, 0x5f, 0x63, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x53, 0x74, 0x63, 0x6b, 0x43, 0x64,
	0x12, 0x1e, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x5f, 0x6f, 0x72, 0x64, 0x72, 0x5f, 0x64, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x4f, 0x72, 0x64, 0x72, 0x44, 0x74,
	0x12, 0x20, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x5f, 0x6f, 0x72, 0x64, 0x72, 0x5f, 0x66, 0x6c, 0x77,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x4f, 0x72, 0x64, 0x72, 0x46,
	0x6c, 0x77, 0x12, 0x20, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x5f, 0x6f, 0x72, 0x64, 0x72, 0x5f, 0x71,
	0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x4f, 0x72, 0x64,
	0x72, 0x51, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x5f, 0x6c, 0x6d, 0x74, 0x5f,
	0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x4c, 0x6d, 0x74,
	0x52, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x6f, 0x72, 0x64, 0x5f, 0x6f, 0x72, 0x64, 0x72, 0x5f, 0x73,
	0x74, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x4f, 0x72,
	0x64, 0x72, 0x53, 0x74, 0x74, 0x73, 0x22, 0x3f, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x5f, 0x64,
	0x74, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x4f, 0x72, 0x64, 0x72, 0x44, 0x74, 0x6c, 0x73, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x44, 0x74, 0x6c, 0x73, 0x32, 0x47, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orders_EquityOrderDtls_proto_rawDescOnce sync.Once
	file_orders_EquityOrderDtls_proto_rawDescData = file_orders_EquityOrderDtls_proto_rawDesc
)

func file_orders_EquityOrderDtls_proto_rawDescGZIP() []byte {
	file_orders_EquityOrderDtls_proto_rawDescOnce.Do(func() {
		file_orders_EquityOrderDtls_proto_rawDescData = protoimpl.X.CompressGZIP(file_orders_EquityOrderDtls_proto_rawDescData)
	})
	return file_orders_EquityOrderDtls_proto_rawDescData
}

var file_orders_EquityOrderDtls_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_orders_EquityOrderDtls_proto_goTypes = []any{
	(*OrderRequest)(nil),  // 0: orders.OrderRequest
	(*OrdOrdrDtls)(nil),   // 1: orders.OrdOrdrDtls
	(*OrderResponse)(nil), // 2: orders.OrderResponse
}
var file_orders_EquityOrderDtls_proto_depIdxs = []int32{
	1, // 0: orders.OrderResponse.ord_dtls:type_name -> orders.OrdOrdrDtls
	0, // 1: orders.OrderService.GetOrder:input_type -> orders.OrderRequest
	2, // 2: orders.OrderService.GetOrder:output_type -> orders.OrderResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_orders_EquityOrderDtls_proto_init() }
func file_orders_EquityOrderDtls_proto_init() {
	if File_orders_EquityOrderDtls_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orders_EquityOrderDtls_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*OrderRequest); i {
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
		file_orders_EquityOrderDtls_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*OrdOrdrDtls); i {
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
		file_orders_EquityOrderDtls_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*OrderResponse); i {
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
			RawDescriptor: file_orders_EquityOrderDtls_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orders_EquityOrderDtls_proto_goTypes,
		DependencyIndexes: file_orders_EquityOrderDtls_proto_depIdxs,
		MessageInfos:      file_orders_EquityOrderDtls_proto_msgTypes,
	}.Build()
	File_orders_EquityOrderDtls_proto = out.File
	file_orders_EquityOrderDtls_proto_rawDesc = nil
	file_orders_EquityOrderDtls_proto_goTypes = nil
	file_orders_EquityOrderDtls_proto_depIdxs = nil
}
