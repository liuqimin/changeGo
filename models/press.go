package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Press struct {
	Id   int    `orm:"pk;auto;unique;column(press_id)" json:"press_id"`
	Name string `orm:"column(name);size(500)" json:"name"`
}

func NewPress() *Press {
	return &Press{}
}

func (p *Press) TableName() string {
	return "press"
}

func (p *Press) Insert(name string) (err error) {
	o := orm.NewOrm()
	fmt.Println(name)
	_, err = o.Insert(&Press{Name: name})
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

func (p *Press) GetAll() (result []*Press, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(p.TableName()).All(&result)
	if err != nil {
		return
	}
	return
}
