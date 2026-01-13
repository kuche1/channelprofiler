package channelprofiler

import "log"

type ChannelData struct {
	name string

	getLength func() int
	capacity  int

	samples    int
	countEmpty int
	countFull  int
}

func NewChannelData(name string, getLength func() int, capacity int) *ChannelData {
	if capacity <= 0 {
		log.Printf("Channel `%v`: Cannot profile a channel with capacity <=0 (in this case %v)", name, capacity)
	}

	return &ChannelData{
		name: name,

		getLength: getLength,
		capacity:  capacity,

		samples:    0,
		countEmpty: 0,
		countFull:  0,
	}
}
