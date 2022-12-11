package main

import (
	"fmt"
	"sort"
)

// generates throughput, average waiting time, and average turnaround
// statistics after a single call of a scheduling algorithm
func GenerateStatistics(elapsedTime int, processes []Process) {
	fmt.Println("STATISTICS:")

	// grabbing information from pass of processes
	var processesCompleted = 0
	var totalWaitingTime = 0
	var totalTurnaroundTime = 0
	for i := 0; i < len(processes); i++ {
		if processes[i].completed == true {
			totalWaitingTime += processes[i].waitingTime
			totalTurnaroundTime += processes[i].turnaroundTime
			processesCompleted++
		}
	}

	// total time
	fmt.Println("elapsed time:                        ", elapsedTime)

	// processes completed
	fmt.Println("processes completed:                 ", processesCompleted)

	// throughput
	// converts processes/nanosecond to processes/second
	throughput := float64(processesCompleted) / float64(elapsedTime)
	fmt.Println("throughput (processes/s):            ", throughput)

	// average waiting time per process, in seconds/process
	var averageWaitingTime = float64(totalWaitingTime) / float64(processesCompleted)
	fmt.Println("average waiting time (s/process):    ", averageWaitingTime)

	// average turnaround time per process, in seconds/process
	var averageTurnaroundTime = float64(totalTurnaroundTime) / float64(processesCompleted)
	fmt.Println("average turnaround time (s/process): ", averageTurnaroundTime)

	// average waiting time per process with each initial priority
	fmt.Println("average waiting time by priority:")
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].priority < processes[j].priority
	})
	lowestInitialPriority := processes[0].priority
	highestInitialPriority := processes[len(processes) - 1].priority

	// one average waiting time statistic per initial priority
	for i := lowestInitialPriority; i <= highestInitialPriority; i++ {
		waitingTimeSum := 0
		totalProcesses := 0
		// finds processes with initial priority of i
		for j := 0; j < len(processes); j++ {
			if processes[j].initialPriority == i {
				waitingTimeSum += processes[j].waitingTime
				totalProcesses += 1
			}
		}

		if totalProcesses == 0 {
			fmt.Printf("     priority %d: %s\n", i, "n/a")
		} else {
			averageWait := float64(waitingTimeSum) / float64(totalProcesses)
			fmt.Printf("     priority %d: %.3f\n", i, averageWait)
		}
	}
}