// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: authorization/authorization.proto

package authorization

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

type CheckAccessResponse_Result int32

const (
	CheckAccessResponse_RESULT_ALLOWED CheckAccessResponse_Result = 0
	CheckAccessResponse_RESULT_DENIED  CheckAccessResponse_Result = 1
)

// Enum value maps for CheckAccessResponse_Result.
var (
	CheckAccessResponse_Result_name = map[int32]string{
		0: "RESULT_ALLOWED",
		1: "RESULT_DENIED",
	}
	CheckAccessResponse_Result_value = map[string]int32{
		"RESULT_ALLOWED": 0,
		"RESULT_DENIED":  1,
	}
)

func (x CheckAccessResponse_Result) Enum() *CheckAccessResponse_Result {
	p := new(CheckAccessResponse_Result)
	*p = x
	return p
}

func (x CheckAccessResponse_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CheckAccessResponse_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_authorization_authorization_proto_enumTypes[0].Descriptor()
}

func (CheckAccessResponse_Result) Type() protoreflect.EnumType {
	return &file_authorization_authorization_proto_enumTypes[0]
}

func (x CheckAccessResponse_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CheckAccessResponse_Result.Descriptor instead.
func (CheckAccessResponse_Result) EnumDescriptor() ([]byte, []int) {
	return file_authorization_authorization_proto_rawDescGZIP(), []int{3, 0}
}

type Relationship struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// relation is the name of the relationship between two resources.
	Relation string `protobuf:"bytes,1,opt,name=relation,proto3" json:"relation,omitempty"`
	// subject_id is the ID of the subject (i.e., "other end") of the relationship.
	SubjectId string `protobuf:"bytes,2,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
}

func (x *Relationship) Reset() {
	*x = Relationship{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorization_authorization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Relationship) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relationship) ProtoMessage() {}

func (x *Relationship) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_authorization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Relationship.ProtoReflect.Descriptor instead.
func (*Relationship) Descriptor() ([]byte, []int) {
	return file_authorization_authorization_proto_rawDescGZIP(), []int{0}
}

func (x *Relationship) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *Relationship) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

type AccessRequestAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// action is the name of the action the subject is attempting to perform an action on.
	Action string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	// resource_id is the ID of the resource the subject is attempting to perform an action on.
	ResourceId string `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *AccessRequestAction) Reset() {
	*x = AccessRequestAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorization_authorization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessRequestAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessRequestAction) ProtoMessage() {}

func (x *AccessRequestAction) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_authorization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessRequestAction.ProtoReflect.Descriptor instead.
func (*AccessRequestAction) Descriptor() ([]byte, []int) {
	return file_authorization_authorization_proto_rawDescGZIP(), []int{1}
}

func (x *AccessRequestAction) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *AccessRequestAction) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type CheckAccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// credential is the literal credential for a subject (such as a bearer token) passed to the
	// application with no transformations applied.
	Credential string `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
	// actions is the set of all actions to check access for. All of these must be allowed for the
	// request itself to be allowed.
	Actions []*AccessRequestAction `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *CheckAccessRequest) Reset() {
	*x = CheckAccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorization_authorization_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAccessRequest) ProtoMessage() {}

func (x *CheckAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_authorization_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAccessRequest.ProtoReflect.Descriptor instead.
func (*CheckAccessRequest) Descriptor() ([]byte, []int) {
	return file_authorization_authorization_proto_rawDescGZIP(), []int{2}
}

func (x *CheckAccessRequest) GetCredential() string {
	if x != nil {
		return x.Credential
	}
	return ""
}

func (x *CheckAccessRequest) GetActions() []*AccessRequestAction {
	if x != nil {
		return x.Actions
	}
	return nil
}

type CheckAccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result CheckAccessResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=runtime.iam.v1.CheckAccessResponse_Result" json:"result,omitempty"`
}

func (x *CheckAccessResponse) Reset() {
	*x = CheckAccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorization_authorization_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAccessResponse) ProtoMessage() {}

func (x *CheckAccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_authorization_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAccessResponse.ProtoReflect.Descriptor instead.
func (*CheckAccessResponse) Descriptor() ([]byte, []int) {
	return file_authorization_authorization_proto_rawDescGZIP(), []int{3}
}

func (x *CheckAccessResponse) GetResult() CheckAccessResponse_Result {
	if x != nil {
		return x.Result
	}
	return CheckAccessResponse_RESULT_ALLOWED
}

type CreateRelationshipsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource_id is the ID of the resource to create relationships for.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// relationships is the set of relationships to create.
	Relationships []*Relationship `protobuf:"bytes,2,rep,name=relationships,proto3" json:"relationships,omitempty"`
}

func (x *CreateRelationshipsRequest) Reset() {
	*x = CreateRelationshipsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorization_authorization_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRelationshipsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRelationshipsRequest) ProtoMessage() {}

func (x *CreateRelationshipsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_authorization_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRelationshipsRequest.ProtoReflect.Descriptor instead.
func (*CreateRelationshipsRequest) Descriptor() ([]byte, []int) {
	return file_authorization_authorization_proto_rawDescGZIP(), []int{4}
}

func (x *CreateRelationshipsRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *CreateRelationshipsRequest) GetRelationships() []*Relationship {
	if x != nil {
		return x.Relationships
	}
	return nil
}

type CreateRelationshipsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateRelationshipsResponse) Reset() {
	*x = CreateRelationshipsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorization_authorization_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRelationshipsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRelationshipsResponse) ProtoMessage() {}

func (x *CreateRelationshipsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_authorization_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRelationshipsResponse.ProtoReflect.Descriptor instead.
func (*CreateRelationshipsResponse) Descriptor() ([]byte, []int) {
	return file_authorization_authorization_proto_rawDescGZIP(), []int{5}
}

type DeleteRelationshipsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource_id is the ID of the resource to delete relationships for.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// relationships is the set of relationships to delete.
	Relationships []*Relationship `protobuf:"bytes,2,rep,name=relationships,proto3" json:"relationships,omitempty"`
}

func (x *DeleteRelationshipsRequest) Reset() {
	*x = DeleteRelationshipsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorization_authorization_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRelationshipsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRelationshipsRequest) ProtoMessage() {}

func (x *DeleteRelationshipsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_authorization_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRelationshipsRequest.ProtoReflect.Descriptor instead.
func (*DeleteRelationshipsRequest) Descriptor() ([]byte, []int) {
	return file_authorization_authorization_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteRelationshipsRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *DeleteRelationshipsRequest) GetRelationships() []*Relationship {
	if x != nil {
		return x.Relationships
	}
	return nil
}

type DeleteRelationshipsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRelationshipsResponse) Reset() {
	*x = DeleteRelationshipsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorization_authorization_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRelationshipsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRelationshipsResponse) ProtoMessage() {}

func (x *DeleteRelationshipsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_authorization_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRelationshipsResponse.ProtoReflect.Descriptor instead.
func (*DeleteRelationshipsResponse) Descriptor() ([]byte, []int) {
	return file_authorization_authorization_proto_rawDescGZIP(), []int{7}
}

var File_authorization_authorization_proto protoreflect.FileDescriptor

var file_authorization_authorization_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x22, 0x49, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x68, 0x69, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x4e,
	0x0a, 0x13, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x73,
	0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x12, 0x3d, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x2f, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0d, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x01,
	0x22, 0x81, 0x01, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x42, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x68, 0x69, 0x70, 0x73, 0x22, 0x1d, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x68, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x22, 0x1d, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xcd, 0x02, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x58, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x22, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x70, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x2a, 0x2e, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x2a, 0x2e, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x62,
	0x6f, 0x78, 0x2f, 0x69, 0x61, 0x6d, 0x2d, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authorization_authorization_proto_rawDescOnce sync.Once
	file_authorization_authorization_proto_rawDescData = file_authorization_authorization_proto_rawDesc
)

func file_authorization_authorization_proto_rawDescGZIP() []byte {
	file_authorization_authorization_proto_rawDescOnce.Do(func() {
		file_authorization_authorization_proto_rawDescData = protoimpl.X.CompressGZIP(file_authorization_authorization_proto_rawDescData)
	})
	return file_authorization_authorization_proto_rawDescData
}

var file_authorization_authorization_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_authorization_authorization_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_authorization_authorization_proto_goTypes = []interface{}{
	(CheckAccessResponse_Result)(0),     // 0: runtime.iam.v1.CheckAccessResponse.Result
	(*Relationship)(nil),                // 1: runtime.iam.v1.Relationship
	(*AccessRequestAction)(nil),         // 2: runtime.iam.v1.AccessRequestAction
	(*CheckAccessRequest)(nil),          // 3: runtime.iam.v1.CheckAccessRequest
	(*CheckAccessResponse)(nil),         // 4: runtime.iam.v1.CheckAccessResponse
	(*CreateRelationshipsRequest)(nil),  // 5: runtime.iam.v1.CreateRelationshipsRequest
	(*CreateRelationshipsResponse)(nil), // 6: runtime.iam.v1.CreateRelationshipsResponse
	(*DeleteRelationshipsRequest)(nil),  // 7: runtime.iam.v1.DeleteRelationshipsRequest
	(*DeleteRelationshipsResponse)(nil), // 8: runtime.iam.v1.DeleteRelationshipsResponse
}
var file_authorization_authorization_proto_depIdxs = []int32{
	2, // 0: runtime.iam.v1.CheckAccessRequest.actions:type_name -> runtime.iam.v1.AccessRequestAction
	0, // 1: runtime.iam.v1.CheckAccessResponse.result:type_name -> runtime.iam.v1.CheckAccessResponse.Result
	1, // 2: runtime.iam.v1.CreateRelationshipsRequest.relationships:type_name -> runtime.iam.v1.Relationship
	1, // 3: runtime.iam.v1.DeleteRelationshipsRequest.relationships:type_name -> runtime.iam.v1.Relationship
	3, // 4: runtime.iam.v1.Authorization.CheckAccess:input_type -> runtime.iam.v1.CheckAccessRequest
	5, // 5: runtime.iam.v1.Authorization.CreateRelationships:input_type -> runtime.iam.v1.CreateRelationshipsRequest
	7, // 6: runtime.iam.v1.Authorization.DeleteRelationships:input_type -> runtime.iam.v1.DeleteRelationshipsRequest
	4, // 7: runtime.iam.v1.Authorization.CheckAccess:output_type -> runtime.iam.v1.CheckAccessResponse
	6, // 8: runtime.iam.v1.Authorization.CreateRelationships:output_type -> runtime.iam.v1.CreateRelationshipsResponse
	8, // 9: runtime.iam.v1.Authorization.DeleteRelationships:output_type -> runtime.iam.v1.DeleteRelationshipsResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_authorization_authorization_proto_init() }
func file_authorization_authorization_proto_init() {
	if File_authorization_authorization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authorization_authorization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Relationship); i {
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
		file_authorization_authorization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessRequestAction); i {
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
		file_authorization_authorization_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAccessRequest); i {
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
		file_authorization_authorization_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAccessResponse); i {
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
		file_authorization_authorization_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRelationshipsRequest); i {
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
		file_authorization_authorization_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRelationshipsResponse); i {
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
		file_authorization_authorization_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRelationshipsRequest); i {
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
		file_authorization_authorization_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRelationshipsResponse); i {
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
			RawDescriptor: file_authorization_authorization_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authorization_authorization_proto_goTypes,
		DependencyIndexes: file_authorization_authorization_proto_depIdxs,
		EnumInfos:         file_authorization_authorization_proto_enumTypes,
		MessageInfos:      file_authorization_authorization_proto_msgTypes,
	}.Build()
	File_authorization_authorization_proto = out.File
	file_authorization_authorization_proto_rawDesc = nil
	file_authorization_authorization_proto_goTypes = nil
	file_authorization_authorization_proto_depIdxs = nil
}
