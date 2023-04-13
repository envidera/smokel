package data

import (
	"fmt"
)

type LineType int

const (
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
)

type Line struct {

	// number is used to store Line number
	// it will logged in case of compiler error.
	Number int

	// Name is used to store the (file name) or (io.Reader name).
	// It helps to identify whats file error come from.
	// It will logged in case of compiler error.
	Name string

	//isType is the identified line type
	Type LineType

	// Text is the line content itself
	Text string

	// Text formatted used by linexer
	Formatted string

	// Can be 0, 1 or 2, only
	Indent int
}

func (t *Line) IsNotIdentified() bool  { return t.Type == NotIdentified }
func (t *Line) IsCommentGroup() bool   { return t.Type == LineCommentGroup }
func (t *Line) IsComment() bool        { return t.Type == LineComment }
func (t *Line) IsIncludeGroup() bool   { return t.Type == LineIncludeGroup }
func (t *Line) IsInclude() bool        { return t.Type == LineInclude }
func (t *Line) IsVarGroup() bool       { return t.Type == LineVarGroup }
func (t *Line) IsVar() bool            { return t.Type == LineVar }
func (t *Line) IsBlockGroup() bool     { return t.Type == LineBlockGroup }
func (t *Line) IsBlockName() bool      { return t.Type == LineBlockName }
func (t *Line) IsBlockProperty() bool  { return t.Type == LineBlockProperty }
func (t *Line) IsBlock() bool          { return t.Type == LineBlock }
func (t *Line) IsExtendGroup() bool    { return t.Type == LineExtendGroup }
func (t *Line) IsExtendElement() bool  { return t.Type == LineExtendElement }
func (t *Line) IsExtendProperty() bool { return t.Type == LineExtendProperty }
func (t *Line) IsRawGroup() bool       { return t.Type == LineRawGroup }
func (t *Line) IsRaw() bool            { return t.Type == LineRaw }
func (t *Line) IsMedia() bool          { return t.Type == LineMedia }
func (t *Line) IsMediaElement() bool   { return t.Type == LineMediaElement }
func (t *Line) IsMediaProperty() bool  { return t.Type == LineMediaProperty }
func (t *Line) IsEmpty() bool          { return t.Type == LineEmpty }
func (t *Line) IsElement() bool        { return t.Type == LineElement }
func (t *Line) IsProperty() bool       { return t.Type == LineProperty }

func (t *Line) NotIdentified()  { t.Type = NotIdentified }
func (t *Line) CommentGroup()   { t.Type = LineCommentGroup }
func (t *Line) Comment()        { t.Type = LineComment }
func (t *Line) IncludeGroup()   { t.Type = LineIncludeGroup }
func (t *Line) Include()        { t.Type = LineInclude }
func (t *Line) VarGroup()       { t.Type = LineVarGroup }
func (t *Line) Var()            { t.Type = LineVar }
func (t *Line) BlockGroup()     { t.Type = LineBlockGroup }
func (t *Line) BlockName()      { t.Type = LineBlockName }
func (t *Line) BlockProperty()  { t.Type = LineBlockProperty }
func (t *Line) Block()          { t.Type = LineBlock }
func (t *Line) ExtendGroup()    { t.Type = LineExtendGroup }
func (t *Line) ExtendElement()  { t.Type = LineExtendElement }
func (t *Line) ExtendProperty() { t.Type = LineExtendProperty }
func (t *Line) RawGroup()       { t.Type = LineRawGroup }
func (t *Line) Raw()            { t.Type = LineRaw }
func (t *Line) Media()          { t.Type = LineMedia }
func (t *Line) MediaElement()   { t.Type = LineMediaElement }
func (t *Line) MediaProperty()  { t.Type = LineMediaProperty }
func (t *Line) Empty()          { t.Type = LineEmpty }
func (t *Line) Element()        { t.Type = LineElement }
func (t *Line) Property()       { t.Type = LineProperty }

// TypeNotIdentified set Line Type to NotIdentified
/*
func (t *Line) TypeNotIdentified() LineType  { return NotIdentified }
func (t *Line) TypeCommentGroup() LineType   { return LineCommentGroup }
func (t *Line) TypeComment() LineType        { return LineComment }
func (t *Line) TypeIncludeGroup() LineType   { return LineIncludeGroup }
func (t *Line) TypeInclude() LineType        { return LineInclude }
func (t *Line) TypeVarGroup() LineType       { return LineVarGroup }
func (t *Line) TypeVar() LineType            { return LineVar }
func (t *Line) TypeBlockGroup() LineType     { return LineBlockGroup }
func (t *Line) TypeBlockName() LineType      { return LineBlockName }
func (t *Line) TypeBlockProperty() LineType  { return LineBlockProperty }
func (t *Line) TypeBlock() LineType          { return LineBlock }
func (t *Line) TypeExtendGroup() LineType    { return LineExtendGroup }
func (t *Line) TypeExtendElement() LineType  { return LineExtendElement }
func (t *Line) TypeExtendProperty() LineType { return LineExtendProperty }
func (t *Line) TypeRawGroup() LineType       { return LineRawGroup }
func (t *Line) TypeRaw() LineType            { return LineRaw }
func (t *Line) TypeMedia() LineType          { return LineMedia }
func (t *Line) TypeMediaElement() LineType   { return LineMediaElement }
func (t *Line) TypeMediaProperty() LineType  { return LineMediaProperty }
func (t *Line) TypeEmpty() LineType          { return LineEmpty }
func (t *Line) TypeElement() LineType        { return LineElement }
func (t *Line) TypeProperty() LineType       { return LineProperty }
*/
//-------------------------------------------------------------------------------------

func (t Line) Describe() {
	fmt.Println(t.SDescribe())
}

func (t Line) SDescribe() string {

	return fmt.Sprint(" Name: ", t.Name, "\n",
		"    |   Number: ", t.Number, "\n",
		"    |   Indent: ", t.Indent, "\n",
		"    |     Text: \"", t.Text, "\" \n",
		"    | Formatted: \"", t.Formatted, "\"\n",
		"    |     Type: ", t.TypeToName(t.Type), "\n\n",
	)

}

func (t Line) TypeToName(lnType LineType) string {
	name := []string{
		"NotIdentified",
		"CommentGroup",
		"Comment",
		"IncludeGroup",
		"Include",
		"VarGroup",
		"Var",
		"BlockGroup",
		"BlockName",
		"BlockProperty",
		"Block",
		"ExtendGroup",
		"ExtendElement",
		"ExtendProperty",
		"RawGroup",
		"Raw",
		"Media",
		"MediaElement",
		"MediaProperty",
		"Empty",
		"Element",
		"Property",
	}

	return name[lnType]
}
