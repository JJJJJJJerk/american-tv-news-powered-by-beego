package models
import (
	"fmt"
	"time"
	"github.com/astaxie/beego/orm"
)

type Articles struct {
	Id          int       `orm:"column(id);auto"`
	Title       string    `orm:"column(title);size(255)"`
	Body        string    `orm:"column(body)"`
	UrlVideo    string    `orm:"column(url_video);size(255)"`
	UrlProvider string    `orm:"column(url_provider);size(255)"`
	UrlFlash    string    `orm:"column(url_flash);size(255)"`
	HtmlCode    string    `orm:"column(html_code)"`
	IsShow      int8      `orm:"column(is_show)"`
	CreatedAt   time.Time `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt   time.Time `orm:"column(updated_at);type(timestamp);null"`
	KeyWord     string    `orm:"column(key_word)"`
	MobileCode  string    `orm:"column(mobile_code)"`
}

//默认的表名规则，使用驼峰转蛇形：

// AuthUser -> auth_user
// Auth_User -> auth__user
// DB_AuthUser -> d_b__auth_user

// func (t *Articles) TableName() string {
// 	return "articles"
// }
func init() {
	orm.RegisterModel(new(Articles))
}

// AddArticles insert a new Articles into database and returns
// last inserted Id on success.
func AddArticles(m *Articles) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetArticlesById retrieves Articles by Id. Returns error if
// Id doesn't exist
func GetArticlesById(id int) (v *Articles, err error) {
	o := orm.NewOrm()
	v = &Articles{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllArticles retrieves all Articles matches certain condition. Returns empty list if
// no records exist
func GetAllArticles() (errs error, artcs []Articles) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Articles))
	var articles []Articles
	_, err := qs.Filter("id__lt", 30).OrderBy("CreatedAt").All(&articles, "Id", "Title")
	return err, articles
}

// UpdateArticles updates Articles by Id and returns error if
// the record to be updated doesn't exist
func UpdateArticlesById(m *Articles) (err error) {
	o := orm.NewOrm()
	v := Articles{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteArticles deletes Articles by Id and returns error if
// the record to be deleted doesn't exist
func DeleteArticles(id int) (err error) {
	o := orm.NewOrm()
	v := Articles{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Articles{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
