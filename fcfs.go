package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

func FirstComeFirstServe(processes []Process, totalTime int) {
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

	time.AfterFunc(time.Duration(totalTime)*time.Second, func() {
		fmt.Println("hellooooo")
		os.Exit(0)
	})

	start := time.Now()

	fmt.Println("about to start with processes")
	for i := 0; i < len(processes); i++ {
		fmt.Println("starting process ", i)
		time.Sleep(time.Duration(processes[i].duration) * time.Second)
	}

	elapsed := time.Since(start)

	fmt.Println("elapsed time: ", elapsed)
}
