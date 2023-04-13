package linexer

import (
	"bufio"
	"io"
	"os"

	"github.com/anstk/smokel/compiler/data"
)

// appendLinesFromFile reads a file, converts it to []line
// and appends to the existing []line
func appendLinesFromFile(lines []data.Line, fileName string) ([]data.Line, error) {

	fileLines, err := fileToLines(fileName)
	if err != nil {
		return fileLines, err
	}

	return append(lines, fileLines...), nil
}

func fileToLines(fileName string) ([]data.Line, error) {

	reader, err := os.Open(fileName)
	if err != nil {
		return []data.Line{}, err
	}

	lines := readerToLines(reader, fileName)

	err = reader.Close()
	if err != nil {
		return []data.Line{}, err
	}

	return lines, nil

}

func readerToLines(r io.Reader, name string) []data.Line {
	lines := []data.Line{}
	count := 1 //count starts at 1 because files opened in editors start at 1 not 0

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {

		txt := scanner.Text()

		l := data.Line{
			Number:    count,
			Name:      name,
			Text:      txt,
			Formatted: cleanUp(txt),
			Indent:    indent(txt),
			// Type : will be defined in next step
		}

		lines = append(lines, l)
		count++
	}
	return lines

}
