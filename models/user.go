package models

import (
	"github.com/beego/beego/v2/adapter/validation"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type User struct {
	Id       int64
	Username string `orm:"size(64);unique" form:"Username" valid:"Required;Username"`
	Email    string `orm:"size(64);unique" form:"Email" valid:"Required;Email"`
	Password string `orm:"size(225)" form:"Password" valid:"Required;MinSize(6)"`
	//Repassword    string    `orm:"-" form:"Repassword" valid:"Required"`
	Lastlogintime time.Time `orm:"type(datetime);null" form:"-"`
	CreatedAt     time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt     time.Time `orm:"auto_now;type(datetime)"`
}

func (u *User) Valid(v *validation.Validation) {
	//if u.Password != u.Repassword {
	//	v.SetError("Repassword", "Does not matched password, repassword")
	//}
}

func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func Users() orm.QuerySeter {
	var table User
	return orm.NewOrm().QueryTable(table).OrderBy("-Id")
}

func init() {
	orm.RegisterModel(new(User))
}
