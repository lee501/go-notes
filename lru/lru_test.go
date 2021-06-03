package lru

import (
	"fmt"
	"testing"
)

func TestNewList(t *testing.T) {
	l := NewList()
	fmt.Println(l.Front())
}