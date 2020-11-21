/*
copypaste.go
-John Taylor
Nov-19-2020

copy / paste text to / from SQS queue

See also:
https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/sqs-example-receive-message.html
https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html

aws configure set copypaste https://sqs.us-east-1.amazonaws.com/078680625362/CopyPaste2.fifo

*/

package copypaste

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/jftuga/copypaste/customlog"
	"github.com/mtraver/base91"
	"github.com/ulikunitz/xz"
)

// CopyPaste contains the AWS SQS queue url
type CopyPaste struct {
	QueueURL string
	sess     session.Session
	svc      sqs.SQS
}

// compress msg before sending it to SQS
func compress(msg string) bytes.Buffer {
	var buf bytes.Buffer
	w, err := xz.NewWriter(&buf)
	if err != nil {
		customlog.Fatalf("xz.NewWriter error %s", err)
	}
	if _, err := io.WriteString(w, msg); err != nil {
		customlog.Fatalf("WriteString error %s", err)
	}
	if err := w.Close(); err != nil {
		customlog.Fatalf("w.Close error %s", err)
	}

	return buf
}

// compress msg before sending it to SQS
func compressBinary(data []byte) bytes.Buffer {
	var buf bytes.Buffer
	w, err := xz.NewWriter(&buf)
	if err != nil {
		customlog.Fatalf("xz.NewWriter error %s", err)
	}
	if _, err := io.WriteString(w, bytes.NewBuffer(data).String()); err != nil {
		customlog.Fatalf("WriteString error %s", err)
	}
	if err := w.Close(); err != nil {
		customlog.Fatalf("w.Close error %s", err)
	}

	return buf
}

// decompress data sent from SQS
func decompress(buf *bytes.Buffer) *bytes.Buffer {
	r, err := xz.NewReader(buf)
	if err != nil {
		customlog.Fatalf("NewReader error %s", err)
	}

	var uncompressed bytes.Buffer
	if _, err = io.Copy(&uncompressed, r); err != nil {
		customlog.Fatalf("io.Copy error %s", err)
	}
	return &uncompressed
}

// New initializes a new CopyPaste object
func New(queueURL string) *CopyPaste {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := sqs.New(sess)
	return &CopyPaste{
		QueueURL: queueURL,
		sess:     *sess,
		svc:      *svc,
	}
}

// Copy a message to the SQS queue
func (cp CopyPaste) Copy(msg string) {
	xz := compress(msg)
	encoding := "text"
	payload := &msg
	if float32(xz.Len()/len(msg)) < 0.85 {
		b91enc := base91.StdEncoding.EncodeToString(xz.Bytes())
		//fmt.Println(len(msg), xz.Len(), len(b91enc))

		maxSize := 256 * 1024
		if len(b91enc) > maxSize {
			customlog.Fatalf("Data of length %d is too big! Should be less than %d.", len(b91enc), maxSize)
		}

		// compare orig text size to b91enc size
		if len(b91enc) < len(msg) {
			encoding = "xzb91"
			payload = &b91enc
		}
	}

	_, err := cp.svc.SendMessage(&sqs.SendMessageInput{
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Encoding": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(encoding),
			},
		},
		MessageGroupId: &cp.QueueURL,
		MessageBody:    aws.String(*payload),
		QueueUrl:       &cp.QueueURL,
	})
	if err != nil {
		customlog.Log(err.Error())
		return
	}
}

// Paste and remove a message from the SQS queue
func (cp CopyPaste) Paste() string {
	var timeout int64
	timeout = 5

	msgResult, err := cp.svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &cp.QueueURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   &timeout,
	})
	if err != nil {
		customlog.Log(err.Error())
		return ""
	}

	if len(msgResult.Messages) == 0 {
		customlog.Fatalf("There are no more messages in the SQS queue!")
	}

	cp.Delete(*msgResult.Messages[0].ReceiptHandle)

	encoding := (*msgResult.Messages[0]).MessageAttributes["Encoding"].StringValue
	if "xzb91" == *encoding {
		body := *msgResult.Messages[0].Body
		b91dec, err := base91.StdEncoding.DecodeString(body)
		if err != nil {
			customlog.Log(err.Error())
			return ""
		}
		compressedData := bytes.NewBuffer(b91dec)
		un := decompress(compressedData)
		return un.String()
	} else if "text" == *encoding {
		return *msgResult.Messages[0].Body
	}

	customlog.Fatalf("Unknown encoding: %s", *encoding)
	return ""
}

// CopySmallFile send a file to the SQS queue
// it must be smaller than 256 KB after XZ compression and b91 encoding
func (cp CopyPaste) CopySmallFile(fileName string, data []byte) {
	xz := compressBinary(data)
	encoding := "xzb91"

	b91enc := base91.StdEncoding.EncodeToString(xz.Bytes())
	//fmt.Println(len(data), xz.Len(), len(b91enc))

	maxSize := 256 * 1024
	if len(b91enc) > maxSize {
		customlog.Fatalf("Data of length %d is too big! Should be less than %d.", len(b91enc), maxSize)
	}

	_, err := cp.svc.SendMessage(&sqs.SendMessageInput{
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Encoding": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(encoding),
			},
			"Filename": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(fileName),
			},
		},
		MessageGroupId: &cp.QueueURL,
		MessageBody:    aws.String(b91enc),
		QueueUrl:       &cp.QueueURL,
	})
	if err != nil {
		customlog.Log(err.Error())
		return
	}
}

// PasteSmallFile retrieves a file from SQS and saves it to the file system
func (cp CopyPaste) PasteSmallFile() string {
	var timeout int64
	timeout = 5

	msgResult, err := cp.svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &cp.QueueURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   &timeout,
	})
	if err != nil {
		customlog.Log(err.Error())
		return ""
	}

	if len(msgResult.Messages) == 0 {
		customlog.Fatalf("There are no more messages in the SQS queue!")
		return ""
	}

	cp.Delete(*msgResult.Messages[0].ReceiptHandle)

	fileName := (*msgResult.Messages[0]).MessageAttributes["Filename"].StringValue
	encoding := (*msgResult.Messages[0]).MessageAttributes["Encoding"].StringValue
	if "xzb91" == *encoding {
		body := *msgResult.Messages[0].Body
		b91dec, err := base91.StdEncoding.DecodeString(body)
		if err != nil {
			customlog.Log(err.Error())
			return ""
		}
		compressedData := bytes.NewBuffer(b91dec)
		un := decompress(compressedData)
		err = ioutil.WriteFile(*fileName, un.Bytes(), 0700)
		if err != nil {
			customlog.Log(err.Error())
			return ""
		}
	} else {
		customlog.Fatalf("Unknown encoding: %s", *encoding)
		return ""
	}
	return *fileName
}

// Delete a message from the SQS queue
func (cp CopyPaste) Delete(messageHandle string) {
	_, err := cp.svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &cp.QueueURL,
		ReceiptHandle: &messageHandle,
	})
	if err != nil {
		customlog.Log(err.Error())
	}
}

// Purge remove all messages from the SQS queue
func (cp CopyPaste) Purge() {
	var input sqs.PurgeQueueInput
	input.QueueUrl = &cp.QueueURL
	_, err := cp.svc.PurgeQueue(&input)
	if err != nil {
		customlog.Log(err.Error())
	}
}
