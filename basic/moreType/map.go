package main

import "fmt"

type Vertex2 struct {
	Lat, Long float64
}

var m = map[string]Vertex2{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func map_main() {
	delete(m, "Google")
	m["Nguyenlt"] = Vertex2{1672.123, 12381.2}
	fmt.Println(m)
}
