package main

import (
	"github.com/atotto/clipboard"
	"github.com/jftuga/copypaste"
	"github.com/jftuga/copypaste/customlog"
	"github.com/jftuga/copypaste/queue"
)

func main() {
	queueURL := queue.GetQueueURL()
	cp := copypaste.New(queueURL)
	data, err := clipboard.ReadAll()
	if err != nil {
		customlog.Log(err.Error())
	}
	if len(data) <= 4 {
		customlog.Log("clipboard must have at least 5 bytes of data to send!")
	}
	cp.Copy(data)
}
