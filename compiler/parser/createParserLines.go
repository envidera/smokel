package parser

import "github.com/anstk/smokel/compiler/data"

func (ln *parser) createParserLines(value string, values ...string) []data.ParserLine {
	lines := []data.ParserLine{
		{
			Number: ln.line.Number,
			Name:   ln.line.Name,
			Type:   int(ln.line.Type),
			Parsed: value,
		},
	}

	if len(values) > 0 {
		for _, value := range values {
			lines = append(lines, data.ParserLine{
				Number: ln.line.Number,
				Name:   ln.line.Name,
				Type:   int(ln.line.Type),
				Parsed: value,
			},
			)
		}
	}

	return lines
}
