//使用内置库encoding/base64编码
package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func base64Demo() {
	input := []byte("hello go")

	//	base64编码
	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Println(encodeString)

	//	base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(string(decodeBytes))

	//	url中，需要使用URLEncoding
}
