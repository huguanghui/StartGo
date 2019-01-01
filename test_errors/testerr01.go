package main

import (
	"fmt"
	"time"
)

// Myerr 定义一个私有的错误格式
type Myerr struct {
	when time.Time
	what string
}

func (m Myerr) Error() string {
	return fmt.Sprintf("%v: %v", m.when, m.what)
}

func oops() error {
	return Myerr{
		time.Date(2018, 12, 21, 8, 5, 30, 0, time.UTC),
		"the file system has gone away",
	}
}

func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}
