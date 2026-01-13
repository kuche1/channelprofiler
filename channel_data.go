package channelprofiler

import (
	"fmt"
	"log"
)

type ChannelData struct {
	name  string
	error string

	getLength func() int
	capacity  int

	samples    int
	countEmpty int
	countFull  int
}

func NewChannelData(name string, getLength func() int, capacity int) *ChannelData {
	error_ := ""

	if capacity <= 0 {
		error_ = fmt.Sprintf("Cannot profile a channel with capacity <=0 (in this case %v)", capacity)
		log.Printf("Channel `%v`: %v", name, error_)
	}

	return &ChannelData{
		name:  name,
		error: error_,

		getLength: getLength,
		capacity:  capacity,

		samples:    0,
		countEmpty: 0,
		countFull:  0,
	}
}
