package main

import (
	"fmt"
	"math/rand"
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

func generateEqualDistributionProcesses(numberOfProcesses int) []Process {
	var processes []Process

	for i := 0; i < numberOfProcesses; i++ {
		p := new(Process)
		p.arrivalTime = rand.Intn(100)
		p.duration = 1
		p.waitingTime = 0
		p.turnaroundTime = 0
		p.completed = false
		p.priority = rand.Intn(10)

		processes = append(processes, *p)
	}

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
