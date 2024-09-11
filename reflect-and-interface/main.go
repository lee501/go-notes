package main

import (
	"fmt"
	"reflect"
)

/*
	通过refelct动态操作结构体
	可以使用反射来查看或赋值结构体的值（赋值的时候结构体首字母必须大写）
	可以通过反射来判断变量的类型
	* 整个reflect包中最重要的两个类型
	  * reflect.Type 类型
	  * reflect.Value 值
	* 获取到Type和Value的函数
	  * reflect.TypeOf(interface{}) 返回Type
	  * reflect.ValueOf(interface{}) 返回值Value
*/
type Person struct {
	Id int64
	Name string `xml:"name"`
}

func main()  {
	//基本示例，查看变量类型和value
	a := 1.5
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))

	//获取结构体属性的值
	people := Person{21, "lee"}
	v := reflect.ValueOf(people)
	fmt.Println(v)
	//获取属性个数
	num := v.NumField()
	fmt.Println(num)
	//获取第0个属性,id,并转换为int64类型: v.FieldByIndex([]int{0})
	fmt.Println(v.Field(0).Int())
	//获取第1个属性,转换换为string类型
	fmt.Println(v.Field(1).String())

	//根据名字获取类型
	isValue := v.FieldByName("Id")
	fmt.Println(isValue.Kind().String())


	/*
		设置结构体属性的值，反射设置结构属性的时候，首字母必须大写
		传递的是结构体的指针
	*/
	peo := new(Person)
	// Elem()获取指针指向地址的封装, 才可以操作value
	v = reflect.ValueOf(peo).Elem()
	fmt.Println(v)
	if v.FieldByName("Id").CanSet() {
		v.FieldByName("Id").SetInt(123)
		v.FieldByName("Name").SetString("hello")
	}


	//通过反射来获取tags， 通过typeof
	t := reflect.TypeOf(Person{})
	name, _ := t.FieldByName("Name")
	fmt.Println(name)
	fmt.Println(name.Tag)
	//获取tag对应的内容
	fmt.Println(name.Tag.Get("xml"))


	type S struct {
		F string `species:"gopher" color:"blue"`
	}
	s := S{}
	st := reflect.TypeOf(s)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))
}
