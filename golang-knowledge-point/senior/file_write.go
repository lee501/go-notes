package main

import (
	"bufio"
	"fmt"
	"os"
)
//写入文件
func writeIntoFile(filename string) {
	outfile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outfile.Close()

	outWriter := bufio.NewWriter(outfile)

	for i := 0; i < 10; i++ {
		outWriter.WriteString("Hello golang")
	}
	outWriter.Flush()
}
