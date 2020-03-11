package main

//switch用法：用于类型判断

func convert(i interface{}) {
	switch t := i.(type) {
	case int:
		println("i is int", t)
	case string:
		println("i is string", t)
	default:
		println("type not found")
	}
}

func main() {
	var i interface{}
	i = 100
	convert(i)
	i = "abc"
	convert(i)
	i = []int{1,2,3,4}
	convert(i)
}
