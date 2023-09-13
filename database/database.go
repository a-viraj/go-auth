package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func DBinstance() *mongo.Client{
	err:=godotenv.Load(".env")
	if err!=nil{
		log.Fatal("Error loading the env")
	}
	MondoDb:=os.Getenv("MONGOURL")
	client,err:=mongo.NewClient(options.Client().ApplyURI(MondoDb))
	if err!=nil{
		log.Fatal(err)
	}
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	err =client.Connect(ctx)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("connected")
	return client
}
var Client *mongo.Client=DBinstance()

func OpenCollection(client *mongo.Client,collectionName string) *mongo.Collection{
	var coll *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	return coll
}