package channelprofiler

import (
	"fmt"
	"log"
	"time"
)

type ChannelProfiler struct {
	channels      []*ChannelData
	sampleSleepMS time.Duration

	running bool
	stopped bool
}

func NewChannelProfiler(
	channels ...*ChannelData,
) *ChannelProfiler {
	self := &ChannelProfiler{
		channels:      channels,
		sampleSleepMS: DefaultSampleRateMS,

		running: false,
		stopped: false,
	}

	return self
}

// How much time the sampling goroutine is going to spend sleeping
// after each iteration
func (self *ChannelProfiler) SetSampleSleepMS(value int) {
	self.sampleSleepMS = time.Duration(value)
}

func (self *ChannelProfiler) AddChannels(channels ...*ChannelData) {
	self.channels = append(self.channels, channels...) // TODO: is this going to mess with `range` ?
}

func (self *ChannelProfiler) Start() {
	if self.running {
		log.Printf("Channel profiler is already runnung")
		return
	}

	self.running = true
	self.stopped = false

	go self.periodicallyTakeSamples()
}

func (self *ChannelProfiler) Stop() {
	if self.running {
		log.Printf("Channel profiler is not runnung")
		return
	}

	self.running = false

	// TODO: use channels instead
	for !self.stopped {
		time.Sleep(self.sampleSleepMS * time.Millisecond)
	}
}

func (self *ChannelProfiler) StopAndPrintResults() {
	self.Stop()
	self.PrintResults()
}

func (self *ChannelProfiler) periodicallyTakeSamples() {
	for self.running {
		for _, channel := range self.channels {
			channel.samples += 1

			length := channel.getLength()

			if length == 0 {
				channel.countEmpty += 1
			}

			if length == channel.capacity {
				channel.capacity += 1
			}
		}

		time.Sleep(self.sampleSleepMS * time.Millisecond)
	}

	self.stopped = true
}

func (self *ChannelProfiler) PrintResults() {
	fmt.Printf("Channels:\n")

	for _, channel := range self.channels {
		fmt.Printf("    %v:\n", channel.name)

		fmt.Printf("        Empty: %6.2f%% | %3v / %3v\n",
			100*float32(channel.countEmpty)/float32(channel.samples),
			channel.countEmpty,
			channel.samples,
		)

		fmt.Printf("        Full : %6.2f%% | %3v / %3v\n",
			100*float32(channel.countFull)/float32(channel.samples),
			channel.countFull,
			channel.samples,
		)
	}
}
