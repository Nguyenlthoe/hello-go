package main

import "fmt"

func defer_test_func() {

	defer fmt.Println("defer 1")

	fmt.Println("defer2")
}

func main() {
	fmt.Println("counting")
	defer_test_func()
	// for i := 0; i < 10; i++ { // last in first out
	// 	defer fmt.Println(i) // lệnh trì hoãn khi tất cả các chương trình chạy xong
	// }
	fmt.Println("done")
}
