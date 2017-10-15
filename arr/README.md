## arr

### 数组特性
- 数组的长度定义后就不可更改； 


### 语法 
```
var arr [n]type

arr  表示数组的名字
n    表示数组的长度
type 表示数组存储元素的类型
```

### 声明
```
var arr [3]byte

var arr1 [3]int

var arr2 [3]string

var arr3 [3]*int  // 长度为3的指针数组

var arr4 [3][2]int //二维数组
var arr5 [3][2][2]int //三维数组

arr6 := [...]int{1, 2, 3, 4, 5} // 编译器自动判断数组的长度
```


### 二维数组初始化 和 循环遍历
```
arr1 := [3][2]string{{"age", "address"}, {"name", "sex"}, {"weight", "high"}}
arr2 := [3][5]string{{"age", "address"}, {"name", "sex"}, {"weight", "high"}}
arr3 := [...][5]string{{"age", "address"}, {"name", "sex"}, {"weight", "high"}}

for _, v := range arr1 {
    for k, v := range v {
        fmt.Println(k, v)
    }
}
```
