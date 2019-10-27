package file_process

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

//读取压缩文件
func ReadGzipFile(filename string) {
	var r *bufio.Reader
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fscanf(os.Stderr, "%v, can't open %s: error: %s\n", os.Args[0], filename, err)
		os.Exit(0)
	}
	fz, err := gzip.NewReader(file)
	if err != nil {
		r = bufio.NewReader(file)
	} else {
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			println("done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}
