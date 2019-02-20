package controllers

import (
	//"changeGo/api"
	"changeGo/common"
	"changeGo/jsontime"
	"changeGo/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type AplloVersioninfoController struct {
	beego.Controller
}

func (a *AplloVersioninfoController) GetData() {
	//var getData api.ApolloGet
	//models.NewApolloInfo()
	fmt.Println("getData")
	common.Consumer(WriteAplloData)
}

func (a *AplloVersioninfoController) PushData() {
	//var getData api.ApolloGet
	//models.NewApolloInfo()
	//fmt.Println(getData)
	fmt.Println("pushData")
	//common.Publish()
}

type User struct {
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	Age       int    `json:"age"`
	AvatarUrl string `json:"avatar_url"`
}

func (a *AplloVersioninfoController) Test() {
	//ReleaseTime ,_ := time.Parse("2006-01-02 15:04:05", "2018-10-12 12:22:00:000000")
	var ReleaseTimeString jsontime.JsonTime
	ReleaseTimeString.UnmarshalJSON([]byte("2018-10-12 12:22:00"))

	fmt.Println(ReleaseTimeString)
	//ReleaseTime ,_ = jsontime.Todatetime("2018-10-12 12:22:00:000000")
	//var mapResult map[string]interface{}
	mapResult := make(map[string]interface{})
	mapResult["releaseOperation"] = "test"
	mapResult["appID"] = "te1a"
	mapResult["env"] = "hllo"
	mapResult["clusterName"] = "clusterName"
	mapResult["namespaceName"] = "namespaceName"
	mapResult["operator"] = "abas"
	mapResult["releaseTitle"] = "afafas"
	mapResult["releaseComment"] = "basfafs"
	mapResult["releaseTime"] = "2018-12-12 12:00:03"
	poshData, err := json.Marshal(mapResult)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
		return
	}
	fmt.Println(poshData)
	fmt.Printf("json:%v\n", string(poshData))
	//fmt.Println([]byte("hello world"))
	//common.Publish([]byte("hello world"))
	common.Publish(poshData)

}

func WriteAplloData(b []byte) {
	var data models.ApolloGet
	err := json.Unmarshal(b, &data)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(12345)
	fmt.Println(data)
	fmt.Println(5678)
	fmt.Println(data.ReleaseTitle)
	var ApolloInfoModel models.ApolloInfo
	//ApolloInfoModel.Env = "test"
	ApolloInfoModel.Add(&data)

}
