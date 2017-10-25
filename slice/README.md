## arr

slice初始化的两种方式

```
1. 对arr/slice切片
2. 赋值初始化 
3. make创建slice
```

```
1. 切片的长度和容量
```


nil切片

> nil切片不能复制
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

