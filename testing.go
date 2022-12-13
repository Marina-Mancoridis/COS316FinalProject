package main

import "errors"

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

//
