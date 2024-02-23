package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func goifmain() {
	fmt.Println(sqrt(2), sqrt(-4))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // biến v được khởi tạo trong if, là local variable, mất khi thực thi xong if.
		return v
	}

	// if v := math.Pow(x, n); v < lim { // if else
	// 	return v
	// } else {
	// 	fmt.Printf("%g >= %g\n", v, lim)
	// }
	return lim
}

func pow_main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
