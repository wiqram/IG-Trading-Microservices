package gmail

import (
	"cloud.google.com/go/pubsub"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// GmailClient represent a connection with Gmail
type GmailClient struct {
	srv   *gmail.Service
	email string
}

// NewGmailClient will set up a new GmailClient for the specified email address.
func NewGmailClient(email string) GmailClient {
	//ctx := context.Background()
	b, err := ioutil.ReadFile("./secrets/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	//client := getClient(ctx, getConfig())
	client := getClient(config)
	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve gmail Client %v", err)
	}
	return GmailClient{srv, email}
}

func (c *GmailClient) SubscribeToGmailLabel(emailId string, gmailLabel string, topicName string) (*gmail.WatchResponse, error) {
	//get labelId for the label that needs to be searched for emails. For this you must tag a label against the sender
	labelToRetrieve, err := c.getLabelId(c.srv, gmailLabel, emailId)
	fmt.Printf("label ID retrieved is  %s\n ", labelToRetrieve)
	//labelToRetrieve := "Label_1929097597847660255"

	//get push notifications on every new email that arrives on the labelId
	watchResponse, err := c.setWatchRequest(c.srv, emailId, labelToRetrieve, topicName)
	fmt.Printf("This is watch response historyid %s\n ", watchResponse.HistoryId)
	if err != nil {
		log.Fatalf("Unable to subscribe to Gmail label: %v", err)
		return nil, err
	}
	/*err = PullEmails(srv, bytes.NewBufferString("your string"), "", "", "", user, *watchResponse)
	if err != nil {
		log.Println("Api: PullMail", "failed to pull any emails from gmail subscription topic", err.Error())
		return nil, err
	}*/
	return watchResponse, nil
	//get the latest email from the emails within the labelId
	//gmailMessage, err := getGmailMessage(srv, user, labelToRetrieve)
	//fmt.Printf("- GMAIL MESSAGE IS ...%s\n", gmailMessage)

}

func (c *GmailClient) getLabelId(srv *gmail.Service, label string, user string) (string, error) {
	r, err := srv.Users.Labels.List(user).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve labels: %v", err)
		return "", err
	}
	if len(r.Labels) == 0 {
		fmt.Println("No labels found.")
		return "", err
	}
	fmt.Println("Labels:")
	for _, l := range r.Labels {
		if l.Name == label {
			fmt.Println("Found the label and now sending back labelId.")
			return l.Id, err
		}
		//fmt.Printf("- Label not found yet...%s\n", l.Name+""+l.Id)
	}
	//label not found
	return "", err
}

func (c *GmailClient) setWatchRequest(srv *gmail.Service, user string, labelToRetrieve string, topicName string) (*gmail.WatchResponse, error) {
	fmt.Printf("inside watchrequest function \n ")
	//var watchRequest *gmail.WatchRequest
	watchRequest := new(gmail.WatchRequest)
	watchRequest.LabelFilterAction = "include"
	watchRequest.LabelIds = append(watchRequest.LabelIds, labelToRetrieve)
	watchRequest.TopicName = topicName
	//fmt.Printf("This is watch request  TopicName %s\n ", watchRequest.TopicName)
	//fmt.Printf("This is watch request  label id %s\n ", watchRequest.LabelIds)
	//stop the watch request first and then start another one
	err := srv.Users.Stop(user).Do()
	if err != nil {
		log.Fatalf("Couldnt stop watch for user me: %v", err)
		return nil, err
	}
	watchResponse, err := srv.Users.Watch(user, watchRequest).Do()
	if err != nil {
		log.Fatalf("Did not get any response back from Watch call: %v", err)
		return nil, err
	}
	//fmt.Printf("This is watch response  %s\n ", watchResponse.HistoryId, watchResponse.Expiration)
	return watchResponse, err
}

//once subscribed successfully pull mails

func (c *GmailClient) PullEmails(googleProjectId, googleSubscriptionName, topicName, user string, secs int, watchResponse gmail.WatchResponse) error {
	//now tht the gmail api has been setup to send messages from greg label to sub id, get messages
	if googleProjectId == "" {
		googleProjectId = "apiservice-1606168155272"
	}
	if googleSubscriptionName == "" {
		googleSubscriptionName = "gmailSubscribe"
	}
	if topicName == "" {
		topicName = "igTradingTopic"
	}

	ctx := context.Background()
	clientPubSub, err := pubsub.NewClient(ctx, googleProjectId)
	if err != nil {
		log.Fatalf("Could not create pubsub Client: %v", err)
	}
	defer clientPubSub.Close()

	sub := clientPubSub.Subscription(googleSubscriptionName)
	// Receive messages for "secs" seconds.
	fmt.Print("\ninside pull emails duration is --- ", time.Duration(secs)*time.Second)
	ctx, cancel := context.WithTimeout(ctx, time.Duration(secs)*time.Second)
	defer cancel()
	// Create a channel to handle messages to as they come in.
	cm := make(chan *pubsub.Message)
	defer close(cm)
	// Handle individual messages in a goroutine.
	go func() {
		for msg := range cm {
			messData := string(msg.Data)
			line := messData[strings.Index(messData, "historyId"):strings.Index(messData, "}")]
			out := strings.TrimSpace(strings.TrimPrefix(line, "historyId"))
			histId := strings.TrimLeft(out, "\\\":")
			historyIdToFetch, err := strconv.Atoi(histId)
			messageIdToFetch, err := c.fetchMessageFromHistoryId(c.srv, uint64(historyIdToFetch), watchResponse, user)
			if err != nil {
				fmt.Errorf("Receive: %v", err)
			}
			fmt.Println("this is the message ID recieved --- Message ID", messageIdToFetch)
			email, err := c.FetchEmailFromMessageId(c.srv, user, messageIdToFetch)
			if err != nil {
				fmt.Errorf("Receive: %v", err)
			}
			if email.TextBody == "" && email.HtmlBody == "" {
				fmt.Printf("\nNo changes found since watch request intitiated")
			} else {
				strHTML, err := email.HTML()
				if err != nil {
					fmt.Errorf("Receive: %v", err)
				}
				strTEXT, err := email.TEXT()
				if err != nil {
					fmt.Errorf("Receive: %v", err)
				}
				fmt.Printf("Got message Body HTML:\n", strHTML)
				fmt.Printf("Got message Body TEXT:\n", strTEXT)
				fmt.Printf("Got message ID:\n", string(email.Id))
				fmt.Printf("Got message Sender:\n", string(email.Sender))
				fmt.Printf("Got message Attachments length:\n", len(email.Attachments))
				fmt.Printf("Got message Subject:\n", email.Subject)
				dbURI := os.Getenv("DB_GCP_URI" + "?tlsCAFile=%s&tlsCertificateKeyFile=%s")
				dbName := os.Getenv("GCP_MONGO_DB_NAME")
				dbCertPem := os.Getenv("DB_GCP_PEM_loc")

				dbClient := NewDBClient(dbURI, dbName, dbCertPem)
				//Add emails to mongo DB and trigger signal extractor
				err = dbClient.AddMessageToDB(email)
				if err != nil {
					fmt.Printf("\nMessage was not loaded onto DB please Check", err)
				}
			}
			msg.Ack()
		}
	}()
	// Receive blocks until the context is cancelled or an error occurs.
	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		cm <- msg
	})
	if err != nil {
		return fmt.Errorf("Receive: %v", err)
	}
	return nil
}

func (c *GmailClient) fetchMessageFromHistoryId(srv *gmail.Service, historyId uint64, watchResponse gmail.WatchResponse, user string) (string, error) {
	//messageRecieved := new(gmail.Message)
	messageIdToFetch := ""
	/*	i64, err := strconv.Atoi(histId)
		if err != nil {
			log.Fatalf("Unable to retrieve labels: %v", err)
			return *messageRecieved, err
		}
		historyId := uint64(i64)*/
	//fmt.Printf("\nthis is the historyId of recent changge >>>>", strconv.Itoa(int(historyId)))
	//fmt.Printf("\nthis is the historyId of when watch was requested >>>>", strconv.Itoa(int(watchResponse.HistoryId)))
	//fmt.Printf("\nthis is the user >>>>", user)
	r, err := srv.Users.History.List(user).StartHistoryId(watchResponse.HistoryId).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve labels: %v", err)
		return "", err
	}
	if len(r.History) == 0 {
		fmt.Printf("\nNo changes found since watch request intitiated")
		return "", err
	}
	//historyArray := r.History
	//fmt.Printf("\nHistory found and its array length is:\n", len(r.History))

	for _, va := range r.History {
		//fmt.Printf("\nMEssages found inside History  and Messages array length is:\n", len(va.Messages))
		//fmt.Printf("we are inside history loop>>this is the hsitory ID inside history", strconv.Itoa(int(va.Id)), ">>>>and this is the history Id to compare against>>>", historyId)
		if va.MessagesAdded != nil && len(va.MessagesAdded) > 0 {
			for _, val := range va.MessagesAdded {
				//fmt.Printf("\nthis is the historyID fetched from message >>>>\n", strconv.Itoa(int(val.Message.HistoryId)))
				//fmt.Printf("\nmessage labelIds matched!!! >>>>\n", val.Message.LabelIds)
				//fmt.Printf("\nmessage RAW FORMAT ----- >>>>\n", val.Message.Raw)
				//fmt.Printf("\nthis is the historyID from request>>>>\n", strconv.Itoa(int(historyId)))
				messageIdToFetch = val.Message.Id
				//messageRecieved = val.Message
				if err != nil {
					log.Fatalf("Unable to retrieve labels: %v", err)
					return "", err
				}
				//fmt.Printf("\nthis is the message that we recieved just NOWWWW >>>>\n", email)
				/*
					messagePart := val.Message.Payload

					if messagePart.Parts == nil && messagePart.Body != nil {
						messagePartBody := messagePart.Body
						decoded, err := base64.StdEncoding.DecodeString(messagePartBody.Data)
						if err != nil {
							fmt.Println("decode error:", err)
							return *messageRecieved,err
						}
						if string(decoded) != "" {
							fmt.Printf("- message load payload Body %s\n", string(decoded))
						}
					} else {
						fmt.Printf("- message load payload has parts so need to parse that %s\n")
						for _, l := range messagePart.Parts {
							fmt.Printf("-inside for loop for Parts %s\n")
							if l.Parts == nil {
								if l.Body != nil && l.Body.Data != "" {
									decoded, err := base64.StdEncoding.DecodeString(l.Body.Data)
									if err != nil {
										fmt.Println("decode error:", err)
										return *messageRecieved, err
									}
									str += string(decoded)
									fmt.Printf("- Found the latest message %s\n", str)
									return *messageRecieved, err
								}
							} else {
								fmt.Printf("- message load payload Body Parts has further parts %s\n")
							}
						}
					}*/
			}
		}

	}
	//fmt.Printf("- Label not found yet...%s\n", l.Name+""+l.Id)
	//label not found
	return messageIdToFetch, err
}

func (c *GmailClient) FetchEmailFromMessageId(srv *gmail.Service, userId string, messageId string) (Email, error) {

	call := c.srv.Users.Messages.Get(userId, messageId)
	res, err := call.Format("full").Do()

	email := Email{}

	if err != nil {
		return email, err
	}
	if res != nil && res.Id != "" {
		email.Id = res.Id
		email.HtmlBody = getMessageHtmlBody(res.Payload.Parts)
		email.TextBody = getMessageTextBody(res.Payload.Parts)
		email.Sender = getMessageSender(res.Payload.Headers)
		email.Subject = getMessageSubject(res.Payload.Headers)
		email.Attachments = c.getMessageAttachments(
			res.Payload.Parts,
			res.Id,
			c.srv.Users.Messages.Attachments,
		)
	}
	return email, nil
}

/*// Email will return an email type that contains the HTML body, subject, sender
// and related attachments of an email specified by it's ID.
func (c *GmailClient) Email(id string) (email.Email, error) {
	call := c.srv.Users.Messages.Get(c.email, id)
	res, err := call.Format("full").Do()

	email := email.Email{}

	if err != nil {
		return email, err
	}

	email.Id = id
	email.Body = getMessageBody(res.Payload.Parts)
	email.Sender = getMessageSender(res.Payload.Headers)
	email.Subject = getMessageSubject(res.Payload.Headers)
	email.Attachments = c.getMessageAttachments(
		res.Payload.Parts,
		id,
		c.srv.Users.Messages.Attachments,
	)

	return email, nil
}
*/
// getMessageBody finds the HTML body of an email.
func getMessageHtmlBody(parts []*gmail.MessagePart) string {
	for _, part := range parts {
		if len(part.Parts) > 0 {
			return getMessageHtmlBody(part.Parts)
		} else {
			if part.MimeType == "text/html" {
				return part.Body.Data
			}
		}
	}

	return ""
}

// getMessageBody finds the TEXT body of an email.
func getMessageTextBody(parts []*gmail.MessagePart) string {
	for _, part := range parts {
		if len(part.Parts) > 0 {
			return getMessageTextBody(part.Parts)
		} else {
			if part.MimeType == "text/plain" {
				return part.Body.Data
			}
		}
	}

	return ""
}

// getMessageAttachments goes through the message parts for a specific message
// to find all the image attachments.
func (c *GmailClient) getMessageAttachments(
	parts []*gmail.MessagePart,
	messageId string,
	s *gmail.UsersMessagesAttachmentsService) []Attachment {

	attachments := make([]Attachment, 0)

	for _, part := range parts {
		if len(part.Parts) == 0 {
			if isImageAttachment(part.MimeType) {
				gmailAttachment := s.Get(
					c.email,
					messageId,
					part.Body.AttachmentId,
				)
				body, err := gmailAttachment.Do()

				if err != nil {
					continue
				}

				attachment := Attachment{
					Body:     body.Data,
					MimeType: part.MimeType,
					Filename: part.Filename,
				}
				attachments = append(attachments, attachment)
			}
		}
	}

	return attachments
}

var imageMimeTypes = []string{
	"image/png",
	"image/jpg",
	"image/jpeg",
	"image/gif",
	"application/pdf",
	"application/msword",
}

// isImageAttachment validates a mime type to see if it is an image MimeType.
func isImageAttachment(mime string) bool {
	for _, tp := range imageMimeTypes {
		if tp == mime {
			return true
		}
	}

	return false
}

// getMessageSender goes through the headers to find the From header.
func getMessageSender(headers []*gmail.MessagePartHeader) string {
	return getMessageHeader(headers, "From")
}

// getMessageSubject goes through the headers to find the Subject header.
func getMessageSubject(headers []*gmail.MessagePartHeader) string {
	return getMessageHeader(headers, "Subject")
}

// getMessageHeader goes through a list of headers and returns the header where
// the name matches the one we want.
func getMessageHeader(headers []*gmail.MessagePartHeader, wanted string) string {
	for _, header := range headers {
		if header.Name == wanted {
			return header.Value
		}

	}

	return ""
}

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClients(ctx context.Context, config *oauth2.Config) *http.Client {
	return config.Client(ctx, oauth2Token())
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func oauth2Token() *oauth2.Token {
	tm, _ := time.Parse("2006-Jan-02", "2015-Nov-13")
	return &oauth2.Token{
		AccessToken:  os.Getenv("GOOGLE_ACCESS_TOKEN"),
		TokenType:    "Bearer",
		RefreshToken: os.Getenv("GOOGLE_REFRESH_TOKEN"),
		Expiry:       tm,
	}
}

func getConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}
}
