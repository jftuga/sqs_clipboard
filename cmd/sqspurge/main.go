//go:generate goversioninfo -icon=sqspurge.ico -platform-specific=true

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
	cp.Purge()
}
