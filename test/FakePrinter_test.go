package test

import (
	"testing"

	"github.com/ironiclensflare/printer/pos"
)

func TestPrintTextDoesNotError(t *testing.T) {
	var printerOutput string
	printer := pos.NewFakePrinter(&printerOutput)
	printer.PrintText("Blah")
	if printerOutput == "Blah" {
		return
	}
	t.Errorf("Received %s", printerOutput)
}

func TestPrintFileDoesNotError(t *testing.T) {
	var printerOutput string
	printer := pos.NewFakePrinter(&printerOutput)
	printer.PrintFile("blah.png")
	if printerOutput == "Just pretend I printed blah.png" {
		return
	}
	t.Errorf("Received %s", printerOutput)
}
