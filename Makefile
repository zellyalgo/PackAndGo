build:
	go build -o trip cmd/main.go

test:
	go test ./... -v -cover

image:
	go build -o trip cmd/main.go
	docker build -t $(label) .