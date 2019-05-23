// main.go
package main

import (
	"flag"
	"fmt"
)

func main() {
	name, ico := parseFlag()
	fmt.Println(name, ico)
}

func parseFlag() (name string, ico string) {

	nameInput := flag.String("name", "GoDoc", "Set docset name")
	icoInput := flag.String("icon", "", "Docset icon .png path")

	flag.Parse()
	name = *nameInput
	ico = *icoInput

	return
}
