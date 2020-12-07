package queue

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/jftuga/copypaste/customlog"
	"gopkg.in/ini.v1"
)

// GetQueueURL retrieves the URL from the OS environemnt
// or from ini config file: $HOME/.aws/sqs_clipboard
func GetQueueURL() string {
	queueURL := os.Getenv("SQS_CLIPBOARD_URL")
	if len(queueURL) > 8 {
		return queueURL
	}

	var configIni string
	if runtime.GOOS == "windows" {
		configIni = filepath.Join(os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"), ".aws", "sqs_clipboard")
	} else {
		configIni = filepath.Join(os.Getenv("HOME"), ".aws", "sqs_clipboard")
	}
	cfg, err := ini.Load(configIni)
	if err != nil {
		customlog.Fatalf("Fail to read file: [%s]: %v", configIni, err)
		return ""
	}

	queueURL = cfg.Section("default").Key("SQS_CLIPBOARD_URL").String()
	if len(queueURL) < 8 {
		var location = "$HOME"
		if runtime.GOOS == "windows" {
			location = "%HOMEDRIVE%//%HOMEPATH%"
		}
		customlog.Fatalf("\n\nThe 'SQS_CLIPBOARD_URL' environment variable is not set correctly.\n\nOr %s/.aws/sqs_clipboard file is not properly configured.\n\nSee: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html#envvars-set\n", location)
		return ""
	}
	return queueURL
}
