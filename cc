CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server1 server.go
scp server1 admin@101.132.107.3:/home/admin/web/server