package main

import (
	"fmt"
	"sort"
)

// runs the list of processes for a maximum of totalTime seconds in
// accordance to the first come first serve scheduling algorithm
func FirstComeFirstServe(processes []Process, totalTime int) {
	fmt.Println("\n\n                         Running First Come First Serve Scheduling Algorithm...")

	// sorts the list of processes by arrival time
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].arrivalTime < processes[j].arrivalTime
	})

	fmt.Println("PROCESSES SORTED BY ARRIVAL TIME")
	printProcesses(processes)
	fmt.Println("--------------")

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
		for j := 0; j < len(processes); j++ {
			// finds first in queue (queue is in time order...)
			if processes[j].isInQueue {
				processId = j
				break
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
