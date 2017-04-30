package spiders

//使用go自带的http client

import (
	//https://github.com/benbjohnson/phantomjs
	"fmt"
	"my_go_web/models"
	"net/http"
	"net/url"

	"encoding/json"

	"github.com/astaxie/beego/orm"
)

func init() {
	//SpiderZimuzuAuth()
	SpiderTtmjAuth()
}

func SpiderZimuzuAuth() {
	//最好过一段时间7天刷新一次

	//读取数据库爬虫配置
	var spiderConfig = new(models.Spiders)
	var o = orm.NewOrm()
	if err := o.QueryTable(spiderConfig).Filter("name", "zimuzu").One(spiderConfig); err != nil {
		fmt.Println(err)
	}

	//post auth request
	url_back := fmt.Sprintf("%s%s", spiderConfig.Host, spiderConfig.Uri)
	var formPara = url.Values{"account": {spiderConfig.UserName}, "passowrd": {spiderConfig.Password}, "remember": {"1"}, "url_back": {url_back}}
	resp, err := http.PostForm(spiderConfig.AuthUrl, formPara)
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode != 200 {
		fmt.Println("auth not 200 success")

	}
	headerJsonBytes, err := json.Marshal(resp.Header)
	if err != nil {
		fmt.Println(err)
	}
	spiderConfig.Headers = string(headerJsonBytes)
	cookiesJsonBytes, err := json.Marshal(resp.Cookies())
	if err != nil {
		fmt.Println(err)
	}
	spiderConfig.Cookies = string(cookiesJsonBytes)
	spiderConfig.Status = 2 //2:已经登录保存了 登录cookie
	o.Update(spiderConfig)
}

func SpiderTtmjAuth() {
	//最好过一段时间7天刷新一次

	//读取数据库爬虫配置
	var spiderConfig = new(models.Spiders)
	var o = orm.NewOrm()
	if err := o.QueryTable(spiderConfig).Filter("name", "ttmj").One(spiderConfig); err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", spiderConfig.AuthUrl, nil)
	if err != nil {
		fmt.Println(err)
	}
	//
	req.Header.Set("Referer", "https://www.baidu.com/link?url=mdtKKBRlTmwpOvtr-wn_HHu58qR_wG3wukSQQAy-jTIRjV_asvHuoKlsL6pIidqt&wd=&eqid=c56896ac00066b56000000025906150b")
	req.Header.Set("User-Agent", spiderConfig.UserAgent)
	//send request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode != 200 {
		fmt.Println("auth not 200 success")

	}

	headerJsonBytes, err := json.Marshal(resp.Header)
	if err != nil {
		fmt.Println(err)
	}
	spiderConfig.Headers = string(headerJsonBytes)
	cookiessss := resp.Cookies()
	cookiesJsonBytes, err := json.Marshal(cookiessss)
	if err != nil {
		fmt.Println(err)
	}
	spiderConfig.Cookies = string(cookiesJsonBytes)
	spiderConfig.Status = 2 //2:已经登录保存了 登录cookie
	o.Update(spiderConfig)
}
