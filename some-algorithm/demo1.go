package main

import "fmt"

func main() {
	bslice := []byte{1, ' ', ' ', 2, ' ', ' ', 3, ' ', ' ', ' ', 4}
	//re := removeExtraSpace(bslice)
	xin(bslice)
}
//func strtrip(s []byte) string {
//	str := string(s)
//	reg := regexp.MustCompile("\\s+")
//	return reg.ReplaceAllString(str, " ")
//}
func test1() {
	nums := []int{1, 2, 3, 4}
	_ = append(nums[:2], 4)
	fmt.Println("test1:", nums)

	//nums changes because the cap is big enought, the original array is modified.

}

func removeExtraSpace(bs []byte) []byte {
	curl := 0
	var res []byte
	for i := 0; i < len(bs); i++ {
		if bs[i] == 32 {
			curl++
			if i == len(bs) - 1 && curl > 1 {
				res = append(res, bs[i])
			}
		} else {
			if curl == 0 {
				res = append(res, bs[:i+1]...)
			} else if curl == 1 {
				res = append(res, bs[i-1:i+1]...)
			} else {
				res = append(res, bs[i-1:i+1]...)
			}
			curl = 0
			continue
		}
	}
	return res
}

func xin(data []byte) int {
	l := len(data)
	index := 0
	for {
		if data[index] == 32 && data[index+1] == 32 && index <= l -1 {
			data = append(data[0:index], data[index+1:]...)
			l--
			continue
		}
		index++
		if index == l {
			break
		}
	}
	/*
	1 [1 32 32 2 32 32 3 32 32 32 4]
	3 [1 32 2 32 32 3 32 32 32 4]
	5 [1 32 2 32 3 32 32 32 4]
	5 [1 32 2 32 3 32 32 4]
	[1 32 2 32 3 32 4]
	[1 32 2 32 3 32 4 4 4 4 4]


	1 [1 32 32 2 32 32 3 32 32 32 4]
	3 [1 32 2 32 32 3 32 32 32 4]
	5 [1 32 2 32 3 32 32 32 4]
	[1 32 2 32 3 32 32 4]

	*/

	//for index := 0; index < len(data)-1; index++ {
	//	if data[index] == 32 && data[index+1] == 32{
	//		fmt.Println(index, data)
	//		data = append(data[0:index], data[index+1:]...)
	//	}
	//}
	fmt.Println(data)
	return len(data)
}