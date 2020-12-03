package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gen2brain/dlgs"
	"github.com/jftuga/copypaste"
	"github.com/jftuga/copypaste/customlog"
	"github.com/jftuga/copypaste/queue"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {
	argsVersion := flag.Bool("v", false, "display program version and then exit")
	flag.Parse()

	if *argsVersion {
		fmt.Println(copypaste.Version())
		return
	}

	fileName := ""
	if len(flag.Args()) == 1 {
		fileName = flag.Arg(0)
		if !fileExists(fileName) {
			customlog.Fatalf("File not found: %s", fileName)
		}
	} else {
		if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
			var success bool
			var err error
			fileName, success, err = dlgs.File("SQS Copy Small File", "", false)
			if !success {
				//customlog.Fatalf("Unable to open file dialog")
				os.Exit(0)
			}
			if err != nil {
				customlog.Log(err.Error())
			}
		}
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		customlog.Log(err.Error())
	}
	if len(data) <= 4 {
		customlog.Log("clipboard must have at least 5 bytes of data to send!")
	}

	queueURL := queue.GetQueueURL()
	cp := copypaste.New(queueURL)
	cp.CopySmallFile(filepath.Base(fileName), data)
}
