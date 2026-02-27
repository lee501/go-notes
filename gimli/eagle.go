package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net"
)

var payloadSeparatorAsBytes = []byte("\n🐵🙈🙉\n")

type GorMessage struct {
	Meta []byte
	Data []byte
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8899")
	if err != nil {
		return
	}
	for {
		conn, err := listener.Accept()
		if err == nil {
			go handleConnection(conn)
			continue
		}
		break
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	var buffer bytes.Buffer

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				logrus.Warnf("[INPUT-TCP] connection error: %q", err)
			}
			break
		}

		if bytes.Equal(payloadSeparatorAsBytes[1:], line) {
			// find gor message separator
			gorMsgBytes := buffer.Bytes()
			// create new buffer
			buffer = bytes.Buffer{}

			dataLen := len(gorMsgBytes)
			if dataLen == 1 {
				fmt.Println("invalid gor message, received data length is too short", "data len:", dataLen)
				continue
			}
			//firstChar := gorMsgBytes[0]
			//if firstChar < '1' || firstChar > '9' {
			//	fmt.Println("invalid gor message, meta type is invalid", "meta type:", string(firstChar))
			//	continue
			//}

			// remove last '\n' charactor, from the monkey separator before monkeys
			gorMsgBytes = gorMsgBytes[:dataLen-1]
			var msg GorMessage
			msg.Meta, msg.Data = payloadMetaWithBody(gorMsgBytes)
			//fmt.Println(strconv.FormatInt(time.Now().UnixMilli(), 10) + "----$$$$----" + string(msg.Meta))
			//fmt.Println("---", string(gorMsgBytes))

		} else {
			buffer.Write(line)
		}
	}
}

func payloadMetaWithBody(payload []byte) (meta, body []byte) {
	if i := bytes.IndexByte(payload, '\n'); i > 0 && len(payload) > i+1 {
		meta = payload[:i]
		body = payload[i+1:]
		return
	}
	// we assume the message did not have meta data
	return nil, payload
}
