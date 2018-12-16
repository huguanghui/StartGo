// slice类型的使用
package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	printSlice(s)
}

func printSlice(x []int) {
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(x), cap(x), x)
}
