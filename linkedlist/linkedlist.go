package main

import "fmt"

/**
 * Created by damilolaolayinka on 16/3/22.
 */

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil}
}

func NewNode(d int) *Node {
	return &Node{data: d, next: nil}
}

func (l *LinkedList) InsertFirst(d int) {
	var newNode = NewNode(d)
	// newNode.next = l.head
	// l.head = newNode

	if l.head != nil {
		newNode.next = l.head
	}
	l.head = newNode
}

func (l *LinkedList) InsertLast(d int) {
	n := NewNode(d)

	if l.head == nil {
		l.head = n
		return
	}

	current := l.head

	for current.next != nil {
		current = current.next
	}

	current.next = n
}

func (l *LinkedList) RemoveFirst() {
	if l.head == nil {
		return
	}

	l.head = l.head.next
}

func (l *LinkedList) RemoveByValue(d int) {

	var (
		currentNode *Node
		prevNode    *Node
	)

	currentNode = l.head

	if currentNode == nil {

		return
	}

	if currentNode != nil {

		if currentNode.data == d {
			l.head = currentNode.next
			return
		}

	}

	for currentNode != nil {
		if currentNode.data == d {

			break
		}

		prevNode = currentNode

		currentNode = currentNode.next

	}

	prevNode.next = currentNode.next
	currentNode = nil

}

func printLinkedList(l *LinkedList) {

	var currentNode *Node

	currentNode = l.head

	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

func main() {

	var ll2 = NewLinkedList()

	ll2.InsertFirst(2)
	ll2.InsertFirst(18)
	ll2.InsertFirst(15)
	fmt.Println("here is the add updated linkedList")
	printLinkedList(ll2)

	ll2.RemoveByValue(18)
	fmt.Println("here is the delete updated linkedList")
	printLinkedList(ll2)

}
