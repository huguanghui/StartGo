// Go中变量的使用的三种方式
// 1）带变量类型
// 2）根据值自动推导
// 3）省略var
// 多变量申明
// 全局变量
package main

import "fmt"

var g_test = "abcdefg"

var (
	g_p1 = 1
	g_p2 = "tome"
)

func main() {
	// 1.指定类型
	var i int
	i = 10
	fmt.Println("i=", i)
	// 2.根据值自动推导
	var num = 10.11
	fmt.Println("mum=", num)
	// 3.省略var
	name := "tom"
	fmt.Println("name=", name)
	// 多变量申明
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)
	var size, ss, sex = 10, "tom", 23
	fmt.Println("size=", size, ", ss=", ss, ",sex=", sex)
	a1, a2, a3 := 1, "peter", 23.2
	fmt.Println("a1=", a1, ",a2=", a2, ",a3=", a3)

	fmt.Println("g_test=", g_test)
	fmt.Println("g_p1=", g_p1, ", g_p2=", g_p2)
	fmt.Printf("Type:%T g_p1=%d, g_p2=%s\n", g_p1, g_p1, g_p2)
}
