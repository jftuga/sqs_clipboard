package main

import (
	"fmt"

	"github.com/jftuga/copypaste"
	"github.com/jftuga/copypaste/queue"
)

func main() {
	queueURL := queue.GetQueueURL()
	cp := copypaste.New(queueURL)
	savedFileName := cp.PasteSmallFile()
	fmt.Println("Saved file: ", savedFileName)
}
