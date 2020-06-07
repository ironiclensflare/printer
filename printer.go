package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	var maxMessage int64 = 10
	var queueURL = os.Getenv("PRINTER_QUEUE_URL")
	var waitTimeSeconds int64 = 20
	var visibilityTimeout int64 = 0

	session, err := session.NewSession()
	client := sqs.New(session)

	request := sqs.ReceiveMessageInput{
		MaxNumberOfMessages: &maxMessage,
		QueueUrl:            &queueURL,
		WaitTimeSeconds:     &waitTimeSeconds,
		VisibilityTimeout:   &visibilityTimeout,
	}
	output, err := client.ReceiveMessage(&request)

	fmt.Println(output)
	fmt.Println(err)
}
