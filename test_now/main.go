package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// fmt.Println("Hello World!")
	fmt.Printf("%v\n", randParse(7))
}

func randParse(a int) []int {
	arr := make([]int, 0)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < a; i++ {
		arr = append(arr, int(rand.Int31n(45)))
	}

	return arr
}
