FROM golang:alpine
WORKDIR /go/src/gotator
COPY . .
RUN go build -o /go/bin/gotator 

ENTRYPOINT ["/go/bin/gotator"]