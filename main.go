package main

import "fmt"
import "time"

// 结构体定义
type book struct {
	name string
	num  int
}

func (f book) Read() {
	fmt.Println(f)
}

type People struct {
	Age  int
	Name string
}

// 方法定义
func (f People) change() {
	f.Age += 10
	fmt.Println(f)
}

func (f People) Read() {
	fmt.Println(f)
}

func (f *People) change01() {
	f.Age += 10
	fmt.Println(*f)
}

// 函数定义
func test() {
	fmt.Println("hello world")
}

func funret() (a, b string) {
	//return "a", "b"
	a = "hello"
	b = "test"
	return a, b
}

// 协程
func pthread() {
	time.Sleep(10 * time.Second)
	fmt.Println("hello a")
}

// 接口定义
type TestRead interface {
	/* TODO: add methods */
	Read()
}

// 反射
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	// 变量定义(整型,浮点型,数组,切片,map)
	tmp := 1
	ta := 0.1
	var book1 book
	book1.name = "Go 语言"
	book1.num = 4
	// 数组
	abc := [5]int{1, 2, 3, 4}
	// 切片
	// 1. 创建切片
	sl := make([]int, 10)
	sl[0] = 12
	// Map集合
	// 1.声明
	var map_var map[string]string
	// 2.定义
	map_var = make(map[string]string)
	map_var["France"] = "巴黎"
	map_var["Italy"] = "罗马"
	map_var["Japan"] = "东京"
	map_var["India"] = "新德里"

	// 3.删除元素
	delete(map_var, "India")

	for country := range map_var {
		fmt.Println(country, "首都是 ", map_var[country])
	}

	fmt.Printf("书名: %s\n", book1.name)
	fmt.Printf("数量: %d\n", book1.num)

	// 判断语句
	if tmp == 1 {
		for _, i := range sl[0:3] {
			fmt.Println(i, " ")
		}
		fmt.Printf("len=%d cap=%d slice=%v\n", len(sl), cap(sl), sl)
		for i := 0; i < len(abc); i++ {
			fmt.Println(abc[i], " ")
		}
		fmt.Println("test01", ta, abc[0])
	} else {
		fmt.Println("test02", ta)
	}
	cnt := 10
	for cnt > 0 {
		fmt.Printf("Cnt :%d\n", cnt)
		cnt--
	}
	switch cnt {
	case 10:
		fmt.Println("cnt == 10")
	default:
		fmt.Println("cnt is ", cnt)
	}

	// 函数调用
	test()

	// 匿名函数
	f := func() {
		fmt.Println("Lambdas function")
	}
	f()

	a, b := funret()
	fmt.Println(a, b)

	// 结构体的使用
	p1 := People{25, "Jake"}
	fmt.Println(p1)
	p2 := People{Name: "Tom", Age: 23}
	fmt.Println(p2)
	p2.change()
	fmt.Println(p2)
	p2.change01()
	fmt.Println(p2)
	p3 := new(People)
	p3.Name = "Stone"
	p3.Age = 22
	fmt.Println(*p3)

	// channel使用
	c := make(chan int)
	go func() {
		c <- 1
	}()
	go func() {
		fmt.Println(<-c)
	}()

	// 带缓存的channel
	bc := make(chan int, 2)
	go func() {
		bc <- 1
		bc <- 2
	}()
	go func() {
		fmt.Println(<-bc)
		fmt.Println(<-bc)
	}()

	do(25)
	do("hello world!")
	do(1.24)
	// 协程使用
	// go pthread()
	// fmt.Println("hello b")
	// time.Sleep(30 * time.Second)
}
