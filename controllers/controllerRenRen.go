package controllers

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"my_go_web/models"
	"net/http"
	"net/url"

	"github.com/astaxie/beego"
)

const SERVER = "http://api.rr.tv"
const SECRET_KEY = "clientSecret=08a30ffbc8004c9a916110683aab0060"
const TOKEN = "5f8f489d12f64488aa310334f32153b4"

var FakeHeader = http.Header{
	"token":         {TOKEN},
	"clientType":    {"android_%E8%B1%8C%E8%B1%86%E8%8D%9A"},
	"clientVersion": {"3.5.3.1"},
	"deviceId":      {"861134030056126"},
	"signature":     {"643c184f77372e364550e77adc0360cd"},
	"t":             {"1491433993933"},
	"Content-Type":  {"application/x-www-form-urlencoded"},
}

type RenRenController struct {
	beego.Controller //集成beego controller
}

func (c *RenRenController) Index() {
	apiURI := "/v3plus/video/indexInfo"
	cacheKey := fmt.Sprint(SERVER, apiURI)
	paraData := url.Values{}
	c.cacheOrPostReturnJson(cacheKey, apiURI, paraData, FakeHeader)
}

//根据episodeId找到对应的电视剧播放地址
func (c *RenRenController) M3u8() {
	//55872
	var episodeSid = c.GetString(":episodeSid")
	//https://github.com/wilsonwen/kanmeiju/blob/adc56d8665b3c99a8d48df3cc2f1eaf623f7be6e/index.js line203
	apiURI := "/video/findM3u8ByEpisodeSid"
	cacheKey := fmt.Sprint(SERVER, apiURI, "/episodeSid/", episodeSid)

	paraData := url.Values{
		"episodeSid": {episodeSid},
		"quality":    {"super"},
		"token":      {TOKEN},
		"seasonId":   {"0"},
	}
	key := fmt.Sprint("episodeSid=", episodeSid, "quality=super", "clientType=", FakeHeader["clientType"][0], "clientVersion=", FakeHeader["clientVersion"][0], "t=", FakeHeader["t"][0], SECRET_KEY)
	signature := GetMD5Hash(key)
	FakeHeader["signature"] = []string{signature}
	c.cacheOrPostReturnJson(cacheKey, apiURI, paraData, FakeHeader)

}
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (c *RenRenController) Search() {
	searchWord := c.GetString(":keyword")
	//https://github.com/wilsonwen/kanmeiju/blob/adc56d8665b3c99a8d48df3cc2f1eaf623f7be6e/index.js line203
	apiURI := "/v3plus/video/search"
	cacheKey := fmt.Sprint(SERVER, apiURI, "/name/", searchWord)
	paraData := url.Values{
		"title":   {searchWord},
		"enTitle": {searchWord},
	}
	c.cacheOrPostReturnJson(cacheKey, apiURI, paraData, FakeHeader)
}

func (c *RenRenController) Hot() {
	apiURI := "/video/seasonRankingList"
	cacheKey := fmt.Sprint(SERVER, apiURI)
	paraData := url.Values{}
	c.cacheOrPostReturnJson(cacheKey, apiURI, paraData, FakeHeader)
}

func (c *RenRenController) Top() {
	apiURI := "/v3plus/season/topList"
	cacheKey := fmt.Sprint(SERVER, apiURI)
	paraData := url.Values{}
	c.cacheOrPostReturnJson(cacheKey, apiURI, paraData, FakeHeader)
}

func (c *RenRenController) Season() {
	apiURI := "/v3plus/season/detail"
	seasonId := c.GetString(":seasonId")
	cacheKey := fmt.Sprint(SERVER, apiURI, "/seasonId/", seasonId)
	paraData := url.Values{
		"seasonId": {seasonId},
	}
	c.cacheOrPostReturnJson(cacheKey, apiURI, paraData, FakeHeader)
}

func (c *RenRenController) Album() {
	apiURI := "/v3plus/video/album"
	albumId := c.GetString(":albumId")
	cacheKey := fmt.Sprint(SERVER, apiURI, "/albumId/", albumId)
	paraData := url.Values{
		"albumId": {"albumId"},
	}
	c.cacheOrPostReturnJson(cacheKey, apiURI, paraData, FakeHeader)
}

func (c *RenRenController) Category() {
	apiURI := "/v3plus/video/search"
	cat := c.GetString(":categoryType")
	pages := c.GetString(":pages")
	cacheKey := fmt.Sprint(SERVER, apiURI, "/category/", cat, "/pages/", pages)
	paraData := url.Values{
		"cate":  {"cat_list"},
		"cat":   {cat},
		"pages": {pages},
	}
	c.cacheOrPostReturnJson(cacheKey, apiURI, paraData, FakeHeader)
}

func (c *RenRenController) cacheOrPostReturnJson(cacheKey string, apiURI string, paraData url.Values, headers http.Header) {
	apiUrl := fmt.Sprint(SERVER, apiURI)
	var content []byte
	if x, found := models.CacheManager.Get(cacheKey); found {
		foo := x.(string)
		content = []byte(foo)
	} else {
		client := &http.Client{}
		datUrl := paraData.Encode()
		req, _ := http.NewRequest("POST", apiUrl, bytes.NewBufferString(datUrl))
		req.Header = headers
		resp, err := client.Do(req)
		defer resp.Body.Close()
		if err == nil && resp.StatusCode == 200 {
			body, _ := ioutil.ReadAll(resp.Body)
			content = body
			//放缓存
			models.CacheManager.Set(cacheKey, string(content), models.C_EXPIRE_TIME_HOUR_01)
		}
	}

	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	c.Ctx.Output.Body(content)
	return
}
