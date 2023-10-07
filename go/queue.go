package main

import "fmt"

type Queue struct {
	Data  []int
	Head  int
	Tail  int
	Count int
}

func (s *Queue) Display() {
	if s.Count <= 0 {
		fmt.Println("\nQueue is empty")
		return
	}
	for _, data := range s.Data {
		fmt.Printf(" %d ", data)
	}
}

type QueueList interface {
	Enqueue(int)
	Dequeue()
	Display()
}

func (s *Queue) Enqueue(value int) {
	if s.Data == nil {
		s.Head = 0
		s.Tail = 0
	}
	s.Data = append(s.Data, value)
	s.Count++
	s.Tail = s.Head
}

func (s *Queue) Dequeue() {

	if s.Count <= 0 {
		fmt.Println("\nQueue is empty")
		return
	}

	s.Data = s.Data[1:]

	s.Count--
	s.Tail = s.Count
}

func NewQueue() QueueList {
	return &Queue{}
}

func main() {

	Queue := NewQueue()

	Queue.Enqueue(1)
	Queue.Enqueue(2)
	Queue.Enqueue(3)

	Queue.Display()

	fmt.Println("\nAfter single pop")

	Queue.Dequeue()

	Queue.Enqueue(4)

	Queue.Display()

	Queue.Dequeue()
	Queue.Dequeue()
	fmt.Println("\nAfter all pop")
	Queue.Display()

	Queue.Dequeue()

}
