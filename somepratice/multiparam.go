package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

//func main() {
//	str := []string{"a","b","c"} //[]string 不可以转换成[]interface{}
//	multi(str...)
//	//var str []int
//	//str[0] = 1
//	//fmt.Println(str)
//	exec.Command()
//}
//
//func multi(s ...interface{}) {
//	for _, v := range s {
//		fmt.Println(v)
//	}
//}

func main() {
	cmd := exec.Command("/bin/bash", "-c", `df -lh`)

	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return
	}

	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return
	}

	//使用带缓冲的读取器
	outputBuf := bufio.NewReader(stdout)

	for {

		//一次获取一行,_ 获取当前行是否被读完
		output, _, err := outputBuf.ReadLine()
		if err != nil {

			// 判断是否到文件的结尾了否则出错
			if err.Error() != "EOF" {
				fmt.Printf("Error :%s\n", err)
			}
			return
		}
		fmt.Printf("%s\n", string(output))
	}

	//wait 方法会一直阻塞到其所属的命令完全运行结束为止
	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return
	}
}

