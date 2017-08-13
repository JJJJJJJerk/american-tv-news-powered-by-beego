package controllers

import (
	"fmt"
	"my_go_web/models"
	"os/exec"

	"github.com/astaxie/beego"
)

type VideoParseController struct {
	beego.Controller //集成beego controller
}

func (c *VideoParseController) Index() {
	cacheKey := c.GetString("video")
	//cacheKey = fmt.Sprint("'", cacheKey, "'")
	var content []byte
	if x, found := models.CacheManager.Get(cacheKey); found {
		foo := x.(string)
		content = []byte(foo)
	} else {
		cmdString := "you-get"

		cmd := exec.Command(cmdString, "-u", cacheKey)
		output, err := cmd.CombinedOutput()
		fmt.Println(string(output))

		if err != nil {
			fmt.Println(err)

			fmt.Println(string(output))
		} else {
			content = output
			models.CacheManager.Set(cacheKey, string(content), models.C_EXPIRE_TIME_HOUR_01)
			fmt.Println(string(output))
		}

	}

	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	c.Ctx.Output.Body(content)
	return
}
