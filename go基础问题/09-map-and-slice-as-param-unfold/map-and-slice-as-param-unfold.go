package main

import "fmt"

/*
	考点：map中不存在key时，返回该元素类型的零值
		delete map中不存在的键值对时， 不会报错
		map的 key: 可以是很多种类型，比如 bool, 数字，string, 指针, channel ,
					还有只包含前面几个类型的 interface types, structs, arrays

					显然，slice， map 还有 function 是不可以了，因为这几个没法用 == 来判断
*/
type person struct {
	name string
}
func mapIssue() {
	var m map[person]int
	p := person{name: "lee"}
	delete(m, p)
	fmt.Println(m[p])
	//m[p] = 1 panic: assignment to entry in nil map
}

/*
	考点： 可变函数
		通过可变函数操作符..., 切片可以展开，多参数在函数中可以转成切片
*/
func params(nums ...int) {
	nums[0] = 1
}

func main() {
	//输出0值
	mapIssue()
	//通过可变函数来修改切片元素值
	i := []int{0,2,3}
	params(i...)
	fmt.Println(i)
}
