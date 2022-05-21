package use

import (
	"fmt"
	"rango/app/controller"

	"github.com/gin-gonic/gin"

	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	controller.Base
}

// 任务的执行时间点
type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

// 一条日志
type LogRecord struct {
	JobName   string    `bson:"jobName"`   // 任务名
	Command   string    `bson:"command"`   // shell命令
	Err       string    `bson:"err"`       // 脚本错误
	Content   string    `bson:"content"`   // 脚本输出
	TimePoint TimePoint `bson:"timePoint"` // 执行时间点
}

func (this Mongo) Mongo(c *gin.Context) {
	var (
		client *mongo.Client
		result *mongo.InsertOneResult
		err    error
	)

	// 建立mongodb连接
	clientOptions := options.Client().ApplyURI("mongodb://192.168.43.123:27017")
	if client, err = mongo.Connect(context.TODO(), clientOptions); err != nil {
		return
	}

	// 2, 选择数据库my_db
	database := client.Database("banana")

	// 3, 选择表my_collection
	collection := database.Collection("crontab")
	// 4, 插入记录(bson)
	record := &LogRecord{
		JobName:   "job10",
		Command:   "echo hello",
		Err:       "",
		Content:   "hello",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}

	result, err = collection.InsertOne(context.TODO(), record)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	// _id: 默认生成一个全局唯一ID, ObjectID：12字节的二进制
	docId := result.InsertedID.(primitive.ObjectID)
	fmt.Println("自增ID:", docId.Hex())

	c.JSON(200, gin.H{
		"code":     200,
		"message":  "success",
		"ObjectID": docId.Hex(),
	})
}
