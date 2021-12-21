package main

import "fmt"

type Proxy func() (m string)

type Backend func() Proxy


type defaultFactory struct {
	backend Backend
}

func main() {
	var p Proxy
	df := defaultFactory{}
	p = df.backend()
	fmt.Println(p)
}