package main

import "fmt"

type node struct {
	value int
	next  *node
}

type linkedlist struct {
	head *node
}

// reverse list and return new head
//***ITERATIVE SOLUTION***//
func reverseList(ll *linkedlist) *node {
	var prevNode *node
	currNode := ll.head

	for currNode != nil {
		nextCurrNode := currNode.next
		currNode.next = prevNode
		prevNode = currNode
		currNode = nextCurrNode
	}
	ll.head = prevNode
	return prevNode

}

//***RECURSIVE SOLUTION***//
func reverseList2(head *node) *node {
	//we call each current node we pick head here cause we assume it's the only thing we are
	//dealing with

	//i'll always be passed a node, which has already been made an
	//head to do heading along the way too(or make sure i only deal with 1 element, tha is head)
	if head == nil {
		return nil
	}

	newHead := head
	if head.next != nil {
		// if condition is still true, second func can be called with other problem half
		newHead = reverseList2(head.next)

		//this shoulld be done also if condition is true
		head.next.next = head
	}
	// what each func is meant to do regardless
	head.next = nil
	return newHead
}

func printLinkedList(ll linkedlist) {
	currNode := ll.head
	// fmt.Printf("currNode   ssssssssssss: %v\n", currNode)
	for currNode != nil {
		fmt.Printf("currNode: %v\n", currNode)
		currNode = currNode.next
	}

}

func main() {
	ll := linkedlist{}
	ll.head = &node{
		value: 1,
		next: &node{
			value: 2,
			next: &node{
				value: 16,
				next:  nil,
			},
		},
	}
	printLinkedList(ll)
	// fmt.Printf("reverseList(ll.head): %v\n", reverseList(&ll))
	fmt.Printf("reverseList2(ll.head): %v\n", reverseList2(ll.head))
	// fmt.Printf("ll.head: %v\n", ll.head)
	printLinkedList(ll)
}
