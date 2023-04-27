package main

import (
	"io"
	"os"
	"reflect"
)

/*
	go语言中每个变量都有唯一个静态类型。
	interface的结构包含类型(type)和数据值(value):
		无函数eface（interface{}）:
				type --> type类型对象
				value --> data
		有函数iface
				type--> 静态类型  --> 静态类型
						动态混合类型 --> 动态混合类型
						方法集   -->  函数列表
				value --> data
*/
//interface例子
func inter() {
	/*
		r --> type --> 静态类型--> io.Reader
	 			   --> 动态混合类型 --> nil
					--> 方法集  --> Read函数
			  value --> nil
	*/
	var r io.Reader
	file, _ := os.OpenFile("/path", os.O_RDWR, 0)
	/*
		r --> type --> 动态混合类型 --> *os.File
 		  --> value --> file
	*/
	r = file
	_= r
	/*
		empty --> type --> nil
	  		  --> value --> nil
	*/
	var empty interface{}
	/*
		empty --> type --> *os.File
			  --> value --> file
	*/
	empty = file
	_ = empty
}

/*
	reflect三大法则：
		1接口数据(interface) ==> 反射对象
				interface --> type <reflect.TypeOf> --> reflect.Type
						  --> value <reflect.ValueOf> --> reflect.Value
				reflect.Value --> <.Type()> reflect.Type
		2反射对象 ===> 接口数据
				reflect.Value  <Interface().(强制类型)> --> interface
		3通过反射对象修改数据
				反射对象需要传递的是指针
*/
func changeValueByreflect() {
	var x float64 = 3.4
	v := reflect.ValueOf(&x) //此处需要传递指针， 反射的是原数据类型的指针

	e := v.Elem() //通过Elem获取当前数据的值类型
	if e.CanSet() {
		e.SetFloat(6.18)
	}
}

