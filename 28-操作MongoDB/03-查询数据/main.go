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

func findData() {
	defer client.Disconnect(context.TODO())
	c := client.Database("go_db").Collection("student")
	// c2, _ := c.Find(context.Background(), bson.D{{Key: "name", Value: "wang fan"}})
	// c2, _ := c.Find(context.Background(), bson.D{{"name", "wang fan"}})
	c2, _ := c.Find(context.Background(), bson.D{})
	defer c2.Close(context.Background())
	for c2.Next(context.Background()) {
		var result bson.D
		err := c2.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("result: %v\n", result)
		fmt.Printf("result.Map(): %v\n", result.Map())
		fmt.Printf("result.Map()[\"name\"]: %v\n", result.Map()["name"])
	}
	if err := c2.Err(); err != nil {
		log.Fatal(err)
	}
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
	findData()
}
