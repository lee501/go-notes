package main

import "fmt"

func main() {
	i := 1
	s := []string{"A", "B", "C"}
	s[i-1], i = "Z", 2
	fmt.Printf("s: %v \n", s)
}

type Fragment interface {
	Exec(transInfo *TransInfo) error
}

type TransInfo struct {

}
type GetPodAction struct {
}
func (g *GetPodAction) Exec(transInfo *TransInfo) error {

	return nil
}

var f Fragment = &GetPodAction{}