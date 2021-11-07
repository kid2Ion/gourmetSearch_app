FROM golang:1.13

WORKDIR /opt

COPY . .

RUN go build

EXPOSE 80

ENTRYPOINT ["/opt/gourmetSearch_app"]