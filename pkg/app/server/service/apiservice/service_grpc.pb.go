// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package apiservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// APIServiceClient is the client API for APIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIServiceClient interface {
	AddApplication(ctx context.Context, in *AddApplicationRequest, opts ...grpc.CallOption) (*AddApplicationResponse, error)
	SyncApplication(ctx context.Context, in *SyncApplicationRequest, opts ...grpc.CallOption) (*SyncApplicationResponse, error)
	GetApplication(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*GetApplicationResponse, error)
	ListApplications(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ListApplicationsResponse, error)
	RenameApplicationConfigFile(ctx context.Context, in *RenameApplicationConfigFileRequest, opts ...grpc.CallOption) (*RenameApplicationConfigFileResponse, error)
	GetDeployment(ctx context.Context, in *GetDeploymentRequest, opts ...grpc.CallOption) (*GetDeploymentResponse, error)
	GetCommand(ctx context.Context, in *GetCommandRequest, opts ...grpc.CallOption) (*GetCommandResponse, error)
	EnablePiped(ctx context.Context, in *EnablePipedRequest, opts ...grpc.CallOption) (*EnablePipedResponse, error)
	DisablePiped(ctx context.Context, in *DisablePipedRequest, opts ...grpc.CallOption) (*DisablePipedResponse, error)
	RegisterEvent(ctx context.Context, in *RegisterEventRequest, opts ...grpc.CallOption) (*RegisterEventResponse, error)
	RequestPlanPreview(ctx context.Context, in *RequestPlanPreviewRequest, opts ...grpc.CallOption) (*RequestPlanPreviewResponse, error)
	GetPlanPreviewResults(ctx context.Context, in *GetPlanPreviewResultsRequest, opts ...grpc.CallOption) (*GetPlanPreviewResultsResponse, error)
	Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error)
}

type aPIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIServiceClient(cc grpc.ClientConnInterface) APIServiceClient {
	return &aPIServiceClient{cc}
}

func (c *aPIServiceClient) AddApplication(ctx context.Context, in *AddApplicationRequest, opts ...grpc.CallOption) (*AddApplicationResponse, error) {
	out := new(AddApplicationResponse)
	err := c.cc.Invoke(ctx, "/grpc.service.apiservice.APIService/AddApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) SyncApplication(ctx context.Context, in *SyncApplicationRequest, opts ...grpc.CallOption) (*SyncApplicationResponse, error) {
	out := new(SyncApplicationResponse)
	err := c.cc.Invoke(ctx, "/grpc.service.apiservice.APIService/SyncApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) GetApplication(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*GetApplicationResponse, error) {
	out := new(GetApplicationResponse)
	err := c.cc.Invoke(ctx, "/grpc.service.apiservice.APIService/GetApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) ListApplications(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ListApplicationsResponse, error) {
	out := new(ListApplicationsResponse)
	err := c.cc.Invoke(ctx, "/grpc.service.apiservice.APIService/ListApplications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) RenameApplicationConfigFile(ctx context.Context, in *RenameApplicationConfigFileRequest, opts ...grpc.CallOption) (*RenameApplicationConfigFileResponse, error) {
	out := new(RenameApplicationConfigFileResponse)
	err := c.cc.Invoke(ctx, "/grpc.service.apiservice.APIService/RenameApplicationConfigFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) GetDeployment(ctx context.Context, in *GetDeploymentRequest, opts ...grpc.CallOption) (*GetDeploymentResponse, error) {
	out := new(GetDeploymentResponse)
	err := c.cc.Invoke(ctx, "/grpc.service.apiservice.APIService/GetDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) GetCommand(ctx context.Context, in *GetCommandRequest, opts ...grpc.CallOption) (*GetCommandResponse, error) {
	out := new(GetCommandResponse)
	err := c.cc.Invoke(ctx, "/grpc.service.apiservice.APIService/GetCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) EnablePiped(ctx context.Context, in *EnablePipedRequest, opts ...grpc.CallOption) (*EnablePipedResponse, error) {
	out := new(EnablePipedResponse)
	err := c.cc.Invoke(ctx, "/grpc.service.apiservice.APIService/EnablePiped", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) DisablePiped(ctx context.Context, in *DisablePipedRequest, opts ...grpc.CallOption) (*DisablePipedResponse, error) {
	out := new(DisablePipedResponse)
	err := c.cc.Invoke(ctx, "/grpc.service.apiservice.APIService/DisablePiped", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) RegisterEvent(ctx context.Context, in *RegisterEventRequest, opts ...grpc.CallOption) (*RegisterEventResponse, error) {
	out := new(RegisterEventResponse)
	err := c.cc.Invoke(ctx, "/grpc.service.apiservice.APIService/RegisterEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) RequestPlanPreview(ctx context.Context, in *RequestPlanPreviewRequest, opts ...grpc.CallOption) (*RequestPlanPreviewResponse, error) {
	out := new(RequestPlanPreviewResponse)
	err := c.cc.Invoke(ctx, "/grpc.service.apiservice.APIService/RequestPlanPreview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) GetPlanPreviewResults(ctx context.Context, in *GetPlanPreviewResultsRequest, opts ...grpc.CallOption) (*GetPlanPreviewResultsResponse, error) {
	out := new(GetPlanPreviewResultsResponse)
	err := c.cc.Invoke(ctx, "/grpc.service.apiservice.APIService/GetPlanPreviewResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error) {
	out := new(EncryptResponse)
	err := c.cc.Invoke(ctx, "/grpc.service.apiservice.APIService/Encrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServiceServer is the server API for APIService service.
// All implementations must embed UnimplementedAPIServiceServer
// for forward compatibility
type APIServiceServer interface {
	AddApplication(context.Context, *AddApplicationRequest) (*AddApplicationResponse, error)
	SyncApplication(context.Context, *SyncApplicationRequest) (*SyncApplicationResponse, error)
	GetApplication(context.Context, *GetApplicationRequest) (*GetApplicationResponse, error)
	ListApplications(context.Context, *ListApplicationsRequest) (*ListApplicationsResponse, error)
	RenameApplicationConfigFile(context.Context, *RenameApplicationConfigFileRequest) (*RenameApplicationConfigFileResponse, error)
	GetDeployment(context.Context, *GetDeploymentRequest) (*GetDeploymentResponse, error)
	GetCommand(context.Context, *GetCommandRequest) (*GetCommandResponse, error)
	EnablePiped(context.Context, *EnablePipedRequest) (*EnablePipedResponse, error)
	DisablePiped(context.Context, *DisablePipedRequest) (*DisablePipedResponse, error)
	RegisterEvent(context.Context, *RegisterEventRequest) (*RegisterEventResponse, error)
	RequestPlanPreview(context.Context, *RequestPlanPreviewRequest) (*RequestPlanPreviewResponse, error)
	GetPlanPreviewResults(context.Context, *GetPlanPreviewResultsRequest) (*GetPlanPreviewResultsResponse, error)
	Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error)
	mustEmbedUnimplementedAPIServiceServer()
}

// UnimplementedAPIServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAPIServiceServer struct {
}

func (UnimplementedAPIServiceServer) AddApplication(context.Context, *AddApplicationRequest) (*AddApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddApplication not implemented")
}
func (UnimplementedAPIServiceServer) SyncApplication(context.Context, *SyncApplicationRequest) (*SyncApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncApplication not implemented")
}
func (UnimplementedAPIServiceServer) GetApplication(context.Context, *GetApplicationRequest) (*GetApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplication not implemented")
}
func (UnimplementedAPIServiceServer) ListApplications(context.Context, *ListApplicationsRequest) (*ListApplicationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApplications not implemented")
}
func (UnimplementedAPIServiceServer) RenameApplicationConfigFile(context.Context, *RenameApplicationConfigFileRequest) (*RenameApplicationConfigFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameApplicationConfigFile not implemented")
}
func (UnimplementedAPIServiceServer) GetDeployment(context.Context, *GetDeploymentRequest) (*GetDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeployment not implemented")
}
func (UnimplementedAPIServiceServer) GetCommand(context.Context, *GetCommandRequest) (*GetCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommand not implemented")
}
func (UnimplementedAPIServiceServer) EnablePiped(context.Context, *EnablePipedRequest) (*EnablePipedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnablePiped not implemented")
}
func (UnimplementedAPIServiceServer) DisablePiped(context.Context, *DisablePipedRequest) (*DisablePipedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisablePiped not implemented")
}
func (UnimplementedAPIServiceServer) RegisterEvent(context.Context, *RegisterEventRequest) (*RegisterEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterEvent not implemented")
}
func (UnimplementedAPIServiceServer) RequestPlanPreview(context.Context, *RequestPlanPreviewRequest) (*RequestPlanPreviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestPlanPreview not implemented")
}
func (UnimplementedAPIServiceServer) GetPlanPreviewResults(context.Context, *GetPlanPreviewResultsRequest) (*GetPlanPreviewResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlanPreviewResults not implemented")
}
func (UnimplementedAPIServiceServer) Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}
func (UnimplementedAPIServiceServer) mustEmbedUnimplementedAPIServiceServer() {}

// UnsafeAPIServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIServiceServer will
// result in compilation errors.
type UnsafeAPIServiceServer interface {
	mustEmbedUnimplementedAPIServiceServer()
}

func RegisterAPIServiceServer(s grpc.ServiceRegistrar, srv APIServiceServer) {
	s.RegisterService(&APIService_ServiceDesc, srv)
}

func _APIService_AddApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).AddApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.service.apiservice.APIService/AddApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).AddApplication(ctx, req.(*AddApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_SyncApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).SyncApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.service.apiservice.APIService/SyncApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).SyncApplication(ctx, req.(*SyncApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_GetApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).GetApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.service.apiservice.APIService/GetApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).GetApplication(ctx, req.(*GetApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_ListApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).ListApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.service.apiservice.APIService/ListApplications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).ListApplications(ctx, req.(*ListApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_RenameApplicationConfigFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameApplicationConfigFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).RenameApplicationConfigFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.service.apiservice.APIService/RenameApplicationConfigFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).RenameApplicationConfigFile(ctx, req.(*RenameApplicationConfigFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_GetDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).GetDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.service.apiservice.APIService/GetDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).GetDeployment(ctx, req.(*GetDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_GetCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).GetCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.service.apiservice.APIService/GetCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).GetCommand(ctx, req.(*GetCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_EnablePiped_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnablePipedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).EnablePiped(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.service.apiservice.APIService/EnablePiped",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).EnablePiped(ctx, req.(*EnablePipedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_DisablePiped_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisablePipedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).DisablePiped(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.service.apiservice.APIService/DisablePiped",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).DisablePiped(ctx, req.(*DisablePipedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_RegisterEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).RegisterEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.service.apiservice.APIService/RegisterEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).RegisterEvent(ctx, req.(*RegisterEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_RequestPlanPreview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPlanPreviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).RequestPlanPreview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.service.apiservice.APIService/RequestPlanPreview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).RequestPlanPreview(ctx, req.(*RequestPlanPreviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_GetPlanPreviewResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlanPreviewResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).GetPlanPreviewResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.service.apiservice.APIService/GetPlanPreviewResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).GetPlanPreviewResults(ctx, req.(*GetPlanPreviewResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.service.apiservice.APIService/Encrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).Encrypt(ctx, req.(*EncryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// APIService_ServiceDesc is the grpc.ServiceDesc for APIService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var APIService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.service.apiservice.APIService",
	HandlerType: (*APIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddApplication",
			Handler:    _APIService_AddApplication_Handler,
		},
		{
			MethodName: "SyncApplication",
			Handler:    _APIService_SyncApplication_Handler,
		},
		{
			MethodName: "GetApplication",
			Handler:    _APIService_GetApplication_Handler,
		},
		{
			MethodName: "ListApplications",
			Handler:    _APIService_ListApplications_Handler,
		},
		{
			MethodName: "RenameApplicationConfigFile",
			Handler:    _APIService_RenameApplicationConfigFile_Handler,
		},
		{
			MethodName: "GetDeployment",
			Handler:    _APIService_GetDeployment_Handler,
		},
		{
			MethodName: "GetCommand",
			Handler:    _APIService_GetCommand_Handler,
		},
		{
			MethodName: "EnablePiped",
			Handler:    _APIService_EnablePiped_Handler,
		},
		{
			MethodName: "DisablePiped",
			Handler:    _APIService_DisablePiped_Handler,
		},
		{
			MethodName: "RegisterEvent",
			Handler:    _APIService_RegisterEvent_Handler,
		},
		{
			MethodName: "RequestPlanPreview",
			Handler:    _APIService_RequestPlanPreview_Handler,
		},
		{
			MethodName: "GetPlanPreviewResults",
			Handler:    _APIService_GetPlanPreviewResults_Handler,
		},
		{
			MethodName: "Encrypt",
			Handler:    _APIService_Encrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/app/server/service/apiservice/service.proto",
}