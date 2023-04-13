package linexer

import (
	"io"

	"github.com/anstk/smokel/compiler/data"
)

type Reader interface {
	Lex(name string, r io.Reader) []data.Line
}
