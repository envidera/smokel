package linexer

import (
	"strings"

	"github.com/anstk/smokel/compiler/data"
)

type lni struct {
	groupType data.LineType
	line      data.Line
}

func (ln *lni) Identify(line *data.Line) error {

	ln.line = *line

	err := ln.detectIndentError()
	if err != nil {
		//line.NotIdentified()
		return err
	}

	err = ln.detect(line)
	if err != nil {
		//line.NotIdentified()
		return err
	}

	return nil
}

// more than two indents is prohibited in smokel syntax
// only raw lines can have many indents
func (ln *lni) detectIndentError() error {
	if ln.moreThanTwoIndent() && !ln.insideRawGroup() {
		return errIndentError
	}

	return nil
}

func (ln *lni) detect(line *data.Line) error {

	// This must be the first
	// include .css always as raw lines
	if strings.HasSuffix(line.Name, ".css") {
		line.Raw()
		return nil
	}

	// .kel will continue and be detected below >

	if ln.isEmpty() {
		line.Empty()
		return nil
	}

	if ln.isComment() {
		line.Comment()
		return nil
	}

	if ln.isCommentGroup() {
		ln.groupType = data.LineCommentGroup
		line.CommentGroup()
		return nil
	}

	if ln.isIncludeGroup() {
		ln.groupType = data.LineIncludeGroup
		line.IncludeGroup()
		return nil
	}

	if ln.isVarGroup() {
		ln.groupType = data.LineVarGroup
		line.VarGroup()
		return nil
	}

	if ln.isBlockGroup() {
		ln.groupType = data.LineBlockGroup
		line.BlockGroup()
		return nil
	}

	if ln.isExtendGroup() {
		ln.groupType = data.LineExtendGroup
		line.ExtendGroup()
		return nil
	}

	if ln.isMedia() {
		ln.groupType = data.LineMedia
		line.Media()
		return nil
	}

	if ln.isRawGroup() {
		ln.groupType = data.LineRawGroup
		line.RawGroup()
		return nil
	}

	// INSIDE GROUPS ------------------------

	if ln.insideCommentGroup() {
		line.Comment()
		return nil
	}

	if ln.insideIncludeGroup() {
		if ln.isInclude() {
			line.Include()
			return nil
		}
	}

	if ln.insideVarGroup() {
		if ln.isVar() {
			line.Var()
			return nil
		}

		line.NotIdentified()
		return errVarGroupJustVar
	}

	if ln.insideBlockGroup() {

		if ln.oneIndent() {
			if ln.isBlockName() {
				line.BlockName()
				return nil
			}
			line.NotIdentified()
			return errBlockGroupJustBlockName
		}

		if ln.twoIndent() {
			if ln.isBlockProperty() {
				line.BlockProperty()
				return nil
			}
			line.NotIdentified()
			return errBlockGroupJustProperty
		}

	}

	if ln.insideExtendGroup() {

		if ln.oneIndent() {
			if ln.isExtendElement() {
				line.ExtendElement()
				return nil
			}

			line.NotIdentified()
			return errExtendJustElement
		}

		if ln.twoIndent() {
			if ln.isExtendProperty() {
				line.ExtendProperty()
				return nil
			}

			if ln.isBlock() {
				line.Block()
				return nil
			}

			line.NotIdentified()
			return errExtendJustThese
		}

	}

	if ln.insideMedia() {

		if ln.oneIndent() {
			if ln.isMediaElement() {
				line.MediaElement()
				return nil
			}
			line.NotIdentified()
			return errMediaJustElement
		}

		if ln.twoIndent() {
			if ln.isMediaProperty() {
				line.MediaProperty()
				return nil
			}

			if ln.isBlock() {
				line.Block()
				return nil
			}

			line.NotIdentified()
			return errMediaJustThese
		}

	}

	if ln.insideRawGroup() {
		line.Raw()
		return nil
	}

	//---------------------------------------------------
	// property and element must be the last because they
	// have weak match options

	if ln.insideElement() {

		if ln.oneIndent() {
			if ln.isElementProperty() {
				line.Property()
				return nil
			}

			if ln.isBlock() {
				line.Block()
				return nil
			}

			line.NotIdentified()
			return errElementJustThese
		}

		if ln.twoIndent() {
			line.NotIdentified()
			return errElementJustOneLevel
		}

	}

	if ln.isElement() {
		ln.groupType = data.LineElement
		line.Element()
		return nil
	}

	// If line is nothing above, so its not identified
	line.NotIdentified()
	return errNotIdentified

}

// ------------------------------------------------------------

func (ln *lni) noIndent() bool {
	return ln.line.Indent == 0
}

func (ln *lni) haveIndent() bool {
	return ln.line.Indent == 1 || ln.line.Indent == 2
}

func (ln *lni) oneIndent() bool {
	return ln.line.Indent == 1
}

func (ln *lni) twoIndent() bool {
	return ln.line.Indent == 2
}

func (ln *lni) moreThanTwoIndent() bool {
	return ln.line.Indent == -1
}

// ------------------------------------------------------------

func (ln *lni) insideCommentGroup() bool {
	return ln.haveIndent() && ln.groupType == data.LineCommentGroup
}

func (ln *lni) insideIncludeGroup() bool {
	return ln.haveIndent() && ln.groupType == data.LineIncludeGroup
}

func (ln *lni) insideVarGroup() bool {
	return ln.haveIndent() && ln.groupType == data.LineVarGroup
}

func (ln *lni) insideBlockGroup() bool {
	return ln.haveIndent() && ln.groupType == data.LineBlockGroup
}

func (ln *lni) insideExtendGroup() bool {
	return ln.haveIndent() && ln.groupType == data.LineExtendGroup
}

func (ln *lni) insideMedia() bool {
	return ln.haveIndent() && ln.groupType == data.LineMedia
}

func (ln *lni) insideRawGroup() bool {
	return ln.groupType == data.LineRawGroup
}

func (ln *lni) insideElement() bool {
	return ln.haveIndent() && ln.groupType == data.LineElement
}

// ------------------------------------------------------------

func (ln *lni) isVarGroup() bool {
	return ln.noIndent() && ln.line.Formatted == "var"
}

func (ln *lni) isVar() bool {
	return ln.insideVarGroup() && ln.matchVar()
}

func (ln *lni) isRawGroup() bool {
	return ln.noIndent() && ln.line.Formatted == "raw"
}

func (ln *lni) isBlockGroup() bool {
	return ln.noIndent() && ln.line.Formatted == "block"
}

func (ln *lni) isBlockName() bool {
	return ln.insideBlockGroup() && ln.oneIndent() && strings.HasPrefix(ln.line.Formatted, "$")
}

func (ln *lni) isBlockProperty() bool {
	return ln.insideBlockGroup() && ln.twoIndent() && ln.matchProperty()
}

func (ln *lni) isBlock() bool {
	return ln.insideElement() && ln.matchBlock() ||
		ln.insideMedia() && ln.matchBlock() ||
		ln.insideExtendGroup() && ln.matchBlock()
}

func (ln *lni) isMedia() bool {
	return ln.noIndent() && strings.HasPrefix(ln.line.Formatted, "@media")
}

func (ln *lni) isMediaElement() bool {
	return ln.insideMedia() && ln.oneIndent()
}

func (ln *lni) isMediaProperty() bool {
	return ln.insideMedia() && ln.twoIndent() && ln.matchProperty()
}

func (ln *lni) isCommentGroup() bool {
	return strings.HasPrefix(ln.line.Formatted, "/*")
}

func (ln *lni) isComment() bool {
	return strings.HasPrefix(ln.line.Formatted, "//") || strings.HasPrefix(ln.line.Formatted, "#")
}

func (ln *lni) isIncludeGroup() bool {
	return ln.noIndent() && ln.line.Formatted == "include"
}

func (ln *lni) isInclude() bool {
	return ln.insideIncludeGroup() && ln.oneIndent()
}

func (ln *lni) isExtendGroup() bool {
	return ln.noIndent() && ln.line.Formatted == "extend"
}

func (ln *lni) isExtendElement() bool {
	return ln.insideExtendGroup() && ln.oneIndent()
}

func (ln *lni) isExtendProperty() bool {
	return ln.insideExtendGroup() && ln.twoIndent() && ln.matchProperty()
}

func (ln *lni) isElement() bool {
	return ln.noIndent()
}

func (ln *lni) isElementProperty() bool {
	return ln.insideElement() && ln.oneIndent() && ln.matchProperty()
}

func (ln *lni) isEmpty() bool {
	return ln.line.Formatted == ""
}

// ------------------------------------------------------------

func (ln *lni) matchBlock() bool {
	return strings.HasPrefix(ln.line.Formatted, "$") && !strings.Contains(ln.line.Formatted, ":")
}

func (ln *lni) matchProperty() bool {
	return !strings.HasPrefix(ln.line.Formatted, "$") && strings.Contains(ln.line.Formatted, ":")
}

func (ln *lni) matchVar() bool {
	return strings.HasPrefix(ln.line.Formatted, "$") && strings.Contains(ln.line.Formatted, ":")
}
