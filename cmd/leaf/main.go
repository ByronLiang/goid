package main

import (
	"log"
	"os"

	"github.com/ByronLiang/goid/pkg/db"

	"github.com/spf13/viper"

	"github.com/ByronLiang/goid/internal/leaf/route"
	"github.com/ByronLiang/goid/pkg/config"

	"github.com/ByronLiang/goid/internal/leaf/handle"
	"github.com/ByronLiang/goid/pkg/service"
	"github.com/ByronLiang/servant"
	"github.com/ByronLiang/servant/net"
)

func main() {
	err := config.NewConfig(os.Getenv("CONFIG_ENV"))
	if err != nil {
		log.Fatal(err)
		return
	}
	if viper.GetBool("leaf.debug") {
		service.Leaf.FakeLeafNode(3, 10)
	} else {
		err = db.InitDb()
		if err != nil {
			log.Fatal(err)
			return
		}
		leafList, err := db.LeafDao.GetAll()
		if err != nil {
			log.Fatal(err)
			return
		}
		service.Leaf.AddLeafNode(leafList,
			viper.GetInt64("leaf.buffer_size"),
			viper.GetInt64("leaf.percent"))
	}
	httpSrv := net.NewDefaultHttpServer(
		net.HttpAddress(viper.GetString("http.address")),
		net.HttpRouteGroup(route.InitHttpRouteGroup()),
	).InitRouteHandle()
	leafSrv := net.NewGRpc(net.GRpcAddress(viper.GetString("grpc.address")),
		net.GRpcIsRegistered(false)).
		SetRegisterHandler(handle.RegisterLeafService)
	serve := servant.NewServant(
		servant.Name(viper.GetString("leaf.name")),
		servant.AddServer(leafSrv),
		servant.AddServer(httpSrv))
	errs := serve.Run()
	for _, err := range errs {
		log.Println(err)
	}
}
