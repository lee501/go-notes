package lru

import (
	"container/list"
)

/*
	LRU: 最近最少使用，核心思想是(如果数据最近被访问过，那么将来被访问的几率也更高)
		1. 新数据插入到链表头部；
		2. 每当缓存命中（即缓存数据被访问），则将数据移到链表头部；
		3. 当链表满的时候，将链表尾部的数据丢弃。
*/

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type Pair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}
func (lc *LRUCache) Get(key int) int {
	if elem, ok := lc.cache[key]; ok {
		lc.list.MoveToFront(elem)
		return elem.Value.(Pair).value
	}
	return -1
}

func (lc *LRUCache) Put(key, value int) {
	//hit cache, remove the elem to list head
	if elem, ok := lc.cache[key]; ok {
		elem.Value = Pair{key, value}
		lc.list.MoveToFront(elem)
	} else {
		//del cache and last elem when list full
		if lc.list.Len() >= lc.capacity {
			delete(lc.cache, lc.list.Back().Value.(Pair).key)
			lc.list.Remove(lc.list.Back())
		}
		lc.list.PushFront(Pair{key, value})
		lc.cache[key] = lc.list.Front()
	}
}
