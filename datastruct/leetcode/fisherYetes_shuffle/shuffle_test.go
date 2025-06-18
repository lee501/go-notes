package fisherYetes_shuffle

import (
	"strconv"
	"testing"
)

func TestShuffle(t *testing.T) {
	array := []int{1,2,3,4,5,6,7,8,9}
	//十次洗牌
	for i := 0; i < 10; i++ {
		t.Run(strconv.FormatInt(int64(i), 10) + "洗牌", func(t *testing.T) {
			Shuffle(array)
			t.Log(array)
		})
	}
}
