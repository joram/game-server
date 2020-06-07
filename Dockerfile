FROM golang:1.10 as builder

#ENV CGO_ENABLED=0 GOOS=linux
ENV GOOS=linux
RUN go get github.com/aquilax/go-perlin
RUN go get github.com/gorilla/websocket
RUN go get github.com/mattn/go-sqlite3

WORKDIR /go/src/github.com/joram/game-server/
ADD . /go/src/github.com/joram/game-server
RUN go build -ldflags "-linkmode external -extldflags -static" -a -installsuffix cgo -o build/game-server github.com/joram/game-server/cmd/game-server

FROM alpine:latest
RUN apk add --no-cache ca-certificates
EXPOSE 2305
WORKDIR /
ENTRYPOINT ["/bin/game-server"]
RUN mkdir /static
COPY ./static/ /static/
COPY --from=builder /go/src/github.com/joram/game-server/build/ /bin/
