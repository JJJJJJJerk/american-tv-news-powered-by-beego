package models

type Sessions struct {
	Id_RENAME    string `orm:"column(id);size(255)"`
	UserId       int    `orm:"column(user_id);null"`
	IpAddress    string `orm:"column(ip_address);size(45);null"`
	UserAgent    string `orm:"column(user_agent);null"`
	Payload      string `orm:"column(payload)"`
	LastActivity int    `orm:"column(last_activity)"`
}
