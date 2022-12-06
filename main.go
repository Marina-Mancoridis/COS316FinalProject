package main

import (
	"fmt"
	"strconv"
)

// A process
type Process struct {
	arrivalTime int
	duration    int
}

func generateProcesses() []Process {
	var allProcesses []Process
	p1 := new(Process)
	p1.arrivalTime = 3
	p1.duration = 5

	p2 := new(Process)
	p2.arrivalTime = 2
	p2.duration = 5

	p3 := new(Process)
	p3.arrivalTime = 1
	p3.duration = 5

	allProcesses = append(allProcesses, *p1)
	allProcesses = append(allProcesses, *p2)
	allProcesses = append(allProcesses, *p3)

	printProcesses(allProcesses)

	return allProcesses
}

func printProcesses(processList []Process) {
	for i := 0; i < len(processList); i++ {
		fmt.Println("(arrivalTime: " + strconv.Itoa(processList[i].arrivalTime) + ", duration: " + strconv.Itoa(processList[i].duration) + ")")
	}
}

func main() {
	allProcesses := generateProcesses()
	FirstComeFirstServe(allProcesses, 3)
}
