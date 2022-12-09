package main

import (
	"fmt"
	"math"
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
}

// runs the list of processes for a maximum of totalTime seconds in
// accordance to the first come first serve scheduling algorithm
func FirstComeFirstServe(processes []Process, totalTime int) {
	fmt.Println("\n\n                         Running First Come First Serve Scheduling Algorithm...")

	// sorts the list of processes by arrival time
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].arrivalTime < processes[j].arrivalTime
	})

	// fmt.Println("processes...")
	// printProcesses(processes)

	i := 0
	currentTime := 0

	for currentTime < (totalTime) {
		// if the process can be executed on time
		if currentTime+processes[i].duration <= totalTime {
			processes[i].waitingTime += currentTime
			currentTime += processes[i].duration
			processes[i].completed = true
			processes[i].turnaroundTime += currentTime
		} else {
			break
		}

		if i >= len(processes)-1 {
			break
		}
		i++
	}
	// fmt.Println("\n")

	// outputs statistics for first come first serve scheduling algorithm
	GenerateStatistics(currentTime, processes)
}

// runs the list of processes for a maximum of totalTime seconds in
// accordance to the shortest job first scheduling algorithm
func ShortestJobFirst(processes []Process, totalTime int) {
	fmt.Println("\n\n                         Running Shortest Job First Scheduling Algorithm...")

	// sorts the list of processes by duration
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].duration < processes[j].duration
	})

	// fmt.Println("processes...")
	// printProcesses(processes)

	i := 0
	currentTime := 0

	for currentTime < totalTime {
		// if the process can be executed on time
		if currentTime+processes[i].duration <= totalTime {
			processes[i].waitingTime += currentTime
			currentTime += processes[i].duration
			processes[i].completed = true
			processes[i].turnaroundTime += currentTime
		} else {
			break
		}

		if i >= len(processes)-1 {
			break
		}
		i++
	}
	// fmt.Println("\n")

	// outputs statistics for shortest job first scheduling algorithm
	GenerateStatistics(currentTime, processes)
}

// // runs the list of processes for a maximum of totalTime seconds in
// // accordance to the priority scheduling algorithm
// // uses low numbers as high priority, with 1 as highest priority
func Priority(processes []Process, totalTime int) {
	fmt.Println("\n\n                         Running Priority Scheduling Algorithm...")

	// sorts the list of processes by priority
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].priority < processes[j].priority
	})

	// fmt.Println("processes...")
	// printProcesses(processes)

	i := 0
	currentTime := 0

	for currentTime < totalTime {
		// if the process can be executed on time
		if currentTime+processes[i].duration <= totalTime {
			processes[i].waitingTime += currentTime // i feel like this is wrong!
			currentTime += processes[i].duration
			processes[i].completed = true
			processes[i].turnaroundTime += currentTime
		} else {
			break
		}

		if i >= len(processes)-1 {
			break
		}
		i++
	}
	// fmt.Println("\n")

	// outputs statistics for shortest job first scheduling algorithm
	GenerateStatistics(currentTime, processes)
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

// 	// outputs statistics for shortest job first scheduling algorithm
// 	GenerateStatistics(currentTime, processes)
// }

func PriorityWithAging(processes []Process, totalTime int) {
	fmt.Println("\n\n                         Running Priority With Aging Scheduling Algorithm...")

	// sorts the list of processes by arrival time
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].arrivalTime < processes[j].arrivalTime
	})

	// fmt.Println("processes...")
	// printProcesses(processes)

	i := 0
	currentTime := 0

	for currentTime < totalTime {
		// update the queue and priorities
		for j := 0; j < len(processes); j++ {
			if processes[j].arrivalTime <= currentTime {
				// process was already in queue -> decrement the priority
				if processes[j].isInQueue {
					if processes[j].priority >= 2 {
						processes[j].priority -= 1
					}
					// process was not already in queue -> add to queue
				} else {
					processes[j].isInQueue = true
				}
				// no need to update any processes who haven't arrived yet
			} else {
				break
			}
		}

		// find the next process to execute
		processId := 0
		lowestPriority := math.MaxInt
		for j := 0; j < len(processes); j++ {
			if processes[j].isInQueue {
				// note: between processes of same priority, processes that arrived first are prioritized
				if processes[j].priority < lowestPriority {
					processId = j
				}
			}
		}

		// execute the next process, mark as completed, take it off the queue
		if currentTime+processes[processId].duration <= totalTime {
			currentTime += processes[processId].duration
			processes[processId].completed = true
			processes[processId].isInQueue = false
			processes[processId].turnaroundTime += processes[processId].duration
		} else {
			break
		}

		// update the waiting time of all processes in the queue
		for j := 0; j < len(processes); j++ {
			if processes[j].isInQueue {
				processes[j].waitingTime += processes[processId].duration
			}
		}

		if i >= len(processes)-1 {
			break
		}
		i++
	}
	// fmt.Println("\n")

	// outputs statistics for shortest job first scheduling algorithm
	// GenerateStatistics(currentTime, processes)
}
