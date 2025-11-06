package main

import (
	"fmt"

	"golang.org/x/exp/slices" // Go 1.18 及以上版本的 slices 包
)

// 位运算符用于websocket协议的判断
var WebsocketOpCodes []byte = []byte{
	0,  // continuation frame
	1,  // text
	2,  // binary
	8,  // close
	9,  // ping
	10, // pong
}

func processWebSocketFrame(data []byte) error {
	if len(data) == 0 {
		return fmt.Errorf("data is empty")
	}

	opCode := data[0] & 0x0F
	if !slices.Contains(WebsocketOpCodes, opCode) {
		return fmt.Errorf("unsupported opcode: %d", opCode)
	}

	// 处理有效的操作码...
	return nil
}

func demoOptCode() {
	data := []byte{9, 11} // 示例数据，代表一个文本帧
	err := processWebSocketFrame(data)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Frame processed successfully.")
	}
}
