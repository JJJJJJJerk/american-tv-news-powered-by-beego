package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Weibos struct {
	Id        int       `orm:"column(id);auto"`
	WbId      uint64    `orm:"column(wb_id)"`
	Body      string    `orm:"column(body)"`
	Images    string    `orm:"column(images)"`
	Type      string    `orm:"column(type)"`
	IsShown   uint8     `orm:"column(is_shown)"`
	Weibor    string    `orm:"column(weibor);size(255)"`
	Avatar    string    `orm:"column(avatar);size(255)"`
	MediaUrl  string    `orm:"column(media_url);size(255)"`
	LongUrl   string    `orm:"column(long_url);size(255);null"`
	MediaType string    `orm:"column(media_type);null"`
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp);null"`
}

func (t *Weibos) TableName() string {
	return "weibos"
}

func init() {
	orm.RegisterModel(new(Weibos))
}

// AddWeibos insert a new Weibos into database and returns
// last inserted Id on success.
func AddWeibos(m *Weibos) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetWeibosById retrieves Weibos by Id. Returns error if
// Id doesn't exist
func GetWeibosById(id int) (v *Weibos, err error) {
	o := orm.NewOrm()
	v = &Weibos{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllWeibos retrieves all Weibos matches certain condition. Returns empty list if
// no records exist
func GetAllWeibos(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Weibos))
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

	var l []Weibos
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

// UpdateWeibos updates Weibos by Id and returns error if
// the record to be updated doesn't exist
func UpdateWeibosById(m *Weibos) (err error) {
	o := orm.NewOrm()
	v := Weibos{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteWeibos deletes Weibos by Id and returns error if
// the record to be deleted doesn't exist
func DeleteWeibos(id int) (err error) {
	o := orm.NewOrm()
	v := Weibos{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Weibos{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
