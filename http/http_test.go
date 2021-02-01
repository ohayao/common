package http

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	res := New("https://www.baidu.com").Timeout(time.Second * 10)
	t.Log(res.GetString())
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		weater := New("http://t.weather.sojson.com/api/weather/city/101280601").Timeout(time.Second * 10)
		weater.GetString()
	}
}
