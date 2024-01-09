package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func doSomething() {
	time.Sleep(time.Second * 2)
	fmt.Println("I did something")
	wg.Done()
}

func doSomethingElse() {
	time.Sleep(time.Second * 2)
	fmt.Println("I did something else")
	wg.Done()
}

func main() {
	wg.Add(2)
	start := time.Now()
	go doSomething()
	go doSomethingElse()
	wg.Wait()
	fmt.Println("I think i am done")
	elapsed := time.Since(start)
	fmt.Printf("Process took %d", elapsed)
}
