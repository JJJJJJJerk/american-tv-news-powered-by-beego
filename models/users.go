package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id             int       `orm:"column(id);auto"`
	Name           string    `orm:"column(name);size(255)"`
	Email          string    `orm:"column(email);size(255)"`
	Password       string    `orm:"column(password);size(60)"`
	HonorPoint     int32     `orm:"column(honor_point)"`
	NickName       string    `orm:"column(nick_name);size(20)"`
	Slogan         string    `orm:"column(slogan);size(40)"`
	AvatarImage    string    `orm:"column(avatar_image);size(255)"`
	RememberToken  string    `orm:"column(remember_token);size(100);null"`
	CreatedAt      time.Time `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt      time.Time `orm:"column(updated_at);type(timestamp);null"`
	GithubId       int       `orm:"column(github_id)"`
	GithubToken    string    `orm:"column(github_token);size(255)"`
	GithubName     string    `orm:"column(github_name);size(255)"`
	GithubNickname string    `orm:"column(github_nickname);size(255)"`
	GithubEmail    string    `orm:"column(github_email);size(255)"`
	GithubAvatar   string    `orm:"column(github_avatar);size(255)"`
	GithubJson     string    `orm:"column(github_json)"`
	WeiboId        int       `orm:"column(weibo_id)"`
	WeiboToken     string    `orm:"column(weibo_token);size(255)"`
	WeiboName      string    `orm:"column(weibo_name);size(255)"`
	WeiboNickname  string    `orm:"column(weibo_nickname);size(255)"`
	WeiboEmail     string    `orm:"column(weibo_email);size(255)"`
	WeiboAvatar    string    `orm:"column(weibo_avatar);size(255)"`
	WeiboJson      string    `orm:"column(weibo_json)"`
	QqId           int       `orm:"column(qq_id)"`
	QqToken        string    `orm:"column(qq_token);size(255)"`
	QqName         string    `orm:"column(qq_name);size(255)"`
	QqNickname     string    `orm:"column(qq_nickname);size(255)"`
	QqEmail        string    `orm:"column(qq_email);size(255)"`
	QqAvatar       string    `orm:"column(qq_avatar);size(255)"`
	QqJson         string    `orm:"column(qq_json)"`
	WeixinId       int       `orm:"column(weixin_id)"`
	WeixinToken    string    `orm:"column(weixin_token);size(255)"`
	WeixinName     string    `orm:"column(weixin_name);size(255)"`
	WeixinNickname string    `orm:"column(weixin_nickname);size(255)"`
	WeixinEmail    string    `orm:"column(weixin_email);size(255)"`
	WeixinAvatar   string    `orm:"column(weixin_avatar);size(255)"`
	WeixinJson     string    `orm:"column(weixin_json)"`
}

func (t *Users) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(new(Users))
}

// AddUsers insert a new Users into database and returns
// last inserted Id on success.
func AddUsers(m *Users) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUsersById retrieves Users by Id. Returns error if
// Id doesn't exist
func GetUsersById(id int) (v *Users, err error) {
	o := orm.NewOrm()
	v = &Users{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUsers retrieves all Users matches certain condition. Returns empty list if
// no records exist
func GetAllUsers(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Users))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Users
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

//根据email用户信息
func GetUsersByEmail(email string) (v *Users, err error) {
	o := orm.NewOrm()
	v = &Users{Email: email}
	if err = o.Read(v,"Email"); err == nil {
		return v, nil
	}
	return nil, err
}

// UpdateUsers updates Users by Id and returns error if
// the record to be updated doesn't exist
func UpdateUsersById(m *Users) (err error) {
	o := orm.NewOrm()
	v := Users{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUsers deletes Users by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUsers(id int) (err error) {
	o := orm.NewOrm()
	v := Users{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Users{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
