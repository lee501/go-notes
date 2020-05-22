package main

import (
	"fmt"
)

type people struct {
	name string
}
func main() {
	//arr := []int{1,2,3,4}
	//fmt.Println(arr[:2])
	//fmt.Println(arr[2:])
	//a :=[...]int{1,2,3}
	//m := reflect.TypeOf(a)
	////fmt.Printf("%v", m)
	//var p *people
	//fmt.Println(p == nil)
	//pe := people{"lee"}
	//fmt.Println(reflect.ValueOf(pe).FieldByName("name"))
	//str := "abc"
	//fmt.Println(str[0])
	//for _,v := range str{
	//	fmt.Println(v)
	//}

	arr := []int{22, 18, 30, 10, 9}
	fmt.Println(arr)
	quickSort(arr, 0, 4)

	i := BinarySearch(arr, 18)
	fmt.Println(i)

	m := []int{2, 5, 3, 6, 8, 8, 7, 5}
	fmt.Println(findNumSumInArray(m, 10))

	var mutiSlice = make([][]int, 2)
	for i := range mutiSlice {
		mutiSlice[i] = make([]int, 2)
	}
	mutiSlice[0][0] = 1
	fmt.Println(mutiSlice)

	s := make([]int, 10)
	s[0] = 1
	fmt.Println(s)
}

type M interface {
	N
}

type N interface {
	Check()
}

func (p *people) Check() {
	fmt.Println("interface")
}

//[]{22, 18, 30, 10, 9}
func quickSort(arr []int, begin, end int) {
	if len(arr) < 0 || begin > end {
		return
	}
	low, high := begin, end
	temp := arr[begin]
	for low < high {
		for low < high {
			if temp < arr[high] {
				high --
				continue
			}
			arr[high], arr[low] = arr[low], arr[high]
			//[]{9, 18, 30, 10, 22} high = 4 low = 0 temp = 22
			//[]{9, 18, 10, 22, 30} high = 3 low = 2 temp = 22
			break
		}

		for low < high {
			if temp > arr[low] {
				low++
				continue
			}
			arr[high], arr[low] = arr[low], arr[high]
			//[]{9, 18, 22, 10, 30} high = 4 low = 2  temp = 22
			//[]{9, 18, 10, 22, 30} high = 3 low = 3 temp = 22
			break
		}
	}
	//[]{9, 18, 10, 22, 30} high = 3 low = 3 temp = 22
	quickSort(arr, begin, low - 1)
	quickSort(arr, high + 1, end)
}

func BinarySearch(array []int, target int) int {
  	if len(array) < 1 {
  		return -1
	}
  	i, j := 0, len(array) - 1
  	for {
  		mid := (j - i) / 2
  		if target > array[mid] {
  			i = mid + 1
		} else if target < array[mid] {
			j = mid - 1
		} else {
			return mid
		}
	}
}

func findNumSumInArray(array []int, sum int) [][]int {
	m := make(map[int]int)
	l := make([][]int, 0)
	for i, v := range array {
		if j, ok := m[sum -v]; ok {
			l = append(l, []int{i, j})
		}
		m[v] = i
	}
	return l
}