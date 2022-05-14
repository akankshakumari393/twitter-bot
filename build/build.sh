#!/bin/bash

help(){
    echo "Please provide one argument which points to relative path of .env file "
    echo "For Example ./build.sh ./.env"
    exit 1
}

if [ $# -ne 1 ] 
then
   help
fi

set -a # export all variables created next
. $1
set +a # stop exporting

go build # build the program

./twitter-bot
