// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: synchronization/core/cache.proto

package core

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

// CacheEntry represents cache data for a file on disk.
type CacheEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Mode stores the value of the POSIX mode bits (i.e. the st_mode member of
	// struct stat). On Windows, this value is computed using the Go os.FileMode
	// value retrieved through the os package (for which bit definitions are
	// guaranteed to be stable).
	Mode uint32 `protobuf:"varint,1,opt,name=mode,proto3" json:"mode,omitempty"`
	// ModificationTime is the cached modification time.
	ModificationTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=modificationTime,proto3" json:"modificationTime,omitempty"`
	// Size is the cached size.
	Size uint64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	// FileID is the file identifier. On POSIX systems it is the inode number.
	// On Windows it is currently 0.
	FileID uint64 `protobuf:"varint,4,opt,name=fileID,proto3" json:"fileID,omitempty"`
	// Digest is the cached digest for file entries.
	Digest []byte `protobuf:"bytes,9,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (x *CacheEntry) Reset() {
	*x = CacheEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synchronization_core_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheEntry) ProtoMessage() {}

func (x *CacheEntry) ProtoReflect() protoreflect.Message {
	mi := &file_synchronization_core_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheEntry.ProtoReflect.Descriptor instead.
func (*CacheEntry) Descriptor() ([]byte, []int) {
	return file_synchronization_core_cache_proto_rawDescGZIP(), []int{0}
}

func (x *CacheEntry) GetMode() uint32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *CacheEntry) GetModificationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ModificationTime
	}
	return nil
}

func (x *CacheEntry) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *CacheEntry) GetFileID() uint64 {
	if x != nil {
		return x.FileID
	}
	return 0
}

func (x *CacheEntry) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

// Cache provides a store for file metadata and digets to allow for efficient
// rescans.
type Cache struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Entries is a map from scan path to cache entry.
	Entries map[string]*CacheEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Cache) Reset() {
	*x = Cache{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synchronization_core_cache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cache) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cache) ProtoMessage() {}

func (x *Cache) ProtoReflect() protoreflect.Message {
	mi := &file_synchronization_core_cache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cache.ProtoReflect.Descriptor instead.
func (*Cache) Descriptor() ([]byte, []int) {
	return file_synchronization_core_cache_proto_rawDescGZIP(), []int{1}
}

func (x *Cache) GetEntries() map[string]*CacheEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

var File_synchronization_core_cache_proto protoreflect.FileDescriptor

var file_synchronization_core_cache_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x0a, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x46, 0x0a, 0x10,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x10, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x22, 0x89, 0x01, 0x0a, 0x05, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x4c, 0x0a, 0x0c, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2d, 0x69, 0x6f, 0x2f, 0x6d, 0x75,
	0x74, 0x61, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72,
	0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_synchronization_core_cache_proto_rawDescOnce sync.Once
	file_synchronization_core_cache_proto_rawDescData = file_synchronization_core_cache_proto_rawDesc
)

func file_synchronization_core_cache_proto_rawDescGZIP() []byte {
	file_synchronization_core_cache_proto_rawDescOnce.Do(func() {
		file_synchronization_core_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_synchronization_core_cache_proto_rawDescData)
	})
	return file_synchronization_core_cache_proto_rawDescData
}

var file_synchronization_core_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_synchronization_core_cache_proto_goTypes = []interface{}{
	(*CacheEntry)(nil),            // 0: core.CacheEntry
	(*Cache)(nil),                 // 1: core.Cache
	nil,                           // 2: core.Cache.EntriesEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_synchronization_core_cache_proto_depIdxs = []int32{
	3, // 0: core.CacheEntry.modificationTime:type_name -> google.protobuf.Timestamp
	2, // 1: core.Cache.entries:type_name -> core.Cache.EntriesEntry
	0, // 2: core.Cache.EntriesEntry.value:type_name -> core.CacheEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_synchronization_core_cache_proto_init() }
func file_synchronization_core_cache_proto_init() {
	if File_synchronization_core_cache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_synchronization_core_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheEntry); i {
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
		file_synchronization_core_cache_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cache); i {
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
			RawDescriptor: file_synchronization_core_cache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_synchronization_core_cache_proto_goTypes,
		DependencyIndexes: file_synchronization_core_cache_proto_depIdxs,
		MessageInfos:      file_synchronization_core_cache_proto_msgTypes,
	}.Build()
	File_synchronization_core_cache_proto = out.File
	file_synchronization_core_cache_proto_rawDesc = nil
	file_synchronization_core_cache_proto_goTypes = nil
	file_synchronization_core_cache_proto_depIdxs = nil
}
