package golang文件处理

import (
	"bufio"
	"io"
	"os"
)

//一次性加载到内存中适用于小文件
func ReadAllIntoMemory(filename string) (content []byte, err error) {
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	fileInfo, err := fp.Stat()
	if err != nil {
		return nil, err
	}
	//分配一个文件大小的空间
	buffer := make([]byte, fileInfo.Size())
	_, err = fp.Read(buffer) // 文件内容读取到buffer中
	if err != nil {
		return nil, err
	}
	return buffer, nil
}

//即给一个缓冲, 分多次读到缓冲中
func readByBlock(filename string) (content []byte, err error) {
	//获取文件指针
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	const bufferSize = 64 // 缓冲大小, 每次读取64个字节
	buffer := make([]byte, bufferSize)

	for {
		bytesRead, err := fp.Read(buffer)
		content = append(content, buffer[:bytesRead]...)
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			} else {
				return nil, err
			}
		}
	}
	return
}

//按行读取， 一行是一个[]byte, 多行就是[][]byte
func readBYLine(filename string) (lines [][]byte, err error) {
	fp, err := os.Open(filename) // 获取文件指针
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	bufReader := bufio.NewReader(fp)

	for {
		line, _, err := bufReader.ReadLine()
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}
		} else {
			lines = append(lines, line)
		}
	}
	return
}
