package messenger

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type messenger struct {
	writer io.Writer
}

func New(writer ...io.Writer) *messenger {

	if len(writer) > 0 {
		return &messenger{
			writer: writer[0],
		}
	}

	return &messenger{
		writer: os.Stdout,
	}
}

func (m *messenger) ErrorX(lineNumber int, lineName string, lineText string, msg Message) {
	fmt.Fprint(m.writer, m.format("ERROR", lineNumber, lineName, lineText, msg))
}

//func (m *messenger) Error(lineNumber int, lineName string, text any) {
//	fmt.Fprint(m.writer, format("ERROR", lineNumber, lineName, text))
//}

//func (m *messenger) Alert(lineNumber int, lineName string, text any) {
//	fmt.Fprint(m.writer, format("GOOD PRACTICE", lineNumber, lineName, text))
//}

func (m *messenger) format(msgPrefix string, lineNumber int, lineName string, lineText string, msg Message) string {
	return fmt.Sprint(msgPrefix, " > ", lineName, "\n",
		msg.text, "\n",
		"\n",
		lineNumber, "|", lineText, "\n",
		"\n",
		formatExample(lineNumber, msg.example),
		"----------------------------------------------------------------\n")
}

/*
format must be text/string in format

-2| //comment test
-1|var
0|    $wrong : variable indentation
1|    ...
2|    ...

0 means the current line error example
-1 means current line error example -1
*/
func formatExample(lineNumber int, example string) string {
	formatted := "\n"

	lines := strings.Split(example, "\n")

	for i, line := range lines {

		//if have some empty line, just skip
		if strings.TrimSpace(line) == "" {
			continue
		}

		values := strings.Split(line, "|")
		number, _ := strconv.Atoi(values[0])
		text := values[1]

		// if number is not "0" the current msg line number
		if number != 0 {
			formatted += strconv.Itoa(lineNumber + number)
		} else {
			formatted += strconv.Itoa(lineNumber)
		}

		// put newline "\n" in all lines, except last one
		if i != len(lines)-1 {
			formatted += "|" + text + "\n"
		} else {
			formatted += "|" + text
		}

	}

	return formatted
}

/*

ERROR > main.kel
inside a "var" group, the first indent level must have only variables

2|    body

example:
1|var
2|    $wrong : variable indentation
----------------------------------------------------------------
ERROR > main.kel
inside a "var" group, the first indent level must have only variables

220|    body

example:
219|var
220|    $wrong : variable indentation
----------------------------------------------------------------
*/

/*
ERROR
checkSyntax:2 > inside a "var" group, the first indent level must have only variables

2|$wrong : variable indentation

example:
1|var
2|    $wrong : variable indentation
----------------------------------------------------------------
*/
