package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

var (
	size = 1024
)

type ICMP struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	Identifier  uint16
	SequenceNum uint16
}

func Ping(ip string) bool {
	conn, err := net.DialTimeout("ip:icmp", ip, 3*time.Second)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer conn.Close()

	fmt.Printf("正在 ping %s 具有 %d 字节的数据:\n", ip, size)

	//填充初始数据
	icmp := ICMP{
		Type:        8,
		Code:        0,
		Checksum:    0,
		Identifier:  1,
		SequenceNum: 1,
	}
	var buffer bytes.Buffer
	//写入icmp数据求校验和
	binary.Write(&buffer, binary.BigEndian, icmp)
	icmp.Checksum = CheckSum(buffer.Bytes())
	//清空buffer，写入求完校验和的icmp数据
	buffer.Reset()
	binary.Write(&buffer, binary.BigEndian, icmp)
	//发送请求
	start := time.Now()
	n, err := conn.Write(buffer.Bytes())
	if err != nil {
		return false
	}
	conn.SetWriteDeadline(time.Now().Add(time.Second * 2))
	buf := make([]byte, 1024)
	//读取结果
	n, err = conn.Read(buf)
	fmt.Println("消耗时间：", time.Now().Sub(start))
	if err != nil {
		return false
	}
	fmt.Println(buf[0:n])
	return true
}

func CheckSum(data []byte) uint16 {
	var sum uint32
	var length = len(data)
	var index int
	for length > 1 { // 溢出部分直接去除
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length == 1 {
		sum += uint32(data[index])
	}
	// CheckSum的值是16位，计算是将高16位加低16位，得到的结果进行重复以该方式进行计算，直到高16位为0
	/*
	   sum的最大情况是：ffffffff
	   第一次高16位+低16位：ffff + ffff = 1fffe
	   第二次高16位+低16位：0001 + fffe = ffff
	   即推出一个结论，只要第一次高16位+低16位的结果，再进行之前的计算结果用到高16位+低16位，即可处理溢出情况
	*/
	sum = uint32(sum>>16) + uint32(sum)
	sum = uint32(sum>>16) + uint32(sum)
	return uint16(^sum)
}

func main() {
	ip := "180.76.76.76"
	fmt.Println(Ping(ip))
}
