package models

import (
	"changeGo/common"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int    `orm:"pk;auto;unique;column(user_id)" json:"id"`
	Username string `orm:"column(username);size(500)" json:"username"`
	Password string `orm:"column(password);size(500)" json:"password"`
	Role     string `orm:"column(role);type(string);default(dev)	" description:"角色(dev: 开发， dba: 数据库管理员)" json:"Role"`
}

func NewUser() *User {
	return &User{}
}

func (a *User) TableName() string {
	return "user"
}

type UserPost struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) Add(name string, password string) (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(&User{Username: name, Password: common.Db_passwd_encode(password)})
	if err != nil {
		return err
	}
	return
}

func (u *User) Verification(name string, password string) (status bool, err error) {
	o := orm.NewOrm()
	queryObj := o.QueryTable(u.TableName()).Filter("username", name)
	count, _ := queryObj.Count()
	if count == 0 {
		status = false
		err = fmt.Errorf("缺少用户")
		return
	}
	var user_obj User
	queryObj.Limit(1).One(&user_obj, "password")
	fmt.Println(user_obj.Password)
	if password == common.Db_passwd_decode(user_obj.Password) {
		status = true
		return
	} else {
		status = false
		err = fmt.Errorf("密码错误")
		return
	}
}

func (u *User) ChangeRole(id int, role string) (err error) {
	if role != "dev" || role != "dba" {
		err = fmt.Errorf("角色信息错误")
		return
	} else {
		o := orm.NewOrm()
		queryObj := o.QueryTable(u.TableName()).Filter("user_id", id)
		count, _ := queryObj.Count()
		if count == 0 {
			err = fmt.Errorf("没有该用户")
			return
		}
		queryObj.Update(orm.Params{
			"Role": role,
		})
		return
	}

}

func (u *User) AddName(name string) (err error) {
	o := orm.NewOrm()
	queryObj := o.QueryTable(u.TableName()).Filter("username", name)
	count, _ := queryObj.Count()
	if count == 0 {
		o.Insert(&User{Username: name})
		return
	}
	return
}

func (u *User) GetAllUser() (result []string, err error) {
	o := orm.NewOrm()
	var users []User
	_, err = o.QueryTable(u.TableName()).All(&users, "username")
	for _, obj := range users {
		result = append(result, obj.Username)
	}

	return
}
