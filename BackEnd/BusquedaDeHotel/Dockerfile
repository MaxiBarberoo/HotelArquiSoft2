FROM golang:1.20

ADD ./ /backend

WORKDIR /backend
RUN go mod init busquedadehotel
RUN go mod tidy
RUN go build -o busquedadehotel .
RUN chmod +x /backend

CMD sleep 10 && go run main.go

ENTRYPOINT ["/backend/busquedadehotel"]