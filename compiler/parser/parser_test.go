package parser

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/anstk/smokel/compiler/data"
	"github.com/anstk/smokel/compiler/linexer"
	"github.com/anstk/smokel/compiler/messenger"
)

func TestParse(t *testing.T) {

	tests := []struct {
		name string
		text string
		want []data.ParserLine
	}{
		{
			name: "var",
			text: `
var
	$color : red
	$color:red
	$color: red
	$color:blue //with comment
	$final : $unknow
	`,
			want: []data.ParserLine{
				//var ------------------
				{
					Name:   "var",
					Number: 3,
					Type:   int(data.LineVar),
					Parsed: "{{$color := red}}",
				},
				{
					Name:   "var",
					Number: 4,
					Type:   int(data.LineVar),
					Parsed: "{{$color := red}}",
				},
				{
					Name:   "var",
					Number: 5,
					Type:   int(data.LineVar),
					Parsed: "{{$color := red}}",
				},
				{
					Name:   "var",
					Number: 6,
					Type:   int(data.LineVar),
					Parsed: "{{$color := blue}}",
				},
				{
					Name:   "var",
					Number: 7,
					Type:   int(data.LineVar),
					Parsed: "{{$final := $color}}",
				},
			},
		},
		{
			name: "var",
			text: `
/*	start here
    more comment
		
var
	$color: red //uia
	$font-size: 14px

body	
	color: red
	font: arial
	border-radius: 3px
	// comment
	// another comment


.btn
.btn2
	color: blue
	border: 1px solid gray

h1 
h2, h3,
	color: yellow
	`,
			want: []data.ParserLine{
				//var ------------------
				{
					Name:   "var",
					Number: 6,
					Type:   int(data.LineVar),
					Parsed: "{{$color := red}}",
				},
				{
					Name:   "var",
					Number: 7,
					Type:   int(data.LineVar),
					Parsed: "{{$font-size := 14px}}",
				},
				{
					Name:   "var",
					Number: 9,
					Type:   int(data.LineElement),
					Parsed: "body" + charSpace + "{" + charNewLine,
				},
				{
					Name:   "element",
					Number: 6,
					Type:   int(data.LineVar),
					Parsed: "{{$color := red}}",
				},
				{
					Name:   "element",
					Number: 6,
					Type:   int(data.LineVar),
					Parsed: "{{$color := red}}",
				},
				{
					Name:   "element",
					Number: 6,
					Type:   int(data.LineVar),
					Parsed: "{{$color := red}}",
				},
				{
					Name:   "element",
					Number: 6,
					Type:   int(data.LineVar),
					Parsed: "{{$color := red}}",
				},
				{
					Name:   "element",
					Number: 6,
					Type:   int(data.LineVar),
					Parsed: "{{$color := red}}",
				},
				{
					Name:   "element",
					Number: 6,
					Type:   int(data.LineVar),
					Parsed: "{{$color := red}}",
				},
				{
					Name:   "element",
					Number: 6,
					Type:   int(data.LineVar),
					Parsed: "{{$color := red}}",
				},
				{
					Name:   "element",
					Number: 6,
					Type:   int(data.LineVar),
					Parsed: "{{$color := red}}",
				},
				{
					Name:   "element",
					Number: 6,
					Type:   int(data.LineVar),
					Parsed: "{{$color := red}}",
				},
				{
					Name:   "element",
					Number: 6,
					Type:   int(data.LineVar),
					Parsed: "{{$color := red}}",
				},
				{
					Name:   "element",
					Number: 6,
					Type:   int(data.LineVar),
					Parsed: "{{$color := red}}",
				},
				{
					Name:   "element",
					Number: 6,
					Type:   int(data.LineVar),
					Parsed: "{{$color := red}}",
				},
				{
					Name:   "element",
					Number: 6,
					Type:   int(data.LineVar),
					Parsed: "{{$color := red}}",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ms := messenger.New()

			lx := linexer.New(ms)
			lines := lx.Lex(tt.name, strings.NewReader(tt.text))

			ps := &parser{}
			have := ps.parse(lines)
			want := tt.want

			eq, eHave, eWant := equalSlice(have, want)

			if !eq {
				t.Error("\n have:", eHave, "\n want:", eWant)
			}
		})

	}

}

func equalSlice(a, b []data.ParserLine) (equal bool, have, want data.ParserLine) {

	if len(a) != len(b) {
		log.Fatal("DEVELOPER: equalSlice(a, b []data.ParserLine) > Different array size ", len(a), " | ", len(b))
	}

	for i, ai := range a {
		if !compare(ai, b[i]) {
			return false, ai, b[i]
		}
	}

	return true, data.ParserLine{}, data.ParserLine{}
}

func compare(a, b data.ParserLine) bool {
	return reflect.DeepEqual(a, b)
}
