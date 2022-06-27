package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func delete() {
	defer client.Disconnect(context.TODO())
	c := client.Database("go_db").Collection("student")
	dr, err := c.DeleteMany(context.TODO(), bson.D{{Key: "name", Value: "王帆"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("dr.DeletedCount: %v\n", dr.DeletedCount)
}

func initDB() {
	connect := options.Client().ApplyURI("mongodb://localhost:27017")
	c, err := mongo.Connect(context.TODO(), connect)
	if err != nil {
		log.Fatal(err)
	}
	err2 := c.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err2)
	}
	fmt.Println("mongodb connected")
	client = c
}

func main() {
	initDB()
	delete()
}
