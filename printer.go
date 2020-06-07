package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/desertbit/fillpdf"
	"github.com/ironiclensflare/printer/pos"
	"github.com/ironiclensflare/printer/queue"
)

var printer pos.Printer

func main() {
	if runtime.GOOS == "darwin" {
		log.Print("OS appears to be Mac - using fake printer")
		printer = &pos.FakePrinter{}
	} else {
		log.Print("OS appears to be Linux - using real printer")
		printer = &pos.ThermalPrinter{}
	}
	messages := queue.GetMessages()
	parseMessages(messages)
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

func createCitation(name, offence, penalty string) {
	fmt.Printf("Name: %v\nCrime: %v\nPenalty: %v\n", name, offence, penalty)
	form := fillpdf.Form{
		"Name":    name,
		"Offence": offence,
		"Penalty": penalty,
	}
	err := fillpdf.Fill(form, "awoo.pdf", "awoo-filled.pdf", true)
	fmt.Println(err)
	printer.PrintFile("awoo-filled.pdf")
}
