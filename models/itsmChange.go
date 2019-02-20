package models

import (
	"changeGo/jsontime"
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
)

type ItsmChangeInfo struct {
	Id              int               `orm:"pk;auto;unique;column(id)" json:"id"`
	Changenumber    string            `orm:"column(changenumber);size(500);comment(变更编号)"`
	Changetitle     string            `orm:"column(changetitle);type(text);size(500);comment(变更编号)"`
	Happentime      jsontime.JsonTime `orm:"column(happentime);type(datetime);size(500);comment(发生时间)"`
	Forcetime       jsontime.JsonTime `orm:"column(forcetime);type(datetime);size(500);comment(期望实施时间)"`
	Changepricrity  string            `orm:"column(changepricrity);size(500);comment(优先级别)"`
	Changedetails   string            `orm:"column(changedetails);type(text);size(500);comment(详情描述)"`
	Changeforce     string            `orm:"column(changeforce);size(500);comment(变更实施人)"`
	Createtime      string            `orm:"column(createtime);size(500);comment(提交时间)"`
	Changerisk      string            `orm:"column(changerisk);size(500);type(text);comment(风险评估)"`
	Backspace       string            `orm:"column(backspace);size(500);type(text);comment(回退方案)"`
	Changestatus    string            `orm:"column(changestatus);type(text);comment(数据状态)"`
	Changeapprove   string            `orm:"column(changeapprove);size(500);type(text);comment(变更审核人)"`
	Changechaeck    string            `orm:"column(changechaeck);size(500);type(text);comment(验证人)"`
	ChangeType      string            `orm:"column(changeType);size(500);type(tinyint);comment(变更类型)"`
	Changuser       string            `orm:"column(changuser);size(500);comment(申请人)"`
	Changemajor     string            `orm:"column(changemajor);size(500);type(tinyint);comment(是否重大变更)"`
	InfluenceSystem string            `orm:"column(influenceSystem);type(text);comment(影响系统)"`
	SupplementOrder string            `orm:"column(supplementOrder);type(tinyint);comment(是否补单)"`
	UpdateDate      string            `orm:"column(updateDate);comment(更新时间)"`
	UpdateBy        string            `orm:"column(updateBy);comment(更新人)"`
	Orderclosetime  string            `orm:"column(orderclosetime);comment(关闭时间)"`
	Environment     string            `orm:"column(environment);comment(环境)"`
	Overtime        string            `orm:"column(overtime);comment(完工时间)"`
	IsChange        string            `orm:"column(isChange);type(tinyint);comment(配置是否变动)"`
	StartTime       string            `orm:"column(startTime);comment(实际开始时间)"`
	EndTime         string            `orm:"column(endTime);comment(实际结束时间)"`
	Exetype         string            `orm:"column(exetype);type(tinyint);comment(单据状态)"`
	Groupname       string            `orm:"column(groupname);comment(团队)"`
}

func NewItsmChangeInfo() *ItsmChangeInfo {
	return &ItsmChangeInfo{}
}

func (a *ItsmChangeInfo) TableName() string {
	return "record_web_itsmchangerecord"
}

func init() {
	orm.RegisterModel(new(ItsmChangeInfo))
}

func (a *ItsmChangeInfo) Insert() (i int64, err error) {

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
