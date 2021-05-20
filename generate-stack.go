package main

const STACK_INIT_SIZE = 100

type Stack struct {
	list *List
	len  int
}

type List struct {
	Data int
	Next *List
}
func InitStack() *Stack {
	return &Stack{
		&List{},
		1,
	}
}

func (s *Stack) push(value int) {
	l := &List{Data: value}
	curl := s.list
	for curl.Next != nil {
		curl = curl.Next
	}
	curl.Next = l
	s.len += 1
}

func (s *Stack) pop() {
	pre := s.list
	curl := s.list.Next
	for curl.Next != nil {
		curl = curl.Next
		pre = pre.Next
	}
	pre.Next = nil
	s.len -= 1
}
//func main() {
//	s := InitStack()
//	s.push(1)
//	s.push(2)
//	fmt.Println(s.len)
//	s.pop()
//	s.pop()
//	fmt.Println(s.len)
//}