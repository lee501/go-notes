package main

import "fmt"

/*
	按位置零操作符： &^
		z = x &^ y 表示如果y中的bit为1，则z对应的bit为0，否则z对应的bit等于x中的bit位的值
       	z = x | y 或操作符可以理解为y的bit为1，则z也为1， 否则z与x的bit位相同
*/
func main() {
	var a int8 = 3
	var b int8 = 5  //00001001
	z := a &^ b
	fmt.Printf("z: %08b\n", z)
	conversePosition()
}

/*
	按位取反之后返回一个每个 bit 位都取反的数，
		对于有符号的整数来说，是按照补码进行取反操作的（快速计算方法：对数 a 取反，结果为 -(a+1) ），
		对于无符号整数来说就是按位取反

	作为二元运算符，^ 表示按位异或，即：对应位相同为 0，相异为 1
*/

func conversePosition() {
	//0000 0011
	//1111 1100
	//1111 1101
	var a int8 = -3
	//0000 0011
	//0000 0011
	//0000 0011
	var x int8 = 3

	b := x^a //11111110 => 11111101 => 0000 0010
	fmt.Printf("%08b\n", a)
	fmt.Println(b)
}
