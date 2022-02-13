package handle

import (
	"context"

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

func RegisterLeafService(s *grpc.Server) {
	pb.RegisterLeafServer(s, &LeafSrv{})
}
