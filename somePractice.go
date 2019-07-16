package main

import (
	"errors"
	"net"
)

type Advert struct {
	ID			int64
	Name 		string
}

func main() {
	//who := "lee"
	//if len(os.Args) > 1 {
	//	who += strings.Join(os.Args[1:], " ")
	//}
	//who := GenerateImageAdvert("/api/image/1562207823.png")
	//fmt.Println(who)
	//s := make([]*Advert, 0)
	//s = append(s, &Advert{Name: "lee"})
	//for _, item := range s {
	//	println(item.Name)
	//}
	e := "59.255.255.255"
	s := "59.192.0.0"
	re,_ := IPString2Long(e)
	rs,_ := IPString2Long(s)
	println(re, rs)
}

func createString(img string) string {
	return `helle ` + img + ` /hello`
}

func GenerateImageAdvert(imgUrl string) string {
	div := `<div id="goimg" src="` + imgUrl + `"</div>`
	return div
}

func IPString2Long(ip string) (uint, error) {
	b := net.ParseIP(ip).To4()
	if b == nil {
		return 0, errors.New("invalid ipv4 format")
	}

	return uint(b[3]) | uint(b[2])<<8 | uint(b[1])<<16 | uint(b[0])<<24, nil
}