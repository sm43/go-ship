package main

import (
	"fmt"
	"sync"
	"time"
)

// using locks

func case6() {
	var mut sync.Mutex

	alice := 10000
	bob := 10000

	go func() {
		for i := 0; i < 1000; i++ {
			mut.Lock()
			alice -= 1

			// without this incre and decre happens at same time
			// so the total would be correct

			// at some given time, the total will not be correct
			// but after completing all routine the total will be correct

			// mut.Unlock()
			// mut.Lock()
			bob += 1
			mut.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mut.Lock()
			bob -= 1
			// mut.Unlock()
			// mut.Lock()
			alice += 1
			mut.Unlock()
		}
	}()

	start := time.Now()
	for time.Since(start) < 1*time.Second {
		mut.Lock()
		if alice+bob != 20000 {
			fmt.Printf("\n observed violation, alice = %v, bob = %v, sum = %v \n", alice, bob, alice+bob)
		}
		mut.Unlock()
	}

}
