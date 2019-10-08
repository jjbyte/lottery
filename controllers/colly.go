package controllers

import (
	"github.com/gocolly/colly"
	"lottery/config"
)

func getUrlBody(url string) (stockList string, err error) {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36"
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Host", "www.cwl.gov.cn")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Referer", "http://www.cwl.gov.cn/kjxx/ssq/kjgg/") //关键头 如果没有 则返回 错误
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
	})
	c.OnResponse(func(resp *colly.Response) {
		stockList = string(resp.Body)
	})
	c.OnError(func(resp *colly.Response, errHttp error) {
		err = errHttp
	})
	err = c.Visit(url)
	return
}

//爬大乐透和双色球开奖记录
func GetLottery(lotterytype uint) {
	var url	string
	if lotterytype == 0 {
		url = config.Conf.Get("app.welfareUrl").(string)
	} else {
		url = config.Conf.Get("app.sportUrl").(string)
	}

	url_body, err := getUrlBody(url)
	if err != nil {
		println(err.Error())
	}
	println(url_body)
}
