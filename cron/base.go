package cron

import (
	commonlogger "changeGo/common/logger"
	"context"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"go.etcd.io/etcd/clientv3"
	"reflect"
	"strings"
	"time"
)

var LoggerInit = commonlogger.NewApplloLog()

func CronInit() {

	defer func() {
		fmt.Println("defer caller")
		if err := recover(); err != nil {
			fmt.Println(err)
			LoggerInit.Error(fmt.Sprintf("%s", err))
			fmt.Println("recover success.")
		}
	}()
	//toolbox.AddTask("hllo", tk)
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{beego.AppConfig.String("emdb::url")},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {

		panic(err)
		//return
	}
	resp, err := cli.Get(context.TODO(), "/monitor", clientv3.WithPrefix())
	if err != nil {
		fmt.Println(12345)

	}
	fmt.Println(len(resp.Kvs))
	fmt.Println(resp.Kvs)
	toolbox.StartTask()
	if len(resp.Kvs) > 0 {
		for _, d := range resp.Kvs {
			//判断下d.value是否 '* * * * * *'
			tk := toolbox.NewTask(string(d.Key), string(d.Value), func() error { err := selectFunc(string(d.Key)); return err })
			toolbox.AddTask(string(d.Key), tk)
		}
	}
}

func selectFunc(f string) (err error) {
	fmt.Println(f)
	k := strings.Split(f, "/")
	RefCron_obj := &RefCron{}
	RefCron_obj.ItsmCron()
	fmt.Println(k[1])
	_, tt := reflect.ValueOf(RefCron_obj).Type().MethodByName(k[2])

	if tt {
		reflect.ValueOf(RefCron_obj).MethodByName(k[2]).Call(nil)
	} else {
		err = fmt.Errorf("输入方法没有相应函数，用户输入内容是%s", k[2])
		LoggerInit.Error(fmt.Sprintf("%s", err))
	}
	return
}

func CronWatch() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{beego.AppConfig.String("emdb::url")},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(123455555)
		panic(err)
		//return
	}
	defer cli.Close()
	for {
		rch := cli.Watch(context.Background(), "/monitor", clientv3.WithPrefix())
		fmt.Println("%v fch is ", &rch)
		for wresp := range rch {
			for _, ev := range wresp.Events {
				fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				fmt.Println(12345)
				Type := fmt.Sprintf("%s", ev.Type)
				SwitchWatchCron(Type, string(ev.Kv.Key), string(ev.Kv.Value))
			}
		}

		time.Sleep(5 * time.Second)
	}
}

func SwitchWatchCron(t, k, v string) {
	//watch任务，当出现变动时候。触发任务
	fmt.Println(t, k, v)

	switch t {
	case "PUT":
		fmt.Println(11)
		ChangeCron(k, v)
	case "DELETE":
		fmt.Println(2222)
		DeleteCron(k)
	}

}

func DeleteCron(name string) {
	fmt.Println(name)
	toolbox.DeleteTask(name)
	fmt.Println(toolbox.AdminTaskList)
}

func ChangeCron(name string, value string) {
	fmt.Println(name)
	fmt.Println(toolbox.AdminTaskList)
	_, ok := toolbox.AdminTaskList[name]
	if ok {
		//在列表中，修改字符串信息
		UpdateCron(name, value)
	} else {
		//新增列表
		AddCron(name, value)
	}

	//toolbox.DeleteTask(name)
	fmt.Println(toolbox.AdminTaskList)
}

func AddCron(k, v string) {
	tk := toolbox.NewTask(k, v, func() error { err := selectFunc(k); return err })
	toolbox.AddTask(k, tk)
}

func UpdateCron(k, v string) {
	toolbox.DeleteTask(k)
	tk := toolbox.NewTask(k, v, func() error { err := selectFunc(k); return err })
	toolbox.AddTask(k, tk)
	fmt.Println(toolbox.AdminTaskList)
}
