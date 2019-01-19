# Build Container
FROM golang:latest as builder
WORKDIR /go/src/github.com/2018-miraikeitai-org/Rakusale-Another-Server
COPY . .
# Set Environment Variable
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV DB_TYPE=mysql
ENV CONNECTION_STRING=rakusale:rakusale@tcp(127.0.0.1:3306)/rakusale?charset=utf8&parseTime=true
ENV PRIVATE_KEY=/go/src/github.com/2018-miraikeitai-org/Rakusale-Another-Server/secret/rakusale_private.key
ENV PUBLIC_KEY=/go/src/github.com/2018-miraikeitai-org/Rakusale-Another-Server/secret/rakusale_pub.key
# Run Command
RUN make

# Runtime Container
FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/2018-miraikeitai-org/Rakusale-Another-Server/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]