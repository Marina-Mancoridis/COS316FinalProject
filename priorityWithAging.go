package main

import (
	"fmt"
	"math"
	"sort"
)

// currently, priorities are decremented every time a process has been completed, not by number of seconds
// that every process has been waiting... we should figure out how we are decrementing priorities!
// can play around with this metric as well...
func PriorityWithAging(processes []Process, totalTime int) {
	fmt.Println("\n\n                         Running Priority With Aging Scheduling Algorithm...")

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
		// update the queue and priorities
		for j := 0; j < len(processes); j++ {
			if processes[j].arrivalTime > currentTime {
				break
			}
			if !processes[j].completed {
				// process was already in queue -> decrement the priority
				if processes[j].isInQueue {
					// cannot have negative priorities
					if processes[j].priority >= 1 {
						processes[j].priority -= 1
					}
					// process was not already in queue -> add to queue
				} else {
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
	fmt.Println("RIGHT BEFORE ENTIRE FUNCTION RETURNS")
	printProcesses(processes)
	fmt.Println("--------------")

	// outputs statistics
	GenerateStatistics(currentTime, processes)
}
