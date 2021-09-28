package main

import (
	"flag"
	"fmt"

	rl "github.com/juju/ratelimit"
	"github.com/urfave/cli/v2"

	etcdv3 "github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/plugins/transport/grpc/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/config"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/web"

	"log"
	cfg "micro-message-system/common/config"
	uploadRpcConfig "micro-message-system/uploadserver/config"
	upProto "micro-message-system/uploadserver/protos"
	"micro-message-system/uploadserver/route"
	upRpc "micro-message-system/uploadserver/rpcserverimpl"
	"os"

	_ "github.com/asim/go-micro/plugins/registry/kubernetes/v3"
	"github.com/asim/go-micro/plugins/wrapper/ratelimiter/ratelimit/v3"
)

func startRPCService() {
	uploadRpcFlag := cli.StringFlag{
		Name:  "f",
		Value: "./config/config_rpc.json",
		Usage: "please use xxx -f config_rpc.json",
	}
	configFile := flag.String(uploadRpcFlag.Name, uploadRpcFlag.Value, uploadRpcFlag.Usage)
	flag.Parse()

	conf := new(uploadRpcConfig.RpcConfig)

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
	b := rl.NewBucketWithRate(float64(conf.Server.RateLimit), int64(conf.Server.RateLimit))
	service := micro.NewService(
		micro.Name(conf.Server.Name),
		micro.Registry(etcdRegisty),
		micro.Version(conf.Version),
		micro.Transport(grpc.NewTransport()),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(b, false)),
		micro.Flags(&uploadRpcFlag),
	)
	service.Init()

	upProto.RegisterUploadServiceHandler(service.Server(), new(upRpc.Upload))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func startAPIService() {
	uploadApiFlag := cli.StringFlag{
		Name:  "f",
		Value: "./config/config_api.json",
		Usage: "please use xxx -f config_api.json",
	}
	configFile := flag.String(uploadApiFlag.Name, uploadApiFlag.Value, uploadApiFlag.Usage)
	flag.Parse()
	conf := new(uploadRpcConfig.ApiConfig)

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
	router := route.Router()
	service := web.NewService(
		web.Registry(etcdRegisty),
		web.Version(conf.Version),
		web.Flags(&uploadApiFlag),
		web.Address(conf.Port),
	)
	service.Handle("/", router)
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	os.MkdirAll(cfg.TempLocalRootDir, 0777)
	os.MkdirAll(cfg.TempPartRootDir, 0777)

	// api 服务
	startAPIService()

	// rpc 服务
	go startRPCService()
}
