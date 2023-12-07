FROM golang:1.20

ADD ./ /backend

WORKDIR /backend
RUN go mod init busquedadehotel
RUN go mod tidy
RUN go build -o busquedadehotel .
RUN go test ./...
RUN chmod +x /backend

ENTRYPOINT ["/backend/busquedadehotel"]