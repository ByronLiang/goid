package main

import (
	"log"

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
	err := config.NewConfig()
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
		err = service.Leaf.InitLeaf()
		if err != nil {
			log.Fatal(err)
			return
		}
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
