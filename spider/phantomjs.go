package spider

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/benbjohnson/phantomjs"
)

type DygodItem struct {
	Name string `json:",string"`
	Href string `json:",string"`
}

func init() {
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
	if err := page.Open("http://www.dy2018.com/html/tv/oumeitv/index.html"); err != nil {
		log.Fatal(err)
	}
	//注入js 文件必须在打开文件之后
	if err := page.InjectJS("jquery.min.js"); err != nil {
		log.Fatal(err)
	}

	// Read first link.
	array, err := page.Evaluate(`function() {
		var array = new Array();
		var host = location.host;
		$('ul > table > tbody > tr:nth-child(2) > td:nth-child(2) > b > a').each(function(index,item){
			var name = $(item).attr('title');
			var href = 'http://'+host + $(item).attr('href');
			array.push({name:name,href:href})
		});
		return array;
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
		//解析页面
		jsonUU, err := page.Evaluate(`function() {
		var h1 = $('h1').text();
		var body =$('#Zoom').text();
		return {title:h1,content:body};
	}`)

		if err != nil {
			log.Fatal(err)
		}

		dic, ok := jsonUU.(map[string]interface{})
		title := dic["title"].(string)
		content := dic["content"].(string)

		//储存数据结果
		fmt.Println(content, title)
	}

	// var array = new Array();
	// var array = $('#header > div > div.bd2 > div.bd3 > div.bd3r > div.co_area2 > div.co_content8 > ul > table > tbody > tr:nth-child(2) > td:nth-child(2) > b > a:nth-child(2)').each(function(index,value){var href = $(value).attr('href');var name=$(value).attr('title');array.push({href:href,name:name})});
}

func fetchDyGodPage(name, href string, page *phantomjs.WebPage) map[string]string {
	if err := page.Open(href); err != nil {
		log.Fatal(err)
	}

	//注入js 文件必须在打开文件之后
	if err := page.InjectJS("jquery.min.js"); err != nil {
		log.Fatal(err)
	}
	//解析页面
	jsonUU, err := page.Evaluate(`function() {
		var h1 = $('h1').text();
		var body =$('#Zoom').text();
		return {title:h1};
	}`)

	if err != nil {
		log.Fatal(err)
	}

	dic := jsonUU.(map[string]interface{})

	//储存数据结果
	fmt.Println(dic)

	info := map[string]string{
		"name": name,
		// "title":       dic["title"].(string),
		// "raw_content": dic["content"].(string),
		"raw_href": href,
	}
	fmt.Println("\n\n\n\n", info)
	return info

}
