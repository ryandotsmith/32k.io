GOARCH=amd64 GOOS=linux go build -o /tmp/server main.go
ssh ubuntu@server.32k.io 'sudo systemctl stop server'
scp /tmp/server ubuntu@server.32k.io:~
ssh ubuntu@server.32k.io 'sudo systemctl restart server'
ssh ubuntu@server.32k.io 'sudo systemctl status server'
