package customlog

import (
	"fmt"
	"log"
	"os"

    "github.com/gen2brain/dlgs"
    "github.com/mattn/go-isatty"
)

func debug(msg string, useGUI bool) {
    f, err := os.Create("/tmp/debug.log")
    if err != nil {
        fmt.Println(err)
        return
    }
    _, err = f.WriteString(fmt.Sprintf("%s\n%v\n%s\n",msg,useGUI,os.Args[0]))
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}

// Log is similar to log.Fatal, but with a pause
func Log(msg string) {
    useGUI := true
    if isatty.IsTerminal(os.Stdout.Fd()) {
        useGUI = false
    } else if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
        useGUI = false
    }
    //debug(msg,useGUI)

    if useGUI {
        dlgs.Error("Error", msg)
    } else {
        log.Print(msg)
    }
	os.Exit(1)
}

// Fatalf is similar to log.Fatalf, but with a pause
func Fatalf(format string, v ...interface{}) {
    useGUI := true
    if isatty.IsTerminal(os.Stdout.Fd()) {
        useGUI = false
    } else if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
        useGUI = false
    }
    msg := fmt.Sprintf(format, v...)
    //debug(msg,useGUI)
    if useGUI {
        dlgs.Error("Error", msg)
    } else {
        log.Print(msg)
    }
	os.Exit(1)
}
