package main

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/xuri/excelize/v2"
)

var export_path = "/Users/lee/workspace/excel/temp/"

func init() {
	_ = os.Mkdir(export_path, fs.ModePerm)
}
func main() {
	f, _ := excelize.OpenFile("/Users/lee/Downloads/report.xlsx")
	defer func() {
		f.Close()
	}()
	cells := []string{"C2", "D2", "E2", "F2", "G2", "H2", "I2", "J2"}
	values := []interface{}{"北京", 40, 20, 10, 5, 20, 4, 8}
	for i, cell := range cells {
		err := f.SetCellValue("index", cell, values[i])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	//if err := f.AddChart("Sheet1", "E1", `{
	//   "type": "colStacked",
	//   "series": [
	//   {
	//       "name": "Sheet1!$A$2",
	//       "categories": "Sheet1!$B$1:$D$1",
	//       "values": "Sheet1!$B$2:$D$2"
	//   },
	//   {
	//       "name": "Sheet1!$A$3",
	//       "categories": "Sheet1!$B$1:$D$1",
	//       "values": "Sheet1!$B$3:$D$3"
	//   },
	//   {
	//       "name": "Sheet1!$A$4",
	//       "categories": "Sheet1!$B$1:$D$1",
	//       "values": "Sheet1!$B$4:$D$4"
	//   }],
	//   "title":
	//   {
	//       "name": "Fruit 3D Clustered Column Chart"
	//   }
	//}`); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//// Save spreadsheet by the given path.
	if err := f.SaveAs("/Users/lee/workspace/excel/demo.xlsx"); err != nil {
		fmt.Println(err)
	}
}
