# Tail

```golang
package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	config := tail.Config{Follow: true}
	tail, _ := tail.TailFile("/tmp/access.log", config)
	for line := range tail.Lines {
		fmt.Println(line.Text, line.Time)
		time.Sleep(time.Second)
	}
}
```
