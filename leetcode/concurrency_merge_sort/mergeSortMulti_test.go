package concurrency_merge_sort

import "testing"

var array []int

func init() {
	for i := 0; i < 100000; i++ {
		array = append(array, i)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MergeSort(array)
	}
}

func BenchmarkMergeSortMulti(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MergeSortMulti(array)
	}
}
