package stdout

import (
	"fmt"
	"sync"

	"github.com/fatih/color"
)

type ColorPrinterInterface interface {
	Println(colorAttr color.Attribute, format string, attr ...interface{})
}

var colorPrinterInterface ColorPrinterInterface

type printer struct{}

var once sync.Once

func GetPrinter() ColorPrinterInterface {
	once.Do(func() {
		colorPrinterInterface = &printer{}
	})

	return colorPrinterInterface
}

func (p *printer) Println(colorAttr color.Attribute, format string, attr ...interface{}) {
	_, _ = color.New(colorAttr).Println(fmt.Sprintf(format, attr...))
}

var ColorPrinter = GetPrinter()
