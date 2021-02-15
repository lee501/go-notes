package main

import (
	"fmt"
)

//
//import (
//	"fmt"
//	"sync"
//)
//
func main() {
	//product chan size
	//producer := &Producer{make(chan *Workload, 100)}
	//Question2(producer)
	s := []int{2,7,9,3,1}
	fmt.Println(massage(s))
	decodeString("2[abc")
}
//
//// IWorkload 请勿修改接口
//type IWorkload interface {
//	// Process内包含一些耗时的处理，可能是密集计算或者外部IO
//	Process()
//}
//
//// IProducer 请勿修改接口
//type IProducer interface {
//	// Produce每次调用会返回一个IWorkload实例
//	// 当返回nil时表示已经生产完毕
//	Produce() IWorkload
//}
//
//// 问题2：请编写函数Question2的实现如下功能
//// 该函数输入一个IProducer实例，请调用Produce()方法生产IWorkload实例
//// 对于每个IWorkload实例，调用其Process()方法执行其业务功能
//// 重复调用Produce()直到生产完毕，对每个IWorkload进行Process()，
//// 所有工作完成后Question2函数返回
////
//// 要求：系统允许最大5个IWorkload同时进行Process()操作
////      请利用golang的并发特性完成此任务，单个并发的实现将不得分
////
//// 提示：请尽量使用规范的代码风格，使代码整洁易读
//// 提示：如果也实现了测试代码，请一并提交，将有利于分数评定
//
//func Question2(producer IProducer) {
//
//	// ====== 在这里书写代码 ====== //
//	if prod, ok := producer.(*Producer); ok {
//		go func() {
//			for i := 0; i< 100; i++ {
//				prod.Produce()
//			}
//			close(prod.W)
//		}()
//		pool := 5 //控制process数量
//		var wait sync.WaitGroup
//		wait.Add(5)
//		for i:= 1; i <= pool; i++ {
//			go func(i int) {
//				fmt.Println("线程", i)
//				for w := range prod.W {
//					w.Process()
//				}
//				wait.Done()
//			}(i)
//		}
//		wait.Wait()
//	}
//}
//
//type Producer struct {
//	W chan *Workload
//}
//
//func (p *Producer) Produce() IWorkload {
//	w :=  &Workload{}
//	p.W <- w
//	return w
//}
//
//type Workload struct {
//
//}
//
//func (w *Workload) Process() {
//	println("w")
//}

func massage(nums []int) int {
	l := len(nums)
	step := 2
	index := 0
	sum := 0
	for index < l {
		sum += nums[index]
		index += step
	}
	return sum
}


//func oddEvenList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	pre := head
//	cur := head.Next
//	for cur != nil && cur.Next != nil {
//		tmp := pre.Next
//		pre.Next = cur.Next
//		cur.Next = cur.Next.Next
//		pre.Next.Next = tmp
//		cur = cur.Next
//		pre = pre.Next
//	}
//	return head
//}

/*
给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例:

s = "3[a]2[bc]", 返回 "aaabcbc".
s = "3[a2[c]]", 返回 "accaccacc".
s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".
*/

func decodeString(s string) string {
  temp := []string{}
  for i, v := range s {
  	if v == '[' {
		fmt.Println(i, v)
	}
  }
  return ""
}