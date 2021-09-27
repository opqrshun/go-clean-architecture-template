# gobackend

### docker
```
docker-compose -p gobackend up -d
```

### マイグレーション
```
mysql -h 127.0.0.1 -P 3338 -u gobackend -pgobackend -D gobackend < schema/base.sql
```

phpmyadmin  
http://localhost:9008/

### 環境変数
.envファイルで指定できます。

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
go build .
```

### test
```
docker exec gobackend_go_1 go test tests/* -count=1
```


