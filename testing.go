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
			return "", errors.New("negative arrivalTime")
		}

		if p.duration < 0 {
			return "", errors.New("negative duration")
		}

		if p.priority < 0 {
			return "", errors.New("negative priority")
		}

		if p.secondsCompleted < 0 {
			return "", errors.New("negative secondsCompleted")
		}

		if p.waitingTime < 0 {
			return "", errors.New("negative waitingTime")
		}

		if p.turnaroundTime < 0 {
			return "", errors.New("negative turnaroundTime")
		}

		if p.initialPriority < p.priority {
			return "", errors.New("priority increased as a result of algorithm")
		}

	}

	success := "TESTING ========> processes are valid"
	return success, nil
}

// test correctness of algorithms on toy processes
func TestCorrectnessOnToyProcesses() {
	processes1 := generateToyProcesses()
	// processes2 := make([]Process, len(processes1))
	// copy(processes2, processes1)

	FirstComeFirstServe(processes1, 10)
	stats := GenerateStatistics(10, processes1)
	fmt.Println(stats)

}
