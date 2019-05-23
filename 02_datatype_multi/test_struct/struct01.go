package main

import "fmt"

// Ctest 用于测试结构类型中的隐式声明
type Ctest struct {
	Name string
}

func (c *Ctest) echo() {
	fmt.Println("Ctest Demo!")
}

// Ctell 用于测试结构类型中的隐式声明
type Ctell struct {
	Value string
}

func (c *Ctell) echo01() {
	fmt.Println("Ctell Demo!")
}

type abc struct {
	*Ctest
	*Ctell
}

func (a abc) Test() {
	a.echo()
	a.echo01()
}

func main() {
	var tmp abc
	tmp.Test()
}
