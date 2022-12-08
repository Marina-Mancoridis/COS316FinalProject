package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

// A process
type Process struct {
	arrivalTime    int
	duration       int
	waitingTime    time.Duration
	turnaroundTime time.Duration
	completed      bool
}



// creates a workload of processes (manually catered, for now)
func generateProcesses() []Process {
	var processes []Process
	p1 := new(Process)
	p1.arrivalTime = 3
	p1.duration = 1
	p1.waitingTime = 0 * time.Second
	p1.turnaroundTime = 0 * time.Second
	p1.completed = false

	p2 := new(Process)
	p2.arrivalTime = 2
	p2.duration = 1
	p2.waitingTime = 0 * time.Second
	p2.turnaroundTime = 0 * time.Second
	p2.completed = false

	p3 := new(Process)
	p3.arrivalTime = 1
	p3.duration = 6
	p3.waitingTime = 0 * time.Second
	p3.turnaroundTime = 0 * time.Second
	p3.completed = false

	processes = append(processes, *p1)
	processes = append(processes, *p2)
	processes = append(processes, *p3)

	// printProcesses(processes)

	// sorts the list of processes by arrival time
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].arrivalTime < processes[j].arrivalTime
	})

	return processes
}


// prints a workload of processes in a readable way
func printProcesses(processList []Process) {
	for i := 0; i < len(processList); i++ {
		fmt.Println("(arrivalTime: " + strconv.Itoa(processList[i].arrivalTime) + ", duration: " + strconv.Itoa(processList[i].duration) + ")")
	}
}



// runs a workload of processes on a CPU with different scheduling
// algorithms, outputting statistics of how the CPU runs under each 
// algorithm
func main() {
	processes := generateProcesses()
	FirstComeFirstServe(processes, 5*time.Second)

	processesAgain := generateProcesses()
	ShortestJobFirst(processesAgain, 5*time.Second)
}
