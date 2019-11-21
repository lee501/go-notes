package main

import "fmt"

/*	考点
		注意map需要初始化才能使用
		指针不支持索引
*/
//定义map类型
type Param map[string]interface{}
//定义结构体，属性为Param指针。
type Show struct {
	*Param
}

func main() {
	s := new(Show)
	//初始化map
	p := make(Param)
	p["day"] = "21"
	s.Param = &p
	//指针不支持索引，需取值
	tmp := *s.Param
	fmt.Println(tmp["day"])
}

/*
	channel:
		单项channel不可关闭

	func Stop(stop <-chan bool) {
		close(stop) //错误示类
	}
*/

/*
	select机制
		* select机制用来处理异步IO问题
		* 每个case语句必须示一个IO操作
		* golang在语言级别上支持select关键字
*/
