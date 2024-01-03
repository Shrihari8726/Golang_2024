package main

import (
	"fmt"
)

// Creating a struct with items integer slice , to utilizze it for stack operations
type Stack struct {
	items []int
}

// Stack is a FILO  or LIFO
// Stack operation to push item
// Push is a receiver function accepts one arg to add to
// items integer slice and receiver accepts Stack struct as pointer
func (s *Stack) push(i int) {
	s.items = append(s.items, i)
}

//Stack operation to pop item
// Pop is also a receiver function accepts no arg
// but returns an int , which is the item to poped
// receiver accepts Stack struct as pointer

func (s *Stack) pop() int {
	l := len(s.items) - 1
	itemToRemove := s.items[l]
	s.items = s.items[:l]
	return itemToRemove
}

func main() {
	fmt.Println("This is Stack Intro...")

	myStack := Stack{}
	myStack.push(100)
	myStack.push(200)
	myStack.push(300)
	fmt.Println(myStack)

	fmt.Println(myStack.pop())
	fmt.Println(myStack)
}
