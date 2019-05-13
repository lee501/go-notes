# go_libraries
*woshi*
---
----
***
****
图片
![图片alt](图片地址 "图片title")

[简书](http://jianshu.com)

```
列表语法
-
+
*

1. wo
   - ai
   + ni
2. ni

```

*golang fmt.Printf 显示类型和值的小技巧*
```$xslt
//如果值是一个结构体，%+v 的格式化输出内容将包括结构体的字段名。
fmt.Printf("%+v\n", p) // {x:1 y:2}

//%#v 形式则输出这个值的 Go 语法表示。例如，值的运行源代码片段。
fmt.Printf("%#v\n", p) // main.point{x:1, y:2}
```