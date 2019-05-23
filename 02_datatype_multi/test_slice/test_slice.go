// slice类型的使用
package main

import "fmt"

// GlobalParam 全局切片变量定义
var GlobalParam = []string{"test01", "test02", "test03"}

func main() {
	s := []int{1, 2, 3}
	printSlice(s)

	for i, elt := range GlobalParam {
		fmt.Printf("index: %d value: %s\n", i, elt)
	}
}

func printSlice(x []int) {
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(x), cap(x), x)
}
