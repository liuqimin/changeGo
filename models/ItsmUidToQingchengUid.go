package models

import (
	"github.com/astaxie/beego/orm"
)

type ItsmUidToQingchengUid struct {
	Id      int    `orm:"pk;auto;unique;column(id)"`
	Quid    string `orm:"column(quid);size(500);comment(青橙用户账号)"`
	Itsmuid string `orm:"column(itsmuid);size(500);comment(itsm用户账号)"`
}

func NewItsmUidToQingchengUid() *ItsmUidToQingchengUid {
	return &ItsmUidToQingchengUid{}
}

func (a *ItsmUidToQingchengUid) TableName() string {
	return "record_web_itsmuidtoqingchenguid"
}

func init() {
	orm.RegisterModel(new(ItsmUidToQingchengUid))
}
