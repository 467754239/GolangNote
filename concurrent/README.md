## concurrent

```golang
package main

import (
	"fmt"
	"net/http"
	"sync"
)

type TaskRunner struct {
	token chan bool
	wg    sync.WaitGroup
}

func NewTaskRunner(concurrent int) *TaskRunner {
	return &TaskRunner{
		token: make(chan bool, concurrent),
	}
}

func (this *TaskRunner) Put(task func()) {
	this.token <- true
	this.wg.Add(1)
	go func() {
		defer this.wg.Done()
		task()
		<-this.token
	}()
}

func (this *TaskRunner) Wait() {
	this.wg.Wait()
}

func main() {
	t := NewTaskRunner(10)
	for i := 0; i < 1000; i++ {
		i := i
		t.Put(func() {
			fmt.Println(i)
			_, _ = http.Get("http://www.google.com")
		})
	}
	t.Wait()
}
```
