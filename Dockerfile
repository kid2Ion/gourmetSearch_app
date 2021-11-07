FROM golang:1.13 AS builder

WORKDIR /opt

COPY . .

RUN go build

ENTRYPOINT ["/opt/gourmetSearch_app"]

FROM ubuntu:20.04

COPY --from=builder /opt /opt

ENTRYPOINT ["/opt/gourmetSearch_app"]

