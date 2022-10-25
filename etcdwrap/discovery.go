package etcdwrap

import "sync"

type HealthyService struct {
	Key   string
	Value []byte
}

type DiscoveryCallback func(string, []HealthyService)

type ServiceDiscovery struct {
	ServiceKey string
	Callback   DiscoveryCallback
	watcher    *Watcher
	serviceMap sync.Map
}

func (sd *ServiceDiscovery) Start() error {
	if sd.watcher != nil {
		return nil
	}

	sd.watcher = &Watcher{
		EtcdKey: sd.ServiceKey,
		Callback: func(s string, wr WatcherResult) {
			switch wr.EventType {
			case WatcherEventTypeCurrent:
				// clear current map
				sd.serviceMap.Range(func(key interface{}, _ interface{}) bool {
					sd.serviceMap.Delete(key)
					return true
				})
				for _, kv := range wr.Kvs {
					sd.serviceMap.Store(kv.Key, kv.Value)
				}
			case WatcherEventTypePut:
				for _, kv := range wr.Kvs {
					sd.serviceMap.Store(kv.Key, kv.Value)
				}
			case WatcherEventTypeDelete:
				sd.serviceMap.Delete(s)
			}
			sd.callback()
		},
	}

	sd.watcher.watchAsync()
	return nil
}

func (sd *ServiceDiscovery) callback() {
	if sd.Callback == nil {
		return
	}
	res := []HealthyService{}
	sd.serviceMap.Range(func(key, value interface{}) bool {
		etcdKeyBytes, ok := key.([]byte)
		if !ok {
			return true
		}
		etcdValueBytes, ok := value.([]byte)
		if !ok {
			return true
		}

		res = append(res, HealthyService{
			Key:   string(etcdKeyBytes),
			Value: etcdValueBytes,
		})
		return true
	})
	sd.Callback(sd.ServiceKey, res)
}
