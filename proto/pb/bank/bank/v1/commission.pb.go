// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: bank/v1/commission.proto

package bankv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Commission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TransactionId string                 `protobuf:"bytes,2,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Amount        *Money                 `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Rate          float64                `protobuf:"fixed64,4,opt,name=rate,proto3" json:"rate,omitempty"`
	RecipientId   string                 `protobuf:"bytes,5,opt,name=recipient_id,json=recipientId,proto3" json:"recipient_id,omitempty"`
	BusinessType  string                 `protobuf:"bytes,6,opt,name=business_type,json=businessType,proto3" json:"business_type,omitempty"`
	Status        string                 `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	PaidAt        *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=paid_at,json=paidAt,proto3" json:"paid_at,omitempty"`
	Description   string                 `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Commission) Reset() {
	*x = Commission{}
	mi := &file_bank_v1_commission_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Commission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commission) ProtoMessage() {}

func (x *Commission) ProtoReflect() protoreflect.Message {
	mi := &file_bank_v1_commission_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commission.ProtoReflect.Descriptor instead.
func (*Commission) Descriptor() ([]byte, []int) {
	return file_bank_v1_commission_proto_rawDescGZIP(), []int{0}
}

func (x *Commission) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Commission) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *Commission) GetAmount() *Money {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Commission) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *Commission) GetRecipientId() string {
	if x != nil {
		return x.RecipientId
	}
	return ""
}

func (x *Commission) GetBusinessType() string {
	if x != nil {
		return x.BusinessType
	}
	return ""
}

func (x *Commission) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Commission) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Commission) GetPaidAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PaidAt
	}
	return nil
}

func (x *Commission) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CalculateCommissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction  *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	BusinessType string       `protobuf:"bytes,2,opt,name=business_type,json=businessType,proto3" json:"business_type,omitempty"`
}

func (x *CalculateCommissionRequest) Reset() {
	*x = CalculateCommissionRequest{}
	mi := &file_bank_v1_commission_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CalculateCommissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateCommissionRequest) ProtoMessage() {}

func (x *CalculateCommissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_v1_commission_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateCommissionRequest.ProtoReflect.Descriptor instead.
func (*CalculateCommissionRequest) Descriptor() ([]byte, []int) {
	return file_bank_v1_commission_proto_rawDescGZIP(), []int{1}
}

func (x *CalculateCommissionRequest) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *CalculateCommissionRequest) GetBusinessType() string {
	if x != nil {
		return x.BusinessType
	}
	return ""
}

type CalculateCommissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commission *Commission `protobuf:"bytes,1,opt,name=commission,proto3" json:"commission,omitempty"`
}

func (x *CalculateCommissionResponse) Reset() {
	*x = CalculateCommissionResponse{}
	mi := &file_bank_v1_commission_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CalculateCommissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateCommissionResponse) ProtoMessage() {}

func (x *CalculateCommissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_v1_commission_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateCommissionResponse.ProtoReflect.Descriptor instead.
func (*CalculateCommissionResponse) Descriptor() ([]byte, []int) {
	return file_bank_v1_commission_proto_rawDescGZIP(), []int{2}
}

func (x *CalculateCommissionResponse) GetCommission() *Commission {
	if x != nil {
		return x.Commission
	}
	return nil
}

type ProcessCommissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommissionId string `protobuf:"bytes,1,opt,name=commission_id,json=commissionId,proto3" json:"commission_id,omitempty"`
}

func (x *ProcessCommissionRequest) Reset() {
	*x = ProcessCommissionRequest{}
	mi := &file_bank_v1_commission_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessCommissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessCommissionRequest) ProtoMessage() {}

func (x *ProcessCommissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_v1_commission_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessCommissionRequest.ProtoReflect.Descriptor instead.
func (*ProcessCommissionRequest) Descriptor() ([]byte, []int) {
	return file_bank_v1_commission_proto_rawDescGZIP(), []int{3}
}

func (x *ProcessCommissionRequest) GetCommissionId() string {
	if x != nil {
		return x.CommissionId
	}
	return ""
}

type ProcessCommissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ProcessCommissionResponse) Reset() {
	*x = ProcessCommissionResponse{}
	mi := &file_bank_v1_commission_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessCommissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessCommissionResponse) ProtoMessage() {}

func (x *ProcessCommissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_v1_commission_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessCommissionResponse.ProtoReflect.Descriptor instead.
func (*ProcessCommissionResponse) Descriptor() ([]byte, []int) {
	return file_bank_v1_commission_proto_rawDescGZIP(), []int{4}
}

func (x *ProcessCommissionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetCommissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommissionId string `protobuf:"bytes,1,opt,name=commission_id,json=commissionId,proto3" json:"commission_id,omitempty"`
}

func (x *GetCommissionRequest) Reset() {
	*x = GetCommissionRequest{}
	mi := &file_bank_v1_commission_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCommissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommissionRequest) ProtoMessage() {}

func (x *GetCommissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_v1_commission_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommissionRequest.ProtoReflect.Descriptor instead.
func (*GetCommissionRequest) Descriptor() ([]byte, []int) {
	return file_bank_v1_commission_proto_rawDescGZIP(), []int{5}
}

func (x *GetCommissionRequest) GetCommissionId() string {
	if x != nil {
		return x.CommissionId
	}
	return ""
}

type GetCommissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commission *Commission `protobuf:"bytes,1,opt,name=commission,proto3" json:"commission,omitempty"`
}

func (x *GetCommissionResponse) Reset() {
	*x = GetCommissionResponse{}
	mi := &file_bank_v1_commission_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCommissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommissionResponse) ProtoMessage() {}

func (x *GetCommissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_v1_commission_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommissionResponse.ProtoReflect.Descriptor instead.
func (*GetCommissionResponse) Descriptor() ([]byte, []int) {
	return file_bank_v1_commission_proto_rawDescGZIP(), []int{6}
}

func (x *GetCommissionResponse) GetCommission() *Commission {
	if x != nil {
		return x.Commission
	}
	return nil
}

type GetPendingCommissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPendingCommissionsRequest) Reset() {
	*x = GetPendingCommissionsRequest{}
	mi := &file_bank_v1_commission_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPendingCommissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPendingCommissionsRequest) ProtoMessage() {}

func (x *GetPendingCommissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_v1_commission_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPendingCommissionsRequest.ProtoReflect.Descriptor instead.
func (*GetPendingCommissionsRequest) Descriptor() ([]byte, []int) {
	return file_bank_v1_commission_proto_rawDescGZIP(), []int{7}
}

type GetPendingCommissionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commissions []*Commission `protobuf:"bytes,1,rep,name=commissions,proto3" json:"commissions,omitempty"`
}

func (x *GetPendingCommissionsResponse) Reset() {
	*x = GetPendingCommissionsResponse{}
	mi := &file_bank_v1_commission_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPendingCommissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPendingCommissionsResponse) ProtoMessage() {}

func (x *GetPendingCommissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_v1_commission_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPendingCommissionsResponse.ProtoReflect.Descriptor instead.
func (*GetPendingCommissionsResponse) Descriptor() ([]byte, []int) {
	return file_bank_v1_commission_proto_rawDescGZIP(), []int{8}
}

func (x *GetPendingCommissionsResponse) GetCommissions() []*Commission {
	if x != nil {
		return x.Commissions
	}
	return nil
}

type GetFailedCommissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFailedCommissionsRequest) Reset() {
	*x = GetFailedCommissionsRequest{}
	mi := &file_bank_v1_commission_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFailedCommissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFailedCommissionsRequest) ProtoMessage() {}

func (x *GetFailedCommissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_v1_commission_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFailedCommissionsRequest.ProtoReflect.Descriptor instead.
func (*GetFailedCommissionsRequest) Descriptor() ([]byte, []int) {
	return file_bank_v1_commission_proto_rawDescGZIP(), []int{9}
}

type GetFailedCommissionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commissions []*Commission `protobuf:"bytes,1,rep,name=commissions,proto3" json:"commissions,omitempty"`
}

func (x *GetFailedCommissionsResponse) Reset() {
	*x = GetFailedCommissionsResponse{}
	mi := &file_bank_v1_commission_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFailedCommissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFailedCommissionsResponse) ProtoMessage() {}

func (x *GetFailedCommissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_v1_commission_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFailedCommissionsResponse.ProtoReflect.Descriptor instead.
func (*GetFailedCommissionsResponse) Descriptor() ([]byte, []int) {
	return file_bank_v1_commission_proto_rawDescGZIP(), []int{10}
}

func (x *GetFailedCommissionsResponse) GetCommissions() []*Commission {
	if x != nil {
		return x.Commissions
	}
	return nil
}

type RetryFailedCommissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RetryFailedCommissionsRequest) Reset() {
	*x = RetryFailedCommissionsRequest{}
	mi := &file_bank_v1_commission_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RetryFailedCommissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryFailedCommissionsRequest) ProtoMessage() {}

func (x *RetryFailedCommissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_v1_commission_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryFailedCommissionsRequest.ProtoReflect.Descriptor instead.
func (*RetryFailedCommissionsRequest) Descriptor() ([]byte, []int) {
	return file_bank_v1_commission_proto_rawDescGZIP(), []int{11}
}

type RetryFailedCommissionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetriedCount int32 `protobuf:"varint,1,opt,name=retried_count,json=retriedCount,proto3" json:"retried_count,omitempty"`
	SuccessCount int32 `protobuf:"varint,2,opt,name=success_count,json=successCount,proto3" json:"success_count,omitempty"`
}

func (x *RetryFailedCommissionsResponse) Reset() {
	*x = RetryFailedCommissionsResponse{}
	mi := &file_bank_v1_commission_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RetryFailedCommissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryFailedCommissionsResponse) ProtoMessage() {}

func (x *RetryFailedCommissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_v1_commission_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryFailedCommissionsResponse.ProtoReflect.Descriptor instead.
func (*RetryFailedCommissionsResponse) Descriptor() ([]byte, []int) {
	return file_bank_v1_commission_proto_rawDescGZIP(), []int{12}
}

func (x *RetryFailedCommissionsResponse) GetRetriedCount() int32 {
	if x != nil {
		return x.RetriedCount
	}
	return 0
}

func (x *RetryFailedCommissionsResponse) GetSuccessCount() int32 {
	if x != nil {
		return x.SuccessCount
	}
	return 0
}

var File_bank_v1_commission_proto protoreflect.FileDescriptor

var file_bank_v1_commission_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x61, 0x6e, 0x6b,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x62, 0x61, 0x6e, 0x6b,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x02, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61,
	0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x70, 0x61, 0x69, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x06, 0x70, 0x61, 0x69, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x79, 0x0a, 0x1a, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62,
	0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x52, 0x0a, 0x1b, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x19, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x3b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x4c, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x61, 0x6e,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x1e, 0x0a, 0x1c, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x56, 0x0a, 0x1d, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0b,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x1d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x55, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x1f, 0x0a, 0x1d, 0x52, 0x65, 0x74,
	0x72, 0x79, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6a, 0x0a, 0x1e, 0x52, 0x65,
	0x74, 0x72, 0x79, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xd9, 0x04, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x13,
	0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a,
	0x0a, 0x11, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x62, 0x61,
	0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x61, 0x6e,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x25, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x62, 0x61, 0x6e,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x63, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x24, 0x2e, 0x62, 0x61, 0x6e,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a, 0x16, 0x52, 0x65, 0x74, 0x72, 0x79,
	0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x26, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x72,
	0x79, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x62, 0x61, 0x6e, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x24, 0x5a, 0x22, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x76,
	0x31, 0x3b, 0x62, 0x61, 0x6e, 0x6b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bank_v1_commission_proto_rawDescOnce sync.Once
	file_bank_v1_commission_proto_rawDescData = file_bank_v1_commission_proto_rawDesc
)

func file_bank_v1_commission_proto_rawDescGZIP() []byte {
	file_bank_v1_commission_proto_rawDescOnce.Do(func() {
		file_bank_v1_commission_proto_rawDescData = protoimpl.X.CompressGZIP(file_bank_v1_commission_proto_rawDescData)
	})
	return file_bank_v1_commission_proto_rawDescData
}

var file_bank_v1_commission_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_bank_v1_commission_proto_goTypes = []any{
	(*Commission)(nil),                     // 0: bank.v1.Commission
	(*CalculateCommissionRequest)(nil),     // 1: bank.v1.CalculateCommissionRequest
	(*CalculateCommissionResponse)(nil),    // 2: bank.v1.CalculateCommissionResponse
	(*ProcessCommissionRequest)(nil),       // 3: bank.v1.ProcessCommissionRequest
	(*ProcessCommissionResponse)(nil),      // 4: bank.v1.ProcessCommissionResponse
	(*GetCommissionRequest)(nil),           // 5: bank.v1.GetCommissionRequest
	(*GetCommissionResponse)(nil),          // 6: bank.v1.GetCommissionResponse
	(*GetPendingCommissionsRequest)(nil),   // 7: bank.v1.GetPendingCommissionsRequest
	(*GetPendingCommissionsResponse)(nil),  // 8: bank.v1.GetPendingCommissionsResponse
	(*GetFailedCommissionsRequest)(nil),    // 9: bank.v1.GetFailedCommissionsRequest
	(*GetFailedCommissionsResponse)(nil),   // 10: bank.v1.GetFailedCommissionsResponse
	(*RetryFailedCommissionsRequest)(nil),  // 11: bank.v1.RetryFailedCommissionsRequest
	(*RetryFailedCommissionsResponse)(nil), // 12: bank.v1.RetryFailedCommissionsResponse
	(*Money)(nil),                          // 13: bank.v1.Money
	(*timestamppb.Timestamp)(nil),          // 14: google.protobuf.Timestamp
	(*Transaction)(nil),                    // 15: bank.v1.Transaction
}
var file_bank_v1_commission_proto_depIdxs = []int32{
	13, // 0: bank.v1.Commission.amount:type_name -> bank.v1.Money
	14, // 1: bank.v1.Commission.created_at:type_name -> google.protobuf.Timestamp
	14, // 2: bank.v1.Commission.paid_at:type_name -> google.protobuf.Timestamp
	15, // 3: bank.v1.CalculateCommissionRequest.transaction:type_name -> bank.v1.Transaction
	0,  // 4: bank.v1.CalculateCommissionResponse.commission:type_name -> bank.v1.Commission
	0,  // 5: bank.v1.GetCommissionResponse.commission:type_name -> bank.v1.Commission
	0,  // 6: bank.v1.GetPendingCommissionsResponse.commissions:type_name -> bank.v1.Commission
	0,  // 7: bank.v1.GetFailedCommissionsResponse.commissions:type_name -> bank.v1.Commission
	1,  // 8: bank.v1.CommissionService.CalculateCommission:input_type -> bank.v1.CalculateCommissionRequest
	3,  // 9: bank.v1.CommissionService.ProcessCommission:input_type -> bank.v1.ProcessCommissionRequest
	5,  // 10: bank.v1.CommissionService.GetCommission:input_type -> bank.v1.GetCommissionRequest
	7,  // 11: bank.v1.CommissionService.GetPendingCommissions:input_type -> bank.v1.GetPendingCommissionsRequest
	9,  // 12: bank.v1.CommissionService.GetFailedCommissions:input_type -> bank.v1.GetFailedCommissionsRequest
	11, // 13: bank.v1.CommissionService.RetryFailedCommissions:input_type -> bank.v1.RetryFailedCommissionsRequest
	2,  // 14: bank.v1.CommissionService.CalculateCommission:output_type -> bank.v1.CalculateCommissionResponse
	4,  // 15: bank.v1.CommissionService.ProcessCommission:output_type -> bank.v1.ProcessCommissionResponse
	6,  // 16: bank.v1.CommissionService.GetCommission:output_type -> bank.v1.GetCommissionResponse
	8,  // 17: bank.v1.CommissionService.GetPendingCommissions:output_type -> bank.v1.GetPendingCommissionsResponse
	10, // 18: bank.v1.CommissionService.GetFailedCommissions:output_type -> bank.v1.GetFailedCommissionsResponse
	12, // 19: bank.v1.CommissionService.RetryFailedCommissions:output_type -> bank.v1.RetryFailedCommissionsResponse
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_bank_v1_commission_proto_init() }
func file_bank_v1_commission_proto_init() {
	if File_bank_v1_commission_proto != nil {
		return
	}
	file_bank_v1_common_proto_init()
	file_bank_v1_transaction_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bank_v1_commission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bank_v1_commission_proto_goTypes,
		DependencyIndexes: file_bank_v1_commission_proto_depIdxs,
		MessageInfos:      file_bank_v1_commission_proto_msgTypes,
	}.Build()
	File_bank_v1_commission_proto = out.File
	file_bank_v1_commission_proto_rawDesc = nil
	file_bank_v1_commission_proto_goTypes = nil
	file_bank_v1_commission_proto_depIdxs = nil
}