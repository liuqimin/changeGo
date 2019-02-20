package api

import (
	"changeGo/api/getStruct"
	"changeGo/common"
	commonlogger "changeGo/common/logger"
	"changeGo/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/smartgo/stgcommon/logger"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"
)

const (
	formatTime     = "15:04:05"
	formatDate     = "2006-01-02"
	formatDateTime = "2006-01-02 15:04:05"
)

var LoggerInit = commonlogger.NewApplloLog()

var KeyWordPublish = [...]string{"应用版本投产", "信用卡版本", "零售版本", "公司版本"}

type responseStruct struct {
	Code    string                  `json:"code"`
	Rows    []getStruct.ItsmGetData `json:"rows"`
	Message string                  `json:"message"`
}

func ChangITSMSyncThreeDay() error {
	now := time.Now()
	m, _ := time.ParseDuration("-72h")
	m1 := now.Add(m)
	payload := fmt.Sprintf("?mintime=%v&maxtime=%v", now.Format(formatDate), m1.Format(formatDate))
	fmt.Println(payload)
	//logger.debug("定时任务{}".format(payload))
	err := ChangeITSMSync(payload)
	return err
}

func ChangITSMSyncTwoMonth() error {
	now := time.Now()
	m, _ := time.ParseDuration("-1440h")
	m1 := now.Add(m)
	payload := fmt.Sprintf("?mintime=%v&maxtime=%v", now.Format(formatDate), m1.Format(formatDate))
	fmt.Println(payload)
	//logger.debug("定时任务{}".format(payload))
	err := ChangeITSMSync(payload)
	return err
}

func ChangITSMSyncOneYear() error {
	/*同步一年方法*/
	now := time.Now()
	m, _ := time.ParseDuration("-8640h")
	m1 := now.Add(m)
	payload := fmt.Sprintf("?mintime=%v&maxtime=%v", now.Format(formatDate), m1.Format(formatDate))
	fmt.Println(payload)
	//logger.debug("定时任务{}".format(payload))
	err := ChangeITSMSync(payload)
	return err
}

func ChangeITSMSync(payload string) error {
	LoggerInit.Info(fmt.Sprintf("%v start", payload))
	postUrl := fmt.Sprintf("%v%v%v", beego.AppConfig.String("itsm::ITSM_URL"),
		beego.AppConfig.String("itsm::ITSM_GET_CHANGE"), payload)
	httpdo := common.HttpDo()
	req, err := http.NewRequest("GET", postUrl, nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpdo.Do(req)
	//responseBase , err := http.Get(postUrl)
	if err != nil {
		LoggerInit.Error(fmt.Sprintf("%v", err))
		return err
	}
	LoggerInit.Info(fmt.Sprintf("%v", err))
	fmt.Println(resp)
	if err != nil {
		LoggerInit.Error(fmt.Sprintf("%v", err))
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			LoggerInit.Error(fmt.Sprintf("%v", err))
			return err
		}
		LoggerInit.Info(fmt.Sprintf("get info %v", string(body)))
		var itsmData responseStruct
		err = json.Unmarshal(body, &itsmData)
		if err != nil {
			LoggerInit.Error(fmt.Sprintf("json解析失败 %v", err))
			return err
		}
		handleData(itsmData)

	}
	LoggerInit.Error(fmt.Sprintf("%v", err))
	return err
}

func findKeyWord(s string) bool {
	for i := 0; i < len(KeyWordPublish); i++ {
		if KeyWordPublish[i] == s {
			return true
		}
	}
	return false

}

func InputPublishData(data getStruct.ItsmGetData) {
	var a = models.NewItsmPublishInfo()
	fields := reflect.ValueOf(a).Elem().NumField()
	for i := 0; i < fields; i++ {
		vv := reflect.TypeOf(models.NewItsmPublishInfo()).Elem().Field(i).Name
		//value := reflect.ValueOf(postData).Elem().Field(i) //get name
		_, tt := reflect.ValueOf(data).Elem().Type().FieldByName(vv) //panic: reflect: call of reflect.Value.Elem on string Value
		logger.Info(tt)
		if tt {
			setNameAddr := reflect.ValueOf(data).Elem().FieldByName(vv)
			reflect.ValueOf(a).Elem().FieldByName(vv).Set(setNameAddr)

		} else {
			continue
		}
	}
	fmt.Println(a)
	_, err := a.Insert()
	if err != nil {
		LoggerInit.Error(fmt.Sprintf("插入失败，失败原因%v", err))
	}
}

func InputChangeData(data getStruct.ItsmGetData) {
	var a = models.NewItsmChangeInfo()
	fields := reflect.ValueOf(a).Elem().NumField()
	for i := 0; i < fields; i++ {
		vv := reflect.TypeOf(models.NewItsmPublishInfo()).Elem().Field(i).Name
		_, tt := reflect.ValueOf(data).Elem().Type().FieldByName(vv) //panic: reflect: call of reflect.Value.Elem on string Value
		if tt {
			setNameAddr := reflect.ValueOf(data).Elem().FieldByName(vv)
			reflect.ValueOf(a).Elem().FieldByName(vv).Set(setNameAddr)

		} else {
			continue
		}
	}
	fmt.Println(a)
	_, err := a.Insert()
	if err != nil {
		LoggerInit.Error(fmt.Sprintf("插入失败，失败原因%v", err))
	}
}

func handleData(itsmData responseStruct) error {
	if len(itsmData.Rows) > 0 {
		for i, v := range itsmData.Rows {
			fmt.Println(i)
			selectObj := strings.Split(v.Changectitype, "-")
			if findKeyWord(selectObj[2]) {
				InputPublishData(v)
			} else {
				InputChangeData(v)
			}
		}
	} else {
		LoggerInit.Error(fmt.Sprintf("本次未获得数据"))
		err := fmt.Errorf("本次未获得数据")
		return err
	}
	var err error
	return err
}
