package main

import "fmt"

type Student struct {
	Name string
}


func main() {
	m := map[string]*Student{"stu": &Student{Name: "lee"}}
	m["stu"].Name = "anne"
	fmt.Println(m["stu"].Name)
}

//const (
//	a   = 1                  //1 (iota=0)
//	b                        //1 (iota=1，同上一行，相当于写b=1)
//	c   = b + iota           //3 (iota=2，b=1)
//	d                        //4 (iota=3，同上一行，相当于写b+iota，b=1)
//	e                        //5 (iota=4，同上一行，相当于写b+iota，b=1)
//	f   = "last one but one" //  (iota=5)
//	end               //6 (iota=6)
//)
//
//func main() {
//	fmt.Println(a, reflect.TypeOf(a))
//	fmt.Println(b, reflect.TypeOf(b))
//	fmt.Println(c, reflect.TypeOf(c))
//	fmt.Println(d, reflect.TypeOf(d))
//	fmt.Println(e, reflect.TypeOf(e))
//	fmt.Println(f, reflect.TypeOf(f))
//	fmt.Println(end, reflect.TypeOf(end))
//}
