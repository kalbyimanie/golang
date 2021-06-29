FROM golang:alpine3.13 AS base_layer
RUN apk update && \
    apk upgrade && \
    apk add bash \
            git
WORKDIR /go/src/app
ENV GOPATH=/go/src/app GO111MODULE=off
EXPOSE 80

FROM base_layer AS config_files
COPY ./src/ /go/src/app/src
COPY ./pkg/ /go/src/app/pkg
COPY ./bin/ /go/src/app/bin
RUN go get -v -u github.com/gorilla/mux
RUN go build -o ./bin/main ./src/main/main.go

CMD ["./bin/main"]