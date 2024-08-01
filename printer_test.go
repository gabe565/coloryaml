package coloryaml

import (
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func Test_format(t *testing.T) {
	t.Parallel()
	type args struct {
		attr color.Attribute
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"reset", args{color.Reset}, "\x1b[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, tt.want, format(tt.args.attr), "format(%v)", tt.args.attr)
		})
	}
}
