package terminal

import (
	"fmt"
	"os"

	"github.com/mdp/qrterminal/v3"
)

// PrintQRCode prints a QR code to the terminal, safely bypassing Windows Sixel bugs
func PrintQRCode(url string) {
	fmt.Println("QR Code:")
	config := qrterminal.Config{
		Level:     qrterminal.L,
		Writer:    os.Stdout,
		BlackChar: qrterminal.BLACK,
		WhiteChar: qrterminal.WHITE,
		QuietZone: 2,
	}
	qrterminal.GenerateWithConfig(url, config)
}
