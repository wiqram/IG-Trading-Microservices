// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	gmailapi "github.com/wiqram/IG-Trading-Microservices/protos/gmailapi"
	igapi "github.com/wiqram/IG-Trading-Microservices/protos/igapi"
	project "github.com/wiqram/IG-Trading-Microservices/protos/project"
	task "github.com/wiqram/IG-Trading-Microservices/protos/task"
	uibackend "github.com/wiqram/IG-Trading-Microservices/protos/uibackend"
	user "github.com/wiqram/IG-Trading-Microservices/protos/user"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIClient interface {
	// RegisterUser creates user
	RegisterUser(ctx context.Context, in *user.RegisterRequest, opts ...grpc.CallOption) (*user.UserResponse, error)
	// LoginUser logs in user
	LoginUser(ctx context.Context, in *user.LoginRequest, opts ...grpc.CallOption) (*user.UserResponse, error)
	// CreateProject creates project
	CreateProject(ctx context.Context, in *project.CreateProjectRequest, opts ...grpc.CallOption) (*project.ProjectResponse, error)
	// GetProject retrives a project
	GetProject(ctx context.Context, in *project.GetProjectRequest, opts ...grpc.CallOption) (*project.ProjectResponse, error)
	// CreateTask creates task
	CreateTask(ctx context.Context, in *task.CreateTaskRequest, opts ...grpc.CallOption) (*task.TaskResponse, error)
	// UpdateTask updates task
	UpdateTask(ctx context.Context, in *task.UpdateTaskRequest, opts ...grpc.CallOption) (*task.TaskResponse, error)
	// ListTasks lists task
	ListTasks(ctx context.Context, in *task.ListTasksRequest, opts ...grpc.CallOption) (*task.ListTasksResponse, error)
	// SubscribeToMail call gmailapi subscribe
	SubscribeToMail(ctx context.Context, in *gmailapi.SubscriptionRequest, opts ...grpc.CallOption) (*gmailapi.SubscriptionResponse, error)
	// TradeAction set a trade on IG Broker based on action from email
	TradeAction(ctx context.Context, in *igapi.OpenTradeRequest, opts ...grpc.CallOption) (*igapi.TradeResponse, error)
	//GetProducts retrieves all products
	GetProducts(ctx context.Context, in *uibackend.ProductsRequest, opts ...grpc.CallOption) (*uibackend.ProductsResponse, error)
	//GetTransactions retrieves all products
	GetTransactions(ctx context.Context, in *uibackend.TransactionsRequest, opts ...grpc.CallOption) (*uibackend.TransactionsResponse, error)
	//openOTCOrder opens a trade over the counter on specific epic
	OpenOTCOrder(ctx context.Context, in *gmailapi.OTCOrderRequest, opts ...grpc.CallOption) (*gmailapi.OTCOrderResponse, error)
	//GetConfirmationDetails retrieves all products
	GetConfirmationDetails(ctx context.Context, in *gmailapi.OTCOrderResponse, opts ...grpc.CallOption) (*gmailapi.ConfirmationResponse, error)
	//GetClientSentiment retrieves all products
	GetClientSentiment(ctx context.Context, in *gmailapi.ClientSentimentRequest, opts ...grpc.CallOption) (*gmailapi.ClientSentimentResponse, error)
	//MarketSearch retrieves all products
	MarketSearch(ctx context.Context, in *gmailapi.ClientSentimentRequest, opts ...grpc.CallOption) (*gmailapi.MarketSearchResponse, error)
	//OpenLightStreamerSubscription retrieves all products
	OpenLightStreamerSubscription(ctx context.Context, in *gmailapi.LightStreamerSubRequest, opts ...grpc.CallOption) (*gmailapi.LightStreamerSubResponse, error)
	// TestAccountTextRazor set a trade on IG Broker based on action from email
	TestAccountTextRazor(ctx context.Context, in *igapi.TextRazorRequest, opts ...grpc.CallOption) (*igapi.TextRazorResponse, error)
	// TestClassifierTextRazor set a trade on IG Broker based on action from email
	TestClassifierTextRazor(ctx context.Context, in *igapi.TextRazorRequest, opts ...grpc.CallOption) (*igapi.TextRazorResponse, error)
	// TestDictionaryTextRazor set a trade on IG Broker based on action from email
	TestDictionaryTextRazor(ctx context.Context, in *igapi.TextRazorRequest, opts ...grpc.CallOption) (*igapi.TextRazorResponse, error)
	// TestAnalysisTextRazor set a trade on IG Broker based on action from email
	TestAnalysisTextRazor(ctx context.Context, in *igapi.TextRazorRequest, opts ...grpc.CallOption) (*igapi.TextRazorResponse, error)
}

type aPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIClient(cc grpc.ClientConnInterface) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) RegisterUser(ctx context.Context, in *user.RegisterRequest, opts ...grpc.CallOption) (*user.UserResponse, error) {
	out := new(user.UserResponse)
	err := c.cc.Invoke(ctx, "/demo_api.API/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) LoginUser(ctx context.Context, in *user.LoginRequest, opts ...grpc.CallOption) (*user.UserResponse, error) {
	out := new(user.UserResponse)
	err := c.cc.Invoke(ctx, "/demo_api.API/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateProject(ctx context.Context, in *project.CreateProjectRequest, opts ...grpc.CallOption) (*project.ProjectResponse, error) {
	out := new(project.ProjectResponse)
	err := c.cc.Invoke(ctx, "/demo_api.API/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetProject(ctx context.Context, in *project.GetProjectRequest, opts ...grpc.CallOption) (*project.ProjectResponse, error) {
	out := new(project.ProjectResponse)
	err := c.cc.Invoke(ctx, "/demo_api.API/GetProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateTask(ctx context.Context, in *task.CreateTaskRequest, opts ...grpc.CallOption) (*task.TaskResponse, error) {
	out := new(task.TaskResponse)
	err := c.cc.Invoke(ctx, "/demo_api.API/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) UpdateTask(ctx context.Context, in *task.UpdateTaskRequest, opts ...grpc.CallOption) (*task.TaskResponse, error) {
	out := new(task.TaskResponse)
	err := c.cc.Invoke(ctx, "/demo_api.API/UpdateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListTasks(ctx context.Context, in *task.ListTasksRequest, opts ...grpc.CallOption) (*task.ListTasksResponse, error) {
	out := new(task.ListTasksResponse)
	err := c.cc.Invoke(ctx, "/demo_api.API/ListTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) SubscribeToMail(ctx context.Context, in *gmailapi.SubscriptionRequest, opts ...grpc.CallOption) (*gmailapi.SubscriptionResponse, error) {
	out := new(gmailapi.SubscriptionResponse)
	err := c.cc.Invoke(ctx, "/demo_api.API/SubscribeToMail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) TradeAction(ctx context.Context, in *igapi.OpenTradeRequest, opts ...grpc.CallOption) (*igapi.TradeResponse, error) {
	out := new(igapi.TradeResponse)
	err := c.cc.Invoke(ctx, "/demo_api.API/TradeAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetProducts(ctx context.Context, in *uibackend.ProductsRequest, opts ...grpc.CallOption) (*uibackend.ProductsResponse, error) {
	out := new(uibackend.ProductsResponse)
	err := c.cc.Invoke(ctx, "/demo_api.API/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetTransactions(ctx context.Context, in *uibackend.TransactionsRequest, opts ...grpc.CallOption) (*uibackend.TransactionsResponse, error) {
	out := new(uibackend.TransactionsResponse)
	err := c.cc.Invoke(ctx, "/demo_api.API/GetTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) OpenOTCOrder(ctx context.Context, in *gmailapi.OTCOrderRequest, opts ...grpc.CallOption) (*gmailapi.OTCOrderResponse, error) {
	out := new(gmailapi.OTCOrderResponse)
	err := c.cc.Invoke(ctx, "/demo_api.API/OpenOTCOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetConfirmationDetails(ctx context.Context, in *gmailapi.OTCOrderResponse, opts ...grpc.CallOption) (*gmailapi.ConfirmationResponse, error) {
	out := new(gmailapi.ConfirmationResponse)
	err := c.cc.Invoke(ctx, "/demo_api.API/GetConfirmationDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetClientSentiment(ctx context.Context, in *gmailapi.ClientSentimentRequest, opts ...grpc.CallOption) (*gmailapi.ClientSentimentResponse, error) {
	out := new(gmailapi.ClientSentimentResponse)
	err := c.cc.Invoke(ctx, "/demo_api.API/GetClientSentiment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) MarketSearch(ctx context.Context, in *gmailapi.ClientSentimentRequest, opts ...grpc.CallOption) (*gmailapi.MarketSearchResponse, error) {
	out := new(gmailapi.MarketSearchResponse)
	err := c.cc.Invoke(ctx, "/demo_api.API/MarketSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) OpenLightStreamerSubscription(ctx context.Context, in *gmailapi.LightStreamerSubRequest, opts ...grpc.CallOption) (*gmailapi.LightStreamerSubResponse, error) {
	out := new(gmailapi.LightStreamerSubResponse)
	err := c.cc.Invoke(ctx, "/demo_api.API/OpenLightStreamerSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) TestAccountTextRazor(ctx context.Context, in *igapi.TextRazorRequest, opts ...grpc.CallOption) (*igapi.TextRazorResponse, error) {
	out := new(igapi.TextRazorResponse)
	err := c.cc.Invoke(ctx, "/demo_api.API/TestAccountTextRazor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) TestClassifierTextRazor(ctx context.Context, in *igapi.TextRazorRequest, opts ...grpc.CallOption) (*igapi.TextRazorResponse, error) {
	out := new(igapi.TextRazorResponse)
	err := c.cc.Invoke(ctx, "/demo_api.API/TestClassifierTextRazor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) TestDictionaryTextRazor(ctx context.Context, in *igapi.TextRazorRequest, opts ...grpc.CallOption) (*igapi.TextRazorResponse, error) {
	out := new(igapi.TextRazorResponse)
	err := c.cc.Invoke(ctx, "/demo_api.API/TestDictionaryTextRazor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) TestAnalysisTextRazor(ctx context.Context, in *igapi.TextRazorRequest, opts ...grpc.CallOption) (*igapi.TextRazorResponse, error) {
	out := new(igapi.TextRazorResponse)
	err := c.cc.Invoke(ctx, "/demo_api.API/TestAnalysisTextRazor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
// All implementations must embed UnimplementedAPIServer
// for forward compatibility
type APIServer interface {
	// RegisterUser creates user
	RegisterUser(context.Context, *user.RegisterRequest) (*user.UserResponse, error)
	// LoginUser logs in user
	LoginUser(context.Context, *user.LoginRequest) (*user.UserResponse, error)
	// CreateProject creates project
	CreateProject(context.Context, *project.CreateProjectRequest) (*project.ProjectResponse, error)
	// GetProject retrives a project
	GetProject(context.Context, *project.GetProjectRequest) (*project.ProjectResponse, error)
	// CreateTask creates task
	CreateTask(context.Context, *task.CreateTaskRequest) (*task.TaskResponse, error)
	// UpdateTask updates task
	UpdateTask(context.Context, *task.UpdateTaskRequest) (*task.TaskResponse, error)
	// ListTasks lists task
	ListTasks(context.Context, *task.ListTasksRequest) (*task.ListTasksResponse, error)
	// SubscribeToMail call gmailapi subscribe
	SubscribeToMail(context.Context, *gmailapi.SubscriptionRequest) (*gmailapi.SubscriptionResponse, error)
	// TradeAction set a trade on IG Broker based on action from email
	TradeAction(context.Context, *igapi.OpenTradeRequest) (*igapi.TradeResponse, error)
	//GetProducts retrieves all products
	GetProducts(context.Context, *uibackend.ProductsRequest) (*uibackend.ProductsResponse, error)
	//GetTransactions retrieves all products
	GetTransactions(context.Context, *uibackend.TransactionsRequest) (*uibackend.TransactionsResponse, error)
	//openOTCOrder opens a trade over the counter on specific epic
	OpenOTCOrder(context.Context, *gmailapi.OTCOrderRequest) (*gmailapi.OTCOrderResponse, error)
	//GetConfirmationDetails retrieves all products
	GetConfirmationDetails(context.Context, *gmailapi.OTCOrderResponse) (*gmailapi.ConfirmationResponse, error)
	//GetClientSentiment retrieves all products
	GetClientSentiment(context.Context, *gmailapi.ClientSentimentRequest) (*gmailapi.ClientSentimentResponse, error)
	//MarketSearch retrieves all products
	MarketSearch(context.Context, *gmailapi.ClientSentimentRequest) (*gmailapi.MarketSearchResponse, error)
	//OpenLightStreamerSubscription retrieves all products
	OpenLightStreamerSubscription(context.Context, *gmailapi.LightStreamerSubRequest) (*gmailapi.LightStreamerSubResponse, error)
	// TestAccountTextRazor set a trade on IG Broker based on action from email
	TestAccountTextRazor(context.Context, *igapi.TextRazorRequest) (*igapi.TextRazorResponse, error)
	// TestClassifierTextRazor set a trade on IG Broker based on action from email
	TestClassifierTextRazor(context.Context, *igapi.TextRazorRequest) (*igapi.TextRazorResponse, error)
	// TestDictionaryTextRazor set a trade on IG Broker based on action from email
	TestDictionaryTextRazor(context.Context, *igapi.TextRazorRequest) (*igapi.TextRazorResponse, error)
	// TestAnalysisTextRazor set a trade on IG Broker based on action from email
	TestAnalysisTextRazor(context.Context, *igapi.TextRazorRequest) (*igapi.TextRazorResponse, error)
	mustEmbedUnimplementedAPIServer()
}

// UnimplementedAPIServer must be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (UnimplementedAPIServer) RegisterUser(context.Context, *user.RegisterRequest) (*user.UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedAPIServer) LoginUser(context.Context, *user.LoginRequest) (*user.UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedAPIServer) CreateProject(context.Context, *project.CreateProjectRequest) (*project.ProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedAPIServer) GetProject(context.Context, *project.GetProjectRequest) (*project.ProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedAPIServer) CreateTask(context.Context, *task.CreateTaskRequest) (*task.TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedAPIServer) UpdateTask(context.Context, *task.UpdateTaskRequest) (*task.TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedAPIServer) ListTasks(context.Context, *task.ListTasksRequest) (*task.ListTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTasks not implemented")
}
func (UnimplementedAPIServer) SubscribeToMail(context.Context, *gmailapi.SubscriptionRequest) (*gmailapi.SubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeToMail not implemented")
}
func (UnimplementedAPIServer) TradeAction(context.Context, *igapi.OpenTradeRequest) (*igapi.TradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeAction not implemented")
}
func (UnimplementedAPIServer) GetProducts(context.Context, *uibackend.ProductsRequest) (*uibackend.ProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (UnimplementedAPIServer) GetTransactions(context.Context, *uibackend.TransactionsRequest) (*uibackend.TransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactions not implemented")
}
func (UnimplementedAPIServer) OpenOTCOrder(context.Context, *gmailapi.OTCOrderRequest) (*gmailapi.OTCOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenOTCOrder not implemented")
}
func (UnimplementedAPIServer) GetConfirmationDetails(context.Context, *gmailapi.OTCOrderResponse) (*gmailapi.ConfirmationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfirmationDetails not implemented")
}
func (UnimplementedAPIServer) GetClientSentiment(context.Context, *gmailapi.ClientSentimentRequest) (*gmailapi.ClientSentimentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientSentiment not implemented")
}
func (UnimplementedAPIServer) MarketSearch(context.Context, *gmailapi.ClientSentimentRequest) (*gmailapi.MarketSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketSearch not implemented")
}
func (UnimplementedAPIServer) OpenLightStreamerSubscription(context.Context, *gmailapi.LightStreamerSubRequest) (*gmailapi.LightStreamerSubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenLightStreamerSubscription not implemented")
}
func (UnimplementedAPIServer) TestAccountTextRazor(context.Context, *igapi.TextRazorRequest) (*igapi.TextRazorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestAccountTextRazor not implemented")
}
func (UnimplementedAPIServer) TestClassifierTextRazor(context.Context, *igapi.TextRazorRequest) (*igapi.TextRazorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestClassifierTextRazor not implemented")
}
func (UnimplementedAPIServer) TestDictionaryTextRazor(context.Context, *igapi.TextRazorRequest) (*igapi.TextRazorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestDictionaryTextRazor not implemented")
}
func (UnimplementedAPIServer) TestAnalysisTextRazor(context.Context, *igapi.TextRazorRequest) (*igapi.TextRazorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestAnalysisTextRazor not implemented")
}
func (UnimplementedAPIServer) mustEmbedUnimplementedAPIServer() {}

// UnsafeAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIServer will
// result in compilation errors.
type UnsafeAPIServer interface {
	mustEmbedUnimplementedAPIServer()
}

func RegisterAPIServer(s grpc.ServiceRegistrar, srv APIServer) {
	s.RegisterService(&API_ServiceDesc, srv)
}

func _API_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_api.API/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).RegisterUser(ctx, req.(*user.RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_api.API/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).LoginUser(ctx, req.(*user.LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(project.CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_api.API/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateProject(ctx, req.(*project.CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(project.GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_api.API/GetProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetProject(ctx, req.(*project.GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(task.CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_api.API/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateTask(ctx, req.(*task.CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(task.UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_api.API/UpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).UpdateTask(ctx, req.(*task.UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(task.ListTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_api.API/ListTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListTasks(ctx, req.(*task.ListTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_SubscribeToMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(gmailapi.SubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).SubscribeToMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_api.API/SubscribeToMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).SubscribeToMail(ctx, req.(*gmailapi.SubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_TradeAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(igapi.OpenTradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).TradeAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_api.API/TradeAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).TradeAction(ctx, req.(*igapi.OpenTradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(uibackend.ProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_api.API/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetProducts(ctx, req.(*uibackend.ProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(uibackend.TransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_api.API/GetTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetTransactions(ctx, req.(*uibackend.TransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_OpenOTCOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(gmailapi.OTCOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).OpenOTCOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_api.API/OpenOTCOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).OpenOTCOrder(ctx, req.(*gmailapi.OTCOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetConfirmationDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(gmailapi.OTCOrderResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetConfirmationDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_api.API/GetConfirmationDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetConfirmationDetails(ctx, req.(*gmailapi.OTCOrderResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetClientSentiment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(gmailapi.ClientSentimentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetClientSentiment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_api.API/GetClientSentiment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetClientSentiment(ctx, req.(*gmailapi.ClientSentimentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_MarketSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(gmailapi.ClientSentimentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).MarketSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_api.API/MarketSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).MarketSearch(ctx, req.(*gmailapi.ClientSentimentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_OpenLightStreamerSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(gmailapi.LightStreamerSubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).OpenLightStreamerSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_api.API/OpenLightStreamerSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).OpenLightStreamerSubscription(ctx, req.(*gmailapi.LightStreamerSubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_TestAccountTextRazor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(igapi.TextRazorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).TestAccountTextRazor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_api.API/TestAccountTextRazor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).TestAccountTextRazor(ctx, req.(*igapi.TextRazorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_TestClassifierTextRazor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(igapi.TextRazorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).TestClassifierTextRazor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_api.API/TestClassifierTextRazor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).TestClassifierTextRazor(ctx, req.(*igapi.TextRazorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_TestDictionaryTextRazor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(igapi.TextRazorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).TestDictionaryTextRazor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_api.API/TestDictionaryTextRazor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).TestDictionaryTextRazor(ctx, req.(*igapi.TextRazorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_TestAnalysisTextRazor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(igapi.TextRazorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).TestAnalysisTextRazor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_api.API/TestAnalysisTextRazor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).TestAnalysisTextRazor(ctx, req.(*igapi.TextRazorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// API_ServiceDesc is the grpc.ServiceDesc for API service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var API_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo_api.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _API_RegisterUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _API_LoginUser_Handler,
		},
		{
			MethodName: "CreateProject",
			Handler:    _API_CreateProject_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _API_GetProject_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _API_CreateTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _API_UpdateTask_Handler,
		},
		{
			MethodName: "ListTasks",
			Handler:    _API_ListTasks_Handler,
		},
		{
			MethodName: "SubscribeToMail",
			Handler:    _API_SubscribeToMail_Handler,
		},
		{
			MethodName: "TradeAction",
			Handler:    _API_TradeAction_Handler,
		},
		{
			MethodName: "GetProducts",
			Handler:    _API_GetProducts_Handler,
		},
		{
			MethodName: "GetTransactions",
			Handler:    _API_GetTransactions_Handler,
		},
		{
			MethodName: "OpenOTCOrder",
			Handler:    _API_OpenOTCOrder_Handler,
		},
		{
			MethodName: "GetConfirmationDetails",
			Handler:    _API_GetConfirmationDetails_Handler,
		},
		{
			MethodName: "GetClientSentiment",
			Handler:    _API_GetClientSentiment_Handler,
		},
		{
			MethodName: "MarketSearch",
			Handler:    _API_MarketSearch_Handler,
		},
		{
			MethodName: "OpenLightStreamerSubscription",
			Handler:    _API_OpenLightStreamerSubscription_Handler,
		},
		{
			MethodName: "TestAccountTextRazor",
			Handler:    _API_TestAccountTextRazor_Handler,
		},
		{
			MethodName: "TestClassifierTextRazor",
			Handler:    _API_TestClassifierTextRazor_Handler,
		},
		{
			MethodName: "TestDictionaryTextRazor",
			Handler:    _API_TestDictionaryTextRazor_Handler,
		},
		{
			MethodName: "TestAnalysisTextRazor",
			Handler:    _API_TestAnalysisTextRazor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/api/api.proto",
}
