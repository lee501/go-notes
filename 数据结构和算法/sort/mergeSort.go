package main

/*
#cgo CFLAGS: -I/opt/homebrew/opt/libpcap/include
#cgo LDFLAGS: -L/opt/homebrew/opt/libpcap/lib -lpcap
#include <pcap.h>
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"os"
)

func main() {
	filter := "tcp and port 80" // 替换为你要验证的 BPF 过滤器
	var errbuf [C.PCAP_ERRBUF_SIZE]C.char

	// 获取所有设备
	var alldevs *C.pcap_if_t
	if C.pcap_findalldevs(&alldevs, &errbuf[0]) == -1 {
		fmt.Printf("Failed to find devices: %s\n", C.GoString(&errbuf[0]))
		os.Exit(1)
	}
	defer C.pcap_freealldevs(alldevs)

	// 选择第一个设备
	if alldevs == nil {
		fmt.Println("No devices found")
		os.Exit(1)
	}
	device := alldevs.name
	fmt.Println(device)
	// 初始化 pcap
	pcapHandle := C.pcap_open_live(device, 1600, 1, 100, &errbuf[0])
	if pcapHandle == nil {
		fmt.Printf("Failed to open pcap handle: %s\n", C.GoString(&errbuf[0]))
		os.Exit(1)
	}
	defer C.pcap_close(pcapHandle)

	// 设置 BPF 过滤器
	var bpfProgram C.struct_bpf_program
	compileResult := C.pcap_compile(pcapHandle, &bpfProgram, (*C.char)(C.CString(filter)), 1, 0)
	fmt.Println(compileResult)
	if compileResult == 0 {
		fmt.Println("BPF filter is valid:", filter)
		C.pcap_freecode(&bpfProgram)
	} else {
		fmt.Printf("BPF filter is invalid: %s\n", C.GoString(&errbuf[0]))
	}
}
