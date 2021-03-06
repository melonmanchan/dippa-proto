// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Username             string   `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Room                 string   `protobuf:"bytes,2,opt,name=room" json:"room,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_fc1b8375ff54072c, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetRoom() string {
	if m != nil {
		return m.Room
	}
	return ""
}

type Emotion struct {
	Sadness              float32  `protobuf:"fixed32,1,opt,name=sadness" json:"sadness,omitempty"`
	Joy                  float32  `protobuf:"fixed32,2,opt,name=joy" json:"joy,omitempty"`
	Fear                 float32  `protobuf:"fixed32,3,opt,name=fear" json:"fear,omitempty"`
	Disgust              float32  `protobuf:"fixed32,4,opt,name=disgust" json:"disgust,omitempty"`
	Anger                float32  `protobuf:"fixed32,5,opt,name=anger" json:"anger,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Emotion) Reset()         { *m = Emotion{} }
func (m *Emotion) String() string { return proto.CompactTextString(m) }
func (*Emotion) ProtoMessage()    {}
func (*Emotion) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_fc1b8375ff54072c, []int{1}
}
func (m *Emotion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Emotion.Unmarshal(m, b)
}
func (m *Emotion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Emotion.Marshal(b, m, deterministic)
}
func (dst *Emotion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Emotion.Merge(dst, src)
}
func (m *Emotion) XXX_Size() int {
	return xxx_messageInfo_Emotion.Size(m)
}
func (m *Emotion) XXX_DiscardUnknown() {
	xxx_messageInfo_Emotion.DiscardUnknown(m)
}

var xxx_messageInfo_Emotion proto.InternalMessageInfo

func (m *Emotion) GetSadness() float32 {
	if m != nil {
		return m.Sadness
	}
	return 0
}

func (m *Emotion) GetJoy() float32 {
	if m != nil {
		return m.Joy
	}
	return 0
}

func (m *Emotion) GetFear() float32 {
	if m != nil {
		return m.Fear
	}
	return 0
}

func (m *Emotion) GetDisgust() float32 {
	if m != nil {
		return m.Disgust
	}
	return 0
}

func (m *Emotion) GetAnger() float32 {
	if m != nil {
		return m.Anger
	}
	return 0
}

type GoogleEmotion struct {
	DetectionConfidence  float32  `protobuf:"fixed32,1,opt,name=detection_confidence,json=detectionConfidence" json:"detection_confidence,omitempty"`
	Anger                float32  `protobuf:"fixed32,2,opt,name=anger" json:"anger,omitempty"`
	Blurred              float32  `protobuf:"fixed32,3,opt,name=blurred" json:"blurred,omitempty"`
	Joy                  float32  `protobuf:"fixed32,4,opt,name=joy" json:"joy,omitempty"`
	Sorrow               float32  `protobuf:"fixed32,5,opt,name=sorrow" json:"sorrow,omitempty"`
	Surprise             float32  `protobuf:"fixed32,6,opt,name=surprise" json:"surprise,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoogleEmotion) Reset()         { *m = GoogleEmotion{} }
func (m *GoogleEmotion) String() string { return proto.CompactTextString(m) }
func (*GoogleEmotion) ProtoMessage()    {}
func (*GoogleEmotion) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_fc1b8375ff54072c, []int{2}
}
func (m *GoogleEmotion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoogleEmotion.Unmarshal(m, b)
}
func (m *GoogleEmotion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoogleEmotion.Marshal(b, m, deterministic)
}
func (dst *GoogleEmotion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoogleEmotion.Merge(dst, src)
}
func (m *GoogleEmotion) XXX_Size() int {
	return xxx_messageInfo_GoogleEmotion.Size(m)
}
func (m *GoogleEmotion) XXX_DiscardUnknown() {
	xxx_messageInfo_GoogleEmotion.DiscardUnknown(m)
}

var xxx_messageInfo_GoogleEmotion proto.InternalMessageInfo

func (m *GoogleEmotion) GetDetectionConfidence() float32 {
	if m != nil {
		return m.DetectionConfidence
	}
	return 0
}

func (m *GoogleEmotion) GetAnger() float32 {
	if m != nil {
		return m.Anger
	}
	return 0
}

func (m *GoogleEmotion) GetBlurred() float32 {
	if m != nil {
		return m.Blurred
	}
	return 0
}

func (m *GoogleEmotion) GetJoy() float32 {
	if m != nil {
		return m.Joy
	}
	return 0
}

func (m *GoogleEmotion) GetSorrow() float32 {
	if m != nil {
		return m.Sorrow
	}
	return 0
}

func (m *GoogleEmotion) GetSurprise() float32 {
	if m != nil {
		return m.Surprise
	}
	return 0
}

type ProcessingData struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Contents             []byte   `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessingData) Reset()         { *m = ProcessingData{} }
func (m *ProcessingData) String() string { return proto.CompactTextString(m) }
func (*ProcessingData) ProtoMessage()    {}
func (*ProcessingData) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_fc1b8375ff54072c, []int{3}
}
func (m *ProcessingData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessingData.Unmarshal(m, b)
}
func (m *ProcessingData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessingData.Marshal(b, m, deterministic)
}
func (dst *ProcessingData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessingData.Merge(dst, src)
}
func (m *ProcessingData) XXX_Size() int {
	return xxx_messageInfo_ProcessingData.Size(m)
}
func (m *ProcessingData) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessingData.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessingData proto.InternalMessageInfo

func (m *ProcessingData) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ProcessingData) GetContents() []byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

type WatsonNLP struct {
	Contents             string               `protobuf:"bytes,1,opt,name=contents" json:"contents,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	User                 *User                `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
	Keywords             []*WatsonNLP_Keyword `protobuf:"bytes,4,rep,name=keywords" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *WatsonNLP) Reset()         { *m = WatsonNLP{} }
func (m *WatsonNLP) String() string { return proto.CompactTextString(m) }
func (*WatsonNLP) ProtoMessage()    {}
func (*WatsonNLP) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_fc1b8375ff54072c, []int{4}
}
func (m *WatsonNLP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatsonNLP.Unmarshal(m, b)
}
func (m *WatsonNLP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatsonNLP.Marshal(b, m, deterministic)
}
func (dst *WatsonNLP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatsonNLP.Merge(dst, src)
}
func (m *WatsonNLP) XXX_Size() int {
	return xxx_messageInfo_WatsonNLP.Size(m)
}
func (m *WatsonNLP) XXX_DiscardUnknown() {
	xxx_messageInfo_WatsonNLP.DiscardUnknown(m)
}

var xxx_messageInfo_WatsonNLP proto.InternalMessageInfo

func (m *WatsonNLP) GetContents() string {
	if m != nil {
		return m.Contents
	}
	return ""
}

func (m *WatsonNLP) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *WatsonNLP) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *WatsonNLP) GetKeywords() []*WatsonNLP_Keyword {
	if m != nil {
		return m.Keywords
	}
	return nil
}

type WatsonNLP_Keyword struct {
	Contents             string   `protobuf:"bytes,1,opt,name=contents" json:"contents,omitempty"`
	Sentiment            float32  `protobuf:"fixed32,2,opt,name=sentiment" json:"sentiment,omitempty"`
	Relevance            float32  `protobuf:"fixed32,3,opt,name=relevance" json:"relevance,omitempty"`
	Emotion              *Emotion `protobuf:"bytes,4,opt,name=emotion" json:"emotion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatsonNLP_Keyword) Reset()         { *m = WatsonNLP_Keyword{} }
func (m *WatsonNLP_Keyword) String() string { return proto.CompactTextString(m) }
func (*WatsonNLP_Keyword) ProtoMessage()    {}
func (*WatsonNLP_Keyword) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_fc1b8375ff54072c, []int{4, 0}
}
func (m *WatsonNLP_Keyword) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatsonNLP_Keyword.Unmarshal(m, b)
}
func (m *WatsonNLP_Keyword) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatsonNLP_Keyword.Marshal(b, m, deterministic)
}
func (dst *WatsonNLP_Keyword) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatsonNLP_Keyword.Merge(dst, src)
}
func (m *WatsonNLP_Keyword) XXX_Size() int {
	return xxx_messageInfo_WatsonNLP_Keyword.Size(m)
}
func (m *WatsonNLP_Keyword) XXX_DiscardUnknown() {
	xxx_messageInfo_WatsonNLP_Keyword.DiscardUnknown(m)
}

var xxx_messageInfo_WatsonNLP_Keyword proto.InternalMessageInfo

func (m *WatsonNLP_Keyword) GetContents() string {
	if m != nil {
		return m.Contents
	}
	return ""
}

func (m *WatsonNLP_Keyword) GetSentiment() float32 {
	if m != nil {
		return m.Sentiment
	}
	return 0
}

func (m *WatsonNLP_Keyword) GetRelevance() float32 {
	if m != nil {
		return m.Relevance
	}
	return 0
}

func (m *WatsonNLP_Keyword) GetEmotion() *Emotion {
	if m != nil {
		return m.Emotion
	}
	return nil
}

type GoogleFacialRecognition struct {
	Emotion              *GoogleEmotion       `protobuf:"bytes,1,opt,name=emotion" json:"emotion,omitempty"`
	Image                []byte               `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	User                 *User                `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GoogleFacialRecognition) Reset()         { *m = GoogleFacialRecognition{} }
func (m *GoogleFacialRecognition) String() string { return proto.CompactTextString(m) }
func (*GoogleFacialRecognition) ProtoMessage()    {}
func (*GoogleFacialRecognition) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_fc1b8375ff54072c, []int{5}
}
func (m *GoogleFacialRecognition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoogleFacialRecognition.Unmarshal(m, b)
}
func (m *GoogleFacialRecognition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoogleFacialRecognition.Marshal(b, m, deterministic)
}
func (dst *GoogleFacialRecognition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoogleFacialRecognition.Merge(dst, src)
}
func (m *GoogleFacialRecognition) XXX_Size() int {
	return xxx_messageInfo_GoogleFacialRecognition.Size(m)
}
func (m *GoogleFacialRecognition) XXX_DiscardUnknown() {
	xxx_messageInfo_GoogleFacialRecognition.DiscardUnknown(m)
}

var xxx_messageInfo_GoogleFacialRecognition proto.InternalMessageInfo

func (m *GoogleFacialRecognition) GetEmotion() *GoogleEmotion {
	if m != nil {
		return m.Emotion
	}
	return nil
}

func (m *GoogleFacialRecognition) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *GoogleFacialRecognition) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *GoogleFacialRecognition) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "types.User")
	proto.RegisterType((*Emotion)(nil), "types.Emotion")
	proto.RegisterType((*GoogleEmotion)(nil), "types.GoogleEmotion")
	proto.RegisterType((*ProcessingData)(nil), "types.ProcessingData")
	proto.RegisterType((*WatsonNLP)(nil), "types.WatsonNLP")
	proto.RegisterType((*WatsonNLP_Keyword)(nil), "types.WatsonNLP.Keyword")
	proto.RegisterType((*GoogleFacialRecognition)(nil), "types.GoogleFacialRecognition")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_types_fc1b8375ff54072c) }

var fileDescriptor_types_fc1b8375ff54072c = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcb, 0x8a, 0xd4, 0x40,
	0x14, 0x25, 0x8f, 0xee, 0x9e, 0xbe, 0xd1, 0x41, 0xca, 0x46, 0x43, 0x23, 0xcc, 0x90, 0x55, 0xaf,
	0x32, 0xd8, 0x8a, 0xe0, 0x52, 0x7c, 0x2d, 0x7c, 0x30, 0x04, 0xc5, 0xe5, 0x50, 0x9d, 0xdc, 0x0e,
	0xd1, 0x4e, 0x55, 0x53, 0x55, 0x71, 0xe8, 0x6f, 0xf0, 0x87, 0x04, 0x3f, 0xc0, 0xdf, 0x92, 0xba,
	0x55, 0x49, 0x1c, 0xa1, 0x5d, 0xcc, 0xae, 0xce, 0x7d, 0x1c, 0xce, 0xad, 0x73, 0x20, 0x31, 0x87,
	0x3d, 0xea, 0x7c, 0xaf, 0xa4, 0x91, 0x6c, 0x42, 0x60, 0x79, 0x56, 0x4b, 0x59, 0xef, 0xf0, 0x82,
	0x8a, 0x9b, 0x6e, 0x7b, 0x61, 0x9a, 0x16, 0xb5, 0xe1, 0xed, 0xde, 0xcd, 0x65, 0xcf, 0x20, 0xfe,
	0xac, 0x51, 0xb1, 0x25, 0x9c, 0x74, 0x1a, 0x95, 0xe0, 0x2d, 0xa6, 0xc1, 0x79, 0xb0, 0x9a, 0x17,
	0x03, 0x66, 0x0c, 0x62, 0x25, 0x65, 0x9b, 0x86, 0x54, 0xa7, 0x77, 0x76, 0x80, 0xd9, 0xeb, 0x56,
	0x9a, 0x46, 0x0a, 0x96, 0xc2, 0x4c, 0xf3, 0x4a, 0xa0, 0xd6, 0xb4, 0x19, 0x16, 0x3d, 0x64, 0xf7,
	0x20, 0xfa, 0x2a, 0x0f, 0xb4, 0x17, 0x16, 0xf6, 0x69, 0xa9, 0xb6, 0xc8, 0x55, 0x1a, 0x51, 0x89,
	0xde, 0x76, 0xbf, 0x6a, 0x74, 0xdd, 0x69, 0x93, 0xc6, 0x6e, 0xdf, 0x43, 0xb6, 0x80, 0x09, 0x17,
	0x35, 0xaa, 0x74, 0x42, 0x75, 0x07, 0xb2, 0x9f, 0x01, 0xdc, 0x7d, 0x4b, 0x67, 0xf5, 0x0a, 0x1e,
	0xc3, 0xa2, 0x42, 0x83, 0xa5, 0x05, 0x57, 0xa5, 0x14, 0xdb, 0xa6, 0x42, 0x51, 0xa2, 0x97, 0x73,
	0x7f, 0xe8, 0xbd, 0x1c, 0x5a, 0x23, 0x75, 0xf8, 0x17, 0xb5, 0x95, 0xb2, 0xd9, 0x75, 0x4a, 0x61,
	0xe5, 0x15, 0xf6, 0xb0, 0x3f, 0x25, 0x1e, 0x4f, 0x79, 0x00, 0x53, 0x2d, 0x95, 0x92, 0xd7, 0x5e,
	0x9d, 0x47, 0xf6, 0x27, 0x75, 0xa7, 0xf6, 0xaa, 0xd1, 0x98, 0x4e, 0xa9, 0x33, 0xe0, 0xec, 0x03,
	0x9c, 0x5e, 0x2a, 0x59, 0xa2, 0xd6, 0x8d, 0xa8, 0x5f, 0x71, 0xc3, 0xd9, 0x19, 0xc4, 0xf6, 0x9f,
	0x49, 0x6a, 0xb2, 0x4e, 0x72, 0xe7, 0xa1, 0xb5, 0xa4, 0xa0, 0x86, 0xa5, 0x2b, 0xa5, 0x30, 0x28,
	0x8c, 0x26, 0xad, 0x77, 0x8a, 0x01, 0x67, 0xbf, 0x43, 0x98, 0x7f, 0xe1, 0x46, 0x4b, 0xf1, 0xf1,
	0xfd, 0xe5, 0x8d, 0x49, 0x6f, 0x61, 0x8f, 0xd9, 0x73, 0x80, 0x52, 0x21, 0x37, 0x58, 0x5d, 0x71,
	0x43, 0x3c, 0xc9, 0x7a, 0x99, 0xbb, 0x70, 0xe4, 0x7d, 0x38, 0xf2, 0x4f, 0x7d, 0x38, 0x8a, 0xb9,
	0x9f, 0x7e, 0x61, 0x06, 0x85, 0xd1, 0x31, 0x85, 0x4f, 0xe1, 0xe4, 0x1b, 0x1e, 0xae, 0xa5, 0xaa,
	0x74, 0x1a, 0x9f, 0x47, 0xab, 0x64, 0x9d, 0xfa, 0xa1, 0x41, 0x5b, 0xfe, 0xce, 0x0d, 0x14, 0xc3,
	0xe4, 0xf2, 0x47, 0x00, 0x33, 0x5f, 0xfd, 0xaf, 0xf2, 0x47, 0x30, 0xd7, 0x28, 0x6c, 0x6c, 0x85,
	0xf1, 0x66, 0x8d, 0x05, 0xdb, 0x55, 0xb8, 0xc3, 0xef, 0xdc, 0xda, 0xed, 0x2c, 0x1b, 0x0b, 0x6c,
	0x05, 0x33, 0x74, 0x11, 0x21, 0xe3, 0x92, 0xf5, 0xa9, 0x17, 0xe6, 0x83, 0x53, 0xf4, 0xed, 0xec,
	0x57, 0x00, 0x0f, 0x5d, 0xa6, 0xde, 0xf0, 0xb2, 0xe1, 0xbb, 0x02, 0x4b, 0x59, 0x8b, 0x86, 0xd2,
	0x95, 0x8f, 0x2c, 0xce, 0xa5, 0x85, 0x67, 0xb9, 0x11, 0xc2, 0x81, 0xcb, 0x46, 0xab, 0x69, 0x79,
	0x8d, 0xde, 0x2e, 0x07, 0xfe, 0x71, 0x20, 0xba, 0x8d, 0x03, 0xf1, 0x11, 0x07, 0x36, 0x53, 0xda,
	0x7f, 0xf2, 0x27, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xdd, 0xec, 0x12, 0x02, 0x04, 0x00, 0x00,
}
