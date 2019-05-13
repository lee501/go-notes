package main

import (
	"fmt"
	"strings"
)

/*
	GO语言中，函数与int string等一样为值类型， 可以声明一个值类型为某个函数的变量（函数变量）
		基本用法函数变量被直接当作值进行传递
*/
func add(a, b int) (sum int) {
	sum = a + b
	return
}

func simple() {
	var number = 3
	str := "abc"
	fmt.Println(number, str)
	//声明一个函数类型的变量
	myFunc := add
	fmt.Println(myFunc(1,2 ))
}

/*
	另外一种使用方法， 使用type关键字定义一个底层类型为函数类型的自定义类型
	函数被当作值赋值给这种变量时（函数的签名（参数和返回值）必须相同）
*/
type ProcessBaseName func(string) string
//base函数 移除字符串的路径部分和.后缀 只取文件名
func basename(str string) string {
	lash := strings.LastIndex(str, "/")
	str = str[lash+1:]
	if dest := strings.LastIndex(str, "."); dest > 0 {
		str = str[:dest]
	}
	return str
}
//使用函数变量作为参数
func GetBaseName(str string, BaseNameFunc func(string)string) string {
	return BaseNameFunc(str)
}
//使用type自定义类型
func GetBaseNameType(str string, BaseNameFunc ProcessBaseName) string {
	return BaseNameFunc(str)
}

func main() {
	str := "root/websocker/test.go"
	//函数签名是指 函数的参数和返回值
	GetBaseName(str, basename)
	GetBaseNameType(str, basename)
}