FROM golang:1.20

ADD ./ /backend

WORKDIR /backend
RUN go mod init urd
RUN go mod tidy
RUN go build -o urd .
RUN go test ./...
RUN chmod +x /backend

CMD sleep 10 && go run main.go

ENTRYPOINT ["/backend/urd"]