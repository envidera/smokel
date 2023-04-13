package linexer

import (
	"strings"
)

// lineFormat removes:
//
//	new lines, tabs, and white space from the beginning and end of the string.
//	right comments on line
func cleanUp(text string) string {
	return removeRightComment(removeSpaces(text))
}

func removeSpaces(text string) string {
	return strings.TrimSpace(text)
}

func removeRightComment(text string) string {
	index := strings.Index(text, "//") + strings.Index(text, "#")
	if index > 1 { // index > 1  = line doesn't start with //, but has // somewhere
		text = removeSpaces(text[:index+1]) // remove // and everything after //
	}

	return text
}
