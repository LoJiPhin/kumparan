FROM golang:latest

RUN apt-get update
RUN apt-get install vim -y
RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/kataras/iris"
RUN go get "github.com/kataras/iris/middleware/logger"
RUN go get "github.com/kataras/iris/middleware/recover"
EXPOSE 3002



# #build stage
# FROM golang:alpine AS builder
# WORKDIR /go/src/app
# COPY . .
# RUN apk add --no-cache git
# RUN go get -d -v ./...
# RUN go install -v ./...

# #final stage
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# COPY --from=builder /go/bin/app /app
# ENTRYPOINT ./app
# LABEL Name=kumparan Version=0.0.1
# EXPOSE 3002

