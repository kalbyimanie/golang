build-builder:
	docker build -t builder:v1 -f Dockerfile .
	docker run --rm --name builder -v $(GOPATH):/go -w /go builder:v1 go build -o /go/src/main/main /go/src/main/main.go

build-service:
	docker build -t web-app:v2 -f service.Dockerfile .
