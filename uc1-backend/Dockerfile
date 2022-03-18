FROM golang:1.17.5-alpine
WORKDIR /go/src/github.com/postgres-go
COPY . .
RUN go get -u github.com/lib/pq
RUN apk add git
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]

