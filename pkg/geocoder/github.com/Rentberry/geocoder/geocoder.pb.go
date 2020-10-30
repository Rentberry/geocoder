// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: geocoder.proto

package geocoder

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

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider         string            `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	FormattedAddress string            `protobuf:"bytes,2,opt,name=formatted_address,json=formattedAddress,proto3" json:"formatted_address,omitempty"`
	StreetNumber     string            `protobuf:"bytes,4,opt,name=street_number,json=streetNumber,proto3" json:"street_number,omitempty"`
	StreetName       string            `protobuf:"bytes,5,opt,name=street_name,json=streetName,proto3" json:"street_name,omitempty"`
	Locality         string            `protobuf:"bytes,6,opt,name=locality,proto3" json:"locality,omitempty"`
	Sublocality      string            `protobuf:"bytes,7,opt,name=sublocality,proto3" json:"sublocality,omitempty"`
	Timezone         string            `protobuf:"bytes,8,opt,name=timezone,proto3" json:"timezone,omitempty"`
	PostalCode       string            `protobuf:"bytes,9,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	Id               string            `protobuf:"bytes,20,opt,name=id,proto3" json:"id,omitempty"`
	Country          *Country          `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	State            *State            `protobuf:"bytes,12,opt,name=state,proto3" json:"state,omitempty"`
	LatLng           *LatLng           `protobuf:"bytes,10,opt,name=latLng,proto3" json:"latLng,omitempty"`
	Bounds           *Bounds           `protobuf:"bytes,13,opt,name=bounds,proto3" json:"bounds,omitempty"`
	AdminLevels      []*AdminLevel     `protobuf:"bytes,11,rep,name=admin_levels,json=adminLevels,proto3" json:"admin_levels,omitempty"`
	Components       map[string]string `protobuf:"bytes,21,rep,name=components,proto3" json:"components,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geocoder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_geocoder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_geocoder_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *Location) GetFormattedAddress() string {
	if x != nil {
		return x.FormattedAddress
	}
	return ""
}

func (x *Location) GetStreetNumber() string {
	if x != nil {
		return x.StreetNumber
	}
	return ""
}

func (x *Location) GetStreetName() string {
	if x != nil {
		return x.StreetName
	}
	return ""
}

func (x *Location) GetLocality() string {
	if x != nil {
		return x.Locality
	}
	return ""
}

func (x *Location) GetSublocality() string {
	if x != nil {
		return x.Sublocality
	}
	return ""
}

func (x *Location) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Location) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *Location) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Location) GetCountry() *Country {
	if x != nil {
		return x.Country
	}
	return nil
}

func (x *Location) GetState() *State {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *Location) GetLatLng() *LatLng {
	if x != nil {
		return x.LatLng
	}
	return nil
}

func (x *Location) GetBounds() *Bounds {
	if x != nil {
		return x.Bounds
	}
	return nil
}

func (x *Location) GetAdminLevels() []*AdminLevel {
	if x != nil {
		return x.AdminLevels
	}
	return nil
}

func (x *Location) GetComponents() map[string]string {
	if x != nil {
		return x.Components
	}
	return nil
}

type Bounds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NorthEast *LatLng `protobuf:"bytes,1,opt,name=northEast,proto3" json:"northEast,omitempty"`
	SouthWest *LatLng `protobuf:"bytes,2,opt,name=southWest,proto3" json:"southWest,omitempty"`
}

func (x *Bounds) Reset() {
	*x = Bounds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geocoder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bounds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bounds) ProtoMessage() {}

func (x *Bounds) ProtoReflect() protoreflect.Message {
	mi := &file_geocoder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bounds.ProtoReflect.Descriptor instead.
func (*Bounds) Descriptor() ([]byte, []int) {
	return file_geocoder_proto_rawDescGZIP(), []int{1}
}

func (x *Bounds) GetNorthEast() *LatLng {
	if x != nil {
		return x.NorthEast
	}
	return nil
}

func (x *Bounds) GetSouthWest() *LatLng {
	if x != nil {
		return x.SouthWest
	}
	return nil
}

type AdminLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level int32  `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code  string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *AdminLevel) Reset() {
	*x = AdminLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geocoder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminLevel) ProtoMessage() {}

func (x *AdminLevel) ProtoReflect() protoreflect.Message {
	mi := &file_geocoder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminLevel.ProtoReflect.Descriptor instead.
func (*AdminLevel) Descriptor() ([]byte, []int) {
	return file_geocoder_proto_rawDescGZIP(), []int{2}
}

func (x *AdminLevel) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *AdminLevel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdminLevel) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type Country struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Country) Reset() {
	*x = Country{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geocoder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Country) ProtoMessage() {}

func (x *Country) ProtoReflect() protoreflect.Message {
	mi := &file_geocoder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Country.ProtoReflect.Descriptor instead.
func (*Country) Descriptor() ([]byte, []int) {
	return file_geocoder_proto_rawDescGZIP(), []int{3}
}

func (x *Country) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Country) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geocoder_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_geocoder_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_geocoder_proto_rawDescGZIP(), []int{4}
}

func (x *State) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *State) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type LatLng struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng float64 `protobuf:"fixed64,2,opt,name=lng,proto3" json:"lng,omitempty"`
}

func (x *LatLng) Reset() {
	*x = LatLng{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geocoder_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatLng) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatLng) ProtoMessage() {}

func (x *LatLng) ProtoReflect() protoreflect.Message {
	mi := &file_geocoder_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatLng.ProtoReflect.Descriptor instead.
func (*LatLng) Descriptor() ([]byte, []int) {
	return file_geocoder_proto_rawDescGZIP(), []int{5}
}

func (x *LatLng) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *LatLng) GetLng() float64 {
	if x != nil {
		return x.Lng
	}
	return 0
}

type Timezone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Timezone) Reset() {
	*x = Timezone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geocoder_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timezone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timezone) ProtoMessage() {}

func (x *Timezone) ProtoReflect() protoreflect.Message {
	mi := &file_geocoder_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timezone.ProtoReflect.Descriptor instead.
func (*Timezone) Descriptor() ([]byte, []int) {
	return file_geocoder_proto_rawDescGZIP(), []int{6}
}

func (x *Timezone) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type LocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider string            `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	Address  string            `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Latlng   *LatLng           `protobuf:"bytes,3,opt,name=latlng,proto3" json:"latlng,omitempty"`
	Query    map[string]string `protobuf:"bytes,4,rep,name=query,proto3" json:"query,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *LocationRequest) Reset() {
	*x = LocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geocoder_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationRequest) ProtoMessage() {}

func (x *LocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_geocoder_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationRequest.ProtoReflect.Descriptor instead.
func (*LocationRequest) Descriptor() ([]byte, []int) {
	return file_geocoder_proto_rawDescGZIP(), []int{7}
}

func (x *LocationRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *LocationRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *LocationRequest) GetLatlng() *LatLng {
	if x != nil {
		return x.Latlng
	}
	return nil
}

func (x *LocationRequest) GetQuery() map[string]string {
	if x != nil {
		return x.Query
	}
	return nil
}

type LocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locations []*Location `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
	Exists    bool        `protobuf:"varint,2,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *LocationResponse) Reset() {
	*x = LocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geocoder_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationResponse) ProtoMessage() {}

func (x *LocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_geocoder_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationResponse.ProtoReflect.Descriptor instead.
func (*LocationResponse) Descriptor() ([]byte, []int) {
	return file_geocoder_proto_rawDescGZIP(), []int{8}
}

func (x *LocationResponse) GetLocations() []*Location {
	if x != nil {
		return x.Locations
	}
	return nil
}

func (x *LocationResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

type TimezoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latlng *LatLng `protobuf:"bytes,1,opt,name=latlng,proto3" json:"latlng,omitempty"`
}

func (x *TimezoneRequest) Reset() {
	*x = TimezoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geocoder_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimezoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimezoneRequest) ProtoMessage() {}

func (x *TimezoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_geocoder_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimezoneRequest.ProtoReflect.Descriptor instead.
func (*TimezoneRequest) Descriptor() ([]byte, []int) {
	return file_geocoder_proto_rawDescGZIP(), []int{9}
}

func (x *TimezoneRequest) GetLatlng() *LatLng {
	if x != nil {
		return x.Latlng
	}
	return nil
}

var File_geocoder_proto protoreflect.FileDescriptor

var file_geocoder_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd2, 0x04, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x75, 0x62, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69,
	0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69,
	0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73,
	0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x6c, 0x61, 0x74,
	0x4c, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4c, 0x61, 0x74, 0x4c,
	0x6e, 0x67, 0x52, 0x06, 0x6c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x06, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x42, 0x6f, 0x75,
	0x6e, 0x64, 0x73, 0x52, 0x06, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x12, 0x2e, 0x0a, 0x0c, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x0b,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x56, 0x0a, 0x06, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x12,
	0x25, 0x0a, 0x09, 0x6e, 0x6f, 0x72, 0x74, 0x68, 0x45, 0x61, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52, 0x09, 0x6e, 0x6f, 0x72,
	0x74, 0x68, 0x45, 0x61, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x74, 0x68, 0x57,
	0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4c, 0x61, 0x74, 0x4c,
	0x6e, 0x67, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x74, 0x68, 0x57, 0x65, 0x73, 0x74, 0x22, 0x4a, 0x0a,
	0x0a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x31, 0x0a, 0x07, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x2f, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x2c, 0x0a,
	0x06, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x22, 0x1e, 0x0a, 0x08, 0x54,
	0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xd5, 0x01, 0x0a, 0x0f,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x06, 0x6c, 0x61, 0x74, 0x6c, 0x6e, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52, 0x06,
	0x6c, 0x61, 0x74, 0x6c, 0x6e, 0x67, 0x12, 0x31, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x38, 0x0a, 0x0a, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x53, 0x0a, 0x10, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x22, 0x32, 0x0a, 0x0f, 0x54, 0x69, 0x6d, 0x65,
	0x7a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x6c,
	0x61, 0x74, 0x6c, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4c, 0x61,
	0x74, 0x4c, 0x6e, 0x67, 0x52, 0x06, 0x6c, 0x61, 0x74, 0x6c, 0x6e, 0x67, 0x32, 0x40, 0x0a, 0x0e,
	0x47, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e,
	0x0a, 0x07, 0x47, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x38,
	0x0a, 0x0f, 0x54, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x25, 0x0a, 0x06, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x10, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x42, 0x34, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x6e, 0x74, 0x62, 0x65, 0x72, 0x72, 0x79,
	0x2f, 0x67, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x72, 0xca, 0x02, 0x12, 0x52, 0x65, 0x6e, 0x74,
	0x62, 0x65, 0x72, 0x72, 0x79, 0x5c, 0x47, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_geocoder_proto_rawDescOnce sync.Once
	file_geocoder_proto_rawDescData = file_geocoder_proto_rawDesc
)

func file_geocoder_proto_rawDescGZIP() []byte {
	file_geocoder_proto_rawDescOnce.Do(func() {
		file_geocoder_proto_rawDescData = protoimpl.X.CompressGZIP(file_geocoder_proto_rawDescData)
	})
	return file_geocoder_proto_rawDescData
}

var file_geocoder_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_geocoder_proto_goTypes = []interface{}{
	(*Location)(nil),         // 0: Location
	(*Bounds)(nil),           // 1: Bounds
	(*AdminLevel)(nil),       // 2: AdminLevel
	(*Country)(nil),          // 3: Country
	(*State)(nil),            // 4: State
	(*LatLng)(nil),           // 5: LatLng
	(*Timezone)(nil),         // 6: Timezone
	(*LocationRequest)(nil),  // 7: LocationRequest
	(*LocationResponse)(nil), // 8: LocationResponse
	(*TimezoneRequest)(nil),  // 9: TimezoneRequest
	nil,                      // 10: Location.ComponentsEntry
	nil,                      // 11: LocationRequest.QueryEntry
}
var file_geocoder_proto_depIdxs = []int32{
	3,  // 0: Location.country:type_name -> Country
	4,  // 1: Location.state:type_name -> State
	5,  // 2: Location.latLng:type_name -> LatLng
	1,  // 3: Location.bounds:type_name -> Bounds
	2,  // 4: Location.admin_levels:type_name -> AdminLevel
	10, // 5: Location.components:type_name -> Location.ComponentsEntry
	5,  // 6: Bounds.northEast:type_name -> LatLng
	5,  // 7: Bounds.southWest:type_name -> LatLng
	5,  // 8: LocationRequest.latlng:type_name -> LatLng
	11, // 9: LocationRequest.query:type_name -> LocationRequest.QueryEntry
	0,  // 10: LocationResponse.locations:type_name -> Location
	5,  // 11: TimezoneRequest.latlng:type_name -> LatLng
	7,  // 12: GeocodeService.Geocode:input_type -> LocationRequest
	9,  // 13: TimezoneService.Lookup:input_type -> TimezoneRequest
	8,  // 14: GeocodeService.Geocode:output_type -> LocationResponse
	6,  // 15: TimezoneService.Lookup:output_type -> Timezone
	14, // [14:16] is the sub-list for method output_type
	12, // [12:14] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_geocoder_proto_init() }
func file_geocoder_proto_init() {
	if File_geocoder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_geocoder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_geocoder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bounds); i {
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
		file_geocoder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminLevel); i {
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
		file_geocoder_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Country); i {
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
		file_geocoder_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
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
		file_geocoder_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatLng); i {
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
		file_geocoder_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timezone); i {
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
		file_geocoder_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationRequest); i {
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
		file_geocoder_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationResponse); i {
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
		file_geocoder_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimezoneRequest); i {
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
			RawDescriptor: file_geocoder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_geocoder_proto_goTypes,
		DependencyIndexes: file_geocoder_proto_depIdxs,
		MessageInfos:      file_geocoder_proto_msgTypes,
	}.Build()
	File_geocoder_proto = out.File
	file_geocoder_proto_rawDesc = nil
	file_geocoder_proto_goTypes = nil
	file_geocoder_proto_depIdxs = nil
}
