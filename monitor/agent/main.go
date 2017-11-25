package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/467754239/GolangNote/monitor/common"
	"github.com/467754239/logrus"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

var (
	transAddr = flag.String("trans", "127.0.0.1:6000", "transfer address")
	log       = logrus.New()
)

func init() {
	log.Formatter = new(logrus.JSONFormatter)
	log.Formatter = new(logrus.TextFormatter) // default
	log.Level = logrus.DebugLevel
}

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
		log.Panic(err)
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

func MemMetric() []*common.Metric {
	var ret []*common.Metric

	mem_stat, err := mem.VirtualMemory()
	if err != nil {
		log.Panic(err)
		return ret
	}

	mem_total := mem_stat.Total / 1024 / 1024
	fmt.Println(reflect.TypeOf(mem_total))
	mem_total_str := float64(mem_total)
	fmt.Println(reflect.TypeOf(mem_total_str))

	metric := NewMetric("mem.total", mem_total_str)
	ret = append(ret, metric)

	metric = NewMetric("mem.used_percent", mem_stat.UsedPercent)
	ret = append(ret, metric)

	return ret

}

func NewUserMetric(cmdstr string) MetricFunc {
	return func() []*common.Metric {
		metrics, err := getUserMetrics(cmdstr)
		if err != nil {
			log.Print(err)
			return []*common.Metric{}
		}
		return metrics
	}
}

func getUserMetrics(cmdstr string) ([]*common.Metric, error) {
	// 构建命令
	// 获取标准输出
	// 按行解析
	// 获取key， value
	// 包装成common.Metric

	var ret []*common.Metric
	cmd := exec.Command("bash", "-c", cmdstr)
	stdout, _ := cmd.StdoutPipe()

	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(stdout)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}
		key, value := fields[0], fields[1]
		n, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Print(err)
			continue
		}
		metric := NewMetric(key, n)
		ret = append(ret, metric)
	}

	return ret, nil
}

func main() {
	flag.Parse()

	log.Info("Start Agent......")
	// 初始化构造函数
	sender := NewSender(*transAddr)

	// 返回构造函数的ch
	ch := sender.Channel()

	sched := NewSched(ch)

	sched.AddMetric(NewUserMetric("./usr.py"), time.Second*5)
	sched.AddMetric(CpuMetric, time.Second*5)
	sched.AddMetric(MemMetric, time.Second*1)

	// 主协成阻塞
	sender.Start()
}
