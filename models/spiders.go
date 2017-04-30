package models

import (
	"github.com/astaxie/beego/orm"
)

type Spiders struct {
	Id          int    `orm:"column(id);auto"`
	Name        string `orm:"column(name);size(255)"`
	Description string `orm:"column(description);size(255)"`
	Host        string `orm:"column(host);size(255)"`
	Uri         string `orm:"column(uri);size(255)"`
	Headers     string `orm:"column(headers)"`
	Cookies     string `orm:"column(cookies)"`
	UserAgent   string `orm:"column(user_agent);size(255)"`
	UserName    string `orm:"column(user_name);size(255)"`
	Password    string `orm:"column(password);size(255)"`
	Status      uint8  `orm:"column(status)"`
	AuthUrl     string `orm:"column(auth_url);size(255)"`
}

func (t *Spiders) TableName() string {
	return "spiders"
}

func init() {
	orm.RegisterModel(new(Spiders))
}

func AddSpiders(m *Spiders) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetSpidersByName(name string) (v *Spiders, err error) {
	o := orm.NewOrm()
	v = &Spiders{Name: name}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
