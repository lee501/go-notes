package main

import (
	"fmt"
	"github.com/robfig/cron"
"log"
	"time"
)

func main() {
	i := 0
	c := cron.New()
	spec := "0 */1 * * * *"
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	c.Start()
	s := time.Second
	m := 2
	r := s * time.Duration(m)
	fmt.Println(r)
	select{}
}
