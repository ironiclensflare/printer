package queue

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// GetMessages returns the first 10 messages available in the SQS queue.
func GetMessages() []*sqs.Message {
	var maxMessage int64 = 10
	var queueURL = os.Getenv("PRINTER_QUEUE_URL")
	var waitTimeSeconds int64 = 20
	var visibilityTimeout int64 = 5

	session, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	client := sqs.New(session)

	request := sqs.ReceiveMessageInput{
		MaxNumberOfMessages: &maxMessage,
		QueueUrl:            &queueURL,
		WaitTimeSeconds:     &waitTimeSeconds,
		VisibilityTimeout:   &visibilityTimeout,
	}
	output, err := client.ReceiveMessage(&request)
	if err != nil {
		log.Fatal(err)
	}

	return output.Messages
}

// DeleteMessage deletes a message from the queue.
func DeleteMessage(receiptHandle string) {
	var queueURL = os.Getenv("PRINTER_QUEUE_URL")

	session, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	client := sqs.New(session)

	request := sqs.DeleteMessageInput{
		QueueUrl:      &queueURL,
		ReceiptHandle: &receiptHandle,
	}
	output, err := client.DeleteMessage(&request)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(output)
}
