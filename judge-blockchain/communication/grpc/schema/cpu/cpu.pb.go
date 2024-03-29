// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: communication/grpc/schema/cpu.proto

package cpu

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

type CPURequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActionId            string   `protobuf:"bytes,1,opt,name=action_id,json=actionId,proto3" json:"action_id,omitempty"`
	ActionName          string   `protobuf:"bytes,2,opt,name=action_name,json=actionName,proto3" json:"action_name,omitempty"`
	Function            string   `protobuf:"bytes,3,opt,name=function,proto3" json:"function,omitempty"`
	SecreteShare        string   `protobuf:"bytes,4,opt,name=secrete_share,json=secreteShare,proto3" json:"secrete_share,omitempty"`
	Sender              string   `protobuf:"bytes,5,opt,name=sender,proto3" json:"sender,omitempty"`
	ParticipantingNodes []string `protobuf:"bytes,6,rep,name=participanting_nodes,json=participantingNodes,proto3" json:"participanting_nodes,omitempty"`
}

func (x *CPURequest) Reset() {
	*x = CPURequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_grpc_schema_cpu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPURequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPURequest) ProtoMessage() {}

func (x *CPURequest) ProtoReflect() protoreflect.Message {
	mi := &file_communication_grpc_schema_cpu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPURequest.ProtoReflect.Descriptor instead.
func (*CPURequest) Descriptor() ([]byte, []int) {
	return file_communication_grpc_schema_cpu_proto_rawDescGZIP(), []int{0}
}

func (x *CPURequest) GetActionId() string {
	if x != nil {
		return x.ActionId
	}
	return ""
}

func (x *CPURequest) GetActionName() string {
	if x != nil {
		return x.ActionName
	}
	return ""
}

func (x *CPURequest) GetFunction() string {
	if x != nil {
		return x.Function
	}
	return ""
}

func (x *CPURequest) GetSecreteShare() string {
	if x != nil {
		return x.SecreteShare
	}
	return ""
}

func (x *CPURequest) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *CPURequest) GetParticipantingNodes() []string {
	if x != nil {
		return x.ParticipantingNodes
	}
	return nil
}

type CPUConsensusResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActionId   string                      `protobuf:"bytes,1,opt,name=action_id,json=actionId,proto3" json:"action_id,omitempty"`
	ActionName string                      `protobuf:"bytes,2,opt,name=action_name,json=actionName,proto3" json:"action_name,omitempty"`
	SenderNode string                      `protobuf:"bytes,3,opt,name=sender_node,json=senderNode,proto3" json:"sender_node,omitempty"`
	Verdict    *CPUConsensusResult_Verdict `protobuf:"bytes,4,opt,name=verdict,proto3" json:"verdict,omitempty"`
}

func (x *CPUConsensusResult) Reset() {
	*x = CPUConsensusResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_grpc_schema_cpu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPUConsensusResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPUConsensusResult) ProtoMessage() {}

func (x *CPUConsensusResult) ProtoReflect() protoreflect.Message {
	mi := &file_communication_grpc_schema_cpu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPUConsensusResult.ProtoReflect.Descriptor instead.
func (*CPUConsensusResult) Descriptor() ([]byte, []int) {
	return file_communication_grpc_schema_cpu_proto_rawDescGZIP(), []int{1}
}

func (x *CPUConsensusResult) GetActionId() string {
	if x != nil {
		return x.ActionId
	}
	return ""
}

func (x *CPUConsensusResult) GetActionName() string {
	if x != nil {
		return x.ActionName
	}
	return ""
}

func (x *CPUConsensusResult) GetSenderNode() string {
	if x != nil {
		return x.SenderNode
	}
	return ""
}

func (x *CPUConsensusResult) GetVerdict() *CPUConsensusResult_Verdict {
	if x != nil {
		return x.Verdict
	}
	return nil
}

type CPUConsensusResult_Nodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ResultHash              string `protobuf:"bytes,2,opt,name=result_hash,json=resultHash,proto3" json:"result_hash,omitempty"`
	DeterministicResultHash string `protobuf:"bytes,3,opt,name=deterministic_result_hash,json=deterministicResultHash,proto3" json:"deterministic_result_hash,omitempty"`
	DeterministicResultRaw  string `protobuf:"bytes,4,opt,name=deterministic_result_raw,json=deterministicResultRaw,proto3" json:"deterministic_result_raw,omitempty"`
	ResultRaw               string `protobuf:"bytes,5,opt,name=result_raw,json=resultRaw,proto3" json:"result_raw,omitempty"`
	Elapsed                 string `protobuf:"bytes,6,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	NodeAddress             string `protobuf:"bytes,7,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty"`
}

func (x *CPUConsensusResult_Nodes) Reset() {
	*x = CPUConsensusResult_Nodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_grpc_schema_cpu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPUConsensusResult_Nodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPUConsensusResult_Nodes) ProtoMessage() {}

func (x *CPUConsensusResult_Nodes) ProtoReflect() protoreflect.Message {
	mi := &file_communication_grpc_schema_cpu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPUConsensusResult_Nodes.ProtoReflect.Descriptor instead.
func (*CPUConsensusResult_Nodes) Descriptor() ([]byte, []int) {
	return file_communication_grpc_schema_cpu_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CPUConsensusResult_Nodes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CPUConsensusResult_Nodes) GetResultHash() string {
	if x != nil {
		return x.ResultHash
	}
	return ""
}

func (x *CPUConsensusResult_Nodes) GetDeterministicResultHash() string {
	if x != nil {
		return x.DeterministicResultHash
	}
	return ""
}

func (x *CPUConsensusResult_Nodes) GetDeterministicResultRaw() string {
	if x != nil {
		return x.DeterministicResultRaw
	}
	return ""
}

func (x *CPUConsensusResult_Nodes) GetResultRaw() string {
	if x != nil {
		return x.ResultRaw
	}
	return ""
}

func (x *CPUConsensusResult_Nodes) GetElapsed() string {
	if x != nil {
		return x.Elapsed
	}
	return ""
}

func (x *CPUConsensusResult_Nodes) GetNodeAddress() string {
	if x != nil {
		return x.NodeAddress
	}
	return ""
}

type CPUConsensusResult_Verdict struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result          string                      `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	ConsensusStatus string                      `protobuf:"bytes,2,opt,name=consensus_status,json=consensusStatus,proto3" json:"consensus_status,omitempty"`
	NodeCount       int32                       `protobuf:"varint,3,opt,name=node_count,json=nodeCount,proto3" json:"node_count,omitempty"`
	Nodes           []*CPUConsensusResult_Nodes `protobuf:"bytes,4,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *CPUConsensusResult_Verdict) Reset() {
	*x = CPUConsensusResult_Verdict{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_grpc_schema_cpu_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPUConsensusResult_Verdict) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPUConsensusResult_Verdict) ProtoMessage() {}

func (x *CPUConsensusResult_Verdict) ProtoReflect() protoreflect.Message {
	mi := &file_communication_grpc_schema_cpu_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPUConsensusResult_Verdict.ProtoReflect.Descriptor instead.
func (*CPUConsensusResult_Verdict) Descriptor() ([]byte, []int) {
	return file_communication_grpc_schema_cpu_proto_rawDescGZIP(), []int{1, 1}
}

func (x *CPUConsensusResult_Verdict) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *CPUConsensusResult_Verdict) GetConsensusStatus() string {
	if x != nil {
		return x.ConsensusStatus
	}
	return ""
}

func (x *CPUConsensusResult_Verdict) GetNodeCount() int32 {
	if x != nil {
		return x.NodeCount
	}
	return 0
}

func (x *CPUConsensusResult_Verdict) GetNodes() []*CPUConsensusResult_Nodes {
	if x != nil {
		return x.Nodes
	}
	return nil
}

var File_communication_grpc_schema_cpu_proto protoreflect.FileDescriptor

var file_communication_grpc_schema_cpu_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x63, 0x70, 0x75, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x70, 0x75, 0x22, 0xd6, 0x01, 0x0a, 0x0a, 0x43,
	0x50, 0x55, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x65, 0x5f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x31, 0x0a, 0x14, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x22, 0xde, 0x04, 0x0a, 0x12, 0x43, 0x50, 0x55, 0x43, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x73, 0x75, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x64, 0x69, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x70, 0x75,
	0x2e, 0x43, 0x50, 0x55, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x2e, 0x56, 0x65, 0x72, 0x64, 0x69, 0x63, 0x74, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x64, 0x69, 0x63, 0x74, 0x1a, 0x8a, 0x02, 0x0a, 0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x3a, 0x0a, 0x19, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x17, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x38, 0x0a, 0x18, 0x64,
	0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x5f, 0x72, 0x61, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x64,
	0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x61, 0x77, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f,
	0x72, 0x61, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x61, 0x77, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x1a, 0xa0, 0x01, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x64, 0x69, 0x63, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73,
	0x75, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x33, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x63, 0x70, 0x75, 0x2e, 0x43, 0x50, 0x55, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x05, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x32, 0x48, 0x0a, 0x08, 0x43, 0x50, 0x55, 0x53, 0x68, 0x61, 0x72, 0x64,
	0x12, 0x3c, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0f, 0x2e, 0x63, 0x70, 0x75, 0x2e, 0x43, 0x50, 0x55, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x70, 0x75, 0x2e, 0x43, 0x50, 0x55, 0x43, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x30,
	0x5a, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x63, 0x70, 0x75,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_communication_grpc_schema_cpu_proto_rawDescOnce sync.Once
	file_communication_grpc_schema_cpu_proto_rawDescData = file_communication_grpc_schema_cpu_proto_rawDesc
)

func file_communication_grpc_schema_cpu_proto_rawDescGZIP() []byte {
	file_communication_grpc_schema_cpu_proto_rawDescOnce.Do(func() {
		file_communication_grpc_schema_cpu_proto_rawDescData = protoimpl.X.CompressGZIP(file_communication_grpc_schema_cpu_proto_rawDescData)
	})
	return file_communication_grpc_schema_cpu_proto_rawDescData
}

var file_communication_grpc_schema_cpu_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_communication_grpc_schema_cpu_proto_goTypes = []interface{}{
	(*CPURequest)(nil),                 // 0: cpu.CPURequest
	(*CPUConsensusResult)(nil),         // 1: cpu.CPUConsensusResult
	(*CPUConsensusResult_Nodes)(nil),   // 2: cpu.CPUConsensusResult.Nodes
	(*CPUConsensusResult_Verdict)(nil), // 3: cpu.CPUConsensusResult.Verdict
}
var file_communication_grpc_schema_cpu_proto_depIdxs = []int32{
	3, // 0: cpu.CPUConsensusResult.verdict:type_name -> cpu.CPUConsensusResult.Verdict
	2, // 1: cpu.CPUConsensusResult.Verdict.nodes:type_name -> cpu.CPUConsensusResult.Nodes
	0, // 2: cpu.CPUShard.ProcessRequest:input_type -> cpu.CPURequest
	1, // 3: cpu.CPUShard.ProcessRequest:output_type -> cpu.CPUConsensusResult
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_communication_grpc_schema_cpu_proto_init() }
func file_communication_grpc_schema_cpu_proto_init() {
	if File_communication_grpc_schema_cpu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_communication_grpc_schema_cpu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CPURequest); i {
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
		file_communication_grpc_schema_cpu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CPUConsensusResult); i {
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
		file_communication_grpc_schema_cpu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CPUConsensusResult_Nodes); i {
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
		file_communication_grpc_schema_cpu_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CPUConsensusResult_Verdict); i {
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
			RawDescriptor: file_communication_grpc_schema_cpu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_communication_grpc_schema_cpu_proto_goTypes,
		DependencyIndexes: file_communication_grpc_schema_cpu_proto_depIdxs,
		MessageInfos:      file_communication_grpc_schema_cpu_proto_msgTypes,
	}.Build()
	File_communication_grpc_schema_cpu_proto = out.File
	file_communication_grpc_schema_cpu_proto_rawDesc = nil
	file_communication_grpc_schema_cpu_proto_goTypes = nil
	file_communication_grpc_schema_cpu_proto_depIdxs = nil
}
