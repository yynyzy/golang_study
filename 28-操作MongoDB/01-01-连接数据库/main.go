package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func initDB() {
	// 设置客户端连接配置
	connect := options.Client().ApplyURI("mongodb://localhost:27017")
	// 连接到MongoDB
	var err error
	c, err2 := mongo.Connect(context.TODO(), connect)
	if err2 != nil {
		log.Fatal(err2)
	}
	// 检查连接
	err3 := c.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err3)
	}
	fmt.Println("mongodb connected")
}

func main() {
	initDB()
}
