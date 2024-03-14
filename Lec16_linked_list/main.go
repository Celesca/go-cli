package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Append(val int) {
	newNode := &Node{val: val}
	if l.head == nil {
		l.head = newNode
		return
	}
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = newNode
}

// Prepend adds a new node to the beginning of the list
func (l *LinkedList) Prepend(val int) {
	newNode := &Node{val: val, next: l.head}
	l.head = newNode
}

func (l *LinkedList) DeleteWithValue(val int) {
	if l.head == nil {
		return
	}
	if l.head.val == val {
		l.head = l.head.next
		return
	}

	cur := l.head
	for cur.next != nil {
		if cur.next.val == val {
			cur.next = cur.next.next
			return
		}
		cur = cur.next
	}

}

func (l *LinkedList) Display() {
	cur := l.head
	for cur != nil {
		fmt.Printf("%d ", cur.val)
		cur = cur.next
	}
	fmt.Println()
}

func main() {
	l := &LinkedList{}
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Display()
	l.Prepend(0)
	l.Display()
	l.DeleteWithValue(2)
	l.Display()
}
