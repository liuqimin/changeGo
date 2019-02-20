package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Author struct {
	Id       int       `orm:"pk;auto;unique;column(author_id)" json:"author_id"`
	Name     string    `orm:"column(name);size(500)" json:"name"`
	Country  string    `orm:"column(country);size(50)" json:"country"`
	Birthday time.Time `orm:"column(birthday);type(birthday)" json:"birthday"`
}

func NewAuthor() *Author {
	return &Author{}
}

func (a *Author) TableName() string {
	return "author"
}
func (a *Author) Insert() (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(a)
	if err != nil {
		return err
	}
	return nil
}

func (a *Author) Getall() (result []*Author, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(a.TableName()).All(&result)
	return
}

func (a *Author) Del(authorId int) (status string, err error) {
	o := orm.NewOrm()
	fmt.Println(o)
	fmt.Println(authorId)
	queryObj := o.QueryTable(a.TableName()).Filter("author_id", authorId)
	count, _ := queryObj.Count()
	if count == 0 {
		err = fmt.Errorf("没有数据")
		return
	}
	_, err = queryObj.Delete()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(err)
	return
}
