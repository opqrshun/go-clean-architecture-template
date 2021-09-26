#!/usr/bin/env bash

PROJECT_DIR=$1
SHELL_DIR=`dirname $0`
# cp -r clean-architecture $PROJECT_DIR
model=$2
Model=${model^}
echo $Model
mkdir $PROJECT_DIR 
rsync -av $SHELL_DIR/ $PROJECT_DIR/ # --exclude "model*"

sed -i "s/go-clean-architecture/${PROJECT_DIR}/g" $(fd -e go --full-path $PROJECT_DIR ) 
sed -i "s/model/${model}/g" $(fd -e go --full-path $PROJECT_DIR ) 
sed -i "s/Model/${Model}/g" $(fd -e go --full-path $PROJECT_DIR ) 

find $PROJECT_DIR -name '*model*' | xargs -L 1 rename model ${model}
