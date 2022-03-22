package handle

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/ByronLiang/goid/pkg/pb"
	"github.com/ByronLiang/goid/pkg/service"
	"google.golang.org/grpc"
)

type LeafSrv struct {
}

func (l LeafSrv) Query(ctx context.Context, req *pb.LeafRequest) (*pb.LeafResponse, error) {
	num, err := service.Leaf.Get(req.Domain)
	if err != nil {
		return nil, err
	}
	return &pb.LeafResponse{Number: num}, nil
}

func (l LeafSrv) Stop(ctx context.Context, req *pb.LeafRequest) (*empty.Empty, error) {
	err := service.Leaf.Stop(req.Domain)
	return nil, err
}

func RegisterLeafService(s *grpc.Server) {
	pb.RegisterLeafServer(s, &LeafSrv{})
}
