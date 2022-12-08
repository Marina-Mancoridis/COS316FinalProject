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

		// in fcfs, you wait as long as it takes to start your process
		processes[i].waitingTime += time.Since(start)

		time.Sleep(time.Duration(processes[i].duration) * time.Second)
		
		processesCompleted = processesCompleted + 1
		processes[i].completed = true

		if i >= len(processes) {
			break
		}

		i++
	}

	elapsed := time.Since(start)
	fmt.Println("elapsed time: ", elapsed)

	fmt.Println("processes completed: ", processesCompleted)

	// throughput
	// converts processes/nanosecond to processes/second
	throughput := float64(processesCompleted) / float64(totalTime) * 1000000000
	fmt.Println("throughput: ", throughput)

	// total time waited for processed processes
	var totalWaitingTime = time.Duration(0)
	for i := 0; i < len(processes); i++ {
		if processes[i].completed == true {
			totalWaitingTime += processes[i].waitingTime
			fmt.Println("process ", i)
			fmt.Println("process waited: ", processes[i].waitingTime)
			fmt.Println()
		}
	}

	// average waiting time per process, in seconds/process
	var averageWaitingTime = float64(totalWaitingTime) / (processesCompleted * 1000000000)
	fmt.Println("average waiting time: ", averageWaitingTime)
}
