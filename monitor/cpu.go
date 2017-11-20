package main

import (
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	fmt.Println("---------cpu load-----------------")
	//cpu_percent, err := cpu.Percent(time.Second, true)
	cpu_percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("cpu percent: %v\n", cpu_percent[0])

	avgstat, err := load.Avg()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("load1 : %v\n", avgstat.Load1)
	fmt.Printf("load5 : %v\n", avgstat.Load5)
	fmt.Printf("load15: %v\n", avgstat.Load15)

	fmt.Println("---------mem-----------------")
	mem_stat, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("mem UsedPercent: %v\n", mem_stat.UsedPercent)
	fmt.Printf("mem total: %v\n", mem_stat.Total/1024/1024)

	fmt.Println("---------disk-----------------")

	disk_stat, err := disk.Usage("/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(disk_stat.Total / 1024 / 1024 / 1024)
	fmt.Println(disk_stat.UsedPercent)
}
