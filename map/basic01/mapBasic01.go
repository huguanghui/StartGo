package main

import "fmt"

// Vertex 用于测试的一个二维数组
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Huguanghui"] = Vertex{
		40.2, 4.6,
	}
	m["abc"] = Vertex{
		35.5, 49.5,
	}
	m["test"] = Vertex{
		23.5, 24.5,
	}
	fmt.Println(m["Huguanghui"])
	fmt.Println(m["abc"])
	fmt.Println(len(m))
}
