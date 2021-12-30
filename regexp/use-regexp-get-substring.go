package main

import (
	"fmt"
	"regexp"
)

/*使用正则提取url字串*/

func main() {
	re :=  `^http://www.flysnow.org/([\d]{4})/([\d]{2})/([\d]{2})/([\w-]+).html$`
	reg := regexp.MustCompile(re)
	params := reg.FindStringSubmatch("http://www.flysnow.org/2018/01/20/golang-goquery-examples-selector.html")
	for _, item := range params {
		fmt.Println(item)
	}
	
	//使用\W首尾确定非字符数字，用()截取分词结果
	pattern := "\\W(1[45][0-9]{7})\\W|\\W([P|p|S|s]\\d{7})\\W|\\W([S|s|G|g]\\d{8})\\W|\\W([Gg|Tt|Ss|Ll|Qq|Dd|Aa|Ff]\\d{8})\\W|\\W([H|h|M|m]\\d{8，10})\\W"
	reg1 := regexp.MustCompile(pattern)
	p := reg1.FindStringSubmatch("测试141234567护照")
	for _, item := range p {
		fmt.Println(item)
	}

	email := "([a-zA-Z0-9_\\-\\.]+)@([a-zA-Z0-9_\\-\\.]+)\\.([a-zA-Z]{2,5})"
	e := regexp.MustCompile(email)
	emailr := e.FindStringSubmatch("23js@163.com")
	fmt.Println(emailr[0])
}
