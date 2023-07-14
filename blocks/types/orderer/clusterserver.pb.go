// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderer/clusterserver.proto

package orderer

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/blcvn/lib-golang-test/blocks/types/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// ClusterNodeServiceStepRequest wraps a message that is sent to a cluster member.
type ClusterNodeServiceStepRequest struct {
	// Types that are valid to be assigned to Payload:
	//	*ClusterNodeServiceStepRequest_NodeConrequest
	//	*ClusterNodeServiceStepRequest_NodeTranrequest
	//	*ClusterNodeServiceStepRequest_NodeAuthrequest
	Payload              isClusterNodeServiceStepRequest_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *ClusterNodeServiceStepRequest) Reset()         { *m = ClusterNodeServiceStepRequest{} }
func (m *ClusterNodeServiceStepRequest) String() string { return proto.CompactTextString(m) }
func (*ClusterNodeServiceStepRequest) ProtoMessage()    {}
func (*ClusterNodeServiceStepRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce2573c2284b220f, []int{0}
}

func (m *ClusterNodeServiceStepRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterNodeServiceStepRequest.Unmarshal(m, b)
}
func (m *ClusterNodeServiceStepRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterNodeServiceStepRequest.Marshal(b, m, deterministic)
}
func (m *ClusterNodeServiceStepRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterNodeServiceStepRequest.Merge(m, src)
}
func (m *ClusterNodeServiceStepRequest) XXX_Size() int {
	return xxx_messageInfo_ClusterNodeServiceStepRequest.Size(m)
}
func (m *ClusterNodeServiceStepRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterNodeServiceStepRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterNodeServiceStepRequest proto.InternalMessageInfo

type isClusterNodeServiceStepRequest_Payload interface {
	isClusterNodeServiceStepRequest_Payload()
}

type ClusterNodeServiceStepRequest_NodeConrequest struct {
	NodeConrequest *NodeConsensusRequest `protobuf:"bytes,1,opt,name=node_conrequest,json=nodeConrequest,proto3,oneof"`
}

type ClusterNodeServiceStepRequest_NodeTranrequest struct {
	NodeTranrequest *NodeTransactionOrderRequest `protobuf:"bytes,2,opt,name=node_tranrequest,json=nodeTranrequest,proto3,oneof"`
}

type ClusterNodeServiceStepRequest_NodeAuthrequest struct {
	NodeAuthrequest *NodeAuthRequest `protobuf:"bytes,3,opt,name=node_authrequest,json=nodeAuthrequest,proto3,oneof"`
}

func (*ClusterNodeServiceStepRequest_NodeConrequest) isClusterNodeServiceStepRequest_Payload() {}

func (*ClusterNodeServiceStepRequest_NodeTranrequest) isClusterNodeServiceStepRequest_Payload() {}

func (*ClusterNodeServiceStepRequest_NodeAuthrequest) isClusterNodeServiceStepRequest_Payload() {}

func (m *ClusterNodeServiceStepRequest) GetPayload() isClusterNodeServiceStepRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ClusterNodeServiceStepRequest) GetNodeConrequest() *NodeConsensusRequest {
	if x, ok := m.GetPayload().(*ClusterNodeServiceStepRequest_NodeConrequest); ok {
		return x.NodeConrequest
	}
	return nil
}

func (m *ClusterNodeServiceStepRequest) GetNodeTranrequest() *NodeTransactionOrderRequest {
	if x, ok := m.GetPayload().(*ClusterNodeServiceStepRequest_NodeTranrequest); ok {
		return x.NodeTranrequest
	}
	return nil
}

func (m *ClusterNodeServiceStepRequest) GetNodeAuthrequest() *NodeAuthRequest {
	if x, ok := m.GetPayload().(*ClusterNodeServiceStepRequest_NodeAuthrequest); ok {
		return x.NodeAuthrequest
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ClusterNodeServiceStepRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ClusterNodeServiceStepRequest_NodeConrequest)(nil),
		(*ClusterNodeServiceStepRequest_NodeTranrequest)(nil),
		(*ClusterNodeServiceStepRequest_NodeAuthrequest)(nil),
	}
}

// ClusterNodeServiceStepResponse is a message received from a cluster member.
type ClusterNodeServiceStepResponse struct {
	// Types that are valid to be assigned to Payload:
	//	*ClusterNodeServiceStepResponse_TranorderRes
	Payload              isClusterNodeServiceStepResponse_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *ClusterNodeServiceStepResponse) Reset()         { *m = ClusterNodeServiceStepResponse{} }
func (m *ClusterNodeServiceStepResponse) String() string { return proto.CompactTextString(m) }
func (*ClusterNodeServiceStepResponse) ProtoMessage()    {}
func (*ClusterNodeServiceStepResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce2573c2284b220f, []int{1}
}

func (m *ClusterNodeServiceStepResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterNodeServiceStepResponse.Unmarshal(m, b)
}
func (m *ClusterNodeServiceStepResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterNodeServiceStepResponse.Marshal(b, m, deterministic)
}
func (m *ClusterNodeServiceStepResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterNodeServiceStepResponse.Merge(m, src)
}
func (m *ClusterNodeServiceStepResponse) XXX_Size() int {
	return xxx_messageInfo_ClusterNodeServiceStepResponse.Size(m)
}
func (m *ClusterNodeServiceStepResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterNodeServiceStepResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterNodeServiceStepResponse proto.InternalMessageInfo

type isClusterNodeServiceStepResponse_Payload interface {
	isClusterNodeServiceStepResponse_Payload()
}

type ClusterNodeServiceStepResponse_TranorderRes struct {
	TranorderRes *TransactionOrderResponse `protobuf:"bytes,1,opt,name=tranorder_res,json=tranorderRes,proto3,oneof"`
}

func (*ClusterNodeServiceStepResponse_TranorderRes) isClusterNodeServiceStepResponse_Payload() {}

func (m *ClusterNodeServiceStepResponse) GetPayload() isClusterNodeServiceStepResponse_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ClusterNodeServiceStepResponse) GetTranorderRes() *TransactionOrderResponse {
	if x, ok := m.GetPayload().(*ClusterNodeServiceStepResponse_TranorderRes); ok {
		return x.TranorderRes
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ClusterNodeServiceStepResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ClusterNodeServiceStepResponse_TranorderRes)(nil),
	}
}

// NodeConsensusRequest is a consensus specific message sent to a cluster member.
type NodeConsensusRequest struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Metadata             []byte   `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeConsensusRequest) Reset()         { *m = NodeConsensusRequest{} }
func (m *NodeConsensusRequest) String() string { return proto.CompactTextString(m) }
func (*NodeConsensusRequest) ProtoMessage()    {}
func (*NodeConsensusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce2573c2284b220f, []int{2}
}

func (m *NodeConsensusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeConsensusRequest.Unmarshal(m, b)
}
func (m *NodeConsensusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeConsensusRequest.Marshal(b, m, deterministic)
}
func (m *NodeConsensusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeConsensusRequest.Merge(m, src)
}
func (m *NodeConsensusRequest) XXX_Size() int {
	return xxx_messageInfo_NodeConsensusRequest.Size(m)
}
func (m *NodeConsensusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeConsensusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeConsensusRequest proto.InternalMessageInfo

func (m *NodeConsensusRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *NodeConsensusRequest) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// NodeTransactionOrderRequest wraps a transaction to be sent for ordering.
type NodeTransactionOrderRequest struct {
	// last_validation_seq denotes the last configuration sequence at which the
	// sender validated this message.
	LastValidationSeq uint64 `protobuf:"varint,1,opt,name=last_validation_seq,json=lastValidationSeq,proto3" json:"last_validation_seq,omitempty"`
	// content is the fabric transaction
	// that is forwarded to the cluster member.
	Payload              *common.Envelope `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *NodeTransactionOrderRequest) Reset()         { *m = NodeTransactionOrderRequest{} }
func (m *NodeTransactionOrderRequest) String() string { return proto.CompactTextString(m) }
func (*NodeTransactionOrderRequest) ProtoMessage()    {}
func (*NodeTransactionOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce2573c2284b220f, []int{3}
}

func (m *NodeTransactionOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeTransactionOrderRequest.Unmarshal(m, b)
}
func (m *NodeTransactionOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeTransactionOrderRequest.Marshal(b, m, deterministic)
}
func (m *NodeTransactionOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeTransactionOrderRequest.Merge(m, src)
}
func (m *NodeTransactionOrderRequest) XXX_Size() int {
	return xxx_messageInfo_NodeTransactionOrderRequest.Size(m)
}
func (m *NodeTransactionOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeTransactionOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeTransactionOrderRequest proto.InternalMessageInfo

func (m *NodeTransactionOrderRequest) GetLastValidationSeq() uint64 {
	if m != nil {
		return m.LastValidationSeq
	}
	return 0
}

func (m *NodeTransactionOrderRequest) GetPayload() *common.Envelope {
	if m != nil {
		return m.Payload
	}
	return nil
}

// TransactionOrderResponse returns a success
// or failure status to the sender.
type TransactionOrderResponse struct {
	Channel string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	TxId    string `protobuf:"bytes,2,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	// Status code, which may be used to programatically respond to success/failure.
	Status common.Status `protobuf:"varint,3,opt,name=status,proto3,enum=common.Status" json:"status,omitempty"`
	// Info string which may contain additional information about the returned status.
	Info                 string   `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionOrderResponse) Reset()         { *m = TransactionOrderResponse{} }
func (m *TransactionOrderResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionOrderResponse) ProtoMessage()    {}
func (*TransactionOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce2573c2284b220f, []int{4}
}

func (m *TransactionOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionOrderResponse.Unmarshal(m, b)
}
func (m *TransactionOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionOrderResponse.Marshal(b, m, deterministic)
}
func (m *TransactionOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionOrderResponse.Merge(m, src)
}
func (m *TransactionOrderResponse) XXX_Size() int {
	return xxx_messageInfo_TransactionOrderResponse.Size(m)
}
func (m *TransactionOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionOrderResponse proto.InternalMessageInfo

func (m *TransactionOrderResponse) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *TransactionOrderResponse) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *TransactionOrderResponse) GetStatus() common.Status {
	if m != nil {
		return m.Status
	}
	return common.Status_UNKNOWN
}

func (m *TransactionOrderResponse) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

// NodeAuthRequest for authenticate the stream
// between the cluster members
type NodeAuthRequest struct {
	// version represents the fields on which the signature is computed
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// signature is verifiable using the initiator's public key
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	// timestamp indicates the freshness of the request; expected to be within the margin
	// of the responsder's local time
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// from_id is the numerical identifier of the initiator of the connection
	FromId uint64 `protobuf:"varint,4,opt,name=from_id,json=fromId,proto3" json:"from_id,omitempty"`
	// to_id is the numerical identifier of the node that is being connected to
	ToId uint64 `protobuf:"varint,5,opt,name=to_id,json=toId,proto3" json:"to_id,omitempty"`
	// session_binding is verifiable using application level protocol
	SessionBinding       []byte   `protobuf:"bytes,6,opt,name=session_binding,json=sessionBinding,proto3" json:"session_binding,omitempty"`
	Channel              string   `protobuf:"bytes,7,opt,name=channel,proto3" json:"channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeAuthRequest) Reset()         { *m = NodeAuthRequest{} }
func (m *NodeAuthRequest) String() string { return proto.CompactTextString(m) }
func (*NodeAuthRequest) ProtoMessage()    {}
func (*NodeAuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce2573c2284b220f, []int{5}
}

func (m *NodeAuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeAuthRequest.Unmarshal(m, b)
}
func (m *NodeAuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeAuthRequest.Marshal(b, m, deterministic)
}
func (m *NodeAuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeAuthRequest.Merge(m, src)
}
func (m *NodeAuthRequest) XXX_Size() int {
	return xxx_messageInfo_NodeAuthRequest.Size(m)
}
func (m *NodeAuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeAuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeAuthRequest proto.InternalMessageInfo

func (m *NodeAuthRequest) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *NodeAuthRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *NodeAuthRequest) GetTimestamp() *timestamppb.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *NodeAuthRequest) GetFromId() uint64 {
	if m != nil {
		return m.FromId
	}
	return 0
}

func (m *NodeAuthRequest) GetToId() uint64 {
	if m != nil {
		return m.ToId
	}
	return 0
}

func (m *NodeAuthRequest) GetSessionBinding() []byte {
	if m != nil {
		return m.SessionBinding
	}
	return nil
}

func (m *NodeAuthRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func init() {
	proto.RegisterType((*ClusterNodeServiceStepRequest)(nil), "orderer.ClusterNodeServiceStepRequest")
	proto.RegisterType((*ClusterNodeServiceStepResponse)(nil), "orderer.ClusterNodeServiceStepResponse")
	proto.RegisterType((*NodeConsensusRequest)(nil), "orderer.NodeConsensusRequest")
	proto.RegisterType((*NodeTransactionOrderRequest)(nil), "orderer.NodeTransactionOrderRequest")
	proto.RegisterType((*TransactionOrderResponse)(nil), "orderer.TransactionOrderResponse")
	proto.RegisterType((*NodeAuthRequest)(nil), "orderer.NodeAuthRequest")
}

func init() { proto.RegisterFile("orderer/clusterserver.proto", fileDescriptor_ce2573c2284b220f) }

var fileDescriptor_ce2573c2284b220f = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x5d, 0x6f, 0xd3, 0x3c,
	0x14, 0x5e, 0xf7, 0x66, 0xed, 0x5b, 0xb3, 0x75, 0xc3, 0x43, 0xa2, 0xea, 0x18, 0x1f, 0x15, 0x6c,
	0x13, 0xd2, 0x12, 0x34, 0x6e, 0xb8, 0xdd, 0xa6, 0x49, 0x9d, 0x84, 0x40, 0xb8, 0x13, 0x42, 0x70,
	0x51, 0xb9, 0xc9, 0x69, 0x6a, 0x29, 0xb1, 0x53, 0xdb, 0xa9, 0xb6, 0x1f, 0xc0, 0x2d, 0xbf, 0x96,
	0x1f, 0x80, 0xfc, 0x91, 0x64, 0x1f, 0x6c, 0x5c, 0x25, 0xe7, 0x39, 0xe7, 0x3c, 0xcf, 0xf1, 0xe3,
	0x63, 0xb4, 0x23, 0x64, 0x02, 0x12, 0x64, 0x14, 0x67, 0xa5, 0xd2, 0x20, 0x15, 0xc8, 0x25, 0xc8,
	0xb0, 0x90, 0x42, 0x0b, 0xdc, 0xf1, 0xc9, 0xc1, 0x76, 0x2c, 0xf2, 0x5c, 0xf0, 0xc8, 0x7d, 0x5c,
	0x76, 0xf0, 0x22, 0x15, 0x22, 0xcd, 0x20, 0xb2, 0xd1, 0xb4, 0x9c, 0x45, 0x9a, 0xe5, 0xa0, 0x34,
	0xcd, 0x0b, 0x57, 0x30, 0xfc, 0xb5, 0x8a, 0x76, 0x4f, 0x1d, 0xed, 0x27, 0x91, 0xc0, 0x18, 0xe4,
	0x92, 0xc5, 0x30, 0xd6, 0x50, 0x10, 0x58, 0x94, 0xa0, 0x34, 0x1e, 0xa1, 0x4d, 0x2e, 0x12, 0x98,
	0xc4, 0x82, 0x4b, 0x07, 0xf5, 0x5b, 0x2f, 0x5b, 0x07, 0x8f, 0x8e, 0x76, 0x43, 0x2f, 0x1d, 0x9a,
	0xce, 0x53, 0xc1, 0x15, 0x70, 0x55, 0x2a, 0xdf, 0x37, 0x5a, 0x21, 0x3d, 0xee, 0x70, 0xdf, 0x86,
	0xbf, 0xa0, 0x2d, 0xcb, 0xa4, 0x25, 0xad, 0xa9, 0x56, 0x2d, 0xd5, 0xeb, 0x1b, 0x54, 0x17, 0x92,
	0x72, 0x45, 0x63, 0xcd, 0x04, 0xff, 0x6c, 0xe0, 0x86, 0xd1, 0x4e, 0x72, 0xd1, 0xb4, 0xe3, 0x33,
	0x4f, 0x49, 0x4b, 0x3d, 0xaf, 0x28, 0xff, 0xb3, 0x94, 0xfd, 0x1b, 0x94, 0xc7, 0xa5, 0x9e, 0xdf,
	0xa2, 0x39, 0x6e, 0x5a, 0x4e, 0xba, 0xa8, 0x53, 0xd0, 0xab, 0x4c, 0xd0, 0x64, 0x58, 0xa2, 0xe7,
	0xf7, 0xf9, 0xa1, 0x0a, 0x73, 0x46, 0x3c, 0x42, 0x1b, 0xe6, 0x04, 0x96, 0x7e, 0x22, 0x41, 0x79,
	0x3b, 0x5e, 0xd5, 0x82, 0x77, 0xe7, 0x77, 0x9d, 0xa3, 0x15, 0xb2, 0x5e, 0x77, 0x12, 0x50, 0xd7,
	0x65, 0x3f, 0xa2, 0x27, 0x7f, 0x73, 0x11, 0xf7, 0xeb, 0x12, 0x2b, 0xb3, 0x4e, 0xaa, 0x10, 0x0f,
	0xd0, 0xff, 0x39, 0x68, 0x9a, 0x50, 0x4d, 0xad, 0x8b, 0xeb, 0xa4, 0x8e, 0x87, 0x57, 0x68, 0xe7,
	0x01, 0x23, 0x71, 0x88, 0xb6, 0x33, 0xaa, 0xf4, 0x64, 0x49, 0x33, 0x96, 0x50, 0x93, 0x9e, 0x28,
	0x58, 0x58, 0x81, 0x80, 0x3c, 0x36, 0xa9, 0xaf, 0x75, 0x66, 0x0c, 0x0b, 0xfc, 0xb6, 0x19, 0xc2,
	0xdd, 0xd7, 0x56, 0xe8, 0xb7, 0xec, 0x8c, 0x2f, 0x21, 0x13, 0x05, 0xd4, 0x63, 0x0d, 0x7f, 0xb6,
	0x50, 0xff, 0x3e, 0x03, 0xcc, 0x69, 0xe2, 0x39, 0xe5, 0x1c, 0x32, 0x2b, 0xd6, 0x25, 0x55, 0x88,
	0xb7, 0xd1, 0x9a, 0xbe, 0x9c, 0x30, 0x27, 0xd0, 0x25, 0x81, 0xbe, 0x3c, 0x4f, 0xf0, 0x1e, 0x6a,
	0x2b, 0x4d, 0x75, 0xa9, 0xec, 0x9d, 0xf6, 0x8e, 0x7a, 0x95, 0xec, 0xd8, 0xa2, 0xc4, 0x67, 0x31,
	0x46, 0x01, 0xe3, 0x33, 0xd1, 0x0f, 0x5c, 0xaf, 0xf9, 0x1f, 0xfe, 0x6e, 0xa1, 0xcd, 0x5b, 0x37,
	0x6f, 0xe4, 0x97, 0x20, 0x15, 0x13, 0xdc, 0xca, 0x6f, 0x90, 0x2a, 0xc4, 0xcf, 0x50, 0x57, 0xb1,
	0x94, 0x53, 0x5d, 0x4a, 0xf0, 0x6e, 0x36, 0x00, 0xfe, 0x80, 0xba, 0xf5, 0xbb, 0xf1, 0xeb, 0x35,
	0x08, 0xdd, 0xcb, 0x0a, 0xab, 0x97, 0x15, 0x5e, 0x54, 0x15, 0xa4, 0x29, 0xc6, 0x4f, 0x51, 0x67,
	0x26, 0x45, 0x6e, 0x0e, 0x16, 0x58, 0x77, 0xdb, 0x26, 0x3c, 0x4f, 0xec, 0x79, 0x85, 0x81, 0xd7,
	0x2c, 0x1c, 0x68, 0x71, 0x9e, 0xe0, 0x7d, 0xb4, 0xa9, 0x40, 0x99, 0x81, 0x26, 0x53, 0xc6, 0x13,
	0xc6, 0xd3, 0x7e, 0xdb, 0xce, 0xd2, 0xf3, 0xf0, 0x89, 0x43, 0xaf, 0xfb, 0xd8, 0xb9, 0xe1, 0xe3,
	0xd1, 0x02, 0xe1, 0xbb, 0xeb, 0x8b, 0x7f, 0xa0, 0xc0, 0xac, 0x30, 0xde, 0xab, 0x77, 0xf4, 0xc1,
	0x37, 0x3f, 0xd8, 0xff, 0x67, 0x9d, 0xbb, 0xd0, 0x83, 0xd6, 0xbb, 0xd6, 0xc9, 0x37, 0xf4, 0x46,
	0xc8, 0x34, 0x9c, 0x5f, 0x15, 0x20, 0x33, 0x48, 0x52, 0x90, 0xe1, 0x8c, 0x4e, 0x25, 0x8b, 0x9d,
	0x37, 0xaa, 0xe2, 0xfa, 0x1e, 0xa5, 0x4c, 0xcf, 0xcb, 0xa9, 0xb9, 0xc4, 0xe8, 0x5a, 0x75, 0xe4,
	0xaa, 0x0f, 0x5d, 0xf5, 0x61, 0x2a, 0x22, 0xdf, 0x30, 0x6d, 0x5b, 0xe8, 0xfd, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd4, 0x82, 0x41, 0x47, 0x01, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClusterNodeServiceClient is the client API for ClusterNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClusterNodeServiceClient interface {
	// Step passes an implementation-specific message to another cluster member.
	Step(ctx context.Context, opts ...grpc.CallOption) (ClusterNodeService_StepClient, error)
}

type clusterNodeServiceClient struct {
	cc *grpc.ClientConn
}

func NewClusterNodeServiceClient(cc *grpc.ClientConn) ClusterNodeServiceClient {
	return &clusterNodeServiceClient{cc}
}

func (c *clusterNodeServiceClient) Step(ctx context.Context, opts ...grpc.CallOption) (ClusterNodeService_StepClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ClusterNodeService_serviceDesc.Streams[0], "/orderer.ClusterNodeService/Step", opts...)
	if err != nil {
		return nil, err
	}
	x := &clusterNodeServiceStepClient{stream}
	return x, nil
}

type ClusterNodeService_StepClient interface {
	Send(*ClusterNodeServiceStepRequest) error
	Recv() (*ClusterNodeServiceStepResponse, error)
	grpc.ClientStream
}

type clusterNodeServiceStepClient struct {
	grpc.ClientStream
}

func (x *clusterNodeServiceStepClient) Send(m *ClusterNodeServiceStepRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clusterNodeServiceStepClient) Recv() (*ClusterNodeServiceStepResponse, error) {
	m := new(ClusterNodeServiceStepResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClusterNodeServiceServer is the server API for ClusterNodeService service.
type ClusterNodeServiceServer interface {
	// Step passes an implementation-specific message to another cluster member.
	Step(ClusterNodeService_StepServer) error
}

// UnimplementedClusterNodeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClusterNodeServiceServer struct {
}

func (*UnimplementedClusterNodeServiceServer) Step(srv ClusterNodeService_StepServer) error {
	return status.Errorf(codes.Unimplemented, "method Step not implemented")
}

func RegisterClusterNodeServiceServer(s *grpc.Server, srv ClusterNodeServiceServer) {
	s.RegisterService(&_ClusterNodeService_serviceDesc, srv)
}

func _ClusterNodeService_Step_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClusterNodeServiceServer).Step(&clusterNodeServiceStepServer{stream})
}

type ClusterNodeService_StepServer interface {
	Send(*ClusterNodeServiceStepResponse) error
	Recv() (*ClusterNodeServiceStepRequest, error)
	grpc.ServerStream
}

type clusterNodeServiceStepServer struct {
	grpc.ServerStream
}

func (x *clusterNodeServiceStepServer) Send(m *ClusterNodeServiceStepResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clusterNodeServiceStepServer) Recv() (*ClusterNodeServiceStepRequest, error) {
	m := new(ClusterNodeServiceStepRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ClusterNodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "orderer.ClusterNodeService",
	HandlerType: (*ClusterNodeServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Step",
			Handler:       _ClusterNodeService_Step_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "orderer/clusterserver.proto",
}
