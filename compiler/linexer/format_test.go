package linexer

import (
	"testing"
)

func TestRemoveRightComment(t *testing.T) {

	tests := []struct {
		name string
		text string
		want string
	}{
		{
			name: "1",
			text: "something // comment",
			want: "something",
		},
		{
			name: "2",
			text: "something //comment",
			want: "something",
		},

		{
			name: "3",
			text: "something//comment",
			want: "something",
		},
		{
			name: "4",
			text: "something#comment",
			want: "something",
		},
		{
			name: "5",
			text: "something		#  	comment",
			want: "something",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			have := removeRightComment(tt.text)
			want := tt.want

			if have != want {
				t.Error("\n have:", have, "\n want:", want)
			}

		})

	}

}
