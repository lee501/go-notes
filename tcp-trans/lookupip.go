package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	hosts := []string{"www.baidu.com", "www.sina.com"}
	chans := make(chan string, len(hosts))
	var wait sync.WaitGroup
	wait.Add(2)
	for _, host := range hosts {
		fmt.Println(host)
		go func(ch chan string, wait *sync.WaitGroup, host string) {
			ips, err := net.LookupIP(host)
			fmt.Println(err)
			//for _, ip := range ips {
			//	chans <- fmt.Sprintf("%s %s\r\n", ip.String(), host)
			//}
			chans <- fmt.Sprintf("%s %s\r\n", ips[0].String(), host)
			//}
			wait.Done()
		}(chans, &wait, host)
	}
	wait.Wait()
	close(chans)

	file, err := os.OpenFile("./hosts.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	w := bufio.NewWriter(file)
	for str := range chans {
		if !strings.Contains(string(content), str) {
			w.WriteString(str)
		}
	}
	w.Flush()
}
