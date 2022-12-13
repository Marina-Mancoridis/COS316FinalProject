package main

import "fmt"

func copyProcesses(processes []Process) []Process {
	var newProcesses []Process

	for i := 0; i < len(processes); i++ {
		p := new(Process)
		p.arrivalTime = processes[i].arrivalTime
		p.duration = processes[i].duration
		p.waitingTime = processes[i].waitingTime
		p.turnaroundTime = processes[i].turnaroundTime
		p.completed = processes[i].completed
		p.priority = processes[i].priority
		p.initialPriority = processes[i].initialPriority
		p.secondsCompleted = processes[i].secondsCompleted
		p.isInQueue = processes[i].isInQueue
		p.roundRobinID = processes[i].roundRobinID

		newProcesses = append(newProcesses, *p)
	}

	return processes
}

// runs a workload of processes on a CPU with different scheduling
// algorithms, outputting statistics of how the CPU runs under each
// algorithm
func main() {
	// processes1 := generatePriorityAgingProcesses(100)
	// processes := generateToyProcesses()
	// processes := generateShortLongProcesses(100, 900)
	// processes1 := generateRandomUniformDurationProcesses(4000)
	processes1 := generateLongInFrontShortInBack(100)
	str, _ := testValidityOfProcesses(processes1)
	fmt.Println(str)
	processes2 := make([]Process, len(processes1))
	copy(processes2, processes1)
	processes3 := make([]Process, len(processes1))
	copy(processes3, processes1)
	processes4 := make([]Process, len(processes1))
	copy(processes4, processes1)
	processes5 := make([]Process, len(processes1))
	copy(processes5, processes1)

	FirstComeFirstServe(processes1, 300)
	ShortestJobFirst(processes2, 300)
	RoundRobin(processes3, 300, 2)
	// Priority(processes4, 150)
	// PriorityWithAging(processes5, 150)
	// MultiLevelQueue(processes, 100, 5, 7)
}
