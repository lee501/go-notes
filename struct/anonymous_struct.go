//匿名字段： 只提供类型，而不写字段名的方式

//struct中，无论使用的是指针的方式声明还是普通方式，访问其成员都使用"."，
//在访问的时候编译器会自动把 stu2.name 转为 (*stu2).name。

// 当匿名字段是一个结构体的时候，
// 那么这个结构体所拥有的全部字段都被隐式地引入了当前定义的这个结构体，
package main

//人
type Person1 struct {
	name string
	sex  byte
	age  int
}
//学生
type Student2 struct {
	Person1 // 匿名字段，那么默认Student就包含了Person的所有字段
	id     int
	addr   string
	name string // 同名字段
}

//非结构体匿名字段
//所有的内置类型和自定义类型都是可以作为匿名字段的
type mystr string //自定义类型
type Teacher struct {
	int   //匿名字段，内置类型
	mystr //匿名字段，自定义类型
}

//结构体指针
//注意成员使用.访问的时候，go是转换成指针后访问的： stu.name 转换为(*stu).name
type Student3 struct {
	*Person1
	id int
}
func main() {
	//初始化
	s1 := Student2{Person1{"mike", 'm', 18}, 1, "sz", "ss"} //顺序初始化，注意Person的初始化方法

	s2 := Student2{Person1: Person1{name: "lee"}, id: 1}  //部分成员初识化

	//成员操作
	var s3 Student2
	s3.name = "lee" //同名字段默认给外层成员赋值
	s3.Person1.name = "anne" //同名字段给内层成员赋值
	s3.sex = 0 //等同于s3.Person.sex
	//Person为结构体，是一个匿名字段
	s3.Person1 = Person1{"lee", 1, 22}

	//非结构体匿名字段的赋值
	t := Teacher{int: 1, mystr: "hello"}

	//	结构体指针的初始化
	s4 := Student3{&Person1{"mike", 1, 18}, 1}
	// 声明变量方式
	var s5 Student3
	s5.Person1 = new(Person1) //new 一个空间
	s5.Person1.name = "lee" //等同于s5.name
}
