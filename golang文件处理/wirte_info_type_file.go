package golang文件处理

import (
	"bufio"
	"encoding/json"
	"os"
)

var filepath = "../test.info"

type Info struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func saveInfo(filepath string) {
	info := Info{Name: "lee", Age: 22, Sex: "boy"}
	data, _ := json.Marshal(&info)
	file, _ := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	buf := bufio.NewWriter(file)
	buf.WriteString(string(data))
	buf.Flush()
}
