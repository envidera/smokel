package linexer

import (
	"testing"
)

func TestIndent(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{
			name: "1 tab",
			text: "\t",
			want: 1,
		},
		{
			name: "2 tabs",
			text: "\t\t",
			want: 2,
		},
		{
			// 3 tabs is a smokel sintax error, so return -1 and a error
			// the error message is ignored in this test
			name: "3 tabs",
			text: "\t\t\t",
			want: -1,
		},
		{
			name: "1 space",
			text: " ",
			want: 0,
		},
		{
			name: "2 spaces",
			text: "  ",
			want: 0,
		},
		{
			name: "3 spaces",
			text: "   ",
			want: 1,
		},
		{
			name: "4 spaces",
			text: "    ",
			want: 1,
		},
		{
			name: "5 spaces",
			text: "     ",
			want: 1,
		},
		{
			name: "6 spaces",
			text: "      ",
			want: 2,
		},
		{
			name: "7 spaces",
			text: "       ",
			want: 2,
		},
		{
			name: "8 spaces",
			text: "        ",
			want: 2,
		},
		{
			// 9 spaces  is a smokel sintax error, so return -1 and a error
			// the error message is ignored in this test
			name: "9 spaces",
			text: "         ",
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			have := indent(tt.text)
			want := tt.want

			if have != want {
				t.Error("\n have:", have, "\n want:", want)
			}

		})

	}

}

func TestCountLeftChar(t *testing.T) {
	tests := []struct {
		name string
		text string
		char string
		want int
	}{
		{
			name: "0 space",
			text: "zzz",
			char: " ", //space
			want: 0,
		},
		{
			name: "1 space",
			text: " zzz",
			char: " ", //space
			want: 1,
		},
		{
			name: "2 spaces",
			text: "  zzz",
			char: " ", //space
			want: 2,
		},
		{
			name: "3 spaces",
			text: "   zzz",
			char: " ", //space
			want: 3,
		},
		{
			name: "1 tab",
			text: "	zzz",
			char: "\t", //space
			want: 1,
		},
		{
			name: "2 tabs",
			text: "		zzz",
			char: "\t", //space
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			have := countLeftChar(tt.text, tt.char)
			want := tt.want

			if have != want {
				t.Error("\n have:", have, "\n want:", want)
			}

		})

	}

}
