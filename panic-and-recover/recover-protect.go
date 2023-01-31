package main

import (
	"fmt"
	"runtime"
)

type panicContext struct {
	context string
}

func ProtectPanic() {
	fmt.Println("开始测试")
	ProtectRun(func() {
		fmt.Println("手动panic前")
		panic(panicContext{"手动触发"})
		fmt.Println("手动panic后")
	})

	//空指针
	ProtectRun(func() {
		fmt.Println("空指针赋值前")
		var a *int
		*a = 1
		fmt.Println("空指针赋值完成")
	})
}

func ProtectRun(entry func()) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Printf("runtime error %s\n", err)
		default:
			fmt.Println("other err:", err)
		}
	}()
	entry()
}
