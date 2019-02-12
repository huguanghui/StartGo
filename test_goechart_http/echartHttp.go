// echartHttp.go
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/chenjiandongx/go-echarts/charts"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	nameItem := []string{"衬衫", "牛仔裤", "运动裤", "袜子", "冲锋衣", "羊毛衫"}
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar示意图"})
	bar.AddXAxis(nameItem).
		AddYAxis("商家A", []int{23, 26, 37, 12, 21, 9}).
		AddYAxis("商家B", []int{22, 23, 20, 18, 27, 33})
	f, err := os.Create("bar.html")
	if err != nil {
		fmt.Println(err)
	}
	bar.Render(w, f)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
