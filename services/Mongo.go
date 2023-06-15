package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func FindOriginURL(originURL string) (bool, string) {
	client, err := initMongo()
	if err != nil {
		log.Fatal(err)
	}
	// 关闭连接
	defer client.Disconnect(context.TODO())
	collection := client.Database("gshort").Collection("code_url_map")
	filter := bson.D{{"url", originURL}}
	res := collection.FindOne(context.TODO(), filter)
	if err := res.Err(); err != nil {
		// 没有找到记录
		if err == mongo.ErrNoDocuments {
			return false, ""
		}
	}
	var result struct {
		Code string `bson:"code"`
	}
	if err := res.Decode(&result); err != nil {
		log.Fatal(err)
	}
	return true, result.Code
}

func initMongo() (*mongo.Client, error) {
	credential := options.Credential{
		Username: "root",
		Password: "example",
	}
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017").SetAuth(credential)
	// 连接到 MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}
