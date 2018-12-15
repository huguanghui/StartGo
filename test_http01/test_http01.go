package main

import (
	"fmt"
	"net/http"
)

func main() {
	// res, err := http.Get("https://www.easy-mock.com/mock/5befe2e91aaa7772d6759b75/hgh/frontendguide/list")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// robots, err := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)
	fmt.Printf("%T\n", http.ErrBodyNotAllowed)
}
