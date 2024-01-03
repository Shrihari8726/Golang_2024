package main

import (
	"fmt"
)

// Creating a struct with items integer slice , to utilizze it for queue operations
type Queue struct {
	items []int
}

// Queue is a FIFO  or LILO
// Queue operation to enqueue item
// Enqueue is a receiver function accepts one arg to add to
// items integer slice and receiver accepts Queue struct as pointer
func (q *Queue) enQueue(i int) {
	q.items = append(q.items, i)
}

// Queue operation to dequeue item
// dequeue is also a receiver function accepts no arg
// but returns an int , which is the item to dequeue
// receiver accepts Queue struct as pointer
func (q *Queue) deQueue() int {
	itemToDequeue := q.items[0]
	q.items = q.items[1:]
	return itemToDequeue
}

func main() {
	fmt.Printf("This is Queue...")

	myQueue := Queue{}

	myQueue.enQueue(100)
	myQueue.enQueue(200)
	myQueue.enQueue(300)
	fmt.Println(myQueue)
	fmt.Println(myQueue.deQueue())
	fmt.Println(myQueue)

}
