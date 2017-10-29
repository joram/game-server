FROM golang:1.9 as builder
WORKDIR /go/src/github.com/joram/game-server/
ENV CGO_ENABLED=0 GOOS=linux

COPY vendor /go/src/github.com/joram/game-server/vendor

RUN ls
RUN pwd

RUN go build -a -installsuffix cgo -o build/game-server github.com/joram/game-server/cmd/game-server

FROM alpine:latest
RUN apk add --no-cache ca-certificates
ENTRYPOINT ["/bin/game-server"]

COPY --from=builder /go/src/github.com/joram/game-server/build/ /bin/
