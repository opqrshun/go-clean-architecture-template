# gobackend

### docker
```
docker-compose -p gobackend up -d
```

### migration

```
mysql -h 127.0.0.1 -P 3338 -u gobackend -pgobackend -D gobackend < schema/base.sql
```

phpmyadmin  
http://localhost:9008/

### 環境変数
setting .env

```
#!/usr/bin/env bash

MYSQL_ROOT_PASSWORD=gobackend
MYSQL_DATABASE=gobackend
MYSQL_USER=gobackend
MYSQL_PASSWORD=gobackend
MYSQL_PORT=3338
SERVER_PORT=8088
GO_PORT=8089
PHPMYADMIN_PORT=9008

MYSQL_PROTOCOL="tcp(mariadb:3306)"


```

### API動作確認


### build
```
go build ./cmd
```

### test

if local

```
export MYSQL_PROTOCOL="tcp(127.0.0.1:3338)"
```


run test
```
go test -v ./...
```


