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
		return
	}
	l.head = n
}

func main() {

}
