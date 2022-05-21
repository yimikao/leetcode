package main

import "fmt"

/******
RECURSION IS JUST SAYING DIVIDING THE PROBLEM INTO TWO HALVES,
WITH TWO FUNCTIONS HANDLING EACH HALF,
THE SECOND ONE IS CALLED ONLY IF A CONDITION IS TRUE
UNTIL ONE OF YOU REACH A CONDITION THAT ISN"T TRUE.
****/

func recurse(arr []int, pos int) {
	arr[pos] = arr[pos] * 2
	if pos == len(arr)-1 {
		return
	}
	recurse(arr, pos+1)
}

// a better way to write it
func recurse2(arr []int, pos int) int {

	// we use the condition that will not cause a return
	if pos != len(arr) {
		recurse2(arr, pos+1)
	}
	// each of the functions does what it needs to do
	arr[pos] = arr[pos] * 2
	return 0
}

func main() {
	arr := []int{11, 6, 9, 4, 19, 3}

	recurse2(arr, 0)

	fmt.Printf("arr: %v\n", arr)

}

// recursion function needs a value that is changing to keep goin
//a state of the value in ie at one point needs to stop it

func reverseList(head *node) *node {
	//we call each current node we pick head here cause we assume it's the only thing we are
	//dealing with

	//i'll always be passed a node, which has already been made an
	//head to do heading along the way too(or make sure i only deal with 1 element, tha is head)
	if head == nil {
		return nil
	}

	newHead = head
	if head.next != nil {
		// if condition is still true, second func can be called with other problem half
		newHead = reverseList(head.next)

		//this shoulld be done also if condition is true
		head.next.next = head
	}
	// what each func is meant to do regardless
	head.next = nil
	return newHead
}
