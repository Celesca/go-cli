package main

import "fmt"

type Stack struct {
	data []int
	top  int
}

func (s *Stack) Push(x int) {
	s.top++
	s.data = append(s.data, x)
}

func (s *Stack) Pop() int {
	if s.top == -1 {
		return -1
	}
	x := s.data[s.top]
	s.top--
	return x
}

func (s *Stack) Top() int {
	if s.top == -1 {
		return -1
	}
	return s.data[s.top]
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func main() {

	s := &Stack{top: -1}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println("Top: ", s.Top())
	fmt.Println("Pop: ", s.Pop())
	fmt.Println("Top: ", s.Top())
	fmt.Println("Stack is empty?", s.IsEmpty())

}
