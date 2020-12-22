FROM golang:1.15.0-alpine3.12 AS golang

RUN apk add --no-cache git
RUN mkdir -p /builds/go/src/github.com/betorvs/event-logger/
ENV GOPATH /go
COPY . /builds/go/src/github.com/betorvs/event-logger/
ENV CGO_ENABLED 0
RUN cd /builds/go/src/github.com/betorvs/event-logger/ && TESTRUN=true go test ./... && go build

FROM alpine:3.12
WORKDIR /
VOLUME /tmp
RUN apk add --no-cache ca-certificates
RUN update-ca-certificates
RUN mkdir -p /app
RUN addgroup -g 1000 -S app && \
    adduser -u 1000 -G app -S -D -h /app app && \
    chmod 755 /app
COPY --from=golang /builds/go/src/github.com/betorvs/event-logger/event-logger /app

EXPOSE 9090
RUN chmod +x /app/event-logger
WORKDIR /app    
USER app
CMD ["/app/event-logger"]
