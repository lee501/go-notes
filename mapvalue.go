package main

import (
	"fmt"
	"regexp"
)

type ExtraConfig map[string]interface{}

var test ExtraConfig = map[string]interface{}{"customer": map[interface{}]interface{}{
	"body-tem": `{"code":601,"data":{"value":""},"message": "对不起，尊敬的旅客，您的访问存在风险，请您稍后重试。如有疑问请拨打0871-96598。感谢您的理解。IP：{{ .ClientIP }}，时间：{{ .Time }}，访问ID：{{ .TraceID }}"}`,
	"code":     601,
	"data":     "no",
}}

func main() {
	re := `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
	reg := regexp.MustCompile(re)
	str := "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
	str2 := "AD80::ABAA:0000:00C2:0002"
	fmt.Println(reg.FindString(str))
	fmt.Println(reg.FindString(str2))
}
