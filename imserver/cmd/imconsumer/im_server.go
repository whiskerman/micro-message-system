package main

/*
相当于是kafka的consumer的服务
对logic层的调用
两个进程分别调用：
	Subscribe订阅kafka的
	Run() 发送消息给websocket显示
*/

import (
	"flag"
	"log"
	"micro-message-system/imserver/logic"
	"micro-message-system/imserver/util"

	"github.com/asim/go-micro/plugins/broker/kafka/v3"
	etcdv3 "github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/broker"
	"github.com/asim/go-micro/v3/config"
	"github.com/asim/go-micro/v3/registry"
	"github.com/urfave/cli/v2"

	imConfig "micro-message-system/imserver/cmd/config"
)

func main() {
	imFlag := cli.StringFlag{
		Name:  "f",
		Value: "./config/config_im.json",
		Usage: "please use xxx -f config_im.json",
	}
	configFile := flag.String(imFlag.Name, imFlag.Value, imFlag.Usage)
	flag.Parse()
	conf := new(imConfig.ImConfig)

	if err := config.LoadFile(*configFile); err != nil {
		log.Fatal(err)
	}
	if err := config.Scan(conf); err != nil {
		log.Fatal(err)
	}
	etcdRegisty := etcdv3.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = conf.Etcd.Address
		})
	kafkaRegistry := kafka.NewBroker(func(options *broker.Options) {
		options.Addrs = conf.Kafka.Address
	})
	service := micro.NewService(
		micro.Name(conf.Server.Name),
		micro.Registry(etcdRegisty),
		micro.Version(conf.Version),
		micro.Flags(&imFlag),
	)

	log.Printf("has start listen topic %s", conf.Kafka.Topic)
	kafkaBroker, err := util.NewKafkaBroker(conf.Kafka.Topic, kafkaRegistry)
	log.Printf("kafkaBroker:%v\n", kafkaBroker)
	if err != nil {
		log.Fatal(err)
	}
	imServer, err := logic.NewImServer(kafkaBroker,
		func(im *logic.ImServer) {
			im.Address = conf.Port
		})
	log.Printf("imServer:%v\n", imServer)
	if err != nil {
		log.Fatal(err)
	}
	go imServer.Subscribe()
	go imServer.Run()
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
