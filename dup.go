package main

import "fmt"

//
//import "fmt"
//
//type People struct {
//	Age int
//	Sex string
//}
//
//type Man interface {
//	Name(string) string
//}
//
//func (*People) Name(str string) string {
//	return str
//}
//
//type Person struct {
//	name string
//	sex  byte
//	age  int
//}
//
////学生
//type Student struct {
//	Person // 匿名字段，那么默认Student就包含了Person的所有字段
//	id     int
//	addr   string
//	name   string //和Person中的name同名
//}
//
////func main() {
////	//counts := make(map[string]int)
////	//inputs := bufio.NewScanner(os.Stdin)
////	//
////	//for inputs.Scan() {
////	//	fmt.Println(counts[inputs.Text()])
////	//	counts[inputs.Text()] ++
////	//}
////
////	//s := []string{"a", "b"}
////	//fmt.Printf("%v\n", s[0:0])
////	//fmt.Println(^uint(0))
////	//const MaxScanTokenSize = 64 * 1024
////	//fmt.Printf("%T", MaxScanTokenSize)
////	//i := int(1)
////	//fmt.Println(unsafe.Sizeof(i)) // 4
////	//j := 1
////	//fmt.Println(unsafe.Sizeof(j)) // 4
////	//u := uint(1)
////	//fmt.Println(unsafe.Sizeof(u)) // 4
////	//str := []string{"a", "b", "c"}
////	//s := [][]string{{"a", "b", "c"}, {"e", "d", "f"} }
////	//for _, value := range s {
////	//	fmt.Println(value)
////	//}
////	//str := "abcd"
////	//for v := range []rune(str) {
////	//	fmt.Println(v)
////	//}
////	//str1 := []rune(str)
////	//for v := range str1 {
////	//	//输出效果跟[]rune的值一样
////	//	fmt.Println(v)
////	//}
////	//
////	//me := interface{}(&People{32, "man"})
////	//if m, ok := me.(People); ok {
////	//	fmt.Println(m)
////	//}
////	//
////	//type Car struct {
////	//	weight int
////	//	name   string
////	//}
////	//
////	//type Bike struct {
////	//	Car
////	//	lunzi int
////	//}
////	////time.After(time.Second * time.Duration(2))
////	//for i := 0; i < 3; i++ {
////	//	for j := 0; j < 5; j++ {
////	//		if j == 2 {
////	//			break
////	//		}
////	//		fmt.Println("hello")
////	//	}
////	//	fmt.Println("hi")
////	//}
////	//
////	//student := Student{Person{"lee", 0, 22}, 1, "beijing", "hh"}
////	for i := 0; i < 2; i++ {
////		createfunc()
////	}
////	var t map[] string
////}
//
//func main() {
//
//	//println(DeferFunc1(1))
//	println(DeferFunc2(1))
//	//println(DeferFunc3(1))
//	c := sum()
//	println(c)
//
//}
//
//func DeferFunc1(i int) (t int) {
//	t = i
//	defer func() {
//		t += 3
//	}()
//	return t
//}
//
//func DeferFunc2(i int) int {
//	t := i
//	defer func() {
//		t += 3
//	}()
//	return t
//}
//
//func DeferFunc3(i int) (t int) {
//	defer func() {
//		t += i
//	}()
//	return 2
//}
//
//func sum() int{
//	a := 2
//	b := 3
//	return a + b
//}
//
//func createfunc() {
//	i := 0
//	fmt.Printf("%v\n", &i)
//}
type Kind uint
const (
	Invalid Kind = iota
	Bool
	Int
	Int8
)

var(
	size = 1024
	max_size = size*2
)

type A struct {
	int
}

type B struct {
	A
}
func main() {
	println(Int)
	a := &A{2}
	b := &B{*a}
	println(b.int)
	b.int = 3
	println(a.int)
}
func a() int {
	var i = 2
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	fmt.Println("return:", i)
	return i
}