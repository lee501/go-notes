package main

import "fmt"

/*
	考点：
		nil的赋值问题
			nil只能赋值给指针, chan, map, slice， interface和func类型的变量
			error是内置接口类型，也可以赋值nil
*/

var m interface{} = nil

/*
	考点： init()函数
		一个包中可以有多个init 函数
		不同包的init函数是根据导入的依赖关系来执行的，如A import B, B import C, 此时执行顺序为C， B， A
		一个包被多次引用， init只会执行一次
*/

/*
	考点： 类型选择
		类型选择语法 i.(type),  i只能为interface
*/

func GetValue() interface{} {
	return 1
}

func main() {
	v := GetValue()
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case interface{}:
		fmt.Println("interface")
	default:
		fmt.Println("unknown")
	}
}
