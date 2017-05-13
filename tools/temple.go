package tools

//防止模板相关的方便函数
import (
	"fmt"
	"my_go_web/models"

	"github.com/astaxie/beego"
)

//注册tmeplate的一些实用函数
//https://beego.me/docs/mvc/view/template.md

//配置cd镜像地址
var cdn string
var imageCdnHost string

func init() {
	cdn = beego.AppConfig.DefaultString("cdnhost", "/static/")
	imageCdnHost = beego.AppConfig.String("imageCdnHost")

	//注册template function
	beego.AddFuncMap("hi", hello)
	beego.AddFuncMap("cdnSrc", cdnSrc) //拼接cdn资源

	beego.AddFuncMap("cdnImageSrc", cdnImageSrc) //拼接cdn资源

	beego.AddFuncMap("articleCoverageSrc", articleCoverageSrc) //获取文章的封面图片

}

func hello(in string) (out string) {
	out = in + "world"
	return
}

func cdnImageSrc(image *models.Image, param string) (out string) {
	out = fmt.Sprintf("%s%s%s", imageCdnHost, image.Key, param)
	return
}
func cdnSrc(in string) (out string) {
	out = fmt.Sprintf("%s%s", cdn, in)
	return
}

func articleCoverageSrc(article *models.Article) (src string) {
	var key string
	if article.Coverage == nil {
		if len(article.Images) == 0 {
			key = "1461329417"
		} else {
			key = article.Images[0].Key
		}
	} else {
		key = article.Coverage.Key
	}
	param := "?imageView2/1/w/120/h/120"
	src = fmt.Sprintf("%s%s%s", imageCdnHost, key, param)
	return
}

func articleExcerpt(article *models.Article) (text string) {
	return ""
}
