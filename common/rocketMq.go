package common

import (
	"flag"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/didapinchegit/go_rocket_mq"
	"github.com/golang/glog"
	rbt "github.com/sevenNt/rocketmq"
	"time"
)

func Consumer(f func([]byte)) {
	flag.Parse()
	conf := &rocketmq.Config{
		Nameserver:   beego.AppConfig.String("RocketmqNameSrv"),
		InstanceName: "DEFAULT",
	}
	fmt.Println("rocketmq")
	consumer, err := rocketmq.NewDefaultConsumer(beego.AppConfig.String("RocketmqBroker"), conf)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	consumer.Subscribe(beego.AppConfig.String("RocketmqTopic"), "*")
	consumer.RegisterMessageListener(func(msgs []*rocketmq.MessageExt) error {
		for i, msg := range msgs {
			fmt.Println(string(msg.Body))
			glog.Info(i, string(msg.Body))
			go f(msg.Body)
		}
		return nil
	})
	fmt.Println("test")
	consumer.Start()
	time.Sleep(1000 * time.Second)
}

func Publish(data []byte) {
	var producerGroup = "Default"
	var topic = "Apollo-publish-messlog"
	var conf = &rbt.Config{
		Namesrv:      beego.AppConfig.String("RocketmqNameSrv"),
		InstanceName: "DEFAULT",
	}
	fmt.Println(producerGroup)
	fmt.Println(topic)
	fmt.Println(conf)
	producer, err := rbt.NewDefaultProducer(beego.AppConfig.String("RocketmqBroker"), conf)
	producer.Start()
	if err != nil {
		fmt.Println(err)
	}
	msg := rbt.NewMessage(topic, data)
	if sendResult, err := producer.Send(msg); err != nil {
		fmt.Println("Sync sending fail!, %s", err.Error())
	} else {
		fmt.Println("sendResult", sendResult)

		//t.Logf("sendResult.msgId", sendResult.msgId)
		//t.Logf("sendResult.messageQueue", sendResult.messageQueue)
		//t.Logf("sendResult.queueOffset", sendResult.queueOffset)
		//t.Logf("sendResult.transactionId", sendResult.transactionId)
		//t.Logf("sendResult.offsetMsgId", sendResult.offsetMsgId)
		//t.Logf("sendResult.regionId", sendResult.regionId)
	}

	fmt.Println("Sync sending success!")
}
