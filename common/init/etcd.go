package initcommon

import (
	"context"
	"fmt"
	"github.com/astaxie/beego"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func Etcdinit() (cli *clientv3.Client, err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{beego.AppConfig.String("emdb::url")},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer cli.Close()
	return

}

func Etcdpush(cli *clientv3.Client, key string, value string) (status bool, err error) {
	_, err = cli.Put(context.Background(), key, value)
	if err != nil {
		fmt.Println(err)
		return
	}
	status = true
	return
}

func EtcdGet(cli *clientv3.Client, key string) {
	rsp, _ := cli.Get(context.Background(), key)
	fmt.Println(cli.KV.Txn(context.Background()))
	fmt.Println(rsp)
}
func EtcdWatch(cli *clientv3.Client, watckKey string) {
	for {
		resultCh := cli.Watch(context.Background(), watckKey, clientv3.WithPrefix())
		fmt.Printf("wacth return, resultCh:%v\n", resultCh)
		for v := range resultCh {
			fmt.Printf("wacth return, v:%v\n", v)
			if v.Err() != nil {
				fmt.Printf("watch failed, err:%v\n", v.Err())
				continue
			}

			for _, e := range v.Events {
				fmt.Printf("event_type:%v key:%v val:%v\n", e.Type, e.Kv.Key, string(e.Kv.Value))
			}
		}
	}
}
