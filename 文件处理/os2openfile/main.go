package main

import (
	"fmt"
	"os"
)

//使用os包中内容进行操作系统文件或目录
/*
	文件结构表示操作系统文件或文件夹
	type File struct {
		*file
	}

	type file struct {
		pfd poll.FD
		name string
		dirinfo *dirInfo
	}

	type FileMode unit32

	ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice

	*FIleInfo是一个interface表示文件的信息
	type FileInfo interface {
		Name() string //文件名字
		Size() int64 文件大小
		Mode() FileMode 文件mode
		ModTime() time.Time 修改时间
		IsDir() bool 是否是路径
		Sys() interface{}
	}
*/

func main()  {
	//创建文件夹
	//Mkdir要求创建的文件夹必须不存在，且父级目录必须存在
	err := os.Mkdir("/Users/richctrl/workspace/os", os.ModeDir)
	if err != nil {
		fmt.Println("创建目录失败", err)
		return
	}
	//使用MkdirAll创建文件夹，如果文件夹存在，不报错
	err = os.MkdirAll("/Users/richctrl/workspace/os", os.ModeDir)

	//使用Create创建文件： 要求文件目录必须存在，文件已存在的时候，会创建空文件覆盖之前的文件
	file, err := os.Create("/Users/richctrl/workspace/os/os.txt")
	if err != nil {
		fmt.Println("文件创建失败", err)
		return
	}
	fmt.Println(file.Name())

	//重命名文件和重命名文件夹os.Rename

	//打开文件
	f, err := os.Open("/Users/richctrl/workspace/os/os.txt")
	if err != nil {
		fmt.Println("打开文件失败", err)
	}
	defer f.Close()
	//获取文件信息
	fileInfo, err := f.Stat()
	if err != nil {
		fmt.Println("获取文件信息失败", err)
		return
	}
	fmt.Println(fileInfo.Name())    //文件名
	fmt.Println(fileInfo.IsDir())   //是否是文件夹,返回bool,true表示文件夹,false表示文件
	fmt.Println(fileInfo.Mode())    //文件权限
	fmt.Println(fileInfo.ModTime()) //修改时间
	fmt.Println(fileInfo.Size())    //文件大小

	//删除文件或文件夹
	/*
	os.Remove
	删除的内容只能是一个文件或空文件夹且必须存在
	*/
	/*
	os.RemoveAll
	只要文件夹存在,删除文件夹.
	无论文件夹是否有内容都会删除
	如果删除目标是文件,则删除文件
	*/
}