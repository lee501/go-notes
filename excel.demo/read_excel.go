package main

import (
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

func readExcel() {
	path, err := os.Getwd()
	fmt.Println(path)
	excel, err := excelize.OpenFile("excel.demo/api.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := excel.GetRows("API清单")
	if err != nil {
		fmt.Println(1)
		return
	}
	for i, row := range rows {
		fmt.Println(i)
		if i > 0 {
			for _, item := range row {
				fmt.Println(item)
			}
		}
	}
}
