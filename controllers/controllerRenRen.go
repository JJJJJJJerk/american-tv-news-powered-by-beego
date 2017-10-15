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

const SERVER = "https://api.rr.tv"
const SECRET_KEY = "clientSecret=08a30ffbc8004c9a916110683aab0060"

//const TOKEN = "945e82b94c08447aafe45e6051159737"
const TOKEN = "35dafc24e73d461a94c2e69daef72c6b" //我的iOStoken

var TOKENS = []string{
	"91b7fc36672548259433ca40bc57e6dd", //我的
	"35dafc24e73d461a94c2e69daef72c6b", //我的
	"99ce6f3f7ff840e586958770472ec893",
	"37570aa5dabe44219ab8a13986778390",
	"6f8443b2d6c54d39bf7ab780bb5df3b2",
	"3739e42b649948febad6dbfeaeecddfe",
	"ffeb43e117d744e6822c3fe6abc12dea",
	"5f6d0d3a49e5439e8492515c49b797c5",
	"41d74adccd4b46389882242949e82cea",
	"ff750a61e2684b7d9d4f783b4e3b14b1",
	"f703dcbae2fb42889e8b111829b36167",
	"9aea227ebb224a8b9bea347c13192706",
	"827af86209464b038552a87f1cd546b2",
	"bfc156b8c9ce48d98f2f9e48868130b3",
	"c0a60ecf198d458d83b475e50b8d7893",
	"22f6e8f0fac049849ce6db26535ecf6d",
	"cce5bae0fcdd446fa87f0e61d86cd345",
	"d1f33f20d83441228adf791215308d6d",
	"803db2c8c9de400a9e90f75c2926d98a",
	"b7c4f78d734f48fb871b3c7667f32019",
	"99b3109e62fa4f1eb449182f5428cb84",
	"10347e508c5946ec9116e3044b728f01",
}

var tokenIndex = 0

func generateFakeHeader() (headers http.Header) {
	// var randomToken = TOKENS[tokenIndex%(len(TOKENS))]
	return http.Header{
		"token":          {"6b6cfdd3e90843c0a0914425638db7ef"},
		"clientType":     {"android_RRMJ"},
		"clientVersion":  {"3.6.3"},
		"deviceId":       {"861134030056126"},
		"signature":      {"643c184f77372e364550e77adc0360cd"},
		"t":              {"1491433993933"},
		"Content-Type":   {"application/x-www-form-urlencoded"},
		"Authentication": {"RRTV 470164b995ea4aa5a53f9e5cbceded472:IxIYBj:LPWfRb:I9gvePR5R2N8muXD7NWPCj"},
	}
}

type RenRenController struct {
	beego.Controller //集成beego controller
}

func (c *RenRenController) Index() {
	apiURI := "/v3plus/video/indexInfo"
	cacheKey := fmt.Sprint(SERVER, apiURI)
	paraData := url.Values{}
	c.cacheOrPostReturnJson(cacheKey, apiURI, paraData, generateFakeHeader())
}

//根据episodeId找到对应的电视剧播放地址
func (c *RenRenController) M3u8() {
	//55872
	var FakeHeader = generateFakeHeader()
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
	c.cacheOrPostReturnJson(cacheKey, apiURI, paraData, generateFakeHeader())
}

func (c *RenRenController) Hot() {
	apiURI := "/video/seasonRankingList"
	cacheKey := fmt.Sprint(SERVER, apiURI)
	paraData := url.Values{}
	c.cacheOrPostReturnJson(cacheKey, apiURI, paraData, generateFakeHeader())
}

func (c *RenRenController) Top() {
	apiURI := "/v3plus/season/topList"
	cacheKey := fmt.Sprint(SERVER, apiURI)
	paraData := url.Values{}
	c.cacheOrPostReturnJson(cacheKey, apiURI, paraData, generateFakeHeader())
}

func (c *RenRenController) Season() {
	apiURI := "/v3plus/season/detail"
	seasonId := c.GetString(":seasonId")
	cacheKey := fmt.Sprint(SERVER, apiURI, "/seasonId/", seasonId)
	paraData := url.Values{
		"seasonId": {seasonId},
	}
	c.cacheOrPostReturnJson(cacheKey, apiURI, paraData, generateFakeHeader())
}

func (c *RenRenController) Album() {
	apiURI := "/v3plus/video/album"
	albumId := c.GetString(":albumId")
	cacheKey := fmt.Sprint(SERVER, apiURI, "/albumId/", albumId)
	paraData := url.Values{
		"albumId": {albumId},
	}
	c.cacheOrPostReturnJson(cacheKey, apiURI, paraData, generateFakeHeader())
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
	c.cacheOrPostReturnJson(cacheKey, apiURI, paraData, generateFakeHeader())
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
