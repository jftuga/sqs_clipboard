# sqs_clipboard
Use [AWS SQS](https://aws.amazon.com/sqs/) as a clipboard to copy and paste across different systems and platforms. Clipboard contents are encrypted in transit and at rest.

Binaries for Windows, MacOS, and Linux can be found on the [Releases Page](https://github.com/jftuga/sqs_clipboard/releases).
___

**Description**

This set of programs can be used to *copy* and *paste* clipboard text by using an [AWS SQS FIFO Queue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html) as an intermediary. To minimize the amount of data transferred, the contents are compressed with the `XZ` algorithm before sending to the SQS queue via `sqscopy` and then decompressed upon arrival via `sqspaste`.  A maximum of `256 KB` of compressed *(and then encoded)* data can be sent to the queue.

**NOTE:** There can be a small AWS cost when using this program.  Each copy / paste operation uses 3 SQS requests, plus the data transferred associated with `sqspaste` and `sqspastesmallfile`.  See [Amazon SQS pricing](https://aws.amazon.com/sqs/pricing/) for more details.

**Programs**

* `sqscopy` - send the system clipboard contents to a user-defined AWS SQS FIFO queue
* `sqspaste` - get the queue contents and then place it onto the system clipboard
* `sqspurge` - remove all entries from the queue
* `sqscopysmallfile` - copy a small file to the queue with file name given on command line
* * after XZ compression and base 91 encoding, the resulting file size must be less than 256 KB
* `sqspastesmallfile` - retrieve a file from the queue
* * **will overwrite** an existing file with the same name
* * file name is stored in the queue along with the file
* * no command line arguments needed

**AWS Queue Creation**

* Open the [Amazon SQS Console](https://console.aws.amazon.com/sqs/v2/home)
* Make sure to create a `fifo` queue instead of a `standard` queue
* * The name of your queue should end in `.fifo`
* Set the `Receive message wait time` aka *long polling* to at least `12` seconds
* **Note:** Do **not** check: `Content-based deduplication`, otherwise you will **not** be able to copy the exact same contents with a 5 minute interval.
* Send / Receive Access: `Only the queue owner`
* Encryption: Optional, but recommended
* * Data key reuse period: `1 hour`
* * A shorter time period provides better security, but results in more calls to AWS KMS, which might incur charges after Free Tier.

**AWS Queue Creation - Automation**

* Please review these scripts before running them and make any adjustments as needed:
* * [Windows](create_stack.ps1)
* * [Linux, MacOS](create_stack.sh)
* Both of these scripts use [cloud_formation_template.json](cloud_formation_template.json)

**AWS IAM Permissions**

* Make sure to change the `Resource` value listed below.

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "VisualEditor0",
            "Effect": "Allow",
            "Action": [
                "sqs:DeleteMessage",
                "sqs:GetQueueUrl",
                "sqs:ChangeMessageVisibility",
                "sqs:PurgeQueue",
                "sqs:ReceiveMessage",
                "sqs:DeleteQueue",
                "sqs:SendMessage",
                "sqs:GetQueueAttributes",
                "sqs:ListQueueTags",
                "sqs:CreateQueue" 
            ],
            "Resource": "arn:aws:sqs:region:account-id:QueueName.fifo"
        }
    ]
}
```

**Setting Environment Variables**

* The `SQS_CLIPBOARD_URL` environment variable should be set to URL of your SQS `FIFO` Queue
* * This URL can be found on the AWS SQS Dashboard for the queue that you have created
* [How to set environment variables](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html#envvars-set)

**Setting Configuration File**
* For Windows and MacOS, you might want to set the `SQS_CLIPBOARD_URL` value in a configuration file instead of using environment variable.
* This will allow you to launch the programs from the *Windows Taskbar* or *MacOS Dock*.
* To do this you will need to create this file:
* * Windows: `%HOMEPATH%\.aws\sqs_clipboard`
* * MacOS: `${HOME}/.aws/sqs_clipboard`

File format:

```ini
[default]
SQS_CLIPBOARD_URL=https://sqs.region.amazonaws.com/account-id/queuename.fifo
```

___

**Compilation**

* Install [GoReleaser](https://goreleaser.com/)
* Run: `goreleaser build --rm-dist --snapshot`
* There should now be 5 resulting binaries found in the `dist` directory:
* * sqscopy
* * sqspaste
* * sqspurge
* * sqscopysmallfile
* * sqspastesmallfile

You can also create a *Mac App* by running [mac_app_builder.sh](cmd/mac_app_builder.sh)

___

**Windows Icons Used**

* [Button Upload Icon](https://www.iconarchive.com/show/soft-scraps-icons-by-hopstarter/Button-Upload-icon.html)
* [Button Download Icon](https://www.iconarchive.com/show/soft-scraps-icons-by-hopstarter/Button-Download-icon.html)
* [Email Delete Icon](https://www.iconarchive.com/show/soft-scraps-icons-by-hopstarter/Email-Delete-icon.html)

**Bundling Icons**

* `Cross Platform`
* [a cross platform Go library to place an icon and menu in the notification area](https://github.com/getlantern/systray)
* `Windows`
* [How do you set the application icon in golang?](https://stackoverflow.com/questions/25602600/how-do-you-set-the-application-icon-in-golang)
* * [goversioninfo](https://github.com/josephspurrier/goversioninfo)
* `Mac`
* [Packaging a Go application for macOS](https://medium.com/@mattholt/packaging-a-go-application-for-macos-f7084b00f6b5)
* * [Distribute your Go program (or any single binary) as a native macOS application](https://gist.github.com/mholt/11008646c95d787c30806d3f24b2c844)
* * [Go library to create menubar apps- programs that live only in OSX's NSStatusBar](https://github.com/caseymrm/menuet)

**AWS Resources**
* https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html
* https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-cli-creating-stack.html
* https://docs.aws.amazon.com/cli/latest/reference/cloudformation/describe-stacks.html

