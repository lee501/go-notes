package utils

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	dialTimeout = 10
	leaseTime   = 30
)

var (
	once   sync.Once
	logger *zap.Logger
)

func init() {
	once.Do(func() {
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.Lock(os.Stdout),
			zapcore.InfoLevel,
		)
		caller := zap.AddCaller()
		dev := zap.Development()
		logger = zap.New(core, caller, dev)
	})
}

//注册服务
type Service struct {
	client *clientv3.Client

	key         string
	addr        string //服务地址
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse
}

func NewService(prefix, serverName, addr string) (*Service, error) {
	//etcd config
	config := clientv3.Config{
		Endpoints:   []string{},
		DialTimeout: dialTimeout * time.Second,
	}
	cli, err := clientv3.New(config)
	if err != nil {
		logger.Info("etcd client init failed: ", zap.String("etdc err", fmt.Sprintf("%s", err)))
		return nil, err
	}
	key := prefix + "/" + serverName + "/" + addr
	return &Service{cli, key, addr, nil}, nil
}

func (s *Service) RegistryServe() error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), dialTimeout*time.Second)
	defer func() {
		if err != nil {
			cancel()
		}
	}()

	kv := clientv3.NewKV(s.client)
	//租约
	lease := clientv3.NewLease(s.client)
	//etcdctl lease grand 30 second
	lresp, err := lease.Grant(ctx, leaseTime)
	if err != nil {
		return err
	}

	//etcdctl put key value --lease=xxxxx(leaseId)
	_, err = kv.Put(ctx, s.key, s.addr, clientv3.WithLease(lresp.ID))
	if err != nil {
		return err
	}
	//s.client.KeepAlive()
	//set leaseTime second expired, 定时续租
	//etcdctl lease keep-alive leaseId
	keepAliveCh, err := lease.KeepAlive(ctx, lresp.ID)
	if err != nil {
		return err
	}
	s.keepAliveCh = keepAliveCh
	go s.keepAlive()
	return nil
}

func (s *Service) keepAlive() {
	ticker := time.NewTicker(dialTimeout * time.Second)
	for {
		select {
		case <-ticker.C:
			if s.keepAliveCh == nil {
				s.RegistryServe()
			}
		case keep := <-s.keepAliveCh:
			if keep == nil {
				s.RegistryServe()
			}

		}
	}
}

func (s *Service) UnRegister() error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(dialTimeout)*time.Second)
	defer func() {
		if err != nil {
			cancel()
		}
	}()
	kv := clientv3.NewKV(s.client)
	_, err = kv.Delete(ctx, s.key)
	return err
}
