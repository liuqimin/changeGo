package models

import (
	"changeGo/jsontime"
	"fmt"
	"github.com/astaxie/beego/orm"
	//	"go.etcd.io/etcd/raft"
	"reflect"
)

var deleteFlag_CH = map[int8]string{
	1: "紧急变更",
	2: "标准变更",
	3: "常规变更",
	9: "无数据",
}
var isDisaster_CH = map[int8]string{
	0: "否",
	1: "是",
	9: "无数据",
}

var eXETYPE_CH = map[int8]string{
	1: "变更成功",
	2: "失败回退",
	3: "完成",
	4: "变更取消",
	9: "无数据",
}

type ItsmPublishInfo struct {
	Id              int               `orm:"pk;auto;unique;column(id)"`
	Changenumber    string            `orm:"column(changenumber);size(500);comment(变更编号);null"`
	Changetitle     string            `orm:"column(changetitle);type(text);size(500);comment(变更编号);null"`
	Happentime      jsontime.JsonTime `orm:"column(happentime);type(datetime);size(500);comment(发生时间);null"`
	Forcetime       jsontime.JsonTime `orm:"column(forcetime);type(datetime);size(500);comment(期望实施时间);null"`
	Changepricrity  string            `orm:"column(changepricrity);size(500);comment(优先级别);null"`
	Changedetails   string            `orm:"column(changedetails);type(text);size(500);comment(详情描述);null"`
	Changeforce     string            `orm:"column(changeforce);size(500);comment(变更实施人);null"`
	Createtime      string            `orm:"column(createtime);size(500);comment(提交时间);null"`
	Changerisk      string            `orm:"column(changerisk);size(500);type(text);comment(风险评估);null"`
	Backspace       string            `orm:"column(backspace);size(500);type(text);comment(回退方案);null"`
	Changestatus    string            `orm:"column(changestatus);type(text);comment(数据状态);null"`
	Changeapprove   string            `orm:"column(changeapprove);size(500);type(text);comment(变更审核人);null"`
	Changechaeck    string            `orm:"column(changechaeck);size(500);type(text);comment(验证人);null"`
	ChangeType      string            `orm:"column(changeType);size(500);type(tinyint);comment(变更类型);null"`
	Changuser       string            `orm:"column(changuser);size(500);comment(申请人);null"`
	Changemajor     string            `orm:"column(changemajor);size(500);type(tinyint);comment(是否重大变更);null"`
	InfluenceSystem string            `orm:"column(influenceSystem);type(text);comment(影响系统);null"`
	SupplementOrder string            `orm:"column(supplementOrder);type(tinyint);comment(是否补单);null"`
	UpdateDate      string            `orm:"column(updateDate);comment(更新时间);null"`
	UpdateBy        string            `orm:"column(updateBy);comment(更新人);null"`
	Orderclosetime  string            `orm:"column(orderclosetime);comment(关闭时间);null"`
	Environment     string            `orm:"column(environment);comment(环境);null"`
	Overtime        string            `orm:"column(overtime);comment(完工时间);null"`
	IsChange        string            `orm:"column(isChange);type(tinyint);comment(配置是否变动);null"`
	StartTime       string            `orm:"column(startTime);comment(实际开始时间);null"`
	EndTime         string            `orm:"column(endTime);comment(实际结束时间);null"`
	Exetype         string            `orm:"column(exetype);type(tinyint);comment(单据状态);null"`
	Groupname       string            `orm:"column(groupname);comment(团队);null"`
}

func NewItsmPublishInfo() *ItsmPublishInfo {
	return &ItsmPublishInfo{}
}

func (a *ItsmPublishInfo) TableName() string {
	return "record_web_itsmrecord"
}

func init() {
	orm.RegisterModel(new(ItsmPublishInfo))
}

func (a *ItsmPublishInfo) Insert() (i int64, err error) {

	o := orm.NewOrm()
	count, _ := o.QueryTable(a.TableName()).Filter("Changenumber", a.Changenumber).Count()
	if count == 1 {
		//更新
		fields := reflect.ValueOf(a).Elem().NumField()
		params := make(orm.Params, fields) //orm.Params为map[string]interface{}
		for i := 0; i < fields; i++ {
			k := reflect.TypeOf(a).Elem().Field(i).Name
			v := reflect.ValueOf(a).Elem().Field(i)
			params[k] = v
		}
		i, err = o.QueryTable(a.TableName()).Filter("Changenumber", a.Changenumber).Update(params)
		if err != nil {
			return
		}
	} else if count == 0 {
		i, err = o.Insert(&a)
		if err != nil {
			return
		}
	} else {
		err = fmt.Errorf("数据有重复，重复单号是%v", a.Changenumber)
		return
	}
	return
	//fields := reflect.ValueOf(a).Elem().NumField()
	//return i,fmt.Errorf("over")
}

func (a *ItsmPublishInfo) FindChangeforce(changenumber string) (user string, status bool, err error) {
	o := orm.NewOrm()
	var data ItsmPublishInfo
	qs := o.QueryTable(a.TableName()).Filter("changenumber", changenumber)
	if qs.Exist() {
		err = qs.Limit(1).One(&data)
		if err != nil {
			status = false

			return
		}
		user = data.Changeforce
		return
	} else {
		user = ""
		status = false
		return
	}
	return
}
