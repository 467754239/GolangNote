# arr

## slice初始化
1. 对arr/slice切片
2. 切片的字面量
3. make创建slice

## 切片的长度和容量


## nil切片
> nil切片不能赋值
```golang
package main

improt "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil")
	}
}

```

## append切片
- 对slice扩容;
- 一定要用原来的slice接收;


## 
