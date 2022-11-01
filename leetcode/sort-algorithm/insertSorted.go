package sort_algorithm

func insertSort(array []int) {
	if len(array) < 2 {
		return
	}
	for i := 1; i < len(array); i++ {
		key := array[i]
		for j := i; j > 0; j-- {
			if array[j-1] > key {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}

	}
}

func binarySearch(s []int, key int) int {
	begin := 0
	end := len(s) - 1
	for {
		mid := (end - begin) / 2
		if s[mid] == key {
			return mid
		} else if s[mid] < key {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

/*
func main() {
	arr := []int{6, 5, 4, 3, 2, 1}
	//i := binarySearch(arr, 3)
	insertSort(arr)
	fmt.Println(arr)
	//fmt.Println(binarySearch(arr, 3))
}
*/
