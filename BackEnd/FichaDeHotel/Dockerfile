FROM golang:1.20

ADD ./ /backend

WORKDIR /backend
RUN go mod init fichadehotel
RUN go mod tidy
RUN go build -o fichadehotel .
RUN go test ./...
RUN chmod +x /backend

ENTRYPOINT ["/backend/fichadehotel"]