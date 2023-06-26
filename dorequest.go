package main

import (
	"fmt"
)

func main() {
	// 准备要发送的数据

	//jsonData := `{"msgtype":"text","text":{"content":"api\n 告警名称: [测试]告警验证\n 告警信息: [{\"metric\":\"AQI=\",\"risk\":\"[测试]告警验证\",\"ts\":1685515430}]"},"at":{"isAtAll":false}}`
	//
	//// 创建一个请求对象
	//req, _ := http.NewRequest("POST", "https://oapi.dingtalk.com/robot/send?access_token=ac36dd5f2c896a44a279d40f8b1487977f4b2bdcdb83ca9a4f539b58d2d7eaff", bytes.NewBuffer([]byte(jsonData)))
	//req.Header.Set("Content-Type", "application/json")
	//
	//// 发送请求
	//client := &http.Client{}
	//resp, err := client.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()

	// ...
	js := "msgi1"
	//b, err := json.Marshal(js)
	//if err != nil {
	//	fmt.Println("1111")
	//}
	fmt.Println(js[:2])
	fmt.Println(js[3:])
}
