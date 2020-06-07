package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/desertbit/fillpdf"
)

func main() {
	var maxMessage int64 = 10
	var queueURL = os.Getenv("PRINTER_QUEUE_URL")
	var waitTimeSeconds int64 = 20
	var visibilityTimeout int64 = 0

	session, _ := session.NewSession()
	client := sqs.New(session)

	request := sqs.ReceiveMessageInput{
		MaxNumberOfMessages: &maxMessage,
		QueueUrl:            &queueURL,
		WaitTimeSeconds:     &waitTimeSeconds,
		VisibilityTimeout:   &visibilityTimeout,
	}
	output, _ := client.ReceiveMessage(&request)
	parseMessages(output.Messages)
}

func parseMessages(messages []*sqs.Message) {
	for _, message := range messages {
		switch getMessageType(message) {
		case "CITATION":
			fmt.Println("Message is a citation")
			body := strings.Replace(*message.Body, "!AWOO", "", 1)
			parts := strings.SplitN(body, "|", 3)
			name, offence, penalty := parts[0], parts[1], parts[2]
			createCitation(name, offence, penalty)
		case "STICKER":
			fmt.Println("Message is a sticker")
		case "UNKNOWN":
			fmt.Println("Message type is unknown")
		}
	}
}

func getMessageType(message *sqs.Message) string {
	if strings.HasPrefix(*message.Body, "!AWOO") {
		return "CITATION"
	} else if strings.HasPrefix(*message.Body, "!STICKER") {
		return "STICKER"
	} else {
		return "UNKNOWN"
	}
}

func sendTextToPrinter(text string) {
	cmd := exec.Command("lp")
	stdin, _ := cmd.StdinPipe()
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, text)
	}()
	out, _ := cmd.CombinedOutput()
	fmt.Printf("%s\n", out)
}

func sendFileToPrinter(filename string) {
	cmd := exec.Command("lp", "-o page-right=12", filename)
	out, _ := cmd.CombinedOutput()
	fmt.Printf("%s\n", out)
}

func createCitation(name, offence, penalty string) {
	fmt.Printf("Name: %v\nCrime: %v\nPenalty: %v\n", name, offence, penalty)
	form := fillpdf.Form{
		"Name":    name,
		"Offence": offence,
		"Penalty": penalty,
	}
	err := fillpdf.Fill(form, "awoo.pdf", "awoo-filled.pdf", true)
	fmt.Println(err)
	sendFileToPrinter("awoo-filled.pdf")
}
