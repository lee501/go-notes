package main

import (
	"context"
	"fmt"
	"github.com/gen2brain/go-fitz"
	"github.com/google/go-tika/tika"
	"os"
)

func main() {
	//filePath := "/Users/lee/Downloads/oakai.pdf"
	////调用函数ReadPdf解析pdf文件
	//content, err := ReadPdf(filePath) // Read local pdf file
	//if err != nil {
	//	panic(err)
	//}
	//f, _ := os.Create("content.html")
	//f.WriteString(content)
	//f.Close()
	//var buff bytes.Buffer
	//pd, _ := os.Open(filePath)
	//
	//pdfReader, _ := pdf2.NewPdfReader(pd)
	//numPages, err := pdfReader.GetNumPages()
	//for i := 0; i < numPages; i++ {
	//	pageNum := i + 1
	//	page, err := pdfReader.GetPage(pageNum)
	//	if err != nil {
	//		return
	//	}
	//	ex, err := extractor.New(page)
	//	if err != nil {
	//		return
	//	}
	//	pt, _, _, _ := ex.ExtractPageText()
	//
	//	buff.WriteString(pt.Text())
	//}
	//te, _ := os.Create("aokai.txt")
	//str := strconv.Quote(buff.String())
	//te.WriteString(str)
	doc, err := fitz.New("/Users/lee/Downloads/en.pdf")
	if err != nil {
		fmt.Println("无法打开 PDF 文件：", err)
		return
	}
	defer doc.Close()

	// 获取页面数量
	pageCount := doc.NumPage()

	// 遍历每个页面并提取文本
	for pageNum := 0; pageNum < pageCount; pageNum++ {
		text, err := doc.Text(pageNum)
		if err != nil {
			fmt.Println("提取文本时出错：", err)
			continue
		}
		fmt.Println("第", pageNum+1, "页的文本：", text)
	}

	var i interface{}
	i = ""
	fmt.Println(i.(string) + "1")
}

//解析PDF文件
func ReadPdf(path string) (string, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}
	client := tika.NewClient(nil, "http://127.0.0.1:9998")
	return client.Parse(context.TODO(), f)
}
