# fmt
- [x] %v 打印变量
- [x] %s 打印字符串
- [x] %d 打印数字
- [x] %f 打印浮点型
- [x] %p 打印指针

```golang
fmt.Println("hello world!!!")
fmt.Printf("%v", "hello world!!!")

var f *os.File
fmt.Fprintf(f, "%v", "hello world")

s := fmt.Sprintf("%v", 1)
```
