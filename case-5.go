package main

import (
	"fmt"
	"sync"
	"time"
)

// Why do we need mutex?

// add a lock
var mu sync.Mutex

func case5() {

	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			// run as end of current func body
			defer mu.Unlock()
			counter = counter + 1
		}()
	}

	// can use wait grp instead of sleep to let other complete
	// their work
	// let everyone do their work
	time.Sleep(time.Second * 1)
	// then print
	mu.Lock()
	fmt.Println(counter)
	mu.Unlock()
}

// Why do we need mutex?

// In this case we will miss some of the increments done
// as all go routine are trying to update same var without
// any lock, so sometime some might try to update var at same
// time
// func main() {

// 	counter := 0
// 	for i := 0; i < 1000; i++ {
// 		go func() {
// 			// will not run atomically
// 			counter = counter + 1
// 		}()
// 	}
// 	time.Sleep(time.Second * 3)
// 	fmt.Println(counter)
// }
