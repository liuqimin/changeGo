package models

import (
	//"changeGo/api"
	"changeGo/jsontime"
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
)

type ApolloGet struct {
	ReleaseOperation string            `json:"releaseOperation"`
	AppID            string            `json:"appID"`
	Env              string            `json:"env"`
	ClusterName      string            `json:"clusterName"`
	NamespaceName    string            `json:"namespaceName"`
	Operator         string            `json:"operator"`
	ReleaseTitle     string            `json:"releaseTitle"`
	ReleaseComment   string            `json:"releaseComment"`
	ReleaseTime      jsontime.JsonTime `json:"releaseTime"`
	//ReleaseTime time.Time `json:"releaseTime"`
}

var postdata ApolloGet

type ApolloInfo struct {
	Id               int               `orm:"pk;auto;unique;column(id)" json:"id"`
	ReleaseOperation string            `orm:"column(releaseOperation);size(500)" json:"releaseOperation"`
	AppID            string            `orm:"column(appID);size(500)" json:"appID"`
	Env              string            `orm:"column(env);size(500)" json:"env"`
	ClusterName      string            `orm:"column(clusterName);size(500)" json:"clusterName"`
	NamespaceName    string            `orm:"column(namespaceName);size(500)" json:"namespaceName"`
	Operator         string            `orm:"column(operator);size(500)" json:"operator"`
	ReleaseTitle     string            `orm:"column(releaseTitle);size(500)" json:"releaseTitle"`
	ReleaseComment   string            `orm:"column(releaseComment);size(500)" json:"releaseComment"`
	ReleaseTime      jsontime.JsonTime `orm:"column(releaseTime);size(500);type(datetime)" json:"releaseTime"`
}

func NewApolloInfo() *ApolloInfo {
	return &ApolloInfo{}
}

func (a *ApolloInfo) TableName() string {
	return "record_web_configurationcenterinfo"
}

func (a *ApolloInfo) Add(postData *ApolloGet) {
	o := orm.NewOrm()
	fields := reflect.ValueOf(a).Elem().NumField()
	for i := 0; i < fields; i++ {
		vv := reflect.TypeOf(a).Elem().Field(i).Name
		//value := reflect.ValueOf(postData).Elem().Field(i) //get name
		_, tt := reflect.ValueOf(postData).Elem().Type().FieldByName(vv) //panic: reflect: call of reflect.Value.Elem on string Value
		fmt.Println(tt)
		if tt {
			setNameAddr := reflect.ValueOf(postData).Elem().FieldByName(vv)
			reflect.ValueOf(a).Elem().FieldByName(vv).Set(setNameAddr)

		} else {
			continue
		}

	}
	fmt.Println(a)
	id, err := o.Insert(a)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(id)
	}

}
