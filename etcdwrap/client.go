package etcdwrap

import (
	"strings"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdKV struct {
	Key   string
	Value []byte
}

type EtcdClientConfig struct {
	clientv3.Config
}

func (cfg *EtcdClientConfig) NewEtcdClient() (*clientv3.Client, error) {
	return clientv3.New(cfg.Config)
}

type etcdClient struct {
	sync.RWMutex
	cli *clientv3.Client
}

func (etcdCli *etcdClient) SetEtcdClient(cli *clientv3.Client) {
	etcdCli.Lock()
	etcdCli.cli = cli
	etcdCli.Unlock()
}

func (etcdCli *etcdClient) GetEtcdClient() *clientv3.Client {
	etcdCli.RLock()
	res := etcdCli.cli
	etcdCli.RUnlock()
	return res
}

var etcdCli etcdClient

func GetClient() *clientv3.Client {
	return etcdCli.GetEtcdClient()
}

func New(endpoints string) (cli *clientv3.Client, err error) {
	cfg := EtcdClientConfig{
		clientv3.Config{
			Endpoints:   strings.Split(endpoints, ","),
			DialTimeout: 5 * time.Second,
		},
	}
	cli, err = cfg.NewEtcdClient()
	if err != nil {
		return
	}

	etcdCli.SetEtcdClient(cli)
	return
}
