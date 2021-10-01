package main

import (
	"fmt"
	"sync"
)

// using condition
func case8() {

	count := 0
	finished := 0
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	for i := 0; i < 10; i++ {
		go func() {
			vote := requestVote()
			mu.Lock()
			defer mu.Unlock()
			if vote {
				count++
			}
			finished++
			cond.Broadcast()
		}()
	}

	mu.Lock()
	// condition avoid using random sleep
	// each time counter increased it calls cond Broadcast
	// which will make routine to check condition
	// if not true then go in wait stage and wait for
	// next broadcast
	for count < 5 && finished != 10 {
		fmt.Println("Wake up")
		cond.Wait()
	}
	// once the condition becomes true then we go ahead otherwise
	// wait it happens
	if count >= 5 {
		fmt.Println("I win !")
	} else {
		fmt.Println("I lose !")
	}
	mu.Unlock()
}

func requestVote() bool {
	return true
}
