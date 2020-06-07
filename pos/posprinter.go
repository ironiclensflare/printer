package pos

import (
	"fmt"
	"io"
	"os/exec"
)

// Printer represents a physical printer.
type Printer interface {
	PrintText(text string)
	PrintFile(filename string)
}

// ThermalPrinter represents a thermal printer.
type ThermalPrinter struct{}

// PrintText sends some text to the printer.
func (p *ThermalPrinter) PrintText(text string) {
	cmd := exec.Command("lp")
	stdin, _ := cmd.StdinPipe()
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, text)
	}()
	out, _ := cmd.CombinedOutput()
	fmt.Printf("%s\n", out)
}

// PrintFile prints a file.
func (p *ThermalPrinter) PrintFile(filename string) {
	cmd := exec.Command("lp", filename)
	out, _ := cmd.CombinedOutput()
	fmt.Printf("%s\n", out)
}
