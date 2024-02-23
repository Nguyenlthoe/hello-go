package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func chanel_main() {
	s := []int{7, 2, 8, -9, 4, 0}

	//c := make(chan int, 2)  Channels can be buffered. Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func range_close_main() {
	c := make(chan int, 10) //A sender can close a channel to indicate that no more values will be sent.
	go fibonacci(cap(c), c)
	for i := range c { // The loop for i := range c receives values from the channel repeatedly until it is closed.
		fmt.Println(i)
	}
}
func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	//for {
	select {
	case c <- x:
		x, y = y, x+y // có chanel sẵn sàng thì thực hiện case đó.
	case <-quit:
		fmt.Println("quit")
		return

	default:
		fmt.Print("allo") // Khi không có channel nào sẵn sàng thì gọi default
	}
	//}
}

func fibo2_main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func end_main() {

	done := make(chan bool, 1)
	go worker(done)

	msg := <-done // Không có cái này chương trình sẽ end trước khi thực hiện xong luồng goroutines
	fmt.Print(msg)
}

func main31() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case res2 := <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
		fmt.Println(res2)
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
