package main

import "fmt"

// Vertex 定义一个二维数组
type Vertex struct {
	Lat, Long float64
}

func main() {
	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68, -75.5,
		},
		"Google": Vertex{
			37.5, -122.2,
		},
	}

	fmt.Println(m)
}
