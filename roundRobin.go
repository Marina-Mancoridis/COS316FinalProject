package main

import (
	"container/list"
	"fmt"
	"math"
	"sort"
)

// runs the list of processes for a maximum of totalTime seconds in
// accordance to the round robin scheduling algorithm
func RoundRobinOld(processes []Process, totalTime int, timeQuantum int) {
	fmt.Println("\n\n                         Running Round Robin Scheduling Algorithm...")

	// sorts the list of processes by arrival time
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].arrivalTime < processes[j].arrivalTime
	})

	// puts processes into a readyQueue
	readyQueue := list.New()
	for i := 0; i < len(processes); i++ {
		readyQueue.PushBack(processes[i])
	}

	fmt.Println("processes...")
	printProcesses(processes)

	i := 0
	currentTime := 0

	for currentTime < totalTime {
		// remove process from front of queue
		front := readyQueue.Front()
		process := *front
		readyQueue.Remove(process)

		// if the process can be executed on time
		if currentTime+process.duration <= totalTime {
			// if process can be executed within time quantum
			if (process.duration - process.secondsCompleted) <= timeQuantum {
				process.waitingTime += currentTime
				currentTime += process.duration - process.secondsCompleted
				process.completed = true
				process.turnaroundTime += currentTime
			} else {
				// if process cannot be executed within time quantum
				process.waitingTime += currentTime
				currentTime += timeQuantum
				process.secondsCompleted += timeQuantum
				readyQueue.PushBack(process)
			}
		} else {
			break
		}

		if i >= len(processes) {
			break
		}
		i++
	}
	fmt.Println("\n")

	// outputs statistics for shortest job first scheduling algorithm
	GenerateStatistics(currentTime, processes)
}

// runs the list of processes for a maximum of totalTime seconds in
// accordance to the round robin scheduling algorithm
func RoundRobin(processes []Process, totalTime int, timeQuantum int) {
	fmt.Println("\n\n                         Running Priority Scheduling Algorithm...")

	// sorts the list of processes by arrival time
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].arrivalTime < processes[j].arrivalTime
	})

	// fmt.Println("PROCESSES SORTED BY ARRIVAL TIME")
	// printProcesses(processes)
	// fmt.Println("--------------")

	// creates the queue of processes
	readyQueue := list.New()

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
	fmt.Println("RIGHT BEFORE ENTIRE FUNCTION RETURNS")
	printProcesses(processes)
	fmt.Println("--------------")

	// outputs statistics
	GenerateStatistics(currentTime, processes)
}
