package main

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) InsertFirst(v int) {

	n := &Node{data: v}

	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}

	n.next = l.head
	l.head.prev = n
	l.head = n
}

func (l *LinkedList) InsertLast(v int) {
	n := &Node{data: v}

	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}

	l.tail.next = n
	n.prev = l.tail
	l.tail = n

}

func (l *LinkedList) RemoveByValue(v int) bool {

	if l.head == nil {

		return false

	}

	if l.head.data == v {

		l.head = l.head.next
		return true

	}

	if l.tail.data == v {
		l.tail = l.tail.prev
		return true
	}

	current := l.head
	for current.next != nil {

		if current.data == v {
			current.prev.next = current.next
			current = nil
			return true
		}
		current = current.next

	}
	return false
}

func printLinkedList(l LinkedList) {
	items := []int{}

	if l.head == nil {
		fmt.Println("empty list")
		return
	}

	current := l.head
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	fmt.Println(items)

}

func main() {

	ll := LinkedList{}

	// ll.InsertFirst(4)
	// ll.InsertFirst(14)
	// ll.InsertFirst(77)
	ll.InsertFirst(98)
	ll.InsertLast(9)

	printLinkedList(ll)
	fmt.Println(ll.head)
	fmt.Println(ll.tail)
}
