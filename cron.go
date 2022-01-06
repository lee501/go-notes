package main

import (
	"github.com/robfig/cron/v3"
	"log"
)

func main() {
	i := 0
	c := cron.New()
	//每分钟
	spec := "*/1 * * * *"
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	c.Start()
	select{}
}
