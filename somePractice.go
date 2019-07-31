package main

import (
	"errors"
	"fmt"
	"net"
	"reflect"
)

type Advert struct {
	ID			int64
	Name 		string
}

type IpRange struct {
	Start string `json: "start"`
	End	  string `json: "end"`
	Source  string `json: "source"`
}

func main() {
	//ip := new(IpRange)
	//m := []byte(`{"start": "60.200.0.0", "end": "60.203.255.255", "source": "\u5e7f\u7535"}`)
	//fmt.Println(string(m))
	//json.Unmarshal(m, &ip)
	//fmt.Println(ip)
	//arr := []int{1,2,3,4}
	//TestArgs(1, arr...)
	var a int8 = 3  //0000 0011 -> 0000 0011 -0000 0011
	fmt.Printf("%b\n", a)
	var b int8 = -3  //0000 0011 -> 1111 1100 -> 1111 1101
	fmt.Printf("%b\n", b)
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

func TestArgs(first int, arg ...interface{}) {
	fmt.Println(reflect.TypeOf(arg))
	fmt.Println(first, arg)
}