package coloryaml

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorize(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"map", args{"a: b"}, "\x1b[36ma\x1b[0m:\x1b[32m b\x1b[0m"},
		{"anchor", args{"&a"}, "\x1b[93m&\x1b[0m\x1b[93ma\x1b[0m"},
		{"alias", args{"*a"}, "\x1b[93m*\x1b[0m\x1b[93ma\x1b[0m"},
		{"bool", args{"true"}, "\x1b[95mtrue\x1b[0m"},
		{"string", args{"test"}, "\x1b[32mtest\x1b[0m"},
		{"number", args{"123"}, "\x1b[95m123\x1b[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, tt.want, Colorize(tt.args.s), "Colorize(%v)", tt.args.s)
		})
	}
}
