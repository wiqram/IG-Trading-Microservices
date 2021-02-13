package signalxtractor

import (
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/wiqram/IG-Trading-Microservices/notificationService/src/gmail"
	"io"
	"os"
	"regexp"
)

func dbConn() (db *sql.DB) {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file")
	}
	sqlDbUri := os.Getenv("SQL_DB_URI")
	dbDriver := os.Getenv("SQL_DRIVER_NAME")
	db, err = sql.Open(dbDriver, sqlDbUri)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func XtractSignalFromGmailMessage(email gmail.Email) error {
	//db := dbConn()
	//textBody := email.TextBody
	subject := email.Subject
	tradeSignal := TradeSignal{
		ClientName:           email.Sender,
		OriginatingMessageID: email.Id,
	}
	//find out if the subject as New word in it . if so it means its likely a new position
	newTradeSubject, err := regexp.MatchString(`(?im)New`, subject)
	//Is the trade speculative or has the client asked remain cautious
	//speculativeTrade, err := regexp.MatchString(`(?im)(cautious|speculative)`, textBody)
	//addedToExistingTrade, err := regexp.MatchString(`(?im)(just added|added)`, textBody)

	//optionStrikePrice, err := regexp.MatchString(`(?im)(calls)`, textBody)
	if err != nil {
		fmt.Printf("error in reg expression", err.Error())
	}
	if newTradeSubject {
		openNewSignal(email)
	} else {
		tradeSignal.ClientID = "Greg Manarino"
	}
	return err
}

func openNewSignal(email gmail.Email) {

}

func closeExistingSignal(email gmail.Email) {

}

// detectText gets text from the Vision API for an image at the given file path.
func detectText(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return err
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return err
	}
	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		return err
	}

	if len(annotations) == 0 {
		fmt.Fprintln(w, "No text found.")
	} else {
		fmt.Fprintln(w, "Text:")
		for _, annotation := range annotations {
			fmt.Fprintf(w, "%q\n", annotation.Description)
		}
	}

	return nil
}
