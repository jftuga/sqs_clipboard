package main

import (
	"flag"
	"fmt"

	"github.com/jftuga/copypaste"
	"github.com/jftuga/copypaste/queue"
)

func main() {
	argsVersion := flag.Bool("v", false, "display program version and then exit")
	flag.Parse()

	if *argsVersion {
		fmt.Println(copypaste.Version())
		return
	}

	queueURL := queue.GetQueueURL()
	cp := copypaste.New(queueURL)
	savedFileName := cp.PasteSmallFile()
	fmt.Println("Saved file: ", savedFileName)
}
