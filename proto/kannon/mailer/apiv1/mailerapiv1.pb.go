// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        (unknown)
// source: kannon/mailer/apiv1/mailerapiv1.proto

package apiv1

import (
	types "github.com/ludusrusso/kannon/proto/kannon/mailer/types"
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

type SendHTMLReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender        *types.Sender          `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Subject       string                 `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Html          string                 `protobuf:"bytes,4,opt,name=html,proto3" json:"html,omitempty"`
	ScheduledTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=scheduled_time,json=scheduledTime,proto3,oneof" json:"scheduled_time,omitempty"`
	Recipients    []*types.Recipient     `protobuf:"bytes,6,rep,name=recipients,proto3" json:"recipients,omitempty"`
}

func (x *SendHTMLReq) Reset() {
	*x = SendHTMLReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kannon_mailer_apiv1_mailerapiv1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendHTMLReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendHTMLReq) ProtoMessage() {}

func (x *SendHTMLReq) ProtoReflect() protoreflect.Message {
	mi := &file_kannon_mailer_apiv1_mailerapiv1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendHTMLReq.ProtoReflect.Descriptor instead.
func (*SendHTMLReq) Descriptor() ([]byte, []int) {
	return file_kannon_mailer_apiv1_mailerapiv1_proto_rawDescGZIP(), []int{0}
}

func (x *SendHTMLReq) GetSender() *types.Sender {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *SendHTMLReq) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SendHTMLReq) GetHtml() string {
	if x != nil {
		return x.Html
	}
	return ""
}

func (x *SendHTMLReq) GetScheduledTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduledTime
	}
	return nil
}

func (x *SendHTMLReq) GetRecipients() []*types.Recipient {
	if x != nil {
		return x.Recipients
	}
	return nil
}

type SendTemplateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender        *types.Sender          `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Subject       string                 `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	TemplateId    string                 `protobuf:"bytes,4,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	ScheduledTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=scheduled_time,json=scheduledTime,proto3,oneof" json:"scheduled_time,omitempty"`
	Recipients    []*types.Recipient     `protobuf:"bytes,6,rep,name=recipients,proto3" json:"recipients,omitempty"`
}

func (x *SendTemplateReq) Reset() {
	*x = SendTemplateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kannon_mailer_apiv1_mailerapiv1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTemplateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTemplateReq) ProtoMessage() {}

func (x *SendTemplateReq) ProtoReflect() protoreflect.Message {
	mi := &file_kannon_mailer_apiv1_mailerapiv1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTemplateReq.ProtoReflect.Descriptor instead.
func (*SendTemplateReq) Descriptor() ([]byte, []int) {
	return file_kannon_mailer_apiv1_mailerapiv1_proto_rawDescGZIP(), []int{1}
}

func (x *SendTemplateReq) GetSender() *types.Sender {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *SendTemplateReq) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SendTemplateReq) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *SendTemplateReq) GetScheduledTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduledTime
	}
	return nil
}

func (x *SendTemplateReq) GetRecipients() []*types.Recipient {
	if x != nil {
		return x.Recipients
	}
	return nil
}

type SendRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId     string                 `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	TemplateId    string                 `protobuf:"bytes,2,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	ScheduledTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=scheduled_time,json=scheduledTime,proto3" json:"scheduled_time,omitempty"`
}

func (x *SendRes) Reset() {
	*x = SendRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kannon_mailer_apiv1_mailerapiv1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRes) ProtoMessage() {}

func (x *SendRes) ProtoReflect() protoreflect.Message {
	mi := &file_kannon_mailer_apiv1_mailerapiv1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRes.ProtoReflect.Descriptor instead.
func (*SendRes) Descriptor() ([]byte, []int) {
	return file_kannon_mailer_apiv1_mailerapiv1_proto_rawDescGZIP(), []int{2}
}

func (x *SendRes) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *SendRes) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *SendRes) GetScheduledTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduledTime
	}
	return nil
}

var File_kannon_mailer_apiv1_mailerapiv1_proto protoreflect.FileDescriptor

var file_kannon_mailer_apiv1_mailerapiv1_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x70, 0x6b, 0x67, 0x2e, 0x6b, 0x61, 0x6e,
	0x6e, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x93, 0x02, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x48, 0x54, 0x4d, 0x4c, 0x52, 0x65,
	0x71, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2e, 0x6d,
	0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x12, 0x46, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x0d,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x42, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x6b, 0x61, 0x6e, 0x6e, 0x6f,
	0x6e, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xa4, 0x02, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x37, 0x0a, 0x06, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12,
	0x46, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x11, 0x0a, 0x0f, 0x5f,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x8c,
	0x01, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0e, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xbc, 0x01,
	0x0a, 0x06, 0x4d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x12, 0x54, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64,
	0x48, 0x54, 0x4d, 0x4c, 0x12, 0x24, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x6b, 0x61, 0x6e, 0x6e, 0x6f,
	0x6e, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x48, 0x54, 0x4d, 0x4c, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x70, 0x6b, 0x67,
	0x2e, 0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x5c,
	0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x28,
	0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x6b,
	0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x75, 0x64, 0x75, 0x73,
	0x72, 0x75, 0x73, 0x73, 0x6f, 0x2f, 0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kannon_mailer_apiv1_mailerapiv1_proto_rawDescOnce sync.Once
	file_kannon_mailer_apiv1_mailerapiv1_proto_rawDescData = file_kannon_mailer_apiv1_mailerapiv1_proto_rawDesc
)

func file_kannon_mailer_apiv1_mailerapiv1_proto_rawDescGZIP() []byte {
	file_kannon_mailer_apiv1_mailerapiv1_proto_rawDescOnce.Do(func() {
		file_kannon_mailer_apiv1_mailerapiv1_proto_rawDescData = protoimpl.X.CompressGZIP(file_kannon_mailer_apiv1_mailerapiv1_proto_rawDescData)
	})
	return file_kannon_mailer_apiv1_mailerapiv1_proto_rawDescData
}

var file_kannon_mailer_apiv1_mailerapiv1_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kannon_mailer_apiv1_mailerapiv1_proto_goTypes = []interface{}{
	(*SendHTMLReq)(nil),           // 0: pkg.kannon.mailer.apiv1.SendHTMLReq
	(*SendTemplateReq)(nil),       // 1: pkg.kannon.mailer.apiv1.SendTemplateReq
	(*SendRes)(nil),               // 2: pkg.kannon.mailer.apiv1.SendRes
	(*types.Sender)(nil),          // 3: pkg.kannon.mailer.types.Sender
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*types.Recipient)(nil),       // 5: pkg.kannon.mailer.types.Recipient
}
var file_kannon_mailer_apiv1_mailerapiv1_proto_depIdxs = []int32{
	3, // 0: pkg.kannon.mailer.apiv1.SendHTMLReq.sender:type_name -> pkg.kannon.mailer.types.Sender
	4, // 1: pkg.kannon.mailer.apiv1.SendHTMLReq.scheduled_time:type_name -> google.protobuf.Timestamp
	5, // 2: pkg.kannon.mailer.apiv1.SendHTMLReq.recipients:type_name -> pkg.kannon.mailer.types.Recipient
	3, // 3: pkg.kannon.mailer.apiv1.SendTemplateReq.sender:type_name -> pkg.kannon.mailer.types.Sender
	4, // 4: pkg.kannon.mailer.apiv1.SendTemplateReq.scheduled_time:type_name -> google.protobuf.Timestamp
	5, // 5: pkg.kannon.mailer.apiv1.SendTemplateReq.recipients:type_name -> pkg.kannon.mailer.types.Recipient
	4, // 6: pkg.kannon.mailer.apiv1.SendRes.scheduled_time:type_name -> google.protobuf.Timestamp
	0, // 7: pkg.kannon.mailer.apiv1.Mailer.SendHTML:input_type -> pkg.kannon.mailer.apiv1.SendHTMLReq
	1, // 8: pkg.kannon.mailer.apiv1.Mailer.SendTemplate:input_type -> pkg.kannon.mailer.apiv1.SendTemplateReq
	2, // 9: pkg.kannon.mailer.apiv1.Mailer.SendHTML:output_type -> pkg.kannon.mailer.apiv1.SendRes
	2, // 10: pkg.kannon.mailer.apiv1.Mailer.SendTemplate:output_type -> pkg.kannon.mailer.apiv1.SendRes
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_kannon_mailer_apiv1_mailerapiv1_proto_init() }
func file_kannon_mailer_apiv1_mailerapiv1_proto_init() {
	if File_kannon_mailer_apiv1_mailerapiv1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kannon_mailer_apiv1_mailerapiv1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendHTMLReq); i {
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
		file_kannon_mailer_apiv1_mailerapiv1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTemplateReq); i {
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
		file_kannon_mailer_apiv1_mailerapiv1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRes); i {
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
	file_kannon_mailer_apiv1_mailerapiv1_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_kannon_mailer_apiv1_mailerapiv1_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kannon_mailer_apiv1_mailerapiv1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kannon_mailer_apiv1_mailerapiv1_proto_goTypes,
		DependencyIndexes: file_kannon_mailer_apiv1_mailerapiv1_proto_depIdxs,
		MessageInfos:      file_kannon_mailer_apiv1_mailerapiv1_proto_msgTypes,
	}.Build()
	File_kannon_mailer_apiv1_mailerapiv1_proto = out.File
	file_kannon_mailer_apiv1_mailerapiv1_proto_rawDesc = nil
	file_kannon_mailer_apiv1_mailerapiv1_proto_goTypes = nil
	file_kannon_mailer_apiv1_mailerapiv1_proto_depIdxs = nil
}
