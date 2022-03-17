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

func (l *LinkedList) addStart(d int) {
	var newNode = NewNode(d)

	newNode.next = l.head

	l.head = newNode
}

func (l *LinkedList) deleteAt(d int) {

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

	ll2.addStart(2)
	ll2.addStart(18)
	ll2.addStart(15)
	fmt.Println("here is the add updated linkedList")
	printLinkedList(ll2)

	ll2.deleteAt(18)
	fmt.Println("here is the delete updated linkedList")
	printLinkedList(ll2)

}
