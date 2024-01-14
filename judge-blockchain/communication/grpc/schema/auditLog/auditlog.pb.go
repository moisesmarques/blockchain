// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: communication/grpc/schema/auditlog.proto

package grpcAuditLog

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

// Send Audit To Admin Request
type SendAuditToAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auditlog     *AuditLogData     `protobuf:"bytes,1,opt,name=auditlog,proto3" json:"auditlog,omitempty"`
	Auditsummary *AuditSummaryData `protobuf:"bytes,2,opt,name=auditsummary,proto3" json:"auditsummary,omitempty"`
}

func (x *SendAuditToAdminRequest) Reset() {
	*x = SendAuditToAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_grpc_schema_auditlog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendAuditToAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendAuditToAdminRequest) ProtoMessage() {}

func (x *SendAuditToAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_communication_grpc_schema_auditlog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendAuditToAdminRequest.ProtoReflect.Descriptor instead.
func (*SendAuditToAdminRequest) Descriptor() ([]byte, []int) {
	return file_communication_grpc_schema_auditlog_proto_rawDescGZIP(), []int{0}
}

func (x *SendAuditToAdminRequest) GetAuditlog() *AuditLogData {
	if x != nil {
		return x.Auditlog
	}
	return nil
}

func (x *SendAuditToAdminRequest) GetAuditsummary() *AuditSummaryData {
	if x != nil {
		return x.Auditsummary
	}
	return nil
}

// The storage nodes response message
type SendAuditToAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendAuditToAdminResponse) Reset() {
	*x = SendAuditToAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_grpc_schema_auditlog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendAuditToAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendAuditToAdminResponse) ProtoMessage() {}

func (x *SendAuditToAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_communication_grpc_schema_auditlog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendAuditToAdminResponse.ProtoReflect.Descriptor instead.
func (*SendAuditToAdminResponse) Descriptor() ([]byte, []int) {
	return file_communication_grpc_schema_auditlog_proto_rawDescGZIP(), []int{1}
}

func (x *SendAuditToAdminResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AuditLogData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity           string `protobuf:"bytes,1,opt,name=Identity,proto3" json:"Identity,omitempty"`
	Chunk              string `protobuf:"bytes,2,opt,name=Chunk,proto3" json:"Chunk,omitempty"`
	TimeInMilliSeconds uint64 `protobuf:"varint,3,opt,name=TimeInMilliSeconds,proto3" json:"TimeInMilliSeconds,omitempty"`
	Result             string `protobuf:"bytes,4,opt,name=Result,proto3" json:"Result,omitempty"`
	Message            string `protobuf:"bytes,5,opt,name=Message,proto3" json:"Message,omitempty"`
	NextAction         string `protobuf:"bytes,6,opt,name=NextAction,proto3" json:"NextAction,omitempty"`
}

func (x *AuditLogData) Reset() {
	*x = AuditLogData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_grpc_schema_auditlog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditLogData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditLogData) ProtoMessage() {}

func (x *AuditLogData) ProtoReflect() protoreflect.Message {
	mi := &file_communication_grpc_schema_auditlog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditLogData.ProtoReflect.Descriptor instead.
func (*AuditLogData) Descriptor() ([]byte, []int) {
	return file_communication_grpc_schema_auditlog_proto_rawDescGZIP(), []int{2}
}

func (x *AuditLogData) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *AuditLogData) GetChunk() string {
	if x != nil {
		return x.Chunk
	}
	return ""
}

func (x *AuditLogData) GetTimeInMilliSeconds() uint64 {
	if x != nil {
		return x.TimeInMilliSeconds
	}
	return 0
}

func (x *AuditLogData) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *AuditLogData) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AuditLogData) GetNextAction() string {
	if x != nil {
		return x.NextAction
	}
	return ""
}

type AuditSummaryData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity      string `protobuf:"bytes,1,opt,name=Identity,proto3" json:"Identity,omitempty"`
	Successes     uint64 `protobuf:"varint,2,opt,name=Successes,proto3" json:"Successes,omitempty"`
	Fails         uint64 `protobuf:"varint,3,opt,name=Fails,proto3" json:"Fails,omitempty"`
	Status        string `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
	Offlines      uint64 `protobuf:"varint,5,opt,name=Offlines,proto3" json:"Offlines,omitempty"`
	TotalCount    uint64 `protobuf:"varint,6,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
	LastAuditTime uint64 `protobuf:"varint,7,opt,name=LastAuditTime,proto3" json:"LastAuditTime,omitempty"`
	Priority      uint64 `protobuf:"varint,8,opt,name=Priority,proto3" json:"Priority,omitempty"`
}

func (x *AuditSummaryData) Reset() {
	*x = AuditSummaryData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_grpc_schema_auditlog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditSummaryData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditSummaryData) ProtoMessage() {}

func (x *AuditSummaryData) ProtoReflect() protoreflect.Message {
	mi := &file_communication_grpc_schema_auditlog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditSummaryData.ProtoReflect.Descriptor instead.
func (*AuditSummaryData) Descriptor() ([]byte, []int) {
	return file_communication_grpc_schema_auditlog_proto_rawDescGZIP(), []int{3}
}

func (x *AuditSummaryData) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *AuditSummaryData) GetSuccesses() uint64 {
	if x != nil {
		return x.Successes
	}
	return 0
}

func (x *AuditSummaryData) GetFails() uint64 {
	if x != nil {
		return x.Fails
	}
	return 0
}

func (x *AuditSummaryData) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AuditSummaryData) GetOfflines() uint64 {
	if x != nil {
		return x.Offlines
	}
	return 0
}

func (x *AuditSummaryData) GetTotalCount() uint64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *AuditSummaryData) GetLastAuditTime() uint64 {
	if x != nil {
		return x.LastAuditTime
	}
	return 0
}

func (x *AuditSummaryData) GetPriority() uint64 {
	if x != nil {
		return x.Priority
	}
	return 0
}

var File_communication_grpc_schema_auditlog_proto protoreflect.FileDescriptor

var file_communication_grpc_schema_auditlog_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x87, 0x01, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x75, 0x64, 0x69, 0x74, 0x54,
	0x6f, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a,
	0x08, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x12, 0x3b,
	0x0a, 0x0c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x34, 0x0a, 0x18, 0x53,
	0x65, 0x6e, 0x64, 0x41, 0x75, 0x64, 0x69, 0x74, 0x54, 0x6f, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0xc2, 0x01, 0x0a, 0x0c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x12, 0x2e, 0x0a, 0x12, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x4d, 0x69,
	0x6c, 0x6c, 0x69, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x12, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4e, 0x65, 0x78, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4e, 0x65, 0x78, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xf8, 0x01, 0x0a, 0x10, 0x41, 0x75, 0x64, 0x69, 0x74,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x46, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x46, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x24, 0x0a, 0x0d, 0x4c, 0x61, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x4c, 0x61, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x32, 0x68, 0x0a, 0x0f, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x54, 0x6f, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x75, 0x64, 0x69, 0x74, 0x54, 0x6f, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x75, 0x64, 0x69, 0x74, 0x54, 0x6f, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x39, 0x0a, 0x16, 0x69,
	0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0d, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0e, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_communication_grpc_schema_auditlog_proto_rawDescOnce sync.Once
	file_communication_grpc_schema_auditlog_proto_rawDescData = file_communication_grpc_schema_auditlog_proto_rawDesc
)

func file_communication_grpc_schema_auditlog_proto_rawDescGZIP() []byte {
	file_communication_grpc_schema_auditlog_proto_rawDescOnce.Do(func() {
		file_communication_grpc_schema_auditlog_proto_rawDescData = protoimpl.X.CompressGZIP(file_communication_grpc_schema_auditlog_proto_rawDescData)
	})
	return file_communication_grpc_schema_auditlog_proto_rawDescData
}

var file_communication_grpc_schema_auditlog_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_communication_grpc_schema_auditlog_proto_goTypes = []interface{}{
	(*SendAuditToAdminRequest)(nil),  // 0: proto.SendAuditToAdminRequest
	(*SendAuditToAdminResponse)(nil), // 1: proto.SendAuditToAdminResponse
	(*AuditLogData)(nil),             // 2: proto.AuditLogData
	(*AuditSummaryData)(nil),         // 3: proto.AuditSummaryData
}
var file_communication_grpc_schema_auditlog_proto_depIdxs = []int32{
	2, // 0: proto.SendAuditToAdminRequest.auditlog:type_name -> proto.AuditLogData
	3, // 1: proto.SendAuditToAdminRequest.auditsummary:type_name -> proto.AuditSummaryData
	0, // 2: proto.AuditLogService.SendAuditToAdmin:input_type -> proto.SendAuditToAdminRequest
	1, // 3: proto.AuditLogService.SendAuditToAdmin:output_type -> proto.SendAuditToAdminResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_communication_grpc_schema_auditlog_proto_init() }
func file_communication_grpc_schema_auditlog_proto_init() {
	if File_communication_grpc_schema_auditlog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_communication_grpc_schema_auditlog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendAuditToAdminRequest); i {
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
		file_communication_grpc_schema_auditlog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendAuditToAdminResponse); i {
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
		file_communication_grpc_schema_auditlog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditLogData); i {
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
		file_communication_grpc_schema_auditlog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditSummaryData); i {
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
			RawDescriptor: file_communication_grpc_schema_auditlog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_communication_grpc_schema_auditlog_proto_goTypes,
		DependencyIndexes: file_communication_grpc_schema_auditlog_proto_depIdxs,
		MessageInfos:      file_communication_grpc_schema_auditlog_proto_msgTypes,
	}.Build()
	File_communication_grpc_schema_auditlog_proto = out.File
	file_communication_grpc_schema_auditlog_proto_rawDesc = nil
	file_communication_grpc_schema_auditlog_proto_goTypes = nil
	file_communication_grpc_schema_auditlog_proto_depIdxs = nil
}
