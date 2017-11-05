# len

```golang
package main

import "fmt"

func main() {
	s := "golang你好"
	fmt.Println("len s: ", len(s))

	cnt := 0
	for _, v := range s {
		fmt.Println(v)
		cnt += 1
	}
	fmt.Printf("cnt1: %v\n", cnt)

	cnt = 0
	for _, v := range []byte(s) {
		fmt.Println(v)
		cnt += 1
	}
	fmt.Printf("cnt2: %v\n", cnt)
}
```

```bash
1. len 统计的是字节数
2. for str 循环 循环的是合法的utf-8字符
3. for byte 循环的是字节数
```
