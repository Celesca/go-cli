package main

import "fmt"

type Set struct {
	elements map[interface{}]bool
}

func NewSet() *Set {
	return &Set{elements: make(map[interface{}]bool)}
}

func (s *Set) Add(element interface{}) {
	s.elements[element] = true
}

func (s *Set) Remove(element interface{}) {
	delete(s.elements, element)
}

func (s *Set) Contains(element interface{}) bool {
	_, ok := s.elements[element]
	return ok
}

func (s *Set) Size() int {
	return len(s.elements)
}

func (s *Set) Clear() {
	s.elements = make(map[interface{}]bool)
}

func main() {
	s := NewSet()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	fmt.Println("Set contains 1 : ", s.Contains(1))
	fmt.Println("Set contains 4 : ", s.Contains(4))

	s.Remove(2)
	fmt.Println("Set contains 2 : ", s.Contains(2))
	fmt.Println("Set size: ", s.Size())

	s.Clear()
	fmt.Println("Set size: ", s.Size())
}
