package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/gen2brain/dlgs"
	"github.com/jftuga/copypaste"
	"github.com/jftuga/copypaste/customlog"
	"github.com/jftuga/copypaste/queue"
)

func dirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func main() {
	argsVersion := flag.Bool("v", false, "display program version and then exit")
	flag.Parse()

	if *argsVersion {
		fmt.Println(copypaste.Version())
		return
	}

	queueURL := queue.GetQueueURL()
	cp := copypaste.New(queueURL)

	destPath := "."
	if len(flag.Args()) == 1 {
		destPath = flag.Arg(0)
		if !dirExists(destPath) {
			customlog.Fatalf("Directory not found: %s", destPath)
		}
	} else {
		if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
			var success bool
			var err error
			destPath, success, err = dlgs.File("SQS Paste Small File", "", true)
			if !success {
				//customlog.Fatalf("Unable to open file dialog")
				os.Exit(0)
			}
			if err != nil {
				customlog.Log(err.Error())
			}
		}
	}

	savedFileName := cp.PasteSmallFile(destPath)
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
		dlgs.Info("SQS Paste Small File", fmt.Sprintf("Saved file: %s", savedFileName))
	} else {
		fmt.Println("Saved file: ", savedFileName)
	}
}
