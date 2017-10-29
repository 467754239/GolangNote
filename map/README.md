## map 
- [x] hash
- [x] 无序的
- [x] O(1)的访问时间

> 时间复杂度

```bash
O(1)  
取值一次直接命中

O(n)
有多少个元素差不多就要花费多少时间来查找到它

hash map
无论有多少个元素，差不多都是一下子直接命中

O(logn)
tree map(C语言)
key是排序的
```

1. 创建map

```

```

2. 空map

```
1. 空map不能做任何操作;
2. 空map和ni是相等的;
3. 如果要对空map操作，必须先要使用make对其进行初始化;

var m map[string]int
fmt.Pritnln(m == nil)

m = make(map[string]int)
```

3. 词频统计

```

```

4. 遍历

```

```

5. 添加

```

```

6. 修改

```

```
7. 删除
```

```
