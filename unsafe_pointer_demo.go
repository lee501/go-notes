package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type SiteConfig struct {
	Host    string                 `json:"host"`
	On      bool                   `json:"on"`
	Configs map[string]interface{} `json:"configs"`
}

func main() {
	i := int64(1)
	var p *int32
	p = (*int32)(unsafe.Pointer(&i))
	fmt.Printf("-----use unsafe pointer转化类型:----%#v\n", *p)
	str := "[{\"host\": \"127.0.0.1\", \"on\": false, \"configs\": {\"time_range_matcher\":  {\"start\": \"2\" , \"end\": \"5\"}, \"region_matcher\": true, \"http_request_matcher\": [\"option\", \"get\"], \"static_html_matcher\": true}}]"
	var sites []SiteConfig
	err := json.Unmarshal([]byte(str), &sites)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sites)
}
