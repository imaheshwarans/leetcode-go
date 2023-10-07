package main

// import (
// 	"fmt"
// 	"math/big"
// )

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func (l *ListNode) add(value int) {
// 	newNode := ListNode{Val: value}

// 	if l == nil {
// 		l = &newNode
// 		// fmt.Println("Head is added")
// 		return
// 	}

// 	// fmt.Println(l)

// 	curr := l
// 	for curr.Next != nil {
// 		curr = curr.Next
// 	}

// 	curr.Next = &newNode
// 	// fmt.Println("New node is added")
// }

// func (l *ListNode) Traverse() {

// 	if l == nil {
// 		// fmt.Printf("List is empty")
// 		return
// 	}

// 	// fmt.Println("\n\nTraverse")

// 	curr := l
// 	for curr != nil {
// 		debug := curr.Val
// 		fmt.Printf("%d", debug)
// 		curr = curr.Next
// 	}
// }

// func (l *ListNode) remove(value int) {
// 	if l == nil {
// 		return
// 	}

// 	if l.Val == value {
// 		l = l.Next
// 		return
// 	}

// 	curr := l
// 	for curr.Next != nil && curr.Next.Val != value {
// 		curr = curr.Next
// 	}

// 	if curr.Next != nil {
// 		curr.Next = curr.Next.Next
// 	}
// }

// func ReverseLinkedList(head *ListNode) {
// 	if head == nil {
// 		return
// 	}
// 	ReverseLinkedList(head.Next)
// 	fmt.Printf("%d ", head.Val)
// }

// // func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

// // 	var i int
// // 	// var tmp *big.Int
// // 	tmp := new(big.Int)
// // 	var result *ListNode
// // 	// var str1, str2 string
// // 	curr1 := l1
// // 	curr2 := l2
// // 	for curr1 != nil {
// // 		// tmp = tmp.Add(tmp, big.NewInt(int64(curr1.Val*(int(math.Pow10(i))))))
// // 		tmp = new(big.Int).Add(tmp, new(big.Int).Mul(big.NewInt(int64(curr1.Val)), new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(i)), nil)))
// // 		curr1 = curr1.Next
// // 		i++
// // 	}

// // 	i = 0
// // 	for curr2 != nil {
// // 		tmp = new(big.Int).Add(tmp, new(big.Int).Mul(big.NewInt(int64(curr2.Val)), new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(i)), nil)))
// // 		// tmp = tmp.Add(tmp, big.NewInt(int64(curr2.Val*(int(math.Pow10(i))))))
// // 		curr2 = curr2.Next
// // 		i++
// // 	}

// // 	// fmt.Println("Maheshwaran - >", tmp)

// // 	// reverse string
// // 	rns := []rune(tmp.String()) // convert to rune
// // 	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
// // 		rns[i], rns[j] = rns[j], rns[i]
// // 	}

// // 	// return the reversed string.
// // 	resultString := string(rns)

// // 	for _, digit := range resultString {
// // 		value := digit - '0'
// // 		if result == nil {
// // 			result = &ListNode{Val: int(value)}
// // 			continue
// // 		}
// // 		curr := result
// // 		for curr.Next != nil {
// // 			curr = curr.Next
// // 		}
// // 		curr.Next = &ListNode{Val: int(value)}
// // 	}

// // 	// if tmp.Uint64() == 0 {
// // 	// 	result = &ListNode{Val: int(tmp.Int64())}
// // 	// 	return result
// // 	// }
// // 	// for tmp.Uint64() != 0 {
// // 	// 	// digit := tmp.Mod(tmp, big.NewInt(10))
// // 	// 	digit := new(big.Int).Mod(tmp, big.NewInt(int64(10)))
// // 	// 	tmp = new(big.Int).Div(tmp, big.NewInt(int64(10)))
// // 	// 	if result == nil {
// // 	// 		result = &ListNode{Val: int(digit.Uint64())}
// // 	// 		continue
// // 	// 	}
// // 	// 	curr := result
// // 	// 	for curr.Next != nil {
// // 	// 		curr = curr.Next
// // 	// 	}
// // 	// 	curr.Next = &ListNode{Val: int(digit.Uint64())}
// // 	// }

// // 	return result
// // }

// func main() {

// 	list := &ListNode{Val: 9}

// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)

// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)
// 	list.add(9)

// 	// list.add(4)
// 	// list.add(3)
// 	// list.Traverse()

// 	list.Traverse()

// 	// fmt.Println("\n\n\n")

// 	// // // ReverseLinkedList(list.head)

// 	fmt.Println("\n\n\n")

// 	list1 := &ListNode{Val: 1}
// 	// list1.add(5)
// 	// list1.add(6)
// 	// list1.add(4)

// 	list1.Traverse()

// 	// // ReverseLinkedList(list.head)

// 	// list2 := &ListNode{Val: 7}
// 	// // list1.add(5)
// 	// list2.add(8)
// 	// list2.add(9)
// 	// fmt.Println("\n\n\n")
// 	// list2.Traverse()

// 	fmt.Println("\n\n\n")
// 	result := addTwoNumbers(list, list1)

// 	fmt.Println("\n\n\n")

// 	result.Traverse()

// 	// strNum := "1000000000000000000000000000001"
// 	// bigNum, success := new(big.Int).SetString(strNum, 10)

// 	// if !success {
// 	// 	fmt.Println("Failed to convert the string to a big integer")
// 	// 	return
// 	// }

// 	// fmt.Println("Successfully converted to big integer:", bigNum)

// 	// fmt.Println(value)
// }
