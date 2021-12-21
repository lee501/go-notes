package main

import (
	"fmt"
	"regexp"
)

func main() {
	addrRe := regexp.MustCompile(`\p{Han}+省|\p{Han}+市|\p{Han}+州|\p{Han}+区|\p{Han}+县`)
	str := `河北省廊坊市三河市`
	re := addrRe.FindAllString(str, 5)
	for _, item := range re {
		fmt.Println(item)
	}
	fmt.Println(len(re))

	//军官
	officeRe := regexp.MustCompile("([\u4E00-\u9FA5])(字第)([0-9a-zA-Z]{4,8})(号?)")
	office := `2001988`
	res := officeRe.FindAllString(office, -1)
	fmt.Println(res)
	//护照
	passPortCardRe := regexp.MustCompile(`1[45][0-9]{7}|([P|p|S|s]\d{7})|([S|s|G|g]\d{8})|([Gg|Tt|Ss|Ll|Qq|Dd|Aa|Ff]\d{8})|([H|h|M|m]\d{8，10})`)
	passPortCard := "G12345678"
	fmt.Println(passPortCardRe.FindString(passPortCard))
}
