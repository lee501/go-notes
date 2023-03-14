package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/antchfx/xmlquery"
)

func main() {
	var buff bytes.Buffer
	path := "/Users/lee/workspace/go/go-notes/program_demo/test.xlsx"
	zr, err := zip.OpenReader(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, f := range zr.Reader.File {
		rc, err := f.Open()
		if err != nil {
			fmt.Println(err)
		}

		content, err := ioutil.ReadAll(rc)
		if err != nil {
			fmt.Println("2", err)
			continue
		}
		doc, err := xmlquery.Parse(strings.NewReader(string(content)))
		if err != nil {
			fmt.Println("1", err)
			continue
		}

		for _, n := range xmlquery.Find(doc, "//t") {
			buff.WriteString(n.InnerText())
			buff.WriteString(" ")
		}
		for _, n := range xmlquery.Find(doc, "//v") {
			buff.WriteString(n.InnerText())
			buff.WriteString(" ")
		}
		rc.Close()
	}
	fmt.Println(buff.String())
}
