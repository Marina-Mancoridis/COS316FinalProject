package main

import (
	"fmt"
	"math"
	"sort"
)

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

	// outputs statistics
	GenerateStatistics(currentTime, processes)
}

// priority algorithm with queue and without aging
func PriorityWithQueue(processes []Process, totalTime int) {
	fmt.Println("\n\n                         Running Priority (With Queue) Scheduling Algorithm...")

	// sorts the list of processes by arrival time
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].arrivalTime < processes[j].arrivalTime
	})

	// fmt.Println("PROCESSES SORTED BY ARRIVAL TIME")
	// printProcesses(processes)
	// fmt.Println("--------------")

	i := 0
	currentTime := 0
	numProcessesComplete := 0

	for currentTime < totalTime {
		fmt.Println("---------------------------------------------------")
		fmt.Println("AT TIME STEP ", currentTime)
		// add to queue
		for j := 0; j < len(processes); j++ {
			if processes[j].arrivalTime > currentTime {
				break
			}
			// this is the only part that I changed ... -marm
			if !processes[j].completed {
				// process was not in queue -> add it
				if !processes[j].isInQueue {
					processes[j].isInQueue = true
				}
				// no need to update any processes who haven't arrived yet
			}
		}

		fmt.Println("queue and priorities supposedly updated...")
		printProcesses(processes)
		fmt.Println("--------------")

		// find the next process to execute
		processId := -1
		lowestPriority := math.MaxInt
		for j := 0; j < len(processes); j++ {
			if processes[j].isInQueue {
				// note: between processes of same priority, processes that arrived first are prioritized
				if processes[j].priority < lowestPriority {
					processId = j
					lowestPriority = processes[j].priority
				}
			}
		}
		fmt.Println("next process to execute has id: ", processId)

		// return if there are no more processes to execute
		if processId == -1 {
			if numProcessesComplete < len(processes) {
				currentTime++
				continue
			}
			return
		}

		// execute the next process, mark as completed, take it off the queue
		if currentTime+processes[processId].duration <= totalTime {
			currentTime += processes[processId].duration
			processes[processId].completed = true
			numProcessesComplete++
			processes[processId].isInQueue = false
			processes[processId].turnaroundTime += processes[processId].duration
		} else {
			break
		}

		// update the waiting and turnaround time of all processes in the queue
		for j := 0; j < len(processes); j++ {
			if processes[j].isInQueue {
				processes[j].waitingTime += processes[processId].duration
				processes[j].turnaroundTime += processes[processId].duration
			}
		}

		if i >= len(processes)-1 {
			break
		}
		i++
	}
	fmt.Println()

	// outputs statistics
	GenerateStatistics(currentTime, processes)
}
