package linexer

import (
	"strings"
)

// ident can be:
// level 0 = no tabs, no spaces
// level 1 = 1 tab = 3 or 4 spaces
// level 2 = 2 tabs = 6 or 8 spaces

func indent(text string) int {

	spaces := countLeftChar(text, " ")
	tabs := countLeftChar(text, "\t")

	if tabs >= 3 || // 3 or more tabs
		spaces > 8 { // more than 8 spaces
		return -1
	}

	if tabs == 2 || // 2 tabs
		spaces >= 6 && spaces <= 8 { // 6, 7 and 8 spaces
		return 2
	}

	if tabs == 1 || // 1 tab
		spaces >= 3 && spaces <= 4 || // 3 and 4 spaces
		spaces == 5 { // 5 spaces
		return 1
	}

	return 0
}

// countLeftChar counts how many chars have in beginning of a text
func countLeftChar(text, char string) int {
	return len(text) - len(strings.TrimLeft(text, char))
}
