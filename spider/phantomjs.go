package spider

import (
	"fmt"
	"log"
	"my_go_web/models"
	"os"
	"path/filepath"

	"strings"

	"github.com/benbjohnson/phantomjs"
)

//爬去美剧天堂的美剧
func RunDygodMeijuSpider() {
	//获取软件的根目录
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	if err := phantomjs.DefaultProcess.Open(); err != nil {
		log.Fatal(err)
	}
	defer phantomjs.DefaultProcess.Close()
	//穿件webpage
	page, err := phantomjs.CreateWebPage()
	if err != nil {
		log.Fatal(err)
	}
	defer page.Close()
	//设置webpage配置
	webPageSettings := phantomjs.WebPageSettings{
		JavascriptEnabled:             true,
		LoadImages:                    false,
		LocalToRemoteURLAccessEnabled: true, //local script can asscess the remote files
		UserAgent:                     "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
		Username:                      "dejavuzhou@qq.com",
		Password:                      "ZHou1987",
		XSSAuditingEnabled:            false,
		WebSecurityEnabled:            false,
		ResourceTimeout:               20}

	if err := page.SetSettings(webPageSettings); err != nil {
		log.Fatal(err)
	}
	// 设置外部js文件的路径
	jqueryFileDir := fmt.Sprintf("%s%s", dir, "/spider")
	if err := page.SetLibraryPath(jqueryFileDir); err != nil {
		log.Fatal(err)
	}
	if err := page.Open("http://www.dygod.net/html/tv/oumeitv/"); err != nil {
		log.Fatal(err)
	}
	//注入js 文件必须在打开文件之后
	if err := page.InjectJS("jquery.min.js"); err != nil {
		log.Fatal(err)
	}

	//注入爬虫内容处理js
	if err := page.InjectJS("dygod.js"); err != nil {
		log.Fatal(err)
	}

	// Read first link.
	array, err := page.Evaluate(`function() {
		return GetDygodMeijuList();
	}`)
	if err != nil {
		log.Fatal(err)
	}
	items := array.([]interface{})

	for _, item := range items {
		node, ok := item.(map[string]interface{})
		// e is the
		name := node["name"].(string)
		href := node["href"].(string)
		fmt.Println(name, href, ok)

		if err := page.Open(href); err != nil {
			log.Fatal(err)
		}

		//注入js 文件必须在打开文件之后
		if err := page.InjectJS("jquery.min.js"); err != nil {
			log.Fatal(err)
		}
		//注入爬虫内容处理js
		if err := page.InjectJS("dygod.js"); err != nil {
			log.Fatal(err)
		}

		//解析页面
		jsonUU, err := page.Evaluate(`function() {
			 return getDygodMeijuDetailJson();
		}`)
		//time.Sleep(time.Second)

		if err != nil {
			log.Fatal(err)
		}

		dic, ok := jsonUU.(map[string]interface{})
		title := dic["title"].(string)
		content := dic["content"].(string)

		//储存数据结果
		//fmt.Println(content, title)
		//写入数据库
		url := strings.TrimSpace(href)
		var article models.Article
		models.Gorm.Where(models.Article{UrlProvider: url, Title: title}).Assign(models.Article{RawTitle: title, Body: content, RawContent: content, Title: title}).FirstOrCreate(&article)
		articleTag := models.ArticleTag{ArticleId: article.ID, TagId: 3}
		models.Gorm.Where(articleTag).Assign(articleTag).FirstOrCreate(&articleTag)
		//fmt.Println(article)
	}

	// var array = new Array();
	// var array = $('#header > div > div.bd2 > div.bd3 > div.bd3r > div.co_area2 > div.co_content8 > ul > table > tbody > tr:nth-child(2) > td:nth-child(2) > b > a:nth-child(2)').each(function(index,value){var href = $(value).attr('href');var name=$(value).attr('title');array.push({href:href,name:name})});
}
