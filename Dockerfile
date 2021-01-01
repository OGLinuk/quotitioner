FROM golang:1.15
ADD . /go/src/quotitioner-12321
WORKDIR /go/src/quotitioner-12321
RUN go build -o quotitioner-12321
EXPOSE 12321
CMD ["./quotitioner-12321"]
