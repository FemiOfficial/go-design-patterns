package adapter

import "log"

type LegacyPrinter interface {
	Print(s string) string
}

type LegacyPrinterInstance struct {}

func (printer *LegacyPrinterInstance) Print(s string) string {
	return "Legacy Printer: " + s
}

type ModernPrinter interface {
	PrintStore() string
}

type PrintterAdapter struct {
	legacyPrinter LegacyPrinter
	message string
}

func (adapter *PrintterAdapter) PrintStore() string {
	if adapter.legacyPrinter != nil {
		return adapter.legacyPrinter.Print(adapter.message)
	} else {
		log.Println(adapter.message)
		return adapter.message
	}

}