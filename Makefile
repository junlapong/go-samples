.EXPORT_ALL_VARIABLES:
GOOS=linux
GOARCH=amd64

asm:
	@go build -o bin/main -gcflags -S main.go

