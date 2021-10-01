package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var done bool

func case4() {
	time.Sleep(1 * time.Second)
	fmt.Println("started")
	go timely()
	time.Sleep(5 * time.Second)

	mu.Lock()
	// As the var is shared between main and go routin
	// change it after starting a lock
	done = true

	fmt.Println("I have changed the value but I will sleep to show that the go routine which has also called lock will be blocked till I unblock")
	time.Sleep(time.Second * 5)
	mu.Unlock()

	fmt.Println("finished the routine")

	// this will let I am completed to print otherwise it will
	// be missed
	time.Sleep(3 * time.Second)
}

// if someone calls lock and then someone else calls lock
// it will be blocked till the firstone unlocks

func timely() {
	for {
		fmt.Println("sm")
		time.Sleep(1 * time.Second)
		fmt.Println("I am calling lock now !")
		mu.Lock()
		if done {
			// we can return from here
			// in general unlock and return but as main is finishing
			// unlock wouldn't matter

			// return

			break
		}
		mu.Unlock()
	}
	fmt.Println("I am completed")
}
