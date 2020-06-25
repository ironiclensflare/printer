package main

import (
	"github.com/cucumber/godog"
	"github.com/ironiclensflare/printer/pos"
)

var fakePrinter pos.FakePrinter

func iHaveAnInstanceOfTheFakePrinter() error {
	fakePrinter = pos.FakePrinter{}
	return nil
}

func iPrint(text string) error {
	return godog.ErrPending
}

func iShouldReceiveASimulatedPrintoutContaining(text string) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I have an instance of the fake printer$`, iHaveAnInstanceOfTheFakePrinter)
	s.Step(`^I print "([^"]*)"$`, iPrint)
	s.Step(`^I should receive a simulated printout containing "([^"]*)"$`, iShouldReceiveASimulatedPrintoutContaining)
}
