// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package purple

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// PurpleClient is the client API for Purple service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PurpleClient interface {
	GetHomePage(ctx context.Context, in *HomePageParam, opts ...grpc.CallOption) (*HomePageResponse, error)
	GetMember(ctx context.Context, in *GetMemberParam, opts ...grpc.CallOption) (*MemberResponse, error)
}

type purpleClient struct {
	cc grpc.ClientConnInterface
}

func NewPurpleClient(cc grpc.ClientConnInterface) PurpleClient {
	return &purpleClient{cc}
}

var purpleGetHomePageStreamDesc = &grpc.StreamDesc{
	StreamName: "GetHomePage",
}

func (c *purpleClient) GetHomePage(ctx context.Context, in *HomePageParam, opts ...grpc.CallOption) (*HomePageResponse, error) {
	out := new(HomePageResponse)
	err := c.cc.Invoke(ctx, "/purple.Purple/GetHomePage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var purpleGetMemberStreamDesc = &grpc.StreamDesc{
	StreamName: "GetMember",
}

func (c *purpleClient) GetMember(ctx context.Context, in *GetMemberParam, opts ...grpc.CallOption) (*MemberResponse, error) {
	out := new(MemberResponse)
	err := c.cc.Invoke(ctx, "/purple.Purple/GetMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PurpleService is the service API for Purple service.
// Fields should be assigned to their respective handler implementations only before
// RegisterPurpleService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type PurpleService struct {
	GetHomePage func(context.Context, *HomePageParam) (*HomePageResponse, error)
	GetMember   func(context.Context, *GetMemberParam) (*MemberResponse, error)
}

func (s *PurpleService) getHomePage(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HomePageParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetHomePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/purple.Purple/GetHomePage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetHomePage(ctx, req.(*HomePageParam))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *PurpleService) getMember(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemberParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/purple.Purple/GetMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetMember(ctx, req.(*GetMemberParam))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterPurpleService registers a service implementation with a gRPC server.
func RegisterPurpleService(s grpc.ServiceRegistrar, srv *PurpleService) {
	srvCopy := *srv
	if srvCopy.GetHomePage == nil {
		srvCopy.GetHomePage = func(context.Context, *HomePageParam) (*HomePageResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method GetHomePage not implemented")
		}
	}
	if srvCopy.GetMember == nil {
		srvCopy.GetMember = func(context.Context, *GetMemberParam) (*MemberResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method GetMember not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "purple.Purple",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "GetHomePage",
				Handler:    srvCopy.getHomePage,
			},
			{
				MethodName: "GetMember",
				Handler:    srvCopy.getMember,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "purple.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewPurpleService creates a new PurpleService containing the
// implemented methods of the Purple service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewPurpleService(s interface{}) *PurpleService {
	ns := &PurpleService{}
	if h, ok := s.(interface {
		GetHomePage(context.Context, *HomePageParam) (*HomePageResponse, error)
	}); ok {
		ns.GetHomePage = h.GetHomePage
	}
	if h, ok := s.(interface {
		GetMember(context.Context, *GetMemberParam) (*MemberResponse, error)
	}); ok {
		ns.GetMember = h.GetMember
	}
	return ns
}

// UnstablePurpleService is the service API for Purple service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstablePurpleService interface {
	GetHomePage(context.Context, *HomePageParam) (*HomePageResponse, error)
	GetMember(context.Context, *GetMemberParam) (*MemberResponse, error)
}
