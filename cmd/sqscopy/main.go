//go:generate goversioninfo -icon=sqscopy.ico -platform-specific=true

package main

import (
	"flag"
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/jftuga/copypaste"
	"github.com/jftuga/copypaste/customlog"
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
	data, err := clipboard.ReadAll()
	if err != nil {
		customlog.Log(err.Error())
	}
	if len(data) <= 4 {
		customlog.Log("clipboard must have at least 5 bytes of data to send!")
	}
	cp.Copy(data)
}
