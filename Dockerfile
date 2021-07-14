FROM golang:alpine3.13 AS base_layer
RUN apk update && \
    apk upgrade && \
    apk add bash \
            git
ENV GOPATH=/go GO111MODULE=off
