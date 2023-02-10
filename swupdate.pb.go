// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: swupdate.proto

package swupdatesvc

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swupdate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_swupdate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_swupdate_proto_rawDescGZIP(), []int{0}
}

type CheckForUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateUrl string `protobuf:"bytes,1,opt,name=updateUrl,proto3" json:"updateUrl,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CheckForUpdateRequest) Reset() {
	*x = CheckForUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swupdate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckForUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckForUpdateRequest) ProtoMessage() {}

func (x *CheckForUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_swupdate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckForUpdateRequest.ProtoReflect.Descriptor instead.
func (*CheckForUpdateRequest) Descriptor() ([]byte, []int) {
	return file_swupdate_proto_rawDescGZIP(), []int{1}
}

func (x *CheckForUpdateRequest) GetUpdateUrl() string {
	if x != nil {
		return x.UpdateUrl
	}
	return ""
}

func (x *CheckForUpdateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CheckForUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EngineVersion string            `protobuf:"bytes,1,opt,name=engineVersion,proto3" json:"engineVersion,omitempty"`
	IsOkay        bool              `protobuf:"varint,2,opt,name=isOkay,proto3" json:"isOkay,omitempty"`
	ErrorMessage  string            `protobuf:"bytes,3,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	UserData      map[string]string `protobuf:"bytes,4,rep,name=userData,proto3" json:"userData,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CheckForUpdateResponse) Reset() {
	*x = CheckForUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swupdate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckForUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckForUpdateResponse) ProtoMessage() {}

func (x *CheckForUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_swupdate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckForUpdateResponse.ProtoReflect.Descriptor instead.
func (*CheckForUpdateResponse) Descriptor() ([]byte, []int) {
	return file_swupdate_proto_rawDescGZIP(), []int{2}
}

func (x *CheckForUpdateResponse) GetEngineVersion() string {
	if x != nil {
		return x.EngineVersion
	}
	return ""
}

func (x *CheckForUpdateResponse) GetIsOkay() bool {
	if x != nil {
		return x.IsOkay
	}
	return false
}

func (x *CheckForUpdateResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *CheckForUpdateResponse) GetUserData() map[string]string {
	if x != nil {
		return x.UserData
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swupdate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_swupdate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_swupdate_proto_rawDescGZIP(), []int{3}
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Progress float32 `protobuf:"fixed32,1,opt,name=progress,proto3" json:"progress,omitempty"`
	Done     bool    `protobuf:"varint,2,opt,name=done,proto3" json:"done,omitempty"`
	Msg      string  `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swupdate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_swupdate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_swupdate_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateResponse) GetProgress() float32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *UpdateResponse) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *UpdateResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_swupdate_proto protoreflect.FileDescriptor

var file_swupdate_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x77, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x73, 0x77, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x51, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x6f, 0x72, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x83, 0x02, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x46, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x4f, 0x6b, 0x61,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4f, 0x6b, 0x61, 0x79, 0x12,
	0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x4a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x77, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x1a,
	0x3b, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x0f, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x52, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x32, 0xd8, 0x01, 0x0a, 0x07, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x12, 0x57, 0x0a,
	0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1f, 0x2e, 0x73, 0x77, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x46, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x73, 0x77, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x46, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x46, 0x0a, 0x0d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x77, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x73, 0x77, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2c,
	0x0a, 0x06, 0x52, 0x65, 0x62, 0x6f, 0x6f, 0x74, 0x12, 0x0f, 0x2e, 0x73, 0x77, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x73, 0x77, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x4b, 0x0a, 0x18,
	0x64, 0x65, 0x2e, 0x69, 0x64, 0x61, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x6d, 0x2e, 0x73, 0x77, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x76, 0x63, 0x42, 0x0b, 0x53, 0x77, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x76, 0x63, 0x50, 0x01, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x64, 0x61, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x6d, 0x2f, 0x73, 0x77,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x76, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_swupdate_proto_rawDescOnce sync.Once
	file_swupdate_proto_rawDescData = file_swupdate_proto_rawDesc
)

func file_swupdate_proto_rawDescGZIP() []byte {
	file_swupdate_proto_rawDescOnce.Do(func() {
		file_swupdate_proto_rawDescData = protoimpl.X.CompressGZIP(file_swupdate_proto_rawDescData)
	})
	return file_swupdate_proto_rawDescData
}

var file_swupdate_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_swupdate_proto_goTypes = []interface{}{
	(*Empty)(nil),                  // 0: swupdate.Empty
	(*CheckForUpdateRequest)(nil),  // 1: swupdate.CheckForUpdateRequest
	(*CheckForUpdateResponse)(nil), // 2: swupdate.CheckForUpdateResponse
	(*UpdateRequest)(nil),          // 3: swupdate.UpdateRequest
	(*UpdateResponse)(nil),         // 4: swupdate.UpdateResponse
	nil,                            // 5: swupdate.CheckForUpdateResponse.UserDataEntry
}
var file_swupdate_proto_depIdxs = []int32{
	5, // 0: swupdate.CheckForUpdateResponse.userData:type_name -> swupdate.CheckForUpdateResponse.UserDataEntry
	1, // 1: swupdate.Updater.CheckForUpdate:input_type -> swupdate.CheckForUpdateRequest
	3, // 2: swupdate.Updater.ExecuteUpdate:input_type -> swupdate.UpdateRequest
	0, // 3: swupdate.Updater.Reboot:input_type -> swupdate.Empty
	2, // 4: swupdate.Updater.CheckForUpdate:output_type -> swupdate.CheckForUpdateResponse
	4, // 5: swupdate.Updater.ExecuteUpdate:output_type -> swupdate.UpdateResponse
	0, // 6: swupdate.Updater.Reboot:output_type -> swupdate.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_swupdate_proto_init() }
func file_swupdate_proto_init() {
	if File_swupdate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_swupdate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_swupdate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckForUpdateRequest); i {
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
		file_swupdate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckForUpdateResponse); i {
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
		file_swupdate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_swupdate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
			RawDescriptor: file_swupdate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_swupdate_proto_goTypes,
		DependencyIndexes: file_swupdate_proto_depIdxs,
		MessageInfos:      file_swupdate_proto_msgTypes,
	}.Build()
	File_swupdate_proto = out.File
	file_swupdate_proto_rawDesc = nil
	file_swupdate_proto_goTypes = nil
	file_swupdate_proto_depIdxs = nil
}
