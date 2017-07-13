package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/astaxie/beego"
)

type VideoController struct {
	beego.Controller //集成beego controller

}

func (c *VideoController) WeiboVideoParse() {
	fid := c.GetString("id")
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
	c.Data["json"] = ""
	if resp, err := client.Do(req); err == nil && resp.StatusCode == 200 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		//<video id=.*?src=[\'"](.*?)[\'"]\W
		reg := regexp.MustCompile(`\<video id=.*?src=[\'"](.*?)[\'"]\W`)
		srcs := reg.FindStringSubmatch(bodyString)
		for k, v := range srcs {
			if k == 1 {
				c.Data["json"] = v
			}
		}
	}
	c.ServeJSON()

}
