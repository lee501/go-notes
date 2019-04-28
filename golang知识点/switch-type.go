//检查v.(type)输出值
package main

import "fmt"

//定义Shaper接口
type Shaper interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return 3.14 * c.Radius *c.Radius
}

type Square struct {
	length float64
}

func (s Square) Area() float64 {
	return s.length * s.length
}

func main() {
	s1 := &Square{3.3}
	whichType(s1)

	s2 := Square{ 3.3 }
	whichType(s2)

	c1 := new(Circle)
	c1.Radius = 2.3
	whichType(c1)
}

func whichType(n Shaper) {
	//v的值包含类型和value
	switch v := n.(type) {
	case *Square:
		fmt.Printf("Type Square %v\n", v)
	case Square:
		fmt.Printf("Type Square %T\n", v)
	case *Circle:
		fmt.Printf("Type Circle %T\n", v)
	case nil:
		fmt.Printf("nil value")
	default:
		fmt.Printf("Unexpected type %T", v)
	}
}