# Build Container
FROM golang:latest as builder
WORKDIR /go/src/github.com/2018-miraikeitai-org/Rakusale-Another-Server
COPY . .
# Set Environment Variable
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Run Command
RUN make

# Runtime Container
FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/2018-miraikeitai-org/Rakusale-Another-Server/app /app
COPY --from=builder /go/src/github.com/2018-miraikeitai-org/Rakusale-Another-Server/secret /secret
ENV DB_TYPE=mysql
ENV CONNECTION_STRING=rakusale:rakusale@tcp(db)/rakusale?charset=utf8&parseTime=true
#ENV CONNECTION_STRING=rakusale:rakusale@tcp(127.0.0.1:3306)/rakusale?charset=utf8&parseTime=true

# Rakusale Private & Public Key
ENV PRIVATE_KEY=/secret/rakusale_private.key
ENV PUBLIC_KEY=/secret/rakusale_pub.key
# Credential Google Cloud Storage
ENV GOOGLE_APPLICATION_CREDENTIALS="/secret/miraikeitai2018.json"
ENV GOOGLE_CLOUD_PROJECT="miraikeitai2018-rakusale"
ENV GOOGLE_CLOUD_STORAGE_BUCKET_NAME="miraikeitai2018-rakusale"
# PATH Info
ENV GOOGLE_CLOUD_STORAGE_PUBLIC_PATH="https://storage.googleapis.com/miraikeitai2018-rakusale/"
ENV SHOP_PATH="shops/" 
ENV VEGETABLE_PATH="vegetables/" 
ENV USER_PATH="users/" 
EXPOSE 3001
EXPOSE 3002
ENTRYPOINT ["/app"]