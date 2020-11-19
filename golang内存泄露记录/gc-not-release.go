package gcunreleased

//1. 字符串与截取的字串共享，暂时性泄露
var s0 string //包级变量

func f(s string) {
	s0 = s[:50]
	//s0和s共享一个内存块, 到这里只有50个字节需要使用，s其他部分无法释放，需要s0在其他地方被重新修改为止
}

func demo() {
	s := "abcedfghijkcclskjsuhdksls"
	f(s)
}

/*
	解决办法，转化成字节切片处理，然后再转化回来
	s0 = string([]byte(s[:50]))
*/

//2. 子切片造成的暂时性内存泄露
var ss []int

func g(s1 []int) {
	// 假设s1的长度远大于30。ss与s1共享切片，所以在ss被重新赋值之前，s1不会被释放
	ss = s1[len(s1)-30:]
}
/*
	通过append方法处理
	ss = append(s1[:0:0], s1[len(s1)-30:]...)
*/

//3. goroutine被永久阻塞会造成永久性内存泄露