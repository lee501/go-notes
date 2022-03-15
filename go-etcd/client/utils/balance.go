package utils

import (
	"math/rand"
	"time"
)

type Balance struct {
	serviceList []*ServiceInfo
}

func NewBalance(serviceList []*ServiceInfo) *Balance {
	return &Balance{serviceList: serviceList}
}

// 随机算法
// 找出与serviceName相匹配的服务
// 保存在一个临时切片中
// 随机从这个临时切片中取出一个服务 并返回
func (this *Balance) Random(serviceName string) *ServiceInfo {
	tmp := make([]*ServiceInfo, 0)
	for _, service := range this.serviceList {
		if service.ServiceName == serviceName {
			tmp = append(tmp, service)
		}
	}

	// 没有匹配到serviceName的相关服务
	if len(tmp) == 0 {
		return nil
	} else if len(tmp) == 1 {
		return this.serviceList[0] // 只有一个服务的时候直接返回  rand.Intn(0) 会抛出异常
	} else {
		// 随机一个服务返回
		rand.Seed(time.Now().UnixNano()) //  设置一个种子 确保每次都是随机的（UnixNano 纳秒）
		i := rand.Intn(len(tmp) - 1)     // 以int型返回   [0,len(tmp)-1)   区间的一个随机数

		return this.serviceList[i]
	}

}
