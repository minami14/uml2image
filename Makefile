build:
	go build -o uml2image cmd/main.go
	docker build -t uml2image .
