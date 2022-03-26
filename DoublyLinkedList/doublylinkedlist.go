package main

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

	if l.head != nil {
		n.next = l.head
		l.head.prev = n
	}
	l.head = n
}

func (l *LinkedList) InsertLast(v int) {
	n := &Node{data: v}

	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}

	if l.tail != nil {
		l.tail.next = n
		n.prev = l.tail
	}

	l.tail = n
}

func main() {

}
