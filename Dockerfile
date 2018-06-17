FROM golang:1.10.3-stretch as builder

COPY * /go/src/ratp-go/

WORKDIR /go/src/ratp-go
RUN go get
RUN go build

FROM debian:latest

RUN apt-get update
RUN apt-get install -y ca-certificates

COPY --from=builder /go/bin/ratp-go /ratp-go

EXPOSE 1323

ENTRYPOINT [ ]
CMD [ "/ratp-go" ]