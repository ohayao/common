package main

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/ohayao/common/db/mongo"
	"github.com/ohayao/common/file"
	"github.com/ohayao/common/http"
	"github.com/ohayao/common/reg"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	eg()
	db()
	select {}
}
func eg() {
	request := http.New("http://localhost:80/api")
	request.PostBytes(bytes.NewBuffer([]byte("")))

	pattern := `^\/user\/(?P<id>[0-9]{1,})?\/career\/(?P<position>.*$)`
	target := `/user/123456/career/officer`
	res := reg.GetNamedMap(pattern, target)
	fmt.Printf("%+v\n", res)
}

func db() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	connStr := file.GetFileContent(".configs/mongodb_str")
	db, err := mongo.Connect(connStr, time.Second*20, ctx)
	if err != nil {
		return
	}
	defer db.Disconnect(ctx)
	fmt.Println(db.SelectDC("sample_analytics", "accounts").CountDocuments(context.TODO(), bson.D{}))
	fmt.Println(db.GetClient().ListDatabaseNames(context.TODO(), bson.D{}))
}
