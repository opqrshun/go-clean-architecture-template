#!/usr/bin/env bash

PROJECT_DIR=sample
# cp -r clean-architecture $PROJECT_DIR
entity=$1
Entity=${entity^}
mkdir $PROJECT_DIR 
rsync -av ./ $PROJECT_DIR/ # --exclude "entity*"

sed -i "s/entity/${entity}/g" $PROJECT_DIR
sed -i "s/Entity/${Entity}/g" $PROJECT_DIR

find $PROJECT_DIR -name '*entity*' | xargs -L 1 rename entity ${entity}
