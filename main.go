package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var done bool

func main() {
	time.Sleep(1 * time.Second)
	fmt.Println("started")
	go timely()
	time.Sleep(5 * time.Second)

	mu.Lock()
	done = true
	mu.Unlock()

	fmt.Println("finished the routine")

	// this will let I am completed to print otherwise it will
	// be missed
	time.Sleep(3 * time.Second)
}

func timely() {
	for {
		fmt.Println("sm")
		time.Sleep(1 * time.Second)
		mu.Lock()
		if done {
			break
		}
		mu.Unlock()
	}
	fmt.Println("I am completed")
}
