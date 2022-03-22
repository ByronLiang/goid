package rpc

import (
	"context"

	"github.com/ByronLiang/goid/pkg/pb"
	"github.com/ByronLiang/servant/net"
	"google.golang.org/grpc"
)

var Leaf *leaf

type leaf struct {
	Cli        pb.LeafClient
	connection *grpc.ClientConn
}

func InitLeafCli(address string) error {
	con, err := net.DialInsecure(context.Background(), net.WithEndpoint(address))
	if err != nil {
		return err
	}
	Leaf = &leaf{
		Cli:        pb.NewLeafClient(con),
		connection: con,
	}
	return nil
}
