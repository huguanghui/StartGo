package main

import "fmt"

func main() {

	var age int

	fmt.Println("请输入年龄:")
	fmt.Scanln(&age)

	if age > 20 {
		fmt.Println("you are so old!")
	} else {
		fmt.Println("you are young!")
	}
}
