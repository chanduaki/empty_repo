// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/enums.proto

package ttnpb

import (
	fmt "fmt"
	math "math"
	strconv "strconv"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DownlinkPathConstraint int32

const (
	// Indicates that the gateway can be selected for downlink without constraints by the Network Server.
	DOWNLINK_PATH_CONSTRAINT_NONE DownlinkPathConstraint = 0
	// Indicates that the gateway can be selected for downlink only if no other or better gateway can be selected.
	DOWNLINK_PATH_CONSTRAINT_PREFER_OTHER DownlinkPathConstraint = 1
	// Indicates that this gateway will never be selected for downlink, even if that results in no available downlink path.
	DOWNLINK_PATH_CONSTRAINT_NEVER DownlinkPathConstraint = 2
)

var DownlinkPathConstraint_name = map[int32]string{
	0: "DOWNLINK_PATH_CONSTRAINT_NONE",
	1: "DOWNLINK_PATH_CONSTRAINT_PREFER_OTHER",
	2: "DOWNLINK_PATH_CONSTRAINT_NEVER",
}

var DownlinkPathConstraint_value = map[string]int32{
	"DOWNLINK_PATH_CONSTRAINT_NONE":         0,
	"DOWNLINK_PATH_CONSTRAINT_PREFER_OTHER": 1,
	"DOWNLINK_PATH_CONSTRAINT_NEVER":        2,
}

func (DownlinkPathConstraint) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e36318a1e2f407cb, []int{0}
}

// State enum defines states that an entity can be in.
type State int32

const (
	// Denotes that the entity has been requested and is pending review by an admin.
	STATE_REQUESTED State = 0
	// Denotes that the entity has been reviewed and approved by an admin.
	STATE_APPROVED State = 1
	// Denotes that the entity has been reviewed and rejected by an admin.
	STATE_REJECTED State = 2
	// Denotes that the entity has been flagged and is pending review by an admin.
	STATE_FLAGGED State = 3
	// Denotes that the entity has been reviewed and suspended by an admin.
	STATE_SUSPENDED State = 4
)

var State_name = map[int32]string{
	0: "STATE_REQUESTED",
	1: "STATE_APPROVED",
	2: "STATE_REJECTED",
	3: "STATE_FLAGGED",
	4: "STATE_SUSPENDED",
}

var State_value = map[string]int32{
	"STATE_REQUESTED": 0,
	"STATE_APPROVED":  1,
	"STATE_REJECTED":  2,
	"STATE_FLAGGED":   3,
	"STATE_SUSPENDED": 4,
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e36318a1e2f407cb, []int{1}
}

type ClusterRole int32

const (
	ClusterRole_NONE                         ClusterRole = 0
	ClusterRole_ENTITY_REGISTRY              ClusterRole = 1
	ClusterRole_ACCESS                       ClusterRole = 2
	ClusterRole_GATEWAY_SERVER               ClusterRole = 3
	ClusterRole_NETWORK_SERVER               ClusterRole = 4
	ClusterRole_APPLICATION_SERVER           ClusterRole = 5
	ClusterRole_JOIN_SERVER                  ClusterRole = 6
	ClusterRole_CRYPTO_SERVER                ClusterRole = 7
	ClusterRole_DEVICE_TEMPLATE_CONVERTER    ClusterRole = 8
	ClusterRole_DEVICE_CLAIMING_SERVER       ClusterRole = 9
	ClusterRole_GATEWAY_CONFIGURATION_SERVER ClusterRole = 10
	ClusterRole_QR_CODE_GENERATOR            ClusterRole = 11
	ClusterRole_PACKET_BROKER_AGENT          ClusterRole = 12
)

var ClusterRole_name = map[int32]string{
	0:  "NONE",
	1:  "ENTITY_REGISTRY",
	2:  "ACCESS",
	3:  "GATEWAY_SERVER",
	4:  "NETWORK_SERVER",
	5:  "APPLICATION_SERVER",
	6:  "JOIN_SERVER",
	7:  "CRYPTO_SERVER",
	8:  "DEVICE_TEMPLATE_CONVERTER",
	9:  "DEVICE_CLAIMING_SERVER",
	10: "GATEWAY_CONFIGURATION_SERVER",
	11: "QR_CODE_GENERATOR",
	12: "PACKET_BROKER_AGENT",
}

var ClusterRole_value = map[string]int32{
	"NONE":                         0,
	"ENTITY_REGISTRY":              1,
	"ACCESS":                       2,
	"GATEWAY_SERVER":               3,
	"NETWORK_SERVER":               4,
	"APPLICATION_SERVER":           5,
	"JOIN_SERVER":                  6,
	"CRYPTO_SERVER":                7,
	"DEVICE_TEMPLATE_CONVERTER":    8,
	"DEVICE_CLAIMING_SERVER":       9,
	"GATEWAY_CONFIGURATION_SERVER": 10,
	"QR_CODE_GENERATOR":            11,
	"PACKET_BROKER_AGENT":          12,
}

func (ClusterRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e36318a1e2f407cb, []int{2}
}

func init() {
	proto.RegisterEnum("ttn.lorawan.v3.DownlinkPathConstraint", DownlinkPathConstraint_name, DownlinkPathConstraint_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.DownlinkPathConstraint", DownlinkPathConstraint_name, DownlinkPathConstraint_value)
	proto.RegisterEnum("ttn.lorawan.v3.State", State_name, State_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.State", State_name, State_value)
	proto.RegisterEnum("ttn.lorawan.v3.ClusterRole", ClusterRole_name, ClusterRole_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.ClusterRole", ClusterRole_name, ClusterRole_value)
}

func init() { proto.RegisterFile("lorawan-stack/api/enums.proto", fileDescriptor_e36318a1e2f407cb) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/enums.proto", fileDescriptor_e36318a1e2f407cb)
}

var fileDescriptor_e36318a1e2f407cb = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xb1, 0x4f, 0xdb, 0x40,
	0x14, 0xc6, 0xef, 0x42, 0xa0, 0xf4, 0x68, 0xc1, 0x1c, 0x2a, 0x55, 0x51, 0x79, 0x6a, 0x2b, 0x75,
	0x28, 0x12, 0xf1, 0xc0, 0x5f, 0x60, 0xec, 0x87, 0x31, 0x09, 0xb6, 0x39, 0x1f, 0x41, 0xe9, 0x62,
	0x25, 0x28, 0x4d, 0xa2, 0x04, 0x3b, 0x4a, 0x2e, 0xb0, 0x32, 0x32, 0x32, 0x76, 0xac, 0x54, 0x55,
	0xa2, 0x1b, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0xb1, 0x17, 0x46, 0x46, 0xc6, 0x2a, 0x21, 0x29,
	0x62, 0x60, 0xbb, 0xf7, 0xbb, 0xef, 0xbd, 0xf7, 0x7d, 0xd2, 0x1d, 0x5b, 0x6e, 0xc5, 0x9d, 0xf2,
	0x51, 0x39, 0x5a, 0xed, 0xaa, 0xf2, 0x7e, 0x53, 0x2f, 0xb7, 0x1b, 0x7a, 0x35, 0xea, 0x1d, 0x74,
	0x73, 0xed, 0x4e, 0xac, 0x62, 0x3e, 0xab, 0x54, 0x94, 0x1b, 0x49, 0x72, 0x87, 0x6b, 0x4b, 0xab,
	0xb5, 0x86, 0xaa, 0xf7, 0x2a, 0xb9, 0xfd, 0xf8, 0x40, 0xaf, 0xc5, 0xb5, 0x58, 0x1f, 0xca, 0x2a,
	0xbd, 0x1f, 0xc3, 0x6a, 0x58, 0x0c, 0x4f, 0x8f, 0xed, 0x2b, 0xa7, 0x94, 0x2d, 0x5a, 0xf1, 0x51,
	0xd4, 0x6a, 0x44, 0x4d, 0xbf, 0xac, 0xea, 0x66, 0x1c, 0x75, 0x55, 0xa7, 0xdc, 0x88, 0x14, 0xff,
	0xcc, 0x96, 0x2d, 0x6f, 0xcf, 0x2d, 0x38, 0x6e, 0x3e, 0xf4, 0x0d, 0xb9, 0x19, 0x9a, 0x9e, 0x1b,
	0x48, 0x61, 0x38, 0xae, 0x0c, 0x5d, 0xcf, 0x45, 0x8d, 0xf0, 0x6f, 0xec, 0xeb, 0x8b, 0x12, 0x5f,
	0xe0, 0x06, 0x8a, 0xd0, 0x93, 0x9b, 0x28, 0x34, 0xca, 0xbf, 0x30, 0x78, 0x79, 0x1a, 0x16, 0x51,
	0x68, 0x99, 0xa5, 0xec, 0xc9, 0x6f, 0x20, 0x2b, 0x1d, 0x36, 0x19, 0xa8, 0xb2, 0xaa, 0xf2, 0x05,
	0x36, 0x17, 0x48, 0x43, 0x62, 0x28, 0x70, 0x67, 0x17, 0x03, 0x89, 0x96, 0x46, 0x38, 0x67, 0xb3,
	0x8f, 0xd0, 0xf0, 0x7d, 0xe1, 0x15, 0xd1, 0xd2, 0xe8, 0x13, 0x13, 0xb8, 0x85, 0xe6, 0x40, 0x97,
	0xe1, 0xf3, 0xec, 0xed, 0x23, 0xdb, 0x28, 0x18, 0xb6, 0x8d, 0x96, 0x36, 0xf1, 0x34, 0x2f, 0xd8,
	0x0d, 0x7c, 0x74, 0x2d, 0xb4, 0xb4, 0xec, 0x68, 0xe7, 0xdf, 0x0c, 0x9b, 0x31, 0x5b, 0xbd, 0xae,
	0xaa, 0x76, 0x44, 0xdc, 0xaa, 0xf2, 0x69, 0x96, 0x1d, 0x45, 0x5c, 0x60, 0x73, 0xe8, 0x4a, 0x47,
	0x96, 0x42, 0x81, 0xb6, 0x13, 0x48, 0x51, 0xd2, 0x28, 0x67, 0x6c, 0xca, 0x30, 0x4d, 0x0c, 0x02,
	0x2d, 0x33, 0x58, 0x6e, 0x1b, 0x12, 0xf7, 0x8c, 0x52, 0x18, 0xa0, 0x18, 0x04, 0x99, 0x18, 0x30,
	0x17, 0xe5, 0x9e, 0x27, 0xf2, 0x63, 0x96, 0xe5, 0x8b, 0x8c, 0x1b, 0xbe, 0x5f, 0x70, 0x4c, 0x43,
	0x3a, 0x9e, 0x3b, 0xe6, 0x93, 0x7c, 0x8e, 0xcd, 0x6c, 0x79, 0xce, 0x7f, 0x30, 0x35, 0x70, 0x6e,
	0x8a, 0x92, 0x2f, 0xbd, 0x31, 0x7a, 0xc5, 0x97, 0xd9, 0x07, 0x0b, 0x8b, 0x8e, 0x89, 0xa1, 0xc4,
	0x6d, 0xbf, 0x30, 0xc8, 0x60, 0x7a, 0x6e, 0x11, 0x85, 0x44, 0xa1, 0x4d, 0xf3, 0x25, 0xb6, 0x38,
	0xba, 0x36, 0x0b, 0x86, 0xb3, 0xed, 0xb8, 0xf6, 0xb8, 0xf5, 0x35, 0xff, 0xc4, 0x3e, 0x8e, 0xed,
	0x99, 0x9e, 0xbb, 0xe1, 0xd8, 0xbb, 0xe2, 0x99, 0x01, 0xc6, 0xdf, 0xb1, 0xf9, 0x1d, 0x11, 0x9a,
	0x9e, 0x85, 0xa1, 0x8d, 0x2e, 0x0a, 0x43, 0x7a, 0x42, 0x9b, 0xe1, 0xef, 0xd9, 0x82, 0x6f, 0x98,
	0x79, 0x94, 0xe1, 0xba, 0xf0, 0xf2, 0x28, 0x42, 0xc3, 0x46, 0x57, 0x6a, 0x6f, 0xd6, 0xff, 0xd0,
	0xab, 0x3e, 0xd0, 0xeb, 0x3e, 0xd0, 0x9b, 0x3e, 0x90, 0xdb, 0x3e, 0x90, 0xbb, 0x3e, 0x90, 0xfb,
	0x3e, 0x90, 0x87, 0x3e, 0xd0, 0xe3, 0x04, 0xe8, 0x49, 0x02, 0xe4, 0x2c, 0x01, 0x7a, 0x9e, 0x00,
	0xb9, 0x48, 0x80, 0x5c, 0x26, 0x40, 0xae, 0x12, 0xa0, 0xd7, 0x09, 0xd0, 0x9b, 0x04, 0xc8, 0x6d,
	0x02, 0xf4, 0x2e, 0x01, 0x72, 0x9f, 0x00, 0x7d, 0x48, 0x80, 0x1c, 0xa7, 0x40, 0x4e, 0x52, 0xa0,
	0xa7, 0x29, 0x90, 0x9f, 0x29, 0xd0, 0x5f, 0x29, 0x90, 0xb3, 0x14, 0xc8, 0x79, 0x0a, 0xf4, 0x22,
	0x05, 0x7a, 0x99, 0x02, 0xfd, 0xae, 0xd7, 0xe2, 0x9c, 0xaa, 0x57, 0x55, 0xbd, 0x11, 0xd5, 0xba,
	0xb9, 0xa8, 0xaa, 0x8e, 0xe2, 0x4e, 0x53, 0x7f, 0xfe, 0x39, 0x0e, 0xd7, 0xf4, 0x76, 0xb3, 0xa6,
	0x2b, 0x15, 0xb5, 0x2b, 0x95, 0xa9, 0xe1, 0x0b, 0x5f, 0xfb, 0x17, 0x00, 0x00, 0xff, 0xff, 0x43,
	0x49, 0xbe, 0x69, 0x41, 0x03, 0x00, 0x00,
}

func (x DownlinkPathConstraint) String() string {
	s, ok := DownlinkPathConstraint_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x State) String() string {
	s, ok := State_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ClusterRole) String() string {
	s, ok := ClusterRole_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}