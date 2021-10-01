package main

import (
	"fmt"
	"sync"
	"time"
)

// using lock
func case7() {

	count := 0
	finished := 0
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		go func() {
			vote := requestVote()
			mu.Lock()
			if vote {
				count++
			}
			finished++
			mu.Unlock()
		}()
	}

	for {
		fmt.Println("Check: ", finished)
		mu.Lock()
		// keep checking till I win or everyone is done voting
		if count >= 5 || finished == 10 {
			fmt.Println("I am breaking now ... ", finished)
			break
		}
		mu.Unlock()

		// without this it would be bad code
		// this will keep checking and keep consuming cpu
		// so adding a sleep would be better
		time.Sleep(50 * time.Millisecond)
	}
	if count >= 5 {
		fmt.Println("I win !")
	} else {
		fmt.Println("I lose !")
	}

}

func requestVote() bool {
	return true
}
