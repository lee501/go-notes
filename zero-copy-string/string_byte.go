package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//通过unsafe指针运算 改变struct field的值
func ptrDemo() {
	x := struct {
		a int
		b int
	}{a: 1, b: 2}
	//获取x.b的地址
	fmt.Println(&x.b)
	fmt.Println((*int)(unsafe.Pointer((uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b))))) //(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	//通过指针修改b的值
	*(*int)(unsafe.Pointer((uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))) = 4
	fmt.Println(x.b)
}

/*
	string和[]byte 在运行时的类型表示为reflect.StringHeader和reflect.SliceHeader
		type StringHeader struct {
			Data uintptr
			Len  int
		}

		type SliceHeader struct {
			Data uintptr
			Len  int
			Cap  int
		}
*/

//string 和 []byte 零拷贝转换
func string2bytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func bytes2string(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	sh := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}

	return *(*string)(unsafe.Pointer(&sh))
}

func main() {
	ptrDemo()
	s := "Hello golang"
	b := string2bytes(s)
	fmt.Println(b)
	s = bytes2string(b)
	fmt.Println(s)
}
