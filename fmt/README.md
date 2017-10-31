# fmt
- [x] %v 打印变量
- [x] %s 打印字符串
- [x] %d 打印数字
- [x] %f 打印浮点型
- [x] %p 打印指针


## 格式化输出

```golang
fmt.Println("hello world!!!")
fmt.Printf("%v", "hello world!!!")

var f *os.File
fmt.Fprintf(f, "%v", "hello world")

s := fmt.Sprintf("%v", 1)
```

## 格式化输入

```golang
var n int
// 要想修改一个变量的值，必须传递这个变量的指针；
fmt.Scanf("%d", &n)
fmt.Println(n)
```
> example 1

```golang
package main

import "fmt"

func main() {
    var s string
    var n int
    fmt.Scanf("%s %d", &s, &n)
    fmt.Println(s, n)

    var s1 string
    var n1 int
    fmt.Scanln(&s1, &n1)
    fmt.Println(s1, n1)
}
```
