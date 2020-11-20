package customlog

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Log is similar to log.Fatal, but with a pause
func Log(msg string) {
	log.Print(msg)
	time.Sleep(5 * time.Second)
	os.Exit(1)
}

// Fatalf is similar to log.Fatalf, but with a pause
func Fatalf(format string, v ...interface{}) {
	log.Println(fmt.Sprintf(format, v...))
	time.Sleep(5 * time.Second)
	os.Exit(1)
}
