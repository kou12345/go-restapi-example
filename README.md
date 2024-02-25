# go-restapi-example

## サーバー起動

```
go run ./api/cmd/main.go
```

## Swagger UI の更新

```
swag init -g api/cmd/main.go --parseDependency
```

## Swagger UI の URL

http://localhost:8080/v1/swagger/index.html
