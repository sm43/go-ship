package main

import (
	"fmt"
	"sync"
	"time"
)

func case1() {
	var a string
	var wg sync.WaitGroup

	// Wait only for one go routing for done !
	wg.Add(1)

	go func() {
		fmt.Println("I am sleeping for 10 sec !")
		time.Sleep(10 * time.Second)
		fmt.Println("done sleeping !")

		a = "hello world"

		// Now I am done you can finish the waiting
		wg.Done()
	}()
	a = "before hello"
	fmt.Println(a)

	// Waits till all go routing are done
	// it looks how many are added in wg.Add()
	wg.Wait()
	fmt.Println(a)
}
