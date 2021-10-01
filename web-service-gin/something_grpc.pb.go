// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package main

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

// CourseManagerClient is the client API for CourseManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourseManagerClient interface {
	GetCourses(ctx context.Context, in *GetCoursesRequest, opts ...grpc.CallOption) (*ReturnCourses, error)
	GetCourseByID(ctx context.Context, in *GetCourseByIDRequest, opts ...grpc.CallOption) (*ReturnCourse, error)
}

type courseManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewCourseManagerClient(cc grpc.ClientConnInterface) CourseManagerClient {
	return &courseManagerClient{cc}
}

func (c *courseManagerClient) GetCourses(ctx context.Context, in *GetCoursesRequest, opts ...grpc.CallOption) (*ReturnCourses, error) {
	out := new(ReturnCourses)
	err := c.cc.Invoke(ctx, "/DISYSMandatoryExercise1.CourseManager/getCourses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseManagerClient) GetCourseByID(ctx context.Context, in *GetCourseByIDRequest, opts ...grpc.CallOption) (*ReturnCourse, error) {
	out := new(ReturnCourse)
	err := c.cc.Invoke(ctx, "/DISYSMandatoryExercise1.CourseManager/getCourseByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourseManagerServer is the server API for CourseManager service.
// All implementations must embed UnimplementedCourseManagerServer
// for forward compatibility
type CourseManagerServer interface {
	GetCourses(context.Context, *GetCoursesRequest) (*ReturnCourses, error)
	GetCourseByID(context.Context, *GetCourseByIDRequest) (*ReturnCourse, error)
	mustEmbedUnimplementedCourseManagerServer()
}

// UnimplementedCourseManagerServer must be embedded to have forward compatible implementations.
type UnimplementedCourseManagerServer struct {
}

func (UnimplementedCourseManagerServer) GetCourses(context.Context, *GetCoursesRequest) (*ReturnCourses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourses not implemented")
}
func (UnimplementedCourseManagerServer) GetCourseByID(context.Context, *GetCourseByIDRequest) (*ReturnCourse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourseByID not implemented")
}
func (UnimplementedCourseManagerServer) mustEmbedUnimplementedCourseManagerServer() {}

// UnsafeCourseManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourseManagerServer will
// result in compilation errors.
type UnsafeCourseManagerServer interface {
	mustEmbedUnimplementedCourseManagerServer()
}

func RegisterCourseManagerServer(s grpc.ServiceRegistrar, srv CourseManagerServer) {
	s.RegisterService(&CourseManager_ServiceDesc, srv)
}

func _CourseManager_GetCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoursesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseManagerServer).GetCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DISYSMandatoryExercise1.CourseManager/getCourses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseManagerServer).GetCourses(ctx, req.(*GetCoursesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseManager_GetCourseByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCourseByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseManagerServer).GetCourseByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DISYSMandatoryExercise1.CourseManager/getCourseByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseManagerServer).GetCourseByID(ctx, req.(*GetCourseByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CourseManager_ServiceDesc is the grpc.ServiceDesc for CourseManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CourseManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DISYSMandatoryExercise1.CourseManager",
	HandlerType: (*CourseManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getCourses",
			Handler:    _CourseManager_GetCourses_Handler,
		},
		{
			MethodName: "getCourseByID",
			Handler:    _CourseManager_GetCourseByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "something.proto",
}