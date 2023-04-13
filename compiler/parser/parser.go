package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/anstk/smokel/compiler/data"
	"github.com/anstk/smokel/compiler/linexer"
	"github.com/anstk/smokel/compiler/messenger"
)

const (
	charIndent  = "\t"
	charNewLine = "\n"
	charSpace   = " "
)

type parser struct {
	// Used to Lex lines to parser
	// import "compiler/linexer"
	lx linexer.Reader

	// Used to handle errors or good practices messages
	// and output to console
	// import "compiler/messenger"
	msg messenger.Writer

	// Current "line" being parsed.
	// Its used in identification functions
	// like inside, next, etc
	line data.Line

	// Used to define which group current line is.
	groupType data.LineType

	// Used to known what will be in next "line",
	// and maybe add } or ; etc... to the returned
	// []data.ParserLine
	nextLineType data.LineType

	// Used to store all declared variable names,
	// and checks if the used variables were previously
	// declared.
	keyVars []string
}

func New(lx linexer.Reader, msg messenger.Writer) *parser {
	return &parser{
		lx:  lx,
		msg: msg,
	}
}

func (ln *parser) parse(lines []data.Line) []data.ParserLine {

	parsed := []data.ParserLine{}

	for i, line := range lines {

		//------------------------
		// SET NEXT LINE TYPE
		if i+1 < len(lines) {
			ln.nextLineType = lines[i+1].Type
		}

		// TODO: nextLine type set to EOT
		// if i == len(lines) {
		//	ln.nextLineType = EOT
		// }

		//------------------------
		// PARSE
		returned, err := ln.parseLine(line)
		if err != nil {
			//TODO messenger here
			fmt.Println(line.Number, err)
		}

		if len(returned) > 0 {
			parsed = append(parsed, returned...)
		}

		//------------------------
		// SET LAST LINE TYPE
		//ln.lastLineType = lines[i].Type

	}

	return parsed
}

/*
    NotIdentified LineType = iota

	LineCommentGroup
	LineComment
	LineIncludeGroup
	LineInclude
	LineVarGroup

	LineVar

	LineBlockGroup
	LineBlockName
	LineBlockProperty
	LineBlock
	LineExtendGroup
	LineExtendElement
	LineExtendProperty
	LineRawGroup
	LineRaw
	LineMedia
	LineMediaElement
	LineMediaProperty
	LineEmpty
	LineElement
	LineProperty
*/

func (ln *parser) parseLine(line data.Line) ([]data.ParserLine, error) {

	var empty = []data.ParserLine{}

	//----------------------------
	// Current "line" in ln.line
	// its used in identification functions
	// like inside, next, etc
	ln.line = line

	//----------------------------
	//CLEAN
	if line.IsEmpty() &&
		line.IsComment() &&
		line.IsCommentGroup() {
		return empty, nil
	}

	//SET GROUP  ------------------------
	if line.IsBlockGroup() {
		ln.groupType = data.LineBlockGroup
		return empty, nil
	}

	if line.IsExtendGroup() {
		ln.groupType = data.LineExtendGroup
		return empty, nil
	}

	if line.IsVarGroup() {
		ln.groupType = data.LineVarGroup
		return empty, nil
	}

	if line.IsElement() {
		ln.groupType = data.LineElement
		// do not use return here
		// element must still be parsed
	}

	// INSIDE GROUPS ------------------------
	if ln.insideVarGroup() {

		if line.IsVar() {
			key, value := ln.split()
			ln.addKeyVar(key)

			// is second parameter is a $variable
			// check if it have in ln.keyVars
			if matchVar(value) {
				if !ln.checkKeyVar(value) {
					return empty, errors.New(fmt.Sprintln("Variable", value, "not declared."))
				}
			}

			return ln.createParserLines(
				fmt.Sprint("{{", key, " := ", value, "}}"),
			), nil
		}
	}

	if ln.insideElement() {

		if line.IsProperty() {

			if !ln.nextIsProperty() {
				return ln.createParserLines(
					fmt.Sprint(charIndent, line.Formatted, ";", charNewLine),
					fmt.Sprint("}", charNewLine),
				), nil
			}

			return ln.createParserLines(
				fmt.Sprint(charIndent, line.Formatted, ";", charNewLine),
			), nil

		}
	}

	// -------------------------

	if line.IsElement() {

		if ln.nextIsProperty() {
			if strings.HasSuffix(line.Formatted, ",") {
				trined := strings.TrimSuffix(line.Formatted, ",")
				return ln.createParserLines(
					fmt.Sprint(trined, charSpace, "{", charNewLine),
				), nil

			}

			return ln.createParserLines(
				fmt.Sprint(line.Formatted, charSpace, "{", charNewLine),
			), nil
		}

		if ln.nextIsElement() {
			if strings.HasSuffix(line.Formatted, ",") {
				return ln.createParserLines(line.Formatted), nil
			}

			return ln.createParserLines(
				fmt.Sprint(line.Formatted, ","),
			), nil
		}

		return ln.createParserLines(line.Formatted), nil

	}

	return empty, nil
}

// -----------------
// INSIDE

func (ln *parser) insideVarGroup() bool {
	return ln.groupType == data.LineVarGroup
}

func (ln *parser) insideElement() bool {
	return ln.groupType == data.LineElement
}

// ----------------------
// NEXT LINE TYPE
func (ln *parser) nextIsProperty() bool {
	return ln.nextLineType == data.LineProperty
}

func (ln *parser) nextIsElement() bool {
	return ln.nextLineType == data.LineElement
}

// ---------------------------
func (ln *parser) split() (string, string) {
	values := strings.Split(ln.line.Formatted, ":")
	return strings.TrimSpace(values[0]), strings.TrimSpace(values[1])
}

// --------------------------
// KEY

func (ln *parser) addKeyVar(key string) {
	ln.keyVars = append(ln.keyVars, key)
}

func (ln *parser) checkKeyVar(key string) bool {
	for _, kv := range ln.keyVars {
		if kv == key {
			return true
		}
	}
	return false
}

// --------------------------
// MATCH
func matchStringInterpolation(text string) bool {
	return strings.Contains(text, "{") && strings.Contains(text, "$") && strings.Contains(text, "}")
}

func matchVar(text string) bool {
	return strings.HasPrefix(text, "$") && len(text) > 1
}

//--------------------------
