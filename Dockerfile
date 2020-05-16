FROM golang:1.9 as builder

ENV CGO_ENABLED=0 GOOS=linux
RUN go get github.com/aquilax/go-perlin
RUN go get github.com/gorilla/websocket

WORKDIR /go/src/github.com/joram/game-server/
COPY . /go/src/github.com/joram/game-server
RUN go build -a -installsuffix cgo -o build/game-server github.com/joram/game-server/cmd/game-server

FROM alpine:latest
RUN apk add --no-cache ca-certificates
EXPOSE 2303
WORKDIR /
ENTRYPOINT ["/bin/game-server"]
RUN mkdir /static
ADD ./static/* /static/
RUN ls -hal /static
COPY --from=builder /go/src/github.com/joram/game-server/build/ /bin/
