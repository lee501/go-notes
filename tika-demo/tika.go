package main

import (
	"context"
	"fmt"
	"github.com/google/go-tika/tika"
	"log"
	"os"
)

func main() {
	pwd, _ := os.Getwd()
	jar := pwd + "/tika-demo/tika-server-standard-2.9.1.jar"
	fmt.Println(jar)
	s, err := tika.NewServer(jar, "")
	if err != nil {
		log.Fatal(err)
	}
	err = s.Start(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer s.Stop()

	f, err := os.Open("/Users/lee/workspace/pcapfiles/test.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	client := tika.NewClient(nil, s.URL())
	body, err := client.Parse(context.Background(), f)
	fmt.Println(body)
}
