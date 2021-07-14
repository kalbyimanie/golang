build-builder:
	docker build -t builder:v1 -f Dockerfile .
	docker run --rm --name builder -v $(PWD)/src:/go/src -v $(PWD)/pkg:/go/pkg -v $(PWD)/bin:/go/bin -w /go builder:v1 go build -o /go/bin/main /go/src/main/main.go

build-service:
	docker build -t web-app:v2 -f service.Dockerfile .
