package main

import (
	"errors"
	"fmt"
	"time"
)

type Slice []interface{}

func (s *Slice) Remove(value interface{}) error{
	for i, v := range *s {
		if isEqual(value, v) {
			*s = append((*s)[:i], (*s)[i+1:]...)
			return nil
		}
	}
	return errors.New("not exist")
}

func isEqual(a, b interface{}) bool {
	return a == b
}

func main() {

	strs := []string{"one", "two", "three"}

	for _, s := range strs {

		go func() {

			//time.Sleep(1 * time.Second)

			fmt.Printf("%s ", s)

		}()

	}

	time.Sleep(3 * time.Second)
}