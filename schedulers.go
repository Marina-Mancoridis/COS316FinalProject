package main

import (
	"fmt"
	"time"
	"sort"
	"container/list"
)

// generates throughput, average waiting time, and average turnaround
// statistics after a single call of a scheduling algorithm
func GenerateStatistics(elapsedTime time.Duration, processes []Process) {
	fmt.Println("STATISTICS:")

	// grabbing information from pass of processes
	var processesCompleted float64 = 0
	var totalWaitingTime = time.Duration(0)
	var totalTurnaroundTime = time.Duration(0)
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
	throughput := float64(processesCompleted) / float64(elapsedTime) * 1000000000
	fmt.Println("throughput (processes/s):            ", throughput)

	// average waiting time per process, in seconds/process
	var averageWaitingTime = float64(totalWaitingTime) / (processesCompleted * 1000000000)
	fmt.Println("average waiting time (s/process):    ", averageWaitingTime)

	// average turnaround time per process, in seconds/process
	var averageTurnaroundTime = float64(totalTurnaroundTime) / (processesCompleted * 1000000000)
	fmt.Println("average turnaround time (s/process): ", averageTurnaroundTime)
}



// runs the list of processes for a maximum of totalTime seconds in 
// accordance to the first come first serve scheduling algorithm
func FirstComeFirstServe(processes []Process, totalTime time.Duration) {
	fmt.Println("\n\n                         Running First Come First Serve Scheduling Algorithm...")

	// sorts the list of processes by arrival time
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].arrivalTime < processes[j].arrivalTime
	})
	
	fmt.Println("processes...")
	printProcesses(processes)
	
	i := 0
	start := time.Now()

	for time.Since(start) < (totalTime) {
		fmt.Println("starting process ", i)

		processes[i].waitingTime += time.Since(start)
		time.Sleep(time.Duration(processes[i].duration) * time.Second)
		processes[i].completed = true
		processes[i].turnaroundTime += time.Since(start)

		if i >= len(processes) {
			break
		}
		i++
	}
	elapsedTime := time.Since(start)
	fmt.Println("\n")

	// outputs statistics for first come first serve scheduling algorithm
	GenerateStatistics(elapsedTime, processes)
}



// runs the list of processes for a maximum of totalTime seconds in 
// accordance to the shortest job first scheduling algorithm
func ShortestJobFirst(processes []Process, totalTime time.Duration) {
	fmt.Println("\n\n                         Running Shortest Job First Scheduling Algorithm...")

	// sorts the list of processes by duration
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].duration < processes[j].duration
	})

	fmt.Println("processes...")
	printProcesses(processes)

	i := 0
	start := time.Now()

	for time.Since(start) < (totalTime) {
		fmt.Println("starting process ", i)

		processes[i].waitingTime += time.Since(start)
		time.Sleep(time.Duration(processes[i].duration) * time.Second)
		processes[i].completed = true
		processes[i].turnaroundTime += time.Since(start)

		if i >= len(processes) {
			break
		}
		i++
	}
	elapsedTime := time.Since(start)
	fmt.Println("\n")

	// outputs statistics for shortest job first scheduling algorithm
	GenerateStatistics(elapsedTime, processes)
}



// runs the list of processes for a maximum of totalTime seconds in 
// accordance to the priority scheduling algorithm
// uses low numbers as high priority, with 1 as highest priority
func Priority(processes []Process, totalTime time.Duration) {
	fmt.Println("\n\n                         Running Priority Scheduling Algorithm...")

	// sorts the list of processes by duration
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].priority < processes[j].priority
	})

	fmt.Println("processes...")
	printProcesses(processes)

	i := 0
	start := time.Now()

	for time.Since(start) < (totalTime) {
		fmt.Println("starting process ", i)

		processes[i].waitingTime += time.Since(start)
		// before sleeping, check if possible to sleep for that long, and
		// if not, we break from the for loop
		
		time.Sleep(time.Duration(processes[i].duration) * time.Second)
		processes[i].completed = true
		processes[i].turnaroundTime += time.Since(start)

		if i >= len(processes) {
			break
		}
		i++
	}
	elapsedTime := time.Since(start)
	fmt.Println("\n")

	// outputs statistics for priority scheduling algorithm
	GenerateStatistics(elapsedTime, processes)
}



// runs the list of processes for a maximum of totalTime seconds in 
// accordance to the round robin scheduling algorithm
// sleep for the time quantum amount of time, then subtract
func RoundRobin(processes []Process, totalTime time.Duration, timeQuantum time.Duration) {
	fmt.Println("\n\n                         Running Round Robin Scheduling Algorithm...")

	// sorts the list of processes by arrival time
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].arrivalTime < processes[j].arrivalTime
	})

	fmt.Println("processes...")
	printProcesses(processes)

	// puts processes into a readyQueue
	readyQueue := list.New()
	for i := 0; i < len(processes); i++ {
		readyQueue.PushBack(processes[i])
	}


	i := 0
	start := time.Now()

	// TODO: needs to "cut" the sleep of a process... we don't know how 
	// to do that still, and it's a problem we've been running into
	for time.Since(start) < (totalTime) {
		fmt.Println("starting process ", i)

		processes[i].waitingTime += time.Since(start)

		// TODO: needs to use time quantum and potentially the queue!!!
		time.Sleep(time.Duration(processes[i].duration) * time.Second)
		processes[i].completed = true
		processes[i].turnaroundTime += time.Since(start)

		if i >= len(processes) {
			break
		}
		i++
	}
	elapsedTime := time.Since(start)
	fmt.Println("\n")

	// outputs statistics for round robin scheduling algorithm
	GenerateStatistics(elapsedTime, processes)
}