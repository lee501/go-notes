package main

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// ==================== 模拟 pcapPkt ====================

type pcapPkt struct {
	data []byte
}

var pcapPktPool = &sync.Pool{
	New: func() interface{} {
		return &pcapPkt{
			data: make([]byte, 1500),
		}
	},
}

func (pkt *pcapPkt) release() {
	pcapPktPool.Put(pkt)
}

// ==================== 模拟 TcpIpPacket ====================

type TcpIpPacket struct {
	Payload []byte
}

var tcpIpPacketPool = &sync.Pool{
	New: func() interface{} {
		return new(TcpIpPacket)
	},
}

func (pkt *TcpIpPacket) release() {
	pkt.Payload = nil
	tcpIpPacketPool.Put(pkt)
}

// ==================== 模拟 TcpMessage ====================

type TcpMessage struct {
	packets []*TcpIpPacket
}

var tcpMessagePool = &sync.Pool{
	New: func() interface{} {
		return new(TcpMessage)
	},
}

func (tcpMsg *TcpMessage) Release() {
	for _, pkt := range tcpMsg.packets {
		pkt.release()
	}
	tcpMsg.packets = nil
	tcpMessagePool.Put(tcpMsg)
}

// 原始实现：直接引用
func (tcpMsg *TcpMessage) getMessageData() [][]byte {
	res := [][]byte{}
	for _, pkt := range tcpMsg.packets {
		res = append(res, pkt.Payload)
	}
	return res
}

// 修复实现：深拷贝
func (tcpMsg *TcpMessage) getMessageDataWithCopy() [][]byte {
	res := make([][]byte, 0, len(tcpMsg.packets))
	for _, pkt := range tcpMsg.packets {
		payloadCopy := make([]byte, len(pkt.Payload))
		copy(payloadCopy, pkt.Payload)
		res = append(res, payloadCopy)
	}
	return res
}

// ==================== 模拟 GorMessage ====================

type GorMessage struct {
	Data [][]byte
}

// ==================== 模拟 extractTcpPacket ====================

func extractTcpPacket(ethData []byte, tcpPkt *TcpIpPacket, payloadLen int) {
	tcpPkt.Payload = ethData[:payloadLen]
}

// ==================== 测试 ====================

func runTest(useCopy bool, totalMessages int) (successCount, failCount int64) {
	msgChan := make(chan *GorMessage, 100)

	var success, fail atomic.Int64

	// 模拟 pcap 抓包 channel（多个）
	pcapChanCount := 4
	pcapChans := make([]chan *pcapPkt, pcapChanCount)
	for i := 0; i < pcapChanCount; i++ {
		pcapChans[i] = make(chan *pcapPkt, 100)
	}

	var producerWg sync.WaitGroup
	var workerWg sync.WaitGroup

	// 生产者：模拟 libpcap 抓包，生成 pcapPkt
	producerWg.Add(1)
	go func() {
		defer producerWg.Done()
		for i := 0; i < totalMessages; i++ {
			pkt := pcapPktPool.Get().(*pcapPkt)
			if pkt.data == nil || cap(pkt.data) < 1500 {
				pkt.data = make([]byte, 1500)
			}

			httpData := fmt.Sprintf("GET /api/test/%d HTTP/1.1\r\nHost: example.com\r\n\r\n", i)
			copy(pkt.data, httpData)

			// 分发到不同的 channel
			pcapChans[i%pcapChanCount] <- pkt
		}
		// 关闭所有 pcap channel
		for i := 0; i < pcapChanCount; i++ {
			close(pcapChans[i])
		}
	}()

	// 多个 worker：模拟 pktWorker 处理 pcapPkt
	workersPerChan := 2
	for chanIdx := 0; chanIdx < pcapChanCount; chanIdx++ {
		for w := 0; w < workersPerChan; w++ {
			workerWg.Add(1)
			go func(ch chan *pcapPkt) {
				defer workerWg.Done()

				for pkt := range ch {
					// 模拟 extractTcpPacket
					tcpPkt := tcpIpPacketPool.Get().(*TcpIpPacket)
					payloadLen := bytes.IndexByte(pkt.data, 0)
					if payloadLen <= 0 {
						payloadLen = 100
					}
					extractTcpPacket(pkt.data, tcpPkt, payloadLen)

					// 模拟 pkt.release()（关键：底层数组被放回池中）
					pkt.release()

					// 模拟 TCP 重组
					tcpMsg := tcpMessagePool.Get().(*TcpMessage)
					tcpMsg.packets = append(tcpMsg.packets, tcpPkt)

					// 模拟 toGorMessage
					var data [][]byte
					if useCopy {
						data = tcpMsg.getMessageDataWithCopy()
					} else {
						data = tcpMsg.getMessageData()
					}
					gorMsg := &GorMessage{Data: data}

					// 发送到 output channel
					msgChan <- gorMsg

					// 模拟 tcpMessage.Release()
					tcpMsg.Release()
				}
			}(pcapChans[chanIdx])
		}
	}

	// 等待 worker 完成后关闭 msgChan
	go func() {
		producerWg.Wait()
		workerWg.Wait()
		close(msgChan)
	}()

	// 多个消费者：模拟 output worker
	var consumerWg sync.WaitGroup
	consumerCount := 2
	for c := 0; c < consumerCount; c++ {
		consumerWg.Add(1)
		go func() {
			defer consumerWg.Done()

			for msg := range msgChan {
				time.Sleep(time.Microsecond * 10) // 模拟网络发送延迟

				if len(msg.Data) > 0 && len(msg.Data[0]) >= 4 {
					if bytes.HasPrefix(msg.Data[0], []byte("GET ")) {
						success.Add(1)
					} else {
						fail.Add(1)
						if fail.Load() <= 5 {
							sample := msg.Data[0]
							if len(sample) > 40 {
								sample = sample[:40]
							}
							fmt.Printf("  数据损坏: %q\n", string(sample))
						}
					}
				} else {
					fail.Add(1)
					if fail.Load() <= 5 {
						fmt.Printf("  数据为空或太短\n")
					}
				}
			}
		}()
	}

	consumerWg.Wait()
	return success.Load(), fail.Load()
}

func main() {
	// totalMessages := 50000

	// fmt.Println("==================== 测试1: 不拷贝（原始实现） ====================")
	// fmt.Println("模拟: 多 pcapChan + 多 pktWorker + 多 outputWorker")
	// fmt.Println("模拟: tcpPkt.Payload = pkt.data[:n] (共享底层数组)")
	// fmt.Println("模拟: pkt.release() 后 pkt.data 被其他 worker 复用")
	// fmt.Println()
	// success1, fail1 := runTest(false, totalMessages)
	// fmt.Printf("总消息数: %d\n", totalMessages)
	// fmt.Printf("成功: %d (%.2f%%)\n", success1, float64(success1)*100/float64(totalMessages))
	// fmt.Printf("失败: %d (%.2f%%)\n", fail1, float64(fail1)*100/float64(totalMessages))

	// fmt.Println()
	// fmt.Println("==================== 测试2: 深拷贝（修复后） ====================")
	// success2, fail2 := runTest(true, totalMessages)
	// fmt.Printf("总消息数: %d\n", totalMessages)
	// fmt.Printf("成功: %d (%.2f%%)\n", success2, float64(success2)*100/float64(totalMessages))
	// fmt.Printf("失败: %d (%.2f%%)\n", fail2, float64(fail2)*100/float64(totalMessages))

	// fmt.Println()
	// fmt.Println("==================== 结论 ====================")
	// if fail1 > 0 && fail2 == 0 {
	// 	fmt.Println("✅ 问题已复现，深拷贝修复有效！")
	// 	fmt.Println()
	// 	fmt.Println("修复方案:")
	// 	fmt.Println("```go")
	// 	fmt.Println("func (tcpMsg *TcpMessage) getMessageData() [][]byte {")
	// 	fmt.Println("    res := make([][]byte, 0, len(tcpMsg.packets))")
	// 	fmt.Println("    for _, pkt := range tcpMsg.packets {")
	// 	fmt.Println("        payloadCopy := make([]byte, len(pkt.Payload))")
	// 	fmt.Println("        copy(payloadCopy, pkt.Payload)")
	// 	fmt.Println("        res = append(res, payloadCopy)")
	// 	fmt.Println("    }")
	// 	fmt.Println("    return res")
	// 	fmt.Println("}")
	// 	fmt.Println("```")
	// } else if fail1 == 0 {
	// 	fmt.Println("⚠️ 未复现，sync.Pool 在当前环境下复用不够激进")
	// 	fmt.Println("   但生产环境高流量下必然出现此问题")
	// 	fmt.Println("   日志 'gor msg data is empty' 就是证据")
	// }
	sql := "SELECT table,min(partition) as partition,if(table='access_raw_local',sum(data_compressed_bytes)*10,sum(data_compressed_bytes)) as sumData FROM system.parts WHERE disk_name = 'default' GROUP BY table ORDER BY sumData DESC FORMAT JSON"
	res := url.QueryEscape(sql)
	fmt.Println(res)

	sqlQuery := strings.ReplaceAll("SELECT table,min(partition) as partition,if(table='access_raw_local',sum(data_compressed_bytes)*10,sum(data_compressed_bytes)) as sumData FROM system.parts WHERE disk_name = 'default' GROUP BY table ORDER BY sumData DESC FORMAT JSON", " ", "%20")
	fmt.Println(sqlQuery)
}
