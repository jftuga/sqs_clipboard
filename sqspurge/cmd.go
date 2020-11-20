package main

import (
	"github.com/jftuga/copypaste"
	"github.com/jftuga/copypaste/queue"
)

func main() {
	queueURL := queue.GetQueueURL()
	cp := copypaste.New(queueURL)
	cp.Purge()
}
