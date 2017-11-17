// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/ml/v1beta1/project_service.proto

package ml

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Requests service account information associated with a project.
type GetConfigRequest struct {
	// Required. The project name.
	//
	// Authorization: requires `Viewer` role on the specified project.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetConfigRequest) Reset()                    { *m = GetConfigRequest{} }
func (m *GetConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*GetConfigRequest) ProtoMessage()               {}
func (*GetConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *GetConfigRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Returns service account information associated with a project.
type GetConfigResponse struct {
	// The service account Cloud ML uses to access resources in the project.
	ServiceAccount string `protobuf:"bytes,1,opt,name=service_account,json=serviceAccount" json:"service_account,omitempty"`
	// The project number for `service_account`.
	ServiceAccountProject int64 `protobuf:"varint,2,opt,name=service_account_project,json=serviceAccountProject" json:"service_account_project,omitempty"`
}

func (m *GetConfigResponse) Reset()                    { *m = GetConfigResponse{} }
func (m *GetConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*GetConfigResponse) ProtoMessage()               {}
func (*GetConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *GetConfigResponse) GetServiceAccount() string {
	if m != nil {
		return m.ServiceAccount
	}
	return ""
}

func (m *GetConfigResponse) GetServiceAccountProject() int64 {
	if m != nil {
		return m.ServiceAccountProject
	}
	return 0
}

func init() {
	proto.RegisterType((*GetConfigRequest)(nil), "google.cloud.ml.v1beta1.GetConfigRequest")
	proto.RegisterType((*GetConfigResponse)(nil), "google.cloud.ml.v1beta1.GetConfigResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ProjectManagementService service

type ProjectManagementServiceClient interface {
	// Get the service account information associated with your project. You need
	// this information in order to grant the service account persmissions for
	// the Google Cloud Storage location where you put your model training code
	// for training the model with Google Cloud Machine Learning.
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error)
}

type projectManagementServiceClient struct {
	cc *grpc.ClientConn
}

func NewProjectManagementServiceClient(cc *grpc.ClientConn) ProjectManagementServiceClient {
	return &projectManagementServiceClient{cc}
}

func (c *projectManagementServiceClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	out := new(GetConfigResponse)
	err := grpc.Invoke(ctx, "/google.cloud.ml.v1beta1.ProjectManagementService/GetConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProjectManagementService service

type ProjectManagementServiceServer interface {
	// Get the service account information associated with your project. You need
	// this information in order to grant the service account persmissions for
	// the Google Cloud Storage location where you put your model training code
	// for training the model with Google Cloud Machine Learning.
	GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error)
}

func RegisterProjectManagementServiceServer(s *grpc.Server, srv ProjectManagementServiceServer) {
	s.RegisterService(&_ProjectManagementService_serviceDesc, srv)
}

func _ProjectManagementService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectManagementServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.ml.v1beta1.ProjectManagementService/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectManagementServiceServer).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProjectManagementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.ml.v1beta1.ProjectManagementService",
	HandlerType: (*ProjectManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _ProjectManagementService_GetConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/ml/v1beta1/project_service.proto",
}

func init() { proto.RegisterFile("google/cloud/ml/v1beta1/project_service.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4a, 0x43, 0x31,
	0x10, 0xc6, 0x79, 0x55, 0x84, 0x66, 0xe1, 0x9f, 0x88, 0xb4, 0x14, 0xc1, 0x52, 0xa4, 0xd6, 0xa2,
	0x09, 0x55, 0x10, 0x54, 0x5c, 0x58, 0x17, 0xae, 0x84, 0x52, 0x77, 0x6e, 0x4a, 0xfa, 0x1c, 0xc3,
	0x93, 0x24, 0x13, 0x5f, 0xd2, 0x6e, 0xc4, 0x8d, 0x27, 0x10, 0x3c, 0x87, 0xa7, 0xf1, 0x0a, 0x1e,
	0x44, 0xfa, 0x92, 0x16, 0x6d, 0x11, 0xdc, 0x0d, 0x33, 0xbf, 0x6f, 0x32, 0xdf, 0x4c, 0xc8, 0xa1,
	0x44, 0x94, 0x0a, 0x78, 0xaa, 0x70, 0x74, 0xcf, 0xb5, 0xe2, 0xe3, 0xce, 0x10, 0xbc, 0xe8, 0x70,
	0x9b, 0xe3, 0x23, 0xa4, 0x7e, 0xe0, 0x20, 0x1f, 0x67, 0x29, 0x30, 0x9b, 0xa3, 0x47, 0x5a, 0x09,
	0x38, 0x2b, 0x70, 0xa6, 0x15, 0x8b, 0x78, 0x6d, 0x3b, 0xf6, 0x11, 0x36, 0xe3, 0xc2, 0x18, 0xf4,
	0xc2, 0x67, 0x68, 0x5c, 0x90, 0x35, 0x9a, 0x64, 0xfd, 0x1a, 0xfc, 0x15, 0x9a, 0x87, 0x4c, 0xf6,
	0xe1, 0x69, 0x04, 0xce, 0x53, 0x4a, 0x96, 0x8d, 0xd0, 0x50, 0x4d, 0xea, 0x49, 0xab, 0xdc, 0x2f,
	0xe2, 0x86, 0x27, 0x1b, 0x3f, 0x38, 0x67, 0xd1, 0x38, 0xa0, 0x7b, 0x64, 0x2d, 0x0e, 0x31, 0x10,
	0x69, 0x8a, 0x23, 0xe3, 0xa3, 0x66, 0x35, 0xa6, 0x2f, 0x43, 0x96, 0x9e, 0x90, 0xca, 0x1c, 0x38,
	0x88, 0x2e, 0xaa, 0xa5, 0x7a, 0xd2, 0x5a, 0xea, 0x6f, 0xfd, 0x16, 0xf4, 0x42, 0xf1, 0xe8, 0x23,
	0x21, 0xd5, 0x18, 0xdf, 0x08, 0x23, 0x24, 0x68, 0x30, 0xfe, 0x36, 0xa0, 0xf4, 0x2d, 0x21, 0xe5,
	0xd9, 0x4c, 0x74, 0x9f, 0xfd, 0xb1, 0x00, 0x36, 0xef, 0xaf, 0xd6, 0xfe, 0x0f, 0x1a, 0x2c, 0x36,
	0x0e, 0x5e, 0x3f, 0xbf, 0xde, 0x4b, 0x4d, 0xba, 0x3b, 0x5b, 0xff, 0xf3, 0x64, 0x1f, 0x17, 0x71,
	0x7c, 0xc7, 0xdb, 0x2f, 0x67, 0x72, 0xaa, 0xea, 0x3a, 0xb2, 0x93, 0xa2, 0x5e, 0x68, 0x2f, 0x6c,
	0x36, 0x7d, 0xa2, 0xbb, 0x19, 0xfd, 0x44, 0x17, 0xbd, 0xc9, 0x15, 0x7a, 0xc9, 0xdd, 0x69, 0xd4,
	0x48, 0x54, 0xc2, 0x48, 0x86, 0xb9, 0xe4, 0x12, 0x4c, 0x71, 0x23, 0x1e, 0x4a, 0xc2, 0x66, 0x6e,
	0xe1, 0x33, 0x9c, 0x6b, 0x35, 0x5c, 0x29, 0xa8, 0xe3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x76,
	0x59, 0xc4, 0x91, 0x31, 0x02, 0x00, 0x00,
}
