package main

// runs a workload of processes on a CPU with different scheduling
// algorithms, outputting statistics of how the CPU runs under each
// algorithm
func main() {
	// processes := generateToyProcesses()
	// printProcesses(processes)
	processes := generateEqualDistributionProcesses(10)
	// FirstComeFirstServe(processes, 10000)
	// ShortestJobFirst(processes, 10000)
	Priority(processes, 10000)
	// RoundRobin(processes, 5, 1)
}
