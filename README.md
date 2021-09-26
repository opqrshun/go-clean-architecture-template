# gobackend

### docker
```
docker-compose -p gobackend up -d
```

### マイグレーション
```
mysql -h 127.0.0.1 -P 3337 -u gobackend -pgobackend -D gobackend < schema/gobackend.sql
```

phpmyadmin  
http://localhost:9007/

### 環境変数
.envファイルで指定できます。

### API動作確認
```
# ユーザー作成
curl -H "Idp-Id: xxxx-xxxx-xxxx-xxx1" -X POST "http://localhost:8088/v1/user" 

# ユーザー更新
curl -d '{"nickname":"taki"}' -H "Parent-Type: application/json" -H "Idp-Id: xxxx-xxxx-xxxx-xxx1" -X PATCH "http://localhost:8088/v1/user" 

# 投稿
curl -d '{"body":"タイトル", "body":"メッセージ１","latitude":11,"longitude":111}' -H "Parent-Type: application/json" -H "Idp-Id: xxxx-xxxx-xxxx-xxx1" -X POST "http://localhost:8088/v1/parents" 

# 取得
curl -X GET "http://localhost:8088/v1/parents/1" 

# 画像アップロード
curl -H "Idp-Id: xxxx-xxxx-xxxx-xxx1" \
  -X PUT http://localhost:8088/v1/parents/1/image\
  -F "file=@test.png" \
  -H "Parent-Type: multipart/form-data"
  
```

### APIドキュメント
https://d2td0y6q4q1g4n.cloudfront.net/document/api.html

user      gobackend
psss      XO7PRVfn

### build
```
go build .
```

### test
```
docker exec gobackend_go_1 go test tests/* -count=1
```


