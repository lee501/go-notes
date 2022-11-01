package main

/*map的声明和初始化
	map的key 可以是bool 数字 string 指针 channel struct array等可以比较类型
       		 不能是slice map func
*/

type People struct {
	Name string
}

func mapDemo() {
	//1声明， 需要分配空间
	var peopleDB map[string]People
	peopleDB = make(map[string]People)
	peopleDB["1234"] = People{"lee"}
	//map的创建
	//使用make
	test := make(map[string]People)
	test["222"] = People{"test"}
	//创建并初始化
	peopleDB = map[string]People{
		"123": People{"anne"},
	}
}
