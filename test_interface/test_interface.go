package main

import "fmt"

// 定义接口
type Phone interface {
	call()
}

// 定义结构体
type NokiaPhone struct {
}

// 实现接口方法
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("NokiaPhone is calling!")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()
}
