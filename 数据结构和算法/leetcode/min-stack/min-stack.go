package main

/*
	实现最小栈，支持push, pop, getMin操作，时间复杂度为O(1)

    思路：
		维护两个栈，一个保存数据，另一个保存最小值的索引
*/
func main() {

}

type Stack struct {
	Data     []int
	MinIndex []int //最小值索引
	L        int   //stack length
}

func (s *Stack) push(item int) {
	s.Data = append(s.Data, item)
	s.L++
	if len(s.MinIndex) == 0 {
		s.MinIndex = append(s.MinIndex, 0)
	} else {
		if s.Data[len(s.MinIndex)-1] > item {
			s.MinIndex = append(s.MinIndex, s.L-1)
		}
	}
}

func (s *Stack) pop() int {
	if s.L == 0 {
		return -1
	}
	t := s.Data[s.L-1]
	s.Data = s.Data[:s.L-1]
	if s.Data[len(s.MinIndex)-1] == t {
		s.MinIndex = s.MinIndex[:len(s.MinIndex)-1]
	}
	return t
}

func (s *Stack) getMin() int {
	return s.Data[s.MinIndex[len(s.MinIndex)-1]]
}

func Construct() *Stack {
	return new(Stack)
}
