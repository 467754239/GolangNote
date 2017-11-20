package main

import (
	"time"

	"github.com/467754239/GolangNote/monitor/common"
)

// 一次性返回多个值
type MetricFunc func() []*common.Metric

type Sched struct {
	ch chan *common.Metric
}

func NewSched(ch chan *common.Metric) *Sched {
	return nil
}

// 每隔多长时间收集一次
func (s *Sched) AddMetric(collecter MetricFunc, step time.Duration) {

}
