package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func readByLine(filepath string) {
	start := time.Now()
	file, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}
	fmt.Println("readEachLine spend : ", time.Now().Sub(start))
}
