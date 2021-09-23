package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"sync"
)

type HttpHost struct {
	Host string
}

func main() {
	hosts := []string{"www.baidu.com", "www.sina.com"}
	url := "192.168.1.20"
	buffer := bytes.NewBufferString("")
	fmt.Fprintf(buffer, `func FindProxyForURL(url, host) {`)
	for _, host := range hosts {
		fmt.Fprintf(buffer, `if (dnsDomainIs(host, "%s"))`, host)
		fmt.Fprintf(buffer, `{ return "PROXY %s"; }`, url)
	}
	fmt.Fprintf(buffer, `return "DIRECT" }`)
	fmt.Println(buffer.String())

	var wait sync.WaitGroup
	wait.Add(1)
	go func() {
		s := "first go routine"
		fmt.Println(s)
		go func() {
			fmt.Println(s + "inner")
			wait.Done()
		}()
	}()
	wait.Wait()

	url = "https://www.jianshu.com/favicon.ico"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))

	sys := runtime.GOOS
	if sys == "linux" {
		fmt.Println("linux")
	} else if sys == "windows" {
		fmt.Println("windows")
	} else {
		fmt.Println(sys)
	}

	httphosts := []*HttpHost{&HttpHost{"http://www.baidu.com"}, &HttpHost{"jira.zshield.int"}, &HttpHost{"http://www.baidu.com"}, &HttpHost{"jira.zshield.int"}}
	re := removeRepeathost(httphosts)
	fmt.Println(len(re))
}

func removeRepeathost(hosts []*HttpHost) []*HttpHost {
	result := []*HttpHost{}
	for i := 0; i < len(hosts); i++ {
		repeat := true
		for j := i + 1; j < len(hosts); j++ {
			if hosts[i].Host == hosts[j].Host {
				repeat = false
				break
			}
		}
		if repeat {
			result = append(result, hosts[i])
		}
	}
	return result
}
