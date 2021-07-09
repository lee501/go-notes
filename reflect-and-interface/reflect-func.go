package main

import (
	"fmt"
	socket "github.com/googollee/go-socket.io"
	"reflect"
)

func main() {
	c := event
	fv := reflect.ValueOf(c)
	ft := fv.Type()
	fmt.Println(ft.NumIn())
	for i := 0; i < ft.NumIn(); i++ {
		fmt.Println(ft.In(i))
	}

	var num float64 = 2.3435

	// 指针 *float64
	pointer := reflect.ValueOf(&num)
	// float64
	value := reflect.ValueOf(num)

	/* 强制转换
	   区分指针和值
	   转换类型要完全符合，否者直接panic
	*/
	convertPointer := pointer.Interface().(*float64) //#=> 一个地址
	convertValue := value.Interface().(float64)      //#=>2。3435
	fmt.Println(convertPointer, convertValue, value.Interface())
}

func event(conn socket.Conn, msg string) {

}
