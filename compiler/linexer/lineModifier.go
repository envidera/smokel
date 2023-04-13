package linexer

import "github.com/anstk/smokel/compiler/data"

type lnm struct{}

// Empty is used to clear the line
// Is used by linexer to empty include lines already included
func (ln *lnm) Empty(line *data.Line) {
	line.Type = data.LineEmpty
	line.Text = ""
	line.Formatted = ""
	line.Indent = 0
}
