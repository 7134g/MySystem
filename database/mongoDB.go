package database

import (
	"MySystem/util"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MONGODBTT *mongo.Database
var MONGODBBG *mongo.Database
var MONGODBDF *mongo.Database

func MongoInit(connString string, DBnames []string) {
	// 建立链接
	var err error
	option := options.Client().ApplyURI(connString)
	MongoClient, err = mongo.Connect(context.Background(), option)
	if err != nil {
		util.Log().Panic("连接数据库不成功", err)
	}

	// 链接数据库
	for _, v := range DBnames {
		MongoConnDatabase(v)
	}

}

func MongoConnDatabase(DBname string) {
	if DBname == "tianting" {
		MONGODBTT = MongoClient.Database(DBname)
	} else if DBname == "blog" {
		MONGODBBG = MongoClient.Database(DBname)
	} else {
		MONGODBDF = MongoClient.Database(DBname)
	}
}
