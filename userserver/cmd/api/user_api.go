package main

import (
	"flag"
	"log"

	etcdv3 "github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3/config"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/web"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli/v2"

	//"github.com/micro/go-plugins/registry/etcdv3"
	//"github.com/micro/micro/v3/cmd/cli"

	userConfig "micro-message-system/userserver/cmd/config"
	"micro-message-system/userserver/controller"
	"micro-message-system/userserver/logic"
	"micro-message-system/userserver/models"
)

func main() {
	userRpcFlag := cli.StringFlag{
		Name:  "f",
		Value: "./config/config_api.json",
		Usage: "please use xxx -f config_rpc.json",
	}
	configFile := flag.String(userRpcFlag.Name, userRpcFlag.Value, userRpcFlag.Usage)
	flag.Parse()
	conf := new(userConfig.ApiConfig)

	if err := config.LoadFile(*configFile); err != nil {
		log.Fatal(err)
	}
	if err := config.Scan(conf); err != nil {
		log.Fatal(err)
	}
	engineUser, err := gorm.Open(conf.Engine.Name, conf.Engine.DataSource)
	if err != nil {
		log.Fatal(err)
	}
	etcdRegisty := etcdv3.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = conf.Etcd.Address
		})
	service := web.NewService(
		web.Name(conf.Server.Name),
		web.Registry(etcdRegisty),
		web.Version(conf.Version),
		web.Flags(&userRpcFlag),
		web.Address(conf.Port),
	)
	router := gin.Default()
	userModel := models.NewMembersModel(engineUser)
	userLogic := logic.NewUserLogic(userModel)
	userController := controller.NewUserController(userLogic)
	userRouterGroup := router.Group("/user")
	{
		userRouterGroup.POST("/login", userController.Login)
		userRouterGroup.POST("/register", userController.Register)
	}
	service.Handle("/", router)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
