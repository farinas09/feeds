ARG GO_VERSION=1.24.3

FROM golang:${GO_VERSION}-alpine as builder

RUN go env -w GOPROXY=direct
RUN apk add --no-cache git
RUN apk add --no-cache add ca-certificates && update-ca-certificates

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go install ./...

FROM alpine:latest
WORKDIR /usr/bin
COPY --from=builder /go/bin .