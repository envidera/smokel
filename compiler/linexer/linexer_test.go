package linexer

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/anstk/smokel/compiler/data"
	"github.com/anstk/smokel/compiler/messenger"
)

func TestLinexer(t *testing.T) {

	tests := []struct {
		name string
		text string
		want []data.Line
	}{

		{
			name: "comment",
			text: `
// comment
// comment with 6 spaces on right      
	// comment with 1 tab
    // comment with 4 spaces
		// comment with 2 tabs

/*	comment group
	more comment 1

	more comment 2
# comment
`,
			want: []data.Line{
				{Number: 1, Name: "comment", Type: data.LineEmpty, Text: "", Formatted: "", Indent: 0},
				{
					Number:    2,
					Name:      "comment",
					Type:      data.LineComment,
					Text:      "// comment",
					Formatted: "// comment",
					Indent:    0,
				},
				{
					Number:    3,
					Name:      "comment",
					Type:      data.LineComment,
					Text:      "// comment with 6 spaces on right      ",
					Formatted: "// comment with 6 spaces on right",
					Indent:    0,
				},
				{
					Number:    4,
					Name:      "comment",
					Type:      data.LineComment,
					Text:      "	// comment with 1 tab",
					Formatted: "// comment with 1 tab",
					Indent:    1,
				},
				{
					Number:    5,
					Name:      "comment",
					Type:      data.LineComment,
					Text:      "    // comment with 4 spaces",
					Formatted: "// comment with 4 spaces",
					Indent:    1,
				},
				{
					Number:    6,
					Name:      "comment",
					Type:      data.LineComment,
					Text:      "		// comment with 2 tabs",
					Formatted: "// comment with 2 tabs",
					Indent:    2,
				},
				{Number: 7, Name: "comment", Type: data.LineEmpty, Text: "", Formatted: "", Indent: 0},
				{
					Number:    8,
					Name:      "comment",
					Type:      data.LineCommentGroup,
					Text:      "/*	comment group",
					Formatted: "/*	comment group",
					Indent:    0,
				},
				{
					Number:    9,
					Name:      "comment",
					Type:      data.LineComment,
					Text:      "	more comment 1",
					Formatted: "more comment 1",
					Indent:    1,
				},
				{Number: 10, Name: "comment", Type: data.LineEmpty, Text: "", Formatted: "", Indent: 0},
				{
					Number:    11,
					Name:      "comment",
					Type:      data.LineComment,
					Text:      "	more comment 2",
					Formatted: "more comment 2",
					Indent:    1,
				},
				{
					Number:    12,
					Name:      "comment",
					Type:      data.LineComment,
					Text:      "# comment",
					Formatted: "# comment",
					Indent:    0,
				},
			},
		},
		{
			name: "include",
			text: `
include
	test/include/btn.kel
	test/include/main.kel
`,
			want: []data.Line{
				{Number: 1, Name: "include", Type: data.LineEmpty, Text: "", Formatted: "", Indent: 0},
				{
					Number:    2,
					Name:      "include",
					Type:      data.LineEmpty, // after using it, include resets line to Type > empty, Text/Formatted > ""  Indent > 0
					Text:      "",
					Formatted: "",
					Indent:    0,
				},
				{
					Number:    3,
					Name:      "include",
					Type:      data.LineEmpty, // after using it, include resets line to Type > empty, Text/Formatted > ""  Indent > 0
					Text:      "",
					Formatted: "",
					Indent:    0,
				},
				{
					Number:    4,
					Name:      "include",
					Type:      data.LineEmpty, // after using it, include resets line to Type > empty, Text/Formatted > ""  Indent > 0
					Text:      "",
					Formatted: "",
					Indent:    0,
				},
				{
					Number:    1,
					Name:      "test/include/btn.kel",
					Type:      data.LineElement,
					Text:      ".btn",
					Formatted: ".btn",
					Indent:    0,
				},
				{
					Number:    2,
					Name:      "test/include/btn.kel",
					Type:      data.LineProperty,
					Text:      "    padding: 6px 12px",
					Formatted: "padding: 6px 12px",
					Indent:    1,
				},
				{
					Number:    1,
					Name:      "test/include/main.kel",
					Type:      data.LineEmpty, // after using it, include resets line to Type > empty, Text/Formatted > ""  Indent > 0
					Text:      "",
					Formatted: "",
					Indent:    0,
				},
				{
					Number:    2,
					Name:      "test/include/main.kel",
					Type:      data.LineEmpty, // after using it, include resets line to Type > empty, Text/Formatted > ""  Indent > 0
					Text:      "",
					Formatted: "",
					Indent:    0,
				},
				{
					Number:    3,
					Name:      "test/include/main.kel",
					Type:      data.LineElement,
					Text:      "body",
					Formatted: "body",
					Indent:    0,
				},
				{
					Number:    4,
					Name:      "test/include/main.kel",
					Type:      data.LineProperty,
					Text:      "    font: arial",
					Formatted: "font: arial",
					Indent:    1,
				},
				{
					Number:    1,
					Name:      "test/include/comment.kel",
					Type:      data.LineComment,
					Text:      "// just a comment",
					Formatted: "// just a comment",
					Indent:    0,
				},
			},
		},
		{
			name: "var",
			text: `
var
	$color : red
	$color:red
	$color: red
	$color:blue //with comment
`,
			want: []data.Line{
				{Number: 1, Name: "var", Type: data.LineEmpty, Text: "", Formatted: "", Indent: 0},
				{
					Number:    2,
					Name:      "var",
					Type:      data.LineVarGroup,
					Text:      "var",
					Formatted: "var",
					Indent:    0,
				},
				{
					Number:    3,
					Name:      "var",
					Type:      data.LineVar,
					Text:      "	$color : red",
					Formatted: "$color : red",
					Indent:    1,
				},
				{
					Number:    4,
					Name:      "var",
					Type:      data.LineVar,
					Text:      "	$color:red",
					Formatted: "$color:red",
					Indent:    1,
				},
				{
					Number:    5,
					Name:      "var",
					Type:      data.LineVar,
					Text:      "	$color: red",
					Formatted: "$color: red",
					Indent:    1,
				},
				{
					Number:    6,
					Name:      "var",
					Type:      data.LineVar,
					Text:      "	$color:blue //with comment",
					Formatted: "$color:blue",
					Indent:    1,
				},
			},
		},
		{
			name: "block",
			text: `
block
	$blockname
		color: red
		color: blue

block

	$blockname
		color: red
		color : $color
`,
			want: []data.Line{
				{Number: 1, Name: "block", Type: data.LineEmpty, Text: "", Formatted: "", Indent: 0},
				{
					Number:    2,
					Name:      "block",
					Type:      data.LineBlockGroup,
					Text:      "block",
					Formatted: "block",
					Indent:    0,
				},
				{
					Number:    3,
					Name:      "block",
					Type:      data.LineBlockName,
					Text:      "	$blockname",
					Formatted: "$blockname",
					Indent:    1,
				},
				{
					Number:    4,
					Name:      "block",
					Type:      data.LineBlockProperty,
					Text:      "		color: red",
					Formatted: "color: red",
					Indent:    2,
				},
				{
					Number:    5,
					Name:      "block",
					Type:      data.LineBlockProperty,
					Text:      "		color: blue",
					Formatted: "color: blue",
					Indent:    2,
				},
				{Number: 6, Name: "block", Type: data.LineEmpty, Text: "", Formatted: "", Indent: 0},
				{
					Number:    7,
					Name:      "block",
					Type:      data.LineBlockGroup,
					Text:      "block",
					Formatted: "block",
					Indent:    0,
				},
				{Number: 8, Name: "block", Type: data.LineEmpty, Text: "", Formatted: "", Indent: 0},

				{
					Number:    9,
					Name:      "block",
					Type:      data.LineBlockName,
					Text:      "	$blockname",
					Formatted: "$blockname",
					Indent:    1,
				},
				{
					Number:    10,
					Name:      "block",
					Type:      data.LineBlockProperty,
					Text:      "		color: red",
					Formatted: "color: red",
					Indent:    2,
				},
				{
					Number:    11,
					Name:      "block",
					Type:      data.LineBlockProperty,
					Text:      "		color : $color",
					Formatted: "color : $color",
					Indent:    2,
				},
			},
		},
		{
			name: "block use",
			text: `
body
	$blockname  // block
	color: red

@media (max-width: 600px)
	body
		color: red
		$blockname // block
`,
			want: []data.Line{
				{Number: 1, Name: "block use", Type: data.LineEmpty, Text: "", Formatted: "", Indent: 0},
				{
					Number:    2,
					Name:      "block use",
					Type:      data.LineElement,
					Text:      "body",
					Formatted: "body",
					Indent:    0,
				},
				{
					Number:    3,
					Name:      "block use",
					Type:      data.LineBlock,
					Text:      "	$blockname  // block",
					Formatted: "$blockname",
					Indent:    1,
				},
				{
					Number:    4,
					Name:      "block use",
					Type:      data.LineProperty,
					Text:      "	color: red",
					Formatted: "color: red",
					Indent:    1,
				},
				{Number: 5, Name: "block use", Type: data.LineEmpty, Text: "", Formatted: "", Indent: 0},
				{
					Number:    6,
					Name:      "block use",
					Type:      data.LineMedia,
					Text:      "@media (max-width: 600px)",
					Formatted: "@media (max-width: 600px)",
					Indent:    0,
				},
				{
					Number:    7,
					Name:      "block use",
					Type:      data.LineMediaElement,
					Text:      "	body",
					Formatted: "body",
					Indent:    1,
				},
				{
					Number:    8,
					Name:      "block use",
					Type:      data.LineMediaProperty,
					Text:      "		color: red",
					Formatted: "color: red",
					Indent:    2,
				},
				{
					Number:    9,
					Name:      "block use",
					Type:      data.LineBlock,
					Text:      "		$blockname // block",
					Formatted: "$blockname",
					Indent:    2,
				},
			},
		},

		{
			name: "media",
			text: `
@media only screen and (max-width: 600px)
	body
		color: red
`,
			want: []data.Line{
				{Number: 1, Name: "media", Type: data.LineEmpty, Text: "", Formatted: "", Indent: 0},
				{
					Number:    2,
					Name:      "media",
					Type:      data.LineMedia,
					Text:      "@media only screen and (max-width: 600px)",
					Formatted: "@media only screen and (max-width: 600px)",
					Indent:    0,
				},
				{
					Number:    3,
					Name:      "media",
					Type:      data.LineMediaElement,
					Text:      "	body",
					Formatted: "body",
					Indent:    1,
				},
				{
					Number:    4,
					Name:      "media",
					Type:      data.LineMediaProperty,
					Text:      "		color: red",
					Formatted: "color: red",
					Indent:    2,
				},
			},
		},
		{
			name: "line raw",
			text: `
raw
	body{
			color: red;
	}
`,
			want: []data.Line{
				{Number: 1, Name: "line raw", Type: data.LineEmpty, Text: "", Formatted: "", Indent: 0},
				{
					Number:    2,
					Name:      "line raw",
					Type:      data.LineRawGroup,
					Text:      "raw",
					Formatted: "raw",
					Indent:    0,
				},
				{
					Number:    3,
					Name:      "line raw",
					Type:      data.LineRaw,
					Text:      "	body{",
					Formatted: "body{",
					Indent:    1,
				},
				{
					Number:    4,
					Name:      "line raw",
					Type:      data.LineRaw,
					Text:      "			color: red;",
					Formatted: "color: red;",
					Indent:    -1,
				},
				{
					Number:    5,
					Name:      "line raw",
					Type:      data.LineRaw,
					Text:      "	}",
					Formatted: "}",
					Indent:    1,
				},
			},
		},
	}

	ln := New(messenger.New())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			have := ln.lex(tt.name, strings.NewReader(tt.text))

			want := tt.want

			equal, cHave, cWant := equalSlice(have, want)
			if !equal {
				t.Error("\n have:", cHave, "\n want:", cWant)
			}

		})

	}
}

//========================================================================

func TestLinexerErrors(t *testing.T) {

	tests := []struct {
		name string
		text string
		want []data.Line
	}{

		{
			name: "errVarGroupJustVar",
			text: `
var
    body //this is a error
`,
			want: []data.Line{
				{
					Number: 1, Name: "errVarGroupJustVar", Type: data.LineEmpty, Text: "", Formatted: "", Indent: 0,
				},
				{
					Number:    2,
					Name:      "errVarGroupJustVar",
					Type:      data.LineVarGroup,
					Text:      "var",
					Formatted: "var",
					Indent:    0,
				},
				{
					Number:    3,
					Name:      "errVarGroupJustVar",
					Type:      data.NotIdentified,
					Text:      "    body //this is a error",
					Formatted: "body",
					Indent:    1,
				},
			},
		},
		{
			name: "errBlockGroupJustBlockName",
			text: `
block
    body //this is a error
`,
			want: []data.Line{
				{
					Number: 1, Name: "errBlockGroupJustBlockName", Type: data.LineEmpty, Text: "", Formatted: "", Indent: 0,
				},
				{
					Number:    2,
					Name:      "errBlockGroupJustBlockName",
					Type:      data.LineBlockGroup,
					Text:      "block",
					Formatted: "block",
					Indent:    0,
				},
				{
					Number:    3,
					Name:      "errBlockGroupJustBlockName",
					Type:      data.NotIdentified,
					Text:      "    body //this is a error",
					Formatted: "body",
					Indent:    1,
				},
			},
		},
		{
			name: "errBlockGroupJustProperty",
			text: `
block
    $blockName
		$var :red //this is a error
		string{$interpolation} //this is a error
`,
			want: []data.Line{
				{
					Number: 1, Name: "errBlockGroupJustProperty", Type: data.LineEmpty, Text: "", Formatted: "", Indent: 0,
				},
				{
					Number:    2,
					Name:      "errBlockGroupJustProperty",
					Type:      data.LineBlockGroup,
					Text:      "block",
					Formatted: "block",
					Indent:    0,
				},
				{
					Number:    3,
					Name:      "errBlockGroupJustProperty",
					Type:      data.LineBlockName,
					Text:      "    $blockName",
					Formatted: "$blockName",
					Indent:    1,
				},
				{
					Number:    4,
					Name:      "errBlockGroupJustProperty",
					Type:      data.NotIdentified,
					Text:      "    body //this is a error",
					Formatted: "body",
					Indent:    2,
				},
			},
		},
	}

	ln := New(messenger.New())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			have := ln.lex(tt.name, strings.NewReader(tt.text))

			want := tt.want

			equal, cHave, cWant := equalSlice(have, want)
			if !equal {
				t.Error("\n have:", cHave, "\n want:", cWant)
			}

		})

	}

}

//========================================================================

func equalSlice(a, b []data.Line) (equal bool, have, want data.Line) {

	if len(a) != len(b) {
		log.Fatal("DEVELOPER: equalSlice(a, b []data.Line) > Different array size ", len(a), " | ", len(b))
	}

	for i, ai := range a {
		if !equalLine(ai, b[i]) {
			return false, ai, b[i]
		}
	}

	return true, data.Line{}, data.Line{}
}

func equalLine(a, b data.Line) bool {
	return reflect.DeepEqual(a, b)
}
