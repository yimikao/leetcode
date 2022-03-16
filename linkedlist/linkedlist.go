package main

import "fmt"

/**
 * Created by damilolaolayinka on 16/3/22.
 */

/**
A linked list is a linear collection of data elements whose order is not given by their physical placement in memory. Instead, each element points to the next. It is a data structure consisting of a collection of nodes which together represent a sequence. In its most basic form, each node contains: data, and a reference (in other words, a link) to the next node in the sequence. This structure allows for efficient insertion or removal of elements from any position in the sequence during iteration. More complex variants add additional links, allowing more efficient insertion or removal of nodes at arbitrary positions. A drawback of linked lists is that access time is linear (and difficult to pipeline). Faster access, such as random access, is not feasible. Arrays have better cache locality compared to linked lists.

The principal benefit of a linked list over a conventional array is that the list elements can be easily inserted or removed without reallocation or reorganization of the entire structure because the data items need not be stored contiguously in memory or on disk, while restructuring an array at run-time is a much more expensive operation. Linked lists allow insertion and removal of nodes at any point in the list, and allow doing so with a constant number of operations by keeping the link previous to the link being added or removed in memory during list traversal.

On the other hand, since simple linked lists by themselves do not allow random access to the data or any form of efficient indexing, many basic operations—such as obtaining the last node of the list, finding a node that contains a given datum, or locating the place where a new node should be inserted—may require iterating through most or all of the list elements

Operations that can be performed on singly linked lists include insertion, deletion and traversal.
*/

// The following code demonstrates how to add a new node with data "value" to the end of a singly linked list:

type Node struct {
	data int   // data to be stored
	next *Node // point to the next node's address
}

type LinkedList struct {
	head *Node // we only need to know the head node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil}
}

func NewNode(d int) *Node {
	return &Node{data: d, next: nil}
}

func (l *LinkedList) addStart(d int) {
	// create a new node
	var newNode = NewNode(d)

	// make current head the next
	newNode.next = l.head

	//then put new node at the beginning(head)
	l.head = newNode
}

func (l *LinkedList) deleteAt(d int) {

	// these are to keep track of position for deletion
	var (
		currentNode *Node
		prevNode    *Node
	)

	currentNode = l.head

	// empty
	if currentNode == nil {

		return
	}

	// if head node holds the data to be deleted
	if currentNode != nil {

		if currentNode.data == d {
			l.head = currentNode.next
			return
		}

	}

	for currentNode != nil {
		// begin to iterate through none nil values i.e while not reached the end
		// and try return early

		// no need to iterate further
		if currentNode.data == d {

			break
		}

		// continue to move down the list
		// keep track of this so i can unlink it from current node easily
		// this node has been confirmed to not hold the value
		prevNode = currentNode

		// check if this node has the value to easily unlink from it
		currentNode = currentNode.next

	}

	// continue from early break
	// unlink prevNode unto currentNode's next
	prevNode.next = currentNode.next
	// erase current node
	currentNode = nil

}

func printLinkedList(l *LinkedList) {

	var currentNode *Node

	// start from the head as usual
	currentNode = l.head

	for currentNode != nil {
		// we continue to print through as long as we haven't reached the end(which points to null)
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
