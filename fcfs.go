package main

import (
	"fmt"
	"sort"
	"time"
)

func FirstComeFirstServe(processes []Process, totalTime time.Duration) {
	fmt.Println("unsorted processes slice:")
	fmt.Printf("%v\n", processes)

	// sorts the list of processes by arrival time
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].arrivalTime < processes[j].arrivalTime
	})

	// start timing now, because CPU wouldn't have to sort processes by
	// their arrival time
	fmt.Println("sorted processes slice:")
	fmt.Printf("%v\n", processes)

	var howManyGotThrough int = 0

	for start := time.Now(); time.Since(start) < (totalTime); {
		for i := 0; i < len(processes); i++ {
			fmt.Println("starting process ", i)
			time.Sleep(time.Duration(processes[i].duration) * time.Second)
			howManyGotThrough = howManyGotThrough + 1
		}

		elapsed := time.Since(start)

		fmt.Println("elapsed time: ", elapsed)
	}

	// calculate latency (for what's done so far) and bandwith
	fmt.Println("got through this many: ", howManyGotThrough)

}
