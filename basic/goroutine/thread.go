// A goroutine is a lightweight thread managed by the Go runtime.
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main_thread() {
	go say("world")
	say("hello")
}
