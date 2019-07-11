package main

import (
	"fmt"
)

func main() {
	//who := "lee"
	//if len(os.Args) > 1 {
	//	who += strings.Join(os.Args[1:], " ")
	//}
	who := GenerateImageAdvert("/api/image/1562207823.png")
	fmt.Println(who)
}

func createString(img string) string {
	return `helle ` + img + ` /hello`
}

func GenerateImageAdvert(imgUrl string) string {
	div := `<div id="goimg" src="` + imgUrl + `"</div>`
	return div
}