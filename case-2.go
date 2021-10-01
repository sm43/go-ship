package main

import (
	"fmt"
	"sync"
)

func case2() {

	var wg sync.WaitGroup

	// If it waits for more than required go routine
	// then after finishing the last one it errors out
	// wg.Add(6)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(x int) {
			// send and done
			sendRPC(x)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func sendRPC(x int) {
	fmt.Println(x)
}
