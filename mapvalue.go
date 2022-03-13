package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type ExtraConfig map[string]interface{}

var test ExtraConfig = map[string]interface{}{"customer": map[interface{}]interface{}{
	"body-tem": `{"code":601,"data":{"value":""},"message": "对不起，尊敬的旅客，您的访问存在风险，请您稍后重试。如有疑问请拨打0871-96598。感谢您的理解。IP：{{ .ClientIP }}，时间：{{ .Time }}，访问ID：{{ .TraceID }}"}`,
	"code":     601,
	"data":     "no",
}}

func main() {
	b, err := jsoniter.Marshal(test)
	if err != nil {
		fmt.Println("-------", err)
	}
	fmt.Println("-----res----", string(b))
}
