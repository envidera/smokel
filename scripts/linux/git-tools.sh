#!/usr/bin/env bash
# github.com/anstk - MIT LICENSE

#----------------------------------------------------------------------
source _.sh
#----------------------------------------------------------------------

function addExecutablesToIgnore(){
	if [ -d ".cmd" ]; then
		for file in "cmd/"*; do # for each file inside dev/cmd do
		if [ -d $file ]; then # if file is a folder then
			echo "bin/$(basename $file)"
		fi
		done
	fi
}

function creategitignore(){
	{			
		echo ".vscode/*"
		echo "#-------"
		echo "# Exclude everything in the tmp folder from being included in git"
		echo "tmp/*"
		echo "# Except this file"
		echo "!tmp/readme.md"
		echo "#-------"
		addExecutablesToIgnore	
		
	} > .gitignore

	echo "[Done] Create .gitignore"
}

function gitInit(){
	if [ ! -d ".git" ]; then

		creategitignore
	
		git init
		git add .
		git commit -m "main"
		
		git branch dev
		git branch -M main
		
		git checkout dev
		echo "[Done]"
		echo "Now in console add github remote, example:"
		echo "  git remote add origin https://github.com/anstk/my-repo.git"
		echo ""
		echo "push to github:"
		echo "  git push -u origin main"
	else
		echo "[Canceled] .git already initialized "
	fi
}

function creategitignoreAsk(){
	if [ -f ".gitignore" ]; then
		echo "Alert! .gitignore will be overwritten. Continue  (y/n)?"
		read CONT
		if [ "$CONT" = "y" ]; then
			creategitignore
		else
			echo "[Canceled]"
		fi
	else
		creategitignore
	fi
}


function gitRecreate(){
	echo "Alert! .git and .gitignore will be deleted. Continue  (y/n)?"
	read CONT
	if [ "$CONT" = "y" ]; then
		rm -drf .git
		rm .gitignore
		gitInit
		echo "[Done]"
	else
		echo "[Canceled]"
	fi
}

function gitCurrentBranchName(){
	echo $(git rev-parse --abbrev-ref HEAD)
}

function gitSoftReset(){
	
	branch=$(gitCurrentBranchName)
	echo " Branch:$branch > Clean how many commits? Type a number:"
		read CONT
		git reset --soft HEAD~$CONT
		git gc
		echo "You cleaned $CONT commits"
}

echo "Choose a number:"
echo " 1 - git Init"
echo " 2 - git Recreate"
echo " 3 - Create .gitignore"
echo '-----------------------'
echo " 4 - git soft reset"
	read CONT
	if [ "$CONT" = "1" ]; then clear;gitInit; fi
	if [ "$CONT" = "2" ]; then clear;gitRecreate; fi
	if [ "$CONT" = "3" ]; then clear;creategitignoreAsk; fi
	if [ "$CONT" = "4" ]; then clear;gitSoftReset; fi


#----------------------------------------------------------------------
echoFinish
#----------------------------------------------------------------------