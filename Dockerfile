# -------------------- Build --------------------
# Start from Alpine Linux image with the latest version of Golang
# Naming build stage as builder
FROM golang:alpine as builder

# Install Git for go get
RUN set -eux;\
  apk add --no-cache --virtual git &&\
  apk add --no-cache --virtual curl

# Set ENV
ENV GOPATH /go/
ENV GO_WORKDIR $GOPATH/src/diapi-mock-server

# Set the Current Working Directory inside the container
WORKDIR $GO_WORKDIR

ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy everything from the current directory to the 
# PWD (Present Working Directory) inside the container
COPY . .

# Build it
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install

# -------------------- Ready --------------------
# Start from a raw Alpine Linux image
FROM alpine:latest

# Set ENV
ENV PORT 5555

# Install ca-certificates for ssl
RUN set -eux; \
  apk add --no-cache --virtual ca-certificates

# Set WORKDIR to go execute directory
WORKDIR /app

# Copy binary from builder stage into image
COPY --from=builder /go/bin/diapi-mock-server /app

EXPOSE $PORT
ENTRYPOINT ./diapi-mock-server
