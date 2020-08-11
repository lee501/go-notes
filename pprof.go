package main

import (
	"encoding/json"
	"fmt"
)

//func main() {
//	var isCPUPprof bool
//	flag.BoolVar(&isCPUPprof, "cpu", false, "cpu pprof")
//	flag.Parse()
//	if isCPUPprof {
//		file, err := os.Create("./cpu.pprof")
//		if err != nil {
//			fmt.Printf("create cpu pprof failed, err:%v\n", err)
//			return
//		}
//		pprof.StartCPUProfile(file)
//		defer pprof.StopCPUProfile()
//	}
//
//	s := file_process.ReadFileByString("/Users/richctrl/Downloads/pproftest.docx")
//	fmt.Println(s)
//	time.Sleep(10 *time.Second)
//}

type Info struct {
	IP string `json:"ip"`
}

func main() {
	str := "{\"ip\":\"8.8.8.8\",\"location\":{\"country\":\"US\",\"region\":\"California\",\"city\":\"Mountain View\",\"lat\":37.4223,\"lng\":-122.085,\"postalCode\":\"94043\",\"timezone\":\"-07:00\",\"geonameId\":5375480},\"domains\":[\"0--9.ru\",\"000.lyxhwy.xyz\",\"000180.top\",\"00049ok.com\",\"001998.com.he2.aqb.so\"],\"as\":{\"asn\":15169,\"name\":\"Google LLC\",\"route\":\"8.8.8.0\\/24\",\"domain\":\"https:\\/\\/about.google\\/intl\\/en\\/\",\"type\":\"Content\"},\"isp\":\"Google LLC\"}\n"
	var info Info
	if err := json.Unmarshal([]byte(str), &info); err !=nil {
		fmt.Println(err)
	}
	fmt.Println(info)
}
