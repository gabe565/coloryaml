package coloryaml

import (
	"io"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/goccy/go-yaml/printer"
	"github.com/mattn/go-isatty"
)

func shouldColor(w io.Writer) bool {
	if os.Getenv("NO_COLOR") != "" || os.Getenv("TERM") == "dumb" {
		return false
	}
	if f, ok := w.(*os.File); ok {
		return isatty.IsTerminal(f.Fd()) || isatty.IsCygwinTerminal(f.Fd())
	}
	return false
}

const escape = "\x1b"

func format(attr color.Attribute) string {
	return escape + "[" + strconv.Itoa(int(attr)) + "m"
}

func Printer() *printer.Printer {
	return &printer.Printer{
		MapKey: func() *printer.Property {
			return &printer.Property{
				Prefix: format(color.FgCyan),
				Suffix: format(color.Reset),
			}
		},
		Anchor: func() *printer.Property {
			return &printer.Property{
				Prefix: format(color.FgHiYellow),
				Suffix: format(color.Reset),
			}
		},
		Alias: func() *printer.Property {
			return &printer.Property{
				Prefix: format(color.FgHiYellow),
				Suffix: format(color.Reset),
			}
		},
		Bool: func() *printer.Property {
			return &printer.Property{
				Prefix: format(color.FgHiMagenta),
				Suffix: format(color.Reset),
			}
		},
		String: func() *printer.Property {
			return &printer.Property{
				Prefix: format(color.FgGreen),
				Suffix: format(color.Reset),
			}
		},
		Number: func() *printer.Property {
			return &printer.Property{
				Prefix: format(color.FgHiMagenta),
				Suffix: format(color.Reset),
			}
		},
	}
}
