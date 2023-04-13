package compiler

import (
	"github.com/anstk/smokel-core/layer/linexer"
	"github.com/anstk/smokel-core/layer/parser"
)

type compiler struct {
	linexer linexer.Reader
	parser  parser.Reader
}

/*
import (
	"io"
)

// Special characters
const (
	cr               = "\r"   // carriage Return
	lb               = "\n"   // line break
	crlb             = "\r\n" // carriage return and line break,
	tab              = "\t"
	space            = " "
	colon            = ":"
	comma            = ","
	openBrace        = "{"
	closeBrace       = "}"
	semicolon        = ";"
	ampersand        = "&"
	atMark           = "@"
	dollarMark       = "$"
	openParenthesis  = "("
	closeParenthesis = ")"
	slash            = "/"
	doubleSlash      = slash + slash
)

func Compile(r io.Reader, w io.Writer) error {

	return nil
}
*/
