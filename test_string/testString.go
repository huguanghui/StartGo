package main

import "fmt"

func main() {
	abc := string("test/abc/")
	n := len(abc)
	fmt.Printf("%c\n", abc[n-1])
}
