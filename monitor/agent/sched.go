package main

import (
	"log"
	"time"

	"github.com/467754239/GolangNote/monitor/common"
)

// 一次性返回多个值
type MetricFunc func() []*common.Metric

type Sched struct {
	ch chan *common.Metric
}

func NewSched(ch chan *common.Metric) *Sched {
	return &Sched{
		ch: ch,
	}
}

// 每隔多长时间收集一次
func (s *Sched) AddMetric(collecter MetricFunc, step time.Duration) {
	go func() {
		ticker := time.NewTicker(step)
		defer ticker.Stop()
		for range ticker.C {

			metrics := collecter()
			for _, metric := range metrics {
				log.Printf("metric:%v\n", metric)
				s.ch <- metric
			}
		}
	}()

}
