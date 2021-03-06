// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: proposal.proto

package protoutil

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ProposalType int32

const (
	ProposalType_Unknown_Proposal         ProposalType = 0
	ProposalType_Channel_AddPeerOrg       ProposalType = 1 // 通道加入peer组织
	ProposalType_Channel_RemovePeerOrg    ProposalType = 2 // 通道删除peer组织
	ProposalType_Channel_ConfigUpdate     ProposalType = 3 // 通道配置更新
	ProposalType_Consortium_AddPeerOrg    ProposalType = 4 // peer组织加入联盟
	ProposalType_Consortium_RemovePeerOrg ProposalType = 5 // 联盟移除peer组织
	ProposalType_Consortium_ConfigUpdate  ProposalType = 6 // 联盟配置更新
)

// Enum value maps for ProposalType.
var (
	ProposalType_name = map[int32]string{
		0: "Unknown_Proposal",
		1: "Channel_AddPeerOrg",
		2: "Channel_RemovePeerOrg",
		3: "Channel_ConfigUpdate",
		4: "Consortium_AddPeerOrg",
		5: "Consortium_RemovePeerOrg",
		6: "Consortium_ConfigUpdate",
	}
	ProposalType_value = map[string]int32{
		"Unknown_Proposal":         0,
		"Channel_AddPeerOrg":       1,
		"Channel_RemovePeerOrg":    2,
		"Channel_ConfigUpdate":     3,
		"Consortium_AddPeerOrg":    4,
		"Consortium_RemovePeerOrg": 5,
		"Consortium_ConfigUpdate":  6,
	}
)

func (x ProposalType) Enum() *ProposalType {
	p := new(ProposalType)
	*p = x
	return p
}

func (x ProposalType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProposalType) Descriptor() protoreflect.EnumDescriptor {
	return file_proposal_proto_enumTypes[0].Descriptor()
}

func (ProposalType) Type() protoreflect.EnumType {
	return &file_proposal_proto_enumTypes[0]
}

func (x ProposalType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProposalType.Descriptor instead.
func (ProposalType) EnumDescriptor() ([]byte, []int) {
	return file_proposal_proto_rawDescGZIP(), []int{0}
}

type Proposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ProposalType `protobuf:"varint,1,opt,name=type,proto3,enum=proposal.ProposalType" json:"type,omitempty"`
	// Types that are assignable to Content:
	//	*Proposal_NewOrg
	//	*Proposal_RemovedOrgName
	Content  isProposal_Content `protobuf_oneof:"content"`
	Deadline int64              `protobuf:"varint,4,opt,name=deadline,proto3" json:"deadline,omitempty"`
}

func (x *Proposal) Reset() {
	*x = Proposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proposal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proposal) ProtoMessage() {}

func (x *Proposal) ProtoReflect() protoreflect.Message {
	mi := &file_proposal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proposal.ProtoReflect.Descriptor instead.
func (*Proposal) Descriptor() ([]byte, []int) {
	return file_proposal_proto_rawDescGZIP(), []int{0}
}

func (x *Proposal) GetType() ProposalType {
	if x != nil {
		return x.Type
	}
	return ProposalType_Unknown_Proposal
}

func (m *Proposal) GetContent() isProposal_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *Proposal) GetNewOrg() *Organization {
	if x, ok := x.GetContent().(*Proposal_NewOrg); ok {
		return x.NewOrg
	}
	return nil
}

func (x *Proposal) GetRemovedOrgName() string {
	if x, ok := x.GetContent().(*Proposal_RemovedOrgName); ok {
		return x.RemovedOrgName
	}
	return ""
}

func (x *Proposal) GetDeadline() int64 {
	if x != nil {
		return x.Deadline
	}
	return 0
}

type isProposal_Content interface {
	isProposal_Content()
}

type Proposal_NewOrg struct {
	NewOrg *Organization `protobuf:"bytes,2,opt,name=new_org,json=newOrg,proto3,oneof"`
}

type Proposal_RemovedOrgName struct {
	RemovedOrgName string `protobuf:"bytes,3,opt,name=removed_org_name,json=removedOrgName,proto3,oneof"`
}

func (*Proposal_NewOrg) isProposal_Content() {}

func (*Proposal_RemovedOrgName) isProposal_Content() {}

type ProposalSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProposalHash []byte `protobuf:"bytes,1,opt,name=proposal_hash,json=proposalHash,proto3" json:"proposal_hash,omitempty"`
	Creator      string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Signature    []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ProposalSignature) Reset() {
	*x = ProposalSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proposal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposalSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposalSignature) ProtoMessage() {}

func (x *ProposalSignature) ProtoReflect() protoreflect.Message {
	mi := &file_proposal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposalSignature.ProtoReflect.Descriptor instead.
func (*ProposalSignature) Descriptor() ([]byte, []int) {
	return file_proposal_proto_rawDescGZIP(), []int{1}
}

func (x *ProposalSignature) GetProposalHash() []byte {
	if x != nil {
		return x.ProposalHash
	}
	return nil
}

func (x *ProposalSignature) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *ProposalSignature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type ProposalEnvelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProposalId []byte             `protobuf:"bytes,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	Proposal   []byte             `protobuf:"bytes,2,opt,name=proposal,proto3" json:"proposal,omitempty"`
	Sign       *ProposalSignature `protobuf:"bytes,3,opt,name=sign,proto3" json:"sign,omitempty"`
	ChannelId  string             `protobuf:"bytes,4,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (x *ProposalEnvelope) Reset() {
	*x = ProposalEnvelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proposal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposalEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposalEnvelope) ProtoMessage() {}

func (x *ProposalEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_proposal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposalEnvelope.ProtoReflect.Descriptor instead.
func (*ProposalEnvelope) Descriptor() ([]byte, []int) {
	return file_proposal_proto_rawDescGZIP(), []int{2}
}

func (x *ProposalEnvelope) GetProposalId() []byte {
	if x != nil {
		return x.ProposalId
	}
	return nil
}

func (x *ProposalEnvelope) GetProposal() []byte {
	if x != nil {
		return x.Proposal
	}
	return nil
}

func (x *ProposalEnvelope) GetSign() *ProposalSignature {
	if x != nil {
		return x.Sign
	}
	return nil
}

func (x *ProposalEnvelope) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

// 发起提案请求
type ProposalInitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signer         *Signer   `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	Orderer        *Orderer  `protobuf:"bytes,2,opt,name=orderer,proto3" json:"orderer,omitempty"`
	ChannelId      string    `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	ConsortiumName string    `protobuf:"bytes,4,opt,name=consortium_name,json=consortiumName,proto3" json:"consortium_name,omitempty"`
	Proposal       *Proposal `protobuf:"bytes,5,opt,name=proposal,proto3" json:"proposal,omitempty"`
}

func (x *ProposalInitRequest) Reset() {
	*x = ProposalInitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proposal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposalInitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposalInitRequest) ProtoMessage() {}

func (x *ProposalInitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proposal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposalInitRequest.ProtoReflect.Descriptor instead.
func (*ProposalInitRequest) Descriptor() ([]byte, []int) {
	return file_proposal_proto_rawDescGZIP(), []int{3}
}

func (x *ProposalInitRequest) GetSigner() *Signer {
	if x != nil {
		return x.Signer
	}
	return nil
}

func (x *ProposalInitRequest) GetOrderer() *Orderer {
	if x != nil {
		return x.Orderer
	}
	return nil
}

func (x *ProposalInitRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *ProposalInitRequest) GetConsortiumName() string {
	if x != nil {
		return x.ConsortiumName
	}
	return ""
}

func (x *ProposalInitRequest) GetProposal() *Proposal {
	if x != nil {
		return x.Proposal
	}
	return nil
}

type ProposalSignRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signer   *Signer           `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	Envelope *ProposalEnvelope `protobuf:"bytes,2,opt,name=envelope,proto3" json:"envelope,omitempty"`
}

func (x *ProposalSignRequest) Reset() {
	*x = ProposalSignRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proposal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposalSignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposalSignRequest) ProtoMessage() {}

func (x *ProposalSignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proposal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposalSignRequest.ProtoReflect.Descriptor instead.
func (*ProposalSignRequest) Descriptor() ([]byte, []int) {
	return file_proposal_proto_rawDescGZIP(), []int{4}
}

func (x *ProposalSignRequest) GetSigner() *Signer {
	if x != nil {
		return x.Signer
	}
	return nil
}

func (x *ProposalSignRequest) GetEnvelope() *ProposalEnvelope {
	if x != nil {
		return x.Envelope
	}
	return nil
}

type ProposalSubmitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signer   *Signer              `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	Orderer  *Orderer             `protobuf:"bytes,2,opt,name=orderer,proto3" json:"orderer,omitempty"`
	Envelope *ProposalEnvelope    `protobuf:"bytes,3,opt,name=envelope,proto3" json:"envelope,omitempty"`
	Sigs     []*ProposalSignature `protobuf:"bytes,4,rep,name=sigs,proto3" json:"sigs,omitempty"`
}

func (x *ProposalSubmitRequest) Reset() {
	*x = ProposalSubmitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proposal_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposalSubmitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposalSubmitRequest) ProtoMessage() {}

func (x *ProposalSubmitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proposal_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposalSubmitRequest.ProtoReflect.Descriptor instead.
func (*ProposalSubmitRequest) Descriptor() ([]byte, []int) {
	return file_proposal_proto_rawDescGZIP(), []int{5}
}

func (x *ProposalSubmitRequest) GetSigner() *Signer {
	if x != nil {
		return x.Signer
	}
	return nil
}

func (x *ProposalSubmitRequest) GetOrderer() *Orderer {
	if x != nil {
		return x.Orderer
	}
	return nil
}

func (x *ProposalSubmitRequest) GetEnvelope() *ProposalEnvelope {
	if x != nil {
		return x.Envelope
	}
	return nil
}

func (x *ProposalSubmitRequest) GetSigs() []*ProposalSignature {
	if x != nil {
		return x.Sigs
	}
	return nil
}

var File_proposal_proto protoreflect.FileDescriptor

var file_proposal_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2e, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x2f, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x5f, 0x6f, 0x72, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x4f,
	0x72, 0x67, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x6f, 0x72,
	0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0e,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x4f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x70, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x9f, 0x01, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x04, 0x73, 0x69, 0x67,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22, 0xe0, 0x01, 0x0a, 0x13, 0x50, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x72, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x75,
	0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f,
	0x6e, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x75, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x22, 0x75, 0x0a, 0x13,
	0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x65, 0x72, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x08, 0x65,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52, 0x08, 0x65, 0x6e, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x65, 0x22, 0xd3, 0x01, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x52, 0x06, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72,
	0x12, 0x36, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2e, 0x50, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52, 0x08,
	0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x04, 0x73, 0x69, 0x67, 0x73, 0x2a, 0xc7, 0x01, 0x0a, 0x0c, 0x50, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x10, 0x00,
	0x12, 0x16, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x41, 0x64, 0x64, 0x50,
	0x65, 0x65, 0x72, 0x4f, 0x72, 0x67, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x5f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x65, 0x72, 0x4f, 0x72,
	0x67, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10, 0x03, 0x12, 0x19, 0x0a,
	0x15, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x75, 0x6d, 0x5f, 0x41, 0x64, 0x64, 0x50,
	0x65, 0x65, 0x72, 0x4f, 0x72, 0x67, 0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x6f, 0x6e, 0x73,
	0x6f, 0x72, 0x74, 0x69, 0x75, 0x6d, 0x5f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x65,
	0x72, 0x4f, 0x72, 0x67, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x72,
	0x74, 0x69, 0x75, 0x6d, 0x5f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x10, 0x06, 0x32, 0xdc, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x53, 0x74, 0x75, 0x62, 0x12, 0x47, 0x0a, 0x08, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65,
	0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a,
	0x04, 0x53, 0x69, 0x67, 0x6e, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2e,
	0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x2e,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x75, 0x74, 0x69, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proposal_proto_rawDescOnce sync.Once
	file_proposal_proto_rawDescData = file_proposal_proto_rawDesc
)

func file_proposal_proto_rawDescGZIP() []byte {
	file_proposal_proto_rawDescOnce.Do(func() {
		file_proposal_proto_rawDescData = protoimpl.X.CompressGZIP(file_proposal_proto_rawDescData)
	})
	return file_proposal_proto_rawDescData
}

var file_proposal_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proposal_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proposal_proto_goTypes = []interface{}{
	(ProposalType)(0),             // 0: proposal.ProposalType
	(*Proposal)(nil),              // 1: proposal.Proposal
	(*ProposalSignature)(nil),     // 2: proposal.ProposalSignature
	(*ProposalEnvelope)(nil),      // 3: proposal.ProposalEnvelope
	(*ProposalInitRequest)(nil),   // 4: proposal.ProposalInitRequest
	(*ProposalSignRequest)(nil),   // 5: proposal.ProposalSignRequest
	(*ProposalSubmitRequest)(nil), // 6: proposal.ProposalSubmitRequest
	(*Organization)(nil),          // 7: common.Organization
	(*Signer)(nil),                // 8: common.Signer
	(*Orderer)(nil),               // 9: common.Orderer
	(*Response)(nil),              // 10: common.Response
}
var file_proposal_proto_depIdxs = []int32{
	0,  // 0: proposal.Proposal.type:type_name -> proposal.ProposalType
	7,  // 1: proposal.Proposal.new_org:type_name -> common.Organization
	2,  // 2: proposal.ProposalEnvelope.sign:type_name -> proposal.ProposalSignature
	8,  // 3: proposal.ProposalInitRequest.signer:type_name -> common.Signer
	9,  // 4: proposal.ProposalInitRequest.orderer:type_name -> common.Orderer
	1,  // 5: proposal.ProposalInitRequest.proposal:type_name -> proposal.Proposal
	8,  // 6: proposal.ProposalSignRequest.signer:type_name -> common.Signer
	3,  // 7: proposal.ProposalSignRequest.envelope:type_name -> proposal.ProposalEnvelope
	8,  // 8: proposal.ProposalSubmitRequest.signer:type_name -> common.Signer
	9,  // 9: proposal.ProposalSubmitRequest.orderer:type_name -> common.Orderer
	3,  // 10: proposal.ProposalSubmitRequest.envelope:type_name -> proposal.ProposalEnvelope
	2,  // 11: proposal.ProposalSubmitRequest.sigs:type_name -> proposal.ProposalSignature
	4,  // 12: proposal.ProposalStub.Initiate:input_type -> proposal.ProposalInitRequest
	5,  // 13: proposal.ProposalStub.Sign:input_type -> proposal.ProposalSignRequest
	6,  // 14: proposal.ProposalStub.Submit:input_type -> proposal.ProposalSubmitRequest
	3,  // 15: proposal.ProposalStub.Initiate:output_type -> proposal.ProposalEnvelope
	2,  // 16: proposal.ProposalStub.Sign:output_type -> proposal.ProposalSignature
	10, // 17: proposal.ProposalStub.Submit:output_type -> common.Response
	15, // [15:18] is the sub-list for method output_type
	12, // [12:15] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_proposal_proto_init() }
func file_proposal_proto_init() {
	if File_proposal_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proposal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proposal); i {
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
		file_proposal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposalSignature); i {
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
		file_proposal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposalEnvelope); i {
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
		file_proposal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposalInitRequest); i {
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
		file_proposal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposalSignRequest); i {
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
		file_proposal_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposalSubmitRequest); i {
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
	file_proposal_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Proposal_NewOrg)(nil),
		(*Proposal_RemovedOrgName)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proposal_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proposal_proto_goTypes,
		DependencyIndexes: file_proposal_proto_depIdxs,
		EnumInfos:         file_proposal_proto_enumTypes,
		MessageInfos:      file_proposal_proto_msgTypes,
	}.Build()
	File_proposal_proto = out.File
	file_proposal_proto_rawDesc = nil
	file_proposal_proto_goTypes = nil
	file_proposal_proto_depIdxs = nil
}
