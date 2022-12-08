package main

import (
	"fmt"
	"strconv"
)

// A process
type Process struct {
	arrivalTime    int
	duration       int
	waitingTime    int
	turnaroundTime int
	completed      bool
	priority       int
}

// creates a workload of processes (manually catered, for now)
func generateProcesses() []Process {
	var processes []Process
	p1 := new(Process)
	p1.arrivalTime = 3
	p1.duration = 1
	p1.waitingTime = 0
	p1.turnaroundTime = 0
	p1.completed = false
	p1.priority = 2

	p2 := new(Process)
	p2.arrivalTime = 2
	p2.duration = 1
	p2.waitingTime = 0
	p2.turnaroundTime = 0
	p2.completed = false
	p2.priority = 3

	p3 := new(Process)
	p3.arrivalTime = 1
	p3.duration = 6
	p3.waitingTime = 0
	p3.turnaroundTime = 0
	p3.completed = false
	p3.priority = 1

	processes = append(processes, *p1)
	processes = append(processes, *p2)
	processes = append(processes, *p3)

	return processes
}

// prints a workload of processes in a readable way
func printProcesses(processList []Process) {
	for i := 0; i < len(processList); i++ {
		fmt.Println("(arrivalTime: " +
			strconv.Itoa(processList[i].arrivalTime) + ", duration: " +
			strconv.Itoa(processList[i].duration) + ", priority: " +
			strconv.Itoa(processList[i].priority) + ")")
	}
}

// runs a workload of processes on a CPU with different scheduling
// algorithms, outputting statistics of how the CPU runs under each
// algorithm
func main() {
	processes := generateProcesses()
	FirstComeFirstServe(processes, 7)
	// ShortestJobFirst(processes, 5)
	// Priority(processes, 5)
	// RoundRobin(processes, 5, 1)
}
