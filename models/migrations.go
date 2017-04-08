package models

type Migrations struct {
	Migration string `orm:"column(migration);size(255)"`
	Batch     int    `orm:"column(batch)"`
}
