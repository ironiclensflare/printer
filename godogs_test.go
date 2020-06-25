package main

import (
	"errors"

	"github.com/cucumber/godog"
	"github.com/ironiclensflare/printer/pos"
)

var fakePrinter pos.FakePrinter
var fakePrinterOutput string = ""

func iHaveAnInstanceOfTheFakePrinter() error {
	fakePrinter = pos.NewFakePrinter(&fakePrinterOutput)
	return nil
}

func iPrint(text string) error {
	fakePrinter.PrintText(text)
	return nil
}

func iPrintTheFile(filename string) error {
	fakePrinter.PrintFile(filename)
	return nil
}

func iShouldReceiveASimulatedPrintoutContaining(text string) error {
	if fakePrinterOutput == text {
		return nil
	}
	return errors.New("Output was " + fakePrinterOutput + " but expected " + text)
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I have an instance of the fake printer$`, iHaveAnInstanceOfTheFakePrinter)
	s.Step(`^I print "([^"]*)"$`, iPrint)
	s.Step(`^I print the file "([^"]*)"$`, iPrintTheFile)
	s.Step(`^I should receive a simulated printout containing "([^"]*)"$`, iShouldReceiveASimulatedPrintoutContaining)
}
