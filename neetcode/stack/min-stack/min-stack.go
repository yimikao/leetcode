package main

import (
	"fmt"
	"math"
)

type Minstack struct {
	stack    []int
	minstack []int
}

// key is to just think see minstack as stack of min(of each element) values only
//e.g stack=[2,-1,3,1,3,-5]  minstack=[2,-1,-1,-1,-1,-5]
// push
// pop
// top
// getmin

func (m *Minstack) push(data int) {
	m.stack = append(m.stack, data)

	var min int
	if len(m.minstack) < 1 {
		min = data
	} else {
		min = int(math.Min(float64(data), float64(m.minstack[len(m.minstack)-1])))
	}
	m.minstack = append(m.minstack, min)
}

func (m *Minstack) pop() {
	m.stack = m.stack[:len(m.stack)-1]
	m.minstack = m.minstack[:len(m.minstack)-1]
}
func (m *Minstack) top() int {
	return m.stack[len(m.stack)-1]
}

func (m *Minstack) getmin() int {
	return m.minstack[len(m.minstack)-1]
}

func main() {
	MS := &Minstack{
		stack:    []int{},
		minstack: []int{},
	}

	MS.push(2)
	MS.push(-1)
	MS.push(3)
	MS.push(1)
	MS.push(3)
	MS.push(-5)
	MS.pop()
	fmt.Printf("MS.top(): %v\n", MS.top())
	fmt.Printf("MS.getmin(): %v\n", MS.getmin())
	fmt.Printf("MS.stack: %v\n", MS.stack)
	fmt.Printf("MS.minstack: %v\n", MS.minstack)
}
