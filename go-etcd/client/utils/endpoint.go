package utils

import (
	"context"
	"net/http"
)

// 发关Http请求的函数   requestParam请求参数  responseResult响应结果
type EndPoint func(ctx context.Context, requestParam interface{}) (responseResult interface{}, err error)

// 解析Http请求函数
type EncodeRequest func(context.Context, *http.Request, interface{}) error
