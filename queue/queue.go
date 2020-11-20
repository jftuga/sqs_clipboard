package queue

import (
	"os"

	"github.com/jftuga/copypaste/customlog"
)

// GetQueueURL retrieves the URL from the OS environemnt
func GetQueueURL() string {
	queueURL := os.Getenv("SQS_CLIPBOARD_URL")
	if len(queueURL) < 8 {
		customlog.Fatalf("The 'SQS_CLIPBOARD_URL' environment variable is not set correctly. See: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html#envvars-set\n")
	}
	return queueURL
}
