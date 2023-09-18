// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: client/proto/directoryInfo/infoDirectory.proto

package directoryInfo

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

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_directoryInfo_infoDirectory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_directoryInfo_infoDirectory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_client_proto_directoryInfo_infoDirectory_proto_rawDescGZIP(), []int{0}
}

func (x *FileInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type DirInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size        int64       `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Files       []*FileInfo `protobuf:"bytes,4,rep,name=files,proto3" json:"files,omitempty"`
	Directories []*DirInfo  `protobuf:"bytes,5,rep,name=directories,proto3" json:"directories,omitempty"`
	Path        string      `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *DirInfoResponse) Reset() {
	*x = DirInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_directoryInfo_infoDirectory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirInfoResponse) ProtoMessage() {}

func (x *DirInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_directoryInfo_infoDirectory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirInfoResponse.ProtoReflect.Descriptor instead.
func (*DirInfoResponse) Descriptor() ([]byte, []int) {
	return file_client_proto_directoryInfo_infoDirectory_proto_rawDescGZIP(), []int{1}
}

func (x *DirInfoResponse) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *DirInfoResponse) GetFiles() []*FileInfo {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *DirInfoResponse) GetDirectories() []*DirInfo {
	if x != nil {
		return x.Directories
	}
	return nil
}

func (x *DirInfoResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type DirInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Size int64  `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *DirInfo) Reset() {
	*x = DirInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_directoryInfo_infoDirectory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirInfo) ProtoMessage() {}

func (x *DirInfo) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_directoryInfo_infoDirectory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirInfo.ProtoReflect.Descriptor instead.
func (*DirInfo) Descriptor() ([]byte, []int) {
	return file_client_proto_directoryInfo_infoDirectory_proto_rawDescGZIP(), []int{2}
}

func (x *DirInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DirInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type PathRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *PathRequest) Reset() {
	*x = PathRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_directoryInfo_infoDirectory_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathRequest) ProtoMessage() {}

func (x *PathRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_directoryInfo_infoDirectory_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathRequest.ProtoReflect.Descriptor instead.
func (*PathRequest) Descriptor() ([]byte, []int) {
	return file_client_proto_directoryInfo_infoDirectory_proto_rawDescGZIP(), []int{3}
}

func (x *PathRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type DirInfoStreamClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DirectoriesInfo []*DirInfoResponse `protobuf:"bytes,1,rep,name=directoriesInfo,proto3" json:"directoriesInfo,omitempty"`
}

func (x *DirInfoStreamClientResponse) Reset() {
	*x = DirInfoStreamClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_directoryInfo_infoDirectory_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirInfoStreamClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirInfoStreamClientResponse) ProtoMessage() {}

func (x *DirInfoStreamClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_directoryInfo_infoDirectory_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirInfoStreamClientResponse.ProtoReflect.Descriptor instead.
func (*DirInfoStreamClientResponse) Descriptor() ([]byte, []int) {
	return file_client_proto_directoryInfo_infoDirectory_proto_rawDescGZIP(), []int{4}
}

func (x *DirInfoStreamClientResponse) GetDirectoriesInfo() []*DirInfoResponse {
	if x != nil {
		return x.DirectoriesInfo
	}
	return nil
}

var File_client_proto_directoryInfo_infoDirectory_proto protoreflect.FileDescriptor

var file_client_proto_directoryInfo_infoDirectory_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2f, 0x69, 0x6e, 0x66,
	0x6f, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x32, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x0f, 0x44, 0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x2a, 0x0a,
	0x0b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x44, 0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x31, 0x0a,
	0x07, 0x44, 0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x22, 0x21, 0x0a, 0x0b, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x22, 0x59, 0x0a, 0x1b, 0x44, 0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x44, 0x69,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0f, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x83,
	0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x2b, 0x0a, 0x07, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x69, 0x72, 0x12, 0x0c, 0x2e, 0x50, 0x61,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x44, 0x69, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a,
	0x13, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x69, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x0c, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x44, 0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x42, 0x19, 0x5a, 0x17, 0x64, 0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_proto_directoryInfo_infoDirectory_proto_rawDescOnce sync.Once
	file_client_proto_directoryInfo_infoDirectory_proto_rawDescData = file_client_proto_directoryInfo_infoDirectory_proto_rawDesc
)

func file_client_proto_directoryInfo_infoDirectory_proto_rawDescGZIP() []byte {
	file_client_proto_directoryInfo_infoDirectory_proto_rawDescOnce.Do(func() {
		file_client_proto_directoryInfo_infoDirectory_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_proto_directoryInfo_infoDirectory_proto_rawDescData)
	})
	return file_client_proto_directoryInfo_infoDirectory_proto_rawDescData
}

var file_client_proto_directoryInfo_infoDirectory_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_client_proto_directoryInfo_infoDirectory_proto_goTypes = []interface{}{
	(*FileInfo)(nil),                    // 0: FileInfo
	(*DirInfoResponse)(nil),             // 1: DirInfoResponse
	(*DirInfo)(nil),                     // 2: DirInfo
	(*PathRequest)(nil),                 // 3: PathRequest
	(*DirInfoStreamClientResponse)(nil), // 4: DirInfoStreamClientResponse
}
var file_client_proto_directoryInfo_infoDirectory_proto_depIdxs = []int32{
	0, // 0: DirInfoResponse.files:type_name -> FileInfo
	2, // 1: DirInfoResponse.directories:type_name -> DirInfo
	1, // 2: DirInfoStreamClientResponse.directoriesInfo:type_name -> DirInfoResponse
	3, // 3: InfoDirectory.InfoDir:input_type -> PathRequest
	3, // 4: InfoDirectory.InfoDirStreamClient:input_type -> PathRequest
	1, // 5: InfoDirectory.InfoDir:output_type -> DirInfoResponse
	4, // 6: InfoDirectory.InfoDirStreamClient:output_type -> DirInfoStreamClientResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_client_proto_directoryInfo_infoDirectory_proto_init() }
func file_client_proto_directoryInfo_infoDirectory_proto_init() {
	if File_client_proto_directoryInfo_infoDirectory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_proto_directoryInfo_infoDirectory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_client_proto_directoryInfo_infoDirectory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirInfoResponse); i {
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
		file_client_proto_directoryInfo_infoDirectory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirInfo); i {
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
		file_client_proto_directoryInfo_infoDirectory_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathRequest); i {
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
		file_client_proto_directoryInfo_infoDirectory_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirInfoStreamClientResponse); i {
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
			RawDescriptor: file_client_proto_directoryInfo_infoDirectory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_proto_directoryInfo_infoDirectory_proto_goTypes,
		DependencyIndexes: file_client_proto_directoryInfo_infoDirectory_proto_depIdxs,
		MessageInfos:      file_client_proto_directoryInfo_infoDirectory_proto_msgTypes,
	}.Build()
	File_client_proto_directoryInfo_infoDirectory_proto = out.File
	file_client_proto_directoryInfo_infoDirectory_proto_rawDesc = nil
	file_client_proto_directoryInfo_infoDirectory_proto_goTypes = nil
	file_client_proto_directoryInfo_infoDirectory_proto_depIdxs = nil
}
