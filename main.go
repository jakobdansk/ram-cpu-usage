package main

import (
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	for {
		v, err := mem.VirtualMemory()
		if err != nil {
			log.Fatalf("Memory error: %v", err)
		}

		c, err := cpu.Percent(0, false)
		if err != nil {
			log.Fatalf("CPU error: %v", err)
		}

		fmt.Printf("RAM usage: %.2f%%, CPU usage: %.2f%%\n", v.UsedPercent, c[0])
		time.Sleep(2 * time.Second)
	}
}
