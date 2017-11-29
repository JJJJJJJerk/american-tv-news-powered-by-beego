package controllers

import (
	"fmt"
	"time"

	"io/ioutil"
	"math/rand"
	"www.mojotv.cn/models"
	"net/http"

	"github.com/astaxie/beego"
)

type FantasyController struct {
	beego.Controller //集成beego controller
}

func (c *FantasyController) Index() {

	myChannelId := c.GetString(":mcid")
	channelId := c.GetString(":cid", "33716")
	vId := c.GetString(":vid", "87816")
	t := fmt.Sprint(rand.ExpFloat64())
	//https://www.fantasy.tv/tv/playDetails.action?myChannelId=33716&id=87816&channelId=33716&t=0.0958195376527875
	cacheKey := fmt.Sprintf("myChannelId=%s&id=%s&channelId=%s", myChannelId, vId, channelId)
	var content []byte
	if x, found := models.CacheManager.Get(cacheKey); found {
		foo := x.(string)
		content = []byte(foo)
	} else {
		client := &http.Client{}
		apiUrl := fmt.Sprintf("https://www.fantasy.tv/tv/playDetails.action?myChannelId=%s&id=%s&channelId=%s&t=%s", myChannelId, vId, channelId, t)
		fmt.Println(apiUrl)
		req, _ := http.NewRequest("GET", apiUrl, nil)
		req.Header = http.Header{
			"Cookie":     {"acw_tc=AQAAAB4/YV0LXwIA8cxxq7tfQBTcUPei"},
			"Referer":    {"https://www.fantasy.tv/newApp/index.html"},
			"User-Agent": {"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1"},
			"Host":       {"www.fantasy.tv"},
		}
		resp, err := client.Do(req)

		defer resp.Body.Close()
		if err == nil && resp.StatusCode == 200 {
			body, _ := ioutil.ReadAll(resp.Body)
			content = body
			//放缓存
			models.CacheManager.Set(cacheKey, string(content), time.Second*90)
		}
	}
	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	c.Ctx.Output.Body(content)
	return
}
