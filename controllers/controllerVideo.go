package controllers

import (
	"fmt"
	"io/ioutil"
	"www.mojotv.cn/models"
	"net/http"
	"regexp"

	"github.com/astaxie/beego"
)

type VideoController struct {
	beego.Controller //集成beego controller

}

func (c *VideoController) WeiboVideoParse() {
	fid := c.GetString("id")
	key := fmt.Sprint("mp4.", fid)
	if x, found := models.CacheManager.Get(key); found {
		c.Data["json"] = x.(string)
	} else {
		mp4Url := c.fetchMp4Url(fid)
		c.Data["json"] = mp4Url
		//这个过期时间也许还需要调整 进过验证有效期为一小时
		models.CacheManager.Set(key, mp4Url, models.C_EXPIRE_TIME_MIN_01*55)
	}
	c.ServeJSON()
}

func (c *VideoController) fetchMp4Url(fid string) (mp4Url string) {
	url := fmt.Sprint("http://video.weibo.com/show?fid=", fid, "&type=mp4")
	//reg := regexp.MustCompile()
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Referer", "http://weibo.com/home?wvr=5")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 4.4.2; Nexus 4 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Mobile Safari/537.36")
	req.Header.Add("Host", "www.miaopai.com")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accept-Language", "en-US,en;q=0.8")
	req.Header.Add("Accept-Charset", "UTF-8,*;q=0.5")
	mp4Url = ""
	if resp, err := client.Do(req); err == nil && resp.StatusCode == 200 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		//<video id=.*?src=[\'"](.*?)[\'"]\W
		reg := regexp.MustCompile(`\<video id=.*?src=[\'"](.*?)[\'"]\W`)
		srcs := reg.FindStringSubmatch(bodyString)
		for k, v := range srcs {
			if k == 1 {
				mp4Url = v
			}
		}
	}
	return
}
