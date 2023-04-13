package messenger

import (
	"testing"
)

func TestFormatExample(t *testing.T) {

	tests := []struct {
		name       string
		lineNumber int
		text       string
		want       string
	}{

		{
			name:       "t1",
			lineNumber: 250,
			text: `
-1|var
0|    $color: red //variable`,

			want: `
249|var
250|    $color: red //variable`,
		},
		{
			name:       "t2",
			lineNumber: 51,
			text: `
-2| //comment
-1|var
0|    $color: red //variable
1|    $color2: green`,

			want: `
49| //comment
50|var
51|    $color: red //variable
52|    $color2: green`,
		},
		{
			name:       "t3",
			lineNumber: 1,
			text: `
0|var
1|    $color: red //variable`,

			want: `
1|var
2|    $color: red //variable`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			have := formatExample(tt.lineNumber, tt.text)
			want := tt.want

			//fmt.Println(len(have), len(want))

			if have != want {
				t.Error("\n have:", have, "\n want:", want)
			}

		})

	}
}
