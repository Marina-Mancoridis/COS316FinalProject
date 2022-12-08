package main

import (
	"fmt"
	"time"
)

// generates throughput, average waiting time, and average turnaround
// statistics after a single call of a scheduling algorithm
func GenerateStatistics(totalTime time.Duration, processes []Process) {
	fmt.Println("STATISTICS:")

	// grabbing information from pass of processes
	var processesCompleted float64 = 0
	var totalWaitingTime = time.Duration(0)
	var totalTurnaroundTime = time.Duration(0)
	for i := 0; i < len(processes); i++ {
		if processes[i].completed == true {
			totalWaitingTime += processes[i].waitingTime
			totalTurnaroundTime += processes[i].turnaroundTime
			processesCompleted++
		}
	}
	
	// throughput
	// converts processes/nanosecond to processes/second
	throughput := float64(processesCompleted) / float64(totalTime) * 1000000000
	fmt.Println("throughput (processes/second):             ", throughput)

	// average waiting time per process, in seconds/process
	var averageWaitingTime = float64(totalWaitingTime) / (processesCompleted * 1000000000)
	fmt.Println("average waiting time (seconds/process):    ", averageWaitingTime)

	// average turnaround time per process, in seconds/process
	var averageTurnaroundTime = float64(totalTurnaroundTime) / (processesCompleted * 1000000000)
	fmt.Println("average turnaround time (seconds/process): ", averageTurnaroundTime)
}



// runs the list of processes for a maximum of totalTime seconds in 
// accordance to the first come first serve scheduling algorithm
func FirstComeFirstServe(processes []Process, totalTime time.Duration) {
	fmt.Println("Running First Come First Serve Algorithm...")
	i := 0
	start := time.Now()

	for time.Since(start) < (totalTime) {
		fmt.Println("starting process ", i)

		processes[i].waitingTime += time.Since(start)
		time.Sleep(time.Duration(processes[i].duration) * time.Second)
		processes[i].completed = true
		processes[i].turnaroundTime += time.Since(start)

		if i >= len(processes) {
			break
		}
		i++
	}
	fmt.Println("elapsed time: ", time.Since(start))
	fmt.Println("\n\n\n")

	// outputs statistics for first come first serve algorithm
	GenerateStatistics(totalTime, processes)
}
