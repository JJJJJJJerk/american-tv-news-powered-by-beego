package controllers

import (
	"my_go_web/spider"
	"strings"
)

type VideoController struct {
	BaseController
}

func (c *VideoController) VideoUrlConvertToMp4Url() {
	url := c.GetString("url", "")
	spider.RunQQVideoParser(url)

	var mp4URL string
	if strings.Contains(url, "youku.com") {
		mp4URL = spider.RunYoukuVideoParser(url)
		c.JsonRetrun("success", "parse video success", mp4URL)

	}
	if strings.Contains(url, "miaopai.com") {
		//todo:miaopai
	}
	c.JsonRetrun("error", "不支持视频解析地址", nil)
}
