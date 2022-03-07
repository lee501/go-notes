package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	args := []string{"BodyMap", "token"}
	//"{\"BodyMap\":[{\"categoryId\":102,\"showType\":2},{\"categoryId\":5,\"showType\":3}],\"ClientIP\":\"114.226.163.45\",\"CookieMap\":{\"UM_distinctid\":\"17f49b577837f-0a90144b215f74-16347940-94354-17f49b5778814eQINGCLOUDELB=8106a03efa6091e099c2566b413b78cadf2bacdcbc445694f77b915de695fb25\"},\"HeaderMap\":{\"Cdn-Src-Ip\":[\"114.226.163.45\"],\"Connection\":[\"close\"],\"Content-Encoding\":[\"UTF-8\"],\"Content-Length\":[\"63\"],\"Content-Type\":[\"application/json; charset=UTF-8\"],\"Cookie\":[\"; UM_distinctid=17f49b577837f-0a90144b215f74-16347940-94354-17f49b5778814eQINGCLOUDELB=8106a03efa6091e099c2566b413b78cadf2bacdcbc445694f77b915de695fb25;\"],\"Referer\":[\"https://wap.airkunming.com/\"],\"User-Agent\":[\"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.5.979.68 Safari/537.36\"],\"X-Cdn-Src-Port\":[\"21758\"],\"X-Forwarded-For\":[\"114.226.163.45, 180.101.30.78, 180.101.30.78\"],\"X-Forwarded-Proto\":[\"https\"],\"X-Real-Ip\":[\"180.101.30.78\"],\"X-Via\":[\"1.1 nxzai40:14 (Cdn Cache Server V2.0), 1.1 PS-000-01tex79:1 (Cdn Cache Server V2.0)\"],\"X-Ws-Request-Id\":[\"621f2852_nxzai40_5846-59385\"]},\"Host\":\"wapnewsearch.kmair.net\",\"Method\":\"POST\",\"Path\":\"/search/cms/index\",\"Payload\":\"[{\\\"categoryId\\\":102,\\\"showType\\\":2},{\\\"categoryId\\\":5,\\\"showType\\\":3}]\",\"Referer\":\"https://wap.airkunming.com/\",\"Timestamp\":1646209106,\"TsMs\":1646209106445,\"UserAgent\":\"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.5.979.68 Safari/537.36\"}"
	//"{\"ClientIP\":\"114.226.163.45\",\"CookieMap\":{\"UM_distinctid\":\"17f49b47eac0-09238e1c8468e68-16347940-36d650-17f49b47eb1a3\"},\"HeaderMap\":{\"Cdn-Src-Ip\":[\"114.226.163.45\"],\"Connection\":[\"close\"],\"Cookie\":[\"; UM_distinctid=17f49b47eac0-09238e1c8468e68-16347940-36d650-17f49b47eb1a3\"],\"User-Agent\":[\"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.6.6512.30 Safari/537.36\"],\"X-Cdn-Src-Port\":[\"20927\"],\"X-Forwarded-For\":[\"114.226.163.45, 114.113.82.32, 114.113.82.32\"],\"X-Forwarded-Proto\":[\"https\"],\"X-Real-Ip\":[\"114.113.82.32\"],\"X-Via\":[\"1.1 inxzai36:11 (Cdn Cache Server V2.0), 1.1 wj32:6 (Cdn Cache Server V2.0)\"],\"X-Ws-Request-Id\":[\"621f2811_inxzai36_16959-19034\"]},\"Host\":\"b2c.kmair.net\",\"Method\":\"GET\",\"Path\":\"/\",\"Timestamp\":1646209041,\"TsMs\":1646209041922,\"UserAgent\":\"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.6.6512.30 Safari/537.36\"}"
	req3 := "{\"BodyMap\":{\"sign\":\"6SWNYGwmr5w1PAr1+sFDYZI+Ncy/kXhuOO8tUPTeJklAOVonOn8UnutTfYKd7MWkOWWiRBFlLcmH 35IjdqAlMA==\",\"timestamp\":1646209130,\"token\":\"a0a81d139a474bcba02f1fa41f5a4e27\"},\"ClientIP\":\"114.226.163.45\",\"CookieMap\":{\"UM_distinctid\":\"17f49b577837f-0a90144b215f74-16347940-94354-17f49b5778814e\"},\"HeaderMap\":{\"Cdn-Src-Ip\":[\"114.226.163.45\"],\"Connection\":[\"close\"],\"Content-Encoding\":[\"UTF-8\"],\"Content-Length\":[\"166\"],\"Content-Type\":[\"application/json; charset=UTF-8\"],\"Cookie\":[\";; UM_distinctid=17f49b577837f-0a90144b215f74-16347940-94354-17f49b5778814e\"],\"Referer\":[\"https://wap.airkunming.com/\"],\"User-Agent\":[\"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.5.979.68 Safari/537.36\"],\"X-Cdn-Src-Port\":[\"21816\"],\"X-Forwarded-For\":[\"114.226.163.45, 180.101.30.71, 180.101.30.71\"],\"X-Forwarded-Proto\":[\"https\"],\"X-Real-Ip\":[\"180.101.30.71\"],\"X-Via\":[\"1.1 nxzai40:7 (Cdn Cache Server V2.0), 1.1 PS-000-01Szz75:5 (Cdn Cache Server V2.0)\"],\"X-Ws-Request-Id\":[\"621f2856_nxzai40_5827-61635\"]},\"Host\":\"wapnewsso.kmair.net\",\"Method\":\"POST\",\"Path\":\"/sso/login/status\",\"Payload\":\"{\\\"timestamp\\\":1646209130,\\\"token\\\":\\\"a0a81d139a474bcba02f1fa41f5a4e27\\\",\\\"sign\\\":\\\"6SWNYGwmr5w1PAr1+sFDYZI+Ncy/kXhuOO8tUPTeJklAOVonOn8UnutTfYKd7MWkOWWiRBFlLcmH 35IjdqAlMA==\\\"}\",\"Referer\":\"https://wap.airkunming.com/\",\"Timestamp\":1646209110,\"TsMs\":1646209110441,\"UserAgent\":\"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.5.979.68 Safari/537.36\"}"
	var root map[string]interface{}
	json.Unmarshal([]byte(req3), &root)
	fmt.Println(json2string(root, args...))

	var d []byte = nil
	r := flattenTo(d, "")
	fmt.Println(r)
}

func json2string(root interface{}, args ...string) string {
	parent := root.(map[string]interface{})
	for _, v := range args {
		if parent[v] == nil {
			return ""
		}
		switch n := parent[v].(type) {
		// string
		case string:
			return n
		// string array: get first string
		case []interface{}:
			if len(n) == 0 {
				return ""
			}
			switch m := n[0].(type) {
			case string:
				return m
			case map[string]interface{}:
				parent = m
			default:
				b, err := json.Marshal(n)
				if err != nil {
					return ""
				}
				return string(b)
			}
		// map: go to next loop
		case map[string]interface{}:
			parent = n

		// number, bool
		case float64, bool:
			b, err := json.Marshal(n)
			if err != nil {
				return ""
			}
			return string(b)
		default:
			b, err := json.Marshal(n)
			if err != nil {
				return ""
			}
			log.Println(string(b))
			return string(b)
		}
	}
	b, err := json.Marshal(parent)
	if err != nil {
		return ""
	}
	return string(b)
}

func flattenTo(dst []byte, parent string) []byte {
	dst = append(dst, '{')
	return dst
}
