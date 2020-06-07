package pos

import "fmt"

// FakePrinter is a dummy printer for debugging.
type FakePrinter struct{}

// PrintText pretends to send some text to the printer.
func (p *FakePrinter) PrintText(text string) {
	fmt.Println("--------")
	fmt.Printf("BEEP BOOP BEEP I'M A FAKE PRINTER\n%s\n", text)
	fmt.Println("--------")
}

// PrintFile pretends to print a file.
func (p *FakePrinter) PrintFile(filename string) {
	fmt.Println("--------")
	fmt.Printf("BEEP BOOP BEEP I'M A FAKE PRINTER\n\nJust pretend I printed %s\n", filename)
	fmt.Println("--------")
}
