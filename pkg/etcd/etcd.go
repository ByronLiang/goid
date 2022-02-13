package etcd

import (
	"strings"

	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var ClusterClient *clientv3.Client

func InitClusterClient() error {
	var err error
	hosts := strings.Split(viper.GetString("etcd.host"), ",")
	ClusterClient, err = clientv3.New(clientv3.Config{
		Endpoints: hosts,
	})
	if err != nil {
		return err
	}
	return nil
}

func CloseClusterClient() error {
	return ClusterClient.Close()
}
