package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() { // GO không có khái niệm class
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func method_nonType() { // có thể định nghĩa cho các đối tượng không phải struct
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
