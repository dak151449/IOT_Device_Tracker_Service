// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.2
// source: api/device_tracker/device_tracker.proto

package dtapi

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

type GetDeviceGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDeviceGroupsRequest) Reset() {
	*x = GetDeviceGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_tracker_device_tracker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceGroupsRequest) ProtoMessage() {}

func (x *GetDeviceGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_tracker_device_tracker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceGroupsRequest.ProtoReflect.Descriptor instead.
func (*GetDeviceGroupsRequest) Descriptor() ([]byte, []int) {
	return file_api_device_tracker_device_tracker_proto_rawDescGZIP(), []int{0}
}

type DeviceGroupData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status      string                 `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Description string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *DeviceGroupData) Reset() {
	*x = DeviceGroupData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_tracker_device_tracker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceGroupData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceGroupData) ProtoMessage() {}

func (x *DeviceGroupData) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_tracker_device_tracker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceGroupData.ProtoReflect.Descriptor instead.
func (*DeviceGroupData) Descriptor() ([]byte, []int) {
	return file_api_device_tracker_device_tracker_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceGroupData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeviceGroupData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceGroupData) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DeviceGroupData) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DeviceGroupData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetDeviceGroupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*DeviceGroupData `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *GetDeviceGroupsResponse) Reset() {
	*x = GetDeviceGroupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_tracker_device_tracker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceGroupsResponse) ProtoMessage() {}

func (x *GetDeviceGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_tracker_device_tracker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceGroupsResponse.ProtoReflect.Descriptor instead.
func (*GetDeviceGroupsResponse) Descriptor() ([]byte, []int) {
	return file_api_device_tracker_device_tracker_proto_rawDescGZIP(), []int{2}
}

func (x *GetDeviceGroupsResponse) GetGroups() []*DeviceGroupData {
	if x != nil {
		return x.Groups
	}
	return nil
}

type GetDevicesFromGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId int64 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *GetDevicesFromGroupRequest) Reset() {
	*x = GetDevicesFromGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_tracker_device_tracker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDevicesFromGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDevicesFromGroupRequest) ProtoMessage() {}

func (x *GetDevicesFromGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_tracker_device_tracker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDevicesFromGroupRequest.ProtoReflect.Descriptor instead.
func (*GetDevicesFromGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_device_tracker_device_tracker_proto_rawDescGZIP(), []int{3}
}

func (x *GetDevicesFromGroupRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type DeviceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status      string                 `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Description string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *DeviceData) Reset() {
	*x = DeviceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_tracker_device_tracker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceData) ProtoMessage() {}

func (x *DeviceData) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_tracker_device_tracker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceData.ProtoReflect.Descriptor instead.
func (*DeviceData) Descriptor() ([]byte, []int) {
	return file_api_device_tracker_device_tracker_proto_rawDescGZIP(), []int{4}
}

func (x *DeviceData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeviceData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceData) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DeviceData) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DeviceData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetDevicesFromGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Devices []*DeviceData `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (x *GetDevicesFromGroupResponse) Reset() {
	*x = GetDevicesFromGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_tracker_device_tracker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDevicesFromGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDevicesFromGroupResponse) ProtoMessage() {}

func (x *GetDevicesFromGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_tracker_device_tracker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDevicesFromGroupResponse.ProtoReflect.Descriptor instead.
func (*GetDevicesFromGroupResponse) Descriptor() ([]byte, []int) {
	return file_api_device_tracker_device_tracker_proto_rawDescGZIP(), []int{5}
}

func (x *GetDevicesFromGroupResponse) GetDevices() []*DeviceData {
	if x != nil {
		return x.Devices
	}
	return nil
}

var File_api_device_tracker_device_tracker_proto protoreflect.FileDescriptor

var file_api_device_tracker_device_tracker_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0xaa, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x52, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x37, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0xa5,
	0x01, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x53, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x32, 0xee, 0x01, 0x0a, 0x14,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x26, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x2a, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x46, 0x72, 0x6f,
	0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4a, 0x5a, 0x48,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6b, 0x31, 0x35,
	0x31, 0x34, 0x34, 0x39, 0x2f, 0x49, 0x4f, 0x54, 0x5f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x72, 0x3b, 0x64, 0x74, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_device_tracker_device_tracker_proto_rawDescOnce sync.Once
	file_api_device_tracker_device_tracker_proto_rawDescData = file_api_device_tracker_device_tracker_proto_rawDesc
)

func file_api_device_tracker_device_tracker_proto_rawDescGZIP() []byte {
	file_api_device_tracker_device_tracker_proto_rawDescOnce.Do(func() {
		file_api_device_tracker_device_tracker_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_device_tracker_device_tracker_proto_rawDescData)
	})
	return file_api_device_tracker_device_tracker_proto_rawDescData
}

var file_api_device_tracker_device_tracker_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_device_tracker_device_tracker_proto_goTypes = []interface{}{
	(*GetDeviceGroupsRequest)(nil),      // 0: device_tracker.GetDeviceGroupsRequest
	(*DeviceGroupData)(nil),             // 1: device_tracker.DeviceGroupData
	(*GetDeviceGroupsResponse)(nil),     // 2: device_tracker.GetDeviceGroupsResponse
	(*GetDevicesFromGroupRequest)(nil),  // 3: device_tracker.GetDevicesFromGroupRequest
	(*DeviceData)(nil),                  // 4: device_tracker.DeviceData
	(*GetDevicesFromGroupResponse)(nil), // 5: device_tracker.GetDevicesFromGroupResponse
	(*timestamppb.Timestamp)(nil),       // 6: google.protobuf.Timestamp
}
var file_api_device_tracker_device_tracker_proto_depIdxs = []int32{
	6, // 0: device_tracker.DeviceGroupData.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: device_tracker.GetDeviceGroupsResponse.groups:type_name -> device_tracker.DeviceGroupData
	6, // 2: device_tracker.DeviceData.created_at:type_name -> google.protobuf.Timestamp
	4, // 3: device_tracker.GetDevicesFromGroupResponse.devices:type_name -> device_tracker.DeviceData
	0, // 4: device_tracker.DeviceTrackerService.GetDeviceGroups:input_type -> device_tracker.GetDeviceGroupsRequest
	3, // 5: device_tracker.DeviceTrackerService.GetDevicesFromGroup:input_type -> device_tracker.GetDevicesFromGroupRequest
	2, // 6: device_tracker.DeviceTrackerService.GetDeviceGroups:output_type -> device_tracker.GetDeviceGroupsResponse
	5, // 7: device_tracker.DeviceTrackerService.GetDevicesFromGroup:output_type -> device_tracker.GetDevicesFromGroupResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_device_tracker_device_tracker_proto_init() }
func file_api_device_tracker_device_tracker_proto_init() {
	if File_api_device_tracker_device_tracker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_device_tracker_device_tracker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceGroupsRequest); i {
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
		file_api_device_tracker_device_tracker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceGroupData); i {
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
		file_api_device_tracker_device_tracker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceGroupsResponse); i {
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
		file_api_device_tracker_device_tracker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDevicesFromGroupRequest); i {
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
		file_api_device_tracker_device_tracker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceData); i {
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
		file_api_device_tracker_device_tracker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDevicesFromGroupResponse); i {
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
			RawDescriptor: file_api_device_tracker_device_tracker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_device_tracker_device_tracker_proto_goTypes,
		DependencyIndexes: file_api_device_tracker_device_tracker_proto_depIdxs,
		MessageInfos:      file_api_device_tracker_device_tracker_proto_msgTypes,
	}.Build()
	File_api_device_tracker_device_tracker_proto = out.File
	file_api_device_tracker_device_tracker_proto_rawDesc = nil
	file_api_device_tracker_device_tracker_proto_goTypes = nil
	file_api_device_tracker_device_tracker_proto_depIdxs = nil
}
