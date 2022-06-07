package main

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

//Map底层的数据结构
type Map struct {
	mu     sync.Mutex
	read   atomic.Value // readOnly
	dirty  map[interface{}]*entry
	misses int
}

// Map.read 属性实际存储的是 readOnly
type readOnly struct {
	m       map[interface{}]*entry
	amended bool // 则会判断 read.readOnly 中的 amended 属性，他会告诉程序 dirty 是否包含 read.readOnly.m 中没有的数据；因此若存在，也就是 amended 为 true，将会进一步到 dirty 中查找数据
}

type entry struct {
	p unsafe.Pointer // *interface{}
}

/*
read 和 dirty 各自维护一套 key，key 指向的都是同一个 value。也就是说，只要修改了这个 entry，对 read 和 dirty 都是可见的
mu: 保护read 和dirty
read: 只读数据，指出并发读取 (atomic.Value 类型) 。如果需要更新 read，需要加锁保护数据安全。
read 实际存储的是 readOnly 结构体，内部是一个原生 map，amended 属性用于标记 read 和 dirty 的数据是否一致
dirty: 读写数据，非线性安全的原生 map。包含新写入的 key，并且包含 read 中所有未被删除的 key
misses: 统计有多少次读取 read 没有被命中。每次 read 读取失败后，misses 的计数加 1
*/
