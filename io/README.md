## 实现简化cat命令

```
io.Copy(dst, src)
把右边类文件对象的数据拷贝到左边，直到读取到EOF.
以流式的方式拷贝.

func Copy(dst Writer, src Reader) (written int64, err error)
https://godoc.org/io#Copy
```


