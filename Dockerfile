FROM golang:1.12.6

ENV GO111MODULE=on

ADD ./app /go/src/app
WORKDIR /go/src/app
RUN go get
RUN go build -o EndpointServer
EXPOSE 8080
#ENV PORT=8080

CMD ["./EndpointServer"]