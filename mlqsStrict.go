package main

import (
	"fmt"
	"math"
	"sort"
)

// runs the list of processes for a maximum of totalTime seconds in
// accordance to the strict multi level queue scheduling algorithm
// low priority = [0, mediumPriorityCutoff)
// medium priority = [mediumPriorityCutoff, highPriorityCutoff)
// high priority = [highPriorityCutoff, inf)
func StrictMultiLevelQueueScheduling(processes []Process, totalTime int,
	mediumPriorityCutoff int, highPriorityCutoff int) {
	fmt.Println("\n\n                         Running Strict Multi Level Queue Scheduling Algorithm...")

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

		// add to queue for all processes
		for j := 0; j < len(processes); j++ {
			if processes[j].arrivalTime > currentTime {
				break
			}
			if !processes[j].completed {
				// process was not in queue -> add it to queue
				if !processes[j].isInQueue {
					processes[j].isInQueue = true
					processes[j].turnaroundTime = currentTime - processes[j].arrivalTime
					processes[j].waitingTime = currentTime - processes[j].arrivalTime
				}
				// no need to update any processes who haven't arrived yet
			}
		}

		fmt.Println("all queues supposedly updated...")
		printProcesses(processes)
		fmt.Println("--------------")

		// find the next process to execute, if exists in low queue
		processId := -1
		lowestDuration := math.MaxInt
		for j := 0; j < len(processes); j++ {
			if processes[j].priority < mediumPriorityCutoff {
				if processes[j].isInQueue {
					// note: between processes of same duration, processes that arrived first are prioritized
					if processes[j].duration < lowestDuration {
						processId = j
						lowestDuration = processes[j].duration
					}
				}
			}
		}

		// find the next process to execute, if exists in med queue
		if processId == -1 {
			lowestDuration := math.MaxInt
			for j := 0; j < len(processes); j++ {
				if processes[j].priority >= mediumPriorityCutoff && processes[j].priority < highPriorityCutoff {
					if processes[j].isInQueue {
						// note: between processes of same duration, processes that arrived first are prioritized
						if processes[j].duration < lowestDuration {
							processId = j
							lowestDuration = processes[j].duration
						}
					}
				}
			}
		}

		// find the next process to execute, if exists in high queue
		if processId == -1 {
			lowestDuration := math.MaxInt
			for j := 0; j < len(processes); j++ {
				if processes[j].priority >= highPriorityCutoff {
					if processes[j].isInQueue {
						// note: between processes of same duration, processes that arrived first are prioritized
						if processes[j].duration < lowestDuration {
							processId = j
							lowestDuration = processes[j].duration
						}
					}
				}
			}
		}

		fmt.Println("next process to execute has id: ", processId)

		// return if there are no more processes to execute
		if processId == -1 {
			if numProcessesComplete < len(processes) {
				currentTime++
				continue
			} else {
				return
			}
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
