package cron

import (
	"changeGo/api"
	"changeGo/common/monitor"
	"fmt"
)

type RefCron struct {
}

func (a *RefCron) ItsmCron() (err error) {
	api.ChangITSMSyncThreeDay()
	return
}

func (a *RefCron) Monitor() (err error) {
	data := [2]string{monitorcommon.GetItsmStatus(), monitorcommon.GetItsmStatus()}
	var dataString string
	for _, value := range data {
		if len(value) != 0 {
			if len(dataString) == 0 {
				dataString = fmt.Sprintf("%s", value)
			} else {
				dataString = fmt.Sprintf("%s\n%s", dataString, value)
			}
		}
	}
	err = monitorcommon.WriteFile(dataString)
	fmt.Println(dataString)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (a *RefCron) Test() (err error) {
	fmt.Println("test")
	return
}
