package gmail

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Message struct {
	MessageId string
	Sender    string
	Email     Email
}

var collection *mongo.Collection
var ctx = context.TODO()

// DBClient represent a connection with MongoDB
type DBClient struct {
	client *mongo.Client
	dbURI  string
	dbName string
}

// NewDBClient will set up a new MongoClient for the specified DBURI address.
func NewDBClient(dbURI string, dbName string, DbGcpPemLoc string) DBClient {
	err := godotenv.Load()
	//	fmt.Println("Connected to MongoDB!URI >>>>>>>>>>>>", dbURI)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	uri := fmt.Sprintf(dbURI, DbGcpPemLoc)
	credential := options.Credential{
		AuthMechanism: "MONGODB-X509",
	}
	clientOptions := options.Client().ApplyURI(uri).SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	//collection = client.Database("tasker").Collection("tasks")
	return DBClient{client, dbURI, dbName}
}

//AddMessageToDB upload messages to mongoDB
func (db *DBClient) AddMessageToDB(email Email) error {
	//we get message info and messagebody to upload to mongoDB
	client := db.client
	collection := client.Database(db.dbName).Collection("notifications")
	message := Message{email.Id, email.Sender, email}
	insertResult, err := collection.InsertOne(context.TODO(), message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
	return err
}
