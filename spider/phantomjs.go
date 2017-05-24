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
		LocalToRemoteURLAccessEnabled: true,
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
	if err := page.Open("http://www.163.com/"); err != nil {
		log.Fatal(err)
	}
	//注入js 文件必须在打开文件之后
	if err := page.InjectJS("jquery.min.js"); err != nil {
		log.Fatal(err)
	}

	// Setup the viewport and render the results view.
	if err := page.SetViewportSize(640, 960); err != nil {
		log.Fatal(err)
	}
	if err := page.Render("haitou.png", "png", 100); err != nil {
		log.Fatal(err)
	}
	// Read first link.
	info, err := page.Evaluate(`function() {
		var link = $('head > meta:nth-child(6)').attr('content');
		return link;
	}`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("json", info)

}
