package etcdwrap

import (
	"context"
	"errors"
	"log"
	"time"

	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type ServiceRegister struct {
	LeaseTTL int64
	Renewal  time.Duration
	Key      string
	Value    string
	stopChan chan bool
}

func (sr *ServiceRegister) Start() error {
	cli := GetClient()
	if cli == nil {
		log.Printf("[etcd error] etcd client is not found")
		return errors.New("etcd client not found")
	}

	// try to stop previous routine
	_ = sr.Stop()

	renewalTimer := time.NewTicker(sr.Renewal)

	// start routine to do service register
	stopChan := make(chan bool, 2)
	go func() {
		var etcdLease clientv3.Lease
		var etcdLeaseId *clientv3.LeaseID

		for {
			if etcdLease == nil {
				etcdLease = clientv3.NewLease(cli)
			}

			if etcdLeaseId == nil {
				ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
				leaseGrantResp, err := etcdLease.Grant(ctx, sr.LeaseTTL)
				cancel()
				if err != nil {
					log.Printf("[etcd error] grant lease failed. key: %s, error: %s\n", sr.Key, err.Error())
					time.Sleep(2 * time.Second)
					continue
				}
				etcdLeaseId = &leaseGrantResp.ID
				err = PutWithTimeout(2*time.Second, sr.Key, sr.Value, clientv3.WithLease(*etcdLeaseId))
				if err != nil {
					log.Printf("[etcd error] lease put kv failed. key: %s, error: %s\n", sr.Key, err.Error())
					time.Sleep(2 * time.Second)
					continue
				}
			}

			select {
			case <-renewalTimer.C:
				ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
				_, err := etcdLease.KeepAliveOnce(ctx, *etcdLeaseId)
				cancel()
				if err == rpctypes.ErrLeaseNotFound {
					// 正常 renewal 时, etcd lease 未找到
					// 停止当前 routine， 启动新的 routine
					// 比如当使用断点调试导致 etcd lease ttl 触发后删除了 lease 的情况
					log.Printf("[etcd error] etcd lease id [%d] is not found.\n", *etcdLeaseId)
					etcdLeaseId = nil
				} else if err != nil {
					log.Printf("[etcd error] lease keep alive failed. key: %s, error: %s\n", sr.Key, err.Error())
				}
			case <-stopChan:
				if etcdLease != nil && etcdLeaseId != nil {
					// revoke etcd lease
					ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
					_, err := etcdLease.Revoke(ctx, *etcdLeaseId)
					cancel()
					if err != nil {
						log.Printf("[etcd error] revoke lease failed. key: %s, error: %s\n", sr.Key, err.Error())
					}
				}
				goto END
			}
		}

	END:
		log.Printf("[etcd] register '%s' routine finish", sr.Key)
	}()

	sr.stopChan = stopChan

	return nil
}

func (sr *ServiceRegister) Stop() (err error) {
	cli := GetClient()
	if cli == nil {
		err = errors.New("etcd client not found")
		return
	}

	if sr.stopChan != nil {
		// stop previous registery routine
		sr.stopChan <- true
		close(sr.stopChan)
		sr.stopChan = nil
	}

	return
}
