package adapter

import "log"

type DefaultDisplayer interface {
	Print(s string) string
}

type Displayer struct {}

func (displayer *Displayer) Print(str string) string {
	result := "result: " + str 
	log.Printf("DEFAULT| %s", str)
	return result
}


type CustomPrint interface {
	PrintCustom(str string) string
}

type CustomDisplayer struct {
	defaultDisplayer DefaultDisplayer
	message string
}

func (customDisplayer *CustomDisplayer) CustomPrint() string {
	if (customDisplayer.defaultDisplayer == nil) {
		log.Print(customDisplayer.message)
		return customDisplayer.message
	}

	result := customDisplayer.defaultDisplayer.Print(customDisplayer.message)
	return result
}
