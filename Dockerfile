FROM golang:1.19.4-alpine3.17
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o main .
ENV PORT 8080
EXPOSE 8080
CMD ["./main"]
