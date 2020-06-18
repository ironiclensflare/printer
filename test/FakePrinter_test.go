package test

import (
	"testing"

	"github.com/ironiclensflare/printer/pos"
)

func TestPrintTextDoesNotError(t *testing.T) {
	printer := pos.FakePrinter{}
	printer.PrintText("Blah")
}

func TestPrintFileDoesNotError(t *testing.T) {
	printer := pos.FakePrinter{}
	printer.PrintFile("blah.png")
}
