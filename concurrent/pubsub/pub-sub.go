package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	p := NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	all := p.Subscribe()
	golang := p.SubscribeTopic(t)

	p.Publish("hello world")
	p.Publish("golang")

	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	}()

	time.Sleep(3 * time.Second)
}

func t(v interface{}) bool {
	if s, ok := v.(string); ok {
		return strings.Contains(s, "golang")
	}
	return false
}
