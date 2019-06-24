package main

import "fmt"

type People struct {

}

func (p *People) ShowA() {
	fmt.Println("ShowA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("ShowB")
}

type Teachers struct {
	People
}

func (t *Teachers) ShowB() {
	fmt.Println("teacher show B")
}

func main() {
	t := &Teachers{}
	t.ShowA()
}
