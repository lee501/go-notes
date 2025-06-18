package main

import (
	"fmt"
	"strings"
)

func main() {
	//worker := []int{1, 2, 3, 4, 5}
	//
	//for i := range worker {
	//	go func() {
	//		fmt.Println(i)
	//	}()
	//}
	//var ch chan int
	//ch <- 1
	//var buffer bytes.Buffer
	//buffer.WriteString("123")
	//buffer.WriteByte('\n')
	//_ = buffer.UnreadByte()
	//fmt.Println(len(buffer.Bytes()), buffer.Bytes()[3])
	//输出 4 10
	//var payloadSeparator = "\n🐵🙈🙉\n"
	//buffer.WriteString(payloadSeparator)
	//buffer.WriteByte('\n')
	////buffer.WriteString("\n")
	////fmt.Printf("%v", payloadSeparator)
	//for _, b := range buffer.Bytes() {
	//	fmt.Printf("%c\n", b)
	//}
	//fmt.Println(buffer.Len())

	var strArr []string
	val := strings.Join(strArr, ", ")
	fmt.Println(val)
}
