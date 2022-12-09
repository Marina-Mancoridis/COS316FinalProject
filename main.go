package main

// runs a workload of processes on a CPU with different scheduling
// algorithms, outputting statistics of how the CPU runs under each
// algorithm
func main() {
	// processes := generateToyProcesses()
	// printProcesses(processes)
	processes := generateShortLongProcesses(5, 5)
	printProcesses(processes)
	// FirstComeFirstServe(processes, 7)
	// ShortestJobFirst(processes, 5)
	// Priority(processes, 5)
	RoundRobin(processes, 5, 1)
}
