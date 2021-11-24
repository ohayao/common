package mongo

import (
	"context"
	"fmt"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	connstr = "....."
)

func TestQuery(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	adb, err1 := Connect(connstr, 20, ctx)
	defer adb.Disconnect(ctx)
	if err1 != nil {
		t.Fatal(err1)
	} else {
		adb.SetDB("sample_analytics")
		_Test(adb)
	}
}

func BenchmarkConnect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()
		db, err1 := Connect(connstr, 20, ctx)
		if err1 != nil {
			b.Fatal(err1)
		} else {
			c2, _ := db.SelectDC("sample_analytics", "accounts").CountDocuments(context.TODO(), bson.D{})
			if c2 != 1746 {
				b.Fail()
			}
		}
	}
}

func _Test(that *MongoDB) {
	ch := make(chan bool)
	res := 0
	that.SetDB("igloohome-mongodb")
	go func() {
		for i := 0; i < 100; i++ {
			datas, err := that.C("c_user").CountDocuments(context.TODO(), bson.D{})
			fmt.Println(i, " user ", datas, err)
		}
		res++
		if res == 2 {
			ch <- true
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			datas, err := that.C("test").CountDocuments(context.TODO(), bson.D{})
			fmt.Println(i, " test ", datas, err)
		}
		res++
		if res == 2 {
			ch <- true
		}
	}()
	<-ch
}
