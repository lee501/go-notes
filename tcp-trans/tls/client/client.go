package main

import (
	"crypto/tls"
	"fmt"
	"log"
)

func main() {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", "127.0.0.1:443", conf)
	if err != nil {
		log.Println(err)
		return
	}
	token := fmt.Sprintf("zs-token: %v", "ndop3gsbjQlxityQLhk57A4CudfY7zbFcSRczc5zJ0c\n\r")
	defer conn.Close()
	n, err := conn.Write([]byte(token))
	if err != nil {
		log.Println(n, err)
		return
	}
	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}
	println(string(buf[:n]))
}
