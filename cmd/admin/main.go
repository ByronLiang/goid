package main

import (
	"log"

	"github.com/ByronLiang/goid/pkg/rpc"

	"github.com/ByronLiang/goid/pkg/db"

	"github.com/ByronLiang/goid/internal/admin/route"
	"github.com/ByronLiang/goid/pkg/config"
	"github.com/ByronLiang/servant"
	"github.com/ByronLiang/servant/net"
	"github.com/spf13/viper"
)

func main() {
	err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.InitDb()
	if err != nil {
		log.Fatal(err)
		return
	}
	// init gRPC leaf client
	rpc.InitLeafCli("")
	httpSrv := net.NewDefaultHttpServer(
		net.HttpAddress(viper.GetString("http.address")),
		net.HttpRouteGroup(route.InitHttpRouteGroup()),
	).InitRouteHandle()
	serve := servant.NewServant(servant.Name(viper.GetString("admin.name")), servant.AddServer(httpSrv))
	errs := serve.Run()
	for _, err := range errs {
		log.Println(err)
	}
}
