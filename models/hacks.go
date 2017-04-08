package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Hacks struct {
	Id           int       `orm:"column(id);auto"`
	Name         string    `orm:"column(name);size(255)"`
	Url          string    `orm:"column(url);size(255)"`
	Cookie       string    `orm:"column(cookie);size(255)"`
	Type         string    `orm:"column(type);size(255)"`
	Query        string    `orm:"column(query);size(255)"`
	Para         string    `orm:"column(para);size(255)"`
	IsSuccess    int8      `orm:"column(is_success)"`
	CommandLines string    `orm:"column(command_lines)"`
	CreatedAt    time.Time `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt    time.Time `orm:"column(updated_at);type(timestamp);null"`
}

func (t *Hacks) TableName() string {
	return "hacks"
}

func init() {
	orm.RegisterModel(new(Hacks))
}

// AddHacks insert a new Hacks into database and returns
// last inserted Id on success.
func AddHacks(m *Hacks) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetHacksById retrieves Hacks by Id. Returns error if
// Id doesn't exist
func GetHacksById(id int) (v *Hacks, err error) {
	o := orm.NewOrm()
	v = &Hacks{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllHacks retrieves all Hacks matches certain condition. Returns empty list if
// no records exist
func GetAllHacks(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Hacks))
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

	var l []Hacks
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

// UpdateHacks updates Hacks by Id and returns error if
// the record to be updated doesn't exist
func UpdateHacksById(m *Hacks) (err error) {
	o := orm.NewOrm()
	v := Hacks{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteHacks deletes Hacks by Id and returns error if
// the record to be deleted doesn't exist
func DeleteHacks(id int) (err error) {
	o := orm.NewOrm()
	v := Hacks{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Hacks{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
