package main

import (
	"fmt"
	"time"

	cache "github.com/patrickmn/go-cache"
)

func myfunc(a string) {
	fmt.Println("myfunc: ", a)
}

func main() {
	c := cache.New(5*time.Minute, 10*time.Minute)
	c.Set("foo", "bar", cache.DefaultExpiration)
	c.Set("baz", 42, cache.NoExpiration)

	foo, found := c.Get("foo")
	if found {
		// fmt.Println(foo)
		myfunc(foo.(string))
	}
}
