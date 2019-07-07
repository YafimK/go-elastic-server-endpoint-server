FROM golang:1.12.6

ENV GO111MODULE=on

ADD . /
WORKDIR /
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o EndpointServer .
EXPOSE 8080

CMD ["/EndpointServer"]