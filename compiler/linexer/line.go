package linexer

type entityLine struct {
	Identifier lni
	Modifier   lnm
}

// newLineEntity() is used to show all required variables
// in lineEntity initialization. In this case it is empty
func newLine() entityLine {
	return entityLine{}
}
