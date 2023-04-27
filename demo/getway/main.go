package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

var average float64

type Response struct {
	Status int
	Usage  float64
}

func CpuUse(w http.ResponseWriter, r *http.Request) {
	usage := &Response{0, 0.0}
	w.Header().Set("content-type", "application/json")
	// CPU使用率
	percent, _ := cpu.Percent(time.Second, false)
	usage.Status = 1
	usage.Usage = percent[0]
	res, _ := json.Marshal(usage)
	w.Header().Set("content-type", "application/json")
	w.Write(res)
}
func main() {
	server := http.Server{Addr: "localhost:8080"}
	http.HandleFunc("/cpu/use", CpuUse)
	go cronTask()
	server.ListenAndServe()
}

func cronTask() {
	t := time.NewTicker(5 * time.Minute)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			t.Reset(5 * time.Minute)
			tmp := average
			percent, _ := cpu.Percent(time.Second, false)
			average = (percent[0] + tmp) / 2
		}
	}
}
