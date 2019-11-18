package main

import "fmt"

/*
	类型实现String方法，当格式化输出时会自动使用String方法
		如果在String方法中使用格式化输出，会导致递归调用，最终抛出错误
*/

type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	return c.Daemon
}

func main() {
	c := &ConfigOne{"test"}
	fmt.Println(c)
}
