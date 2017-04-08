package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Episodes struct {
	Id          int       `orm:"column(id);auto"`
	RawName     string    `orm:"column(raw_name);size(255)"`
	Provider    string    `orm:"column(provider);size(255)"`
	UrlProvider string    `orm:"column(url_provider);size(255)"`
	Season      int16     `orm:"column(season)"`
	Episode     int16     `orm:"column(episode)"`
	ShowId      uint      `orm:"column(show_id)"`
	Size        string    `orm:"column(size);size(255)"`
	Quality     string    `orm:"column(quality);size(255)"`
	IsAuth      int8      `orm:"column(is_auth)"`
	UrlTvmaze   string    `orm:"column(url_tvmaze);size(255)"`
	UrlDouban   string    `orm:"column(url_douban);size(255)"`
	UrlTorrent  string    `orm:"column(url_torrent);size(255)"`
	UrlMagnet   string    `orm:"column(url_magnet);size(255)"`
	UrlEd2k     string    `orm:"column(url_ed2k);size(255)"`
	UrlBaidupan string    `orm:"column(url_baidupan);size(255)"`
	UrlOther    string    `orm:"column(url_other);size(255)"`
	UrlVideo    string    `orm:"column(url_video);size(255)"`
	LikeCount   uint32    `orm:"column(like_count)"`
	CreatedAt   time.Time `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt   time.Time `orm:"column(updated_at);type(timestamp);null"`
	MovieId     uint      `orm:"column(movie_id)"`
}

func (t *Episodes) TableName() string {
	return "episodes"
}

func init() {
	orm.RegisterModel(new(Episodes))
}

// AddEpisodes insert a new Episodes into database and returns
// last inserted Id on success.
func AddEpisodes(m *Episodes) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEpisodesById retrieves Episodes by Id. Returns error if
// Id doesn't exist
func GetEpisodesById(id int) (v *Episodes, err error) {
	o := orm.NewOrm()
	v = &Episodes{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEpisodes retrieves all Episodes matches certain condition. Returns empty list if
// no records exist
func GetAllEpisodes(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Episodes))
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

	var l []Episodes
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

// UpdateEpisodes updates Episodes by Id and returns error if
// the record to be updated doesn't exist
func UpdateEpisodesById(m *Episodes) (err error) {
	o := orm.NewOrm()
	v := Episodes{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEpisodes deletes Episodes by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEpisodes(id int) (err error) {
	o := orm.NewOrm()
	v := Episodes{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Episodes{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
