package main

import (
	"errors"
)

type Slice []interface{}

func (s *Slice) Remove(value interface{}) error {
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
