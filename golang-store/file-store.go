package main

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	//PlainStore()
	CsvStore()
}

//纯文本存储
//只需要将保存的字符串写入文本保存即可。golang提供了ioutil库用于读写文件，也提供了os相关的文件创建，写入，保存的工具函数
func PlainStore() {
	data := []byte("plain text store")
	err := ioutil.WriteFile("data", data, 0644)
	if err != nil {
		panic(err)
	}

	read, _ := ioutil.ReadFile("data")
	fmt.Println(string(read))
}

//csv文件是一种以逗号分割单元数据的文件, 需要通过os创建一个文件句柄，然后调用相关的csv函数读写数据：
type Post struct {
	Id      int
	Content string
	Author  string
}

func CsvStore() {
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
		Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"},
		Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		fmt.Println(line)
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	//读取
	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content:item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}

//gob二进制存储
func GobStore(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	//NewEncoder方法创建一个encoder对象，
	encoder := gob.NewEncoder(buffer)
	//然后对数据进行二进制编码
	err := encoder.Encode(data)
	if err != nil{
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil{
		panic(err)
	}
}

func load(data interface{}, filename string){
	//先读取文件的内容
	raw, err := ioutil.ReadFile(filename)
	if err != nil{
		panic(err)
	}
	//然后把这个二进制内容转换成一个buffer对象
	buffer := bytes.NewBuffer(raw)
	//NewDecoder. 最后再解码
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil{
		panic(err)
	}
}