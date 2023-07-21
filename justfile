default: fmt

fmt:
    go fmt
    go vet

buildserver:
    GOOS=linux GOARCH=amd64 go build -o bin/linux/requestor
