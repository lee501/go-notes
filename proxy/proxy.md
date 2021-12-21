####go 设置代理

1. 通过http.client指定对象
```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
)

func main() {
	//通过设置环境变量，go会自动配置代理
	//os.Setenv("HTTP_PROXY", "http://127.0.0.1:9743")
	//os.Setenv("HTTPS_PROXY", "https://127.0.0.1:9743")
    urli := url.URL{}
    urlproxy, _ := urli.Parse("https://127.0.0.1:9743")
    c := http.Client{
        Transport: &http.Transport{
            Proxy: http.ProxyURL(urlproxy),
        },
    }
    if resp, err := c.Get("https://www.google.com"); err != nil {
    log.Fatalln(err)
    } else {
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Printf("%s\n", body)
    }
}

```
2. Socket5代理

```go
    package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/proxy"
)

func main() {
	// create a socks5 dialer
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9742", nil, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		os.Exit(1)
	}
	// setup a http client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	// set our socks5 as the dialer
	httpTransport.Dial = dialer.Dial
	if resp, err := httpClient.Get("https://www.google.com"); err != nil {
		log.Fatalln(err)
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%s\n", body)
	}
}
```
