package main

/*
	for-range的解析
		1. 遍历一个容器时候，遍历的是此容器的一个副本
			*当容器为数组时，此数组的所有元素也被复制
			*复制一个切片(或数组的指针)的时候，复制的是地址，切片和数组的元素并不会被复制，而是与原切片（或数组）共享元素
		2. 循环变量是被遍历的容器的每个键值(或索引)和元素对的副本， 如果元素为结构体值，这此副本的字段和容器的字段是两个不同的值
*/
import "fmt"

type T struct {
	n int
}

func main() {
	ts := [2]T{}
	//遍历数组
	for i, t := range ts {
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, "  ")
		}
	}
	fmt.Println(ts) //0 [{0} {9}]
	fmt.Println("------------------")

	//遍历数组指针
	tl := [2]T{}
	for i, t := range &tl {
		switch i {
		case 0:
			t.n = 3
			tl[1].n = 9
		case 1:
			fmt.Print(t.n, "  ")
		}
	}
	fmt.Println(tl) //9  [{0} {9}]
	fmt.Println("------------------")

	//遍历切片
	tm := [2]T{}
	for i := range tm[:] {
		t := &tm[i]
		switch i {
		case 0:
			t.n = 3
			tm[1].n = 9
		case 1:
			fmt.Print(t.n, "  ")
		}
	}
	fmt.Println(tm)
	fmt.Println("------------------")
}
