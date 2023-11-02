package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	url := "mongodb://localhost:27017"

	cli, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Disconnect(context.Background())

	err = cli.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	g := newGeneral()

	collection := cli.Database("ss").Collection("general")

	result, err := collection.InsertOne(context.TODO(), g)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var results []bson.M

	err = cursor.All(context.Background(), &results)
	if err != nil {
		log.Fatal(err)
	}

	printJson(results)

}

func printJson(v any) {
	result, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("bind json failed")
	}
	fmt.Println(string(result))
}
