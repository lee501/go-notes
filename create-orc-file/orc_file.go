package main

import (
	"fmt"
	"github.com/scritchley/orc"
	"log"
	"os"
)

func main() {
	file, err := os.Create("test.orc")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	schema, err := orc.ParseSchema("struct<col1:string,col2:string>")
	if err != nil {
		log.Fatal(err)
	}
	writer, err := orc.NewWriter(file, orc.SetSchema(schema))
	if err != nil {
		log.Fatal(err)
	}
	err = writer.Write([]interface{}{"张三", "13021015695"}...)
	if err != nil {
		log.Fatal(err)
	}

	// 完成写入
	err = writer.Close()
	if err != nil {
		log.Fatal(err)
	}

	r, err := orc.Open("test.orc")
	if err != nil {
		log.Fatal(err)
	}
	selected := r.Schema().Columns()
	c := r.Select(selected...)
	defer c.Close()
	vals := make([]interface{}, len(selected))
	ptrVals := make([]interface{}, len(selected))
	strVals := make([]string, len(selected))
	for i := range vals {
		ptrVals[i] = &vals[i]
	}
	for c.Stripes() {
		for c.Next() {
			err := c.Scan(ptrVals...)
			if err != nil {
				log.Fatal(err)
			}
			for i := range ptrVals {
				strVals[i] = fmt.Sprint(ptrVals[i])
				log.Println(strVals[i])
			}
		}
	}
	if err := c.Err(); err != nil {
		log.Fatal(err)
	}
}
