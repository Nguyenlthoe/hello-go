package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() { //cơ chế chính trong quản lý trạng thái trong Go là giao tiếp giữa các channel.
	// Có 1 số lựa chọn để quản lý trạng thái. 1 cái là sử dụng package sync/atomic
	// Biến truy cập bởi nhiều luồng.
	var ops atomic.Uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {

				ops.Add(1)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops.Load())
}
