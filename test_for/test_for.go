package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println("Hello World!")
	}
	// for-range遍历字符串
	address := "hello world!"
	// for i := 0; i < len(address); i++ {
	// 	fmt.Printf("%c \n", address[i])
	// }
	for index, val := range address {
		fmt.Printf("index=%d, val=%c\n", index, val)
	}
}
