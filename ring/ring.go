package ring

import (
	"container/ring"
	"fmt"
)

//双向链表
func testRing() {
	//初始化Ring结构
	r := ring.New(5)
	r.Value = 0
	r.Next().Value = 1
	r.Next().Value = 2
	r.Prev().Value = 3
	r.Prev().Value = 4

	//遍历循环链表
	r.Do(func(i interface{}) {
		fmt.Println(i)
	})
	//移动
	_ = r.Move(1).Value

	//添加
	r1 := ring.New(1)
	r1.Value = 5
	r.Next().Link(r1)

	//删除 n%r.len()
	r.Unlink(1)
}
