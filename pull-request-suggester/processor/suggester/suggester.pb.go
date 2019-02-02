// Code generated by protoc-gen-go. DO NOT EDIT.
// source: suggester/suggester.proto

package suggester

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
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

type GetCommentRequest struct {
	PullRequestURL       string   `protobuf:"bytes,1,opt,name=pullRequestURL,proto3" json:"pullRequestURL,omitempty"`
	PullRequestPatchURL  string   `protobuf:"bytes,2,opt,name=pullRequestPatchURL,proto3" json:"pullRequestPatchURL,omitempty"`
	RepoName             string   `protobuf:"bytes,3,opt,name=repoName,proto3" json:"repoName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCommentRequest) Reset()         { *m = GetCommentRequest{} }
func (m *GetCommentRequest) String() string { return proto.CompactTextString(m) }
func (*GetCommentRequest) ProtoMessage()    {}
func (*GetCommentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_426ac8f376608516, []int{0}
}

func (m *GetCommentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCommentRequest.Unmarshal(m, b)
}
func (m *GetCommentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCommentRequest.Marshal(b, m, deterministic)
}
func (m *GetCommentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCommentRequest.Merge(m, src)
}
func (m *GetCommentRequest) XXX_Size() int {
	return xxx_messageInfo_GetCommentRequest.Size(m)
}
func (m *GetCommentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCommentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCommentRequest proto.InternalMessageInfo

func (m *GetCommentRequest) GetPullRequestURL() string {
	if m != nil {
		return m.PullRequestURL
	}
	return ""
}

func (m *GetCommentRequest) GetPullRequestPatchURL() string {
	if m != nil {
		return m.PullRequestPatchURL
	}
	return ""
}

func (m *GetCommentRequest) GetRepoName() string {
	if m != nil {
		return m.RepoName
	}
	return ""
}

type GetCommentResponse struct {
	PullRequestURL            string                      `protobuf:"bytes,1,opt,name=pullRequestURL,proto3" json:"pullRequestURL,omitempty"`
	FileNameCommitIDPositions []*FileNameCommitIDPosition `protobuf:"bytes,2,rep,name=fileNameCommitIDPositions,proto3" json:"fileNameCommitIDPositions,omitempty"`
	Message                   string                      `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                    `json:"-"`
	XXX_unrecognized          []byte                      `json:"-"`
	XXX_sizecache             int32                       `json:"-"`
}

func (m *GetCommentResponse) Reset()         { *m = GetCommentResponse{} }
func (m *GetCommentResponse) String() string { return proto.CompactTextString(m) }
func (*GetCommentResponse) ProtoMessage()    {}
func (*GetCommentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_426ac8f376608516, []int{1}
}

func (m *GetCommentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCommentResponse.Unmarshal(m, b)
}
func (m *GetCommentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCommentResponse.Marshal(b, m, deterministic)
}
func (m *GetCommentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCommentResponse.Merge(m, src)
}
func (m *GetCommentResponse) XXX_Size() int {
	return xxx_messageInfo_GetCommentResponse.Size(m)
}
func (m *GetCommentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCommentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCommentResponse proto.InternalMessageInfo

func (m *GetCommentResponse) GetPullRequestURL() string {
	if m != nil {
		return m.PullRequestURL
	}
	return ""
}

func (m *GetCommentResponse) GetFileNameCommitIDPositions() []*FileNameCommitIDPosition {
	if m != nil {
		return m.FileNameCommitIDPositions
	}
	return nil
}

func (m *GetCommentResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type FileNameCommitIDPosition struct {
	FileName             string   `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	CommitID             string   `protobuf:"bytes,2,opt,name=commitID,proto3" json:"commitID,omitempty"`
	Position             int32    `protobuf:"varint,3,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileNameCommitIDPosition) Reset()         { *m = FileNameCommitIDPosition{} }
func (m *FileNameCommitIDPosition) String() string { return proto.CompactTextString(m) }
func (*FileNameCommitIDPosition) ProtoMessage()    {}
func (*FileNameCommitIDPosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_426ac8f376608516, []int{2}
}

func (m *FileNameCommitIDPosition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileNameCommitIDPosition.Unmarshal(m, b)
}
func (m *FileNameCommitIDPosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileNameCommitIDPosition.Marshal(b, m, deterministic)
}
func (m *FileNameCommitIDPosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileNameCommitIDPosition.Merge(m, src)
}
func (m *FileNameCommitIDPosition) XXX_Size() int {
	return xxx_messageInfo_FileNameCommitIDPosition.Size(m)
}
func (m *FileNameCommitIDPosition) XXX_DiscardUnknown() {
	xxx_messageInfo_FileNameCommitIDPosition.DiscardUnknown(m)
}

var xxx_messageInfo_FileNameCommitIDPosition proto.InternalMessageInfo

func (m *FileNameCommitIDPosition) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *FileNameCommitIDPosition) GetCommitID() string {
	if m != nil {
		return m.CommitID
	}
	return ""
}

func (m *FileNameCommitIDPosition) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

func init() {
	proto.RegisterType((*GetCommentRequest)(nil), "suggester.service.GetCommentRequest")
	proto.RegisterType((*GetCommentResponse)(nil), "suggester.service.GetCommentResponse")
	proto.RegisterType((*FileNameCommitIDPosition)(nil), "suggester.service.FileNameCommitIDPosition")
}

func init() { proto.RegisterFile("suggester/suggester.proto", fileDescriptor_426ac8f376608516) }

var fileDescriptor_426ac8f376608516 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xdd, 0x4e, 0xf2, 0x40,
	0x10, 0xfd, 0x0a, 0xf9, 0x54, 0x46, 0xa3, 0x61, 0xbd, 0x29, 0x5c, 0x91, 0x46, 0x0d, 0x89, 0xa1,
	0x1a, 0xf4, 0x09, 0xd0, 0xf8, 0x93, 0xa8, 0x69, 0x6a, 0xf4, 0x42, 0xaf, 0x60, 0x19, 0xcb, 0x86,
	0xb6, 0xbb, 0xee, 0x6c, 0x7d, 0x07, 0x1f, 0xcb, 0x37, 0x33, 0x5b, 0xda, 0x42, 0x04, 0xa2, 0x77,
	0x7b, 0x66, 0xce, 0x9e, 0xdd, 0x73, 0x66, 0xa0, 0xa5, 0xb4, 0xe4, 0x48, 0x24, 0xf5, 0x09, 0x65,
	0x51, 0x84, 0x64, 0x50, 0xfb, 0x4a, 0x4b, 0x23, 0x59, 0x73, 0x5e, 0x20, 0xd4, 0x1f, 0x82, 0xa3,
	0xf7, 0xe9, 0x40, 0xf3, 0x1a, 0xcd, 0x85, 0x4c, 0x12, 0x4c, 0x4d, 0x88, 0xef, 0x19, 0x92, 0x61,
	0x47, 0xb0, 0xab, 0xb2, 0x38, 0x2e, 0xe0, 0x53, 0x78, 0xe7, 0x3a, 0x1d, 0xa7, 0xdb, 0x08, 0x7f,
	0x54, 0xd9, 0x29, 0xec, 0x2f, 0x54, 0x82, 0xa1, 0xe1, 0x13, 0x4b, 0xae, 0xe5, 0xe4, 0x55, 0x2d,
	0xd6, 0x86, 0x2d, 0x8d, 0x4a, 0x3e, 0x0c, 0x13, 0x74, 0xeb, 0x39, 0xad, 0xc2, 0xde, 0x97, 0x03,
	0x6c, 0xf1, 0x2f, 0xa4, 0x64, 0x4a, 0xf8, 0xe7, 0xcf, 0x08, 0x68, 0xbd, 0x89, 0x18, 0xad, 0x94,
	0x95, 0x10, 0xe6, 0xf6, 0x32, 0x90, 0x24, 0x8c, 0x90, 0x29, 0xb9, 0xb5, 0x4e, 0xbd, 0xbb, 0xdd,
	0x3f, 0xf6, 0x97, 0x12, 0xf0, 0xaf, 0xd6, 0xdc, 0x09, 0xd7, 0xab, 0x31, 0x17, 0x36, 0x13, 0x24,
	0x1a, 0x46, 0xa5, 0x89, 0x12, 0x7a, 0x29, 0xb8, 0xeb, 0x04, 0xad, 0xf7, 0x52, 0xb2, 0xb0, 0x50,
	0x61, 0xdb, 0xe3, 0x05, 0xbf, 0x88, 0xaf, 0xc2, 0xb6, 0xa7, 0x0a, 0x8d, 0xfc, 0xb9, 0xff, 0x61,
	0x85, 0xfb, 0x53, 0xd8, 0xb9, 0x97, 0x63, 0x2c, 0x73, 0x60, 0xaf, 0x00, 0xf3, 0x08, 0xd9, 0xc1,
	0x0a, 0xbf, 0x4b, 0xd3, 0x6e, 0x1f, 0xfe, 0xc2, 0x9a, 0xcd, 0xc1, 0xfb, 0x37, 0x78, 0x86, 0x73,
	0x2e, 0x13, 0x3f, 0x12, 0x66, 0x92, 0x8d, 0x7c, 0x7b, 0x9c, 0xc8, 0x78, 0x8c, 0xe9, 0xd4, 0x57,
	0x1a, 0xc7, 0x82, 0x9b, 0x9e, 0xd2, 0x3d, 0x3e, 0xbb, 0x46, 0xf3, 0xed, 0x1b, 0xec, 0x3d, 0x96,
	0xc7, 0xc0, 0xee, 0x21, 0xdd, 0x38, 0x81, 0xf3, 0xd2, 0xa8, 0xfa, 0xa3, 0x8d, 0x7c, 0x3d, 0xcf,
	0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x74, 0x3d, 0xdd, 0x36, 0xbb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ModelRequestClient is the client API for ModelRequest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ModelRequestClient interface {
	GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentResponse, error)
}

type modelRequestClient struct {
	cc *grpc.ClientConn
}

func NewModelRequestClient(cc *grpc.ClientConn) ModelRequestClient {
	return &modelRequestClient{cc}
}

func (c *modelRequestClient) GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentResponse, error) {
	out := new(GetCommentResponse)
	err := c.cc.Invoke(ctx, "/suggester.service.ModelRequest/GetComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelRequestServer is the server API for ModelRequest service.
type ModelRequestServer interface {
	GetComment(context.Context, *GetCommentRequest) (*GetCommentResponse, error)
}

func RegisterModelRequestServer(s *grpc.Server, srv ModelRequestServer) {
	s.RegisterService(&_ModelRequest_serviceDesc, srv)
}

func _ModelRequest_GetComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRequestServer).GetComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/suggester.service.ModelRequest/GetComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRequestServer).GetComment(ctx, req.(*GetCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ModelRequest_serviceDesc = grpc.ServiceDesc{
	ServiceName: "suggester.service.ModelRequest",
	HandlerType: (*ModelRequestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetComment",
			Handler:    _ModelRequest_GetComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "suggester/suggester.proto",
}
