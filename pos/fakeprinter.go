package pos

// FakePrinter is a dummy printer for debugging.
type FakePrinter struct {
	output *string
}

// PrintText pretends to send some text to the printer.
func (p *FakePrinter) PrintText(text string) {
	*p.output = text
}

// PrintFile pretends to print a file.
func (p *FakePrinter) PrintFile(filename string) {
	*p.output = "Just pretend I printed " + filename
}

// NewFakePrinter returns an instance of FakePrinter which sends its output to the given string.
func NewFakePrinter(output *string) FakePrinter {
	var fakePrinter = FakePrinter{}
	fakePrinter.output = output
	return fakePrinter
}
