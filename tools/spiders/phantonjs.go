package spiders

import (
	//https://github.com/benbjohnson/phantomjs
	"encoding/json"
	"fmt"
	"my_go_web/models"
	"net/http"
	"os"

	"github.com/astaxie/beego/orm"
	"github.com/benbjohnson/phantomjs"
)

func init() {

	//读取数据库爬虫配置
	var spiderConfig = new(models.Spiders)
	var o = orm.NewOrm()
	if err := o.QueryTable(spiderConfig).Filter("name", "ttmj").One(spiderConfig); err != nil {
		fmt.Println(err)
	}
	if err := phantomjs.DefaultProcess.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer phantomjs.DefaultProcess.Close()

	//初始化页面
	var page, err = phantomjs.CreateWebPage()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//设置web page
	web_setting := phantomjs.WebPageSettings{
		JavascriptEnabled:             true,
		LoadImages:                    false,
		LocalToRemoteURLAccessEnabled: true,
		UserAgent:                     spiderConfig.UserAgent,
		Username:                      spiderConfig.UserName,
		Password:                      spiderConfig.Password,
		XSSAuditingEnabled:            false,
		WebSecurityEnabled:            false,
		ResourceTimeout:               15000,
	}
	if err := page.SetSettings(web_setting); err != nil {
		fmt.Println(err)
	}
	//设置headers
	web_headers := http.Header{}
	if err := page.SetCustomHeaders(web_headers); err != nil {
		fmt.Println(err)
	}

	//如果需要登录验证
	web_cookies := []*http.Cookie{}
	if err := page.SetCookies(web_cookies); err != nil {
		fmt.Println(err)
	}
	// Open a URL.
	var url = fmt.Sprintf("%s%s", spiderConfig.Host, spiderConfig.Uri)
	if err := page.Open(url); err != nil {
		fmt.Println(err)
	}
	//测试外部应用js
	page.IncludeJS("https://cdn.bootcss.com/jquery/3.2.1/jquery.min.js")

	//使用js代码来处理页面信息 存储json数据到数据库
	info, err := page.Evaluate(`function() {
		return  $('body > div.g-wrap > div.m-botnav.rel > div > div > ul > li:nth-child(5) > a').attr('href');
	}`)
	if err != nil {
		fmt.Println(err)
	}
	// Print title and URL.
	link := info.(string)
	fmt.Println(link)
	title, _ := page.FrameURL()
	fmt.Println(title)
	//获取cookie 储存到db
	if cookies, err := page.Cookies(); err != nil {
		fmt.Println(err)
	} else {
		if json, err := json.Marshal(cookies); err == nil {
			spiderConfig.Cookies = string(json)
		} else {
			fmt.Println(err)
		}

		if num, err := o.Update(spiderConfig, "Cookies"); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(num)
		}
	}
}
