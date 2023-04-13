
# Functions used by other shell scripts.

#----------------------------------------------------------------------
set -e # abort on error
cd "../../" # go to root
#----------------------------------------------------------------------

# USAGE
# path (optional)
# example
#  echoFinish
#  or
#  echoFinish path-to-check
function echoFinish(){
    echo "--------------------------------------------------------"
    echo " Finished $(basename $0)" 
    if [ -n "$1" ]
    then
        echo " Check $1 path"
    fi
    echo "--------------------------------------------------------"
}


