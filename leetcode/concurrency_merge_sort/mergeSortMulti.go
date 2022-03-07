package concurrency_merge_sort

import "sync"

/*
	how correct to process concurrency merge sort
	solve too much concurrency slow issue
*/
func MergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	mid := len(s) / 2

	var l []int
	var r []int

	l = MergeSort(s[:mid])
	r = MergeSort(s[mid:])

	return merge(l, r)
}

func merge(l, r []int) []int {
	res := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(res, r...)
		}
		if len(r) == 0 {
			return append(res, l...)
		}
		if l[0] <= r[0] {
			res = append(res, l[0])
			l = l[1:]
		} else {
			res = append(res, r[0])
			r = r[1:]
		}
	}
	return res
}

//并发容器
var sem = make(chan struct{}, 100)

//concurrency merge
func MergeSortMulti(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	mid := len(s) / 2

	//wait group
	wg := sync.WaitGroup{}
	wg.Add(2)

	var l []int
	var r []int

	select {
	case sem <- struct{}{}:
		go func() {
			l = MergeSortMulti(s[:mid])
			<-sem
			wg.Done()
		}()
	default:
		l = MergeSort(s[:mid])
		wg.Done()
	}

	select {
	case sem <- struct{}{}:
		go func() {
			r = MergeSortMulti(s[mid:])
			<-sem
			wg.Done()
		}()
	default:
		l = MergeSort(s[mid:])
		wg.Done()
	}

	wg.Wait()
	return merge(l, r)
}
