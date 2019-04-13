package main
/*Golang关于类型设计的一些原则
  1·变量包括 (type, value) 两部分
    。所以为什么nil != nil了

  2·type包括static type 和 concrete type, 前者是编码过程中的类型(如int、string), 后者是runtime系统的类型

  3·类型断言是否成功，取决于concrete type, 因此一个reader变量如果它的concrete type 也实现了write方法, 就可以断言为writer

@反射主要与Golang的 interface 类型相关（它的 type 是concrete type），只有interface类型才有反射一说。

@每个 interface 变量都有一个对应pair， interface{}类型的变量包含两个指针，一个指向值的的类型 concrete type, 另一个指向实际的值 value
  (value, type)


  // 创建类型为*os.File的变量，将其赋值给接口变量
    testfile, err := os.OpenFile("dev/test", os.O_RDWR, 0)
  // 将其赋值给接口变量
    var r = io.Reader
    r = testfile
    此时接口变量r的pair为(testfile, *os.File), 这个pair在接口变量连续赋值的过程中是不变的
  // 将r赋值给另一接口变量w
    var w = io.Writer
    // 通过接口类型断言来赋值
    w = r.(io.Writer) //接口变量w的pair与r的pair相同，都是:(tty, *os.File

//总结： 反射就是用来检测存储在接口变量内部(值value；类型concrete type) pair对的一种机制。
*/
