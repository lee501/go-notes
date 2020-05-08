//未知数据结构的JSON数据解码
/*
	解码未知结构的JSON，输出到一个空接口

	1.JSON的布尔值转换为GO中的bool类型
	2.数值转换为Go中的float64
	3.字符串转换为Go中的string
	4.JSON数组转换为 []interface{}类型
	5.JSON对象转换为 map[string]interface{}类型
	6.null转换为 nil
*/
package json

import (
	"encoding/json"
	"fmt"
)

func main() {
	b := []byte(`{
		"Title": "go语言编程",
		"Authors": ["lee", "anne"],
		"Publisher": "北京大学",
		"IsPublished": true,
		"Price": 32.9,
		"Sales": 100000
	}`)

	var res interface{}
	//res将会是一个键值对的map[string]interface{}结构
	/*
		map[string]interface{}{
			"Title": "go语言编程",
			"Authors": ["lee", "anne"],
			"Publisher": "北京大学",
			"IsPublished": true,
			"Price": 32.9,
			"Sales": 100000
		}
	*/
	err := json.Unmarshal(b, &res)
	if err != nil {
		fmt.Printf("解码错误：v%\n", err.Error())
	}

	//访问解码后的数据
	fmt.Println(res)
	book, ok := res.(map[string]interface{})
	fmt.Println(book)
	if ok {
		for k, v := range book {
			fmt.Println(v, "-----------")
			//v1的值为map中的value
			switch vl := v.(type) {
			case string:
				fmt.Println(k, "is string", vl)
			case int:
				fmt.Println(k, "is int", vl)
			case bool:
				fmt.Println(k, "is bool", vl)
			case []interface{}:
				fmt.Println(k, "is array", vl)
				for i, item := range vl {
					fmt.Println(i, item)
				}
			default:
				fmt.Println(k, "is other type not handle yet")
			}
		}
	}
}
