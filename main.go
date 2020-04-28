package main

import (
	"fmt"

	"github.com/huguanghui/StartGo/ui"
)

func main() {
	n := ui.UiPrintf()
	n = ui.UiTest()
	fmt.Printf("%d\n", n)
}
