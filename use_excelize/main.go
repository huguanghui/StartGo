package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	// createXLSX()
	readXLSX()
}

func createXLSX() {
	f := excelize.NewFile()
	// Create New sheet
	index := f.NewSheet("Sheet2")
	f.SetCellValue("Sheet2", "A2", "Hello World!")
	f.SetCellValue("Sheet1", "B2", 100)
	f.SetActiveSheet(index)
	err := f.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func readXLSX() {
	f, err := excelize.OpenFile("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given sheetname and axis
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// get all rows
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

