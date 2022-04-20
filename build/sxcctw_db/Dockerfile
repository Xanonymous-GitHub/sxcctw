FROM golang:1.18-alpine

WORKDIR /usr/src/app

RUN apk add --no-cache git

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/sxcctw_db/ \
    && go clean -fuzzcache \
    && go clean -testcache \
    && go clean -modcache

CMD ["app"]