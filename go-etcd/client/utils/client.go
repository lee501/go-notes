package utils

import (
	"context"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

//服务发现
type Client struct {
	client   *clientv3.Client
	services []*ServiceInfo
}

type ServiceInfo struct {
	ServiceName    string
	ServiceAddress string
}

func NewServiceInfo(serviceName string, serviceAddress string) *ServiceInfo {
	return &ServiceInfo{
		ServiceName:    serviceName,
		ServiceAddress: serviceAddress,
	}
}

func NewClient() (*Client, error) {
	// 配置 etcd
	config := clientv3.Config{
		Endpoints:   []string{"192.168.137.132:23791", "192.168.137.132:23792", "192.168.137.132:23793"}, // 端点
		DialTimeout: 10 * time.Second,                                                                    // 超时时间
	}

	cli, err := clientv3.New(config) // 初始化 etcd 客户端
	if err != nil {
		return nil, err
	}

	return &Client{client: cli}, nil
}

func (this *Client) LoadService() error {
	kv := clientv3.NewKV(this.client)

	services, err := kv.Get(context.Background(), "/services/", clientv3.WithPrefix())

	if err != nil {
		return err
	}

	for _, serviceKvs := range services.Kvs {
		this.ParseService(serviceKvs.Key, serviceKvs.Value)
	}

	return nil
}

// 解析key  --- 服务注册时约定的key规则是  /services/(服务ID)/(服务名)
func (this *Client) ParseService(key []byte, value []byte) {
	reg := regexp.MustCompile("/services/(.*)/(\\w+)") // 初始华需要编译的正则表达式

	//匹配key是否包含正则表达式re的任何匹配项。
	if reg.Match(key) {
		names := reg.FindSubmatch(key) // 找到key与正则re的匹配项
		serviceName := names[2]

		// 将匹配到的服务保存到services切片中
		this.services = append(this.services, NewServiceInfo(string(serviceName), string(value)))
	}

}

// 随机获取一个服务
func (this *Client) GetService(serviceName string) *ServiceInfo {
	balance := NewBalance(this.services)
	serviceInfo := balance.Random(serviceName) //随机获取一个服务
	return serviceInfo
}

// 调用服务方法 （把发送HTTP请求的一个过程封装在一个端点函数里）
// info 服务信息 | method 请求方法 | request 解析请求参数(就是拼接url参数的函数)
func (this *Client) Call(info *ServiceInfo, method string, request EncodeRequest) EndPoint {
	return func(ctx context.Context, requestParam interface{}) (responseResult interface{}, err error) {
		httpClient := http.DefaultClient                                                // 初始化client
		httpRequest, err := http.NewRequest(method, "http://"+info.ServiceAddress, nil) // 返回一个新的服务器访问请求
		if err != nil {
			return nil, err
		}
		err = request(ctx, httpRequest, requestParam) // 将访问请求的urlpath带上参数
		if err != nil {
			return nil, err
		}
		// 发送HTTP请求并返回HTTP响应
		resp, err := httpClient.Do(httpRequest) // 通常将使用Get，Post或PostForm代替Do
		if err != nil {
			return nil, err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return string(body), nil
	}
}
