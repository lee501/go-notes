package red_packet

import (
	"strconv"
	"testing"
)

func TestRedPacket(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(strconv.FormatInt(int64(i), 10)+"随机", func(t *testing.T) {
			RedPacket(10, 500)
		})
	}
}
