//string rune byte三者的关系
//uint8       the set of all unsigned  8-bit integers (0 to 255)
//byte        alias for uint8
/*
	1.byte: unit8,用于区分字节值和8位无符号整数值
	2.rune: int32, 用于区分字符值和整数值
	3.string：8位字节字符串的集合
			  底层是一个byte数组
			  字符串不能为nil值
			  字符串的值是不可变的
			  字符串可以判断向不向等

	字节的概念：
		1.计算机的存储的基本单位是byte(字节)， 存储信息的最小单位是位(bit 二进制0或1)，一个byte由8个二进制位组成
		2.ACSCII编码中一个英文占一个字节byte， 汉子占两个字节byte
		3.UTF-8编码 常见的汉子是3个字节
	从编码角度来分析：byte用来代表一个字节代表的数据，而不是整数（比如字符a就是97）
					rune用来表示unicode的码点，即一个字符
*/
//示例代码
package main

import "fmt"

//通过rune修改字符串
func StrChangeByRune(str *string, i int, ch rune) {
	temp := []rune(*str)
	//更改下标i的元素
	temp[i] = ch
	//转换成字符串
	*str = string(temp)
}
//通过byte修改字符串
func StrChangeByByte(str *string, i int, ch byte) {
	//转换为字节
	temp := []byte(*str)
	//更改下标i的元素
	temp[i] = ch
	//转换成字符串
	*str = string(temp)
}
func main() {
	s := "GO编程"
	//字符串的长度为字节的长度，因为string的底层是[]byte, len(s)值为8，一个中文占了三个字节
	fmt.Println(len(s))
	//与[]byte的结果是一致的
	fmt.Println(len([]byte(s)))
	//查看字符串长度，转换成[]rune, len值是4
	fmt.Println(len([]rune(s)))

	//通过遍历来查看string []byte []rune 值为代表该字节或字符的整数
	//for _, v := range []byte(s) {
	//	//输出的是8位的字节值
	//	fmt.Println(v)
	//}
	for _, v := range []rune(s) {
		//fmt.Println(string([]rune(s)[i]))
		fmt.Println(string(v))
	}
	for _, v := range s {
		//输出效果跟[]rune的值一样
		fmt.Println(string(v))
	}
	for i := 0; i < len(s); i++ {
		//输出结果跟[]byte(s)
		fmt.Println(s[i])
	}

	//测试SteChangeByByte和StrChangeByRune
	str1 := "快乐的学Go"
	str2 := "快乐的学Go"
	//出现乱码，一个中文是三个字节；快
	StrChangeByByte(&str1, 1, 'A')
	//乐被A替换掉了
	StrChangeByRune(&str2, 1, 'A')
	fmt.Println(str1)
	fmt.Println(str2)
}