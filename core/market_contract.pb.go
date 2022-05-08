// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: core/contract/market_contract.proto

package core

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

type MarketSellAssetContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress      []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	SellTokenId       []byte `protobuf:"bytes,2,opt,name=sell_token_id,json=sellTokenId,proto3" json:"sell_token_id,omitempty"`
	SellTokenQuantity int64  `protobuf:"varint,3,opt,name=sell_token_quantity,json=sellTokenQuantity,proto3" json:"sell_token_quantity,omitempty"`
	BuyTokenId        []byte `protobuf:"bytes,4,opt,name=buy_token_id,json=buyTokenId,proto3" json:"buy_token_id,omitempty"`
	BuyTokenQuantity  int64  `protobuf:"varint,5,opt,name=buy_token_quantity,json=buyTokenQuantity,proto3" json:"buy_token_quantity,omitempty"` // min to receive
}

func (x *MarketSellAssetContract) Reset() {
	*x = MarketSellAssetContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_contract_market_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketSellAssetContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketSellAssetContract) ProtoMessage() {}

func (x *MarketSellAssetContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_market_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketSellAssetContract.ProtoReflect.Descriptor instead.
func (*MarketSellAssetContract) Descriptor() ([]byte, []int) {
	return file_core_contract_market_contract_proto_rawDescGZIP(), []int{0}
}

func (x *MarketSellAssetContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *MarketSellAssetContract) GetSellTokenId() []byte {
	if x != nil {
		return x.SellTokenId
	}
	return nil
}

func (x *MarketSellAssetContract) GetSellTokenQuantity() int64 {
	if x != nil {
		return x.SellTokenQuantity
	}
	return 0
}

func (x *MarketSellAssetContract) GetBuyTokenId() []byte {
	if x != nil {
		return x.BuyTokenId
	}
	return nil
}

func (x *MarketSellAssetContract) GetBuyTokenQuantity() int64 {
	if x != nil {
		return x.BuyTokenQuantity
	}
	return 0
}

type MarketCancelOrderContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	OrderId      []byte `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *MarketCancelOrderContract) Reset() {
	*x = MarketCancelOrderContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_contract_market_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketCancelOrderContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketCancelOrderContract) ProtoMessage() {}

func (x *MarketCancelOrderContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_market_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketCancelOrderContract.ProtoReflect.Descriptor instead.
func (*MarketCancelOrderContract) Descriptor() ([]byte, []int) {
	return file_core_contract_market_contract_proto_rawDescGZIP(), []int{1}
}

func (x *MarketCancelOrderContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *MarketCancelOrderContract) GetOrderId() []byte {
	if x != nil {
		return x.OrderId
	}
	return nil
}

var File_core_contract_market_contract_proto protoreflect.FileDescriptor

var file_core_contract_market_contract_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22,
	0xe2, 0x01, 0x0a, 0x17, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x6c, 0x6c, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x22, 0x0a, 0x0d, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x73, 0x65, 0x6c, 0x6c, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x11, 0x73, 0x65, 0x6c, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0c, 0x62, 0x75, 0x79, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x62, 0x75, 0x79, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x62, 0x75, 0x79, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x10, 0x62, 0x75, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x22, 0x5b, 0x0a, 0x19, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x42, 0x42, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x75, 0x73, 0x6d, 0x61,
	0x67, 0x6e, 0x61, 0x2f, 0x74, 0x72, 0x6f, 0x6e, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_contract_market_contract_proto_rawDescOnce sync.Once
	file_core_contract_market_contract_proto_rawDescData = file_core_contract_market_contract_proto_rawDesc
)

func file_core_contract_market_contract_proto_rawDescGZIP() []byte {
	file_core_contract_market_contract_proto_rawDescOnce.Do(func() {
		file_core_contract_market_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_contract_market_contract_proto_rawDescData)
	})
	return file_core_contract_market_contract_proto_rawDescData
}

var file_core_contract_market_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_core_contract_market_contract_proto_goTypes = []interface{}{
	(*MarketSellAssetContract)(nil),   // 0: protocol.MarketSellAssetContract
	(*MarketCancelOrderContract)(nil), // 1: protocol.MarketCancelOrderContract
}
var file_core_contract_market_contract_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_core_contract_market_contract_proto_init() }
func file_core_contract_market_contract_proto_init() {
	if File_core_contract_market_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_contract_market_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketSellAssetContract); i {
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
		file_core_contract_market_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketCancelOrderContract); i {
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
			RawDescriptor: file_core_contract_market_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_contract_market_contract_proto_goTypes,
		DependencyIndexes: file_core_contract_market_contract_proto_depIdxs,
		MessageInfos:      file_core_contract_market_contract_proto_msgTypes,
	}.Build()
	File_core_contract_market_contract_proto = out.File
	file_core_contract_market_contract_proto_rawDesc = nil
	file_core_contract_market_contract_proto_goTypes = nil
	file_core_contract_market_contract_proto_depIdxs = nil
}
