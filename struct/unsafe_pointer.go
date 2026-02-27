package Struct

import (
	"fmt"
	"unsafe"
)

//
//func main() {
//	s := "123"
//	b := ([]byte)(s)
//	change(s)
//	fmt.Println(b)
//	ss := "hello"
//	fmt.Println(*(*string)(unsafe.Pointer(&ss)))
//}
//

//func change(s string){
//	s = "456"
//}
/*
	首先，要理解go的string类型是什么

	通过reflect包，我们可以知道，在Golang底层，string其实是struct：

	type StringHeader struct { Data uintptr Len int }

	其实一个string是一个保存了真实字符数组指针和长度的结构体

	对于fun1 ， a是一个int变量，unsafe.Pointer(&a) 是一个原生指针，指向的内存保存了一个int

	c := (*string)(unsafe.Pointer(&a)) 将这个原生指针转换为了指向string的指针

	注意c 现在是指向string的指针，换句话说，c指向的内存被当做了一个StringHeader 结构体来处理

	但是c指向的内存，只分配了一个int的长度，并且这个int的值是2

	所以目前c中 Data 值为2， Len是无效的，因为Len现在已经溢出了有效内存范围，其内的值是随机的(该溢出区域内存的原始值)

	这时对 *c 进行赋值 "44", "44"是一个字符串常量，保存在只读的常量存储区中。

	赋值时 其实是把常量字符串"44"的指针赋给了c的Data字段，把其长度2赋给了c的Len字段。 注意这里是一个野指针赋值，因为 Len 字段已经溢出到有效内存之外了，但是这里panic的概率很低，因为内存padding机制

	然后输出 fmt.Println(*c) 这里，因为c的Len字段是溢出到无效内存中，所以如果其他指令有变量定义的操作，很可能把Len字段给覆盖掉 一旦Len被覆盖，那么其中保存的长度就不在是2，这时输出c，很大概率就会panic。 如果运气好，Len没有被覆盖，那么输出c，就能正常输出44

	对于fun2，可以根据上述原理自行分析。因为a变量也是一个StringHeader，c其实与unsafe.Pointer(&a)指向同一个StringHeader。

	对于fun3, 因为c本来就是一个新创建的string变量，也就是一个独立的StringHeader，没有复用变量a的内存，所以是正常的赋值和输出
*/
func demoUnsafePointer() {
	//num := 5
	//numPointer := &num
	//
	//flnum := (*string)(unsafe.Pointer(numPointer))
	//fmt.Println(*flnum)
	fun1()
	fun2()
	fun3()
	//var m *(*string)
	//fmt.Printf("%p", m)
	//fmt.Printf("%p", *m)
	//var s *string
	//s =
	//fmt.Printf("%#v", s)
	a := 1
	b := a
	var c *int
	c = &a
	fmt.Printf("%#v, %#v, %#v", a, b, *c)
}

func fun1() {
	a := 2
	c := (*string)(unsafe.Pointer(&a))
	*c = "2222"
	fmt.Println(*c)
}

func fun2() {
	a := "654"
	c := (*string)(unsafe.Pointer(&a))
	*c = "44"
	fmt.Println(*c)
}
func fun3() {
	a := 3
	c := *(*string)(unsafe.Pointer(&a))
	c = "445"
	fmt.Println(c)
}
