package queue

import (
    "path/filepath"
	"os"

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

    config_ini := filepath.Join(os.Getenv("HOME"),".aws","sqs_clipboard")
    cfg, err := ini.Load(config_ini)
    if err != nil {
        customlog.Fatalf("Fail to read file: [%s]: %v", config_ini, err)
        return ""
    }

    queueURL = cfg.Section("default").Key("SQS_CLIPBOARD_URL").String()
	if len(queueURL) < 8 {
        customlog.Fatalf("\n\nThe 'SQS_CLIPBOARD_URL' environment variable is not set correctly.\n\nOr $HOME/.aws/sqs_clipbord file is not properly configured.\n\nSee: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html#envvars-set\n")
        return ""
    }
    return queueURL
}

