package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

// A process
type Process struct {
	arrivalTime int
	duration    int
	waitingTime time.Duration
	completed   bool
}

func generateProcesses() []Process {
	var allProcesses []Process
	p1 := new(Process)
	p1.arrivalTime = 3
	p1.duration = 5
	p1.waitingTime = 0 * time.Second
	p1.completed = false

	p2 := new(Process)
	p2.arrivalTime = 2
	p2.duration = 5
	p2.waitingTime = 0 * time.Second
	p2.completed = false

	p3 := new(Process)
	p3.arrivalTime = 1
	p3.duration = 5
	p3.waitingTime = 0 * time.Second
	p3.completed = false

	allProcesses = append(allProcesses, *p1)
	allProcesses = append(allProcesses, *p2)
	allProcesses = append(allProcesses, *p3)

	// printProcesses(allProcesses)

	// sorts the list of processes by arrival time
	sort.Slice(allProcesses, func(i, j int) bool {
		return allProcesses[i].arrivalTime < allProcesses[j].arrivalTime
	})

	return allProcesses
}

func printProcesses(processList []Process) {
	for i := 0; i < len(processList); i++ {
		fmt.Println("(arrivalTime: " + strconv.Itoa(processList[i].arrivalTime) + ", duration: " + strconv.Itoa(processList[i].duration) + ")")
	}
}

func main() {
	start := time.Now()
	allProcesses := generateProcesses()

	FirstComeFirstServe(allProcesses, 8*time.Second)
	fmt.Println("STOPPED")
	elapsed := time.Since(start)
	fmt.Println("TOTAL ELAPSED TIME OF MAIN FN: ", elapsed)
}
