package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"github.com/wiqram/IG-Trading-Microservices/gmailWatchService/src/gmail"
	"log"
	"os"
)

func main() {
	sched := gocron.NewScheduler()
	sched.Every(24).Hours().Do(autoSubscribeToMail())
	<-sched.Start()
}

func autoSubscribeToMail() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
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
