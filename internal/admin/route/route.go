package route

import (
	"net/http"

	"github.com/ByronLiang/goid/internal/admin/handle"
	"github.com/ByronLiang/servant/net"
)

func InitHttpRouteGroup() []net.ApiGroupPath {
	leafPaths := []net.ApiPath{
		{
			Method:  http.MethodPost,
			Path:    "",
			Handler: handle.AddLeaf,
		},
		{
			Method:  http.MethodGet,
			Path:    "",
			Handler: handle.GetLeaf,
		},
		{
			Method:  http.MethodPut,
			Path:    "",
			Handler: handle.UpdateLeaf,
		},
	}
	leafRouteGroup := net.ApiGroupPath{
		Prefix: "/api/admin/leaf",
		Paths:  leafPaths,
	}
	return []net.ApiGroupPath{leafRouteGroup}
}
