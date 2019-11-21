package main

//import "fmt"
////
////func main() {
////	//i := 1
////	//s := []string{"A", "B", "C"}
////	//s[i-1], i = "Z", 2
////	//fmt.Printf("s: %v \n", s)
////
////	x := []string{"a", "b", "c"}
////	for v := range x {
////		fmt.Print(v)
////	}
////}
//
//type Fragment interface {
//	Exec(transInfo *TransInfo) error
//}
//
//type TransInfo struct {
//
//}
//type GetPodAction struct {
//}
//func (g *GetPodAction) Exec(transInfo *TransInfo) error {
//
//	return nil
//}
//
//var f Fragment = &GetPodAction{}
//
//type User struct{}
//type User1 User
//type User2 = User
//
//func (i User1) m1() {
//	  fmt.Println("m1")
//	}
//func (i User) m2() {
//	  fmt.Println("m2")
//	}

//func main() {
//	var i1 User1
//	var i2 User2
//	i1.m1()
//	i2.m2()
//}

type smallStruct struct {
	a, b int64
	c, d float64
}

func main() {
	smallAllocation()
}

//go:noinline
func smallAllocation() *smallStruct {
	return &smallStruct{}
}