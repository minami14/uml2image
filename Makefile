build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build -o uml2image cmd/main.go
	docker build -t uml2image .
