//字符串的链接方式
package main

import (
	"bytes"
	"fmt"
	"strings"
)

//使用+号拼接
func AddStringWithOperator() string{
	hello := "hello"
	world := "world"
	result := hello + world
	return result
}
//使用fmt.Sprintf拼接，包含数字之类的时候使用此方法
func ConnectStringWithNum(first string, last int) string{
	first = first
	last = last
	result := fmt.Sprintf("%s d%", first, last)
	return result
}

//string.Join
func JoinString(first, last string) string{
	first = first
	last = last
	return strings.Join([]string{first, last}, "")
}

//buffer.WriteString()
func ConnectStringWithBuffer(first, last string) string{
	var buffer bytes.Buffer
	buffer.WriteString(first)
	buffer.WriteString(last)
	return buffer.String()
}
func main() {
	name := "lee"
	age := 34
	result := ConnectStringWithNum(name, age)
	fmt.Println(result)
}
