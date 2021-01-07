FROM golang:1.15
ADD . /go/src/quotitioner-9429
WORKDIR /go/src/quotitioner-9429
RUN go build -o quotitioner-9429
EXPOSE 9429
CMD ["./quotitioner-9429"]
