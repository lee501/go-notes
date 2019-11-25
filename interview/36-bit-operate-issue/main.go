package main

import "fmt"

/*
	按位置零操作符： &^
		z = x &^ y 表示如果y中的bit为1，则z对应的bit为0，否则z对应的bit等于x中的bit位的值
       	z = x | y 或操作符可以理解为y的bit为1，则z也为1， 否则z与x的bit位相同
*/
func main() {
	var a int8 = 3
	var b int8 = 5
	z := a &^ b
	fmt.Printf("z: %08b\n", z)
}
