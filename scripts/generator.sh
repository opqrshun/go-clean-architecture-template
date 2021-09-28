#!/usr/bin/env bash

PROJECT_DIR=$1
SHELL_DIR=`dirname $0`
# cp -r clean-architecture $PROJECT_DIR
attribute=$2
Attribute=${attribute^}
echo $Attribute
mkdir $PROJECT_DIR 
rsync -av $SHELL_DIR/ $PROJECT_DIR/ # --exclude "attribute*"

sed -i "s/go-clean-architecture/${PROJECT_DIR}/g" $(fd -e go --full-path $PROJECT_DIR ) 
sed -i "s/attribute/${attribute}/g" $(fd -e go --full-path $PROJECT_DIR ) 
sed -i "s/Attribute/${Attribute}/g" $(fd -e go --full-path $PROJECT_DIR ) 

find $PROJECT_DIR -name '*attribute*' | xargs -L 1 rename attribute ${attribute}
