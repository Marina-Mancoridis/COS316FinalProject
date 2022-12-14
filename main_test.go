package main

import (
	"fmt"
	"testing"
)

// ensure processes are valid
func TestValidityOfProcesses(t *testing.T) {
	processes := generateEqualDistributionProcesses(1000)
	for i := 0; i < len(processes); i++ {
		p := processes[i]
		if p.arrivalTime < 0 {
			panic("negative arrivalTime")
		}

		if p.duration < 0 {
			panic("negative duration")
		}

		if p.priority < 0 {
			panic("negative priority")
		}

		if p.secondsCompleted < 0 {
			panic("negative secondsCompleted")
		}

		if p.waitingTime < 0 {
			panic("negative waitingTime")
		}

		if p.turnaroundTime < 0 {
			panic("negative turnaroundTime")
		}

		if p.initialPriority < p.priority {
			panic("priority increased as a result of algorithm")
		}

	}
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
func TestCorrectnessOnToyProcesses(t *testing.T) {
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
		panic("shortest job first does not work for toy processes")
	}

	// test first come first serve
	FirstComeFirstServe(processes2, 20)
	stats = GenerateStatistics(20, processes2)
	fmt.Println("fcfs stats: ", stats)
	if stats.numCompleted != 4 || stats.throughput != 0.2 || stats.avgWaitingTime != 5.75 ||
		stats.averageTurnaroundTime != 10.5 || stats.numStarved != 3 || stats.avgDurationCompleted != 4.75 ||
		!closeTo(stats.avgDurationStarved, 5.333333333333333) || stats.avgPriorityCompleted != 3.75 || !closeTo(stats.avgPriorityStarved, 3.6666666666666665) {
		panic("fcfs does not work for toy processes")
	}

	// test round robin
	RoundRobin(processes3, 20, 3)
	stats = GenerateStatistics(20, processes3)
	fmt.Println("rr stats: ", stats)
	if stats.numCompleted != 4 || stats.throughput != 0.2 || stats.avgWaitingTime != 6.5 ||
		stats.averageTurnaroundTime != 9.25 || stats.numStarved != 3 || stats.avgDurationCompleted != 4.25 ||
		stats.avgDurationStarved != 6 || stats.avgPriorityCompleted != 3.5 || stats.avgPriorityStarved != 4 {
		panic("round robin does not work for toy processes")

	}

	// test priority
	Priority(processes4, 20)
	stats = GenerateStatistics(20, processes4)
	fmt.Println("priority stats: ", stats)
	if stats.numCompleted != 3 || stats.throughput != 0.15 || !closeTo(stats.avgWaitingTime, 4.666666666666667) ||
		!closeTo(stats.averageTurnaroundTime, 9.333333333333334) || stats.numStarved != 4 || !closeTo(stats.avgDurationCompleted, 4.666666666666667) ||
		stats.avgDurationStarved != 5.25 || !closeTo(stats.avgPriorityCompleted, 3.3333333333333335) || stats.avgPriorityStarved != 4 {
		panic("priority does not work for toy processes")
	}

	// test priority with aging
	PriorityWithAging(processes5, 20)
	stats = GenerateStatistics(20, processes5)
	fmt.Println("priorityWithAging stats: ", stats)
	if stats.numCompleted != 3 || stats.throughput != 0.15 || !closeTo(stats.avgWaitingTime, 4.666666666666667) ||
		!closeTo(stats.averageTurnaroundTime, 9.333333333333334) || stats.numStarved != 4 || !closeTo(stats.avgDurationCompleted, 4.666666666666667) ||
		stats.avgDurationStarved != 5.25 || !closeTo(stats.avgPriorityCompleted, 3.3333333333333335) || stats.avgPriorityStarved != 4 {
		panic("priorityWithAging does not work for toy processes")
	}

	// test mlq
	MultiLevelQueue(processes6, 20, 5, 7)
	stats = GenerateStatistics(20, processes6)
	fmt.Println("mlq stats: ", stats)
	if stats.numCompleted != 4 || stats.throughput != 0.2 || stats.avgWaitingTime != 2.75 ||
		stats.averageTurnaroundTime != 7 || stats.numStarved != 3 || stats.avgDurationCompleted != 4.25 ||
		stats.avgDurationStarved != 6 || stats.avgPriorityCompleted != 3.5 || stats.avgPriorityStarved != 4 {
		panic("mlq does not work for toy processes")
	}

}

func TestCompareSchedulingAlgorithms(t *testing.T) {
	// Compare FCFS and SJF on workload with long processes in front, short in back
	processes1 := generateLongInFrontShortInBack(1000)
	processes2 := make([]Process, len(processes1))
	copy(processes2, processes1)
	FirstComeFirstServe(processes1, 500)
	stats1 := GenerateStatistics(500, processes1)
	ShortestJobFirst(processes2, 500)
	stats2 := GenerateStatistics(500, processes2)

	// we expect FCFS to starve more processes
	if stats1.numStarved < stats2.numStarved {
		panic("FCFS starved fewer processes than SJF for a workload with long processes in front, short in back")
	}

	// we expect the average duration of FCFS completed processes to be greater than SJF
	if stats1.avgDurationCompleted < stats2.avgDurationCompleted {
		panic("FCFS had lower average duration than SJF for a workload with long processes in front, short in back")
	}

	// we expect avg waiting time of FCFS to be greater than SJF
	if stats1.avgWaitingTime < stats2.avgWaitingTime {
		panic("FCFS had lower average waiting time than SJF for a workload with long processes in front, short in back")
	}

	// Compare RR and SJF on a workload of random uniform duration processes
	processes1 = generateRandomUniformDurationProcesses(1000)
	processes2 = make([]Process, len(processes1))
	copy(processes2, processes1)
	RoundRobin(processes1, 500, 3)
	stats1 = GenerateStatistics(500, processes1)
	ShortestJobFirst(processes2, 500)
	stats2 = GenerateStatistics(500, processes2)

	// we expect RR to starve fewer processes
	if stats1.numStarved < stats2.numStarved {
		panic("RR starves more processes than sjf on a random uniform duration workload")
	}

	// we expect RR to have a higher avg waiting time than sjf
	if stats1.avgWaitingTime < stats2.avgWaitingTime {
		panic("RR has a lower average waiting time than SJF on a random uniform duration workload")
	}

	// we expect RR to complete fewer processes than SJF
	if stats1.numCompleted > stats2.numCompleted {
		panic("RR completes more processes than SJF on a random uniform duration workload")
	}

	// Compare Priority and FCFS on 1000 processes w/ priorities 1-10, arrival 1-100, duration 1-10
	processes1 = generatePriorityAgingProcesses(1000)
	processes2 = make([]Process, len(processes1))
	copy(processes2, processes1)
	Priority(processes1, 500)
	stats1 = GenerateStatistics(500, processes1)
	FirstComeFirstServe(processes2, 500)
	stats2 = GenerateStatistics(500, processes2)

	// we expect priority to have lower average priority of completed processes
	if stats1.avgPriorityCompleted > stats2.avgPriorityCompleted {
		panic("Priority has higher avg priority of completed processes than fcfs")
	}

	// we expect priority to have a higher avg priority for starved processes
	if stats1.avgPriorityStarved < stats2.avgPriorityStarved {
		panic("Priority has lower avg priority of starved processes than fcfs")
	}

	// Compare Priority and FCFS on 1000 processes w/ priorities 1-10, arrival 1-100, duration 1-10
	processes1 = generatePriorityAgingProcesses(1000)
	processes2 = make([]Process, len(processes1))
	copy(processes2, processes1)
	PriorityWithAging(processes1, 500)
	stats1 = GenerateStatistics(500, processes1)
	FirstComeFirstServe(processes2, 500)
	stats2 = GenerateStatistics(500, processes2)

	// we expect priority to have lower average priority of completed processes
	if stats1.avgPriorityCompleted > stats2.avgPriorityCompleted {
		panic("PriorityWithAging has higher avg priority of completed processes than fcfs")
	}

	// we expect priority to have a higher avg priority for starved processes
	if stats1.avgPriorityStarved < stats2.avgPriorityStarved {
		panic("PriorityWithAging has lower avg priority of starved processes than fcfs")
	}

	// Compare Priority and Priority with Aging on 1000 processes w/ priorities 1-10,
	// arrival 1-100, duration 1-10
	processes1 = generatePriorityAgingProcesses(1000)
	processes2 = make([]Process, len(processes1))
	copy(processes2, processes1)
	Priority(processes1, 500)
	stats1 = GenerateStatistics(500, processes1)
	PriorityWithAging(processes2, 500)
	stats2 = GenerateStatistics(500, processes2)

	// we expect priority to have lower avg priority of completed processes
	if stats1.avgPriorityCompleted > stats2.avgPriorityCompleted {
		panic("Priority has higher avg priority of completed processes than priorityWithAging")
	}

	// Compare MLQ and FCFS on 1000 processes w/ priorities 1-10, arrival 1-100, duration 1-10
	processes1 = generatePriorityAgingProcesses(1000)
	processes2 = make([]Process, len(processes1))
	copy(processes2, processes1)
	MultiLevelQueue(processes1, 500, 4, 7)
	stats1 = GenerateStatistics(500, processes1)
	FirstComeFirstServe(processes2, 500)
	stats2 = GenerateStatistics(500, processes2)

	// we expect priority to have lower average priority of completed processes
	if stats1.avgPriorityCompleted > stats2.avgPriorityCompleted {
		panic("mlq has higher avg priority of completed processes than fcfs")
	}

	// we expect priority to have a higher avg priority for starved processes
	if stats1.avgPriorityStarved < stats2.avgPriorityStarved {
		panic("mlq has lower avg priority of starved processes than fcfs")
	}

	// Compare MLQ and Priority on workload with only 3 priorities
	processes1 = generateThreePriorityProcesses(1000)
	processes2 = make([]Process, len(processes1))

	for i := 0; i < 100; i++ {
		fmt.Print(processes1[i].priority)
	}
	copy(processes2, processes1)
	MultiLevelQueue(processes1, 500, 1, 2)
	stats1 = GenerateStatistics(500, processes1)
	Priority(processes2, 500)
	stats2 = GenerateStatistics(500, processes2)

	// we expect MLQ to have a shorter waiting time (since MLQ is optimizing for priority AND duration)
	if stats1.avgWaitingTime > stats2.avgWaitingTime {
		panic("MLQ has longer waiting time than Priority with 3-priority workload")
	}

	// we expect MLQ to have a higher number of completed processes than priority
	if stats1.numCompleted < stats2.numCompleted {
		panic("MLQ has lower number of completed processes than priority for 3-priority workload")
	}

	// we expect MLQ to have a lower average priority completed since it will complete more
	if stats1.avgPriorityCompleted > stats2.avgPriorityCompleted {
		panic("MLQ has higher avg priority of completed processes than priority for 3-priority workload")
	}

}
