package main

import (
	"fmt"
	"reflect"
)

//interface与nil判等的问题
//interface的内部包含type和value
//interface==nil， 仅当type和value均为nil（nil, nil）

//检查以下题目，type有值，value==nil
type Animals interface {
	Meow()
}

type Catty struct {
}

func (*Catty) Meow() {
	fmt.Println("meow")
}

func GetAnimal() Animals {
	var catty *Catty
	catty = &Catty{}
	return catty
}

func demointerface() {
	var cat = GetAnimal()
	m := reflect.TypeOf(cat)
	fmt.Println(m)
	kind := m.Kind()
	fmt.Println(kind)
}
