#!/usr/bin/env bash

sed -i "s/ParentV1Category/Category/g" $(fd -e go --full-path ./ ) 
sed -i "s/ParentV1Tag/Tag/g" $(fd -e go --full-path ./ ) 
sed -i "s/ParentV1Latlng/Latlng/g" $(fd -e go --full-path ./ ) 
sed -i "s/ModelV1/Model/g" $(fd -e go --full-path ./ ) 
sed -i "s/ParentV1/Parent/g" $(fd -e go --full-path ./ ) 
sed -i "s/ParentV1User/User/g" $(fd -e go --full-path ./ ) 

sed -i "s/UpdatedTime string/UpdatedAt time.Time/g" $(fd -e go --full-path ./ ) 
sed -i "s/CreatedTime string/CreatedAt time.Time/g" $(fd -e go --full-path ./ ) 
sed -i "s/DeletedTime string/DeletedAt time.Time/g" $(fd -e go --full-path ./ ) 

sed -i "s/package openapi/package domain/g" $(fd -e go --full-path ./ ) 
sed -i 's/int32/int/g' $(fd -e go)
sed -i 's/int64/int/g' $(fd -e go)

find ./ -name '*model_parent_v1_*' | xargs -L 1 rename 'model_parent_v1_' ''
find ./ -name '*model_*' | xargs -L 1 rename 'model_' ''
find ./ -name '*_v1*' | xargs -L 1 rename '_v1' ''   
