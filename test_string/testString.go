package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	abc := string("test/abc/")
// 	n := len(abc)
// 	fmt.Printf("%c\n", abc[n-1])
// }

func main() {
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("abc", "d"))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
}
