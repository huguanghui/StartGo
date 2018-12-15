package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	hellohandle := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, huguanghui!\n")
	}

	http.HandleFunc("/hello", hellohandle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
