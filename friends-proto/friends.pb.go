// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: friends.proto

package friends_proto

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

// Data for add peer_1 and peer_2
type SetFriendsIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Peer_1 id
	Peer_1 string `protobuf:"bytes,1,opt,name=peer_1,json=peer1,proto3" json:"peer_1,omitempty"`
	// Peer_2 id
	Peer_2 string `protobuf:"bytes,2,opt,name=peer_2,json=peer2,proto3" json:"peer_2,omitempty"`
}

func (x *SetFriendsIn) Reset() {
	*x = SetFriendsIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetFriendsIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetFriendsIn) ProtoMessage() {}

func (x *SetFriendsIn) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetFriendsIn.ProtoReflect.Descriptor instead.
func (*SetFriendsIn) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{0}
}

func (x *SetFriendsIn) GetPeer_1() string {
	if x != nil {
		return x.Peer_1
	}
	return ""
}

func (x *SetFriendsIn) GetPeer_2() string {
	if x != nil {
		return x.Peer_2
	}
	return ""
}

// Response add friend
type SetFriendsOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Result of the operation
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SetFriendsOut) Reset() {
	*x = SetFriendsOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetFriendsOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetFriendsOut) ProtoMessage() {}

func (x *SetFriendsOut) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetFriendsOut.ProtoReflect.Descriptor instead.
func (*SetFriendsOut) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{1}
}

func (x *SetFriendsOut) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// Data for peer
type Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Peer uuid
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *Peer) Reset() {
	*x = Peer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer) ProtoMessage() {}

func (x *Peer) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peer.ProtoReflect.Descriptor instead.
func (*Peer) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{2}
}

func (x *Peer) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type RemoveSubscribeIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peer_1 string `protobuf:"bytes,1,opt,name=peer_1,json=peer1,proto3" json:"peer_1,omitempty"`
	Peer_2 string `protobuf:"bytes,2,opt,name=peer_2,json=peer2,proto3" json:"peer_2,omitempty"`
}

func (x *RemoveSubscribeIn) Reset() {
	*x = RemoveSubscribeIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveSubscribeIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveSubscribeIn) ProtoMessage() {}

func (x *RemoveSubscribeIn) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveSubscribeIn.ProtoReflect.Descriptor instead.
func (*RemoveSubscribeIn) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveSubscribeIn) GetPeer_1() string {
	if x != nil {
		return x.Peer_1
	}
	return ""
}

func (x *RemoveSubscribeIn) GetPeer_2() string {
	if x != nil {
		return x.Peer_2
	}
	return ""
}

type RemoveSubscribeOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveSubscribeOut) Reset() {
	*x = RemoveSubscribeOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveSubscribeOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveSubscribeOut) ProtoMessage() {}

func (x *RemoveSubscribeOut) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveSubscribeOut.ProtoReflect.Descriptor instead.
func (*RemoveSubscribeOut) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{4}
}

type GetInvitePeerIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid  string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *GetInvitePeerIn) Reset() {
	*x = GetInvitePeerIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvitePeerIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvitePeerIn) ProtoMessage() {}

func (x *GetInvitePeerIn) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvitePeerIn.ProtoReflect.Descriptor instead.
func (*GetInvitePeerIn) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{5}
}

func (x *GetInvitePeerIn) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetInvitePeerIn) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetInvitePeerOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetInvitePeerOut) Reset() {
	*x = GetInvitePeerOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvitePeerOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvitePeerOut) ProtoMessage() {}

func (x *GetInvitePeerOut) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvitePeerOut.ProtoReflect.Descriptor instead.
func (*GetInvitePeerOut) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{6}
}

// Response subscription
type GetPeerFollowOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Result of the operation
	Subscription []*Peer `protobuf:"bytes,1,rep,name=subscription,proto3" json:"subscription,omitempty"`
}

func (x *GetPeerFollowOut) Reset() {
	*x = GetPeerFollowOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPeerFollowOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPeerFollowOut) ProtoMessage() {}

func (x *GetPeerFollowOut) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPeerFollowOut.ProtoReflect.Descriptor instead.
func (*GetPeerFollowOut) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{7}
}

func (x *GetPeerFollowOut) GetSubscription() []*Peer {
	if x != nil {
		return x.Subscription
	}
	return nil
}

// Response subscribers
type GetWhoFollowPeerOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Result of the operation
	Subscribers []*Peer `protobuf:"bytes,1,rep,name=subscribers,proto3" json:"subscribers,omitempty"`
}

func (x *GetWhoFollowPeerOut) Reset() {
	*x = GetWhoFollowPeerOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWhoFollowPeerOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWhoFollowPeerOut) ProtoMessage() {}

func (x *GetWhoFollowPeerOut) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWhoFollowPeerOut.ProtoReflect.Descriptor instead.
func (*GetWhoFollowPeerOut) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{8}
}

func (x *GetWhoFollowPeerOut) GetSubscribers() []*Peer {
	if x != nil {
		return x.Subscribers
	}
	return nil
}

// Request for subscription
type GetPeerFollowIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Peer uuid
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetPeerFollowIn) Reset() {
	*x = GetPeerFollowIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPeerFollowIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPeerFollowIn) ProtoMessage() {}

func (x *GetPeerFollowIn) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPeerFollowIn.ProtoReflect.Descriptor instead.
func (*GetPeerFollowIn) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{9}
}

func (x *GetPeerFollowIn) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// Request for subscribers
type GetWhoFollowPeerIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Peer uuid
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetWhoFollowPeerIn) Reset() {
	*x = GetWhoFollowPeerIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWhoFollowPeerIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWhoFollowPeerIn) ProtoMessage() {}

func (x *GetWhoFollowPeerIn) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWhoFollowPeerIn.ProtoReflect.Descriptor instead.
func (*GetWhoFollowPeerIn) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{10}
}

func (x *GetWhoFollowPeerIn) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

var File_friends_proto protoreflect.FileDescriptor

var file_friends_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x3c, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x49, 0x6e, 0x12,
	0x15, 0x0a, 0x06, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x65, 0x65, 0x72, 0x31, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x32,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x32, 0x22, 0x29, 0x0a,
	0x0d, 0x53, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x1a, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x49, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x65, 0x65,
	0x72, 0x5f, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x31,
	0x12, 0x15, 0x0a, 0x06, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x32, 0x22, 0x14, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4f, 0x75, 0x74, 0x22, 0x3b, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x22, 0x3d,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4f,
	0x75, 0x74, 0x12, 0x29, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52,
	0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3e, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x57, 0x68, 0x6f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x65, 0x65,
	0x72, 0x4f, 0x75, 0x74, 0x12, 0x27, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x65, 0x65, 0x72,
	0x52, 0x0b, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x22, 0x25, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x57, 0x68, 0x6f, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x32, 0xae,
	0x02, 0x0a, 0x0e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12,
	0x0d, 0x2e, 0x53, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x49, 0x6e, 0x1a, 0x0e,
	0x2e, 0x53, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x4f, 0x75, 0x74, 0x22, 0x00,
	0x12, 0x36, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x49, 0x6e, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x57,
	0x68, 0x6f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x65, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x47,
	0x65, 0x74, 0x57, 0x68, 0x6f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x65, 0x65, 0x72, 0x49,
	0x6e, 0x1a, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x68, 0x6f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x50, 0x65, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0f, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x12, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x49, 0x6e,
	0x1a, 0x13, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x42,
	0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_friends_proto_rawDescOnce sync.Once
	file_friends_proto_rawDescData = file_friends_proto_rawDesc
)

func file_friends_proto_rawDescGZIP() []byte {
	file_friends_proto_rawDescOnce.Do(func() {
		file_friends_proto_rawDescData = protoimpl.X.CompressGZIP(file_friends_proto_rawDescData)
	})
	return file_friends_proto_rawDescData
}

var file_friends_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_friends_proto_goTypes = []any{
	(*SetFriendsIn)(nil),        // 0: SetFriendsIn
	(*SetFriendsOut)(nil),       // 1: SetFriendsOut
	(*Peer)(nil),                // 2: Peer
	(*RemoveSubscribeIn)(nil),   // 3: RemoveSubscribeIn
	(*RemoveSubscribeOut)(nil),  // 4: RemoveSubscribeOut
	(*GetInvitePeerIn)(nil),     // 5: GetInvitePeerIn
	(*GetInvitePeerOut)(nil),    // 6: GetInvitePeerOut
	(*GetPeerFollowOut)(nil),    // 7: GetPeerFollowOut
	(*GetWhoFollowPeerOut)(nil), // 8: GetWhoFollowPeerOut
	(*GetPeerFollowIn)(nil),     // 9: GetPeerFollowIn
	(*GetWhoFollowPeerIn)(nil),  // 10: GetWhoFollowPeerIn
}
var file_friends_proto_depIdxs = []int32{
	2,  // 0: GetPeerFollowOut.subscription:type_name -> Peer
	2,  // 1: GetWhoFollowPeerOut.subscribers:type_name -> Peer
	0,  // 2: FriendsService.SetFriends:input_type -> SetFriendsIn
	9,  // 3: FriendsService.GetPeerFollow:input_type -> GetPeerFollowIn
	10, // 4: FriendsService.GetWhoFollowPeer:input_type -> GetWhoFollowPeerIn
	3,  // 5: FriendsService.RemoveSubscribe:input_type -> RemoveSubscribeIn
	5,  // 6: FriendsService.GetInvitePeer:input_type -> GetInvitePeerIn
	1,  // 7: FriendsService.SetFriends:output_type -> SetFriendsOut
	7,  // 8: FriendsService.GetPeerFollow:output_type -> GetPeerFollowOut
	8,  // 9: FriendsService.GetWhoFollowPeer:output_type -> GetWhoFollowPeerOut
	4,  // 10: FriendsService.RemoveSubscribe:output_type -> RemoveSubscribeOut
	6,  // 11: FriendsService.GetInvitePeer:output_type -> GetInvitePeerOut
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_friends_proto_init() }
func file_friends_proto_init() {
	if File_friends_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_friends_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SetFriendsIn); i {
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
		file_friends_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SetFriendsOut); i {
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
		file_friends_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Peer); i {
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
		file_friends_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RemoveSubscribeIn); i {
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
		file_friends_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*RemoveSubscribeOut); i {
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
		file_friends_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetInvitePeerIn); i {
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
		file_friends_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetInvitePeerOut); i {
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
		file_friends_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetPeerFollowOut); i {
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
		file_friends_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetWhoFollowPeerOut); i {
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
		file_friends_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GetPeerFollowIn); i {
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
		file_friends_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*GetWhoFollowPeerIn); i {
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
			RawDescriptor: file_friends_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_friends_proto_goTypes,
		DependencyIndexes: file_friends_proto_depIdxs,
		MessageInfos:      file_friends_proto_msgTypes,
	}.Build()
	File_friends_proto = out.File
	file_friends_proto_rawDesc = nil
	file_friends_proto_goTypes = nil
	file_friends_proto_depIdxs = nil
}
