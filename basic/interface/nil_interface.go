package main

import "fmt"

type Ii interface {
	Mi()
}

type Ti struct {
	S string
}

func (t *Ti) Mi() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func nil_main() {
	var i Ii

	var t *Ti
	i = t
	nil_describe(i)
	i.Mi()

	i = &Ti{"hello"}
	nil_describe(i)
	i.Mi()
}

func nil_describe(i Ii) {
	fmt.Printf("(%v, %T)\n", i, i)
}
