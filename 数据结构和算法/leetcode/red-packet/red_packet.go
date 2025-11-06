package red_packet

import (
	"fmt"
	"math/rand"
	"time"
)

//a simple red packet algorithm
/*
	count 红包数量
	money 红包金额
*/
func RedPacket(count, money int)  {
	for i := 0; i < count; i++ {
		m := randomMoney(count - i, money)
		fmt.Printf("%d ", m)
		money -= m
	}
}

func randomMoney(remainCount, remainMoney int) int {
	if remainCount == 1 {
		return remainMoney
	}
	rand.Seed(time.Now().UnixNano())

	//min最小红包
	min := 1
	max := remainMoney / remainCount * 2
	//rand 0～max
	money := rand.Intn(max) + min
	return money
}
