package main

import (
	"flag"
	"fmt"
	"go-notes/file-process"
	"log"
	"os"
	"runtime/pprof"
)
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		// 根据命令行指定文件名创建 profile 文件
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		// 开启 CPU profiling
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	s := file_process.ReadFileByString("/Users/lee/Documents/study.rb")
	fmt.Println(s)
}
