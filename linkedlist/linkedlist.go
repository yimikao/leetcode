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
	data int   // assume we're storing integers
	next *Node // we point to the next node's address
}

type LinkedList struct {
	head *Node // we only need to know the head node
}

func (l *LinkedList) addNode(value int) {

	var newNode = &Node{data: value} // new node to be added

	// why didn't this work tho
	// var newNode *Node
	// newNode.data = value

	var currentNode *Node // we use this to keep track of node we have/curentky on when interating

	// when linked-list is empty just add head node
	if l.head == nil {
		l.head = newNode
	} else {
		// if not empty, check where next is null(i.e end) so we add there
		// we have to start from the head
		currentNode = l.head

		// we begin iterating through the linkedlist
		for {
			// that is we aren't at the end yet
			if currentNode.next != nil {
				// we move to the next node
				currentNode = currentNode.next
			} else {
				// we add the current node to the last node
				// and ofcourse, the lastest node will now point to null
				currentNode.next = newNode

				// we stop iterating
				break
			}
		}
	}
}

func printLinkedList(l *LinkedList) {
	if l.head == nil {
		fmt.Println("nothing to print")
		return
	}

	var currentNode *Node

	// we start from the head as usual
	currentNode = l.head

	for {

		// we continue to print through as long as we haven't reached the end(which points to null)
		fmt.Println(currentNode.data)
		if currentNode.next != nil {
			currentNode = currentNode.next
		} else {
			// we've reached the end so we break
			break
		}
	}
}

func main() {
	var ll1 = LinkedList{}
	var ll2 = LinkedList{head: &Node{data: 99}}

	// ll1.addNode(1)
	ll2.addNode(2)
	printLinkedList(&ll1)
	printLinkedList(&ll2)

}
