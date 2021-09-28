# gobackend

### docker
```
docker-compose --env-file .env.docker -p gobackend up -d
```

### migration

```
mysql -h 127.0.0.1 -P 3338 -u gobackend -pgobackend -D gobackend < schema/base.sql
```

phpmyadmin  
http://localhost:9008/



### API動作確認


### build
```
go build ./cmd
```

### test

if local, install godotenv
```
go get github.com/joho/godotenv/cmd/godotenv

```

run test
```
godotenv -f .env.local go test -v ./...
```


