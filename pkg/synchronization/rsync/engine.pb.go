// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: synchronization/rsync/engine.proto

package rsync

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

// BlockHash represents a pair of weak and strong hash for a base block.
type BlockHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Weak is the weak hash for the block.
	Weak uint32 `protobuf:"varint,1,opt,name=weak,proto3" json:"weak,omitempty"`
	// Strong is the strong hash for the block.
	Strong []byte `protobuf:"bytes,2,opt,name=strong,proto3" json:"strong,omitempty"`
}

func (x *BlockHash) Reset() {
	*x = BlockHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synchronization_rsync_engine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockHash) ProtoMessage() {}

func (x *BlockHash) ProtoReflect() protoreflect.Message {
	mi := &file_synchronization_rsync_engine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockHash.ProtoReflect.Descriptor instead.
func (*BlockHash) Descriptor() ([]byte, []int) {
	return file_synchronization_rsync_engine_proto_rawDescGZIP(), []int{0}
}

func (x *BlockHash) GetWeak() uint32 {
	if x != nil {
		return x.Weak
	}
	return 0
}

func (x *BlockHash) GetStrong() []byte {
	if x != nil {
		return x.Strong
	}
	return nil
}

// Signature represents an rsync base signature. It encodes the block size used
// to generate the signature, the size of the last block in the signature (which
// may be smaller than a full block), and the hashes for the blocks of the file.
type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// BlockSize is the block size used to compute the signature.
	BlockSize uint64 `protobuf:"varint,1,opt,name=blockSize,proto3" json:"blockSize,omitempty"`
	// LastBlockSize is the size of the last block in the signature.
	LastBlockSize uint64 `protobuf:"varint,2,opt,name=lastBlockSize,proto3" json:"lastBlockSize,omitempty"`
	// Hashes are the hashes of the blocks in the base.
	Hashes []*BlockHash `protobuf:"bytes,3,rep,name=hashes,proto3" json:"hashes,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synchronization_rsync_engine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_synchronization_rsync_engine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_synchronization_rsync_engine_proto_rawDescGZIP(), []int{1}
}

func (x *Signature) GetBlockSize() uint64 {
	if x != nil {
		return x.BlockSize
	}
	return 0
}

func (x *Signature) GetLastBlockSize() uint64 {
	if x != nil {
		return x.LastBlockSize
	}
	return 0
}

func (x *Signature) GetHashes() []*BlockHash {
	if x != nil {
		return x.Hashes
	}
	return nil
}

// Operation represents an rsync operation, which can be either a data operation
// or a block operation.
type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data contains data for data operations. If its length is 0, the operation
	// is assumed to be a non-data operation. Operation transmitters and
	// receivers may thus treat a length-0 buffer as semantically equivalent to
	// a nil buffer and utilize that fact to efficiently re-use buffer capacity,
	// e.g. by truncating the buffer and doing a gob receive into it.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Start is the 0-indexed starting block for block operations.
	Start uint64 `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	// Count is the number of blocks for block operations.
	Count uint64 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synchronization_rsync_engine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_synchronization_rsync_engine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_synchronization_rsync_engine_proto_rawDescGZIP(), []int{2}
}

func (x *Operation) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Operation) GetStart() uint64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Operation) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_synchronization_rsync_engine_proto protoreflect.FileDescriptor

var file_synchronization_rsync_engine_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x72, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72, 0x73, 0x79, 0x6e, 0x63, 0x22, 0x37, 0x0a, 0x09, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x65, 0x61, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x77, 0x65, 0x61, 0x6b, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x74,
	0x72, 0x6f, 0x6e, 0x67, 0x22, 0x79, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x52, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x22,
	0x4b, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x39, 0x5a, 0x37,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67,
	0x65, 0x6e, 0x2d, 0x69, 0x6f, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x72, 0x73, 0x79, 0x6e, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_synchronization_rsync_engine_proto_rawDescOnce sync.Once
	file_synchronization_rsync_engine_proto_rawDescData = file_synchronization_rsync_engine_proto_rawDesc
)

func file_synchronization_rsync_engine_proto_rawDescGZIP() []byte {
	file_synchronization_rsync_engine_proto_rawDescOnce.Do(func() {
		file_synchronization_rsync_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_synchronization_rsync_engine_proto_rawDescData)
	})
	return file_synchronization_rsync_engine_proto_rawDescData
}

var file_synchronization_rsync_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_synchronization_rsync_engine_proto_goTypes = []interface{}{
	(*BlockHash)(nil), // 0: rsync.BlockHash
	(*Signature)(nil), // 1: rsync.Signature
	(*Operation)(nil), // 2: rsync.Operation
}
var file_synchronization_rsync_engine_proto_depIdxs = []int32{
	0, // 0: rsync.Signature.hashes:type_name -> rsync.BlockHash
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_synchronization_rsync_engine_proto_init() }
func file_synchronization_rsync_engine_proto_init() {
	if File_synchronization_rsync_engine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_synchronization_rsync_engine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockHash); i {
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
		file_synchronization_rsync_engine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
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
		file_synchronization_rsync_engine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
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
			RawDescriptor: file_synchronization_rsync_engine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_synchronization_rsync_engine_proto_goTypes,
		DependencyIndexes: file_synchronization_rsync_engine_proto_depIdxs,
		MessageInfos:      file_synchronization_rsync_engine_proto_msgTypes,
	}.Build()
	File_synchronization_rsync_engine_proto = out.File
	file_synchronization_rsync_engine_proto_rawDesc = nil
	file_synchronization_rsync_engine_proto_goTypes = nil
	file_synchronization_rsync_engine_proto_depIdxs = nil
}
