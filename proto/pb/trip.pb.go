// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.21.12
// source: trip.proto

package proto

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

type TripVehicle struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TripId        uint64                 `protobuf:"varint,1,opt,name=trip_id,json=tripId,proto3" json:"trip_id,omitempty"`
	VehicleId     uint64                 `protobuf:"varint,2,opt,name=vehicle_id,json=vehicleId,proto3" json:"vehicle_id,omitempty"`
	VehicleSpeed  uint64                 `protobuf:"varint,3,opt,name=vehicle_speed,json=vehicleSpeed,proto3" json:"vehicle_speed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TripVehicle) Reset() {
	*x = TripVehicle{}
	mi := &file_trip_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TripVehicle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TripVehicle) ProtoMessage() {}

func (x *TripVehicle) ProtoReflect() protoreflect.Message {
	mi := &file_trip_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TripVehicle.ProtoReflect.Descriptor instead.
func (*TripVehicle) Descriptor() ([]byte, []int) {
	return file_trip_proto_rawDescGZIP(), []int{0}
}

func (x *TripVehicle) GetTripId() uint64 {
	if x != nil {
		return x.TripId
	}
	return 0
}

func (x *TripVehicle) GetVehicleId() uint64 {
	if x != nil {
		return x.VehicleId
	}
	return 0
}

func (x *TripVehicle) GetVehicleSpeed() uint64 {
	if x != nil {
		return x.VehicleSpeed
	}
	return 0
}

type TripVehicleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TripVehicleResponse) Reset() {
	*x = TripVehicleResponse{}
	mi := &file_trip_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TripVehicleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TripVehicleResponse) ProtoMessage() {}

func (x *TripVehicleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trip_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TripVehicleResponse.ProtoReflect.Descriptor instead.
func (*TripVehicleResponse) Descriptor() ([]byte, []int) {
	return file_trip_proto_rawDescGZIP(), []int{1}
}

func (x *TripVehicleResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_trip_proto protoreflect.FileDescriptor

var file_trip_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x0b,
	0x54, 0x72, 0x69, 0x70, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74,
	0x72, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x72,
	0x69, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x73,
	0x70, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x76, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x65, 0x64, 0x22, 0x2f, 0x0a, 0x13, 0x54, 0x72, 0x69, 0x70,
	0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x3f, 0x0a, 0x0b, 0x54, 0x72, 0x69,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x56,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0c, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x56, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x1a, 0x14, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x69, 0x2d, 0x6e, 0x61,
	0x62, 0x61, 0x62, 0x61, 0x2f, 0x67, 0x6f, 0x6c, 0x69, 0x62, 0x61, 0x62, 0x61, 0x2d, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_trip_proto_rawDescOnce sync.Once
	file_trip_proto_rawDescData = file_trip_proto_rawDesc
)

func file_trip_proto_rawDescGZIP() []byte {
	file_trip_proto_rawDescOnce.Do(func() {
		file_trip_proto_rawDescData = protoimpl.X.CompressGZIP(file_trip_proto_rawDescData)
	})
	return file_trip_proto_rawDescData
}

var file_trip_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_trip_proto_goTypes = []any{
	(*TripVehicle)(nil),         // 0: TripVehicle
	(*TripVehicleResponse)(nil), // 1: TripVehicleResponse
}
var file_trip_proto_depIdxs = []int32{
	0, // 0: TripService.SetVehicle:input_type -> TripVehicle
	1, // 1: TripService.SetVehicle:output_type -> TripVehicleResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_trip_proto_init() }
func file_trip_proto_init() {
	if File_trip_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_trip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trip_proto_goTypes,
		DependencyIndexes: file_trip_proto_depIdxs,
		MessageInfos:      file_trip_proto_msgTypes,
	}.Build()
	File_trip_proto = out.File
	file_trip_proto_rawDesc = nil
	file_trip_proto_goTypes = nil
	file_trip_proto_depIdxs = nil
}
