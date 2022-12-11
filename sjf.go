package main

import (
	"fmt"
	"sort"
)

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

	// outputs statistics
	GenerateStatistics(currentTime, processes)
}

// runs the list of processes for a maximum of totalTime seconds in
// accordance to the shortest job first scheduling algorithm
func ShortestJobFirstWithQueue(processes []Process, totalTime int) {
	fmt.Println("\n\n                         Running Shortest Job First (With Queue) Scheduling Algorithm...")

	// sorts the list of processes by duration
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].duration < processes[j].duration
	})

	fmt.Println("PROCESSES SORTED BY DURATION")
	printProcesses(processes)
	fmt.Println("--------------")

	i := 0
	currentTime := 0

	for currentTime < totalTime {
		fmt.Println("---------------------------------------------------")
		fmt.Println("AT TIME STEP ", currentTime)

		// add first process to queue !! this is new.
		processes[0].isInQueue = true

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

		fmt.Println("queue supposedly updated...")
		printProcesses(processes)
		fmt.Println("--------------")

		// find the next process to execute
		processId := -1
		for j := 0; j < len(processes); j++ {
			// finds first in queue (queue is in duration order...)
			if processes[j].isInQueue {
				processId = j
				break
			}
		}
		fmt.Println("next process to execute has id: ", processId)

		// return if there are no more processes to execute
		if processId == -1 {
			return
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
	fmt.Println("\n")

	// outputs statistics
	GenerateStatistics(currentTime, processes)
}
