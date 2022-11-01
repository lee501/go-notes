package main

/*
	已关闭的chan只能读不能写
	声明语法
		var 名称 chan 类型
		var 名称 chan <- 类型 //只写
		var 名称 <- chan 类型//只读
		名称:=make(chan int) //无缓存channel
		名称:=make(chan int,0)//无缓存channel
		名称:=make(chan int,100)//有缓存channel
	取值
		ch <- 值 //向ch中添加一个值
		<- ch //从ch中取出一个值
		a:=<-ch //从ch中取出一个值并赋值给a
		a,b:=<-ch//从ch中取出一个值赋值给a,如果ch已经关闭或ch中没有值,b为false
*/

//主线程和子线程的同步； 子线程至今通信
/*
func main() {
	//声明chan用于子线程通信
	ch := make(chan string)
	//同步
	chi := make(chan int)

	go func() {
		fmt.Println("开始传递信息")
		ch <- "来自星星的你"
		close(ch)
		chi <- 1
	}()

	go func() {
		fmt.Println("开始接收数据")
		content := <- ch
		fmt.Println("接收到的数据为:", content)
		chi <- 2
	}()
	//查看当前同步的子线程
	result := <- chi
	fmt.Println(result)
	result = <- chi
	fmt.Println(result)
}
*/
