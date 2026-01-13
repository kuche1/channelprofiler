package main

import (
	"fmt"
	"time"

	"github.com/kuche1/channelprofiler"
)

func main() {
	channelProfiler := channelprofiler.NewChannelProfiler()
	// channelProfiler.SetSampleSleepMS(1_234)

	chanA := make(chan int, 2)

	channelProfiler.AddChannels(channelprofiler.NewChannelData(
		"chanA",
		func() int { return len(chanA) },
		cap(chanA),
	))

	channelProfiler.Start()
	defer channelProfiler.StopAndPrintResults()

	time.Sleep(500 * time.Millisecond)

	fmt.Printf("chanA: 0 item(s)\n")
	channelProfiler.PrintResults()

	chanA <- 5

	time.Sleep(500 * time.Millisecond)

	fmt.Printf("\n")
	fmt.Printf("chanA: 1 item(s)\n")
	channelProfiler.PrintResults()

	chanA <- 8

	time.Sleep(500 * time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	time.Sleep(500 * time.Millisecond)

	fmt.Printf("\n")
	fmt.Printf("chanA: 2 item(s)\n")

	chanB := make(chan int, 2)

	channelProfiler.AddChannels(channelprofiler.NewChannelData(
		"chanB",
		func() int { return len(chanB) },
		cap(chanB),
	))

	time.Sleep(500 * time.Millisecond)

	fmt.Printf("\n")
	fmt.Printf("chanA: 2 item(s)\n")
	fmt.Printf("chanB: 0 item(s)\n")
	channelProfiler.PrintResults()

	<-chanA
	chanB <- 26

	time.Sleep(500 * time.Millisecond)

	fmt.Printf("\n")
	fmt.Printf("chanA: 1 item(s)\n")
	fmt.Printf("chanB: 1 item(s)\n")
	channelProfiler.PrintResults()

	chanB <- 432

	time.Sleep(500 * time.Millisecond)

	fmt.Printf("\n")
	fmt.Printf("chanA: 1 item(s)\n")
	fmt.Printf("chanB: 2 item(s)\n")
	channelProfiler.PrintResults()
}
