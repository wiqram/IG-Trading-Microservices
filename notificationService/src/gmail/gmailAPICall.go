package gmail

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//"cloud.google.com/go/pubsub"
	//"sync"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

/* func getTradeAction(srv gmail.Service, user string, sender string) (*gmail.Message, error) {
        fmt.Printf("got gmail client, got user and sender",user+sender)
        // New message for our gmail service to send
        	//var message srv.Message

        message, err := srv.message.retrieve(user)
        fmt.Printf("this is the email received", message)
        return message, err
} */

func getGmailMessageByMessageId(srv *gmail.Service, user string, messageId string) (string, error) {
	strTextMessage := ""
	strHTMLMessage := ""
	attachmentId := ""
	attachmentName := ""
	message, err := srv.Users.Messages.Get(user, messageId).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve labels: %v", err)
	}

	messagePart := message.Payload
	//first check if the messagepart is null and if the Body not null, that means all info is available in body
	if messagePart.Parts == nil && messagePart.Body.Size > 0 {
		messagePartBody := messagePart.Body
		decoded, err := base64.StdEncoding.DecodeString(messagePartBody.Data)
		if err != nil {
			fmt.Println("decode error:", err)
			return "", err
		}
		if string(decoded) != "" {
			if messagePart.MimeType == "text/plain" {
				strTextMessage += string(decoded)
			} else if messagePart.MimeType == "text/html" {
				strHTMLMessage += string(decoded)
			} else {
				if messagePart.Filename != "" {
					attachmentId = messagePart.Body.AttachmentId
					attachmentName = messagePart.Filename
				}
			}

			//fmt.Printf("- message load payload Body %s\n", string(decoded))
		}
		if nil != messagePartBody.Header && len(messagePartBody.Header) > 0 {
			for key, value := range messagePartBody.Header {
				fmt.Printf("\nThis is Key NAME++ \n", key)
				fmt.Printf("\nThis is Key VALUE++ \n", value)
			}
		}

	} else
	//message has parts hence we need to now parse those parts and their body 1st Loop
	{
		//fmt.Printf("- message load payload has parts so need to parse that \n")
		for _, l := range messagePart.Parts {
			//fmt.Printf("-inside for loop for Parts \n")
			//Check if the Part has further Parts
			if l.Parts == nil && l.Body.Size > 0 {
				//Parts doesnt have further parts and the Body can be used as the content of message
				decoded, err := base64.StdEncoding.DecodeString(l.Body.Data)
				if err != nil {
					fmt.Println("decode error:", err)
					return "", err
				}
				if string(decoded) != "" {
					if l.MimeType == "text/plain" {
						strTextMessage += string(decoded)
					} else if l.MimeType == "text/html" {
						strHTMLMessage += string(decoded)
					} else {
						if l.Filename != "" {
							attachmentId = l.Body.AttachmentId
							attachmentName = l.Filename
						}
					}

					//fmt.Printf("- message load payload Body %s\n", string(decoded))
				}
				//fmt.Printf("- Found the latest message \n")
				/*					if nil != l.Headers && len(l.Headers) > 0 {
									for _, value := range l.Headers {
										fmt.Printf("\nThis is Key NAME++ \n", value.Name)
										fmt.Printf("\nThis is Key VALUE++ \n", value.Value)
									}
								}*/
			} else
			//Part has further parts hence need to loop the 2nd time to check for body
			{
				//fmt.Printf("- message load payload Body Parts has further parts \n")
				for _, lP := range l.Parts {
					if lP.Parts == nil && lP.Body.Size > 0 {
						decoded, err := base64.StdEncoding.DecodeString(lP.Body.Data)
						if err != nil {
							fmt.Println("decode error:", err)
							return "", err
						}
						if string(decoded) != "" {
							if lP.MimeType == "text/plain" {
								strTextMessage += string(decoded)
							} else if lP.MimeType == "text/html" {
								strHTMLMessage += string(decoded)
							} else {
								if lP.Filename != "" {
									attachmentId = lP.Body.AttachmentId
									attachmentName = lP.Filename
								}
							}

							//fmt.Printf("- message load payload Body %s\n", string(decoded))
						}
						/*							fmt.Printf("- Found the latest message %s\n")
													if nil != l.Headers && len(l.Headers) > 0 {
														for _, value := range l.Headers {
															fmt.Printf("\nThis is Key NAME++ \n", value.Name)
															fmt.Printf("\nThis is Key VALUE++ \n", value.Value)
														}
													}*/
					}
				}
			}
		}
	}
	//messagePart := message.Payload
	fmt.Printf("\nThis is Attachment ID \n", attachmentId)
	fmt.Printf("\nThis is Attachment Filename\n", attachmentName)

	return strTextMessage, err

}

func getGmailMessageByLabelId(srv *gmail.Service, user string, labelToRetrieve string) (string, error) {
	fmt.Printf("inside get gmail message function")
	// New message for our gmail service to send
	//var message srv.Message
	str := ""
	messages, err := srv.Users.Messages.List(user).LabelIds(labelToRetrieve).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve labels: %v", err)
	}
	//fmt.Printf("this is the email received in main function", messages)
	//fmt.Println("%s\n Messages are as below: %s\n")
	for _, l := range messages.Messages {
		fmt.Printf("message ID- %s\n ", l.Id)
		message, err := srv.Users.Messages.Get(user, l.Id).Do()
		if err != nil {
			log.Fatalf("Unable to retrieve labels: %v", err)
		}
		//fmt.Printf("- message load label %s\n", message.LabelIds)
		//fmt.Printf("- message load snippet %s\n", message.Snippet)

		messagePart := message.Payload

		if messagePart.Parts == nil && messagePart.Body != nil {
			messagePartBody := messagePart.Body
			decoded, err := base64.StdEncoding.DecodeString(messagePartBody.Data)
			if err != nil {
				fmt.Println("decode error:", err)
				return "", err
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
							return "", err
						}
						str += string(decoded)
						fmt.Printf("- Found the latest message %s\n")
						return string(decoded), err
					}
				} else {
					fmt.Printf("- message load payload Body Parts has further parts %s\n")
				}
			}
		}
		//messagePart := message.Payload
	}
	return "", err
}

func getLabelIds(srv *gmail.Service, label string, user string) (string, error) {
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

func fetchMessageFromHistoryIds(srv *gmail.Service, historyId uint64, watchResponse gmail.WatchResponse, user string) (string, error) {
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

func setWatchRequests(srv *gmail.Service, user string, labelToRetrieve string, topicName string) (*gmail.WatchResponse, error) {
	fmt.Printf("inside watchrequest function %s\n ")
	//var watchRequest *gmail.WatchRequest
	watchRequest := new(gmail.WatchRequest)
	watchRequest.LabelFilterAction = "include"
	watchRequest.LabelIds = append(watchRequest.LabelIds, labelToRetrieve)
	watchRequest.TopicName = topicName
	//fmt.Printf("This is watch request  TopicName %s\n ", watchRequest.TopicName)
	//fmt.Printf("This is watch request  label id %s\n ", watchRequest.LabelIds)
	watchResponse, err := srv.Users.Watch(user, watchRequest).Do()
	if err != nil {
		log.Fatalf("Did not get any response back from Watch call: %v", err)
		return nil, err
	}
	//fmt.Printf("This is watch response  %s\n ", watchResponse.HistoryId, watchResponse.Expiration)
	return watchResponse, err
}

/*
//pull messages from subscription of pub/subscription
func pullMsgsSync(w io.Writer, projectID, subID string) error {
        // projectID := "my-project-id"
        // subID := "my-sub"
        ctx := context.Background()
        client, err := pubsub.NewClient(ctx, projectID)
        if err != nil {
                return fmt.Errorf("pubsub.NewClient: %v", err)
        }
        defer client.Close()

        sub := client.Subscription(subID)

        // Turn on synchronous mode. This makes the subscriber use the Pull RPC rather
        // than the StreamingPull RPC, which is useful for guaranteeing MaxOutstandingMessages,
        // the max number of messages the client will hold in memory at a time.
        sub.ReceiveSettings.Synchronous = true
        sub.ReceiveSettings.MaxOutstandingMessages = 10

        // Receive messages for 5 seconds.
        ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
        defer cancel()

        // Create a channel to handle messages to as they come in.
        cm := make(chan *pubsub.Message)
        defer close(cm)
        // Handle individual messages in a goroutine.
        go func() {
                for msg := range cm {
                        fmt.Fprintf(w, "Got message :%q\n", string(msg.Data))
                        msg.Ack()
                }
        }()

        // Receive blocks until the passed in context is done.
        err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
                cm <- msg
        })
        if err != nil && status.Code(err) != codes.Canceled {
                return fmt.Errorf("Receive: %v", err)
        }

        return nil
}
*/

/* func pullMsgs(w io.Writer, projectID, subID string) error {
        // projectID := "my-project-id"
        // subID := "my-sub"
        ctx := context.Background()
        client, err := pubsub.NewClient(ctx, projectID)
        if err != nil {
                return fmt.Errorf("pubsub.NewClient: %v", err)
        }

        // Consume 10 messages.
        var mu sync.Mutex
        received := 0
        sub := client.Subscription(subID)
        cctx, cancel := context.WithCancel(ctx)
        err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
                mu.Lock()
                defer mu.Unlock()
                fmt.Fprintf(w, "Got message: %q\n", string(msg.Data))
                msg.Ack()
                received++
                if received == 10 {
                        cancel()
                }
        })
        if err != nil {
                return fmt.Errorf("Receive: %v", err)
        }
        return nil
} */

// pullMsgs - Get 10 messages for a given pubsub topic.
//
// Params:
//     client *pubsub.Client - The pubsub client.
//     name string - The name of the subscription to use.
//     topic *pubsub.Topic - The topic to get messages from.
//
// Return:
//     error - An error if it occured or nil.
/*
func pullMsgs(client *pubsub.Client, name string, topic *pubsub.Topic) error {
    ctx := context.Background()

    // Get the subscription.
    sub := client.Subscription(name)
    it, err := sub.Pull(ctx)
    if err != nil {
        return err
    }
    defer it.Stop()

    // Consume messages.
    for {
        msg, err := it.Next()
        if err == iterator.Done {
            break
        }
        if err != nil {
            return err
        }

        fmt.Printf("Got message: %q\n", string(msg.Data))
        msg.Done(true)
    }

    return nil
}
*/

func SubscribeToGmailLabels(emailId string, gmailLabel string, topicName string) (*gmail.WatchResponse, error) {
	user := "me"
	emailId = "me"
	//gmailLabel := "GregMana"
	//googleProjectId := "apiservice-1606168155272"
	//googleSubscriptionName := "projects/apiservice-1606168155272/subscriptions/gmailSubscribe"
	//topicName := "projects/apiservice-1606168155272/topics/igTradingTopic"
	b, err := ioutil.ReadFile("./secrets/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	//get labelId for the label that needs to be searched for emails. For this you must tag a label against the sender
	labelToRetrieve, err := getLabelIds(srv, gmailLabel, user)
	fmt.Printf("label ID retrieved is  %s\n ", labelToRetrieve)
	//labelToRetrieve := "Label_1929097597847660255"

	//get push notifications on every new email that arrives on the labelId
	watchResponse, err := setWatchRequests(srv, user, labelToRetrieve, topicName)
	fmt.Printf("This is watch response historyid %s\n ", watchResponse.HistoryId)
	if err != nil {
		log.Fatalf("Unable to subscribe to Gmail label: %v", err)
		return nil, err
	}
	/*err = PullEmailss(srv, bytes.NewBufferString("your string"), "", "", "", user, *watchResponse)
	if err != nil {
		log.Println("Api: PullMail", "failed to pull any emails from gmail subscription topic", err.Error())
		return nil, err
	}*/
	return watchResponse, nil
	//get the latest email from the emails within the labelId
	//gmailMessage, err := getGmailMessage(srv, user, labelToRetrieve)
	//fmt.Printf("- GMAIL MESSAGE IS ...%s\n", gmailMessage)

}
