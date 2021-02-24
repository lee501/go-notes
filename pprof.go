package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

//import (
//	"context"
//	"fmt"
//)
//
////func main() {
////	var isCPUPprof bool
////	flag.BoolVar(&isCPUPprof, "cpu", false, "cpu pprof")
////	flag.Parse()
////	if isCPUPprof {
////		file, err := os.Create("./cpu.pprof")
////		if err != nil {
////			fmt.Printf("create cpu pprof failed, err:%v\n", err)
////			return
////		}
////		pprof.StartCPUProfile(file)
////		defer pprof.StopCPUProfile()
////	}
////
////	s := file_process.ReadFileByString("/Users/richctrl/Downloads/pproftest.docx")
////	fmt.Println(s)
////	time.Sleep(10 *time.Second)
////}
//
//type Info struct {
//	IP string `json:"ip"`
//}
//
////func main() {
////	str := "{\"ip\":\"8.8.8.8\",\"location\":{\"country\":\"US\",\"region\":\"California\",\"city\":\"Mountain View\",\"lat\":37.4223,\"lng\":-122.085,\"postalCode\":\"94043\",\"timezone\":\"-07:00\",\"geonameId\":5375480},\"domains\":[\"0--9.ru\",\"000.lyxhwy.xyz\",\"000180.top\",\"00049ok.com\",\"001998.com.he2.aqb.so\"],\"as\":{\"asn\":15169,\"name\":\"Google LLC\",\"route\":\"8.8.8.0\\/24\",\"domain\":\"https:\\/\\/about.google\\/intl\\/en\\/\",\"type\":\"Content\"},\"isp\":\"Google LLC\"}\n"
////	var info Info
////	if err := json.Unmarshal([]byte(str), &info); err !=nil {
////		fmt.Println(err)
////	}
////	fmt.Println(info)
////}
//
//func f(ctx context.Context) {
//	context.WithValue(ctx, "foo", 6)
//}
//
//func main() {
//	done:=make(chan int ,100)
//	defer close(done)
//	//开启线程
//	for i := 1; i <= cap(done); i++ {
//		go func(i int) {
//			fmt.Println("开启线程", i)
//			done <- i
//		}(i)
//	}
//	//使用channel阻塞的方式来出来同步
//	//此处不能使用range 会引起主线程deadline
//	//for i := 0; i< 100; i++ {
//	//	m := <-done
//	//	fmt.Println(m, "线程关闭")
//	//}
//	for v := range done {
//		fmt.Println(v, "线程关闭")
//	}
//	fmt.Println("执行完毕")
//}
func GetFbSeries(n int) []int {
	res := make([]int, 2, n)
	res[0] = 1
	res[1] = 1
	for i := 2; i < n; i++ {
		res = append(res, res[i-1]+res[i-2])
	}
	return res
}

func createFBS(w http.ResponseWriter, r *http.Request) {
	var fbs []int
	for i := 0; i < 1000000; i++ {
		fbs = GetFbSeries(50)
	}
	w.Write([]byte(fmt.Sprintf("%v", fbs)))
}

func main() {
	http.HandleFunc("/fb", createFBS)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
