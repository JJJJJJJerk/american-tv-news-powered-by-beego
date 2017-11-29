package tasks

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"www.mojotv.cn/models"

	"github.com/astaxie/beego/toolbox"
)

const EZTV_XML_URL = "https://eztv.ag/ezrss.xml"

type Rss struct {
	XMLName xml.Name    `xml:"rss"`
	Channel ChannelNode `xml:"channel"`
}
type ChannelNode struct {
	XMLName  xml.Name         `xml:"channel"`
	Episodes []EztvXmlEpisode `xml:"item"`
}
type EztvXmlEpisode struct {
	XMLName   xml.Name `xml:"item"`
	Title     string   `xml:"title"`
	MagnetURI string   `xml:"magnetURI"`
}

//启动定时任务
func init() {
	//创建定时任务
	taskFetchEztvXml := toolbox.NewTask("fetch-eztv", "43 13 8-20 * * *", fetchEztvXmlThenParse)
	err := taskFetchEztvXml.Run()
	//检测定时任务
	if err != nil {
		log.Fatal(err)
	}
	//添加定时任务
	toolbox.AddTask("fetch-eztv", taskFetchEztvXml)
	//启动定时任务
	toolbox.StartTask()
	//defer toolbox.StopTask()
}

func fetchEztvXmlThenParse() error {
	resp, err := http.Get(EZTV_XML_URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err == nil && resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		v := Rss{}
		err = xml.Unmarshal(body, &v)
		if err != nil {
			return err
		} else {
			for _, vo := range v.Channel.Episodes {
				temp := models.Episode{}
				err := models.Gorm.Where(models.Episode{RawName: vo.Title}).Assign(models.Episode{Name: vo.Title, UrlMagnet: vo.MagnetURI, Provider: "eztv"}).FirstOrCreate(&temp).Error
				fmt.Println(err)
			}
			return nil
		}
	} else {
		return err
	}
}
