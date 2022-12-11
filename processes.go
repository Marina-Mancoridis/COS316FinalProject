package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

// A process
type Process struct {
	arrivalTime      int
	duration         int
	waitingTime      int
	turnaroundTime   int
	completed        bool
	priority         int
	secondsCompleted int
	isInQueue        bool
}

// creates a workload of processes (manually catered, for now)
func generateToyProcesses() []Process {
	var processes []Process
	p1 := new(Process)
	p1.arrivalTime = 3
	p1.duration = 1
	p1.waitingTime = 0
	p1.turnaroundTime = 0
	p1.completed = false
	p1.priority = 2
	p1.secondsCompleted = 0
	p1.isInQueue = false

	p2 := new(Process)
	p2.arrivalTime = 2
	p2.duration = 1
	p2.waitingTime = 0
	p2.turnaroundTime = 0
	p2.completed = false
	p2.priority = 3
	p2.secondsCompleted = 0
	p2.isInQueue = false

	p3 := new(Process)
	p3.arrivalTime = 1
	p3.duration = 6
	p3.waitingTime = 0
	p3.turnaroundTime = 0
	p3.completed = false
	p3.priority = 1
	p3.secondsCompleted = 0
	p3.isInQueue = false

	processes = append(processes, *p1)
	processes = append(processes, *p2)
	processes = append(processes, *p3)

	return processes
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
		p.secondsCompleted = 0
		p.isInQueue = false

		processes = append(processes, *p)
	}

	return processes
}

func generatePriorityAgingProcesses(numberOfProcesses int) []Process {
	var processes []Process

	for i := 0; i < numberOfProcesses; i++ {
		p := new(Process)
		p.arrivalTime = rand.Intn(10)
		p.duration = rand.Intn(9) + 1
		p.waitingTime = 0
		p.turnaroundTime = 0
		p.completed = false
		p.priority = rand.Intn(10)
		p.secondsCompleted = 0
		p.isInQueue = false

		processes = append(processes, *p)
	}

	return processes
}

func generateRandomUniformDurationProcesses(numberOfProcesses int) []Process {
	var processes []Process

	for i := 0; i < numberOfProcesses; i++ {
		p := new(Process)
		p.arrivalTime = rand.Intn(100)
		p.duration = rand.Intn(99) + 1
		p.waitingTime = 0
		p.turnaroundTime = 0
		p.completed = false
		p.priority = rand.Intn(10)
		p.secondsCompleted = 0
		p.isInQueue = false

		processes = append(processes, *p)
	}

	return processes
}

func generateShortLongProcesses(numberOfShortProcesses int, numberOfLongProcesses int) []Process {
	var processes []Process

	for i := 0; i < numberOfShortProcesses; i++ {
		p := new(Process)
		p.arrivalTime = rand.Intn(100)
		p.duration = rand.Intn(9) + 1
		p.waitingTime = 0
		p.turnaroundTime = 0
		p.completed = false
		p.priority = rand.Intn(10)
		p.secondsCompleted = 0
		p.isInQueue = false

		processes = append(processes, *p)
	}

	for i := 0; i < numberOfLongProcesses; i++ {
		p := new(Process)
		p.arrivalTime = rand.Intn(100)
		p.duration = rand.Intn(100) + 100
		p.waitingTime = 0
		p.turnaroundTime = 0
		p.completed = false
		p.priority = rand.Intn(10)
		p.secondsCompleted = 0
		p.isInQueue = false

		processes = append(processes, *p)
	}

	return processes
}

// prints a workload of processes in a readable way
func printProcesses(processList []Process) {
	for i := 0; i < len(processList); i++ {
		fmt.Println("(id: " + strconv.Itoa(i) + ", arrivalTime: " +
			strconv.Itoa(processList[i].arrivalTime) + ", turnaroundTime: " +
			strconv.Itoa(processList[i].turnaroundTime) + ", waitingTime: " +
			strconv.Itoa(processList[i].waitingTime) + ", duration: " +
			strconv.Itoa(processList[i].duration) + ", priority: " +
			strconv.Itoa(processList[i].priority) + ", isInQueue: " +
			strconv.FormatBool(processList[i].isInQueue) + ", completed: " +
			strconv.FormatBool(processList[i].completed) + ")")
	}
}
