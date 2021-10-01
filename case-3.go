package main

import (
	"fmt"
	"time"
)

// doing things periodically
// go routine till the main func doesn't finish
// i.e. after completing `time.Sleep(5 * time.Second)`
// go routine will be terminated so need to use wait
// if it wanna wait
func case3() {
	time.Sleep(1 * time.Second)
	fmt.Println("do")
	go periodic()
	fmt.Println("done")
	time.Sleep(5 * time.Second)
}

func periodic() {
	for {
		fmt.Println("sm")
		time.Sleep(3 * time.Second)
	}
}
