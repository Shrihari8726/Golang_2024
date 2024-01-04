package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *linkedList) append(n *node) {
	current := l.head
	prevLength := l.length
	for l.length != 1 {
		current = current.next
		l.length--
	}
	current.next = n
	l.length = prevLength + 1
}

func (l *linkedList) addAfter(n *node, aa int) {
	current := l.head
	next := current.next
	prevLength := l.length
	for l.length != 0 {
		if current.data != aa {
			temp := next
			next = next.next
			current = temp
			l.length--
		} else if current.data == aa {
			current.next = n
			n.next = next
			l.length--
		}
	}
	l.length = prevLength + 1

}

func (l linkedList) print() {
	current := l.head
	for l.length != 0 {
		fmt.Printf("%d ", current.data)
		current = current.next
		l.length--
	}
}

func main() {
	fmt.Println("This is linked list...")
	myList := linkedList{}
	node1 := &node{data: 23}
	node2 := &node{data: 34}
	node3 := &node{data: 56}
	node4 := &node{data: 72}
	node5 := &node{data: 47}
	node6 := &node{data: 64}
	node7 := &node{data: 38}
	node8 := &node{data: 54}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.print()
	fmt.Println(" ")
	myList.append(node5)
	myList.append(node6)
	myList.print()
	fmt.Println(" ")
	myList.addAfter(node7, 23)
	myList.addAfter(node8, 72)
	myList.print()
}
