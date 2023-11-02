FROM golang:1.19

WORKDIR /usr/src/app

COPY . .
RUN go mod download
RUN go build -o server ./cmd/main.go

CMD [ "./server" ]