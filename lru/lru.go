package lru

import (
	"bufio"
	"container/list"
	"net"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

func NewList() *list.List {
	return list.New()
}

type Packet struct {

}

type Channel struct {
	conn net.Conn    // WebSocket 连接
	send chan Packet
}

func (c *Channel) reader() {
	buf := bufio.NewReader(c.conn)
}