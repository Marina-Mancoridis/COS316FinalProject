package main

import (
	"fmt"
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
}

// runs the list of processes for a maximum of totalTime seconds in
// accordance to the round robin scheduling algorithm
// func RoundRobin(processes []Process, totalTime int, timeQuantum int) {
// 	fmt.Println("\n\n                         Running Round Robin Scheduling Algorithm...")

// 	// sorts the list of processes by arrival time
// 	sort.Slice(processes, func(i, j int) bool {
// 		return processes[i].arrivalTime < processes[j].arrivalTime
// 	})

// 	// puts processes into a readyQueue
// 	readyQueue := list.New()
// 	for i := 0; i < len(processes); i++ {
// 		readyQueue.PushBack(processes[i])
// 	}

// 	fmt.Println("processes...")
// 	printProcesses(processes)

// 	i := 0
// 	currentTime := 0

// 	for currentTime < totalTime {
// 		// remove process from front of queue
// 		front := readyQueue.Front()
// 		process := *front
// 		readyQueue.Remove(process)

// 		// if the process can be executed on time
// 		if currentTime+process.duration <= totalTime {
// 			// if process can be executed within time quantum
// 			if (process.duration - process.secondsCompleted) <= timeQuantum {
// 				process.waitingTime += currentTime
// 				currentTime += process.duration - process.secondsCompleted
// 				process.completed = true
// 				process.turnaroundTime += currentTime
// 			} else {
// 				// if process cannot be executed within time quantum
// 				process.waitingTime += currentTime
// 				currentTime += timeQuantum
// 				process.secondsCompleted += timeQuantum
// 				readyQueue.PushBack(process)
// 			}
// 		} else {
// 			break
// 		}

// 		if i >= len(processes)-1 {
// 			break
// 		}
// 		i++
// 	}
// 	fmt.Println("\n")

// 	// outputs statistics
// 	GenerateStatistics(currentTime, processes)
// }