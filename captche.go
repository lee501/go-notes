package main

import (
	"fmt"
	"io/ioutil"

	"github.com/VictoriaMetrics/fastcache"
)

func main() {
	path := "./captcha.html"
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("----", err)
	}
	fmt.Println(len(bs))
	pc := fastcache.New(1)
	pc.Set([]byte("page"), bs)
	fmt.Println(pc.Get(nil, []byte("page")))
	var val string
	var v interface{} = 1
	val = v.(string)
	fmt.Println(val)

}
