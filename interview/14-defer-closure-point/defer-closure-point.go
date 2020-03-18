package main

import "fmt"

/*
	考点：defer 闭包 指针变量
*/

type Person struct {
	age int
}

func main() {
	//p是一个指针变量
	p := &Person{25}

	//此处的defer中p将当前age 25作为参数，缓存到栈中
	defer fmt.Println(p.age)

	//此处将p的引用地址，age最终被重新赋值
	defer func(p *Person) {
		fmt.Println(p.age)
	}(p)

	//闭包 最终引用的是外部变量
	defer func() {
		fmt.Println(p.age)
	}()

	p.age = 27
	fmt.Println(p)
}
