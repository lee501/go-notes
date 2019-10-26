package main

import (
	"fmt"
	"testing"
)

/*
	Go中的字符串是一个只读的字节序列(底层为字节数组)
		正常情况下防止修改字节切片的元素导致字符串被修改
		*字符串转字节： 将字符串中的字节序列复制一份存储在目标字节切片中
		*字节转字符串： 将字节切片的底层元素序列复制一份存储在结果字符串中

*/
//避免相互转换的几种情况
//*1 for range循环中跟随range关键字： 从字符串到字符切片转换不会复制字符串中的字节序列
var s = "01234567890123456789"
var S = s + s
var nf, ng int

func f() {
	//复制底层字节数组
	bs := []byte(S)
	for _, b := range bs {
		nf += int(b)
	}
}

func g() {
	//不复制底层字节数组
	for _, b := range []byte(S) {
		ng += int(b)
	}
}

//*2 字节切片转换成字符串用作map的键时，不会复制自己的底层元素序列
var k = []byte("012345678901234567890123456789")
var m map[string]int

func c() {
	//复制底层字节序列
	key := string(k)
	fmt.Println(m[key])
}

func d() {
	//不复制底层字节序列
	fmt.Println(m[string(k)])
}

//3* 字节切片被用作字符串表达式比较值时， 不会复制字节切片的底层元素序列
func t(x, y []byte) bool {
	return string(x) == string(y)
}

//4* 字符串连接表达式： 如果有一个被连接的字符串值为非空字符串时，此字符串连接表达式从字节切片转换到字符串的时候不会开辟内存复制底层字节切片元素序列
var x = []byte{3, 'x'}
var y = []byte{5, 'x'}

func n() {
//	每个转换都要开辟一次内存
   s := string(x) + string(y)
   fmt.Println(s)
}

func v() {
	//不会单独开辟内存空间
	s := " " + string(x) + string(y)
	fmt.Println(s[1:])
}


func main() {
	//f与g的运行效率 f < g
	fmt.Println(testing.AllocsPerRun(1, f))
	fmt.Println(testing.AllocsPerRun(1, g))

	c()
	//c与d的运行效率 c < d
	fmt.Println(testing.AllocsPerRun(1, c))
	fmt.Println(testing.AllocsPerRun(1, d))

	//n与v的运行效率  n < v
	fmt.Println(testing.AllocsPerRun(1, n))
	fmt.Println(testing.AllocsPerRun(1, v))
}
