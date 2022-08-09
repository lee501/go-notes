package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	values := map[string]interface{}{
		"C2": "北京", "D2": 8, "E2": 10, "F2": 20}
	f, _ := excelize.OpenFile("/Users/lee/Downloads/report_template.xlsx")
	defer func() {
		f.Close()
	}()
	for k, v := range values {
		err := f.SetCellValue("index", k, v)
		fmt.Println(err, k)
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
	if err := f.SaveAs("demo.xlsx"); err != nil {
		fmt.Println(err)
	}
}
