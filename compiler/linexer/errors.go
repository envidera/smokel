package linexer

import (
	"errors"

	"github.com/anstk/smokel/compiler/messenger"
)

const (
	errCheckManual = "\nPlease, check the smokel syntax manual. If error persist, contact the developer"
)

var (
	errVarGroupJustVar = messenger.NewMessage(`inside a "var" group, the first indent level must contain only variable declarations`,
		`
-1|var
0|    $color: red    # variable`)

	errBlockGroupJustBlockName = messenger.NewMessage(`inside a "block" group, the first indent level must be the block name`,
		`
-1|block
0|    $blockName
1|        color: red`)

	errBlockGroupJustProperty = messenger.NewMessage(`inside a block group, the second level is block properties`,
		`	
-2|block
-1|    $blockName
0|        color: red    # block property`)
)

var (
	errIndentError   = errors.New("smokel indent syntax error, max indent is two" + errCheckManual)
	errNotIdentified = errors.New("type not identified." + errCheckManual)

	errMediaJustElement = errors.New(`inside media, the first level is element	
 | example:	
 | @media (max-width: 600px)
 |    body                # element
 |        color : red
 |
 |    .nav                # element
 |        color : blue`)

	errMediaJustThese = errors.New(`inside media, the second level must be a property	or a block
 | example:	
 | @media (max-width: 600px)
 |    body 
 |        color : red     # property
 |        $myBlock        # block
 |
 |    .nav
 |        color : blue    # property `)

	errElementJustThese = errors.New(`inside element, the first level must be a property	or a block
 | example:	
 | body 
 |     color : red    # property
 |     $myBlock       # block`)

	errElementJustOneLevel = errors.New(`inside element, must have just one indentation level
 | example:	
 | body               # element in level 0
 |     color : red    # property in level 1`)

	errExtendJustElement = errors.New(`inside extend, the first level is element	
| example:	
| extend
|    .btn    # element`)

	errExtendJustThese = errors.New(`inside extend, the second level must be a property or a block
| example:	
| extend
|    .btn
|        color : red    # property 
|        $myBlock       # block`)
)
