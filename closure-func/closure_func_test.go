package main

import (
	"testing"
)

//func TestDeferPrint(t *testing.T) {
//	DeferPrint()
//}

func TestProcessFunc(t *testing.T) {
	s := ProcessFunc()
	for _, v := range s {
		v()
	}
}


func TestProcessFunc1(t *testing.T) {
	s := ProcessFunc1()
	for _, v := range s {
		v()
	}
}