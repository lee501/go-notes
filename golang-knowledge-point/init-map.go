package main

/*map的声明和初始化*/

type People struct {
	Name string
}

func main() {
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
