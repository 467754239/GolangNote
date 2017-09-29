## 操作文件的多种方式

- [x] file.Read
- [x] ioutil.ReadFile 
- [x] bufio.Scanner 
- [x] bufio.Reader
- [x] io.Copy 

1. 按块的方式，原生读取，也称为裸读取，很少使用。
> cpu寄存器 > l1/l2/l3 cache > Mem > Disk
```
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fd, err := os.Open("/etc/passwd")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	/*
		方式1. 裸读取，很少使用。
	*/
	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := fd.Read(buf) // 每次Read都是从磁盘上。
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		//fmt.Println(string(buf[:n]))
		chunks = append(chunks, buf[:n]...)
	}
	fmt.Print(string(chunks))
}
```

2. 加上buffer的读取，很高效。
> 预读取，bufio就是一个高效读取的方式。对普通文件读取的封装。  
> 把数据预先放到内存buffer里，以后读数据从内存中读取，而不是从磁盘上。
```
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fd, err := os.Open("/etc/passwd")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	/*
		方式2. 加上buffer的读取，很高效。
	*/
	r := bufio.NewReader(fd) // 预先加载数据到内存buffer

	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf) // 每次Read是从内存buffer而不是从磁盘上。
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		//fmt.Println(string(buf[:n]))
		chunks = append(chunks, buf[:n]...)
	}
	fmt.Print(string(chunks))
}
```

3.1 小文件一次性全部读取
> Version 1

```
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	/*
		方式3.1 小文件一次性读取.
	*/
	buf, err := ioutil.ReadFile("/var/log/messages")
	//buf, err := ioutil.ReadFile("/etc/passwd")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(buf))

}
```

3.2 小文件一次性全部读取 

```
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fd, err := os.Open("/etc/passwd")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	/*
		方式3.2 小文件一次性读取.
	*/
	buf, err := ioutil.ReadAll(fd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(buf))

}
```

4. 按行读取，按分隔符读取

```
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fd, err := os.Open("/etc/passwd")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	/*
		方式4 按行读取，按分隔符读取.
	*/
	r := bufio.NewReader(fd)
	for {
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}

}
```

5. 操作类文件的神器

```
package main

import (
	"io"
	"os"
)

func main() {
	/*
		5. 操作类文件的神器
	*/
	io.Copy(os.Stdout, os.Stdin)
}
```
