package spider

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/benbjohnson/phantomjs"
)

type VideoJson struct {
	code  string
	video string
}

//爬去优酷视频真实地址
func RunQQVideoParser(youkuVideoUrl string) (mp4url string) {
	//源地址 http://v.youku.com/v_show/id_XMjY5MjQ3NDQ2NA==.html?spm=a2hfu.20010077.m_210490.5~5!2~5~5!3~5~5!3~5~A&f=29102521&from=y1.3-fun-fun-904-10077.90023-210490.3-4
	//提换之后的地址 http://m.youku.com/video/id_XMjY5MjQ3NDQ2NA==.html?spm=a2hfu.20010077.m_210490.5~5!2~5~5!3~5~5!3~5~A&f=29102521&from=y1.3-fun-fun-904-10077.90023-210490.3-4
	youkuVideoUrl = "https://v.qq.com/iframe/player.html?vid=s0521l6pfaf&tiny=1&auto=1"

	if err := phantomjs.DefaultProcess.Open(); err != nil {
		log.Fatal(err)
	}
	defer phantomjs.DefaultProcess.Close()
	//创建webpage
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
		UserAgent:                     "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.81 Mobile Safari/537.36",
		Username:                      "dejavuzhou@qq.com",
		Password:                      "ZHou1987",
		XSSAuditingEnabled:            true,
		WebSecurityEnabled:            true,
		ResourceTimeout:               20,
	}

	if err := page.SetSettings(webPageSettings); err != nil {
		log.Fatal(err)
	}

	//http://www.dy2018.com/html/gndy/oumei/
	//http://www.dy2018.com/html/tv/oumeitv/index.html
	// if err := page.IncludeJS("http://cdn.bootcss.com/jquery/3.2.1/jquery.slim.min.js"); err != nil {
	// 	log.Fatal(err)
	// }
	if err := page.Open(youkuVideoUrl); err != nil {
		log.Fatal(err)
	}
	// 不支持 promise phantomjs
	// Read first link.
	if err := page.Render("hackernews.png", "png", 100); err != nil {
		log.Fatal(err)
	}
	page.Evaluate(`function() {
		document.querySelector('span.tvp_button_play').click();
	}`)
	nodes, err := page.Evaluate(`function() {
		document.querySelector('span.tvp_button_play').click();
		var videoURL = document.getElementsByTagName('video')[0].src;
		var html = markup = document.querySelector('div.tvp_video').innerHTML;
  		return {video:videoURL,code:html};
	}`)
	if err != nil {
		log.Fatal(err)
	}

	//失败 这里面有加密的地方
	fmt.Println(nodes)
	var resss VideoJson
	json.Unmarshal(nodes.([]byte), &resss)
	return "data"
	// var array = new Array();
	// var array = $('#header > div > div.bd2 > div.bd3 > div.bd3r > div.co_area2 > div.co_content8 > ul > table > tbody > tr:nth-child(2) > td:nth-child(2) > b > a:nth-child(2)').each(function(index,value){var href = $(value).attr('href');var name=$(value).attr('title');array.push({href:href,name:name})});
}
