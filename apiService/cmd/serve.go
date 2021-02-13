package cmd

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/wiqram/IG-Trading-Microservices/apiService/interceptor"
	"github.com/wiqram/IG-Trading-Microservices/apiService/server"
	pb "github.com/wiqram/IG-Trading-Microservices/protos/api"
	"github.com/wiqram/IG-Trading-Microservices/protos/gmailapi"
	"github.com/wiqram/IG-Trading-Microservices/protos/igapi"
	"github.com/wiqram/IG-Trading-Microservices/protos/uibackend"
	"github.com/wiqram/IG-Trading-Microservices/protos/user"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// srvCmd is the serve sub command to start the api server
var srvCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve serves the api server",
	RunE:  serve,
}

func init() {
	//
}

func serve(cmd *cobra.Command, args []string) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := ":" + os.Getenv("PORT")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userAddress := os.Getenv("USER_ADDRESS")
	log.Println("Connecting to user service on " + userAddress)
	userConn, err := grpc.Dial(userAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer userConn.Close()
	userSvcClient := user.NewUserSvcClient(userConn)

	/*projectAddress := os.Getenv("PROJECT_ADDRESS")
	log.Println("Connecting to project service on " + projectAddress)
	projectConn, err := grpc.Dial(projectAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer projectConn.Close()
	projectSvcClient := project.NewProjectSvcClient(projectConn)

	taskAddress := os.Getenv("TASK_ADDRESS")
	log.Println("Connecting to task service on " + taskAddress)
	taskConn, err := grpc.Dial(taskAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer projectConn.Close()
	taskSvcClient := task.NewTaskSvcClient(taskConn)*/

	gmailapiAddress := os.Getenv("NOTIFICATIONSERVICE_ADDRESS")
	log.Println("Connecting to notificationservice service on " + gmailapiAddress)
	gmailapiConn, err := grpc.Dial(gmailapiAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer gmailapiConn.Close()
	gmailapiSvcClient := gmailapi.NewGmailapiSvcClient(gmailapiConn)
	/*t := &notificationservice.SubscriptionRequest{EmailId: defaultName}
	resp, err := gmailapiSvcClient.SubscribeToMail(context.Background(), t)
	log.Println("thsi is the response " + resp.HistoryId)*/

	igapiAddress := os.Getenv("IGAPI_ADDRESS")
	log.Println("Connecting to igapi service on " + igapiAddress)
	igapiConn, err := grpc.Dial(igapiAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer igapiConn.Close()
	igapiSvcClient := igapi.NewIgSvcClient(igapiConn)

	uibackendAddress := os.Getenv("UIBACKEND_ADDRESS")
	log.Println("Connecting to UIBACKEND service on " + uibackendAddress)
	//grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "401"))
	uibackendConn, err := grpc.Dial(uibackendAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer uibackendConn.Close()
	uibackendSvcClient := uibackend.NewUibackendSvcClient(uibackendConn)
	/*	t := &uibackend.ProductsRequest{}
		resp, err := uibackendSvcClient.GetProducts(context.Background(), t)
		if resp != nil {
			log.Println("this is the response " + resp.Products[0].AccountNumber)
		}*/

	unaryInterceptors := []grpc.UnaryServerInterceptor{
		interceptor.UnaryAuthenticate(userSvcClient),
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryInterceptors...)))
	pb.RegisterAPIServer(
		s,
		server.New(userSvcClient, gmailapiSvcClient, igapiSvcClient, uibackendSvcClient),
	)
	//pb.RegisterAPIServer(s, server.New(userSvcClient, projectSvcClient, taskSvcClient, gmailapiSvcClient))
	log.Println("Starting GRPC server at: " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}
