package server

import (
	"context"
	"google.golang.org/grpc"
	"log"

	/*	"github.com/wiqram/IG-Trading-Microservices/protos/task"*/

	pb "github.com/wiqram/IG-Trading-Microservices/protos/api"
	"github.com/wiqram/IG-Trading-Microservices/protos/gmailapi"
	"github.com/wiqram/IG-Trading-Microservices/protos/igapi"
	/*	"github.com/wiqram/IG-Trading-Microservices/protos/project"*/
	"github.com/wiqram/IG-Trading-Microservices/protos/uibackend"
	"github.com/wiqram/IG-Trading-Microservices/protos/user"
)

var (
	clientStreamDescForProxying = &grpc.StreamDesc{
		ServerStreams: true,
		ClientStreams: false,
	}
)

// Server holds necessary values for server
type Server struct {
	userSvcClient user.UserSvcClient
	/*projectSvcClient project.ProjectSvcClient
	taskSvcClient    task.TaskSvcClient*/
	pb.UnimplementedAPIServer
	gmailapiSvcClient  gmailapi.GmailapiSvcClient
	igapiSvcClient     igapi.IgSvcClient
	uibackendSvcClient uibackend.UibackendSvcClient
}

// New returns new instance of Server
func New(userServiceClient user.UserSvcClient, gmailSvcClient gmailapi.GmailapiSvcClient, igSvcClnt igapi.IgSvcClient, uibackSvcClient uibackend.UibackendSvcClient) *Server {
	s := &Server{
		userSvcClient: userServiceClient,
		/*		projectSvcClient:       projectServiceClient,
				taskSvcClient:          taskServiceClient,*/
		gmailapiSvcClient:      gmailSvcClient,
		igapiSvcClient:         igSvcClnt,
		uibackendSvcClient:     uibackSvcClient,
		UnimplementedAPIServer: pb.UnimplementedAPIServer{},
	}
	return s
}

// RegisterUser directs to user service register method
func (s *Server) RegisterUser(ctx context.Context, in *user.RegisterRequest) (*user.UserResponse, error) {
	return s.userSvcClient.Register(ctx, in)
}

/*// CreateProject creates project for user
func (s *Server) CreateProject(ctx context.Context, in *project.CreateProjectRequest) (*project.ProjectResponse, error) {
	resp := &project.ProjectResponse{}
	userID, err := interceptor.GetUserID(ctx)
	if err != nil {
		log.Println("Api: CreateProject", "failed to get user ID", err.Error())
		return resp, err
	}
	in.UserId = userID
	return s.projectSvcClient.CreateProject(ctx, in)
}

// CreateTask creates task
func (s *Server) CreateTask(ctx context.Context, in *task.CreateTaskRequest) (*task.TaskResponse, error) {
	resp := &task.TaskResponse{}
	userID, err := interceptor.GetUserID(ctx)
	if err != nil {
		log.Println("Api: CreateProject", "failed to get user ID", err.Error())
		return resp, err
	}
	in.UserId = userID
	return s.taskSvcClient.CreateTask(ctx, in)
}*/

// LoginUser directs to user service register method
func (s *Server) LoginUser(ctx context.Context, in *user.LoginRequest) (*user.UserResponse, error) {
	return s.userSvcClient.Login(ctx, in)
}

/*// GetProject gets detail about projects
func (s *Server) GetProject(ctx context.Context, in *project.GetProjectRequest) (*project.ProjectResponse, error) {
	resp := &project.ProjectResponse{}
	userID, err := interceptor.GetUserID(ctx)
	if err != nil {
		log.Println("Api: CreateProject", "failed to get user ID", err.Error())
		return resp, err
	}
	in.UserId = userID
	return s.projectSvcClient.GetProject(ctx, in)
}

// ListTasks lists tasks
func (s *Server) ListTasks(ctx context.Context, in *task.ListTasksRequest) (*task.ListTasksResponse, error) {
	resp := &task.ListTasksResponse{}
	_, err := interceptor.GetUserID(ctx)
	if err != nil {
		log.Println("Api: CreateProject", "failed to get user ID", err.Error())
		return resp, err
	}
	return s.taskSvcClient.ListTasks(ctx, in)
}
*/
// SubscribeToMail subscribes to the mailbox with specific label id
func (s *Server) SubscribeToMail(ctx context.Context, in *gmailapi.SubscriptionRequest) (*gmailapi.SubscriptionResponse, error) {
	/*fmt.Printf("we are in subscribe to mail server method in api")
	resp := &notificationservice.SubscriptionResponse{}
	userID, err := interceptor.GetUserID(ctx)
	labelName, err := interceptor.GetLabelName(ctx)
	topicName, err := interceptor.GetTopicName(ctx)
	if err != nil {
		log.Println("Api: SubscribeToMail", "failed to subscribe to mailbox", err.Error())
		return resp, err
	}
	in.EmailId = userID*/
	return s.gmailapiSvcClient.SubscribeToMail(ctx, in)
}

// TradeAction Open Trade using IGAPI Trade Action
func (s *Server) TradeAction(ctx context.Context, in *igapi.OpenTradeRequest) (*igapi.TradeResponse, error) {
	/*resp := &igapi.TradeResponse{}
	userID, err := interceptor.GetUserID(ctx)
	if err != nil {
		log.Println("Api: TradeAction", "failed to get user ID", err.Error())
		return resp, err
	}
	in.UserId = userID*/
	return s.igapiSvcClient.TradeAction(ctx, in)
}

// GetProducts Open Trade using UIbackend Trade Action
func (s *Server) GetProducts(ctx context.Context, in *uibackend.ProductsRequest) (*uibackend.ProductsResponse, error) {
	/*resp := &uibackend.ProductsResponse{}
	userID, err := interceptor.GetUserID(ctx)
	if err != nil {
		log.Println("Api: TradeAction", "failed to get user ID", err.Error())
		return resp, err
	}
	resp. = userID
	*/
	/*	headers := metadata.Pairs(
		"Content-Type", "text/html; charset=utf-8",
		"Access-Control-Allow-Origin", "*",
	/*"Access-Control-Allow-Methods", "GET",*/
	/*)
	grpc.SendHeader(ctx, headers)
	log.Println("Inside getproducts of apiservice server")*/
	return s.uibackendSvcClient.GetProducts(ctx, in)
}

// GetTransactions Open Trade using UIbackend Trade Action
func (s *Server) GetTransactions(ctx context.Context, in *uibackend.TransactionsRequest) (*uibackend.TransactionsResponse, error) {
	//resp := &uibackend.TransactionsResponse{}
	/*userID, err := interceptor.GetUserID(ctx)
	if err != nil {
		log.Println("Api: TradeAction", "failed to get user ID", err.Error())
		return resp, err
	}*/
	/*	log.Println("Inside gettransactions of apiservice server")*/
	return s.uibackendSvcClient.GetTransactions(ctx, in)
}

// OpenOTCOrder Open Trade using UIbackend Trade Action
func (s *Server) OpenOTCOrder(ctx context.Context, in *gmailapi.OTCOrderRequest) (*gmailapi.OTCOrderResponse, error) {
	return s.gmailapiSvcClient.OpenOTCOrder(ctx, in)
}

// GetConfirmationDetails Open Trade using UIbackend Trade Action
func (s *Server) GetConfirmationDetails(ctx context.Context, in *gmailapi.OTCOrderResponse) (*gmailapi.ConfirmationResponse, error) {
	return s.gmailapiSvcClient.GetConfirmationDetails(ctx, in)
}

// GetClientSentiment Open Trade using UIbackend Trade Action
func (s *Server) GetClientSentiment(ctx context.Context, in *gmailapi.ClientSentimentRequest) (*gmailapi.ClientSentimentResponse, error) {
	return s.gmailapiSvcClient.GetClientSentiment(ctx, in)
}

// OpenLightStreamerSubscription Open Trade using UIbackend Trade Action
func (s *Server) OpenLightStreamerSubscription(ctx context.Context, in *gmailapi.LightStreamerSubRequest) (*gmailapi.LightStreamerSubResponse, error) {
	resp, err := s.gmailapiSvcClient.OpenLightStreamerSubscription(ctx, in)
	if err != nil {
		log.Println("Api: TradeAction", "failed to get user ID", err.Error())
		return nil, err
	}
	lSRespresp, err := resp.Recv()
	return lSRespresp, nil
}

// TestAccountTextRazor Open Trade using IGAPI Trade Action
func (s *Server) TestAccountTextRazor(ctx context.Context, in *igapi.TextRazorRequest) (*igapi.TextRazorResponse, error) {
	return s.igapiSvcClient.TestAccountTextRazor(ctx, in)
}

// TestClassifierTextRazor Open Trade using IGAPI Trade Action
func (s *Server) TestClassifierTextRazor(ctx context.Context, in *igapi.TextRazorRequest) (*igapi.TextRazorResponse, error) {
	return s.igapiSvcClient.TestClassifierTextRazor(ctx, in)
}

// TestDictionaryTextRazor Open Trade using IGAPI Trade Action
func (s *Server) TestDictionaryTextRazor(ctx context.Context, in *igapi.TextRazorRequest) (*igapi.TextRazorResponse, error) {
	return s.igapiSvcClient.TestDictionaryTextRazor(ctx, in)
}

// TestAnalysisTextRazor Open Trade using IGAPI Trade Action
func (s *Server) TestAnalysisTextRazor(ctx context.Context, in *igapi.TextRazorRequest) (*igapi.TextRazorResponse, error) {
	return s.igapiSvcClient.TestAnalysisTextRazor(ctx, in)
}
