// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.11.4
// source: ml.proto

package api

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

type Frames struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Frames []string `protobuf:"bytes,1,rep,name=frames,proto3" json:"frames,omitempty"`
}

func (x *Frames) Reset() {
	*x = Frames{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ml_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Frames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Frames) ProtoMessage() {}

func (x *Frames) ProtoReflect() protoreflect.Message {
	mi := &file_ml_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Frames.ProtoReflect.Descriptor instead.
func (*Frames) Descriptor() ([]byte, []int) {
	return file_ml_proto_rawDescGZIP(), []int{0}
}

func (x *Frames) GetFrames() []string {
	if x != nil {
		return x.Frames
	}
	return nil
}

type Box struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Box map[string]float32 `protobuf:"bytes,1,rep,name=box,proto3" json:"box,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
}

func (x *Box) Reset() {
	*x = Box{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ml_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Box) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Box) ProtoMessage() {}

func (x *Box) ProtoReflect() protoreflect.Message {
	mi := &file_ml_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Box.ProtoReflect.Descriptor instead.
func (*Box) Descriptor() ([]byte, []int) {
	return file_ml_proto_rawDescGZIP(), []int{1}
}

func (x *Box) GetBox() map[string]float32 {
	if x != nil {
		return x.Box
	}
	return nil
}

type FrameBoxes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Boxes []*Box `protobuf:"bytes,1,rep,name=boxes,proto3" json:"boxes,omitempty"`
}

func (x *FrameBoxes) Reset() {
	*x = FrameBoxes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ml_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrameBoxes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrameBoxes) ProtoMessage() {}

func (x *FrameBoxes) ProtoReflect() protoreflect.Message {
	mi := &file_ml_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrameBoxes.ProtoReflect.Descriptor instead.
func (*FrameBoxes) Descriptor() ([]byte, []int) {
	return file_ml_proto_rawDescGZIP(), []int{2}
}

func (x *FrameBoxes) GetBoxes() []*Box {
	if x != nil {
		return x.Boxes
	}
	return nil
}

type Prediction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*FrameBoxes `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *Prediction) Reset() {
	*x = Prediction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ml_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prediction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prediction) ProtoMessage() {}

func (x *Prediction) ProtoReflect() protoreflect.Message {
	mi := &file_ml_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prediction.ProtoReflect.Descriptor instead.
func (*Prediction) Descriptor() ([]byte, []int) {
	return file_ml_proto_rawDescGZIP(), []int{3}
}

func (x *Prediction) GetResult() []*FrameBoxes {
	if x != nil {
		return x.Result
	}
	return nil
}

type Classes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FoodClasses []int64 `protobuf:"varint,1,rep,packed,name=food_classes,json=foodClasses,proto3" json:"food_classes,omitempty"`
}

func (x *Classes) Reset() {
	*x = Classes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ml_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Classes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Classes) ProtoMessage() {}

func (x *Classes) ProtoReflect() protoreflect.Message {
	mi := &file_ml_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Classes.ProtoReflect.Descriptor instead.
func (*Classes) Descriptor() ([]byte, []int) {
	return file_ml_proto_rawDescGZIP(), []int{4}
}

func (x *Classes) GetFoodClasses() []int64 {
	if x != nil {
		return x.FoodClasses
	}
	return nil
}

var File_ml_proto protoreflect.FileDescriptor

var file_ml_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6d, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x06, 0x46, 0x72,
	0x61, 0x6d, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x5e, 0x0a, 0x03,
	0x42, 0x6f, 0x78, 0x12, 0x1f, 0x0a, 0x03, 0x62, 0x6f, 0x78, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x42, 0x6f, 0x78, 0x2e, 0x42, 0x6f, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x03, 0x62, 0x6f, 0x78, 0x1a, 0x36, 0x0a, 0x08, 0x42, 0x6f, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x29, 0x0a, 0x0b,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x05, 0x62,
	0x6f, 0x78, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x42, 0x6f, 0x78,
	0x52, 0x05, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x22, 0x32, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x62, 0x6f,
	0x78, 0x65, 0x73, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2c, 0x0a, 0x07, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x6f,
	0x6f, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x32, 0x52, 0x0a, 0x0c, 0x46, 0x6f, 0x6f,
	0x64, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x07, 0x70, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x12, 0x07, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x1a, 0x0b, 0x2e,
	0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0b, 0x73, 0x65,
	0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x12, 0x08, 0x2e, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x65, 0x73, 0x1a, 0x08, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x42, 0x0e, 0x5a,
	0x0c, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ml_proto_rawDescOnce sync.Once
	file_ml_proto_rawDescData = file_ml_proto_rawDesc
)

func file_ml_proto_rawDescGZIP() []byte {
	file_ml_proto_rawDescOnce.Do(func() {
		file_ml_proto_rawDescData = protoimpl.X.CompressGZIP(file_ml_proto_rawDescData)
	})
	return file_ml_proto_rawDescData
}

var file_ml_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ml_proto_goTypes = []interface{}{
	(*Frames)(nil),     // 0: Frames
	(*Box)(nil),        // 1: Box
	(*FrameBoxes)(nil), // 2: Frame_boxes
	(*Prediction)(nil), // 3: Prediction
	(*Classes)(nil),    // 4: Classes
	nil,                // 5: Box.BoxEntry
}
var file_ml_proto_depIdxs = []int32{
	5, // 0: Box.box:type_name -> Box.BoxEntry
	1, // 1: Frame_boxes.boxes:type_name -> Box
	2, // 2: Prediction.result:type_name -> Frame_boxes
	0, // 3: FoodDetector.predict:input_type -> Frames
	4, // 4: FoodDetector.set_classes:input_type -> Classes
	3, // 5: FoodDetector.predict:output_type -> Prediction
	4, // 6: FoodDetector.set_classes:output_type -> Classes
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ml_proto_init() }
func file_ml_proto_init() {
	if File_ml_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ml_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Frames); i {
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
		file_ml_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Box); i {
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
		file_ml_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrameBoxes); i {
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
		file_ml_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prediction); i {
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
		file_ml_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Classes); i {
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
			RawDescriptor: file_ml_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ml_proto_goTypes,
		DependencyIndexes: file_ml_proto_depIdxs,
		MessageInfos:      file_ml_proto_msgTypes,
	}.Build()
	File_ml_proto = out.File
	file_ml_proto_rawDesc = nil
	file_ml_proto_goTypes = nil
	file_ml_proto_depIdxs = nil
}
