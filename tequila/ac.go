package main

import (
	"fmt"
	ahocorasick "github.com/petar-dambovaliev/aho-corasick"
)

var builder = ahocorasick.NewAhoCorasickBuilder(ahocorasick.Opts{
	AsciiCaseInsensitive: true,
	MatchOnlyWholeWords:  false,
	MatchKind:            ahocorasick.LeftMostFirstMatch,
	DFA:                  true,
})

var dict = []string{"air", "flight", "login", "signin", "signup", "register", "password", "sms", "captcha", "upload", "download", "query", "list", "search", "pay", "checkout", "book", "order", "reserve", "calendar", "coupon"}

// ExampleAhoCorasick 演示Aho-Corasick算法的示例函数
func ExampleAhoCorasick() {
	var res []interface{}
	res = append(res, builder.Build(dict))
	acid := 0
	prefix := "szffp-webloginloginMcodeSF"
	matches := res[acid].(ahocorasick.AhoCorasick).FindAll(prefix)
	var ret []int
	for _, match := range matches {
		ret = append(ret, match.Pattern())
	}
	fmt.Println("AhoCorasick 匹配结果:", ret)
}
