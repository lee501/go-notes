package main

import "fmt"

func main() {
	// 先声明map
	var m1 map[string]string
	// 再使用make函数创建一个非nil的map，nil map不能赋值
	m1 = make(map[string]string)
	// 最后给已声明的map赋值
	m1["a"] = "aa"
	m1["b"] = "bb"

	// 直接创建
	m2 := make(map[string]string)
	// 然后赋值
	m2["a"] = "aa"
	m2["b"] = "bb"

	// 初始化 + 赋值一体化
	m3 := map[string]string{
		"a": "aa",
		"b": "bb",
	}

	// ==========================================
	// 查找键值是否存在
	if v, ok := m1["a"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key Not Found")
	}

	// 遍历map
	for k, v := range m3 {
		fmt.Println(k, v)
	}
}
