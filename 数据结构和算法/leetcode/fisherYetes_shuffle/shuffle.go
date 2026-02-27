package fisherYetes_shuffle

import (
	"math/rand"
	"time"
)

//洗牌算法
func Shuffle(array []int)  {
	rand.Seed(time.Now().UnixNano())
	var i, j int
	for i = len(array) - 1; i > 0; i -- {
		j = rand.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
}
