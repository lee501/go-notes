//JSON数据的流式读写
//go内建的encode/json包的Encoder和Decoder类型
//func NewDecoder(r io.Reader) *Decoder
//func NewEncoder(w io.Writer) *Encoder
package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	//从标准输入流中读取json
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string] interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		for k := range v {
			//字符串可以进行相等判断
			if k != "Title" {
				//删除map中的键
				delete(v, k)
			}
		}
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
			return
		}
	}
}
