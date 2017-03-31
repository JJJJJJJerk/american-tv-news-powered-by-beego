package main

import (
	"fmt"                 // package for joining parameters from config/app.conf to mysql db connection configuration string
	_ "my_go_web/routers" //init reuters
	_ "my_go_web/tools"   //init(register) html-template function ...

	_ "github.com/go-sql-driver/mysql" // import your used driver

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	//config mysql database with config/app.config file
	user := beego.AppConfig.String("mysqluser")
	passwd := beego.AppConfig.String("mysqlpass")
	host := beego.AppConfig.String("mysqlurls")
	port, err := beego.AppConfig.Int("mysqlport")
	dbname := beego.AppConfig.String("mysqldb")
	if nil != err {
		port = 3306
	}
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, passwd, host, port, dbname))
}

func main() {
	beego.Run()
}
