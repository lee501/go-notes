/*反射就是用来检测存储在接口变量内部(值value；类型concrete type) pair对的一种机制*/

/*1.reflect基本功能 TypeOf()和 ValueOf(), 从接口中获取目标对象信息
  a. ValueOf()用来获取输入参数接口中的数据的值，如果接口为空则返回0
  b. TypeOf()用来动态获取输入参数接口中的值的类型，如果接口为空则返回nil*/

  // 示例代码
  package main

import(
    "fmt"
    "reflect"
  )

  func main(){
    var num float64 = 1.2345

    fmt.Printf("Value is %v, Type is %t", reflect.ValueOf(num), reflect.TypeOf(num))
  }
  /*运行接口
    type:  float64
    value:  1.2345
  */

2. 反射可以将"反射类型对象" 转换为"接口类型变量"
  a. reflect.value.Interface().(已知的类型)
  b. 遍历reflect.Type的Field获取其Field
  /* 1，执行 relect.ValueOf(interface) 之后， 得到一个类型为reflect.Value变量
     2，通过自身的interface()方法获得接口变量的真实内容
     3， 通过类型判断进行转换， 转换为原有真实类型
   */ 
  1).已知原有类型
     realValue := value.Interface().(已知类型)

    // 示例代码
    package main 

    import(
       "fmt"
       "reflect"
    )

    func main() {
      var num float64 = 2.3435

      // 指针 *float64
      pointer := reflect.ValueOf(&num)
      // float64
      value := reflect.ValueOf(num)

      /* 强制转换
        区分指针和值
        转换类型要完全符合，否者直接panic
      */
      convertPointer := pointer.Interface().(*float64)  //#=> 一个地址
      convertValue := value.Interface().(float64)  //#=>2。3435
    }

  2).未知原有类型，需遍历其filed, 例如结构体 struct
    // 获取字段
    // 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
    // 2. 再通过reflect.Type的Field获取其Field
    // 3. 最后通过Field的Interface()得到对应的value
    // 获取方法
    // 先获取interface的reflect.Type, 通过.NumMethod()进行遍历
    // 示例
    package main

    import (
      "fmt"
      "reflect"
    )

    type User struct {
      Id   int
      Name string
      Age  int
    }

    func (u User) RelectCallFunc() {
      fmt.Println("Allen.Wu ReflectCallFunc")
    }

    func main(){
      user := User{1, "Allen.Wu", 25}

      DoFieldAndMethod(user)
    }
    // 参数对象为intereface{}  
    func DoFieldAndMethod(input interface{}){
      // 获取类型
      getType := reflect.TypeOf(input)
      fmt.Println("get Type is :", getType.Name())
      // 获取值
      getValue := reflect.ValueOf(input)
      fmt.Println("get all Fields is:", getValue)

      // 获取字段
      // 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
      // 2. 再通过reflect.Type的Field获取其Field
      // 3. 最后通过Field的Interface()得到对应的value
      for i := 0; i < getType.NumField(); i++ {
        // 通过reflect.Type的Field获取其Field
        field := getType.Field(i)
        // 通过Field的Interface()得到对应的value
        value := getType.Field(i).Interface()
        fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
      }
      // 获取方法
      // 先获取interface的reflect.Type, 通过.NumMethod()进行遍历
      for i := 0; i < getType.NumMethod(); i++ {
        m := getType.Method(i)
        fmt.Printf("%s: %v\n", m.Name, m.Type)
      }
    }
    /*运行结果
      get Type is : User
      get all Fields is: {1 Allen.Wu 25}
      Id: int = 1
      Name: string = Allen.Wu
      Age: int = 25
      ReflectCallFunc: func(main.User)
    */

3. 通过reflect.Value设置实际变量的值
  // reflect.Value是通过reflect.ValueOf(X)获得的，只有当X是指针的时候，才可以通过reflec.Value修改实际变量X的值
  pointer := reflect.ValueOf(&num)
  newValue := pointer.Elem()

  newValue.SetFloat(77)

4. 做框架工程的时候，需要可以随意扩展方法, 通过reflect搞定
  
  // 示例代码
  user := User{1, "Allen.Wu", 25}
  
  // 1. 要通过反射来调用起对应的方法，必须要先通过reflect.ValueOf(interface)来获取到reflect.Value，得到“反射类型对象”后才能做下一步处理
  getValue := reflect.ValueOf(user)

  // 一定要指定参数为正确的方法名
  // 2. 先看看带有参数的调用方法
  methodValue := getValue.MethodByName("ReflectCallFuncHasArgs")
  args := []reflect.Value{reflect.ValueOf("wudebao"), reflect.ValueOf(30)}
  methodValue.Call(args)

  // 一定要指定参数为正确的方法名
  // 3. 再看看无参数的调用方法
  methodValue = getValue.MethodByName("ReflectCallFuncNoArgs")
  args = make([]reflect.Value, 0)
  methodValue.Call(args)















