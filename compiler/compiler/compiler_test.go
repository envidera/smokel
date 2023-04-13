package compiler

/*
import (
	"bytes"
	"strings"
	"testing"
)

func TestCompile(t *testing.T) {

	tests := []struct {
		name string
		kel  string
		css  string
	}{
		{
			name: "first one",
			kel: `
			//comment
			body
				color: red
			nav
				color:blue

			footer
				color:yellow
			`,
			css: `body{color:red;}`,
		},
	}

	css := &bytes.Buffer{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := Compile(strings.NewReader(tt.kel), css)
			if err != nil {
				t.Fatal(err)
			}
			have := oneLine(css.String())
			want := oneLine(tt.css)

			if have != want {
				t.Error("\n have:", have, "\n want:", want)
			}
		})

	}

}

func oneLine(css string) string {
	x := strings.Replace(css, " ", "", -1)
	x = strings.Replace(x, "\t", "", -1)
	x = strings.Replace(x, "\n", "", -1)
	return x
}
*/
