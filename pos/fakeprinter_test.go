package pos

import (
	"os"
	"testing"
)

var printerOutput string
var fakePrinter FakePrinter

func TestMain(m *testing.M) {
	fakePrinter = NewFakePrinter(&printerOutput)
	os.Exit(m.Run())
}

func TestPrintTextDoesNotError(t *testing.T) {
	fakePrinter.PrintText("Blah")
	if printerOutput == "Blah" {
		return
	}
	t.Errorf("Received %s", printerOutput)
}

func TestPrintFileDoesNotError(t *testing.T) {
	fakePrinter.PrintFile("blah.png")
	if printerOutput == "Just pretend I printed blah.png" {
		return
	}
	t.Errorf("Received %s", printerOutput)
}
