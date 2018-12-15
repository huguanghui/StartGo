package main

import "fmt"

func main() {
	var b byte
	fmt.Println("请输入a,b,c,d")
	fmt.Scanf("%c", &b)

	switch b {
	case 'a':
		{
			fmt.Println("This is a!")
		}
	case 'b':
		{
			fmt.Println("This is b")
		}
	default:
		{
			fmt.Println("This is default!")
		}
	}
}
