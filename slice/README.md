# arr

## slice初始化
1. 对arr/slice切片
2. 切片的字面量
3. make创建slice

## make切片的长度和容量
- slice
- map
- channel

> cap

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
- 对slice扩容, 动态扩容;
- 一定要用原来的slice接收;


## 反转
```golang
package main

import (
	"fmt"
	"log"
)

// 反转字符串
func reverseString(s string) (string, error) {
	var ss string
	for i := len(s); i > 0; i-- {
		ss = fmt.Sprintf("%v%v", ss, string(s[i-1]))
	}
	return ss, nil
}

// 反转切片
func reverseSlice(s []int) ([]int, error) {
	var ss []int
	for i := len(s); i > 0; i-- {
		ss = append(ss, s[i-1])
	}
	return ss, nil

}

// 反转切片的2部分
func reverseSSlice(s []int, n int) ([]int, error) {

	s1 := s[n:]
	s2, _ := reverseSlice(s[:n])
	fmt.Println(s1)
	fmt.Println(s2)
	s2 = append(s2, s1...)
	return s2, nil
}

func main() {
	sli := []int{1, 2, 3, 8, 10}
	s1, err := reverseSlice(sli)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s1)

	str := "1x3u4aa566"
	s2, err := reverseString(str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s2)

	var n int = 2
	s := []int{1, 3, 5, 7, 9}
	s3, _ := reverseSSlice(s, n)
	fmt.Println(s3)

}
```


## 计算器
```golang
 package main
 
 import (
     "fmt"
     "strconv"
 )
 
 func operationalHandler(x int, oper string, y int) (sum int) {
     switch oper {
     case "/":
         sum = x / y
     case "+":
         sum = x + y
     case "-":
         sum = x - y
     case "*":
         sum = x * y
     }
     return
 
 }
 
 func calc(args ...string) (sum int) {
     sum, _ = strconv.Atoi(args[0])
 
     for i := 1; i < len(args); i++ {
         if i+1 < len(args) {
             if ii, err := strconv.Atoi(args[i+1]); err == nil {
                 sum = operationalHandler(sum, args[i], ii)
             }
         }
     }
     return
 }
 
 func main() {
     fmt.Println(calc("4", "+", "2", "/", "3", "+", "1", "*", "10"))
 }
```


