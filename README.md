# sqs_clipboard
Use AWS SQS as a clipboard to copy and paste across different systems and platforms.

**Description**

This set of programs can be used to *copy* and *paste* by using an [AWS SQS FIFO Queue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html) as an intermediary. To minimize the amount of data transferred, the contents are compressed before sending with `sqscopy` and then decompressed with `sqspaste`.  **WARNING:** There can be a small cost when using this program.  Each copy / paste operation uses 3 SQS requests, plus the data transferred associated with *sqspaste*.  See [Amazon SQS pricing](https://aws.amazon.com/sqs/pricing/) for more details.

**Programs**

* sqscopy - send the system clipboard contents to a AWS SQS FIFO queue
* sqspaste - get the SQS queue contents and then put it into the system clipboard
* sqspurhe - remove all entries from the SQS queue

**Setting Environment Variables**

* [How to set environment variables](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html#envvars-set)
* * `SQS_CLIPBOARD_URL` should be set to URL of your SQS `FIFO` Queue

**Icons**

* [Button Upload Icon](https://www.iconarchive.com/show/soft-scraps-icons-by-hopstarter/Button-Upload-icon.html)
* [Button Download Icon](https://www.iconarchive.com/show/soft-scraps-icons-by-hopstarter/Button-Download-icon.html)
* [Email Delete Icon](https://www.iconarchive.com/show/soft-scraps-icons-by-hopstarter/Email-Delete-icon.html)

**Bundling Icons**

* `Cross Platform`
* [a cross platform Go library to place an icon and menu in the notification area](https://github.com/getlantern/systray)
* `Windows`
* [How do you set the application icon in golang?](https://stackoverflow.com/questions/25602600/how-do-you-set-the-application-icon-in-golang)
* * [Tool for embedding .ico & manifest resources in Go programs for Windows](https://github.com/akavel/rsrc)
* `Mac` - **TODO**
* [Packaging a Go application for macOS](https://medium.com/@mattholt/packaging-a-go-application-for-macos-f7084b00f6b5)
* * [Distribute your Go program (or any single binary) as a native macOS application](https://gist.github.com/mholt/11008646c95d787c30806d3f24b2c844)
* * [Go library to create menubar apps- programs that live only in OSX's NSStatusBar](https://github.com/caseymrm/menuet)
