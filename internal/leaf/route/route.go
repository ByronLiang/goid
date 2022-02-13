package route

import (
	"net/http"

	"github.com/ByronLiang/goid/internal/leaf/handle"

	"github.com/ByronLiang/servant/net"
)

func InitHttpRouteGroup() []net.ApiGroupPath {
	leafPaths := []net.ApiPath{
		{
			Method:  http.MethodGet,
			Path:    "/:domain",
			Handler: handle.QueryLeaf,
		},
	}
	leafRouteGroup := net.ApiGroupPath{
		Prefix: "/api/leaf",
		Paths:  leafPaths,
	}
	return []net.ApiGroupPath{leafRouteGroup}
}
