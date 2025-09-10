package adapter

import (
	"strings"
	"testing"
)

func TestAdapter(t *testing.T) {
	legacyPrinter := &LegacyPrinterInstance{}
	modernPrinter := &PrintterAdapter{ legacyPrinter: legacyPrinter, message: "test message"}

	if (!strings.Contains(modernPrinter.PrintStore(), "Legacy Printer")) {
		t.Errorf("Modern printer is not printing the message correctly")
	}
}

func TestNilAdapter(t *testing.T) {
	modernPrinter := &PrintterAdapter{ legacyPrinter: nil, message: "test message"}

	if (strings.Contains(modernPrinter.PrintStore(), "Legacy Printer")) {
		t.Errorf("Modern printer is not printing the message correctly")
	}
}