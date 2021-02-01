package main

import (
	"fmt"
	"time"

	wf "github.com/mark2b/wpa-connect"
	"github.com/ohayao/common/http"
	"github.com/ohayao/common/reg"
)

func main() {
	pattern := `^\/user\/(?P<id>[0-9]{1,})?\/career\/(?P<position>.*$)`
	target := `/user/123456/career/officer`
	res := reg.GetNamedMap(pattern, target)
	fmt.Printf("%+v\n", res)
	//test()
	wifis()
	select {}
}

func test() {
	type day struct {
		Date          string `json:"date"`
		Low           string `json:"low"`
		High          string `json:"high"`
		WindDirection string `json:"fengxiang"`
		WindLevel     string `json:"fengli"`
		Message       string `json:"type"`
	}
	type DData struct {
		Yesterday *day   `json:"yesterday"`
		Forecast  []*day `json:"forecast"`
	}
	type data struct {
		Data   *DData
		Status int    `json:"status"`
		Desc   string `json:"desc"`
	}
	var d = new(data)
	res := http.New("http://wthrcdn.etouch.cn/weather_mini?city=深圳").Timeout(time.Second * 10)
	res.GetJSON(d)
	if res.Result.Err == nil && d.Data != nil && d.Data.Forecast != nil {
		for _, v := range d.Data.Forecast {
			fmt.Printf("%s %s %s %s %s %s\n", v.Date, v.Message, v.High, v.Low, v.WindDirection, v.WindLevel)
		}
	}
}

func wifis() {
	if bssList, err := wf.ScanManager.Scan(); err == nil {
		for _, bss := range bssList {
			print(bss.SSID, bss.Signal, bss.KeyMgmt)
		}
	} else {
		fmt.Println(err)
	}
}
