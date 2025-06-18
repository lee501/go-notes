package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

type Proxy func() (m string)

type Backend func() Proxy

type defaultFactory struct {
	backend Backend
}

func backendDemo() {

	str := `{
	"Body": "{\"code\": 200, \"orderId\": 1, \"orderStatus\": 0, \"orderTotalAmount\": 256.0, \"productId\": \"27e44731-0e1b-4244-9933-003bae53eff0\", \"productCount\": 4, \"buyerName\": \"\\u5f90\\u6ce2\", \"buyerPhone\": \"18612345678\", \"receiverName\": \"\\u6768\\u82f1\", \"receiverPhone\": \"15887654321\", \"receiverAddress\": \"\\u6e56\\u5317\\u7701\\u6b66\\u6c49\\u5e02\\u5357\\u957f\\u9ec4\\u8857q\\u5ea7 357996\"}",
	"Headers": {
		"Connection": [
			"Keep-Alive"
		],
		"Content-Length": [
			"363"
		],
		"Content-Type": [
			"application/json;charset=utf-8"
		],
		"Date": [
			"Fri, 13 Jan 2023 08:00:36 GMT"
		]
	},
	"Status": 200,
	"BodyMap": {
		"code": 200,
		"orderId": 1,
		"orderStatus": 0,
		"orderTotalAmount": 256,
		"productId": "27e44731-0e1b-4244-9933-003bae53eff0",
		"productCount": 4,
		"buyerName": "徐波",
		"buyerPhone": "18612345678",
		"receiverName": "杨英",
		"receiverPhone": "15887654321",
		"receiverAddress": "湖北省武汉市南长黄街q座 357996"
	}
}`
	fmt.Println(Json2string(str, "BodyMap"))
}

func Json2string(root interface{}, args ...string) string {
	if root == nil {
		return ""
	}

	if len(args) == 0 {
		switch val := root.(type) {
		case string:
			return val
		default:
			if b, err := json.Marshal(root); err == nil && !bytes.Equal(b, []byte("null")) {
				return string(b)
			}
			return ""
		}
	}

	switch val := root.(type) {
	case map[string]interface{}:
		return Json2string(val[args[0]], args[1:]...)
	case []interface{}:
		n, err := strconv.Atoi(args[0])
		if err != nil && len(val) > 0 {
			return Json2string(val[0], args...)
		} else if n < len(val) {
			return Json2string(val[n], args[1:]...)
		} else {
			return ""
		}
	case string:
		var result []interface{}
		err := json.Unmarshal([]byte(val), &result)
		if err != nil {
			return val
		}
		return Json2string(result, args...)
	default:
		return ""
	}
}
