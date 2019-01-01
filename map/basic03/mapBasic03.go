package main

import "fmt"

// Vertex 定义了一个二维数组
type Vertex struct {
	Lat, Long float64
}

func main() {
	var m = map[string]Vertex{
		"Bell Labs": {40.34, 78.5},
		"Google":    {37.5, -122.2},
	}
	fmt.Println(m)
	m["abcd"] = Vertex{
		22.3, 25.5,
	}
	fmt.Println(len(m))
}
