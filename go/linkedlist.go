package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type List struct {
	Head *Node
}

func (l *List) Add(value int) {

	newNode := Node{Data: value, Next: nil}

	if l.Head == nil {
		l.Head = &newNode
		return
	}

	current := l.Head

	for current.Next != nil {
		current = current.Next
	}

	current.Next = &newNode
	return
}

func (l *List) Remove(data int) {

	if l.Head == nil {
		return
	}

	if l.Head.Data == data {
		l.Head = l.Head.Next
	}

	tmp := l.Head

	for tmp.Next != nil && tmp.Next.Data != data {
		tmp = tmp.Next
	}

	if tmp.Next != nil {
		tmp.Next = tmp.Next.Next
	}

}

func (l *List) Traverse() {
	if l.Head == nil {
		return
	}

	curr := l.Head

	for curr != nil {
		fmt.Printf(" %d -> ", curr.Data)
		curr = curr.Next
	}

	fmt.Printf("null")

}

func main() {

	list := &List{}

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)

	fmt.Println("\nInitial List: ")
	list.Traverse()

	list.Remove(3)

	fmt.Println("\nAfter removing 3 in List: ")
	list.Traverse()

	list.Remove(1)

	fmt.Println("\nAfter removing 1 in List: ")
	list.Traverse()

	list.Remove(5)

	fmt.Println("\nAfter removing 5 in List: ")
	list.Traverse()
}
