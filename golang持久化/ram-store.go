package main

import "fmt"

//内存存储，即定义一些数据结构，数组切片，图或者其他自定义结构，把需要持久化的数据存储在这些数据结构中
type Post struct {
	Id      int
	Content string
	Author  string
}

var (
	PostById map[int]*Post
	PostsByAuthor map[string][]*Post
)

func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}
func main() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	post2 := Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"}
	post4 := Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Printf("%p\n", &post1)
	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}

	for _, post := range PostsByAuthor["Pedro"] {
		fmt.Println(post)
	}
}
