package linexer

import (
	"io"

	"github.com/anstk/smokel/compiler/data"
	"github.com/anstk/smokel/compiler/messenger"
)

type lnx struct {
	line  entityLine
	msg   messenger.Writer
	debug bool
}

func New(msg messenger.Writer, debug ...bool) *lnx {
	return &lnx{
		line:  newLine(),
		msg:   msg,
		debug: len(debug) > 0,
	}
}

func (lx *lnx) Lex(name string, r io.Reader) []data.Line {

	lines := lx.lex(name, r)

	if lx.debug {
		for _, l := range lines {
			l.Describe()
		}
	}

	return lines
}

func (lx *lnx) lex(name string, r io.Reader) []data.Line {

	lines := readerToLines(r, name)

	i := 0

	// Used "for" and not "for range", because we know initial length of "[]line",
	// but we don't know the final length of "[]line" .
	//
	// When append new "[]line" from files, we don't known how many "[]line" will have,
	// When append new "[]line" from files, we don't how many other append those files will have.
	// At each loop, we check the length of "[]line"  with "len(lines)" function.
	// And we break the loop when all "[]line" have been processed. "if len(lines) == i"
	for {

		current := &lines[i] // important, pointer to outside for loop, to persist data changes inside loop

		//---------------------------

		err := lx.line.Identifier.Identify(current)
		if err != nil {
			lx.msg.Error(current.Number, current.Name, err)
			break
		}

		//lx.line.Modifier.Type(current, identified)

		//---------------------------

		if current.IsIncludeGroup() {
			lx.line.Modifier.Empty(current)
		}

		if current.IsInclude() {
			fileName := current.Formatted
			lx.line.Modifier.Empty(current)

			lines, err = appendLinesFromFile(lines, fileName)
			if err != nil {
				lx.msg.Error(current.Number, current.Name, err)
				break
			}
		}

		//---------------------------

		i++
		if len(lines) == i {
			break
		}

	}

	return lines
}
