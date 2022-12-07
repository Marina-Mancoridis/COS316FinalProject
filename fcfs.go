package main

import (
	"fmt"
	"time"
)

func FirstComeFirstServe(processes []Process, totalTime time.Duration) {
	var processesCompleted float64 = 0
	i := 0

	start := time.Now()

	for time.Since(start) < (totalTime) {
		fmt.Println("starting process ", i)
		time.Sleep(time.Duration(processes[i].duration) * time.Second)
		processesCompleted = processesCompleted + 1

		if i >= len(processes) {
			break
		}

		i++
	}

	elapsed := time.Since(start)
	fmt.Println("elapsed time: ", elapsed)

	// calculate latency (for what's done so far) and throughput
	fmt.Println("processes completed: ", processesCompleted)

	// throughput
	// converts processes/nanosecond to processes/second
	throughput := float64(processesCompleted) / float64(totalTime) * 1000000000
	fmt.Println("throughput: ", throughput)

	// latency

}
