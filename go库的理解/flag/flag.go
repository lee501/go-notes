package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	name = flag.String("name", "Tom", "Input your name")
	age = flag.Int("age", 18, "Input your age")
	f = flag.Bool("isVIP", false, "Is VIP")
	postCode int
)

func init() {
	flag.IntVar(&postCode, "postcode", 1234, "Input your post code")
}

func main() {
	flag.Parse()

	fmt.Println("name:", *name)
	fmt.Println("age:", *age)
	fmt.Println("VIP:", *f)
	fmt.Println("postCode:", postCode)

	fmt.Println("tail:", flag.Args())
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())

	args := os.Args
	fmt.Println("Args:", args)

	//paramCnt := flag.NArg()
	//for cnt := 0; cnt < paramCnt; cnt++ {
	//	fmt.Println(flag.Arg(cnt))
	//}
}