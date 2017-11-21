package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/467754239/GolangNote/monitor/common"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
)

var (
	transAddr = flag.String("trans", "127.0.0.1:6000", "transfer address")
)

func NewMetric(metric string, value float64) *common.Metric {
	hostname, _ := os.Hostname()
	return &common.Metric{
		Metric:    metric,
		Endpoint:  hostname,
		Value:     value,
		Tag:       []string{runtime.GOOS},
		Timestamp: time.Now().Unix(),
	}

}

func CpuMetric() []*common.Metric {
	var ret []*common.Metric

	cpus, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Fatal(err)
	}

	metric := NewMetric("cpu.percent", cpus[0])
	ret = append(ret, metric)

	cpuload, err := load.Avg()
	if err == nil {

		metric = NewMetric("cpu.load1", cpuload.Load1)
		ret = append(ret, metric)

		metric = NewMetric("cpu.load5", cpuload.Load5)
		ret = append(ret, metric)
	}

	return ret
}

func main() {
	flag.Parse()

	log.Println(*transAddr)
	// 初始化构造函数
	sender := NewSender(*transAddr)

	// 返回构造函数的ch
	ch := sender.Channel()

	sched := NewSched(ch)
	sched.AddMetric(CpuMetric, time.Second*5)

	fmt.Println("sched finish.")
	sender.Start()
}