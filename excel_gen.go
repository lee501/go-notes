package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {
	str := `[
    {
        "key": "身份证",
        "pattern": "(?:^|\\W|\\\\u0027)([1-9]{2}\\d(?:\\d{12}|\\*{12})\\d{2}[0-9Xx])(?:$|\\W)",
        "on": 1,
        "desc": "中华人民共和国居民身份证",
        "include": "",
        "exclude": "",
        "source": "",
        "level": "信息安全规范:个人身份信息:3级;数据出境规范:个人敏感信息:监管;金融数据安全规范:个人基本概况:3级;电信数据安全标准:用户身份相关信息:3级",
        "include_val": "",
        "exclude_val": "",
        "score": 80
    },
    {
        "key": "手机",
        "pattern": "(?:^|\\W|\\\\u0027)(1[3-9][0-9](?:[0-9]{4}|\\*{4})\\d{4})(?:$|\\W)",
        "on": 1,
        "desc": "号码为11位,其中各段有不同的编码方向:第1-3位—网络识别号;第4-7位—地区编码;第8-11位—用户号码",
        "include": "",
        "exclude": "",
        "source": "",
        "level": "信息安全规范:个人联系信息:3级;数据出境规范:个人敏感信息:监管;金融数据安全规范:个人联系信息:3级;电信数据安全标准:用户身份相关信息:3级",
        "include_val": "",
        "exclude_val": "",
        "score": 80
    }
]`
	type S struct {
		Key   string `json:"key"`
		Desc  string `json:"desc"`
		Level string `json:"level"`
		On    int    `json:"on"`
	}

	var res []S
	if err := json.Unmarshal([]byte(str), &res); err != nil {
		fmt.Println(err)
	}
	values := make([][]interface{}, 0)
	for _, r := range res {
		if strings.Contains(r.Level, "电信数据安全标准") {
			value := make([]interface{}, 0)
			ll := strings.Split(r.Level, ";")
			level := ll[len(ll)-1]
			value = append(value, []interface{}{r.Key, r.Desc, level, r.On}...)
			values = append(values, value)
		}
	}
	excel := excelize.NewFile()
	excels := make([]string, 4)
	defer excel.Close()
	for i, vv := range values {
		for j, cell := range []string{"A", "B", "C", "D"} {
			excels[j] = cell + strconv.Itoa(i+1)
		}
		for k, cell := range excels {
			if err := excel.SetCellValue("Sheet1", cell, vv[k]); err != nil {
				fmt.Println(err)
			}
		}
	}
	if err := excel.SaveAs("电信数据安全标准.xlsx"); err != nil {
		fmt.Println("2 ", err)
	}
}
