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

	c, err := mongo.Connect(context.TODO(), connect)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err2 := c.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err2)
	}
	fmt.Println("mongodb connected")
	client = c
}

type Student struct {
	Name string
	Age  int
}

//添加一条
func insertOne(s Student) {
	c := client.Database("go_db").Collection("student")
	ior, err := c.InsertOne(context.TODO(), s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ior.InsertedID: %v\n", ior.InsertedID)
}

//添加多条
func insertMore(students []interface{}) {
	initDB()
	c := client.Database("go_db").Collection("student")
	imr, err := c.InsertMany(context.TODO(), students)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("imr.InsertedIDs: %v\n", imr.InsertedIDs)
}

func main() {
	initDB()
	s := Student{Name: "wang fan", Age: 26}
	insertOne(s)

	// s1 := Student{Name: "wang fan1", Age: 26}
	// s2 := Student{Name: "wang fan2", Age: 26}
	// s3 := Student{Name: "wang fan3", Age: 26}
	// students := []interface{}{s1, s2, s3}
	// insertMore(students)
}
