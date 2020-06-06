#!/usr/bin/env bash

PROJECT_DIR=sample
SHELL_DIR=`dirname $0`
# cp -r clean-architecture $PROJECT_DIR
entity=$1
Entity=${entity^}
echo $Entity
mkdir $PROJECT_DIR 
rsync -av $SHELL_DIR/ $PROJECT_DIR/ # --exclude "entity*"

sed -i "s/entity/${entity}/g" $(fd -e go --full-path $PROJECT_DIR ) 
sed -i "s/Entity/${Entity}/g" $(fd -e go --full-path $PROJECT_DIR ) 

find $PROJECT_DIR -name '*entity*' | xargs -L 1 rename entity ${entity}
