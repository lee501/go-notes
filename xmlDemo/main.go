package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Peoples struct {
	//xml根节点
	XMLName xml.Name `xml:"peoples"`
	//属性
	Version string `xml:"version,attr"`
	//子节点
	Peos []People `xml:"people"`
}

type People struct {
	XMLName xml.Name `xml:"people"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Address string   `xml:"address"`
}

func main() {
	people := new(Peoples)
	xmlfile, _ := ioutil.ReadFile("./peoples.xml")
	xml.Unmarshal(xmlfile, people)
	fmt.Println(people)

	//生成xml文件
	p := People{Id: 1, Name: "lee", Address: "北京"}
	//格式化
	b, _ := xml.MarshalIndent(p, "", "		")
	//追加头信息
	b = append([]byte(xml.Header), b...)
	ioutil.WriteFile("./people.xml", b, 0666)
}
