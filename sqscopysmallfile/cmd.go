package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

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
	if len(os.Args) != 2 {
		customlog.Log("Give file name on command line.")
	}
	fileName := os.Args[1]
	if !fileExists(fileName) {
		customlog.Fatalf("File not found: %s", fileName)
	}

	data, err := ioutil.ReadFile(os.Args[1])
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
