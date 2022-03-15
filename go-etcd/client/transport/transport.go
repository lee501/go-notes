package transport

import (
	"context"
	"net/http"
	"strconv"
)

// 产品详情请求的参数   /product/id
type ProductDetailRequest struct {
	ProductId int
}

func ProductDetailEncode(ctx context.Context, httpRequest *http.Request, requestParam interface{}) error {
	pdr := requestParam.(ProductDetailRequest)
	httpRequest.URL.Path += "/product/" + strconv.Itoa(pdr.ProductId)
	return nil
}
