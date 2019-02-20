package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type ITSMPowerCheck struct {
	Id           int              `orm:"pk;auto;unique;column(id)"`
	User         string           `orm:"column(user);size(500);comment(用户名)"`
	Power        int              `orm:"column(changetitle);type(tinyint);comment(权限)"`
	Changenumber *ItsmPublishInfo `orm:"column(changenumber_id);rel(fk)"`
}

func NewITSMPowerCheck() *ITSMPowerCheck {
	return &ITSMPowerCheck{}
}

func (a *ITSMPowerCheck) TableName() string {
	return "record_web_itsmpowercheck"
}

func init() {
	orm.RegisterModel(new(ITSMPowerCheck))
}

func (a *ITSMPowerCheck) Insert(data ItsmPublishInfo, id int) error {
	o := orm.NewOrm()
	var ItsmPublishInfoObj ItsmPublishInfo
	o.QueryTable(data.TableName()).Filter("Changenumber", data.Changenumber).One(&ItsmPublishInfoObj)
	query := o.QueryTable(a.TableName()).Filter("Changenumber__id", id)
	if query.Exist() {
		Count, _ := query.Count()
		if Count == 2 {
			//未写
		} else {
			//未写
		}
	} else {
		if len(data.Changeforce) == 0 {
			objChangeForce := new(ITSMPowerCheck)
			objChangeForce.Changenumber = &ItsmPublishInfoObj
			objChangeForce.Power = 3
			objChangeForce.User = data.Changeforce
			o.Insert(objChangeForce)
		} else if len(data.Changechaeck) == 0 {
			objChangechaeck := new(ITSMPowerCheck)
			objChangechaeck.Changenumber = &ItsmPublishInfoObj
			objChangechaeck.Power = 4
			objChangechaeck.User = data.Changeforce
			o.Insert(objChangechaeck)
		} else {

		}
	}
	return fmt.Errorf("over")
}
