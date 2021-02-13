package server

import (
	"context"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"github.com/wiqram/IG-Trading-Microservices/gmailWatchService/src/gmail"
	pb "github.com/wiqram/IG-Trading-Microservices/protos/gmailapi"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
)

// Server holds necessary values for server
type server struct {
	pb.UnimplementedGmailapiSvcServer
}

func Execute() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	sched := gocron.NewScheduler()
	sched.Every(1).Minute().Do(autoSubscribeToMail())
	<-sched.Start()
	port := ":" + os.Getenv("PORT")
	//proxyPort := ":" + os.Getenv("PROXY_PORT")
	fmt.Printf("Gmail API server is listenining on port %s\n ", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGmailapiSvcServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Printf("NotificationService GRPC server is listenining on port %s\n ", port)
}

// autoSubscribeToMail subscribes to the mailbox with specific label id

func autoSubscribeToMail() error {
	fmt.Print("-----------------START OF AUTO GMAIL SUBSCRIBE SERVICE-----------------------\n ")
	emailId := "me"
	labelName := os.Getenv("LABELTORETRIEVE")
	topicName := os.Getenv("GMAIL_TOPIC_NAME")
	c := gmail.NewGmailClient(emailId)
	subResp, err := c.SubscribeToGmailLabel(emailId, labelName, topicName)
	if err != nil {
		log.Println("Api: SubscribeToMail", "failed to subscribe to mailbox", err.Error())
		return err
	}
	err = c.PullEmails("", "", "", emailId, 88888, *subResp)
	if err != nil {
		log.Println("Api: PullMail", "failed to pull any emails from gmail subscription topic", err.Error())
		return err
	}
	fmt.Print("------------------END OF AUTO GMAIL SUBSCRIBE SERVICE--------------------------------\n ")
	return nil
}

// subscribeToMail subscribes to the mailbox with specific label id

func (s *server) SubscribeToMail(ctx context.Context, in *pb.SubscriptionRequest) (*pb.SubscriptionResponse, error) {
	fmt.Print("-----------------START OF GMAIL SUBSCRIBE SERVICE-----------------------\n ")
	in.EmailId = "me"
	in.LabelName = os.Getenv("LABELTORETRIEVE")
	in.TopicName = os.Getenv("GMAIL_TOPIC_NAME")
	c := gmail.NewGmailClient(in.EmailId)
	subResp, err := c.SubscribeToGmailLabel(in.EmailId, in.LabelName, in.TopicName)
	if err != nil {
		log.Println("Api: SubscribeToMail", "failed to subscribe to mailbox", err.Error())
		return nil, err
	}
	fmt.Print("This is request Duration at server level -- \n ", in.Duration)
	fmt.Print("Now calling the service to keep Watching for any changes since watchrequest was initiated\n ", strconv.Itoa(int(subResp.HistoryId)))
	err = c.PullEmails("", "", "", in.EmailId, int(in.Duration), *subResp)
	if err != nil {
		log.Println("Api: PullMail", "failed to pull any emails from gmail subscription topic", err.Error())
		return nil, err
	}
	fmt.Print("------------------END OF GMAIL SUBSCRIBE SERVICE--------------------------------\n ")
	return &pb.SubscriptionResponse{TransactionId: fmt.Sprintf("%f", subResp.Expiration), HistoryId: fmt.Sprintf("%f", subResp.HistoryId)}, nil
}
