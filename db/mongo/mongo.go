package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	MongoDB struct {
		client   *mongo.Client
		database *mongo.Database
	}
)

// Connect 创建一个新的连接
// uri 连接字符串
// timeout 超时时间
// ctx 上下文 eg：context.WithTimeout(context.Background(), time.Second*20)
func Connect(uri string, timeout time.Duration, ctx context.Context) (*MongoDB, error) {
	if client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri)); err != nil {
		return nil, err
	} else {
		return &MongoDB{client: client}, nil
	}
}

// DisConnect 断开连接
func (that *MongoDB) Disconnect(ctx context.Context) error {
	return that.client.Disconnect(ctx)
}

// SelectDC 选择数据库以及集合
func (that *MongoDB) SelectDC(dbName, collectionName string) *mongo.Collection {
	return that.client.Database(dbName).Collection(collectionName)
}

// SetDB 设置需要连接的数据库
// 多数据库使用时慎重
func (that *MongoDB) SetDB(dbName string) *MongoDB {
	that.database = that.client.Database(dbName)
	return that
}

// C 集合 需要在SetDB后使用
// collectionName 集合名字
func (that *MongoDB) C(collectionName string) *mongo.Collection {
	if that.database == nil {
		return nil
	}
	return that.database.Collection(collectionName)
}

// GetClient
func (that *MongoDB) GetClient() *mongo.Client {
	return that.client
}
