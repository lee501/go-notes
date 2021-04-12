package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := int64(1)
	var p *int32
	p = (*int32)(unsafe.Pointer(&i))
	fmt.Printf("-----use unsafe pointer转化类型:----%#v\n", *p)
}
