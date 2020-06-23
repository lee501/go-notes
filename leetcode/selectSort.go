package main

func main() {
	
}

func selectSort(array []int) {
	if len(array) < 2 {
		return
	}
	for i:=0; i < len(array); i++ {
		min := i
		for j := i+1; j < len(array); j++ {
			if array[min] > array[j] {
				min = j
			}
		}
		if min != i{
			array[min], array[i] = array[i], array[min]
		}
	}
}
