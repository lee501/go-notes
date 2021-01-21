package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	bufferedWord := &BufferedWordReader{
		word: make(chan string),
	}
	go func() {
		bufferedWord.ReadFileByString("./test.txt")
	}()
	for v := range bufferedWord.word {
		fmt.Println(v)
	}
}

type BufferedWordReader struct {
	word   chan string
}

func (b *BufferedWordReader)ReadFileByString(filename string)  {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		sl, _, err := reader.ReadLine()
		if err == io.EOF {
			close(b.word)
			break
		}
		str := string(sl)
		split := strings.Split(str, " ")
		for _, v := range split {
			b.word <- v
		}
	}
}
