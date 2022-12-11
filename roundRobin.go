package main

import (
	"container/list"
	"fmt"
	"sort"
)

// prints the contents of the queue
func PrintQueue(readyQueue *list.List) {
	fmt.Println("-------------------------------------------")
	fmt.Println("THIS IS THE QUEUE")
	for e := readyQueue.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
		if e.Next() != nil {
			fmt.Print(", ")
		}
	}
	fmt.Println("\n-------------------------------------------")
}

// runs the list of processes for a maximum of totalTime seconds in
// accordance to the round robin scheduling algorithm
func RoundRobin(processes []Process, totalTime int, timeQuantum int) {
	fmt.Println("\n\n                         Running Round Robin Scheduling Algorithm...")

	// sorts the list of processes by arrival time
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].arrivalTime < processes[j].arrivalTime
	})

	// fmt.Println("PROCESSES SORTED BY ARRIVAL TIME")
	// printProcesses(processes)
	// fmt.Println("--------------")

	// update round robin IDs
	for j := 0; j < len(processes); j++ {
		processes[j].roundRobinID = j
	}

	// creates the queue of process IDs
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
			if !processes[j].completed {
				// process was not in queue -> add it
				if !processes[j].isInQueue {
					processes[j].isInQueue = true
					readyQueue.PushBack(j)
				}
				// no need to update any processes who haven't arrived yet
			}
		}

		fmt.Println("queue supposedly updated...")
		PrintQueue(readyQueue)
		printProcesses(processes)
		fmt.Println("--------------")

		// find the next process to execute
		processId := -1
		if readyQueue.Len() != 0 {
			processId = readyQueue.Front().Value.(int)
		}

		fmt.Println("next process to execute has id: ", processId)

		// return if there are no more processes to execute
		if processId == -1 {
			if numProcessesComplete < len(processes) {
				currentTime++
				continue
			} else {
				fmt.Println("no more processes to execute")
				return
			}
		}

		timeIncrement := 0

		// execute the next process, mark as completed, take it off the queue
		if currentTime+processes[processId].duration <= totalTime {
			// if process can be executed within time quantum
			if (processes[processId].duration - processes[processId].secondsCompleted) <= timeQuantum {
				timeIncrement = processes[processId].duration - processes[processId].secondsCompleted
				processes[processId].secondsCompleted += timeIncrement
				currentTime += timeIncrement
				processes[processId].completed = true
				numProcessesComplete++
				processes[processId].isInQueue = false
				processes[processId].turnaroundTime += timeIncrement
				readyQueue.Remove(readyQueue.Front())
			} else {
				// if remaining time of process is longer than time quantum
				timeIncrement = timeQuantum
				currentTime += timeIncrement
				processes[processId].turnaroundTime += timeIncrement
				processes[processId].secondsCompleted += timeIncrement
				processes[processId].waitingTime += timeIncrement
				frontVal := readyQueue.Front().Value
				readyQueue.Remove(readyQueue.Front())
				readyQueue.PushBack(frontVal)
			}
		} else {
			break
		}

		// update the waiting and turnaround time of all processes in the queue
		for e := readyQueue.Front(); e != nil; e = e.Next() {
			id := e.Value.(int)
			if id != processId {
				processes[id].waitingTime += timeIncrement
				processes[id].turnaroundTime += timeIncrement
			}
		}

		if numProcessesComplete >= len(processes) {
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
