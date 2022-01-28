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

func main() {
	var res []interface{}
	res = append(res, builder.Build(dict))
	acid := 0
	prefix := "szffp-webloginloginMcodeSF"
	matches := res[acid].(ahocorasick.AhoCorasick).FindAll(prefix)
	var ret []int
	for _, match := range matches {
		ret = append(ret, match.Pattern())
	}
	fmt.Println(ret)
}
