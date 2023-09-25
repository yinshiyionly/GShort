package services

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gshort/internal/api/model"
	"log"
	"time"
)

var database string

func FindOriginURL(originURL string) (bool, string) {

	client, err := initMongo()
	if err != nil {
		log.Fatal(err)
	}
	// 关闭连接
	defer client.Disconnect(context.TODO())
	collection := client.Database(database).Collection("code_url_map")
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

func FindOriginURLByCode(code string) (bool, string) {

	client, err := initMongo()
	if err != nil {
		log.Fatal(err)
	}
	// 关闭连接
	defer client.Disconnect(context.TODO())
	collection := client.Database(database).Collection("code_url_map")
	filter := bson.D{{"code", code}}
	res := collection.FindOne(context.TODO(), filter)
	if err := res.Err(); err != nil {
		// 没有找到记录
		if err == mongo.ErrNoDocuments {
			return false, ""
		}
	}
	var result struct {
		OriginURL string `bson:"url"`
	}
	if err := res.Decode(&result); err != nil {
		log.Fatal(err)
	}
	return true, result.OriginURL
}

func InsertToMongo(params model.CodeUrlMap) interface{} {
	client, err := initMongo()
	if err != nil {
		log.Fatal(err)
	}
	database = viper.GetString("mongo.database")
	// 关闭连接
	fmt.Println(database)
	//defer client.Disconnect(context.TODO())
	collection := client.Database(database).Collection(model.TableName)
	objectId := primitive.NewObjectIDFromTimestamp(time.Now()).Hex()
	params.Code = GenerateTinyUrlByPrimaryKey(objectId)
	objectID, err := primitive.ObjectIDFromHex(objectId)
	if err != nil {
		log.Fatalf("Transformer objectID failed: %v", err)
	}
	fmt.Println(objectID)
	// 设置保存的数据
	insertParams := bson.D{
		{Key: "_id", Value: objectId},
		{Key: "code", Value: params.Code},
		{Key: "url", Value: params.Url},
	}
	//fmt.Println(params,
	//    primitive.NewObjectIDFromTimestamp(time.Now()).Hex(),
	//)
	result, err := collection.InsertOne(context.TODO(), insertParams)
	if err != nil {
		log.Fatalf("Insert Item Failed: %v", err)
	}
	return result
}

func initMongo() (*mongo.Client, error) {
	// 配置文件读取 user password
	user, password := viper.GetString("mongo.user"), viper.GetString("mongo.password")
	credential := options.Credential{
		Username: user,
		Password: password,
	}
	// 设置客户端连接配置
	host, port := viper.GetString("mongo.host"), viper.GetString("mongo.port")
	clientOptions := options.Client().ApplyURI("mongodb://" + host + ":" + port).SetAuth(credential)
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

//func init() {
//	database = viper.GetString("mongo.database")
//}
