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

func update() {
	defer client.Disconnect(context.TODO())
	c := client.Database("go_db").Collection("student")
	update := bson.D{{"$set", bson.D{{Key: "name", Value: "王帆"}, {Key: "age", Value: 18}}}}
	// 更新条件 name wang fan 更新问 王帆 年龄 18
	updateResult, err := c.UpdateMany(context.TODO(), bson.D{{"name", "wang fan"}}, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("updateResult.ModifiedCount: %v\n", updateResult.ModifiedCount)
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
	update()
}
