package main

// runs a workload of processes on a CPU with different scheduling
// algorithms, outputting statistics of how the CPU runs under each
// algorithm
func main() {
	processes := generatePriorityAgingProcesses(15)
	// processes := generateToyProcesses()
	// processes := generateShortLongProcesses(100, 900)
	// Priority(processes, 500)
	// PriorityWithAging(processes, 100)
	// FirstComeFirstServe(processes, 100)
	ShortestJobFirst(processes, 100)
	// PriorityWithQueue(processes, 100)
	// RoundRobin(processes, 5, 1)
}
