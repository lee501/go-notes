package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	filePath = "/app/data/"
	docx     = "docx"
	pdf      = "pdf"
)

func main() {
	v, _ := strconv.ParseFloat("18.66", 64)
	fmt.Println(v)
	//pdf := gofpdf.New("P", "mm", "A4", "")
	//pdf.AddPage()
	//pdf.SetFont("Arial", "B", 16)
	//pdf.Cell(40, 10, "Hello World")
	//err := pdf.OutputFileAndClose("hello.pdf")
	//if err != nil {
	//	panic(err)
	//}
	//f, _ := os.Open("/Users/lee/workspace/xin.pdf")
	//
	//var buff bytes.Buffer
	//var b = make([]byte, 512)
	//for {
	//	n, err := f.Read(b)
	//	if n > 0 {
	//		buff.Write(b)
	//	}
	//	if err == io.EOF || n == 0 {
	//		break
	//	}
	//}
	//out, _ := os.Create("ttttt.pdf")
	//buff.WriteTo(out)
	s := fmt.Sprintf("%s.%s", "xin", "pdf")
	fmt.Println(strings.Split(s, ".")[1])
	switch "pdf" {
	case docx:
		fmt.Println(docx)
	case pdf:
		fmt.Println(pdf)
	default:
		fmt.Println("unknown")
	}
	var i any = ""
	fmt.Println(i != "")
}
