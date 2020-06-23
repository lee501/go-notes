package main

import "fmt"

func main() {
	//var n int
	//step  := 0
	//fmt.Println("请输入n值：")
	//fmt.Scan(&n)
	//n, step = getStep(n, step)
	//fmt.Println(n, step)
	var a,b int
	fmt.Scanf("%d%d", &a, &b)
	fmt.Println(a, b)
}

func getStep(n, step int) (int, int) {
	for n != 1{
		if n % 2 == 0 {
			n = n / 2
		} else {
			n = (3 * n + 1) / 2
		}
		step++
	}
	return n, step
}