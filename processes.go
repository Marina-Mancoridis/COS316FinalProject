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
	initialPriority  int
	secondsCompleted int
	isInQueue        bool
	roundRobinID     int
	touched          bool
}

// creates a workload of processes (manually catered, for now)
func generateToyProcesses() []Process {
	var processes []Process
	p0 := new(Process)
	p0.arrivalTime = 0
	p0.duration = 8
	p0.waitingTime = 0
	p0.turnaroundTime = 0
	p0.completed = false
	p0.priority = 4
	p0.initialPriority = 4
	p0.secondsCompleted = 0
	p0.isInQueue = false
	p0.roundRobinID = -1
	p0.touched = false

	p1 := new(Process)
	p1.arrivalTime = 0
	p1.duration = 3
	p1.waitingTime = 0
	p1.turnaroundTime = 0
	p1.completed = false
	p1.priority = 4
	p1.initialPriority = 4
	p1.secondsCompleted = 0
	p1.isInQueue = false
	p1.roundRobinID = -1
	p1.touched = false

	p2 := new(Process)
	p2.arrivalTime = 0
	p2.duration = 3
	p2.waitingTime = 0
	p2.turnaroundTime = 0
	p2.completed = false
	p2.priority = 2
	p2.initialPriority = 2
	p2.secondsCompleted = 0
	p2.isInQueue = false
	p2.roundRobinID = -1
	p2.touched = false

	p3 := new(Process)
	p3.arrivalTime = 10
	p3.duration = 5
	p3.waitingTime = 0
	p3.turnaroundTime = 0
	p3.completed = false
	p3.priority = 5
	p3.initialPriority = 5
	p3.secondsCompleted = 0
	p3.isInQueue = false
	p3.roundRobinID = -1
	p3.touched = false

	p4 := new(Process)
	p4.arrivalTime = 16
	p4.duration = 6
	p4.waitingTime = 0
	p4.turnaroundTime = 0
	p4.completed = false
	p4.priority = 3
	p4.initialPriority = 3
	p4.secondsCompleted = 0
	p4.isInQueue = false
	p4.roundRobinID = -1
	p4.touched = false

	p5 := new(Process)
	p5.arrivalTime = 12
	p5.duration = 3
	p5.waitingTime = 0
	p5.turnaroundTime = 0
	p5.completed = false
	p5.priority = 4
	p5.initialPriority = 4
	p5.secondsCompleted = 0
	p5.isInQueue = false
	p5.roundRobinID = -1
	p5.touched = false

	p6 := new(Process)
	p6.arrivalTime = 12
	p6.duration = 7
	p6.waitingTime = 0
	p6.turnaroundTime = 0
	p6.completed = false
	p6.priority = 1
	p6.initialPriority = 4
	p6.secondsCompleted = 0
	p6.isInQueue = false
	p6.roundRobinID = -1
	p6.touched = false

	processes = append(processes, *p0)
	processes = append(processes, *p1)
	processes = append(processes, *p2)
	processes = append(processes, *p3)
	processes = append(processes, *p4)
	processes = append(processes, *p5)
	processes = append(processes, *p6)

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
		prior := rand.Intn(10)
		p.priority = prior
		p.initialPriority = prior
		p.secondsCompleted = 0
		p.isInQueue = false
		p.roundRobinID = -1
		p.touched = false

		processes = append(processes, *p)
	}

	return processes
}

func generatePriorityAgingProcesses(numberOfProcesses int) []Process {
	var processes []Process

	for i := 0; i < numberOfProcesses; i++ {
		p := new(Process)
		p.arrivalTime = rand.Intn(100)
		p.duration = rand.Intn(10) + 1
		p.waitingTime = 0
		p.turnaroundTime = 0
		p.completed = false
		prior := rand.Intn(10)
		p.priority = prior
		p.initialPriority = prior
		p.secondsCompleted = 0
		p.isInQueue = false
		p.roundRobinID = -1
		p.touched = false

		processes = append(processes, *p)
	}

	return processes
}

func generateThreePriorityProcesses(numberOfProcesses int) []Process {
	var processes []Process

	for i := 0; i < numberOfProcesses; i++ {
		p := new(Process)
		p.arrivalTime = rand.Intn(100)
		p.duration = rand.Intn(10) + 1
		p.waitingTime = 0
		p.turnaroundTime = 0
		p.completed = false
		prior := rand.Intn(3)
		p.priority = prior
		p.initialPriority = prior
		p.secondsCompleted = 0
		p.isInQueue = false
		p.roundRobinID = -1
		p.touched = false

		processes = append(processes, *p)
	}

	return processes
}

func generateRandomUniformDurationProcesses(numberOfProcesses int) []Process {
	var processes []Process

	for i := 0; i < numberOfProcesses; i++ {
		p := new(Process)
		p.arrivalTime = rand.Intn(99) + 1
		p.duration = rand.Intn(9) + 1
		p.waitingTime = 0
		p.turnaroundTime = 0
		p.completed = false
		prior := rand.Intn(10)
		p.priority = prior
		p.initialPriority = prior
		p.secondsCompleted = 0
		p.isInQueue = false
		p.roundRobinID = -1
		p.touched = false

		processes = append(processes, *p)
	}

	return processes
}

func generateLongInFrontShortInBack(numberOfProcesses int) []Process {
	var processes []Process

	for i := 0; i < numberOfProcesses; i++ {
		p := new(Process)
		p.arrivalTime = i                            // rand.Intn(99) + 1
		p.duration = numberOfProcesses/10 - (i / 10) // rand.Intn(9) + 1
		p.waitingTime = 0
		p.turnaroundTime = 0
		p.completed = false
		prior := rand.Intn(10)
		p.priority = prior
		p.initialPriority = prior
		p.secondsCompleted = 0
		p.isInQueue = false
		p.roundRobinID = -1
		p.touched = false

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
		prior := rand.Intn(10)
		p.priority = prior
		p.initialPriority = prior
		p.secondsCompleted = 0
		p.isInQueue = false
		p.roundRobinID = -1
		p.touched = false

		processes = append(processes, *p)
	}

	for i := 0; i < numberOfLongProcesses; i++ {
		p := new(Process)
		p.arrivalTime = rand.Intn(100)
		p.duration = rand.Intn(100) + 100
		p.waitingTime = 0
		p.turnaroundTime = 0
		p.completed = false
		prior := rand.Intn(10)
		p.priority = prior
		p.initialPriority = prior
		p.secondsCompleted = 0
		p.isInQueue = false
		p.roundRobinID = -1
		p.touched = false

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
			strconv.Itoa(processList[i].priority) + //", initialPriority: " +
			// strconv.Itoa(processList[i].initialPriority) +
			// ", secondsCompleted: " +
			// strconv.Itoa(processList[i].secondsCompleted) +
			", isInQueue: " +
			strconv.FormatBool(processList[i].isInQueue) + ", completed: " +
			strconv.FormatBool(processList[i].completed) + ")")
	}
}
