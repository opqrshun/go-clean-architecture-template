#!/usr/bin/env bash

sed -i "s/EntityV1Category/Category/g" $(fd -e go --full-path ./ ) 
sed -i "s/EntityV1Tag/Tag/g" $(fd -e go --full-path ./ ) 
sed -i "s/EntityV1Latlng/Latlng/g" $(fd -e go --full-path ./ ) 
sed -i "s/AttributeV1/Attribute/g" $(fd -e go --full-path ./ ) 
sed -i "s/EntityV1/Entity/g" $(fd -e go --full-path ./ ) 
sed -i "s/EntityV1User/User/g" $(fd -e go --full-path ./ ) 

sed -i "s/UpdatedTime string/UpdatedAt time.Time/g" $(fd -e go --full-path ./ ) 
sed -i "s/CreatedTime string/CreatedAt time.Time/g" $(fd -e go --full-path ./ ) 
sed -i "s/DeletedTime string/DeletedAt time.Time/g" $(fd -e go --full-path ./ ) 

sed -i "s/package http/package domain/g" $(fd -e go --full-path ./ ) 
sed -i 's/int32/int/g' $(fd -e go)
sed -i 's/int64/int/g' $(fd -e go)

find ./ -name '*attribute_entity_v1_*' | xargs -L 1 rename 'attribute_entity_v1_' ''
find ./ -name '*attribute_*' | xargs -L 1 rename 'attribute_' ''
find ./ -name '*_v1*' | xargs -L 1 rename '_v1' ''   
