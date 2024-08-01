package coloryaml

import (
	"bytes"
	"io"
	"os"
	"strings"

	"github.com/goccy/go-yaml/lexer"
	"gopkg.in/yaml.v3"
)

func Colorize(s string) string {
	// https://github.com/mikefarah/yq/blob/v4.43.1/pkg/yqlib/color_print.go
	tokens := lexer.Tokenize(s)
	return Printer().PrintTokens(tokens)
}

func Sprintln(obj any) (string, error) {
	b, err := yaml.Marshal(obj)
	if err != nil {
		return "", err
	}
	s := Colorize(string(b))
	if !strings.HasSuffix(s, "\n") {
		s += "\n"
	}
	return s, nil
}

func Fprintln(w io.Writer, obj any) error {
	if !shouldColor(w) {
		b, err := yaml.Marshal(obj)
		if err != nil {
			return err
		}
		if !bytes.HasSuffix(b, []byte{'\n'}) {
			b = append(b, '\n')
		}
		_, err = w.Write(b)
		return err
	}
	s, err := Sprintln(obj)
	if err != nil {
		return err
	}
	_, err = io.WriteString(w, s)
	return err
}

func Println(obj any) error {
	return Fprintln(os.Stdout, obj)
}
