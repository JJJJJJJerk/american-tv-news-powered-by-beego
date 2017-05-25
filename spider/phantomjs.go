package spider

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/benbjohnson/phantomjs"
)

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
		$('ul > table > tbody > tr:nth-child(2) > td:nth-child(2) > b > a').each(function(index,item){
			var name = $(item).attr('title');
			var href = $(item).attr('href');
			array[index] = {name:name,href:href}
		});
		return array;
	}`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("json", array)

}
