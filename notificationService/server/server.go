package server

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/wiqram/IG-Trading-Microservices/notificationService/src/gmail"
	"github.com/wiqram/IG-Trading-Microservices/notificationService/src/igmarkets"
	pb "github.com/wiqram/IG-Trading-Microservices/protos/gmailapi"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"reflect"
	"strconv"
	"time"
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

// openOTCOrder subscribes to the mailbox with specific label id

func (s *server) OpenOTCOrder(ctx context.Context, in *pb.OTCOrderRequest) (*pb.OTCOrderResponse, error) {
	httpTimeout := time.Duration(15 * time.Second)
	ig := igmarkets.New(igmarkets.DemoAPIURL, "1ee808908ed01577d3ab2b19ef493d0f80e24612", "XXXK3", "wiqramdemo", "DemoPWD123", httpTimeout)
	if err := ig.Login(); err != nil {
		fmt.Println("Unable to login into IG account", err)
	}
	fmt.Printf("\nthis is what we get back from calling new igmarkets -- AuthToken", ig.OAuthToken.AccessToken)
	fmt.Printf("\nthis is what we get back from calling new igmarkets -- AuthToken", ig.OAuthToken.ExpiresIn)
	fmt.Printf("\nthis is what we get back from calling new igmarkets -- ", ig.APIKey)
	fmt.Printf("\nthis is what we get back from calling new igmarkets -- ", ig.Identifier)

	oTCOrderRequest := igmarkets.OTCOrderRequest{
		Epic: in.Epic,
		//Level:        in.Level,
		ForceOpen:    in.ForceOpen,
		OrderType:    in.OrderType,
		CurrencyCode: in.CurrencyCode,
		Direction:    in.Direction,
		Size:         float64(in.Size),
		//GuaranteedStop: in.GuaranteedStop,
		Expiry: in.Expiry,
	}

	mSresp, err := ig.MarketSearch("FTSE")
	if err != nil {
		log.Println("Api: PullMail", "failed to pull any emails from gmail subscription topic", err.Error())
		return nil, err
	}
	for _, n := range mSresp.Markets {
		fmt.Printf("\nthis is what we get back from calling new igmarkets EPC-- ", n.Epic)
		fmt.Printf("\nthis is what we get back from calling new igmarkets EPC-- ", n.Expiry)
		fmt.Printf("\nthis is what we get back from calling new igmarkets EPC-- ", n.InstrumentName)
	}
	oTCOrderRequest.Epic = "IX.D.FTSE.DAILY.IP"
	oTCOrderRequest.CurrencyCode = "GBP"
	oTCOrderRequest.Size = 1.0
	oTCOrderRequest.Direction = "BUY"
	oTCOrderRequest.OrderType = "MARKET"
	oTCOrderRequest.Expiry = "DFB"
	/*m := mapFields(in)
	o := applyMap(u, m)*/
	//o := make(M)
	/*vi := reflect.ValueOf(in).Elem()
	vo := reflect.ValueOf(u).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := t.FieldByIndex([]int{i})
		// skip unexported fields
		if f.PkgPath != "" {
			continue
		}
		u.[f.Name] = v.FieldByIndex([]int{i}).Interface()
	}*/
	fmt.Printf("\nthis is the value inside mapped orderRequest after mapping -- ")

	otcOrderResp, err := ig.PlaceOTCOrder(oTCOrderRequest)
	if err != nil {
		log.Println("Api: OTC Order Req", "failed to open any orders on OTC", err.Error())
		return nil, err
	}
	fmt.Printf("\nthis is the order response deal id -- ", otcOrderResp.DealReference)

	// Check order status
	confirmation, err := ig.GetDealConfirmation(otcOrderResp.DealReference)
	if err != nil {
		fmt.Println("Cannot get deal confirmation for:", otcOrderResp.DealReference, err)
		return nil, err
	}

	fmt.Println("Order dealRef", otcOrderResp.DealReference)
	fmt.Println("DealStatus", confirmation.DealStatus) // "ACCEPTED"
	fmt.Println("Profit", confirmation.Profit, confirmation.ProfitCurrency)
	fmt.Println("Status", confirmation.Status) // "OPEN"
	fmt.Println("Reason", confirmation.Reason)
	fmt.Println("Level", confirmation.Level) // Buy price

	return &pb.OTCOrderResponse{DealRef: otcOrderResp.DealReference}, nil
}

// getConfirmationDetails subscribes to the mailbox with specific label id

func (s *server) GetConfirmationDetails(ctx context.Context, in *pb.OTCOrderResponse) (*pb.ConfirmationResponse, error) {
	httpTimeout := time.Duration(5 * time.Second)
	ig := igmarkets.New(igmarkets.DemoAPIURL, "1ee808908ed01577d3ab2b19ef493d0f80e24612", "XXXK3", "wiqramdemo", "DemoPWD123", httpTimeout)
	if err := ig.Login(); err != nil {
		fmt.Println("Unable to login into IG account", err)
	}
	resp, err := ig.GetDealConfirmation(in.DealRef)
	if err != nil {
		log.Println("Api: OTC Order Req", "failed to open any orders on OTC", err.Error())
		return nil, err
	}

	confResp := pb.ConfirmationResponse{
		DealStatus:     resp.DealStatus,
		Profit:         float32(resp.Profit),
		ProfitCurrency: resp.ProfitCurrency,
		Status:         resp.Status,
		Reason:         resp.Reason,
		Level:          float32(resp.Level),
	}

	fmt.Printf("\nthis is what we get back from calling new confirmation details -- ", resp.Profit)
	fmt.Printf("\nthis is what we get back from calling new igmarkets EPC-- ", resp.Expiry)
	fmt.Printf("\nthis is what we get back from calling new igmarkets EPC-- ", resp.CurrencyCode)
	fmt.Printf("\nthis is what we get back from calling new igmarkets EPC-- ", resp.ForceOpen)
	fmt.Printf("\nthis is what we get back from calling new igmarkets EPC-- ", resp.QuoteID)

	return &confResp, nil
}

// getClientSentiment subscribes to the mailbox with specific label id

func (s *server) GetClientSentiment(ctx context.Context, in *pb.ClientSentimentRequest) (*pb.ClientSentimentResponse, error) {
	httpTimeout := time.Duration(5 * time.Second)
	ig := igmarkets.New(igmarkets.DemoAPIURL, "1ee808908ed01577d3ab2b19ef493d0f80e24612", "XXXK3", "wiqramdemo", "DemoPWD123", httpTimeout)
	if err := ig.Login(); err != nil {
		fmt.Println("Unable to login into IG account", err)
	}
	resp, err := ig.GetClientSentiment(in.MarketId)
	if err != nil {
		log.Println("Api: OTC Order Req", "failed to open any orders on OTC", err.Error())
		return nil, err
	}
	cSResp := pb.ClientSentimentResponse{
		LongPositionPercentage:  float32(resp.LongPositionPercentage),
		ShortPositionPercentage: float32(resp.ShortPositionPercentage),
	}
	return &cSResp, nil
}

// getClientSentiment subscribes to the mailbox with specific label id

func (s *server) OpenLightStreamerSubscription(in *pb.LightStreamerSubRequest, srv pb.GmailapiSvc_OpenLightStreamerSubscriptionServer) error {
	log.Printf("fetch marketid for request : %d", in.MarketId)
	httpTimeout := time.Duration(5 * time.Second)
	ig := igmarkets.New(igmarkets.DemoAPIURL, "1ee808908ed01577d3ab2b19ef493d0f80e24612", "XXXK3", "wiqramdemo", "DemoPWD123", httpTimeout)
	if err := ig.Login(); err != nil {
		fmt.Println("Unable to login into IG account", err)
	}
	for {
		tickChan := make(chan igmarkets.LightStreamerTick)
		err := ig.OpenLightStreamerSubscription([]string{"CS.D.BITCOIN.CFD.IP"}, tickChan)
		if err != nil {
			fmt.Println("OpenLightStreamerSubscription() failed", err)
		}

		for tick := range tickChan {
			fmt.Println("tick: %+v", tick)
		}

		fmt.Println("Server closed stream, restarting...")
		return nil
	}
}

type M map[string]interface{} // just an alias

func applyMap(u *igmarkets.OTCOrderRequest, m M) M {
	t := reflect.TypeOf(u).Elem()
	o := make(M)
	for i := 0; i < t.NumField(); i++ {
		f := t.FieldByIndex([]int{i})
		// skip unexported fields
		if f.PkgPath != "" {
			continue
		}
		if x, ok := m[f.Name]; ok {
			k := f.Tag.Get("json")
			o[k] = x
		}
	}
	return o
}

func mapFields(x *pb.OTCOrderRequest) M {
	o := make(M)
	v := reflect.ValueOf(x).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := t.FieldByIndex([]int{i})
		// skip unexported fields
		if f.PkgPath != "" {
			continue
		}
		o[f.Name] = v.FieldByIndex([]int{i}).Interface()
	}
	return o
}
