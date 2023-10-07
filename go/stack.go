package main

import "fmt"

type Stack struct {
	Data  []int
	Count int
}

func (s *Stack) Display() {
	if s.Count <= 0 {
		fmt.Println("\nStack is empty")
		return
	}
	for _, data := range s.Data {
		fmt.Printf(" %d ", data)
	}
}

type OrderedList interface {
	Push(int)
	Pop()
	Display()
}

func (s *Stack) Push(value int) {
	s.Data = append(s.Data, value)
	s.Count++
}

func (s *Stack) Pop() {

	if s.Count <= 0 {
		fmt.Println("\nStack is empty")
		return
	}

	s.Count--
	s.Data = s.Data[:s.Count]

}

func NewStack() OrderedList {
	return &Stack{}
}

func main() {

	stack := NewStack()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.Display()

	fmt.Println("\nAfter single pop")

	stack.Pop()

	stack.Display()
	stack.Push(4)
	// stack.Pop()
	// stack.Pop()
	fmt.Println("\nAfter all pop")

	stack.Display()

	stack.Pop()

}
