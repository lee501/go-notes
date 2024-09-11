package main

import (
	"fmt"
	"regexp"
)

func main() {
	//f, _ := os.Open("/Users/lee/workspace/pcapfiles/files/test.zip")
	//res, err := docconv.Convert(f, "zip", false)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(res.Body)
	fmt.Println([][]string{{"Url"}, {"BodyMap"}})
	for _, path := range append([][]string{{"Url"}, {"BodyMap"}}, []string{""}) {
		fmt.Println(path, len(path))
	}
	regexp.MustCompile("`?\\u0001123`")

	v := make([]string, 3)
	fmt.Println(len(v))

}
