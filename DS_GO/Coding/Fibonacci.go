package main

import "fmt"

type FibSlice struct {
	items []int
}

func (fs *FibSlice) fib(n int) {

	i := 0
	j := 1
	fs.items = append(fs.items, i)
	fs.items = append(fs.items, j)
	for n != 0 {
		temp := i
		i = j
		j = i + temp
		fs.items = append(fs.items, j)
		n--
	}
	fmt.Printf("%d ", fs.items)
}

func main() {
	fmt.Println("This is Fibonacci series:")
	myList := FibSlice{}
	myList.fib(9)
}
