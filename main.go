package main

// runs a workload of processes on a CPU with different scheduling
// algorithms, outputting statistics of how the CPU runs under each
// algorithm
func main() {
	processes := generatePriorityAgingProcesses(15)
	// printProcesses(processes)
	// processes := generateShortLongProcesses(100, 900)
	// FirstComeFirstServe(processes, 500)
	// ShortestJobFirst(processes, 500)
	// Priority(processes, 500)
	PriorityWithAging(processes, 100)
	// RoundRobin(processes, 5, 1)
}
