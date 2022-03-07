package main

import (
	"fmt"

	"github.com/cloudflare/ahocorasick"
)

func main() {
	dictionary := []string{"hello", "world", "世界", "google", "golang", "c++", "love"}
	ac := ahocorasick.NewStringMatcher(dictionary)

	ret := ac.Match([]byte("hello世界, hello google, i love golang!!!"))

	for index, i := range ret {
		fmt.Println(index, ": ", dictionary[i])
	}
	//params := []interface{}{"2", "5", "2", "2"}
	//s := "SELECT toUnixTimestamp(max(ts)) AS _ts, ip, countIf(event='book') AS book_count, uniq(contact_name) AS name_count, uniq(contact_phone) AS phone_count, SUM(passengers_count) AS pass_number, 1 AS _result FROM order_airline WHERE ts > now() - 3600 GROUP BY ip HAVING book_count > %s AND pass_number > %s AND ( name_count > %s OR phone_count > s% )"
	//
	//fmt.Println(FormatRule(s, params))
	//
	//fmt.Sprintf("SELECT toUnixTimestamp(max(ts)) AS _ts, ip, countIf(event='book') AS book_count, uniq(contact_name) AS name_count, uniq(contact_phone) AS phone_count, SUM(passengers_count) AS pass_number, 1 AS _result FROM order_airline WHERE ts > now() - 3600 GROUP BY ip HAVING book_count > %s AND pass_number > %s AND ( name_count > %s OR phone_count > s% )", params...)
}

func FormatRule(format string, params []interface{}) string {
	return fmt.Sprintf(format, params...)
}
