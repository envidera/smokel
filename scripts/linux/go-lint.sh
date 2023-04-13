#!/usr/bin/env bash
# github.com/anstk - MIT LICENSE

#-------------------------------------------------------------------------------------
# Run multiple lint tests
#-------------------------------------------------------------------------------------

# CONFIG

outputPath="tmp/lints"

commands_to_run=(
	# "command ; file name; about"
	"go mod tidy; go-mod-tidy; Ensure that all imports are satisfied and detects when assignments to existing variables are not used"
	"go vet ./...; go-vet; Finds subtle issues where your code may not work as intended."
	"revive ./...; revive; Makes code style recommendations - Fast, configurable, extensible, flexible, and beautiful linter for Go. Drop-in replacement of golint. revive.run"
	"errcheck ./...; errcheck; Check for unchecked errors in go programs. These unchecked errors can be critical bugs in some cases. github.com/kisielk/errcheck" 
	"staticcheck ./...; staticcheck; Finds bugs, performance issues, offers simplifications, and enforces style rules. staticcheck.io"
	"goconst ./...; goconst; Find repeated strings that could be replaced by a constant github.com/jgautheron/goconst"	
	"misspell ./*; misspell; Correct commonly misspelled English words github.com/client9/misspell"  # misspell ./.. doesn't work properly, use misspel ./* instead
	"gosec ./...; gosec; Inspects source code for security problems by scanning the Go AST github.com/securego/gosec"
	"gocyclo -over 3 -avg .; gocyclo; Is a code quality metric which can be used to identify code that needs refactoring. github.com/fzipp/gocyclo"
)


#----------------------------------------------------------------------
source _.sh
#----------------------------------------------------------------------


# install or update needed linters
 go install github.com/mgechev/revive@latest
 go install github.com/kisielk/errcheck@latest
 go install honnef.co/go/tools/cmd/staticcheck@latest
 go install github.com/jgautheron/goconst/cmd/goconst@latest
 go install github.com/client9/misspell/cmd/misspell@latest
 go install github.com/securego/gosec/v2/cmd/gosec@latest
 go install github.com/fzipp/gocyclo/cmd/gocyclo@latest

# more
# https://github.com/golangci/awesome-go-linters


# mkdir only if not exists
mkdir -p $outputPath

for com in "${commands_to_run[@]}"; do
	
	var=${com%;*}

	command=${com%%;*}
	filename=${var#*;}
	about=${com##*;}

	echo $command
	#echo $filename
	#echo $about
	{			
		echo "-----------------------------------------------------------------------"
		echo " $command "
		echo " $about "
		echo "-----------------------------------------------------------------------"
		$command 
	} > "$outputPath/${filename}.txt" 2>&1
done

#----------------------------------------------------------------------
echoFinish $outputPath
#----------------------------------------------------------------------



