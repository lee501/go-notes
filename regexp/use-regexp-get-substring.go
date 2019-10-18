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
}
