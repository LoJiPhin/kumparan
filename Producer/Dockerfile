FROM golang:latest

WORKDIR /go/src/app

COPY main.go main.go

RUN go get github.com/streadway/amqp
RUN go get github.com/LoJiPhin/kumparan/Contracts
RUN go get github.com/kataras/iris
RUN go get github.com/kataras/iris/middleware/logger
RUN go get github.com/kataras/iris/middleware/recover
RUN go get github.com/go-sql-driver/mysql

RUN go build -o main .
RUN ls 

CMD ./main