// Code generated by protoc-gen-go. DO NOT EDIT.
// source: acct.proto

package protos

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Quality Of Service data
type QoS struct {
	DownloadMbps         float32  `protobuf:"fixed32,1,opt,name=download_mbps,json=downloadMbps,proto3" json:"download_mbps,omitempty"`
	UploadMbps           float32  `protobuf:"fixed32,2,opt,name=upload_mbps,json=uploadMbps,proto3" json:"upload_mbps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QoS) Reset()         { *m = QoS{} }
func (m *QoS) String() string { return proto.CompactTextString(m) }
func (*QoS) ProtoMessage()    {}
func (*QoS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9e587f19cf4167b, []int{0}
}

func (m *QoS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QoS.Unmarshal(m, b)
}
func (m *QoS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QoS.Marshal(b, m, deterministic)
}
func (m *QoS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QoS.Merge(m, src)
}
func (m *QoS) XXX_Size() int {
	return xxx_messageInfo_QoS.Size(m)
}
func (m *QoS) XXX_DiscardUnknown() {
	xxx_messageInfo_QoS.DiscardUnknown(m)
}

var xxx_messageInfo_QoS proto.InternalMessageInfo

func (m *QoS) GetDownloadMbps() float32 {
	if m != nil {
		return m.DownloadMbps
	}
	return 0
}

func (m *QoS) GetUploadMbps() float32 {
	if m != nil {
		return m.UploadMbps
	}
	return 0
}

// user session descriptor
type Session struct {
	// user identity
	//
	// Types that are valid to be assigned to User:
	//	*Session_IMSI
	//	*Session_CertificateSerialNumber
	//	*Session_HardwareAddr
	//	*Session_Name
	User isSession_User `protobuf_oneof:"user"`
	// ID of the user network (MNO, Identity Provider, etc. - user network ID)
	ConsumerId string `protobuf:"bytes,5,opt,name=consumer_id,json=consumerId,proto3" json:"consumer_id,omitempty"`
	// unique session ID
	SessionId string `protobuf:"bytes,6,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// ID of the service provider network (serving network ID)
	ProviderId string `protobuf:"bytes,7,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
	// ID/Name of the serving provider Access Point (WiFi AP, LTE enodeb, etc.) - optional
	ProviderApn string `protobuf:"bytes,8,opt,name=provider_apn,json=providerApn,proto3" json:"provider_apn,omitempty"`
	// ID of the service provider gateway  (AGW, CWAG, etc.) - optional
	// Note: this is the ID of Magma access gateway, multiple APs/enodebs may be connected to a single access gateway
	ProviderGatewayId string `protobuf:"bytes,9,opt,name=provider_gateway_id,json=providerGatewayId,proto3" json:"provider_gateway_id,omitempty"`
	// available QoS at the caller's site (optional - can be ignored by SC)
	AvailableQos         *QoS     `protobuf:"bytes,10,opt,name=available_qos,json=availableQos,proto3" json:"available_qos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9e587f19cf4167b, []int{1}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

type isSession_User interface {
	isSession_User()
}

type Session_IMSI struct {
	IMSI string `protobuf:"bytes,1,opt,name=IMSI,proto3,oneof"`
}

type Session_CertificateSerialNumber struct {
	CertificateSerialNumber []byte `protobuf:"bytes,2,opt,name=certificate_serial_number,json=certificateSerialNumber,proto3,oneof"`
}

type Session_HardwareAddr struct {
	HardwareAddr []byte `protobuf:"bytes,3,opt,name=hardware_addr,json=hardwareAddr,proto3,oneof"`
}

type Session_Name struct {
	Name string `protobuf:"bytes,4,opt,name=name,proto3,oneof"`
}

func (*Session_IMSI) isSession_User() {}

func (*Session_CertificateSerialNumber) isSession_User() {}

func (*Session_HardwareAddr) isSession_User() {}

func (*Session_Name) isSession_User() {}

func (m *Session) GetUser() isSession_User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Session) GetIMSI() string {
	if x, ok := m.GetUser().(*Session_IMSI); ok {
		return x.IMSI
	}
	return ""
}

func (m *Session) GetCertificateSerialNumber() []byte {
	if x, ok := m.GetUser().(*Session_CertificateSerialNumber); ok {
		return x.CertificateSerialNumber
	}
	return nil
}

func (m *Session) GetHardwareAddr() []byte {
	if x, ok := m.GetUser().(*Session_HardwareAddr); ok {
		return x.HardwareAddr
	}
	return nil
}

func (m *Session) GetName() string {
	if x, ok := m.GetUser().(*Session_Name); ok {
		return x.Name
	}
	return ""
}

func (m *Session) GetConsumerId() string {
	if m != nil {
		return m.ConsumerId
	}
	return ""
}

func (m *Session) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *Session) GetProviderId() string {
	if m != nil {
		return m.ProviderId
	}
	return ""
}

func (m *Session) GetProviderApn() string {
	if m != nil {
		return m.ProviderApn
	}
	return ""
}

func (m *Session) GetProviderGatewayId() string {
	if m != nil {
		return m.ProviderGatewayId
	}
	return ""
}

func (m *Session) GetAvailableQos() *QoS {
	if m != nil {
		return m.AvailableQos
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Session) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Session_IMSI)(nil),
		(*Session_CertificateSerialNumber)(nil),
		(*Session_HardwareAddr)(nil),
		(*Session_Name)(nil),
	}
}

// start session response
type SessionResp struct {
	ReportingAdvisory *SessionRespReportLimits `protobuf:"bytes,1,opt,name=reporting_advisory,json=reportingAdvisory,proto3" json:"reporting_advisory,omitempty"`
	// minimal required QoS, the session has to be terminated if service provider's site
	// cannot guarantee the requested QoS (optional)
	MinAcceptableQos     *QoS     `protobuf:"bytes,2,opt,name=min_acceptable_qos,json=minAcceptableQos,proto3" json:"min_acceptable_qos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionResp) Reset()         { *m = SessionResp{} }
func (m *SessionResp) String() string { return proto.CompactTextString(m) }
func (*SessionResp) ProtoMessage()    {}
func (*SessionResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9e587f19cf4167b, []int{2}
}

func (m *SessionResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionResp.Unmarshal(m, b)
}
func (m *SessionResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionResp.Marshal(b, m, deterministic)
}
func (m *SessionResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionResp.Merge(m, src)
}
func (m *SessionResp) XXX_Size() int {
	return xxx_messageInfo_SessionResp.Size(m)
}
func (m *SessionResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionResp.DiscardUnknown(m)
}

var xxx_messageInfo_SessionResp proto.InternalMessageInfo

func (m *SessionResp) GetReportingAdvisory() *SessionRespReportLimits {
	if m != nil {
		return m.ReportingAdvisory
	}
	return nil
}

func (m *SessionResp) GetMinAcceptableQos() *QoS {
	if m != nil {
		return m.MinAcceptableQos
	}
	return nil
}

// optional update triggers
// user identity provider will use report_limits to express its update
// frequency preferences. Service provider is encouraged, but not required
// to comply with specified reporting preferences
type SessionRespReportLimits struct {
	// octets_in - trigger update when RX bytes were consumed by the user from the last update event
	// default is 0, no RX trigger
	OctetsIn uint32 `protobuf:"varint,1,opt,name=octets_in,json=octetsIn,proto3" json:"octets_in,omitempty"`
	// octets_out - trigger update when TX bytes were consumed by the user from the last update event
	// default is 0, no TX trigger
	OctetsOut uint32 `protobuf:"varint,2,opt,name=octets_out,json=octetsOut,proto3" json:"octets_out,omitempty"`
	// elapsed_time_sec - trigger update when elapsed_time_sec seconds passed from the last update event
	ElapsedTimeSec       uint32   `protobuf:"varint,3,opt,name=elapsed_time_sec,json=elapsedTimeSec,proto3" json:"elapsed_time_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionRespReportLimits) Reset()         { *m = SessionRespReportLimits{} }
func (m *SessionRespReportLimits) String() string { return proto.CompactTextString(m) }
func (*SessionRespReportLimits) ProtoMessage()    {}
func (*SessionRespReportLimits) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9e587f19cf4167b, []int{2, 0}
}

func (m *SessionRespReportLimits) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionRespReportLimits.Unmarshal(m, b)
}
func (m *SessionRespReportLimits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionRespReportLimits.Marshal(b, m, deterministic)
}
func (m *SessionRespReportLimits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionRespReportLimits.Merge(m, src)
}
func (m *SessionRespReportLimits) XXX_Size() int {
	return xxx_messageInfo_SessionRespReportLimits.Size(m)
}
func (m *SessionRespReportLimits) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionRespReportLimits.DiscardUnknown(m)
}

var xxx_messageInfo_SessionRespReportLimits proto.InternalMessageInfo

func (m *SessionRespReportLimits) GetOctetsIn() uint32 {
	if m != nil {
		return m.OctetsIn
	}
	return 0
}

func (m *SessionRespReportLimits) GetOctetsOut() uint32 {
	if m != nil {
		return m.OctetsOut
	}
	return 0
}

func (m *SessionRespReportLimits) GetElapsedTimeSec() uint32 {
	if m != nil {
		return m.ElapsedTimeSec
	}
	return 0
}

// update_req is relying information on user's ongoing session consumption
type UpdateReq struct {
	// ongoing session information
	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	// octets_in indicates how many octets have been received by the user
	// from the service provider over the course of this session (accumulative)
	// The accumulative nature of this field should compensate for intermittent
	// losses of connectivity
	OctetsIn uint32 `protobuf:"varint,2,opt,name=octets_in,json=octetsIn,proto3" json:"octets_in,omitempty"`
	// octets_out indicates how many octets have been sent on behalf of the user
	// by the service provider over the course of this session (accumulative)
	// The accumulative nature of this field should compensate for intermittent
	// losses of connectivity
	OctetsOut uint32 `protobuf:"varint,3,opt,name=octets_out,json=octetsOut,proto3" json:"octets_out,omitempty"`
	// session_time indicates how many seconds the user/session has received service for
	SessionTime          uint32   `protobuf:"varint,4,opt,name=session_time,json=sessionTime,proto3" json:"session_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateReq) Reset()         { *m = UpdateReq{} }
func (m *UpdateReq) String() string { return proto.CompactTextString(m) }
func (*UpdateReq) ProtoMessage()    {}
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9e587f19cf4167b, []int{3}
}

func (m *UpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateReq.Unmarshal(m, b)
}
func (m *UpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateReq.Marshal(b, m, deterministic)
}
func (m *UpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateReq.Merge(m, src)
}
func (m *UpdateReq) XXX_Size() int {
	return xxx_messageInfo_UpdateReq.Size(m)
}
func (m *UpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateReq proto.InternalMessageInfo

func (m *UpdateReq) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *UpdateReq) GetOctetsIn() uint32 {
	if m != nil {
		return m.OctetsIn
	}
	return 0
}

func (m *UpdateReq) GetOctetsOut() uint32 {
	if m != nil {
		return m.OctetsOut
	}
	return 0
}

func (m *UpdateReq) GetSessionTime() uint32 {
	if m != nil {
		return m.SessionTime
	}
	return 0
}

type StopResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopResp) Reset()         { *m = StopResp{} }
func (m *StopResp) String() string { return proto.CompactTextString(m) }
func (*StopResp) ProtoMessage()    {}
func (*StopResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9e587f19cf4167b, []int{4}
}

func (m *StopResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopResp.Unmarshal(m, b)
}
func (m *StopResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopResp.Marshal(b, m, deterministic)
}
func (m *StopResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopResp.Merge(m, src)
}
func (m *StopResp) XXX_Size() int {
	return xxx_messageInfo_StopResp.Size(m)
}
func (m *StopResp) XXX_DiscardUnknown() {
	xxx_messageInfo_StopResp.DiscardUnknown(m)
}

var xxx_messageInfo_StopResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QoS)(nil), "protos.QoS")
	proto.RegisterType((*Session)(nil), "protos.session")
	proto.RegisterType((*SessionResp)(nil), "protos.session_resp")
	proto.RegisterType((*SessionRespReportLimits)(nil), "protos.session_resp.report_limits")
	proto.RegisterType((*UpdateReq)(nil), "protos.update_req")
	proto.RegisterType((*StopResp)(nil), "protos.stop_resp")
}

func init() { proto.RegisterFile("acct.proto", fileDescriptor_f9e587f19cf4167b) }

var fileDescriptor_f9e587f19cf4167b = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x4b, 0x6f, 0x13, 0x3d,
	0x14, 0x6d, 0x1e, 0x4d, 0x9b, 0x9b, 0xcc, 0xf7, 0xb5, 0xa6, 0x12, 0xa1, 0x08, 0xd1, 0x06, 0x21,
	0x95, 0x45, 0x27, 0x28, 0x20, 0x24, 0x24, 0x36, 0xe9, 0x86, 0x8e, 0x50, 0x41, 0x99, 0xb0, 0x62,
	0x33, 0xf2, 0x8c, 0x4d, 0x6a, 0x31, 0x7e, 0xd4, 0xf6, 0x34, 0xea, 0x3f, 0x60, 0x8f, 0xc4, 0x96,
	0xbf, 0x8a, 0xec, 0x79, 0xa4, 0x0d, 0x85, 0x55, 0xe4, 0x73, 0xcf, 0xb9, 0xf7, 0xdc, 0x47, 0x06,
	0x00, 0x67, 0x99, 0x0d, 0x95, 0x96, 0x56, 0xa2, 0x9e, 0xff, 0x31, 0xe3, 0x0f, 0xd0, 0x99, 0xcb,
	0x05, 0x7a, 0x06, 0x01, 0x91, 0x2b, 0x91, 0x4b, 0x4c, 0x12, 0x9e, 0x2a, 0x33, 0x6a, 0x1d, 0xb5,
	0x4e, 0xda, 0xf1, 0xb0, 0x06, 0x2f, 0x52, 0x65, 0xd0, 0x53, 0x18, 0x14, 0x6a, 0x4d, 0x69, 0x7b,
	0x0a, 0x94, 0x90, 0x23, 0x8c, 0xbf, 0x77, 0x60, 0xc7, 0x50, 0x63, 0x98, 0x14, 0xe8, 0x00, 0xba,
	0xd1, 0xc5, 0x22, 0xf2, 0x89, 0xfa, 0xe7, 0x5b, 0xb1, 0x7f, 0xa1, 0x77, 0xf0, 0x28, 0xa3, 0xda,
	0xb2, 0xaf, 0x2c, 0xc3, 0x96, 0x26, 0x86, 0x6a, 0x86, 0xf3, 0x44, 0x14, 0x3c, 0xa5, 0xda, 0x27,
	0x1c, 0x9e, 0x6f, 0xc5, 0x0f, 0x6f, 0x51, 0x16, 0x9e, 0xf1, 0xd1, 0x13, 0xd0, 0x73, 0x08, 0x2e,
	0xb1, 0x26, 0x2b, 0xac, 0x69, 0x82, 0x09, 0xd1, 0xa3, 0x4e, 0xa5, 0x18, 0xd6, 0xf0, 0x8c, 0x10,
	0xed, 0x4a, 0x0b, 0xcc, 0xe9, 0xa8, 0x5b, 0x97, 0x76, 0x2f, 0xe7, 0x3e, 0x93, 0xc2, 0x14, 0x9c,
	0xea, 0x84, 0x91, 0xd1, 0xb6, 0x0b, 0xc6, 0x50, 0x43, 0x11, 0x41, 0x4f, 0x00, 0x2a, 0xf3, 0x2e,
	0xde, 0xf3, 0xf1, 0x7e, 0x85, 0x44, 0xc4, 0xe9, 0x95, 0x96, 0xd7, 0x8c, 0x94, 0xfa, 0x9d, 0x52,
	0x5f, 0x43, 0x11, 0x41, 0xc7, 0x30, 0x6c, 0x08, 0x58, 0x89, 0xd1, 0xae, 0x67, 0x34, 0xa2, 0x99,
	0x12, 0x28, 0x84, 0x07, 0x0d, 0x65, 0x89, 0x2d, 0x5d, 0xe1, 0x1b, 0x97, 0xab, 0xef, 0x99, 0xfb,
	0x75, 0xe8, 0x7d, 0x19, 0x89, 0x08, 0x7a, 0x09, 0x01, 0xbe, 0xc6, 0x2c, 0xc7, 0x69, 0x4e, 0x93,
	0x2b, 0x69, 0x46, 0x70, 0xd4, 0x3a, 0x19, 0x4c, 0x07, 0xe5, 0x12, 0x4d, 0x38, 0x97, 0x8b, 0x78,
	0xd8, 0x30, 0xe6, 0xd2, 0x9c, 0xf5, 0xa0, 0x5b, 0x18, 0xaa, 0xc7, 0x3f, 0xda, 0x30, 0xac, 0xbb,
	0xd1, 0xd4, 0x28, 0x34, 0x07, 0xa4, 0xa9, 0x92, 0xda, 0x32, 0xb1, 0x4c, 0x30, 0xb9, 0x66, 0x46,
	0xea, 0x1b, 0xbf, 0x9d, 0xc1, 0x74, 0x5c, 0xe7, 0xbb, 0xad, 0x08, 0x4b, 0x7a, 0x92, 0x33, 0xce,
	0xac, 0x89, 0xf7, 0x1b, 0xf5, 0xac, 0x12, 0xa3, 0xb7, 0x80, 0x38, 0x13, 0x09, 0xce, 0x32, 0xaa,
	0x6c, 0x63, 0xb1, 0xfd, 0xa7, 0xc5, 0x3d, 0xce, 0xc4, 0xac, 0x61, 0xcd, 0xa5, 0x39, 0x2c, 0x20,
	0xb8, 0x93, 0x1e, 0x3d, 0x86, 0xbe, 0xcc, 0x2c, 0xb5, 0x26, 0x61, 0xc2, 0xbb, 0x0a, 0xe2, 0xdd,
	0x12, 0x88, 0x84, 0xdb, 0x4c, 0x15, 0x94, 0x85, 0xf5, 0x05, 0x82, 0xb8, 0xa2, 0x7f, 0x2a, 0x2c,
	0x3a, 0x81, 0x3d, 0x9a, 0x63, 0x65, 0x28, 0x49, 0x2c, 0xe3, 0xee, 0xaa, 0x32, 0x7f, 0x19, 0x41,
	0xfc, 0x5f, 0x85, 0x7f, 0x66, 0x9c, 0x2e, 0x68, 0x36, 0xfe, 0xd9, 0x02, 0x28, 0x14, 0x71, 0xa7,
	0xa7, 0xe9, 0x15, 0x7a, 0xd1, 0x9c, 0x6b, 0x35, 0x88, 0xff, 0x37, 0x06, 0x11, 0x37, 0xe7, 0x7c,
	0xc7, 0x5f, 0xfb, 0x9f, 0xfe, 0x3a, 0x9b, 0xfe, 0x8e, 0xd7, 0xab, 0x70, 0xfe, 0xfc, 0x5d, 0x06,
	0xf1, 0xa0, 0xc2, 0x9c, 0xb7, 0xf1, 0x00, 0xfa, 0xc6, 0x4a, 0xe5, 0x07, 0x3f, 0xfd, 0xd5, 0xf2,
	0x7f, 0x55, 0x59, 0x08, 0x37, 0x6e, 0x14, 0xc2, 0xb6, 0xb1, 0x58, 0x5b, 0xb4, 0xe9, 0xee, 0xf0,
	0xe0, 0xbe, 0xbd, 0xa1, 0x29, 0xf4, 0xca, 0x1e, 0x11, 0xaa, 0xe3, 0xeb, 0x9e, 0xff, 0xa2, 0x39,
	0x85, 0xae, 0xab, 0x7f, 0xaf, 0x62, 0xbf, 0x51, 0xd4, 0x0e, 0xcf, 0xde, 0x7c, 0x79, 0xbd, 0x64,
	0xf6, 0xb2, 0x48, 0xc3, 0x4c, 0xf2, 0x09, 0xc7, 0x4b, 0x8e, 0x27, 0xb8, 0x58, 0x72, 0x2a, 0x2c,
	0x25, 0xa7, 0x82, 0xda, 0x95, 0xd4, 0xdf, 0x26, 0xeb, 0x1e, 0x26, 0x65, 0x86, 0xb4, 0xfc, 0xea,
	0xbc, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x6c, 0xe1, 0xe1, 0x8a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AccountingClient is the client API for Accounting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountingClient interface {
	// start will be called at the end of every new user session creation
	// start is responsible for verification & initiation of a smart contract
	// between the user identity provider/MNO and service provider (ISP/WISP/PLTE)
	// A non-error return will indicate successful contract establishment and will
	// result in the beginning of service for the user
	Start(ctx context.Context, in *Session, opts ...grpc.CallOption) (*SessionResp, error)
	// update should be continuously called for every ongoing service session to update
	// the user bandwidth usage as well as current quality of provided service.
	// If update returns error the session should be terminated and the user disconnected,
	// In the case of unsuccessful update completion, service provider is spouse to follow up
	// with final stop call
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*SessionResp, error)
	// stop is a notification call to communicate to identity provider
	// user/network  initiated service termination.
	// stop will provide final used bandwidth count. stop call is issued
	// after the user session was terminated.
	Stop(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*StopResp, error)
}

type accountingClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountingClient(cc grpc.ClientConnInterface) AccountingClient {
	return &accountingClient{cc}
}

func (c *accountingClient) Start(ctx context.Context, in *Session, opts ...grpc.CallOption) (*SessionResp, error) {
	out := new(SessionResp)
	err := c.cc.Invoke(ctx, "/protos.accounting/start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountingClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*SessionResp, error) {
	out := new(SessionResp)
	err := c.cc.Invoke(ctx, "/protos.accounting/update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountingClient) Stop(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*StopResp, error) {
	out := new(StopResp)
	err := c.cc.Invoke(ctx, "/protos.accounting/stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountingServer is the server API for Accounting service.
type AccountingServer interface {
	// start will be called at the end of every new user session creation
	// start is responsible for verification & initiation of a smart contract
	// between the user identity provider/MNO and service provider (ISP/WISP/PLTE)
	// A non-error return will indicate successful contract establishment and will
	// result in the beginning of service for the user
	Start(context.Context, *Session) (*SessionResp, error)
	// update should be continuously called for every ongoing service session to update
	// the user bandwidth usage as well as current quality of provided service.
	// If update returns error the session should be terminated and the user disconnected,
	// In the case of unsuccessful update completion, service provider is spouse to follow up
	// with final stop call
	Update(context.Context, *UpdateReq) (*SessionResp, error)
	// stop is a notification call to communicate to identity provider
	// user/network  initiated service termination.
	// stop will provide final used bandwidth count. stop call is issued
	// after the user session was terminated.
	Stop(context.Context, *UpdateReq) (*StopResp, error)
}

// UnimplementedAccountingServer can be embedded to have forward compatible implementations.
type UnimplementedAccountingServer struct {
}

func (*UnimplementedAccountingServer) Start(ctx context.Context, req *Session) (*SessionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedAccountingServer) Update(ctx context.Context, req *UpdateReq) (*SessionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedAccountingServer) Stop(ctx context.Context, req *UpdateReq) (*StopResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}

func RegisterAccountingServer(s *grpc.Server, srv AccountingServer) {
	s.RegisterService(&_Accounting_serviceDesc, srv)
}

func _Accounting_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.accounting/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServer).Start(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounting_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.accounting/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounting_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.accounting/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServer).Stop(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Accounting_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.accounting",
	HandlerType: (*AccountingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "start",
			Handler:    _Accounting_Start_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Accounting_Update_Handler,
		},
		{
			MethodName: "stop",
			Handler:    _Accounting_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "acct.proto",
}
