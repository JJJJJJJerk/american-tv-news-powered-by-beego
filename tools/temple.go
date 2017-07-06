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
	beego.AddFuncMap("cdnSrc", cdnSrc) //拼接cdn资源

	beego.AddFuncMap("cdnImageSrc", cdnImageSrc) //拼接cdn资源

}

func cdnImageSrc(image *models.Image, param string) (out string) {
	out = fmt.Sprintf("%s%s%s", imageCdnHost, image.Key, param)
	return
}
func cdnSrc(in string) (out string) {
	out = fmt.Sprintf("%s%s", cdn, in)
	return
}
