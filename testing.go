package main

import (
	"errors"
	"fmt"
)

// ensure processes are valid
func testValidityOfProcesses(processes []Process) (string, error) {
	for i := 0; i < len(processes); i++ {
		p := processes[i]
		if p.arrivalTime < 0 {
			return "TESTING ========> TEST FAILED", errors.New("negative arrivalTime")
		}

		if p.duration < 0 {
			return "TESTING ========> TEST FAILED", errors.New("negative duration")
		}

		if p.priority < 0 {
			return "TESTING ========> TEST FAILED", errors.New("negative priority")
		}

		if p.secondsCompleted < 0 {
			return "TESTING ========> TEST FAILED", errors.New("negative secondsCompleted")
		}

		if p.waitingTime < 0 {
			return "TESTING ========> TEST FAILED", errors.New("negative waitingTime")
		}

		if p.turnaroundTime < 0 {
			return "TESTING ========> TEST FAILED", errors.New("negative turnaroundTime")
		}

		if p.initialPriority < p.priority {
			return "TESTING ========> TEST FAILED", errors.New("priority increased as a result of algorithm")
		}

	}

	success := "TESTING ========> processes are valid"
	return success, nil
}

// test if two values are close to one another
func closeTo(val1 float64, val2 float64) bool {
	if val1 > val2 {
		return (val1-val2 < 0.00000001)
	} else {
		return (val2-val1 < 0.00000001)
	}
}

// test correctness of algorithms on toy processes
func TestCorrectnessOnToyProcesses() (string, error) {
	processes1 := generateToyProcesses()
	processes2 := make([]Process, len(processes1))
	copy(processes2, processes1)
	processes3 := make([]Process, len(processes1))
	copy(processes3, processes1)
	processes4 := make([]Process, len(processes1))
	copy(processes4, processes1)
	processes5 := make([]Process, len(processes1))
	copy(processes5, processes1)
	processes6 := make([]Process, len(processes1))
	copy(processes6, processes1)

	// test shortest job first
	ShortestJobFirst(processes1, 20)
	stats := GenerateStatistics(20, processes1)
	fmt.Println("sjf stats: ", stats)
	if stats.numCompleted != 4 || stats.throughput != 0.2 || stats.avgWaitingTime != 2.75 ||
		stats.averageTurnaroundTime != 7 || stats.numStarved != 3 || stats.avgDurationCompleted != 4.25 ||
		stats.avgDurationStarved != 6 || stats.avgPriorityCompleted != 3.5 || stats.avgPriorityStarved != 4 {
		return "TESTING ========> TEST FAILED", errors.New("shortest job first does not work for toy processes")
	}

	// test first come first serve
	FirstComeFirstServe(processes2, 20)
	stats = GenerateStatistics(20, processes2)
	fmt.Println("fcfs stats: ", stats)
	if stats.numCompleted != 4 || stats.throughput != 0.2 || stats.avgWaitingTime != 5.75 ||
		stats.averageTurnaroundTime != 10.5 || stats.numStarved != 3 || stats.avgDurationCompleted != 4.75 ||
		!closeTo(stats.avgDurationStarved, 5.333333333333333) || stats.avgPriorityCompleted != 3.75 || !closeTo(stats.avgPriorityStarved, 3.6666666666666665) {
		return "TESTING ========> TEST FAILED123", errors.New("shortest job first does not work for toy processes")

	}

	// test round robin
	RoundRobin(processes3, 20, 3)
	stats = GenerateStatistics(20, processes3)
	fmt.Println("rr stats: ", stats)
	if stats.numCompleted != 4 || stats.throughput != 0.2 || stats.avgWaitingTime != 6.5 ||
		stats.averageTurnaroundTime != 9.25 || stats.numStarved != 3 || stats.avgDurationCompleted != 4.25 ||
		stats.avgDurationStarved != 6 || stats.avgPriorityCompleted != 3.5 || stats.avgPriorityStarved != 4 {
		return "TESTING ========> TEST FAILED123", errors.New("shortest job first does not work for toy processes")

	}

	// test priority
	Priority(processes4, 20)
	stats = GenerateStatistics(20, processes4)
	fmt.Println("priority stats: ", stats)
	if stats.numCompleted != 3 || stats.throughput != 0.15 || !closeTo(stats.avgWaitingTime, 4.666666666666667) ||
		!closeTo(stats.averageTurnaroundTime, 9.333333333333334) || stats.numStarved != 4 || !closeTo(stats.avgDurationCompleted, 4.666666666666667) ||
		stats.avgDurationStarved != 5.25 || !closeTo(stats.avgPriorityCompleted, 3.3333333333333335) || stats.avgPriorityStarved != 4 {
		return "TESTING ========> TEST FAILED123", errors.New("shortest job first does not work for toy processes")
	}

	// test priority with aging
	PriorityWithAging(processes5, 20)
	stats = GenerateStatistics(20, processes5)
	fmt.Println("priorityWithAging stats: ", stats)
	if stats.numCompleted != 3 || stats.throughput != 0.15 || !closeTo(stats.avgWaitingTime, 4.666666666666667) ||
		!closeTo(stats.averageTurnaroundTime, 9.333333333333334) || stats.numStarved != 4 || !closeTo(stats.avgDurationCompleted, 4.666666666666667) ||
		stats.avgDurationStarved != 5.25 || !closeTo(stats.avgPriorityCompleted, 3.3333333333333335) || stats.avgPriorityStarved != 4 {
		return "TESTING ========> TEST FAILED123", errors.New("shortest job first does not work for toy processes")
	}

	// test mlq
	MultiLevelQueue(processes6, 20, 5, 7)
	stats = GenerateStatistics(20, processes6)
	fmt.Println("mlq stats: ", stats)
	if stats.numCompleted != 4 || stats.throughput != 0.2 || stats.avgWaitingTime != 2.75 ||
		stats.averageTurnaroundTime != 7 || stats.numStarved != 3 || stats.avgDurationCompleted != 4.25 ||
		stats.avgDurationStarved != 6 || stats.avgPriorityCompleted != 3.5 || stats.avgPriorityStarved != 4 {
		return "TESTING ========> TEST FAILED123", errors.New("shortest job first does not work for toy processes")
	}

	success := "TESTING ========> all algorithms work for toy processes"
	return success, nil
}
