install:
	go get ./...
	go install github.com/dmarkham/enumer@latest
	go generate ./...
test:
	go test ./... -count=1
