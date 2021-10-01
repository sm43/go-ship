package main

import (
	"fmt"
	"time"
)

// Channels
func case9() {
	c := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		<-c
	}()

	start := time.Now()
	c <- true
	fmt.Println("waited for ", time.Since(start))
}

// Sending and receiving should always be there
// having only one doesn't work

/*
	c:= make(chan bool)
	<- c // this will wait till someone is receiving but no one so it never reaches next
    c <- true

	deadlock !
*/
