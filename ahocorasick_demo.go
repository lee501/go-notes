package main

import (
	"fmt"
	"github.com/cloudflare/ahocorasick"
)

func main()  {
	dictionary := []string{"hello", "world", "世界", "google", "golang", "c++", "love"}
	ac := ahocorasick.NewStringMatcher(dictionary)


	ret := ac.Match([]byte("hello世界, hello google, i love golang!!!"))

	for _, i := range ret {
		fmt.Println(dictionary[i])
	}
	params := []interface{}{"0.75", "1"}
	s := "SELECT toUnixTimestamp(ts) AS _ts, divide(minus(total_space, free_space), total_space) >= %s AS _result FROM system_nodes WHERE _result = %s AND _ts = (SELECT toUnixTimestamp(max(ts)) FROM system_nodes )"

	fmt.Println(FormatRule(s, params))
}

func FormatRule(format string, params []interface{}) string {
	return fmt.Sprintf(format, params...)
}