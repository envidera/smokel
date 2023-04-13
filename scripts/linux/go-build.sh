#!/usr/bin/env bash
# github.com/anstk - MIT LICENSE


# Build each folder inside cmd path to buildPath.
# Its not necessary to inform the name of the folder, 
# it will be read automatically

# CONFIG
buildPath="bin"
cmdPath="cmd "

#----------------------------------------------------------------------
source _.sh
#----------------------------------------------------------------------



if [ ! -d $cmdPath ]; then
	echo "Exiting, no cmd folder found on $cmdPath"
	exit
fi

cd "$cmdPath" # go cmd folder


for file in *; do # for each file inside dev/cmd do
	if [ -d $file ]; then # if file is a folder then
		cd "${file}"
		echo "building ${file}"

		#build with clear flag
		go build  -o "../../$buildPath/${file}" -ldflags="-s -w" 

		# go up
		cd ..
	fi
done

#----------------------------------------------------------------------
echoFinish $buildPath
#----------------------------------------------------------------------