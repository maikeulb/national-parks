FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/maikeulb/national-parks

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -o national-parks -a -installsuffix cgo main.go 


FROM alpine:latest

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/maikeulb/national-parks .

CMD ["./national-parks"]
