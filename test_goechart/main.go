// main.go
package main

import (
	"log"
	"os"

	"github.com/chenjiandongx/go-echarts/charts"
)

func main() {
	nameItem := []string{"衬衫", "牛仔裤", "运动裤", "袜子", "冲锋衣", "羊毛衫"}
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-实例图"})
	bar.AddXAxis(nameItem).
		AddYAxis("商家A", []int{20, 30, 40, 10, 24, 36}).
		AddYAxis("商家B", []int{35, 14, 25, 60, 44, 23})
	f, err := os.Create("bar.html")
	if err != nil {
		log.Println(err)
	}
	bar.Render(f)
}
