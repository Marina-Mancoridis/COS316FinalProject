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

	// outputs statistics
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

	// outputs statistics
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

	// outputs statistics
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

// 	// outputs statistics
// 	GenerateStatistics(currentTime, processes)
// }

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

	for currentTime < totalTime {
		// fmt.Println("---------------------------------------------------")
		// fmt.Println("AT TIME STEP ", currentTime)
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

		// fmt.Println("queue and priorities supposedly updated...")
		// printProcesses(processes)
		// fmt.Println("--------------")

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
		// fmt.Println("next process to execute has id: ", processId)

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

	for currentTime < totalTime {
		// fmt.Println("---------------------------------------------------")
		// fmt.Println("AT TIME STEP ", currentTime)
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

		// fmt.Println("queue and priorities supposedly updated...")
		// printProcesses(processes)
		// fmt.Println("--------------")

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
		// fmt.Println("next process to execute has id: ", processId)

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


// runs the list of processes for a maximum of totalTime seconds in
// accordance to the first come first serve scheduling algorithm
func FirstComeFirstServeQueue(processes []Process, totalTime int) {
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

	for currentTime < totalTime {
		// fmt.Println("---------------------------------------------------")
		// fmt.Println("AT TIME STEP ", currentTime)

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

		// fmt.Println("queue and priorities supposedly updated...")
		// printProcesses(processes)
		// fmt.Println("--------------")

		// find the next process to execute
		processId := -1
		for j := 0; j < len(processes); j++ {
			// finds first in queue (queue is in time order...)
			if processes[j].isInQueue {
				processId = j
				break
			}
		}
		// fmt.Println("next process to execute has id: ", processId)

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