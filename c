CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server server.go
scp server admin@101.132.107.3:/home/admin/web/server