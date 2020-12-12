FROM golang:1.15
ADD . /go/src
WORKDIR /go/src
RUN go build -o main
EXPOSE 8080
CMD ["./main"]
