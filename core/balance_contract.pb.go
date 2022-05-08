// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: core/balance_contract.proto

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

type FreezeBalanceContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress    []byte       `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	FrozenBalance   int64        `protobuf:"varint,2,opt,name=frozen_balance,json=frozenBalance,proto3" json:"frozen_balance,omitempty"`
	FrozenDuration  int64        `protobuf:"varint,3,opt,name=frozen_duration,json=frozenDuration,proto3" json:"frozen_duration,omitempty"`
	Resource        ResourceCode `protobuf:"varint,10,opt,name=resource,proto3,enum=protocol.ResourceCode" json:"resource,omitempty"`
	ReceiverAddress []byte       `protobuf:"bytes,15,opt,name=receiver_address,json=receiverAddress,proto3" json:"receiver_address,omitempty"`
}

func (x *FreezeBalanceContract) Reset() {
	*x = FreezeBalanceContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_balance_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FreezeBalanceContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FreezeBalanceContract) ProtoMessage() {}

func (x *FreezeBalanceContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_balance_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FreezeBalanceContract.ProtoReflect.Descriptor instead.
func (*FreezeBalanceContract) Descriptor() ([]byte, []int) {
	return file_core_balance_contract_proto_rawDescGZIP(), []int{0}
}

func (x *FreezeBalanceContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *FreezeBalanceContract) GetFrozenBalance() int64 {
	if x != nil {
		return x.FrozenBalance
	}
	return 0
}

func (x *FreezeBalanceContract) GetFrozenDuration() int64 {
	if x != nil {
		return x.FrozenDuration
	}
	return 0
}

func (x *FreezeBalanceContract) GetResource() ResourceCode {
	if x != nil {
		return x.Resource
	}
	return ResourceCode_BANDWIDTH
}

func (x *FreezeBalanceContract) GetReceiverAddress() []byte {
	if x != nil {
		return x.ReceiverAddress
	}
	return nil
}

type UnfreezeBalanceContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress    []byte       `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Resource        ResourceCode `protobuf:"varint,10,opt,name=resource,proto3,enum=protocol.ResourceCode" json:"resource,omitempty"`
	ReceiverAddress []byte       `protobuf:"bytes,15,opt,name=receiver_address,json=receiverAddress,proto3" json:"receiver_address,omitempty"`
}

func (x *UnfreezeBalanceContract) Reset() {
	*x = UnfreezeBalanceContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_balance_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnfreezeBalanceContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnfreezeBalanceContract) ProtoMessage() {}

func (x *UnfreezeBalanceContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_balance_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnfreezeBalanceContract.ProtoReflect.Descriptor instead.
func (*UnfreezeBalanceContract) Descriptor() ([]byte, []int) {
	return file_core_balance_contract_proto_rawDescGZIP(), []int{1}
}

func (x *UnfreezeBalanceContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *UnfreezeBalanceContract) GetResource() ResourceCode {
	if x != nil {
		return x.Resource
	}
	return ResourceCode_BANDWIDTH
}

func (x *UnfreezeBalanceContract) GetReceiverAddress() []byte {
	if x != nil {
		return x.ReceiverAddress
	}
	return nil
}

type WithdrawBalanceContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
}

func (x *WithdrawBalanceContract) Reset() {
	*x = WithdrawBalanceContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_balance_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawBalanceContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawBalanceContract) ProtoMessage() {}

func (x *WithdrawBalanceContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_balance_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawBalanceContract.ProtoReflect.Descriptor instead.
func (*WithdrawBalanceContract) Descriptor() ([]byte, []int) {
	return file_core_balance_contract_proto_rawDescGZIP(), []int{2}
}

func (x *WithdrawBalanceContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

type TransferContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ToAddress    []byte `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Amount       int64  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransferContract) Reset() {
	*x = TransferContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_balance_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferContract) ProtoMessage() {}

func (x *TransferContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_balance_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferContract.ProtoReflect.Descriptor instead.
func (*TransferContract) Descriptor() ([]byte, []int) {
	return file_core_balance_contract_proto_rawDescGZIP(), []int{3}
}

func (x *TransferContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *TransferContract) GetToAddress() []byte {
	if x != nil {
		return x.ToAddress
	}
	return nil
}

func (x *TransferContract) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransactionBalanceTrace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionIdentifier []byte                               `protobuf:"bytes,1,opt,name=transaction_identifier,json=transactionIdentifier,proto3" json:"transaction_identifier,omitempty"`
	Operation             []*TransactionBalanceTrace_Operation `protobuf:"bytes,2,rep,name=operation,proto3" json:"operation,omitempty"`
	Type                  string                               `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Status                string                               `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *TransactionBalanceTrace) Reset() {
	*x = TransactionBalanceTrace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_balance_contract_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionBalanceTrace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionBalanceTrace) ProtoMessage() {}

func (x *TransactionBalanceTrace) ProtoReflect() protoreflect.Message {
	mi := &file_core_balance_contract_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionBalanceTrace.ProtoReflect.Descriptor instead.
func (*TransactionBalanceTrace) Descriptor() ([]byte, []int) {
	return file_core_balance_contract_proto_rawDescGZIP(), []int{4}
}

func (x *TransactionBalanceTrace) GetTransactionIdentifier() []byte {
	if x != nil {
		return x.TransactionIdentifier
	}
	return nil
}

func (x *TransactionBalanceTrace) GetOperation() []*TransactionBalanceTrace_Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *TransactionBalanceTrace) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TransactionBalanceTrace) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type BlockBalanceTrace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockIdentifier         *BlockBalanceTrace_BlockIdentifier `protobuf:"bytes,1,opt,name=block_identifier,json=blockIdentifier,proto3" json:"block_identifier,omitempty"`
	Timestamp               int64                              `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TransactionBalanceTrace []*TransactionBalanceTrace         `protobuf:"bytes,3,rep,name=transaction_balance_trace,json=transactionBalanceTrace,proto3" json:"transaction_balance_trace,omitempty"` //  BlockIdentifier parent_block_identifier = 4;
}

func (x *BlockBalanceTrace) Reset() {
	*x = BlockBalanceTrace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_balance_contract_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockBalanceTrace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockBalanceTrace) ProtoMessage() {}

func (x *BlockBalanceTrace) ProtoReflect() protoreflect.Message {
	mi := &file_core_balance_contract_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockBalanceTrace.ProtoReflect.Descriptor instead.
func (*BlockBalanceTrace) Descriptor() ([]byte, []int) {
	return file_core_balance_contract_proto_rawDescGZIP(), []int{5}
}

func (x *BlockBalanceTrace) GetBlockIdentifier() *BlockBalanceTrace_BlockIdentifier {
	if x != nil {
		return x.BlockIdentifier
	}
	return nil
}

func (x *BlockBalanceTrace) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *BlockBalanceTrace) GetTransactionBalanceTrace() []*TransactionBalanceTrace {
	if x != nil {
		return x.TransactionBalanceTrace
	}
	return nil
}

type AccountTrace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance     int64 `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
	Placeholder int64 `protobuf:"varint,99,opt,name=placeholder,proto3" json:"placeholder,omitempty"`
}

func (x *AccountTrace) Reset() {
	*x = AccountTrace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_balance_contract_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountTrace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountTrace) ProtoMessage() {}

func (x *AccountTrace) ProtoReflect() protoreflect.Message {
	mi := &file_core_balance_contract_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountTrace.ProtoReflect.Descriptor instead.
func (*AccountTrace) Descriptor() ([]byte, []int) {
	return file_core_balance_contract_proto_rawDescGZIP(), []int{6}
}

func (x *AccountTrace) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *AccountTrace) GetPlaceholder() int64 {
	if x != nil {
		return x.Placeholder
	}
	return 0
}

type AccountIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *AccountIdentifier) Reset() {
	*x = AccountIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_balance_contract_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountIdentifier) ProtoMessage() {}

func (x *AccountIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_core_balance_contract_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountIdentifier.ProtoReflect.Descriptor instead.
func (*AccountIdentifier) Descriptor() ([]byte, []int) {
	return file_core_balance_contract_proto_rawDescGZIP(), []int{7}
}

func (x *AccountIdentifier) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

type AccountBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountIdentifier *AccountIdentifier                 `protobuf:"bytes,1,opt,name=account_identifier,json=accountIdentifier,proto3" json:"account_identifier,omitempty"`
	BlockIdentifier   *BlockBalanceTrace_BlockIdentifier `protobuf:"bytes,2,opt,name=block_identifier,json=blockIdentifier,proto3" json:"block_identifier,omitempty"`
}

func (x *AccountBalanceRequest) Reset() {
	*x = AccountBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_balance_contract_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountBalanceRequest) ProtoMessage() {}

func (x *AccountBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_balance_contract_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountBalanceRequest.ProtoReflect.Descriptor instead.
func (*AccountBalanceRequest) Descriptor() ([]byte, []int) {
	return file_core_balance_contract_proto_rawDescGZIP(), []int{8}
}

func (x *AccountBalanceRequest) GetAccountIdentifier() *AccountIdentifier {
	if x != nil {
		return x.AccountIdentifier
	}
	return nil
}

func (x *AccountBalanceRequest) GetBlockIdentifier() *BlockBalanceTrace_BlockIdentifier {
	if x != nil {
		return x.BlockIdentifier
	}
	return nil
}

type AccountBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance         int64                              `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
	BlockIdentifier *BlockBalanceTrace_BlockIdentifier `protobuf:"bytes,2,opt,name=block_identifier,json=blockIdentifier,proto3" json:"block_identifier,omitempty"`
}

func (x *AccountBalanceResponse) Reset() {
	*x = AccountBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_balance_contract_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountBalanceResponse) ProtoMessage() {}

func (x *AccountBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_balance_contract_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountBalanceResponse.ProtoReflect.Descriptor instead.
func (*AccountBalanceResponse) Descriptor() ([]byte, []int) {
	return file_core_balance_contract_proto_rawDescGZIP(), []int{9}
}

func (x *AccountBalanceResponse) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *AccountBalanceResponse) GetBlockIdentifier() *BlockBalanceTrace_BlockIdentifier {
	if x != nil {
		return x.BlockIdentifier
	}
	return nil
}

type TransactionBalanceTrace_Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationIdentifier int64  `protobuf:"varint,1,opt,name=operation_identifier,json=operationIdentifier,proto3" json:"operation_identifier,omitempty"`
	Address             []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Amount              int64  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransactionBalanceTrace_Operation) Reset() {
	*x = TransactionBalanceTrace_Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_balance_contract_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionBalanceTrace_Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionBalanceTrace_Operation) ProtoMessage() {}

func (x *TransactionBalanceTrace_Operation) ProtoReflect() protoreflect.Message {
	mi := &file_core_balance_contract_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionBalanceTrace_Operation.ProtoReflect.Descriptor instead.
func (*TransactionBalanceTrace_Operation) Descriptor() ([]byte, []int) {
	return file_core_balance_contract_proto_rawDescGZIP(), []int{4, 0}
}

func (x *TransactionBalanceTrace_Operation) GetOperationIdentifier() int64 {
	if x != nil {
		return x.OperationIdentifier
	}
	return 0
}

func (x *TransactionBalanceTrace_Operation) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *TransactionBalanceTrace_Operation) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type BlockBalanceTrace_BlockIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash   []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Number int64  `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *BlockBalanceTrace_BlockIdentifier) Reset() {
	*x = BlockBalanceTrace_BlockIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_balance_contract_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockBalanceTrace_BlockIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockBalanceTrace_BlockIdentifier) ProtoMessage() {}

func (x *BlockBalanceTrace_BlockIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_core_balance_contract_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockBalanceTrace_BlockIdentifier.ProtoReflect.Descriptor instead.
func (*BlockBalanceTrace_BlockIdentifier) Descriptor() ([]byte, []int) {
	return file_core_balance_contract_proto_rawDescGZIP(), []int{5, 0}
}

func (x *BlockBalanceTrace_BlockIdentifier) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *BlockBalanceTrace_BlockIdentifier) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

var File_core_balance_contract_proto protoreflect.FileDescriptor

var file_core_balance_contract_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x11, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x01, 0x0a, 0x15, 0x46,
	0x72, 0x65, 0x65, 0x7a, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x72, 0x6f,
	0x7a, 0x65, 0x6e, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x27, 0x0a, 0x0f, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x66, 0x72, 0x6f, 0x7a, 0x65,
	0x6e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x29, 0x0a,
	0x10, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x9d, 0x01, 0x0a, 0x17, 0x55, 0x6e, 0x66,
	0x72, 0x65, 0x65, 0x7a, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x32, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x29, 0x0a,
	0x10, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x3e, 0x0a, 0x17, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x6e, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xb9, 0x02, 0x0a, 0x17, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x16, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x09, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x1a, 0x70, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x31, 0x0a, 0x14, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0xa7, 0x02, 0x0a, 0x11, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x10, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x52, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x5d, 0x0a, 0x19, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x17, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x72, 0x61, 0x63, 0x65, 0x1a,
	0x3d, 0x0a, 0x0f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x4a,
	0x0a, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x63, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x11, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xbb, 0x01, 0x0a, 0x15, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x12, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x11, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x56, 0x0a, 0x10, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x54, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x8a, 0x01, 0x0a, 0x16, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x10,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x52, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x42, 0x42, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x72, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x75,
	0x73, 0x6d, 0x61, 0x67, 0x6e, 0x61, 0x2f, 0x74, 0x72, 0x6f, 0x6e, 0x2d, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_balance_contract_proto_rawDescOnce sync.Once
	file_core_balance_contract_proto_rawDescData = file_core_balance_contract_proto_rawDesc
)

func file_core_balance_contract_proto_rawDescGZIP() []byte {
	file_core_balance_contract_proto_rawDescOnce.Do(func() {
		file_core_balance_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_balance_contract_proto_rawDescData)
	})
	return file_core_balance_contract_proto_rawDescData
}

var file_core_balance_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_core_balance_contract_proto_goTypes = []interface{}{
	(*FreezeBalanceContract)(nil),             // 0: protocol.FreezeBalanceContract
	(*UnfreezeBalanceContract)(nil),           // 1: protocol.UnfreezeBalanceContract
	(*WithdrawBalanceContract)(nil),           // 2: protocol.WithdrawBalanceContract
	(*TransferContract)(nil),                  // 3: protocol.TransferContract
	(*TransactionBalanceTrace)(nil),           // 4: protocol.TransactionBalanceTrace
	(*BlockBalanceTrace)(nil),                 // 5: protocol.BlockBalanceTrace
	(*AccountTrace)(nil),                      // 6: protocol.AccountTrace
	(*AccountIdentifier)(nil),                 // 7: protocol.AccountIdentifier
	(*AccountBalanceRequest)(nil),             // 8: protocol.AccountBalanceRequest
	(*AccountBalanceResponse)(nil),            // 9: protocol.AccountBalanceResponse
	(*TransactionBalanceTrace_Operation)(nil), // 10: protocol.TransactionBalanceTrace.Operation
	(*BlockBalanceTrace_BlockIdentifier)(nil), // 11: protocol.BlockBalanceTrace.BlockIdentifier
	(ResourceCode)(0),                         // 12: protocol.ResourceCode
}
var file_core_balance_contract_proto_depIdxs = []int32{
	12, // 0: protocol.FreezeBalanceContract.resource:type_name -> protocol.ResourceCode
	12, // 1: protocol.UnfreezeBalanceContract.resource:type_name -> protocol.ResourceCode
	10, // 2: protocol.TransactionBalanceTrace.operation:type_name -> protocol.TransactionBalanceTrace.Operation
	11, // 3: protocol.BlockBalanceTrace.block_identifier:type_name -> protocol.BlockBalanceTrace.BlockIdentifier
	4,  // 4: protocol.BlockBalanceTrace.transaction_balance_trace:type_name -> protocol.TransactionBalanceTrace
	7,  // 5: protocol.AccountBalanceRequest.account_identifier:type_name -> protocol.AccountIdentifier
	11, // 6: protocol.AccountBalanceRequest.block_identifier:type_name -> protocol.BlockBalanceTrace.BlockIdentifier
	11, // 7: protocol.AccountBalanceResponse.block_identifier:type_name -> protocol.BlockBalanceTrace.BlockIdentifier
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_core_balance_contract_proto_init() }
func file_core_balance_contract_proto_init() {
	if File_core_balance_contract_proto != nil {
		return
	}
	file_core_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_core_balance_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FreezeBalanceContract); i {
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
		file_core_balance_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnfreezeBalanceContract); i {
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
		file_core_balance_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawBalanceContract); i {
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
		file_core_balance_contract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferContract); i {
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
		file_core_balance_contract_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionBalanceTrace); i {
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
		file_core_balance_contract_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockBalanceTrace); i {
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
		file_core_balance_contract_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountTrace); i {
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
		file_core_balance_contract_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountIdentifier); i {
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
		file_core_balance_contract_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountBalanceRequest); i {
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
		file_core_balance_contract_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountBalanceResponse); i {
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
		file_core_balance_contract_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionBalanceTrace_Operation); i {
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
		file_core_balance_contract_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockBalanceTrace_BlockIdentifier); i {
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
			RawDescriptor: file_core_balance_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_balance_contract_proto_goTypes,
		DependencyIndexes: file_core_balance_contract_proto_depIdxs,
		MessageInfos:      file_core_balance_contract_proto_msgTypes,
	}.Build()
	File_core_balance_contract_proto = out.File
	file_core_balance_contract_proto_rawDesc = nil
	file_core_balance_contract_proto_goTypes = nil
	file_core_balance_contract_proto_depIdxs = nil
}
