# build container
FROM golang:latest as builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/2018-miraikeitai-org/Rakusale-Another-Server
COPY . .
RUN make

# runtime container
FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/2018-miraikeitai-org/Rakusale-Another-Server/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]