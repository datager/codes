CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -ldflags "-s -w" -v -o bin/dg main.go || (echo "编译错误" && exit 1) || exit 1
