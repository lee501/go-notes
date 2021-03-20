package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//常量
	const (
		Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
		Ltime                         // the time in the local time zone: 01:23:23
		Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
		Llongfile  	  = 6                 // full file name and line number: /a/b/c/d.go:23
		Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
		LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
		LstdFlags     = Ldate | Ltime // initial values for the standard logger
	)
	fmt.Println( Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC, LstdFlags)
	const (
		a = 1
		b
		c
	)
	fmt.Println(a,b,c)
	//log
	/*
		log.Println()
		func Println(v ...interface{}) {
			std.Output(2, fmt.Sprintln(v...))
		}
	*/
	//打印日志到文件
	file, _ := os.OpenFile("./logger/logger.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
	defer func() {
		file.Close()
	}()
	logger := log.New(file, "[info]\n", log.Ltime)
	logger.Println("输出日志信息")
}
