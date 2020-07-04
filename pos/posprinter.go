package pos

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/ironiclensflare/printer/telegram"
)

// Printer represents a physical printer.
type Printer interface {
	PrintText(text string)
	PrintFile(filename string)
	PrintSticker(stickerID string)
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

func (p *ThermalPrinter) PrintSticker(stickerID string) {
	sticker := telegram.GetSticker()
	filename, _ := sticker.Get(stickerID)
	p.PrintFile(filename)
}
