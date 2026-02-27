package main

import (
	"fmt"
	"html/template"
	"os"
)

type Info struct {
	RiskName string
	ClientIP string
	Time     string
	TraceID  string
}

func templateDemo() {
	body := `{{ if eq .RiskName "IP提现频繁" "设备提现频繁" "会员提现频繁" "H5-开启调试工具"}}
              {"code":418,"data":{"value":""},"message": "5分钟内提现用户数已达上线，请稍后再试!, IP：{{ .ClientIP }}，时间：{{ .Time }}，访问ID：{{ .TraceID }}"}
            {{ else if eq .RiskName "会员提现金额"}}
              {"code":418,"data":{"value":""},"message": "超过最大提现金额，请稍后再试!, IP：{{ .ClientIP }}，时间：{{ .Time }}，访问ID：{{ .TraceID }}"}
            {{ else if eq .RiskName "会员提现时效"}}
              {"code":418,"data":{"value":""},"message": "注册时间满10分钟后可进行提现，请稍后再试！, IP：{{ .ClientIP }}，时间：{{ .Time }}，访问ID：{{ .TraceID }}"}
            {{ else if eq .RiskName "批量垃圾账号注册"}}
              {"code":418,"data":{"value":""},"message": "您的设备环境异常，无法注册！, IP：{{ .ClientIP }}，时间：{{ .Time }}，访问ID：{{ .TraceID }}"}
            {{ else }}
              {"code":418,"data":{"value":""},"message": "您的访问存在风险，请您稍后重试。, IP：{{ .ClientIP }}，时间：{{ .Time }}，访问ID：{{ .TraceID }}"}
            {{end}}`
	infos := Info{RiskName: "H5-开启调试工具"}

	//m := []string{"1", "2"}
	t, err := template.New("resBody").Parse(body)
	if err != nil {
		fmt.Println("-----", err)
	}
	if err := t.Execute(os.Stdout, infos); err != nil {
		fmt.Println(err)
	}
}
