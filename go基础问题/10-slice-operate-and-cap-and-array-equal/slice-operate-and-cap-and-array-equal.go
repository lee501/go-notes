package main

import "fmt"

/*
	考点：切片操作
		操作符 [i:j], 若原数组的长度为l， 截取后的切片大小和容量分别为 j-i, l-i
		操作符还有第三个参数[i:j:k], k用来限制新切片的容量，长度：j-i, 容量：k-i

*/

func sliceIssue() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:5]
	fmt.Println(t[0], len(t), cap(t))
}

/*
	考点：数组判断相等问题
		数组为值类型可以比较，同时数组的长度也是数组类型的组成部分，只有长度相等才可以比较
*/
func arrayIssue() {
	a := [2]int{5, 6}
	b := [2]int{5, 6}
	if a == b {
		fmt.Println("equal")
		return
	}
	fmt.Println("not equal")
}

/*
	考点：
		interface等于nil问题
		接口类型变量都包含一个type value对，(value, type)
			例如var r io.Reader
			file, err := os.OpenFile("dev/test", os.O_RDWR, 0)
			r = file
			此时r的pair对(file, *os.File)
		只有当type 和 value都为nil的时候，interface才为nil
*/
func interfaceIssue() {
	//声明一个interface的变量，此时=nil
	var i interface{}
	if i == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println("not nil")
}

/*
	考点：cap函数适用的；类型
		适用array， slice, channel

*/

func main() {
	sliceIssue()
	arrayIssue()

	interfaceIssue()
	/*
		切片操作符 [low,high]
			规则0 <= low <= high <= cap(原切片)
	*/
	s := make([]int, 3, 9)
	fmt.Println(len(s), cap(s))

	s1 := s[4:8]
	fmt.Println(len(s1), cap(s1))
}
